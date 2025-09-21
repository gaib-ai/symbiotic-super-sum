# LayerZero DVN Secured by Symbiotic

This project demonstrates a powerful integration between LayerZero's omnichain messaging protocol and Symbiotic's economic security framework. It replaces a standard LayerZero Decentralized Verifier Network (DVN) with one whose security is derived directly from the economic stake managed by a Symbiotic validator set.

This repository is a fusion of two core concepts:
-   **[aid-contracts](https://github.com/gaib-ai/aid-contracts)**: Provides a high-fidelity local testing environment for a complete LayerZero V2 deployment, including a cross-chain application (OApp).
-   **[symbiotic-super-sum](https://github.com/symbioticfi/symbiotic-super-sum)**: Provides a Dockerized, multi-chain local testnet and an off-chain worker infrastructure powered by the Symbiotic Relay SDK.

By combining these, we create a fully-functional prototype of a DVN where packet verification is backed by real economic stake, rather than a simple multi-sig of trusted addresses.

---

## Architecture Overview

The system runs a local, containerized network consisting of two independent blockchains (Anvil instances), a full Symbiotic Relay network, and a dedicated off-chain worker that acts as the DVN.

### Key On-Chain Components

-   **`SymbioticLzDVN.sol`**: A custom LayerZero DVN contract.
    -   **On the source chain**, it implements the `IDVN` interface to provide fee quotes for its verification services.
    -   **On the destination chain**, it exposes a `verifyWithSymbiotic` function. This function accepts a proof from the Symbiotic network, verifies it against the on-chain `Settlement` contract, and—if valid—calls the standard `ReceiveUln302.verify()` function to inform LayerZero of the successful verification.
-   **`ReceiveUln302.sol`**: The **standard** LayerZero "receive" Message Library.
    - It is *not* a custom contract. We use the official, audited version.
    - Through standard LayerZero configuration for the `AID` OApp, it is set up to trust *only* the `SymbioticLzDVN` contract's address as a valid verifier.
-   **Symbiotic & LayerZero Stacks**: The full suite of Symbiotic contracts (`Settlement`, `ValSetDriver`, etc.) and the standard LayerZero contracts (`EndpointV2`, `SendUln302`, etc.) are deployed. The Symbiotic contracts are deployed by the `LocalDeploy.s.sol` script during the Docker network setup, while the LayerZero stack and the `AID` OApp are deployed by the `SymbioticLzDvnDeploy.s.sol` script.

### Off-Chain Nodes (`dvn-node`)

This project runs multiple instances of a Go application that acts as the bridge between the two protocols. Each `dvn-node` performs the following steps:
1.  **Listens** for `PacketSent` events on the source chain's LayerZero `EndpointV2`.
2.  Upon receiving an event, all nodes will **construct a task** for the Symbiotic relay network.
3.  Each node **requests a proof** (an aggregated BLS signature) from its local `relay-sidecar` network.
4.  The nodes then **race to submit a transaction** to the `SymbioticLzDVN` contract on the destination chain. Only the first transaction will succeed; others will fail, which is expected behavior in a decentralized network simulation.
5.  This action completes the verification, allowing the cross-chain message to be executed by a separate Executor.

##### DVN Node Identity and Funding

It's important to note that the `generate_network.sh` script configures each `dvn-node` instance to use the **private key of its corresponding Symbiotic operator**.

-   **How it Works**: The script generates a set of operator keys for the Symbiotic relay network. In a real-world scenario, these operators would stake assets to secure the network. In our local simulation, the `genesis-generator` service automatically funds these operator addresses with plenty of test ETH on both Anvil chains.
-   **Why it Matters**: By reusing the funded operator keys, the `dvn-node` has a pre-funded wallet that can pay the gas fees required to submit verification transactions to the destination chain. This eliminates the need to generate and manually fund separate wallets for the DVN workers.
-   **On Other Testnets**: If you adapt this project for a public testnet, you would need to ensure the operator addresses are funded on all participating chains.

### Workflow Sequence Diagram

```mermaid
sequenceDiagram
    participant User
    participant ChainA as "Anvil Chain A (Src)"
    participant dvn-node as "Off-Chain DVN Node"
    participant relay-network as "Symbiotic Relay Network"
    participant ChainB as "Anvil Chain B (Dst)"
    participant Executor as "Executor Script"

    User->>ChainA: 1. OApp.send() (pays fee)
    activate ChainA
    ChainA-->>dvn-node: 2. Emits PacketSent Event
    deactivate ChainA

    dvn-node->>dvn-node: 3. Hashes header, packs with payload hash
    dvn-node->>relay-network: 4. Request proof for packed data
    activate relay-network
    relay-network-->>relay-network: 5. Hashes packed data, gets signatures
    relay-network-->>dvn-node: 6. Returns aggregated BLS signature (proof)
    deactivate relay-network
    
    dvn-node->>ChainB: 7. SymbioticLzDVN.verifyWithSymbiotic(proof)
    activate ChainB
    ChainB-->>ChainB: 8. DVN re-creates hash, calls Settlement.verify() & ReceiveUln302.verify()
    Note over ChainB: ReceiveUln302 emits <br/> PayloadVerified event
    ChainB-->>Executor: 9. DVN worker's action is now complete
    deactivate ChainB

    User->>Executor: 10. Runs executor script
    Executor->>ChainB: 11. commitVerification() & lzReceive()
    activate ChainB
    ChainB-->>User: 12. Message Delivered
    deactivate ChainB
```

### Deployed Contracts in the Local Environment

The `SymbioticLzDvnDeploy.s.sol` script deploys a complete set of contracts on *each* of the two local chains.

#### Application Layer Contracts (OApp)
| Contract | Role |
| :--- | :--- |
| `AID` | The core ERC-20 token contract, using upgradeable proxies. |
| `Minter` | Mints new `AID` tokens against a stablecoin (`MockERC20`). |
| `MockERC20` | A mock stablecoin used as collateral. |
| `AidOFTAdapter` | The OApp entry point, responsible for initiating cross-chain `AID` transfers. |
| `AidOFTMintBurner` | Handles the burning and minting of `AID` tokens on behalf of the adapter. |

#### LayerZero v2 Protocol Stack
| Contract | Role |
| :--- | :--- |
| `EndpointV2` | The main LayerZero entry point on each chain for sending and receiving messages. |
| `SendUln302` | **Standard** "send" message library for fee calculation and packet formatting. |
| `ReceiveUln302`| **Standard** "receive" library, configured to trust our `SymbioticLzDVN` for this OApp. |
| `PriceFeed` | Provides gas price data for cross-chain fee calculation. |
| `SymbioticLzDVN`| **Custom** DVN contract that provides fee quotes and verifies proofs against the Symbiotic `Settlement` contract. |
| `Executor` | A standard LayerZero contract responsible for final message delivery. |

#### Symbiotic Protocol Stack
| Contract | Role |
| :--- | :--- |
| `Settlement` | Per-chain contract that verifies aggregated BLS signatures from the relay network. |
| `ValSetDriver` | Exposes epoched validator sets to the off-chain relay nodes. |
| `KeyRegistry` | Stores operators’ BLS and ECDSA keys. |
| `VotingPowerProvider` | Derives validator voting power from vault stakes. |
| *Core Contracts* | Other core contracts managing vaults, operators, and networks, deployed via the local setup scripts. |

---

## Local Simulation vs. Production System

It is crucial to understand that this project provides a **high-fidelity logical simulation**, not a production-ready system. The `dvn-node` and Forge-based `Executor` are designed for local testing and verification of on-chain logic. Building and operating resilient, secure off-chain services requires significant additional engineering effort.

**What's Not Included (The Path to a Production System):**

*   **Production-Grade Service Architecture:** The `dvn-node` is a simple application. A real-world system requires continuously running, fault-tolerant services with robust process management, automated restarts, and comprehensive logging.
*   **Persistent State Management:** Production workers need a robust database (e.g., PostgreSQL) to track in-flight messages, transaction statuses, and retry counts to ensure data integrity during service restarts.
*   **Secure Private Key Management:** Using private keys from a local `.env` file is insecure. A live system demands a secure key management solution like HashiCorp Vault.
*   **RPC Redundancy and Error Handling:** Production services must handle RPC provider downtime by implementing logic for failover to redundant nodes and sophisticated retry mechanisms.
*   **Dynamic Gas Price Management:** A production worker must implement a dynamic gas fee strategy, likely integrating with gas station APIs, to ensure transactions are mined in a timely and cost-effective manner.
*   **Real Economic Security:** This simulation uses mock staking tokens. A production system involves real assets and a carefully designed economic model with staking rewards and slashing conditions for misbehavior.

---

## Quick Start: End-to-End Local Test

This guide will walk you through setting up the local network, deploying all contracts, and sending a cross-chain transaction.

### Prerequisites

-   [Docker](https://www.docker.com/get-started) and Docker Compose
-   [Foundry](https://getfoundry.sh/)
-   [Node.js](https://nodejs.org/en) and [npm](https://www.npmjs.com/get-npm)
-   [yarn](https://classic.yarnpkg.com/en/docs/install)

### Installation and Setup

1.  **Clone the Repository**
    ```bash
    git clone --recursive <your-repo-url>
    cd <your-repo-name>
    ```
    *Note: The `--recursive` flag is important as it automatically initializes and clones the git submodules.*

2.  **Install Submodules Manually (if needed)**
    If you cloned the repository without the `--recursive` flag, you can initialize the submodules manually:
    ```bash
    git submodule update --init --recursive
    ```

3.  **Install NodeJS Dependencies**
    Install the root project dependencies and the required packages for the LayerZero v2 library.
    ```bash
    npm i
    cd lib/LayerZero-v2
    yarn
    cd ../..
    ```

4.  **Fix OpenZeppelin v5 Import Paths**
    This project uses both v4 and v5 of OpenZeppelin contracts to maintain compatibility with its dependencies while using modern features. The default import path points to v4, which causes compilation errors in the v5 upgradeable contracts. Run the following command to fix the import paths:
    ```bash
    find lib/openzeppelin-contracts-upgradeable-v5 -type f -name "*.sol" -exec sed -i '' 's|from "@openzeppelin/contracts/|from "@openzeppelin-v5/contracts/|g' {} +
    ```

### Step 1: Setup Environment File

The Forge scripts require a `PRIVATE_KEY` to be set in the environment. Copy the provided example file.

```bash
cp .env.example .env
```

### Step 2: Generate Network Configuration

This script creates a `temp-network` directory containing a `docker-compose.yml` file and a `.env` file for the DVN worker. You can configure the number of operators in the Symbiotic network.

```bash
./generate_network.sh
```

### Step 3: Start the Local Network

Navigate into the newly created directory and start all services in the background. This will pull necessary Docker images, build the `dvn-node`, and start the two blockchains, the Symbiotic relay sidecars, and the DVN worker.

```bash
cd temp-network && docker compose up --build -d && cd ..
```
The first startup may take a few minutes. You can monitor the progress with `docker compose logs -f`.

### Step 4: Deploy and Configure Contracts

Once all services are running, run the unified deployment script. This Forge script will deploy and configure the entire Symbiotic and LayerZero stack on both local chains.

```bash
forge script script/SymbioticLzDvnDeploy.s.sol --rpc-url http://localhost:8545 --broadcast --ffi
```
This will create a `dvn_deployment.json` file in `temp-network/deploy-data/` with the addresses of all deployed contracts.

### Step 5: Initiate a Cross-Chain Transaction

With everything deployed and running, use the `Bridge` script to send 50 `AID` tokens from Chain A (31337) to Chain B (31338).

```bash
forge script script/Bridge.s.sol --rpc-url http://localhost:8545 --broadcast --ffi
```

### Step 6: Observe the DVN and Run the Executor

This step involves two distinct parts: the automated off-chain DVN worker processing the packet, and the manual execution of the delivery.

1.  **Observe the Automated DVN Worker**: The `dvn-node` services, which you started with Docker, are constantly monitoring Chain A. Once they detect the `PacketSent` event from the previous step, they will automatically fetch a proof from the Symbiotic network and submit a verification transaction to Chain B.

    You can watch this happen in real-time by viewing the logs:
    ```bash
    docker compose logs -f dvn-node-1 dvn-node-2 #... and so on
    ```
    You should see one node successfully submit the verification, while the others will report a (safe and expected) failure because the packet was already verified by the winner of the race.

2.  **Manually Run the Executor**: After a DVN node has successfully submitted its verification on Chain B, the packet is ready for delivery. Run the `Executor` script to complete the process. This script finds the verified packet and calls `lzReceive` to deliver the message to the destination OApp.
    ```bash
    forge script script/Executor.s.sol --rpc-url http://localhost:8546 --broadcast --ffi
    ```

### Step 7: Verify the Result

Check your `AID` token balance on the destination chain (Chain B) to confirm the cross-chain transfer was successful.

```bash
# Get your address from the private key in your .env file
SENDER=$(cast wallet address --private-key $(grep PRIVATE_KEY .env | cut -d '=' -f2))

# Get the AID contract address on Chain B from the deployment file
AID_B=$(jq -r '.chainB.aid' temp-network/deploy-data/dvn_deployment.json)

# Check the balance
cast call $AID_B "balanceOf(address)" $SENDER --rpc-url http://localhost:8546 | cast --to-dec
```
The result should be `50000000000000000000` (which is 50 tokens, as `AID` has 18 decimals).

---

## Network Services

The local environment consists of the following services:

-   `anvil`: The source blockchain (Chain A, EID 31337) running on port `8545`.
-   `anvil-settlement`: The destination blockchain (Chain B, EID 31338) running on port `8546`.
-   `deployer`: A short-lived service that waits for the chains to be healthy.
-   `genesis-generator`: A short-lived service that generates the configuration for the Symbiotic relay network.
-   `relay-sidecar-*`: The nodes of the Symbiotic relay network.
-   `dvn-node-*`: The off-chain nodes that listen for LayerZero events and compete to submit Symbiotic-backed verifications.

## Cleanup

To stop and remove all running containers and networks, run the following command from within the `temp-network` directory:
```bash
docker compose down -v
```
