package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"dvn-node/internal/contracts"
	"dvn-node/internal/dvn"
)

type DvnDeployment struct {
	ChainA struct {
		Endpoint string `json:"endpoint"`
		Dvn      string `json:"dvn"`
	} `json:"chainA"`
	ChainB struct {
		Endpoint string `json:"endpoint"`
		Dvn      string `json:"dvn"`
	} `json:"chainB"`
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	// Read deployment configuration
	configFile, err := ioutil.ReadFile("temp-network/deploy-data/dvn_deployment.json")
	if err != nil {
		log.Fatalf("Failed to read deployment config: %v", err)
	}
	var deployment DvnDeployment
	err = json.Unmarshal(configFile, &deployment)
	if err != nil {
		log.Fatalf("Failed to parse deployment config: %v", err)
	}

	// Initialize clients and processor
	processor, err := setupProcessor()
	if err != nil {
		log.Fatalf("Failed to setup processor: %v", err)
	}

	// Create a context that is cancelled on interruption
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the DVN listener
	go func() {
		err := processor.ListenForPackets(ctx, common.HexToAddress(deployment.ChainA.Endpoint))
		if err != nil {
			log.Printf("Listener exited with error: %v", err)
			cancel() // Stop main if listener fails
		}
	}()
	
	log.Println("DVN Worker started. Listening for PacketSent events...")

	// Wait for termination signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	log.Println("Shutting down DVN Worker...")
}

func setupProcessor() (*dvn.Processor, error) {
	// Setup for Chain A (Source)
	clientA, err := ethclient.Dial(os.Getenv("RPC_URL_A"))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to source chain: %w", err)
	}

	// Setup for Chain B (Destination)
	clientB, err := ethclient.Dial(os.Getenv("RPC_URL_B"))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to destination chain: %w", err)
	}

	// Get private key for signing transactions on the destination chain
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: %w", err)
	}

	// Create authenticated transaction signer
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31338)) // Assuming Chain B ID is 31338
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}
	
	// Read deployment configuration again for contract addresses
	configFile, err := ioutil.ReadFile("temp-network/deploy-data/dvn_deployment.json")
	if err != nil {
		log.Fatalf("Failed to read deployment config: %v", err)
	}
	var deployment DvnDeployment
	err = json.Unmarshal(configFile, &deployment)
	if err != nil {
		log.Fatalf("Failed to parse deployment config: %v", err)
	}

	dvnBAddress := common.HexToAddress(deployment.ChainB.Dvn)
	dvnB, err := contracts.NewSymbioticLzDVN(dvnBAddress, clientB)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate DVN contract: %w", err)
	}

	// Get relay address from environment
	relayAddr := os.Getenv("RELAY_ADDR")
	if relayAddr == "" {
		return nil, fmt.Errorf("RELAY_ADDR environment variable not set")
	}

	return dvn.NewProcessor(clientA, clientB, auth, dvnB, relayAddr), nil
}
