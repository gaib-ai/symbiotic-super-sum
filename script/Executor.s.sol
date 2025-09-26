// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {Script, console} from "forge-std/Script.sol";
import {Vm} from "forge-std/Vm.sol";
import "forge-std/StdJson.sol";

// LayerZero Contracts & Interfaces
import {EndpointV2} from "@layerzerolabs/lz-evm-protocol-v2/contracts/EndpointV2.sol";
import {Packet} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ISendLib.sol";
import {Origin} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";
import {AddressCast} from "@layerzerolabs/lz-evm-protocol-v2/contracts/libs/AddressCast.sol";
import {ReceiveUln302} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/uln302/ReceiveUln302.sol";
import {EndpointV2View} from "@layerzerolabs/lz-evm-protocol-v2/contracts/EndpointV2View.sol";
import {ExecutionState} from "@layerzerolabs/lz-evm-protocol-v2/contracts/EndpointV2ViewUpgradeable.sol";
// Using the 302 view as it's compatible for checking verification state
import {ReceiveUln302View, VerificationState} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/uln302/ReceiveUln302View.sol";
import {UlnConfig} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/UlnBase.sol";


contract Executor is Script {
    using stdJson for string;

    string constant DEPLOYMENT_INFO_FILE = "temp-network/deploy-data/dvn_deployment.json";
    
    // PacketSent(bytes encodedPacket, bytes options, address sendLibrary)
    bytes32 constant PACKET_SENT_TOPIC = keccak256("PacketSent(bytes,bytes,address)");
    // ExecutorFeePaid(address executor, uint256 fee)
    bytes32 constant EXECUTOR_FEE_PAID_TOPIC = keccak256("ExecutorFeePaid(address,uint256)");
    // In our custom DVN, this event is emitted when the DVN worker verifies a payload
    bytes32 constant PAYLOAD_VERIFIED_TOPIC = keccak256("PayloadVerified(address,bytes,uint256,bytes32)");

    function run() external {
        uint256 deployerPrivateKey = vm.envOr("PRIVATE_KEY", uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80));

        // --- 1. Load deployment info from file ---
        string memory deploymentJson = vm.readFile(DEPLOYMENT_INFO_FILE);

        // --- Chain A Setup (to find the event) ---
        string memory rpcA = "http://localhost:8545";
        uint256 forkA = vm.createFork(rpcA);
        address endpointA_Addr = deploymentJson.readAddress(".chainA.endpoint");
        // These addresses are needed to check if the executor was paid.
        // Ensure your deployment script writes these to the deployment file.
        address sendLibA_Addr = deploymentJson.readAddress(".chainA.sendLib");
        address executorA_Addr = deploymentJson.readAddress(".chainA.executor");

        // --- Chain B Setup ---
        string memory rpcB = "http://localhost:8546";
        uint256 forkB = vm.createFork(rpcB);
        address endpointB_Addr = deploymentJson.readAddress(".chainB.endpoint");
        address receiveLibB_Addr = deploymentJson.readAddress(".chainB.receiveUln");

        // --- 2. Find the latest PacketSent event on Chain A ---
        vm.selectFork(forkA);
        console.log("Switched to Chain A to find the latest PacketSent event...");

        bytes32[] memory topics = new bytes32[](1);
        topics[0] = PACKET_SENT_TOPIC;
        Vm.EthGetLogs[] memory logs = vm.eth_getLogs(0, block.number, endpointA_Addr, topics);
        

        if (logs.length == 0) {
            console.log("Error: No PacketSent event found on Chain A. Did you run Bridge.s.sol first?");
            return;
        }
        
        for (uint i = 0; i < logs.length; i++) {
            vm.selectFork(forkA); // Re-select Chain A fork at the start of each iteration
            (bytes memory encodedPacket, , ) = abi.decode(logs[i].data, (bytes, bytes, address));
            console.log("\nProcessing PacketSent event %d...", i);

            bool success = _processPacket(
                encodedPacket,
                deployerPrivateKey,
                forkB,
                endpointB_Addr,
                receiveLibB_Addr,
                sendLibA_Addr,
                executorA_Addr
            );

            if (!success) {
                console.log("Halting script due to an error in packet processing.");
                return; // Stop the entire script
            }
        }
    }

    function _processPacket(
        bytes memory encodedPacket,
        uint256 deployerPrivateKey,
        uint256 forkB,
        address endpointB_Addr,
        address receiveLibB_Addr,
        address sendLibA_Addr,
        address executorA_Addr
    ) private returns (bool) {
        // --- 2b. Check for executor fee on Chain A ---
        {
            bytes32[] memory topics = new bytes32[](1);
            topics[0] = EXECUTOR_FEE_PAID_TOPIC;
            Vm.EthGetLogs[] memory feeLogs = vm.eth_getLogs(0, block.number, sendLibA_Addr, topics);
            bool executorPaid = false;
            if (feeLogs.length > 0) {
                for (uint j = 0; j < feeLogs.length; j++) {
                    (address executor, ) = abi.decode(feeLogs[j].data, (address, uint256));
                    if (executor == executorA_Addr) {
                        executorPaid = true;
                        break;
                    }
                }
            }
            if (!executorPaid) {
                console.log("Error: Executor was not paid for this packet.");
                return false;
            }
            console.log("Found ExecutorFeePaid event, Executor is assigned and paid.");
        }

        // --- 3. Decode packet and prepare data ---
        Packet memory packet = _parsePacket(encodedPacket);
        Origin memory origin = Origin({srcEid: packet.srcEid, sender: AddressCast.toBytes32(packet.sender), nonce: packet.nonce});
        address receiver = AddressCast.toAddress(packet.receiver);
        bytes memory packetHeader = slice(encodedPacket, 0, 81);
        bytes32 packetHash = keccak256(packetHeader);
        bytes32 payloadHash = keccak256(abi.encodePacked(packet.guid, packet.message));
        
        // --- 4. Switch to Chain B ---
        vm.selectFork(forkB);
        console.log("Switched to Chain B to simulate Executor...");

        EndpointV2View endpointViewB = new EndpointV2View();
        endpointViewB.initialize(endpointB_Addr);

        // --- 5a. Check if already executed ---
        if (endpointViewB.executable(origin, receiver) == ExecutionState.Executed) {
            console.log("Info: Packet has already been executed. Skipping.");
            return true; // Success, continue to next packet
        }

        // --- 5b. Commit verification ---
        bool committed = _commitVerificationOnB(
            packetHeader, payloadHash, packetHash, receiver, packet.srcEid, 
            endpointB_Addr, receiveLibB_Addr, deployerPrivateKey
        );
        if (!committed) return false; // Error during commit

        // --- 6. Execute Message ---
        return _executeLzReceiveOnB(
            origin, receiver, packet.guid, packet.message,
            endpointViewB, endpointB_Addr, deployerPrivateKey
        );
    }

    function _commitVerificationOnB(
        bytes memory packetHeader,
        bytes32 payloadHash,
        bytes32 packetHash,
        address receiver,
        uint32 srcEid,
        address endpointB_Addr,
        address receiveLibB_Addr,
        uint256 deployerPrivateKey
    ) private returns (bool) {
        console.log("\n--- Committer Role ---");
        ReceiveUln302View receiveUln302ViewB = new ReceiveUln302View();
        receiveUln302ViewB.initialize(endpointB_Addr, receiveLibB_Addr);
        ReceiveUln302 receiveLibB = ReceiveUln302(receiveLibB_Addr);
        
        UlnConfig memory ulnConfig = receiveLibB.getUlnConfig(receiver, srcEid);
        uint8 requiredDVNs = uint8(ulnConfig.confirmations);
        console.log("Required DVN confirmations for this packet: %d", requiredDVNs);

        uint foundVerifications = _findVerifications(receiveLibB_Addr, payloadHash, packetHash);

        if (foundVerifications < requiredDVNs) {
            console.log("Error: Not enough DVN verifications found. Expected %d, found %d.", requiredDVNs, foundVerifications);
            return false;
        }
        console.log("Found %d verification submissions from the DVN for this packet. Proceeding to check verifiable state.", foundVerifications);

        VerificationState vState = receiveUln302ViewB.verifiable(packetHeader, payloadHash);
        if (vState == VerificationState.Verified) {
            console.log("Info: Packet verification already committed. Skipping commit.");
        } else if (vState == VerificationState.Verifying) {
            console.log("Error: Packet is still verifying. Not enough DVN signatures aggregated?");
            return false;
        } else { // Verifiable
            console.log("Packet is verifiable. Broadcasting 'commitVerification'...");
            vm.startBroadcast(deployerPrivateKey);
            receiveLibB.commitVerification(packetHeader, payloadHash);
            vm.stopBroadcast();
            console.log("Verification committed successfully.");
        }
        return true;
    }

    function _executeLzReceiveOnB(
        Origin memory origin,
        address receiver,
        bytes32 guid,
        bytes memory message,
        EndpointV2View endpointViewB,
        address endpointB_Addr,
        uint256 deployerPrivateKey
    ) private returns (bool) {
        console.log("\n--- Executor Role ---");
        ExecutionState eState = endpointViewB.executable(origin, receiver);
        
        if (eState == ExecutionState.NotExecutable) {
            console.log("Error: Packet is not executable even after commit. Check nonce order or other issues.");
            return false;
        } else if (eState == ExecutionState.Executable) {
            console.log("Packet is executable. Broadcasting 'lzReceive'...");
            EndpointV2 endpointB = EndpointV2(endpointB_Addr);
            vm.startBroadcast(deployerPrivateKey);
            bytes memory extraData = abi.encodePacked(bytes32(0), bytes32(0));
            endpointB.lzReceive(origin, receiver, guid, message, extraData);
            vm.stopBroadcast();
            console.log("Message delivered successfully via lzReceive.");
        }
        
        console.log("-------------------------------------------------");
        console.log("Executor successful. Message has been delivered!");
        console.log("-------------------------------------------------");
        return true;
    }

    function _findVerifications(address receiveLibB_Addr, bytes32 payloadHash, bytes32 packetHash) private returns (uint) {
        bytes32[] memory verificationTopics = new bytes32[](1);
        verificationTopics[0] = PAYLOAD_VERIFIED_TOPIC;
        Vm.EthGetLogs[] memory logs = vm.eth_getLogs(0, block.number, receiveLibB_Addr, verificationTopics);
        
        uint foundVerifications = 0;
        for (uint i = 0; i < logs.length; i++) {
            (, bytes memory header, , bytes32 proofHash) =
                abi.decode(logs[i].data, (address, bytes, uint256, bytes32));

            if (proofHash == payloadHash && keccak256(header) == packetHash) {
                foundVerifications++;
            }
        }
        return foundVerifications;
    }

    function _parsePacket(bytes memory encodedPacket) internal pure returns (Packet memory packet) {
        // The packet is abi.encodePacked with the following structure:
        // 1 byte: version
        // 8 bytes: nonce
        // 4 bytes: srcEid
        // 32 bytes: sender
        // 4 bytes: dstEid
        // 32 bytes: receiver
        // 32 bytes: guid
        // N bytes: message
        uint64 nonce;
        uint32 srcEid;
        bytes32 sender;
        uint32 dstEid;
        bytes32 receiver;
        bytes32 guid;

        assembly {
            let data_ptr := add(encodedPacket, 0x20)
            nonce := shr(192, mload(add(data_ptr, 1)))
            srcEid := shr(224, mload(add(data_ptr, 9)))
            sender := mload(add(data_ptr, 13))
            dstEid := shr(224, mload(add(data_ptr, 45)))
            receiver := mload(add(data_ptr, 49))
            guid := mload(add(data_ptr, 81))
        }

        packet.nonce = nonce;
        packet.srcEid = srcEid;
        packet.sender = AddressCast.toAddress(sender);
        packet.dstEid = dstEid;
        packet.receiver = receiver;
        packet.guid = guid;
        
        uint256 messageOffset = 113; // 1 + 8 + 4 + 32 + 4 + 32 + 32
        uint256 messageLength = encodedPacket.length - messageOffset;
        packet.message = slice(encodedPacket, messageOffset, messageLength);
    }

    function slice(bytes memory data, uint256 start, uint256 length) internal pure returns (bytes memory) {
        bytes memory result = new bytes(length);
        for (uint256 i = 0; i < length; i++) {
            result[i] = data[start + i];
        }
        return result;
    }
}
