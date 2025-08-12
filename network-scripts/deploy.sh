#!/bin/sh
set -e

echo "Waiting for anvil to be ready..."
until cast client --rpc-url http://anvil:8545 > /dev/null 2>&1; do sleep 1; done

echo "Deploying contracts..."
forge script script/LocalDeploy.s.sol:LocalDeploy --rpc-url http://anvil:8545 -vv --broadcast --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 | tee /deploy-data/deployment.log

# Extract contract addresses and save to shared files
DRIVER_ADDRESS=$(grep "Driver (chainId: 31337 ):" /deploy-data/deployment.log | awk '{print $5}')
SUMTASK_ADDRESS=$(grep "SumTask (chainId: 31337 ):" /deploy-data/deployment.log | awk '{print $5}')
echo "$DRIVER_ADDRESS" > /deploy-data/driver_address.txt
echo "$SUMTASK_ADDRESS" > /deploy-data/sumtask_address.txt
echo "Driver address saved: $DRIVER_ADDRESS"
echo "SumTask address saved: $SUMTASK_ADDRESS"

echo "Setting interval mining..."
cast rpc --rpc-url http://anvil:8545 evm_setIntervalMining 1

echo "Mine a single block to finalize the deployment..."
cast rpc --rpc-url http://anvil:8545 evm_mine

echo "Deployment completed successfully!"

# Create deployment completion marker
echo "$(date): Deployment completed successfully" > /deploy-data/deployment-complete.marker
echo "Deployment completion marker created" 
