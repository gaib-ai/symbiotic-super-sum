package dvn

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"dvn-node/internal/contracts"
)

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
	// 1. Construct the task for Symbiotic
	payloadHash := sha256.Sum256(event.Payload)
	taskData := append(event.PacketHeader, payloadHash[:]...)
	taskID := sha256.Sum256(taskData)

	log.Printf("Processing task: %x", taskID)

	// 2. Get proof from Symbiotic sidecar
	symbioticProof, err := p.getSymbioticProof(taskID)
	if err != nil {
		log.Printf("Failed to get Symbiotic proof for task %x: %v", taskID, err)
		return
	}
	log.Printf("Successfully obtained proof for task %x from %s", taskID, p.relayAddr)

	// 3. Get ULN Config to find required confirmations
	ulnConfig, err := p.receiveUlnB.GetUlnConfig(&bind.CallOpts{Context: ctx}, header.Sender, header.SrcEid)
	if err != nil {
		log.Printf("Failed to get ULN config for task %x: %v", taskID, err)
		return
	}
	log.Printf("Required confirmations for this message: %d", ulnConfig.Confirmations)


	// 4. Submit verification to the destination chain
	tx, err := p.dvnB.VerifyWithSymbiotic(p.auth, event.PacketHeader, payloadHash, ulnConfig.Confirmations, symbioticProof)
	if err != nil {
		log.Printf("Failed to submit verification for task %x: %v", taskID, err)
		return
	}

	log.Printf("Submitted verification for task %x. Transaction: %s", taskID, tx.Hash().Hex())

	// 5. Wait for transaction receipt
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
func (p *Processor) getSymbioticProof(taskID [32]byte) ([]byte, error) {
	proofURL := fmt.Sprintf("http://%s/proof/%s", p.relayAddr, hex.EncodeToString(taskID[:]))
	log.Printf("Requesting proof from: %s", proofURL)

	// In a real-world scenario, you would implement a robust retry mechanism.
	time.Sleep(5 * time.Second) // Simulate network delay and proof aggregation time

	// Mocking the HTTP call for now, as the sidecar proof endpoint might not be immediately available.
	// In a real implementation:
	// resp, err := http.Get(proofURL)
	// if err != nil {
	// 	 return nil, fmt.Errorf("failed to get proof from sidecar: %w", err)
	// }
	// defer resp.Body.Close()
	// proof, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to read proof from response: %w", err)
	// }
	// return proof, nil

	// Return a placeholder proof
	return []byte("mock-symbiotic-proof-from-" + p.relayAddr), nil
}
