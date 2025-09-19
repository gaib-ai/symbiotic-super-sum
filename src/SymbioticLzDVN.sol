// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {IWorker} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/interfaces/IWorker.sol";
import {IDVN} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/interfaces/IDVN.sol";
import {ILayerZeroPriceFeed} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/interfaces/ILayerZeroPriceFeed.sol";
import {ISettlement} from "@symbioticfi/relay-contracts/interfaces/modules/settlement/ISettlement.sol";
import {ReceiveUlnBase} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/ReceiveUlnBase.sol";

contract SymbioticLzDVN is IDVN {
    // =================================================================================================================
    // ================================================== Structs ======================================================
    // =================================================================================================================
    struct SymbioticProofData {
        uint48 epoch;
        bytes proof;
    }
    // =================================================================================================================
    // ================================================== Storage ======================================================
    // =================================================================================================================

    uint32 public immutable localEid;
    address public immutable override priceFeed;
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

    function setDstConfig(DstConfigParam[] calldata _params) external onlyOwner {
        for (uint i = 0; i < _params.length; ++i) {
            dstConfigLookup[_params[i].dstEid] = DstConfig({
                gas: _params[i].gas,
                multiplierBps: _params[i].multiplierBps,
                floorMarginUSD: _params[i].floorMarginUSD
            });
        }
        emit SetDstConfig(_params);
    }

    function getFee(uint32 _dstEid, uint64 _confirmations, address _sender, bytes calldata _options) external view override returns (uint256 nativeFee) {
        DstConfig memory config = dstConfigLookup[_dstEid];
        // Apply multiplier and add floor margin.
        nativeFee = (config.gas * config.multiplierBps) / 10000;
        // Convert floor margin in USD to native fee.
        if (config.floorMarginUSD > 0) {
            // Assumes floorMarginUSD is in USD with 8 decimals.
            // nativeTokenPriceUSD is in USD with 20 decimals.
            // nativeFee is in wei (18 decimals).
            uint256 nativePriceUSD = ILayerZeroPriceFeed(priceFeed).nativeTokenPriceUSD();
            if (nativePriceUSD > 0) {
                uint256 floorMarginNative = (config.floorMarginUSD * 1e18 * 1e20) / (nativePriceUSD * 1e8);
                nativeFee += floorMarginNative;
            }
        }
    }

    function getFee(
        address,
        bytes calldata,
        bytes calldata,
        bytes calldata
    ) external view override returns (uint256) {
        revert("SymbioticLzDVN: not supported");
    }

    function assignJob(AssignJobParam calldata _param, bytes calldata _options)
        external
        payable
        override
        returns (uint256 fee)
    {
        fee = this.getFee(_param.dstEid, _param.confirmations, _param.sender, _options);
        require(msg.value >= fee, "SymbioticLzDVN: insufficient fee");
    }

    function assignJob(address, bytes calldata, bytes calldata, bytes calldata) external payable override returns (uint256) {
        revert("SymbioticLzDVN: not supported");
    }

    function verifyWithSymbiotic(
        bytes calldata _packetHeader,
        bytes32 _payloadHash,
        uint64 _confirmations,
        bytes calldata _symbioticProof
    ) external {
        // Step 1: Verify the proof against the Symbiotic Settlement contract
        SymbioticProofData memory proofData = abi.decode(_symbioticProof, (SymbioticProofData));
        bytes32 messageHash = keccak256(abi.encodePacked(keccak256(abi.encodePacked(_packetHeader, _payloadHash))));

        ISettlement settlement = ISettlement(settlementContract);
        bool success = settlement.verifyQuorumSigAt(
            abi.encode(messageHash),
            settlement.getRequiredKeyTagFromValSetHeaderAt(proofData.epoch),
            settlement.getQuorumThresholdFromValSetHeaderAt(proofData.epoch),
            proofData.proof,
            proofData.epoch,
            new bytes(0)
        );
        require(success, "SymbioticLzDVN: symbiotic verification failed");


        // Step 2: If symbiotic verification is successful, submit the verification to the custom ReceiveUln library.
        // We use a low-level call here because the custom ReceiveUlnSymbiotic contract will have the logic to handle this.
        // The custom library will ensure that only this DVN contract can call its special verification function.
        (success, ) = receiveUlnContract.call(abi.encodeWithSignature(
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

    function worker() external pure returns (address) {
        return address(0);
    }

    function feeLib() external pure returns (address) {
        return address(0);
    }

    function setWorkerFeeLib(address) external pure {
        revert("SymbioticLzDVN: not supported");
    }

    // IWorker functions
    function setPriceFeed(address) external override {
        revert("SymbioticLzDVN: not supported");
    }

    function setDefaultMultiplierBps(uint16) external override {
        revert("SymbioticLzDVN: not supported");
    }

    function defaultMultiplierBps() external view override returns (uint16) {
        revert("SymbioticLzDVN: not supported");
    }

    function withdrawFee(address, address, uint256) external override {
        revert("SymbioticLzDVN: not supported");
    }

    function setSupportedOptionTypes(uint32, uint8[] calldata) external override {
        revert("SymbioticLzDVN: not supported");
    }

    function getSupportedOptionTypes(uint32) external view override returns (uint8[] memory) {
        revert("SymbioticLzDVN: not supported");
    }
}
