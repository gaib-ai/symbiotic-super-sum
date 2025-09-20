// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {Vm} from "forge-std/Vm.sol";
import {console} from "forge-std/console.sol";
import {Script} from "forge-std/Script.sol";

// Symbiotic Contracts
import {ISettlement} from "@symbioticfi/relay-contracts/interfaces/modules/settlement/ISettlement.sol";
// Copied from LocalDeploy.s.sol to support loading existing deployment
import {IValSetDriver} from "@symbioticfi/relay-contracts/interfaces/modules/valset-driver/IValSetDriver.sol";
import {Network} from "@symbioticfi/relay-contracts/modules/network/Network.sol";

// LayerZero Contracts
import {EndpointV2} from "@layerzerolabs/lz-evm-protocol-v2/contracts/EndpointV2.sol";
import {SendUln302} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/uln302/SendUln302.sol";
import {UlnConfig, SetDefaultUlnConfigParam} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/UlnBase.sol";
import {SetConfigParam} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/IMessageLibManager.sol";
import {ExecutorConfig} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/SendLibBase.sol";
import {PriceFeed, ILayerZeroPriceFeed} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/PriceFeed.sol";
import {Executor, IExecutor} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/Executor.sol";
import {ExecutorFeeLib} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/ExecutorFeeLib.sol";
import {DVNFeeLib} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/dvn/DVNFeeLib.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IDVN} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/interfaces/IDVN.sol";
import {ReceiveUln302} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/uln302/ReceiveUln302.sol";

// Custom Integrated Contracts
import {SymbioticLzDVN} from "../src/SymbioticLzDVN.sol";

// Application Contracts
import {AID} from "../src/AID.sol";
import {Minter} from "../src/Minter.sol";
import {AidOFTAdapter} from "../src/OFT/AidOFTAdapter.sol";
import {AidOFTMintBurner} from "../src/OFT/AidOFTMintBurner.sol";
import {MockERC20} from "./mock/MockERC20.sol";
import {EnumerableMap} from "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";

contract SymbioticLzDvnDeploy is Script {
    using EnumerableMap for EnumerableMap.UintToAddressMap;

    // --- State and Structs from LocalDeploy for loading contracts ---
    struct CrossChainAddress {
        address addr;
        uint64 chainId;
    }

    struct RelayContracts {
        CrossChainAddress driver;
        CrossChainAddress keyRegistry;
        address network;
        CrossChainAddress[] settlements;
        CrossChainAddress[] stakingTokens;
        CrossChainAddress[] votingPowerProviders;
    }

    Network internal network;
    IValSetDriver.CrossChainAddress internal keyRegistry;
    IValSetDriver.CrossChainAddress internal driver;
    EnumerableMap.UintToAddressMap internal stakingTokens;
    EnumerableMap.UintToAddressMap internal votingPowerProviders;
    EnumerableMap.UintToAddressMap internal settlements;

    // Local Chain A (e.g., Anvil Fork 1)
    uint32 internal constant localChainA_Eid = 31337;
    string internal localChainA_Rpc = "http://localhost:8545";
    uint256 internal localChainA_PrivateKey;

    // Local Chain B (e.g., Anvil Fork 2)
    uint32 internal constant localChainB_Eid = 31338;
    string internal localChainB_Rpc = "http://localhost:8546";
    uint256 internal localChainB_PrivateKey;
    
    string constant DEPLOYMENT_INFO_FILE = "temp-network/deploy-data/dvn_deployment.json";

    function run() public {
        localChainA_PrivateKey = vm.envOr("PRIVATE_KEY", uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80));
        localChainB_PrivateKey = vm.envOr("PRIVATE_KEY", uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80));

        uint256 forkA = vm.createFork(localChainA_Rpc);
        uint256 forkB = vm.createFork(localChainB_Rpc);

        address deployer = vm.addr(localChainA_PrivateKey);
        console.log("Deploying contracts with the account:", deployer);
        
        // --- Load existing Symbiotic deployment ---
        loadRelayContracts();
        
        // --- Step 1: Deploy on Local Chain A ---
        vm.selectFork(forkA);
        vm.startBroadcast(localChainA_PrivateKey);
        console.log("\n--- Deploying on Local Chain A (%d) ---", block.chainid);
        // Deploy Symbiotic Stack
        // _deploySymbioticStack(deployer);
        // Deploy LayerZero Stack
        (
            EndpointV2 endpointA,
            SendUln302 sendLibA,
            ReceiveUln302 receiveLibA,
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
        // _deploySymbioticStack(deployer);
        // Deploy LayerZero Stack
        (
            EndpointV2 endpointB,
            SendUln302 sendLibB,
            ReceiveUln302 receiveLibB,
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
        
        // --- Step 4: Configure Contracts ---
        console.log("\n--- Configuring Contracts ---");

        // Set the SymbioticDVN address on the ReceiveUlnSymbiotic contracts
        vm.selectFork(forkA);
        vm.startBroadcast(localChainA_PrivateKey);
        console.log("Configuring ReceiveUlnSymbiotic on Chain A...");
        vm.stopBroadcast();

        vm.selectFork(forkB);
        vm.startBroadcast(localChainB_PrivateKey);
        console.log("Configuring ReceiveUlnSymbiotic on Chain B...");
        vm.stopBroadcast();


        // --- Step 5: Write deployment file ---
        string memory root = '{"chainA":{';
        // Serialize for chain A
        root = string.concat(root, '"endpoint":"', vm.toString(address(endpointA)), '",');
        root = string.concat(root, '"receiveUln":"', vm.toString(address(receiveLibA)), '",');
        root = string.concat(root, '"dvn":"', vm.toString(address(dvnA)), '",');
        root = string.concat(root, '"adapter":"', vm.toString(address(adapterA)), '",');
        root = string.concat(root, '"aid":"', vm.toString(address(aidA)), '",');
        root = string.concat(root, '"minter":"', vm.toString(address(minterA)), '",');
        root = string.concat(root, '"asset":"', vm.toString(address(assetA)), '",');
        root = string.concat(root, '"eid":', vm.toString(localChainA_Eid));
        root = string.concat(root, '},"chainB":{');

        // Serialize for chain B
        root = string.concat(root, '"endpoint":"', vm.toString(address(endpointB)), '",');
        root = string.concat(root, '"receiveUln":"', vm.toString(address(receiveLibB)), '",');
        root = string.concat(root, '"dvn":"', vm.toString(address(dvnB)), '",');
        root = string.concat(root, '"adapter":"', vm.toString(address(adapterB)), '",');
        root = string.concat(root, '"aid":"', vm.toString(address(aidB)), '",');
        root = string.concat(root, '"minter":"', vm.toString(address(minterB)), '",');
        root = string.concat(root, '"asset":"', vm.toString(address(assetB)), '",');
        root = string.concat(root, '"eid":', vm.toString(localChainB_Eid));
        root = string.concat(root, "}}");

        vm.writeFile(DEPLOYMENT_INFO_FILE, root);

        console.log("\n-----------------------------------------");
        console.log("Full local deployment and setup complete!");
        console.log("Deployment info saved to:", DEPLOYMENT_INFO_FILE);
    }

    // --- Function from LocalDeploy for loading contracts ---
    function loadRelayContracts() public {
        string memory root = vm.projectRoot();
        string memory path = string.concat(root, "/temp-network/deploy-data/relay_contracts.json");
        string memory json = vm.readFile(path);
        bytes memory data = vm.parseJson(json);
        RelayContracts memory relayContracts = abi.decode(data, (RelayContracts));

        network = Network(payable(relayContracts.network));
        keyRegistry = IValSetDriver.CrossChainAddress({
            chainId: relayContracts.keyRegistry.chainId,
            addr: relayContracts.keyRegistry.addr
        });
        driver =
            IValSetDriver.CrossChainAddress({chainId: relayContracts.driver.chainId, addr: relayContracts.driver.addr});

        // Manual clear since OpenZeppelin v4 EnumerableMap doesn't have a clear() function.
        uint256[] memory stakingTokensKeys = stakingTokens.keys();
        for (uint256 i = 0; i < stakingTokensKeys.length; i++) {
            stakingTokens.remove(stakingTokensKeys[i]);
        }

        for (uint256 i; i < relayContracts.stakingTokens.length; ++i) {
            stakingTokens.set(relayContracts.stakingTokens[i].chainId, relayContracts.stakingTokens[i].addr);
        }

        // Manual clear since OpenZeppelin v4 EnumerableMap doesn't have a clear() function.
        uint256[] memory votingPowerProvidersKeys = votingPowerProviders.keys();
        for (uint256 i = 0; i < votingPowerProvidersKeys.length; i++) {
            votingPowerProviders.remove(votingPowerProvidersKeys[i]);
        }

        for (uint256 i; i < relayContracts.votingPowerProviders.length; ++i) {
            votingPowerProviders.set(
                relayContracts.votingPowerProviders[i].chainId, relayContracts.votingPowerProviders[i].addr
            );
        }

        // Manual clear since OpenZeppelin v4 EnumerableMap doesn't have a clear() function.
        uint256[] memory settlementsKeys = settlements.keys();
        for (uint256 i = 0; i < settlementsKeys.length; i++) {
            settlements.remove(settlementsKeys[i]);
        }

        for (uint256 i; i < relayContracts.settlements.length; ++i) {
            settlements.set(relayContracts.settlements[i].chainId, relayContracts.settlements[i].addr);
        }
    }

    function _deployLzStack(address _deployer)
        internal
        returns (EndpointV2, SendUln302, ReceiveUln302, SymbioticLzDVN, PriceFeed, Executor)
    {
        uint32 eid = uint32(block.chainid);
        EndpointV2 endpoint = new EndpointV2(eid, _deployer);
        SendUln302 sendLib = new SendUln302(address(endpoint), 200000, 250000);
        ReceiveUln302 receiveLib = new ReceiveUln302(address(endpoint));
        PriceFeed priceFeed = new PriceFeed();
        priceFeed.initialize(_deployer);

        address settlementContract = settlements.get(block.chainid);
        // For the base DVN constructor, we provide minimal viable parameters,
        // as the core multi-sig functionality is being replaced by Symbiotic verification.
        address[] memory messageLibs = new address[](1);
        messageLibs[0] = address(sendLib);
        address[] memory signers = new address[](1);
        signers[0] = _deployer; // Placeholder signer
        address[] memory admins = new address[](1);
        admins[0] = _deployer;
        
        SymbioticLzDVN dvn = new SymbioticLzDVN(
            eid,
            address(priceFeed),
            settlementContract,
            address(receiveLib),
            messageLibs,
            signers,
            1, // threshold
            admins
        );
        DVNFeeLib dvnFeeLib = new DVNFeeLib(eid, 10 ** 18);
        dvn.setWorkerFeeLib(address(dvnFeeLib));

        // Deploy Executor
        ExecutorFeeLib executorFeeLib = new ExecutorFeeLib(eid, 10 ** 18);
        Executor executorImpl = new Executor();
        bytes memory executorInitData = abi.encodeWithSelector(
            Executor.initialize.selector, address(endpoint), address(0x0), messageLibs, address(priceFeed), address(0x0), admins
        );
        ERC1967Proxy executorProxy = new ERC1967Proxy(address(executorImpl), executorInitData);
        Executor executor = Executor(address(executorProxy));
        executor.setWorkerFeeLib(address(executorFeeLib));

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
        ReceiveUln302 _receiveLib,
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

        // Set a default ULN configuration for the peer. This also enables the EID.
        SetDefaultUlnConfigParam[] memory defaultConfig = new SetDefaultUlnConfigParam[](1);
        defaultConfig[0] = SetDefaultUlnConfigParam({
            eid: _peerEid,
            config: UlnConfig({
                confirmations: 1,
                requiredDVNCount: 1,
                optionalDVNCount: 0,
                optionalDVNThreshold: 0,
                requiredDVNs: requiredDvns,
                optionalDVNs: new address[](0)
            })
        });
        _sendLib.setDefaultUlnConfigs(defaultConfig);
        _receiveLib.setDefaultUlnConfigs(defaultConfig);

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

        // --- Configure PriceFeed ---
        _priceFeed.setPriceUpdater(vm.addr(localChainA_PrivateKey), true);
        ILayerZeroPriceFeed.UpdatePrice[] memory prices = new ILayerZeroPriceFeed.UpdatePrice[](1);
        prices[0] = ILayerZeroPriceFeed.UpdatePrice({
            eid: _peerEid,
            price: ILayerZeroPriceFeed.Price({
                priceRatio: 1e20,
                gasPriceInUnit: 1e8,
                gasPerByte: 1
            })
        });
        _priceFeed.setPrice(prices);

        // --- Configure DVN ---
        IDVN.DstConfigParam[] memory dvnConfig = new IDVN.DstConfigParam[](1);
        dvnConfig[0] = IDVN.DstConfigParam({
            dstEid: _peerEid,
            gas: 200000,
            multiplierBps: 10000,
            floorMarginUSD: 0
        });
        _dvn.setDstConfig(dvnConfig);

        // --- Configure Executor ---
        IExecutor.DstConfigParam[] memory execConfig = new IExecutor.DstConfigParam[](1);
        execConfig[0] = IExecutor.DstConfigParam({
            dstEid: _peerEid,
            lzReceiveBaseGas: 200000,
            multiplierBps: 10000,
            floorMarginUSD: 0,
            nativeCap: 1 ether,
            lzComposeBaseGas: 0
        });
        _executor.setDstConfig(execConfig);
    }
}
