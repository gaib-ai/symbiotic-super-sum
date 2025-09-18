#!/bin/sh
set -e

echo "Starting DVN Node..."

# Default values
RPC_URL_A=${RPC_URL_A:-"http://anvil:8545"}
RPC_URL_B=${RPC_URL_B:-"http://anvil-settlement:8546"}
PRIVATE_KEY=${PRIVATE_KEY:-"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"}

export RPC_URL_A
export RPC_URL_B
export PRIVATE_KEY

# Execute the main application
/go/bin/dvn-node
