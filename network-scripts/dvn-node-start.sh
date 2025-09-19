#!/bin/sh
set -e

echo "Starting DVN Node..."

# Default values
RELAY_ADDR=${1:-"relay-sidecar-1:8080"}
PRIVATE_KEY=${2:-"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"} # Default Anvil key
RPC_URL_A=${RPC_URL_A:-"http://anvil:8545"}
RPC_URL_B=${RPC_URL_B:-"http://anvil-settlement:8546"}

export RELAY_ADDR
export RPC_URL_A
export RPC_URL_B
export PRIVATE_KEY

# Execute the main application
/app/dvn-node
