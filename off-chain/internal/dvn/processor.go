package dvn

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"

	"dvn-node/internal/contracts"
)

// SymbioticProofData matches the structure expected by the DVN contract's decoder.
// struct SymbioticProofData {
//     uint48 epoch;
//     bytes proof;
// }
type SymbioticProofData struct {
	Epoch uint64 `json:"epoch"`
	Proof []byte `json:"proof"`
}

// SymbioticProofResponse matches the JSON structure returned by the relay sidecar.
type SymbioticProofResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Epoch uint64 `json:"epoch"`
		Proof string `json:"proof"` // Proof is hex-encoded string
	} `json:"result"`
}

// PacketHeader matches the packed structure of the LayerZero packet header.
// The packetHeader is a tightly packed byte array.
// We are not using ABI unpacking for performance and simplicity, just byte slicing.
// See: lib/LayerZero-v2/packages/layerzero-v2/evm/protocol/contracts/libraries/Packet.sol
type PacketHeader struct {
	Nonce    uint64
	SrcEid   uint32
	Sender   common.Address
	DstEid   uint32
	Receiver [32]byte
	Guid     [32]byte
}

func unpackPacketHeader(encodedHeader []byte) (*PacketHeader, error) {
	if len(encodedHeader) < 112 { // 8 + 4 + 32 + 4 + 32 + 32 = 112
		return nil, fmt.Errorf("invalid packet header length: got %d, want at least 112", len(encodedHeader))
	}

	header := &PacketHeader{}
	var offset int

	// Nonce (uint64)
	header.Nonce = binary.BigEndian.Uint64(encodedHeader[offset : offset+8])
	offset += 8

	// SrcEid (uint32)
	header.SrcEid = binary.BigEndian.Uint32(encodedHeader[offset : offset+4])
	offset += 4

	// Sender (address -> bytes32)
	// In packed format, address is padded to 32 bytes. We take the last 20 bytes.
	header.Sender = common.BytesToAddress(encodedHeader[offset+12 : offset+32])
	offset += 32

	// DstEid (uint32)
	header.DstEid = binary.BigEndian.Uint32(encodedHeader[offset : offset+4])
	offset += 4

	// Receiver (bytes32)
	copy(header.Receiver[:], encodedHeader[offset:offset+32])
	offset += 32

	// Guid (bytes32)
	copy(header.Guid[:], encodedHeader[offset:offset+32])

	return header, nil
}

type Processor struct {
	clientA     *ethclient.Client // Source chain
	clientB     *ethclient.Client // Destination chain
	auth        *bind.TransactOpts
	dvnB        *contracts.SymbioticLzDVN
	receiveUlnB *contracts.ReceiveUln
	relayAddr   string
}

func NewProcessor(clientA, clientB *ethclient.Client, auth *bind.TransactOpts, dvnB *contracts.SymbioticLzDVN, receiveUlnB *contracts.ReceiveUln, relayAddr string) *Processor {
	return &Processor{
		clientA:     clientA,
		clientB:     clientB,
		auth:        auth,
		dvnB:        dvnB,
		receiveUlnB: receiveUlnB,
		relayAddr:   relayAddr,
	}
}

func (p *Processor) ListenForPackets(ctx context.Context, endpointAddress common.Address) error {
	endpoint, err := contracts.NewEndpointV2(endpointAddress, p.clientA)
	if err != nil {
		return fmt.Errorf("failed to instantiate EndpointV2 contract: %w", err)
	}

	eventCh := make(chan *contracts.EndpointV2PacketSent)
	sub, err := endpoint.WatchPacketSent(&bind.WatchOpts{Context: ctx}, eventCh)
	if err != nil {
		return fmt.Errorf("failed to subscribe to PacketSent events: %w", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-sub.Err():
			return err
		case event := <-eventCh:
			header, err := unpackPacketHeader(event.PacketHeader)
			if err != nil {
				log.Printf("Error unpacking packet header, skipping packet: %v", err)
				continue
			}
			log.Printf("Received PacketSent event. From: %s, Nonce: %d", header.Sender.Hex(), header.Nonce)
			go p.processPacket(ctx, event, header)
		}
	}
}

func (p *Processor) processPacket(ctx context.Context, event *contracts.EndpointV2PacketSent, header *PacketHeader) {
	// 1. Construct the message hash for Symbiotic, matching the on-chain logic
	// On-chain logic: keccak256(abi.encodePacked(keccak256(abi.encodePacked(_packetHeader, _payloadHash))))
	payloadHash := sha256.Sum256(event.Payload)

	// Inner hash: keccak256(packetHeader, payloadHash)
	innerHashData := append(event.PacketHeader, payloadHash[:]...)
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(innerHashData)
	innerHash := hasher.Sum(nil)

	// Outer hash (the final message hash): keccak256(innerHash)
	hasher.Reset()
	hasher.Write(innerHash)
	messageHash := hasher.Sum(nil)
	var taskID [32]byte
	copy(taskID[:], messageHash)

	log.Printf("Processing task: %x", taskID)

	// 2. Get proof from Symbiotic sidecar
	symbioticProof, err := p.getSymbioticProof(taskID)
	if err != nil {
		log.Printf("Failed to get Symbiotic proof for task %x: %v", taskID, err)
		return
	}
	log.Printf("Successfully obtained proof for task %x from %s", taskID, p.relayAddr)

	// 3. Get ULN Config to find required confirmations
	ulnConfig, err := p.receiveUlnB.GetUlnConfig(&bind.CallOpts{Context: ctx}, common.BytesToAddress(header.Receiver[:]), header.SrcEid)
	if err != nil {
		log.Printf("Failed to get ULN config for task %x: %v", taskID, err)
		return
	}
	log.Printf("Required confirmations for this message: %d", ulnConfig.Confirmations)

	// 4. ABI-encode the proof data for the contract call
	encodedProof, err := p.abiEncodeProof(symbioticProof)
	if err != nil {
		log.Printf("Failed to ABI-encode symbiotic proof for task %x: %v", taskID, err)
		return
	}

	// 5. Submit verification to the destination chain
	tx, err := p.dvnB.VerifyWithSymbiotic(p.auth, event.PacketHeader, payloadHash, ulnConfig.Confirmations, encodedProof)
	if err != nil {
		log.Printf("Failed to submit verification for task %x: %v", taskID, err)
		return
	}

	log.Printf("Submitted verification for task %x. Transaction: %s", taskID, tx.Hash().Hex())

	// 6. Wait for transaction receipt
	receipt, err := bind.WaitMined(context.Background(), p.clientB, tx)
	if err != nil {
		log.Printf("Failed to get receipt for verification tx %s: %v", tx.Hash().Hex(), err)
		return
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Printf("Verification transaction %s failed", tx.Hash().Hex())
		return
	}

	log.Printf("Successfully verified task %x on destination chain.", taskID)
}

// getSymbioticProof interacts with the specific Symbiotic sidecar for this node.
func (p *Processor) getSymbioticProof(taskID [32]byte) (*SymbioticProofData, error) {
	proofURL := fmt.Sprintf("http://%s/proof/0x%s", p.relayAddr, hex.EncodeToString(taskID[:]))

	// The relay sidecar might take some time to aggregate the proof.
	// A production system would use a more robust backoff/retry strategy.
	const maxRetries = 5
	const retryDelay = 3 * time.Second

	var resp *http.Response
	var err error

	for i := 0; i < maxRetries; i++ {
		log.Printf("Requesting proof (attempt %d/%d) from: %s", i+1, maxRetries, proofURL)
		resp, err = http.Get(proofURL)
		if err == nil && resp.StatusCode == http.StatusOK {
			break
		}

		if resp != nil {
			resp.Body.Close()
		}
		if i < maxRetries-1 {
			time.Sleep(retryDelay)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get proof from sidecar after %d attempts: %w", maxRetries, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sidecar returned non-200 status: %s", resp.Status)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read proof from response body: %w", err)
	}

	var proofResponse SymbioticProofResponse
	err = json.Unmarshal(body, &proofResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON proof response: %w", err)
	}

	// The proof in the JSON response is a hex string (e.g., "0x..."), so we need to decode it.
	proofBytes, err := hex.DecodeString(proofResponse.Result.Proof[2:]) // Strip "0x" prefix
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex proof string: %w", err)
	}

	return &SymbioticProofData{
		Epoch: proofResponse.Result.Epoch,
		Proof: proofBytes,
	}, nil
}

func (p *Processor) abiEncodeProof(proofData *SymbioticProofData) ([]byte, error) {
	// The SymbioticLzDVN contract expects the proof data to be ABI-encoded.
	// struct SymbioticProofData { uint48 epoch; bytes proof; }
	// We manually construct this encoding.

	// uint48 epoch (encoded as a uint256, Solidity will handle the downcast)
	epochBytes := make([]byte, 32)
	binary.BigEndian.PutUint64(epochBytes[24:], proofData.Epoch) // Right-align the uint64 in a 32-byte slice

	// bytes proof (offset and length)
	offsetBytes := make([]byte, 32)
	binary.BigEndian.PutUint64(offsetBytes[24:], 64) // Offset to the data part is 64 bytes (2 * 32 bytes)

	lengthBytes := make([]byte, 32)
	binary.BigEndian.PutUint64(lengthBytes[24:], uint64(len(proofData.Proof)))

	// Actual proof data, padded to a multiple of 32 bytes
	paddedProofLen := (len(proofData.Proof) + 31) / 32 * 32
	paddedProof := make([]byte, paddedProofLen)
	copy(paddedProof, proofData.Proof)

	// Concatenate all parts
	encoded := append(epochBytes, offsetBytes...)
	encoded = append(encoded, lengthBytes...)
	encoded = append(encoded, paddedProof...)

	return encoded, nil
}
