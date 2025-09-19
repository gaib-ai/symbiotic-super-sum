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
import {ReceiveUlnSymbiotic} from "../src/uln/ReceiveUlnSymbiotic.sol";
import {EndpointV2View} from "@layerzerolabs/lz-evm-protocol-v2/contracts/EndpointV2View.sol";
import {ExecutionState} from "@layerzerolabs/lz-evm-protocol-v2/contracts/EndpointV2ViewUpgradeable.sol";
import {ReceiveUln302View, VerificationState} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/uln302/ReceiveUln302View.sol";


contract Executor is Script {
    using stdJson for string;

    string constant DEPLOYMENT_INFO_FILE = "temp-network/deploy-data/dvn_deployment.json";
    
    bytes32 constant PACKET_SENT_TOPIC = keccak256("PacketSent(bytes,bytes,address)");

    function run() external {
        uint256 deployerPrivateKey = vm.envOr("PRIVATE_KEY", uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80));

        // --- 1. Load deployment info from file ---
        string memory deploymentJson = vm.readFile(DEPLOYMENT_INFO_FILE);

        // --- Chain A Setup (to find the event) ---
        string memory rpcA = "http://localhost:8545";
        address endpointA_Addr = deploymentJson.readAddress(".chainA.endpoint");

        // --- Chain B Setup ---
        string memory rpcB = "http://localhost:8546";
        address endpointB_Addr = deploymentJson.readAddress(".chainB.endpoint");
        address receiveLibB_Addr = deploymentJson.readAddress(".chainB.receiveUln");

        // --- 2. Find the latest PacketSent event on Chain A ---
        vm.createSelectFork(rpcA);
        console.log("Switched to Chain A to find the latest PacketSent event...");

        bytes32[] memory topics = new bytes32[](1);
        topics[0] = PACKET_SENT_TOPIC;
        Vm.EthGetLogs[] memory logs = vm.eth_getLogs(0, block.number, endpointA_Addr, topics);
        

        if (logs.length == 0) {
            console.log("Error: No PacketSent event found on Chain A. Did you run Bridge.s.sol first?");
            return;
        }
        
        (bytes memory encodedPacket, , ) = abi.decode(logs[logs.length - 1].data, (bytes, bytes, address));
        console.log("Found PacketSent event!");
        
        // --- 3. Decode packet and prepare data ---
        Packet memory packet = _parsePacket(encodedPacket);
        bytes32 payloadHash = keccak256(abi.encodePacked(packet.guid, packet.message));
        Origin memory origin = Origin({srcEid: packet.srcEid, sender: AddressCast.toBytes32(packet.sender), nonce: packet.nonce});
        address receiver = AddressCast.toAddress(packet.receiver);
        bytes memory packetHeader = slice(encodedPacket, 1, 112); // Skip 1 version byte, take 112 bytes for header

        // --- 4. Switch to Chain B for execution ---
        vm.createSelectFork(rpcB);
        console.log("\nSwitched to Chain B to simulate Executor...");

        // Deploy View contracts for off-chain simulation
        EndpointV2View endpointViewB = new EndpointV2View();
        endpointViewB.initialize(endpointB_Addr);
        ReceiveUln302View receiveUln302ViewB = new ReceiveUln302View();
        // Our ReceiveUlnSymbiotic is compatible with the view for getUlnConfig and verifiable states
        receiveUln302ViewB.initialize(endpointB_Addr, receiveLibB_Addr); 
        
        ReceiveUlnSymbiotic receiveLibB = ReceiveUlnSymbiotic(receiveLibB_Addr);
        EndpointV2 endpointB = EndpointV2(endpointB_Addr);

        // --- 5. Check Verification Status and Commit ---
        console.log("\n--- Committer Role ---");
        VerificationState vState = receiveUln302ViewB.verifiable(packetHeader, payloadHash);

        if (vState == VerificationState.Verified) {
            console.log("Info: Packet verification already committed. Skipping commit.");
        } else if (vState == VerificationState.Verifying) {
            // In our setup, with only one DVN, it should go from NotVerified to Verifiable directly.
            // This state implies something is wrong or more DVNs are needed.
            console.log("Error: Packet is still verifying. Not enough DVN signatures aggregated?");
            return;
        } else { // State is Verifiable (or NotVerified and we will try to commit)
            console.log("Packet is verifiable. Broadcasting 'commitVerification'...");
            vm.startBroadcast(deployerPrivateKey);
            receiveLibB.commitVerification(packetHeader, payloadHash);
            vm.stopBroadcast();
            console.log("Verification committed successfully.");
        }

        // --- 6. Check Executable Status and Execute Message ---
        console.log("\n--- Executor Role ---");
        ExecutionState eState = endpointViewB.executable(origin, receiver);
        
        if (eState == ExecutionState.Executed) {
            console.log("Info: Packet has already been executed. Skipping execution.");
            return;
        } else if (eState == ExecutionState.NotExecutable) {
            console.log("Error: Packet is not executable even after commit. Check nonce order or other issues.");
            return;
        } else { // State is Executable
            console.log("Packet is executable. Broadcasting 'lzReceive'...");
            vm.startBroadcast(deployerPrivateKey);
            bytes memory extraData = abi.encodePacked(bytes32(0), bytes32(0));
            endpointB.lzReceive(origin, receiver, packet.guid, packet.message, extraData);
            vm.stopBroadcast();
            console.log("Message delivered successfully via lzReceive.");
        }

        console.log("\n-------------------------------------------------");
        console.log("Done. Executor successful. Message has been delivered!");
        console.log("-------------------------------------------------");
    }

    function _parsePacket(bytes memory encodedPacket) internal pure returns (Packet memory packet) {
        uint64 nonce;
        uint32 srcEid;
        address sender;
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
        packet.sender = sender;
        packet.dstEid = dstEid;
        packet.receiver = receiver;
        packet.guid = guid;
        
        // The packetHeader is everything up to and including the guid.
        // There is no version byte in the packed header for verify, but there is in the full encodedPacket for parsing.
        uint256 messageOffset = 113; // 1 byte version + 112 byte header
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
