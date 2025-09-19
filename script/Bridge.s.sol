// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {Script, console} from "forge-std/Script.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IOFT, SendParam, MessagingFee} from "@layerzerolabs/oft-evm/contracts/interfaces/IOFT.sol";
import {OptionsBuilder} from "@layerzerolabs/oapp-evm/contracts/oapp/libs/OptionsBuilder.sol";
import "forge-std/StdJson.sol";

import {AID} from "../src/AID.sol";
import {Minter} from "../src/Minter.sol";
import {AidOFTAdapter} from "../src/OFT/AidOFTAdapter.sol";
import {MockERC20} from "./mock/MockERC20.sol";

contract Bridge is Script {
    using SafeERC20 for MockERC20;
    using OptionsBuilder for bytes;
    using stdJson for string;

    // Deployment info file for our DVN deployment
    string constant DEPLOYMENT_INFO_FILE = "temp-network/deploy-data/dvn_deployment.json";
    uint256 constant MINT_AMOUNT = 100 * 1e18;
    uint256 constant BRIDGE_AMOUNT = 50 * 1e18;

    function run() external {
        // Load deployment info
        string memory deploymentInfo = vm.readFile(DEPLOYMENT_INFO_FILE);
        address aidAddress = vm.parseJsonAddress(deploymentInfo, ".chainA.aid");
        address adapterAddress = vm.parseJsonAddress(deploymentInfo, ".chainA.adapter");
        uint32 dstEid = uint32(vm.parseJsonUint(deploymentInfo, ".chainB.eid"));
        
        // --- Setup & broadcast config ---
        uint256 deployerPrivateKey = vm.envOr("PRIVATE_KEY", uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80));
        address sender = vm.addr(deployerPrivateKey);
        address recipient = sender;

        // Attach to existing contracts on Chain A
        AID aid = AID(aidAddress);
        AidOFTAdapter adapter = AidOFTAdapter(adapterAddress);

        // --- Step 1: Mint & Approve ---
        // For this test, we assume the deployer already has AID tokens from the deployment script.
        // If not, you would add minting logic here.
        vm.startBroadcast(deployerPrivateKey);
        console.log("Approving Adapter to spend AID tokens...");
        aid.approve(address(adapter), BRIDGE_AMOUNT);
        vm.stopBroadcast();

        // --- Step 2: Quote the cross-chain fee ---
        bytes32 toAddress = bytes32(uint256(uint160(recipient)));
        bytes memory options = OptionsBuilder.newOptions().addExecutorLzReceiveOption(200000, 0);

        SendParam memory sendParam = SendParam({
            dstEid: dstEid,
            to: toAddress,
            amountLD: BRIDGE_AMOUNT,
            minAmountLD: BRIDGE_AMOUNT,
            extraOptions: options,
            composeMsg: bytes(""),
            oftCmd: bytes("")
        });

        console.log("\nQuoting cross-chain fee...");
        MessagingFee memory fee = adapter.quoteSend(sendParam, false);
        
        console.log("  - Quoted Native Fee:", fee.nativeFee, "wei");
        console.log("  - Quoted LzToken Fee:", fee.lzTokenFee, "wei");

        // --- Step 3: Pay the fee and send the message ---
        vm.startBroadcast(deployerPrivateKey);
        console.log("\nSending AID tokens with total fee...");
        adapter.send{value: fee.nativeFee}(sendParam, fee, sender);
        vm.stopBroadcast();

        console.log("---------------------------------------------------------------------");
        console.log("Done! Cross-chain `send` transaction submitted successfully on Chain A!");
        console.log("From (Chain A): %s", sender);
        console.log("To (Chain B): %s", recipient);
        console.log("Amount: %d AID", BRIDGE_AMOUNT / (10 ** 18));
        console.log("---------------------------------------------------------------------");
        console.log("\nNext Steps:");
        console.log("1. Check the dvn-worker logs to see the packet being processed.");
        console.log("2. Once delivered, check your balance on Chain B!");
        console.log("---------------------------------------------------------------------");
    }
}
