// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {IWorker} from "@layerzerolabs/lz-evm-protocol-v2/contracts/messagelib/interfaces/IWorker.sol";
import {IDVN} from "@layerzerolabs/lz-evm-protocol-v2/contracts/messagelib/uln/interfaces/IDVN.sol";
import {ILayerZeroPriceFeed} from "@layerzerolabs/lz-evm-protocol-v2/contracts/messagelib/interfaces/ILayerZeroPriceFeed.sol";
import {ISettlement} from "@symbioticfi/relay-contracts/interfaces/modules/settlement/ISettlement.sol";
import {ReceiveUlnBase} from "@layerzerolabs/lz-evm-protocol-v2/contracts/messagelib/uln/ReceiveUlnBase.sol";

contract SymbioticLzDVN is IDVN {
    // =================================================================================================================
    // ================================================== Structs ======================================================
    // =================================================================================================================

    // =================================================================================================================
    // ================================================== Storage ======================================================
    // =================================================================================================================

    uint32 public immutable localEid;
    address public immutable priceFeed;
    address public immutable settlementContract;
    address public immutable receiveUlnContract; // The custom ReceiveUlnSymbiotic contract
    address public immutable owner;

    mapping(uint32 => DstConfig) public dstConfigLookup;

    // =================================================================================================================
    // ================================================== Events =======================================================
    // =================================================================================================================

    // =================================================================================================================
    // ================================================ Modifiers ======================================================
    // =================================================================================================================
    modifier onlyOwner() {
        require(msg.sender == owner, "SymbioticLzDVN: only owner");
        _;
    }

    // =================================================================================================================
    // ================================================== Logic ========================================================
    // =================================================================================================================

    constructor(
        uint32 _localEid,
        address _priceFeed,
        address _settlementContract,
        address _receiveUlnContract
    ) {
        localEid = _localEid;
        priceFeed = _priceFeed;
        settlementContract = _settlementContract;
        receiveUlnContract = _receiveUlnContract;
        owner = msg.sender;
    }

    function setDstConfig(DstConfigParam[] calldata _params) external override onlyOwner {
        for (uint i = 0; i < _params.length; ++i) {
            dstConfigLookup[_params[i].dstEid] = DstConfig({
                gas: _params[i].gas,
                multiplierBps: _params[i].multiplierBps,
                floorMarginUSD: _params[i].floorMarginUSD
            });
        }
        emit SetDstConfig(_params);
    }

    function getFee(uint32 _dstEid, uint64, address, bytes calldata) external view override returns (uint256 nativeFee) {
        DstConfig memory config = dstConfigLookup[_dstEid];
        // Apply multiplier and add floor margin.
        nativeFee = (config.gas * config.multiplierBps) / 10000;
        // Convert floor margin in USD to native fee.
        if (config.floorMarginUSD > 0) {
            uint256 floorMarginNative = ILayerZeroPriceFeed(priceFeed).getPrice(
                localEid,
                config.floorMarginUSD
            );
            nativeFee += floorMarginNative;
        }
    }

    function assignJob(AssignJobParam calldata _param, bytes calldata _options)
        external
        payable
        override
        returns (uint256 fee)
    {
        fee = getFee(_param.dstEid, _param.confirmations, _param.sender, _options);
        require(msg.value >= fee, "SymbioticLzDVN: insufficient fee");
    }

    function verifyWithSymbiotic(
        bytes calldata _packetHeader,
        bytes32 _payloadHash,
        uint64 _confirmations,
        bytes calldata _symbioticProof
    ) external {
        // Step 1: Verify the proof against the Symbiotic Settlement contract
        ISettlement(settlementContract).verify(_symbioticProof);

        // Step 2: If symbiotic verification is successful, submit the verification to the custom ReceiveUln library.
        // We use a low-level call here because the custom ReceiveUlnSymbiotic contract will have the logic to handle this.
        // The custom library will ensure that only this DVN contract can call its special verification function.
        (bool success,) = receiveUlnContract.call(abi.encodeWithSignature(
            "verifyWithSymbiotic(bytes,bytes32,uint64,address)",
            _packetHeader,
            _payloadHash,
            _confirmations,
            address(this) // Pass this contract's address as the "signer"
        ));
        require(success, "SymbioticLzDVN: verification failed in ReceiveUln");
    }

    // =================================================================================================================
    // ================================================== Views ========================================================
    // =================================================================================================================

    function dstConfig(uint32 _dstEid) external view override returns (uint64 gas, uint16 multiplierBps, uint128 floorMarginUSD) {
        DstConfig memory config = dstConfigLookup[_dstEid];
        gas = config.gas;
        multiplierBps = config.multiplierBps;
        floorMarginUSD = config.floorMarginUSD;
    }

    function worker() external pure override returns (address) {
        return address(0);
    }

    function feeLib() external pure override returns (address) {
        return address(0);
    }

    function setWorkerFeeLib(address) external pure override {
        revert("SymbioticLzDVN: not supported");
    }
}
