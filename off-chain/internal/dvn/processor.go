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
	"github.com/ethereum/go-ethereum/rpc"
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
		Executor   string `json:"executor"`
	} `json:"chainA"`
	ChainB struct {
		Endpoint   string `json:"endpoint"`
		ReceiveUln string `json:"receiveUln"`
		Dvn        string `json:"dvn"`
		SendLib    string `json:"sendLib"`
		Executor   string `json:"executor"`
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

// PacketStatus defines the state of a packet in the processing pipeline.
type PacketStatus uint8

const (
	PacketStatus_None PacketStatus = iota
	PacketStatus_Detected
	PacketStatus_SignatureRequested
	PacketStatus_ProofReceived
	PacketStatus_Submitted
	PacketStatus_Failed // for terminal failure states
)

// PacketState holds the state and data for a packet being processed.
type PacketState struct {
	Packet         *Packet
	Status         PacketStatus
	SigEpoch       uint64
	SigRequestHash string
	AggProof       []byte
	TxHash         common.Hash // To store the submission tx hash
	LastError      string      // To store the last error message
	Retries        int         // To count retries for certain operations
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
	packets           map[common.Hash]*PacketState // Key: PayloadHash
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
		packets:           make(map[common.Hash]*PacketState),
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
			return ctx.Err()
		case <-ticker.C:
			// 1. Poll for new packets and add them to the state map sequentially
			for eid := range p.ethClients {
				if err := p.pollForPackets(ctx, eid); err != nil {
					// Log and continue to the next chain; a single chain failure shouldn't stop the whole worker.
					p.logger.Error("Error polling for packets", "eid", eid, "error", err)
				}
			}

			// 2. Process all packets based on their current state
			p.processPackets(ctx)
		}
	}
}

func (p *Processor) pollForPackets(ctx context.Context, eid uint32) error {
	client := p.ethClients[eid]
	// Get the latest finalized block number for reorg protection
	finalizedBlock, err := client.BlockByNumber(ctx, big.NewInt(rpc.FinalizedBlockNumber.Int64()))
	if err != nil {
		// Fallback to latest if finalized is not supported (e.g., some local testnets)
		finalizedBlock, err = client.BlockByNumber(ctx, nil)
		if err != nil {
			return errors.Errorf("failed to get latest/finalized block number: %w", err)
		}
	}
	latestBlockNumber := finalizedBlock.NumberU64()

	fromBlock := p.lastBlocks[eid]
	if fromBlock == 0 { // First run, start from the current block to avoid processing history
		fromBlock = latestBlockNumber
	}

	// Don't poll too many blocks at once, and ensure fromBlock is not ahead of toBlock
	toBlock := fromBlock + 49 // Poll a chunk of 50 blocks max
	if toBlock > latestBlockNumber {
		toBlock = latestBlockNumber
	}

	if fromBlock > toBlock {
		p.lastBlocks[eid] = latestBlockNumber // Still update last block to keep pace
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

		receipt, err := client.TransactionReceipt(ctx, event.Raw.TxHash)
		if err != nil {
			p.logger.Error("Failed to get receipt for PacketSent tx", "txHash", event.Raw.TxHash.Hex(), "error", err)
			continue
		}

		if !p.isDvnAssigned(receipt, eid) {
			p.logger.Debug("DVN not assigned for this packet, skipping", "txHash", event.Raw.TxHash.Hex())
			continue
		}

		p.logger.Info("DVN is assigned to packet, processing...", "txHash", event.Raw.TxHash.Hex())

		packet, err := p.unpackAndParsePacket(endpointABI, event)
		if err != nil {
			p.logger.Error("Error unpacking and parsing packet", "txHash", event.Raw.TxHash.Hex(), "error", err)
			continue
		}

		// Check if packet is already being processed
		if _, exists := p.packets[packet.PayloadHash]; exists {
			p.logger.Debug("Packet already detected, skipping", "payloadHash", hexutil.Encode(packet.PayloadHash[:]))
			continue
		}

		p.logger.Info("New packet detected, adding to processing queue", "payloadHash", hexutil.Encode(packet.PayloadHash[:]))
		p.packets[packet.PayloadHash] = &PacketState{
			Packet: packet,
			Status: PacketStatus_Detected,
		}
	}

	if err := iterator.Error(); err != nil {
		return errors.Errorf("iterator error: %w", err)
	}

	// Update last processed block
	p.lastBlocks[eid] = toBlock + 1

	return nil
}

// processPackets iterates through tracked packets and advances them through the state machine.
func (p *Processor) processPackets(ctx context.Context) {
	for payloadHash, state := range p.packets {
		var err error
		currentState := state

		l := p.logger.With("payloadHash", hexutil.Encode(payloadHash[:]), "status", currentState.Status)

		switch currentState.Status {
		case PacketStatus_Detected:
			l.Info("Processing detected packet")
			err = p.handleDetectedState(ctx, currentState)
		case PacketStatus_SignatureRequested:
			l.Info("Processing signature request")
			err = p.handleSignatureRequestedState(ctx, currentState)
		case PacketStatus_ProofReceived:
			l.Info("Processing received proof")
			err = p.handleSubmitProofState(ctx, currentState)
		case PacketStatus_Submitted:
			l.Info("Packet already submitted, cleaning up")
			delete(p.packets, payloadHash) // Clean up finished tasks
			continue
		case PacketStatus_Failed:
			l.Warn("Packet is in a failed state, ignoring", "reason", currentState.LastError)
			continue
		}

		if err != nil {
			l.Error("Error processing packet", "error", err)
			currentState.Retries++
			currentState.LastError = err.Error()
			if currentState.Retries > 5 {
				l.Error("Packet failed after multiple retries, moving to Failed state")
				currentState.Status = PacketStatus_Failed
			}
		}
	}
}

// handleDetectedState performs pre-flight checks and requests a signature from the relay.
func (p *Processor) handleDetectedState(ctx context.Context, state *PacketState) error {
	packet := state.Packet
	_, ok := p.ethClients[packet.DstEid]
	if !ok {
		return errors.Errorf("no client found for destination EID: %d", packet.DstEid)
	}

	// Check if this DVN has already verified this packet.
	isVerified, err := p.isAlreadyVerified(ctx, packet)
	if err != nil {
		p.logger.Warn("Failed to lookup packet hash, proceeding with verification attempt anyway", "error", err)
		// Don't return error, as the on-chain submission will correctly revert if already verified.
	}
	if isVerified {
		p.logger.Info("Packet already verified by this DVN, marking as submitted")
		state.Status = PacketStatus_Submitted
		return nil
	}

	p.logger.Info("Packet not yet verified, requesting signature from Symbiotic Relay")

	// The message sent to the relay for signing must match what the SymbioticLzDVN contract constructs.
	packetHeaderHash := crypto.Keccak256Hash(packet.PacketHeader)
	messageForRelay, err := p.constructMessageForRelay(packetHeaderHash, packet.PayloadHash)
	if err != nil {
		return errors.Errorf("failed to construct message for relay: %w", err)
	}

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

	p.logger.Info("Signature request submitted to relay", "requestHash", signResp.RequestHash, "epoch", signResp.Epoch)

	state.Status = PacketStatus_SignatureRequested
	state.SigRequestHash = signResp.RequestHash
	state.SigEpoch = signResp.Epoch
	state.Retries = 0 // Reset retries for the next stage
	state.LastError = ""

	return nil
}

// handleSignatureRequestedState polls the relay for the aggregation proof.
func (p *Processor) handleSignatureRequestedState(ctx context.Context, state *PacketState) error {
	proofResp, err := p.relayClient.GetAggregationProof(ctx, &v1.GetAggregationProofRequest{
		RequestHash: state.SigRequestHash,
	})
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			p.logger.Debug("Proof not ready yet, waiting...", "requestHash", state.SigRequestHash)
			return nil // Not an error, just not ready
		}
		return errors.Errorf("failed to get aggregation proof: %w", err)
	}

	p.logger.Info("<- Aggregated proof received from Symbiotic Relay", "epoch", state.SigEpoch)
	state.AggProof = proofResp.AggregationProof.Proof
	state.Status = PacketStatus_ProofReceived
	state.Retries = 0
	state.LastError = ""

	return nil
}

// handleSubmitProofState submits the verification proof to the destination chain.
func (p *Processor) handleSubmitProofState(ctx context.Context, state *PacketState) error {
	packet := state.Packet
	dstEid := packet.DstEid
	dvnContract, ok := p.dvnContracts[dstEid]
	if !ok {
		return errors.Errorf("no DVN contract found for destination EID: %d", dstEid)
	}

	// Get required confirmations from the destination ULN.
	requiredConfirmations, err := p.getRequiredConfirmations(ctx, packet)
	if err != nil {
		return errors.Errorf("failed to get required confirmations: %w", err)
	}
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

	p.logger.Info("Submitting verification transaction to destination chain", "dstEID", dstEid, "epoch", state.SigEpoch)

	tx, err := dvnContract.VerifyWithSymbiotic(
		txOpts,
		packet.PacketHeader,
		packet.PayloadHash,
		requiredConfirmations,
		new(big.Int).SetUint64(state.SigEpoch),
		state.AggProof,
	)
	if err != nil {
		if strings.Contains(err.Error(), "MessageAlreadyVerified") || strings.Contains(err.Error(), "ValidVerification") {
			p.logger.Warn("Transaction failed because message was already verified, which is expected in a race.")
			state.Status = PacketStatus_Submitted // Mark as done
			return nil
		}
		return errors.Errorf("failed to submit verification transaction: %w", err)
	}

	p.logger.Info("Verification transaction submitted successfully!", "txHash", tx.Hash().Hex())
	state.TxHash = tx.Hash()

	// Asynchronously wait for the transaction to be mined and update the final status.
	go func() {
		receipt, err := bind.WaitMined(context.Background(), p.ethClients[dstEid], tx)
		if err != nil {
			p.logger.Error("Failed to wait for transaction to be mined", "txHash", tx.Hash().Hex(), "error", err)
			state.LastError = "mining failed: " + err.Error() // Store error, but don't move to failed state yet.
			return
		}

		if receipt.Status == types.ReceiptStatusFailed {
			p.logger.Error("Verification transaction failed on-chain!", "txHash", tx.Hash().Hex())
			state.Status = PacketStatus_Failed
			state.LastError = "transaction reverted"
		} else {
			p.logger.Info("Verification transaction mined successfully!", "txHash", tx.Hash().Hex(), "blockNumber", receipt.BlockNumber)
			state.Status = PacketStatus_Submitted
			state.Retries = 0
			state.LastError = ""
		}
	}()

	// We optimistically move to Submitted state; the goroutine will handle the final outcome.
	// This prevents re-submission while waiting for mining.
	state.Status = PacketStatus_Submitted
	return nil
}

// --- Helper Functions ---

func (p *Processor) isDvnAssigned(receipt *types.Receipt, eid uint32) bool {
	myDvnAddress := p.dvnAddrs[eid]
	sendLibAddress := p.sendLibAddrs[eid]
	for _, log := range receipt.Logs {
		if log.Address == sendLibAddress && len(log.Topics) > 0 && log.Topics[0] == dvnFeePaidTopic {
			var feePaidEvent struct {
				RequiredDVNs []common.Address
				OptionalDVNs []common.Address
				Fees         []*big.Int
			}
			if err := sendLibABI.UnpackIntoInterface(&feePaidEvent, "DVNFeePaid", log.Data); err != nil {
				p.logger.Error("Failed to unpack DVNFeePaid event", "txHash", receipt.TxHash.Hex(), "error", err)
				continue
			}
			for _, dvn := range feePaidEvent.RequiredDVNs {
				if dvn == myDvnAddress {
					return true
				}
			}
		}
	}
	return false
}

func (p *Processor) unpackAndParsePacket(endpointABI abi.ABI, event *contracts.EndpointV2PacketSent) (*Packet, error) {
	results, err := endpointABI.Events["PacketSent"].Inputs.Unpack(event.Raw.Data)
	if err != nil {
		return nil, errors.Errorf("failed to unpack PacketSent event data: %w", err)
	}
	encodedPacket := results[0].([]byte)
	return parsePacket(encodedPacket)
}

func (p *Processor) isAlreadyVerified(ctx context.Context, packet *Packet) (bool, error) {
	dstEid := packet.DstEid
	client, ok := p.ethClients[dstEid]
	if !ok {
		return false, errors.Errorf("no client found for destination EID: %d", dstEid)
	}

	receiveUlnAddrString, err := p.config.GetReceiveUln(dstEid)
	if err != nil {
		return false, err
	}
	receiveUlnAddr := common.HexToAddress(receiveUlnAddrString)
	receiveUlnContract, err := contracts.NewReceiveUln302(receiveUlnAddr, client)
	if err != nil {
		return false, errors.Errorf("failed to instantiate ReceiveUln302 contract: %w", err)
	}

	packetHash := crypto.Keccak256Hash(packet.PacketHeader)
	dvnAddrOnDest := p.dvnAddrs[dstEid]

	lookupResult, err := receiveUlnContract.HashLookup(&bind.CallOpts{Context: ctx}, packetHash, packet.PayloadHash, dvnAddrOnDest)
	if err != nil {
		return false, errors.Errorf("failed to call hashLookup: %w", err)
	}

	return lookupResult.Submitted, nil
}

func (p *Processor) constructMessageForRelay(packetHeaderHash, payloadHash common.Hash) ([]byte, error) {
	bytes32T, _ := abi.NewType("bytes32", "", nil)
	args := abi.Arguments{
		{Type: bytes32T},
		{Type: bytes32T},
	}
	return args.Pack(packetHeaderHash, payloadHash)
}

func (p *Processor) getRequiredConfirmations(ctx context.Context, packet *Packet) (uint64, error) {
	dstEid := packet.DstEid
	client, ok := p.ethClients[dstEid]
	if !ok {
		return 0, errors.Errorf("no client found for destination EID: %d", dstEid)
	}
	receiveUlnAddrString, err := p.config.GetReceiveUln(dstEid)
	if err != nil {
		return 0, err
	}
	receiveUlnAddr := common.HexToAddress(receiveUlnAddrString)
	receiveUlnContract, err := contracts.NewReceiveUln302(receiveUlnAddr, client)
	if err != nil {
		return 0, errors.Errorf("failed to instantiate ReceiveUln302 contract: %w", err)
	}

	receiverAddress := common.BytesToAddress(packet.Receiver.Bytes())
	ulnConfig, err := receiveUlnContract.GetUlnConfig(&bind.CallOpts{Context: ctx}, receiverAddress, packet.SrcEid)
	if err != nil {
		return 0, errors.Errorf("failed to get ULN config for oapp %s: %w", receiverAddress.Hex(), err)
	}
	return ulnConfig.Confirmations, nil
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
