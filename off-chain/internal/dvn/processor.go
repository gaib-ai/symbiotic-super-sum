package dvn

import (
	"context"
	"encoding/json"
	"log/slog"
	"math/big"
	"os"
	"strings"
	"time"

	contracts "dvn-node/internal/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-errors/errors"
	v1 "github.com/symbioticfi/relay/api/client/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Use a fixed, absolute path inside the container
const deploymentConfigPath = "/app/temp-network/deploy-data/dvn_deployment.json"

var (
	dvnFeePaidTopic = crypto.Keccak256Hash([]byte("DVNFeePaid(address[],address[],uint256[])"))
	sendLibABI      abi.ABI
)

func init() {
	// A simplified ABI for just the DVNFeePaid event we need to check.
	const abiJSON = `[{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address[]","name":"requiredDVNs","type":"address[]"},{"indexed":false,"internalType":"address[]","name":"optionalDVNs","type":"address[]"},{"indexed":false,"internalType":"uint256[]","name":"fees","type":"uint256[]"}],"name":"DVNFeePaid","type":"event"}]`
	var err error
	sendLibABI, err = abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		// This should not happen with a static correct ABI
		panic("failed to parse SendLib ABI: " + err.Error())
	}
}

type DvnConfig struct {
	ChainA struct {
		Endpoint   string `json:"endpoint"`
		ReceiveUln string `json:"receiveUln"`
		Dvn        string `json:"dvn"`
		SendLib    string `json:"sendLib"`
	} `json:"chainA"`
	ChainB struct {
		Endpoint   string `json:"endpoint"`
		ReceiveUln string `json:"receiveUln"`
		Dvn        string `json:"dvn"`
		SendLib    string `json:"sendLib"`
	} `json:"chainB"`
}

func (c *DvnConfig) GetReceiveUln(eid uint32) (string, error) {
	if eid == 31337 {
		return c.ChainA.ReceiveUln, nil
	}
	if eid == 31338 {
		return c.ChainB.ReceiveUln, nil
	}
	return "", errors.Errorf("unknown eid for receive uln: %d", eid)
}

func (c *DvnConfig) GetSendLib(eid uint32) (string, error) {
	if eid == 31337 {
		return c.ChainA.SendLib, nil
	}
	if eid == 31338 {
		return c.ChainB.SendLib, nil
	}
	return "", errors.Errorf("unknown eid for send lib: %d", eid)
}

// Packet struct mirrors the structure in LayerZero's ISendLib.sol
type Packet struct {
	Nonce        uint64
	SrcEid       uint32
	Sender       common.Address
	DstEid       uint32
	Receiver     common.Hash
	Guid         common.Hash
	Message      []byte
	PacketHeader []byte
	PayloadHash  common.Hash
}

type Processor struct {
	relayClient       *v1.SymbioticClient
	ethClients        map[uint32]*ethclient.Client
	endpointContracts map[uint32]*contracts.EndpointV2
	dvnContracts      map[uint32]*contracts.SymbioticLzDVN
	dvnAddrs          map[uint32]common.Address
	sendLibAddrs      map[uint32]common.Address
	privateKey        string
	logger            *slog.Logger
	config            *DvnConfig
	lastBlocks        map[uint32]uint64
}

func NewProcessor(
	ctx context.Context,
	relayApiURL string,
	evmRpcURLs map[uint32]string,
	privateKey string,
	logger *slog.Logger,
) (*Processor, error) {
	// 1. Read the deployment config file
	configFile, err := os.ReadFile(deploymentConfigPath)
	if err != nil {
		return nil, errors.Errorf("failed to read deployment config: %w", err)
	}
	var config DvnConfig
	if err := json.Unmarshal(configFile, &config); err != nil {
		return nil, errors.Errorf("failed to parse deployment config: %w", err)
	}

	// 2. Connect to Relay
	conn, err := grpc.Dial(relayApiURL, grpc.WithTransportCredentials(insecure.NewCredentials())) // Use appropriate security options in production
	if err != nil {
		return nil, errors.Errorf("failed to connect to relay API: %w", err)
	}
	relayClient := v1.NewSymbioticClient(conn)

	// 3. Connect to ETH clients and instantiate contracts
	ethClients := make(map[uint32]*ethclient.Client)
	endpointContracts := make(map[uint32]*contracts.EndpointV2)
	dvnContracts := make(map[uint32]*contracts.SymbioticLzDVN)
	dvnAddrs := make(map[uint32]common.Address)
	sendLibAddrs := make(map[uint32]common.Address)
	lastBlocks := make(map[uint32]uint64)

	for eid, rpcURL := range evmRpcURLs {
		client, err := ethclient.DialContext(ctx, rpcURL)
		if err != nil {
			return nil, errors.Errorf("failed to connect to RPC URL '%s' for EID %d: %w", rpcURL, eid, err)
		}
		ethClients[eid] = client

		var endpointAddr, dvnAddr, sendLibAddr common.Address
		if eid == 31337 { // Chain A
			endpointAddr = common.HexToAddress(config.ChainA.Endpoint)
			dvnAddr = common.HexToAddress(config.ChainA.Dvn)
			sendLibAddr = common.HexToAddress(config.ChainA.SendLib)
		} else if eid == 31338 { // Chain B
			endpointAddr = common.HexToAddress(config.ChainB.Endpoint)
			dvnAddr = common.HexToAddress(config.ChainB.Dvn)
			sendLibAddr = common.HexToAddress(config.ChainB.SendLib)
		} else {
			return nil, errors.Errorf("unsupported EID: %d", eid)
		}

		endpoint, err := contracts.NewEndpointV2(endpointAddr, client)
		if err != nil {
			return nil, errors.Errorf("failed to instantiate EndpointV2 contract for EID %d: %w", eid, err)
		}
		endpointContracts[eid] = endpoint

		dvn, err := contracts.NewSymbioticLzDVN(dvnAddr, client)
		if err != nil {
			return nil, errors.Errorf("failed to instantiate SymbioticLzDVN contract for EID %d: %w", eid, err)
		}
		dvnContracts[eid] = dvn
		dvnAddrs[eid] = dvnAddr
		sendLibAddrs[eid] = sendLibAddr
		lastBlocks[eid] = 0 // Initialize last block tracker
	}

	return &Processor{
		relayClient:       relayClient,
		ethClients:        ethClients,
		endpointContracts: endpointContracts,
		dvnContracts:      dvnContracts,
		dvnAddrs:          dvnAddrs,
		sendLibAddrs:      sendLibAddrs,
		privateKey:        privateKey,
		logger:            logger,
		config:            &config,
		lastBlocks:        lastBlocks,
	}, nil
}

// Listen starts the main event processing loop.
func (p *Processor) Listen(ctx context.Context) error {
	p.logger.Info("DVN Worker started. Polling for PacketSent events...")

	ticker := time.NewTicker(2 * time.Second) // Poll every 2 seconds
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			p.logger.Info("Listener shutting down.")
			return nil
		case <-ticker.C:
			for eid, client := range p.ethClients {
				if err := p.pollForPackets(ctx, eid, client); err != nil {
					p.logger.Error("Error polling for packets", "eid", eid, "error", err)
					// In a production system, you might want more robust error handling here.
				}
			}
		}
	}
}

func (p *Processor) pollForPackets(ctx context.Context, eid uint32, client *ethclient.Client) error {
	// Get the latest finalized block number
	latestBlock, err := client.BlockByNumber(ctx, nil) // nil for latest
	if err != nil {
		return errors.Errorf("failed to get latest block number: %w", err)
	}
	latestBlockNumber := latestBlock.NumberU64()

	fromBlock := p.lastBlocks[eid]
	if fromBlock == 0 { // First run, start from the current block
		fromBlock = latestBlockNumber
	}

	// Don't poll too many blocks at once, and ensure fromBlock is not ahead of toBlock
	toBlock := fromBlock + 49 // Poll a chunk of 50 blocks max
	if toBlock > latestBlockNumber {
		toBlock = latestBlockNumber
	}

	if fromBlock > toBlock {
		// Nothing to process yet
		p.lastBlocks[eid] = latestBlockNumber
		return nil
	}

	p.logger.Debug("Polling for events", "eid", eid, "fromBlock", fromBlock, "toBlock", toBlock)

	endpoint := p.endpointContracts[eid]
	iterator, err := endpoint.FilterPacketSent(&bind.FilterOpts{
		Start:   fromBlock,
		End:     &toBlock,
		Context: ctx,
	})

	if err != nil {
		return errors.Errorf("failed to filter PacketSent events: %w", err)
	}
	defer iterator.Close()

	endpointABI, err := abi.JSON(strings.NewReader(contracts.EndpointV2ABI))
	if err != nil {
		return errors.Errorf("failed to parse endpoint ABI: %w", err)
	}

	for iterator.Next() {
		event := iterator.Event
		p.logger.Debug("Found PacketSent event, checking if assigned", "txHash", event.Raw.TxHash.Hex())

		// Check if this DVN was assigned and paid in the same transaction.
		receipt, err := client.TransactionReceipt(ctx, event.Raw.TxHash)
		if err != nil {
			p.logger.Error("Failed to get receipt for PacketSent tx", "txHash", event.Raw.TxHash.Hex(), "error", err)
			continue
		}

		myDvnAddress := p.dvnAddrs[eid]
		sendLibAddress := p.sendLibAddrs[eid]
		assigned := false
		for _, log := range receipt.Logs {
			if log.Address == sendLibAddress && len(log.Topics) > 0 && log.Topics[0] == dvnFeePaidTopic {
				var feePaidEvent struct {
					RequiredDVNs []common.Address
					OptionalDVNs []common.Address
					Fees         []*big.Int
				}
				// The actual event data is in log.Data, and topics are for indexed fields (none here)
				if err := sendLibABI.UnpackIntoInterface(&feePaidEvent, "DVNFeePaid", log.Data); err != nil {
					p.logger.Error("Failed to unpack DVNFeePaid event", "txHash", event.Raw.TxHash.Hex(), "error", err)
					continue // continue to next log
				}
				for _, dvn := range feePaidEvent.RequiredDVNs {
					if dvn == myDvnAddress {
						assigned = true
						break
					}
				}
				if assigned {
					break
				}
			}
		}

		if !assigned {
			p.logger.Debug("DVN not assigned for this packet, skipping", "txHash", event.Raw.TxHash.Hex(), "dvn_address", myDvnAddress.Hex())
			continue
		}

		p.logger.Info("DVN is assigned to packet, processing...", "txHash", event.Raw.TxHash.Hex())

		// Process each event in a separate goroutine to avoid blocking the polling loop.
		go func(e *contracts.EndpointV2PacketSent) {
			p.logger.Info("-> Detected PacketSent event", "srcEID", eid, "txHash", e.Raw.TxHash.Hex())

			// The PacketSent event in LZv2 doesn't have named fields in its data.
			// It's an ABI-encoded tuple of (bytes, bytes, address). We need the first element.
			results, err := endpointABI.Events["PacketSent"].Inputs.Unpack(e.Raw.Data)
			if err != nil {
				p.logger.Error("Error unpacking PacketSent event data", "txHash", e.Raw.TxHash.Hex(), "error", err)
				return
			}
			encodedPacket := results[0].([]byte)

			// The PacketSent event contains the full encoded packet (header + message).
			// We need to parse it to extract the components for verification.
			packet, err := parsePacket(encodedPacket)
			if err != nil {
				p.logger.Error("Error parsing packet", "txHash", e.Raw.TxHash.Hex(), "error", err)
				return
			}

			if err := p.handlePacket(context.Background(), packet); err != nil {
				p.logger.Error("Error handling packet", "payloadHash", hexutil.Encode(packet.PayloadHash[:]), "error", err)
			}
		}(event)
	}

	if err := iterator.Error(); err != nil {
		return errors.Errorf("iterator error: %w", err)
	}

	// Update last processed block
	p.lastBlocks[eid] = toBlock + 1

	return nil
}

func (p *Processor) handlePacket(ctx context.Context, packet *Packet) error {
	// --- 1. Destination chain setup and pre-flight checks ---
	dstEid := packet.DstEid
	dvnContract, ok := p.dvnContracts[dstEid]
	if !ok {
		return errors.Errorf("no DVN contract found for destination EID: %d", dstEid)
	}
	client := p.ethClients[dstEid]

	receiveUlnAddrString, err := p.config.GetReceiveUln(dstEid)
	if err != nil {
		return err
	}
	receiveUlnAddr := common.HexToAddress(receiveUlnAddrString)
	receiveUlnContract, err := contracts.NewReceiveUln302(receiveUlnAddr, client)
	if err != nil {
		return errors.Errorf("failed to instantiate ReceiveUln302 contract: %w", err)
	}

	// --- 2. Check if this DVN has already verified this packet (hashLookup) ---
	packetHash := crypto.Keccak256Hash(packet.PacketHeader)
	dvnAddrOnDest := p.dvnAddrs[dstEid] // The address of our SymbioticLzDVN contract on the destination chain.

	// The hashLookup function is part of the ReceiveUln302 standard. We expect our binding to have it.
	lookupResult, err := receiveUlnContract.HashLookup(&bind.CallOpts{Context: ctx}, packetHash, packet.PayloadHash, dvnAddrOnDest)
	if err != nil {
		// We log the error but proceed. A failure to lookup shouldn't block verification,
		// as the on-chain submission will correctly revert if already verified.
		p.logger.Warn("Failed to lookup packet hash, proceeding with verification attempt anyway", "payloadHash", hexutil.Encode(packet.PayloadHash[:]), "error", err)
	} else if lookupResult.Submitted {
		p.logger.Info("Packet already verified by this DVN, skipping", "payloadHash", hexutil.Encode(packet.PayloadHash[:]))
		return nil // Success, nothing more to do.
	}
	p.logger.Info("Packet not yet verified by this DVN, proceeding...", "payloadHash", hexutil.Encode(packet.PayloadHash[:]))

	// --- 3. Request signature from Symbiotic Relay ---
	// The message sent to the relay for signing must match what the SymbioticLzDVN contract
	// constructs for verification. Based on the LayerZero DVN standard, this should be
	// the ABI encoding of the 81-byte packet header and the 32-byte payload hash.
	bytesT, _ := abi.NewType("bytes", "", nil)
	bytes32T, _ := abi.NewType("bytes32", "", nil)
	args := abi.Arguments{
		{Type: bytesT},
		{Type: bytes32T},
	}
	messageForRelay, err := args.Pack(packet.PacketHeader, packet.PayloadHash)
	if err != nil {
		return errors.Errorf("failed to pack header and payload hash for relay: %w", err)
	}

	p.logger.Info("Requesting signature from Symbiotic Relay",
		"payloadHash", hexutil.Encode(packet.PayloadHash[:]),
		"messageForRelay", hexutil.Encode(messageForRelay))

	// Request the signature from the Relay using the raw packed data.
	suggestedEpoch, err := p.relayClient.GetSuggestedEpoch(ctx, &v1.GetSuggestedEpochRequest{})
	if err != nil {
		return errors.Errorf("failed to get suggested epoch from relay: %w", err)
	}
	signResp, err := p.relayClient.SignMessage(ctx, &v1.SignMessageRequest{
		KeyTag:        15, // Default BLS key tag
		Message:       messageForRelay,
		RequiredEpoch: &suggestedEpoch.Epoch,
	})
	if err != nil {
		return errors.Errorf("failed to request signature from relay: %w", err)
	}
	p.logger.Info("Signature request submitted", "requestHash_from_relay", signResp.RequestHash, "epoch", signResp.Epoch)

	// --- 4. Poll the relay for the aggregated proof ---
	var aggProof *v1.AggregationProof
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(2 * time.Second): // Poll every 2 seconds
			proofResp, err := p.relayClient.GetAggregationProof(ctx, &v1.GetAggregationProofRequest{
				RequestHash: signResp.RequestHash,
			})
			if err != nil {
				// Not ready yet, wait and retry.
				if strings.Contains(err.Error(), "not found") {
					p.logger.Debug("Proof not ready yet, waiting...", "requestHash", signResp.RequestHash)
					continue
				}
				return errors.Errorf("failed to get aggregation proof: %w", err)
			}
			aggProof = proofResp.AggregationProof
			p.logger.Info("<- Aggregated proof received from Symbiotic Relay", "epoch", signResp.Epoch, "proof_len", len(aggProof.Proof))
			goto SUBMIT_PROOF
		}
	}

SUBMIT_PROOF:
	// --- 5. Submit the proof to the SymbioticLzDVN contract on the destination chain ---
	// The config is keyed by the receiver OApp address on the destination chain and the source EID.
	receiverAddress := common.BytesToAddress(packet.Receiver.Bytes())
	ulnConfig, err := receiveUlnContract.GetUlnConfig(&bind.CallOpts{Context: ctx}, receiverAddress, packet.SrcEid)
	if err != nil {
		return errors.Errorf("failed to get ULN config from destination chain for oapp %s: %w", receiverAddress.Hex(), err)
	}
	requiredConfirmations := ulnConfig.Confirmations
	p.logger.Info("Got required confirmations from on-chain config", "confirmations", requiredConfirmations)

	pk, err := crypto.HexToECDSA(p.privateKey)
	if err != nil {
		return errors.Errorf("failed to parse private key: %w", err)
	}

	chainID, err := p.ethClients[dstEid].ChainID(ctx)
	if err != nil {
		return errors.Errorf("failed to get chain ID for EID %d: %w", dstEid, err)
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		return errors.Errorf("failed to create transactor: %w", err)
	}

	p.logger.Info("Submitting verification transaction to destination chain", "dstEID", dstEid, "epoch", signResp.Epoch)

	// We now have all the correct components to call the verification function.
	symbioticProof := abiEncodeSymbioticProof(signResp.Epoch, aggProof.Proof)

	tx, err := dvnContract.VerifyWithSymbiotic(
		txOpts,
		packet.PacketHeader,
		packet.PayloadHash,
		requiredConfirmations, // Use the dynamically fetched confirmations
		symbioticProof,
	)

	if err != nil {
		// It's possible another node submitted the transaction first.
		if strings.Contains(err.Error(), "MessageAlreadyVerified") || strings.Contains(err.Error(), "ValidVerification") {
			p.logger.Warn("Transaction failed because message was already verified, which is expected in a race.", "payloadHash", hexutil.Encode(packet.PayloadHash[:]))
			return nil
		}
		return errors.Errorf("failed to submit verification transaction: %w", err)
	}

	p.logger.Info("Verification transaction submitted successfully!", "txHash", tx.Hash().Hex())

	receipt, err := bind.WaitMined(ctx, p.ethClients[dstEid], tx)
	if err != nil {
		return errors.Errorf("failed to wait for transaction to be mined: %w", err)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		p.logger.Error("Verification transaction failed!", "txHash", tx.Hash().Hex())
	} else {
		p.logger.Info("Verification transaction mined successfully!", "txHash", tx.Hash().Hex(), "blockNumber", receipt.BlockNumber)
	}

	return nil
}

// parsePacket decodes the raw encoded packet from the PacketSent event.
func parsePacket(encodedPacket []byte) (*Packet, error) {
	if len(encodedPacket) < 113 { // Minimum length: 81 (header) + 32 (guid)
		return nil, errors.New("encoded packet is too short")
	}

	p := &Packet{}
	// Nonce starts at offset 1 (after version byte)
	p.Nonce = new(big.Int).SetBytes(encodedPacket[1:9]).Uint64()
	p.SrcEid = uint32(new(big.Int).SetBytes(encodedPacket[9:13]).Uint64())
	// Sender is a 32-byte field in the packet, but the address is the last 20 bytes.
	p.Sender = common.BytesToAddress(encodedPacket[25:45]) // Corrected offset from 13 to 25
	p.DstEid = uint32(new(big.Int).SetBytes(encodedPacket[45:49]).Uint64())
	// Receiver is a full 32-byte hash.
	p.Receiver = common.BytesToHash(encodedPacket[49:81])
	// GUID is the next 32 bytes after the 81-byte header.
	p.Guid = common.BytesToHash(encodedPacket[81:113])
	p.Message = encodedPacket[113:]

	// The PacketHeader must match the logic in LayerZero's reference DVN scripts.
	// It is a fixed-size 81-byte slice of the encoded packet, from version to receiver.
	p.PacketHeader = encodedPacket[0:81]

	// The PayloadHash must also match the reference scripts: keccak256(abi.encodePacked(guid, message))
	// In Go, this is equivalent to concatenating the byte slices and then hashing.
	packedPayload := append(p.Guid.Bytes(), p.Message...)
	p.PayloadHash = crypto.Keccak256Hash(packedPayload)

	return p, nil
}

func abiEncodeSymbioticProof(epoch uint64, proof []byte) []byte {
	// This function ABI-encodes the epoch and proof to match what the SymbioticLzDVN.verifyWithSymbiotic expects.
	uint48T, _ := abi.NewType("uint48", "", nil)
	bytesT, _ := abi.NewType("bytes", "", nil)

	args := abi.Arguments{
		{Type: uint48T}, // epoch
		{Type: bytesT},  // proof
	}

	// The go-ethereum ABI packer expects a *big.Int for uint types that are not native Go sizes (e.g., uint48).
	encoded, err := args.Pack(new(big.Int).SetUint64(epoch), proof)
	if err != nil {
		// This should not happen with correct types
		panic(err)
	}
	return encoded
}
