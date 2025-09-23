#!/bin/sh
set -e

echo "Starting DVN Node..."

# Wait for the deployment configuration file to be available.
# This file is created by the 'deployer' service and is essential for the node to know contract addresses.
echo "Waiting for dvn_deployment.json file..."
until [ -f /app/temp-network/deploy-data/dvn_deployment.json ]; do
    echo "Config file not found, sleeping..."
    sleep 2
done
echo "Found dvn_deployment.json file. Starting node."

# Arguments from docker-compose command
RELAY_ADDR=${1:-"relay-sidecar-1:8080"}
PRIVATE_KEY=${2:-"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"} # Default Anvil key

# RPC URLs are now sourced from the container's environment (from .env file in temp-network)
RPC_URL_A=${RPC_URL_A:-"http://anvil:8545"}
RPC_URL_B=${RPC_URL_B:-"http://anvil-settlement:8546"}

# Construct the EVM RPC URL flag string in the format "eid:url,eid:url"
EVM_RPC_URLS="31337:${RPC_URL_A},31338:${RPC_URL_B}"

# Execute the main application with flags
# Use --log-level=debug for more verbose output during testing
/app/dvn-node \
    --relay-api-url="${RELAY_ADDR}" \
    --private-key="${PRIVATE_KEY}" \
    --evm-rpc-urls="${EVM_RPC_URLS}" \
    --log-level="debug"
