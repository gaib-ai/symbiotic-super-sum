package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
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

type ChainConfig struct {
	Endpoint   string `json:"endpoint"`
	ReceiveUln string `json:"receiveUln"`
	DVN        string `json:"dvn"`
	Adapter    string `json:"adapter"`
	AID        string `json:"aid"`
}

type DeploymentConfig struct {
	ChainA ChainConfig `json:"chainA"`
	ChainB ChainConfig `json:"chainB"`
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
	var deployment DeploymentConfig
	err = json.Unmarshal(configFile, &deployment)
	if err != nil {
		log.Fatalf("Failed to parse deployment config: %v", err)
	}

	// Initialize clients and processor
	processor, err := setupProcessor(deployment)
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

func setupProcessor(config DeploymentConfig) (*dvn.Processor, error) {
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
		return nil, fmt.Errorf("failed to get private key: %w", err)
	}

	// Create authenticated transaction signer
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31338)) // Chain B ID
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}
	
	// Read deployment configuration again for contract addresses
	configFile, err := ioutil.ReadFile("temp-network/deploy-data/dvn_deployment.json")
	if err != nil {
		log.Fatalf("Failed to read deployment config: %v", err)
	}
	var deployment DeploymentConfig
	err = json.Unmarshal(configFile, &deployment)
	if err != nil {
		log.Fatalf("Failed to parse deployment config: %v", err)
	}

	dvnBAddr := common.HexToAddress(deployment.ChainB.DVN)
	dvnB, err := contracts.NewSymbioticLzDVN(dvnBAddr, clientB)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate DVN contract: %w", err)
	}

	receiveUlnBAddr := common.HexToAddress(deployment.ChainB.ReceiveUln)
	receiveUlnB, err := contracts.NewReceiveUln(receiveUlnBAddr, clientB)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate ReceiveUln contract: %w", err)
	}

	endpointAAddr := common.HexToAddress(deployment.ChainA.Endpoint)

	// Get relay address from environment
	relayAddr := os.Getenv("RELAY_ADDR")
	if relayAddr == "" {
		return nil, fmt.Errorf("RELAY_ADDR environment variable not set")
	}

	processor := dvn.NewProcessor(clientA, clientB, auth, dvnB, receiveUlnB, relayAddr)

	go func() {
		err := processor.ListenForPackets(context.Background(), endpointAAddr)
		if err != nil {
			log.Fatalf("Processor listener failed: %v", err)
		}
	}()

	return processor, nil
}
