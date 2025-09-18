// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {Script, console} from "forge-std/Script.sol";
import {Vm} from "forge-std/Vm.sol";
import "forge-std/StdJson.sol";

// LayerZero Contracts & Interfaces
import {EndpointV2} from "@layerzerolabs/lz-evm-protocol-v2/contracts/EndpointV2.sol";
import {Packet} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ISendLib.sol";
import {Origin} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";
import {AddressCast} from "@layerzerolabs/lz-evm-protocol-v2/contracts/protocol/libs/AddressCast.sol";
import {ReceiveUlnSymbiotic} from "../src/uln/ReceiveUlnSymbiotic.sol";

contract Executor is Script {
    using stdJson for string;

    string constant DEPLOYMENT_INFO_FILE = "temp-network/deploy-data/dvn_deployment.json";
    
    bytes32 constant PACKET_SENT_TOPIC = keccak256("PacketSent(bytes,bytes,address)");

    function run() external {
        uint256 deployerPrivateKey = vm.envOr("PRIVATE_KEY", 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);

        // --- 1. Load deployment info from file ---
        string memory deploymentJson = vm.readFile(DEPLOYMENT_INFO_FILE);

        // --- Chain A Setup (to find the event) ---
        string memory rpcA = "http://localhost:8545";
        address endpointA_Addr = deploymentJson.readAddress(".chainA.endpoint");

        // --- Chain B Setup ---
        string memory rpcB = "http://localhost:8456"; // Note: Port mismatch corrected from 8546 in original script
        address endpointB_Addr = deploymentJson.readAddress(".chainB.endpoint");
        address receiveLibB_Addr = deploymentJson.readAddress(".chainB.receiveLib");

        // --- 2. Find the latest PacketSent event on Chain A ---
        vm.createSelectFork(rpcA);
        console.log("Switched to Chain A to find the latest PacketSent event...");

        bytes32[] memory topics = new bytes32[](1);
        topics[0] = PACKET_SENT_TOPIC;
        Vm.Log[] memory logs = vm.getLogs({
            address: endpointA_Addr,
            topics: topics,
            fromBlock: 0,
            toBlock: block.number
        });
        

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
        bytes memory packetHeader = slice(encodedPacket, 0, 81);

        // --- 4. Switch to Chain B for execution ---
        vm.createSelectFork(rpcB);
        console.log("\nSwitched to Chain B to simulate Executor...");

        ReceiveUlnSymbiotic receiveLibB = ReceiveUlnSymbiotic(receiveLibB_Addr);
        EndpointV2 endpointB = EndpointV2(endpointB_Addr);

        // --- 5. Commit Verification ---
        console.log("Broadcasting 'commitVerification'...");
        vm.startBroadcast(deployerPrivateKey);
        receiveLibB.commitVerification(packetHeader, payloadHash);
        vm.stopBroadcast();
        console.log("Verification committed successfully.");
        
        // --- 6. Execute Message ---
        console.log("Broadcasting 'lzReceive'...");
        vm.startBroadcast(deployerPrivateKey);
        bytes memory extraData = abi.encodePacked(bytes32(0), bytes32(0));
        endpointB.lzReceive(origin, receiver, packet.guid, packet.message, extraData);
        vm.stopBroadcast();
        console.log("Message delivered successfully via lzReceive.");

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
        
        uint256 messageOffset = 113;
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
