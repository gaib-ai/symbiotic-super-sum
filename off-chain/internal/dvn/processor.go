package dvn

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"dvn-node/internal/contracts"
)

type Processor struct {
	clientA   *ethclient.Client // Source chain
	clientB   *ethclient.Client // Destination chain
	auth      *bind.TransactOpts
	dvnB      *contracts.SymbioticLzDVN
	relayAddr string
}

func NewProcessor(clientA, clientB *ethclient.Client, auth *bind.TransactOpts, dvnB *contracts.SymbioticLzDVN, relayAddr string) *Processor {
	return &Processor{
		clientA:   clientA,
		clientB:   clientB,
		auth:      auth,
		dvnB:      dvnB,
		relayAddr: relayAddr,
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
			log.Printf("Received PacketSent event. From: %x, Nonce: %d", event.PacketHeader.Sender, event.PacketHeader.Nonce)
			go p.processPacket(event)
		}
	}
}

func (p *Processor) processPacket(event *contracts.EndpointV2PacketSent) {
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


	// 3. Submit verification to the destination chain
	tx, err := p.dvnB.VerifyWithSymbiotic(p.auth, event.PacketHeader, payloadHash, event.PacketHeader.Confirmations, symbioticProof)
	if err != nil {
		log.Printf("Failed to submit verification for task %x: %v", taskID, err)
		return
	}

	log.Printf("Submitted verification for task %x. Transaction: %s", taskID, tx.Hash().Hex())

	// 4. Wait for transaction receipt
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
