// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {Vm} from "forge-std/Vm.sol";
import {console} from "forge-std/console.sol";

// Symbiotic Contracts
import {LocalDeploy} from "./LocalDeploy.s.sol";
import {ISettlement} from "@symbioticfi/relay-contracts/interfaces/modules/settlement/ISettlement.sol";

// LayerZero Contracts
import {EndpointV2} from "@layerzerolabs/lz-evm-protocol-v2/contracts/EndpointV2.sol";
import {SendUln302} from "@layerzerolabs/lz-evm-protocol-v2/contracts/messagelib/uln/uln302/SendUln302.sol";
import {UlnConfig, SetDefaultUlnConfigParam} from "@layerzerolabs/lz-evm-protocol-v2/contracts/messagelib/uln/UlnBase.sol";
import {SetConfigParam} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/IMessageLibManager.sol";
import {ExecutorConfig} from "@layerzerolabs/lz-evm-protocol-v2/contracts/messagelib/SendLibBase.sol";
import {PriceFeed} from "@layerzerolabs/lz-evm-protocol-v2/contracts/messagelib/PriceFeed.sol";
import {Executor} from "@layerzerolabs/lz-evm-protocol-v2/contracts/messagelib/Executor.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

// Custom Integrated Contracts
import {SymbioticLzDVN} from "../src/SymbioticLzDVN.sol";
import {ReceiveUlnSymbiotic} from "../src/uln/ReceiveUlnSymbiotic.sol";

// Application Contracts
import {AID} from "../src/AID.sol";
import {Minter} from "../src/Minter.sol";
import {AidOFTAdapter} from "../src/OFT/AidOFTAdapter.sol";
import {AidOFTMintBurner} from "../src/OFT/AidOFTMintBurner.sol";
import {MockERC20} from "./mock/MockERC20.sol";

contract SymbioticLzDvnDeploy is LocalDeploy {
    // Local Chain A (e.g., Anvil Fork 1)
    uint32 internal constant localChainA_Eid = 31337;
    string internal localChainA_Rpc = "http://localhost:8545";
    uint256 internal localChainA_PrivateKey;

    // Local Chain B (e.g., Anvil Fork 2)
    uint32 internal constant localChainB_Eid = 31338;
    string internal localChainB_Rpc = "http://localhost:8546";
    uint256 internal localChainB_PrivateKey;
    
    string constant DEPLOYMENT_INFO_FILE = "temp-network/deploy-data/dvn_deployment.json";

    function run() public override {
        localChainA_PrivateKey = vm.envOr("PRIVATE_KEY", 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);
        localChainB_PrivateKey = vm.envOr("PRIVATE_KEY", 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);

        uint256 forkA = vm.createFork(localChainA_Rpc, 1);
        uint256 forkB = vm.createFork(localChainB_Rpc, 1);

        address deployer = vm.addr(localChainA_PrivateKey);
        console.log("Deploying contracts with the account:", deployer);
        
        // --- Step 1: Deploy on Local Chain A ---
        vm.selectFork(forkA);
        vm.startBroadcast(localChainA_PrivateKey);
        console.log("\n--- Deploying on Local Chain A (%d) ---", block.chainid);
        // Deploy Symbiotic Stack
        _deploySymbioticStack(deployer);
        // Deploy LayerZero Stack
        (
            EndpointV2 endpointA,
            SendUln302 sendLibA,
            ReceiveUlnSymbiotic receiveLibA,
            SymbioticLzDVN dvnA,
            PriceFeed priceFeedA,
            Executor executorA
        ) = _deployLzStack(deployer);
        // Deploy Application
        (AID aidA, Minter minterA, AidOFTAdapter adapterA, MockERC20 assetA) = _deployAppContracts(deployer, address(endpointA));
        vm.stopBroadcast();


        // --- Step 2: Deploy on Local Chain B ---
        vm.selectFork(forkB);
        vm.startBroadcast(localChainB_PrivateKey);
        console.log("\n--- Deploying on Local Chain B (%d) ---", block.chainid);
         // Deploy Symbiotic Stack
        _deploySymbioticStack(deployer);
        // Deploy LayerZero Stack
        (
            EndpointV2 endpointB,
            SendUln302 sendLibB,
            ReceiveUlnSymbiotic receiveLibB,
            SymbioticLzDVN dvnB,
            PriceFeed priceFeedB,
            Executor executorB
        ) = _deployLzStack(deployer);
        // Deploy Application
        (AID aidB, Minter minterB, AidOFTAdapter adapterB, MockERC20 assetB) = _deployAppContracts(deployer, address(endpointB));
        vm.stopBroadcast();


        // --- Step 3: Configure Both Chains ---
        console.log("\n--- Configuring Chains ---");
        // Configure Chain A
        vm.selectFork(forkA);
        vm.startBroadcast(localChainA_PrivateKey);
        console.log("Configuring Chain A...");
        _configureLayerZero(endpointA, sendLibA, receiveLibA, dvnA, priceFeedA, executorA, adapterA, address(adapterB), localChainB_Eid);
        vm.stopBroadcast();
        
        // Configure Chain B
        vm.selectFork(forkB);
        vm.startBroadcast(localChainB_PrivateKey);
        console.log("Configuring Chain B...");
        _configureLayerZero(endpointB, sendLibB, receiveLibB, dvnB, priceFeedB, executorB, adapterB, address(adapterA), localChainA_Eid);
        vm.stopBroadcast();
        
        // --- Step 4: Write deployment file ---
        string memory root = '{"chainA":{';
        // Serialize for chain A
        root = string.concat(root, '"endpoint":"', vm.toString(address(endpointA)), '",');
        root = string.concat(root, '"dvn":"', vm.toString(address(dvnA)), '",');
        root = string.concat(root, '"adapter":"', vm.toString(address(adapterA)), '",');
        root = string.concat(root, '"aid":"', vm.toString(address(aidA)), '"');
        root = string.concat(root, '},"chainB":{');

        // Serialize for chain B
        root = string.concat(root, '"endpoint":"', vm.toString(address(endpointB)), '",');
        root = string.concat(root, '"dvn":"', vm.toString(address(dvnB)), '",');
        root = string.concat(root, '"adapter":"', vm.toString(address(adapterB)), '",');
        root = string.concat(root, '"aid":"', vm.toString(address(aidB)), '"');

        root = string.concat(root, "}}");

        vm.writeFile(DEPLOYMENT_INFO_FILE, root);

        console.log("\n-----------------------------------------");
        console.log("Full local deployment and setup complete!");
        console.log("Deployment info saved to:", DEPLOYMENT_INFO_FILE);
    }
    
    function _deploySymbioticStack(address _deployer) internal {
        deployer = _deployer; // Set deployer for LocalDeploy context
        SYMBIOTIC_CORE_PROJECT_ROOT = "node_modules/@symbioticfi/core/";
        setupCore();
        setupStakingToken();
        setupNetwork();
        setupKeyRegistry();
        setupVotingPowers();
        setupSettlement();
        setupDriver();
    }

    function _deployLzStack(address _deployer)
        internal
        returns (EndpointV2, SendUln302, ReceiveUlnSymbiotic, SymbioticLzDVN, PriceFeed, Executor)
    {
        EndpointV2 endpoint = new EndpointV2(uint32(block.chainid), _deployer);
        SendUln302 sendLib = new SendUln302(address(endpoint), 200000, 250000);
        ReceiveUlnSymbiotic receiveLib = new ReceiveUlnSymbiotic(address(endpoint));
        PriceFeed priceFeed = new PriceFeed();
        priceFeed.initialize(_deployer);

        address settlementContract = settlements.get(block.chainid);
        SymbioticLzDVN dvn = new SymbioticLzDVN(uint32(block.chainid), address(priceFeed), settlementContract, address(receiveLib));

        // Deploy Executor
        Executor executorImpl = new Executor();
        address[] memory messageLibs = new address[](1);
        messageLibs[0] = address(sendLib);
        address[] memory admins = new address[](1);
        admins[0] = _deployer;
        bytes memory executorInitData = abi.encodeWithSelector(
            Executor.initialize.selector, address(endpoint), address(0x0), messageLibs, address(priceFeed), address(0x0), admins
        );
        ERC1967Proxy executorProxy = new ERC1967Proxy(address(executorImpl), executorInitData);
        Executor executor = Executor(address(executorProxy));

        return (endpoint, sendLib, receiveLib, dvn, priceFeed, executor);
    }
    
    function _deployAppContracts(address _deployer, address _lzEndpoint)
        internal
        returns (AID, Minter, AidOFTAdapter, MockERC20)
    {
        MockERC20 asset = new MockERC20("Mock USDC", "mUSDC");
        asset.mint(_deployer, 1_000_000 * 1e18);

        AID aidImpl = new AID();
        bytes memory aidInitData = abi.encodeWithSelector(AID.initialize.selector, "AID Token", "AID", _deployer, 10);
        ERC1967Proxy aidProxy = new ERC1967Proxy(address(aidImpl), aidInitData);
        AID aid = AID(address(aidProxy));

        Minter minterImpl = new Minter();
        bytes memory minterInitData =
            abi.encodeWithSelector(Minter.initialize.selector, _deployer, address(aid), address(asset));
        ERC1967Proxy minterProxy = new ERC1967Proxy(address(minterImpl), minterInitData);
        Minter minter = Minter(address(minterProxy));

        AidOFTMintBurner mintBurner = new AidOFTMintBurner(address(aid));
        AidOFTAdapter adapter = new AidOFTAdapter(address(aid), mintBurner, _lzEndpoint, _deployer);

        aid.grantRole(aid.MINTER_ROLE(), address(minter));
        aid.grantRole(aid.BRIDGE_MINTER_ROLE(), address(mintBurner));
        mintBurner.setAdapter(address(adapter));
        mintBurner.transferOwnership(address(adapter));

        return (aid, minter, adapter, asset);
    }
    
    function _configureLayerZero(
        EndpointV2 _endpoint,
        SendUln302 _sendLib,
        ReceiveUlnSymbiotic _receiveLib,
        SymbioticLzDVN _dvn,
        PriceFeed _priceFeed,
        Executor _executor,
        AidOFTAdapter _adapter,
        address _peerAdapter,
        uint32 _peerEid
    ) internal {
        // Register libraries
        _endpoint.registerLibrary(address(_sendLib));
        _endpoint.registerLibrary(address(_receiveLib));

        // Set adapter peer
        _adapter.setPeer(_peerEid, bytes32(uint256(uint160(_peerAdapter))));

        address[] memory requiredDvns = new address[](1);
        requiredDvns[0] = address(_dvn);

        // Configure Send Library for Quoting
        UlnConfig memory sendConfig = UlnConfig({
            confirmations: 1,
            requiredDVNCount: 1,
            optionalDVNCount: 0,
            optionalDVNThreshold: 0,
            requiredDVNs: requiredDvns,
            optionalDVNs: new address[](0)
        });
        
        SetConfigParam[] memory sendConfigParam = new SetConfigParam[](2);
        sendConfigParam[0] =
            SetConfigParam({
                eid: _peerEid,
                configType: 1, // CONFIG_TYPE_EXECUTOR
                config: abi.encode(ExecutorConfig({maxMessageSize: 10000, executor: address(_executor)}))
            });
        sendConfigParam[1] = SetConfigParam({eid: _peerEid, configType: 2, config: abi.encode(sendConfig)});
        _endpoint.setConfig(address(_adapter), address(_sendLib), sendConfigParam);

        // Configure Receive Library for Verification
        UlnConfig memory receiveConfig = UlnConfig({
            confirmations: 1,
            requiredDVNCount: 1,
            optionalDVNCount: 0,
            optionalDVNThreshold: 0,
            requiredDVNs: requiredDvns, // Critical: The DVN contract itself is the verifier
            optionalDVNs: new address[](0)
        });
        SetConfigParam[] memory receiveConfigParam = new SetConfigParam[](1);
        receiveConfigParam[0] = SetConfigParam({eid: _peerEid, configType: 2, config: abi.encode(receiveConfig)});
        _endpoint.setConfig(address(_adapter), address(_receiveLib), receiveConfigParam);

        // Set default libraries for the OApp-peer pair
        _endpoint.setSendLibrary(address(_adapter), _peerEid, address(_sendLib));
        _endpoint.setReceiveLibrary(address(_adapter), _peerEid, address(_receiveLib), 0);
    }
}
