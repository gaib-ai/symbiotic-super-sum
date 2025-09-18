package dvn

import (
	"context"
	"crypto/sha256"
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

	"dvn-worker/internal/contracts"
)

type Processor struct {
	clientA *ethclient.Client // Source chain
	clientB *ethclient.Client // Destination chain
	auth    *bind.TransactOpts
	dvnB    *contracts.SymbioticLzDVN
}

func NewProcessor(clientA, clientB *ethclient.Client, auth *bind.TransactOpts, dvnB *contracts.SymbioticLzDVN) *Processor {
	return &Processor{
		clientA: clientA,
		clientB: clientB,
		auth:    auth,
		dvnB:    dvnB,
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
	// In a real implementation, this would involve an HTTP call to the sidecar.
	// For this example, we'll mock it and assume we get a valid proof.
	// We'll use a placeholder proof.
	symbioticProof, err := p.getSymbioticProof(taskID)
	if err != nil {
		log.Printf("Failed to get Symbiotic proof for task %x: %v", taskID, err)
		return
	}
	log.Printf("Successfully obtained proof for task %x", taskID)


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

// getSymbioticProof mocks the interaction with a Symbiotic sidecar.
func (p *Processor) getSymbioticProof(taskID [32]byte) ([]byte, error) {
	// In a real-world scenario, you would make an HTTP request to the sidecar API.
	// e.g., resp, err := http.Get("http://relay-sidecar-1:8081/proof/" + hex.EncodeToString(taskID[:]))
	// Here we will just wait for a moment and return a dummy proof.
	log.Printf("Requesting proof for task %x from sidecar...", taskID)
	time.Sleep(5 * time.Second) // Simulate network delay

	// Return a placeholder proof
	return []byte("mock-symbiotic-proof"), nil
}
