// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {UUPSUpgradeable} from "@openzeppelin-v5/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {AccessControlUpgradeable} from "@openzeppelin-v5/contracts-upgradeable/access/AccessControlUpgradeable.sol";

import {IAID} from "./interfaces/IAID.sol";
import {IERC20} from "@openzeppelin-v5/contracts/token/ERC20/IERC20.sol";

/// @title Minter Contract
/// @dev 1 instance for 1 chain, inherits UUPS upgradeable and AccessControl contracts
contract Minter is UUPSUpgradeable, AccessControlUpgradeable {
    // keccak256(abi.encode(uint256(keccak256("gaib.storage.Minter")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant MINTER_STORAGE_LOCATION =
        0x59ac81088878db12231eae56e5ff71d1f9dc58f6cd8d7f21d5972ac426945e00;

    /// @notice Redemption status constants
    uint8 constant REDEMPTION_STATUS_PENDING = 1;
    uint8 constant REDEMPTION_STATUS_COMPLETED = 2;

    /// @notice Redemption request
    struct RedemptionRequest {
        uint8 status;
        address to;
        uint128 amount;
        uint128 timestamp;
    }

    /// @custom:storage-location erc7201:gaib.storage.Minter
    struct MinterStorage {
        /// @notice Address of the AID token
        address aid;
        /// @notice Address of the asset to deposit
        address asset;
        /// @notice Maximum amount of AID that can be minted per block
        uint256 maxMintPerBlock;
        /// @notice Amount of AID minted per block
        mapping(uint256 blockNumber => uint256 amount) mintedPerBlock;
        /// @notice Maximum amount of AID that can be redeemed per block
        uint256 maxRedeemPerBlock;
        /// @notice Amount of AID redeemed per block
        mapping(uint256 blockNumber => uint256 amount) redeemedPerBlock;
        /// @notice Cooldown duration for redemption
        uint24 cooldownDuration;
        /// @notice Redemption requests
        mapping(address user => RedemptionRequest request) redemptionRequests;
    }

    event MaxMintPerBlockSet(uint256 maxAmount);
    event MaxRedeemPerBlockSet(uint256 maxAmount);
    event CooldownDurationSet(uint24 duration);
    event RedemptionInitiated(address user, address to, uint256 amount);
    event RedemptionCompleted(address user, address to, uint256 amount);

    error RedemptionAlreadyInitiated();
    error RedemptionNotInitiated();
    error RedemptionOnCooldown();

    function _getMinterStorage() private pure returns (MinterStorage storage $) {
        assembly {
            $.slot := MINTER_STORAGE_LOCATION
        }
    }

    function initialize(address _admin, address _aid, address _asset) public initializer {
        __AccessControl_init();
        __UUPSUpgradeable_init();
        _grantRole(DEFAULT_ADMIN_ROLE, _admin);
        MinterStorage storage $ = _getMinterStorage();
        $.aid = _aid;
        $.asset = _asset;
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) {}

    /// @notice Mint AID to a given address
    /// @param to The address to mint AID to
    /// @param amount The amount of AID to mint
    /// @dev transfer asset in, calls AID to mint AID out
    /// @return amount of AID minted
    function mint(address to, uint256 amount) external returns (uint256) {
        // check if the amount is greater than the max mint per block
        MinterStorage storage $ = _getMinterStorage();

        IERC20($.asset).transferFrom(msg.sender, address(this), amount);
        IERC20($.asset).approve($.aid, amount);
        IAID($.aid).mint($.asset, to, amount);
        return amount;
    }

    /// @notice Set the maximum amount of AID that can be minted per block
    /// @param maxAmount The maximum amount of AID that can be minted per block
    function setMaxMintPerBlock(uint256 maxAmount) external onlyRole(DEFAULT_ADMIN_ROLE) {
        MinterStorage storage $ = _getMinterStorage();
        $.maxMintPerBlock = maxAmount;
        emit MaxMintPerBlockSet(maxAmount);
    }

    /// @notice Set the maximum amount of AID that can be redeemed per block
    /// @param maxAmount The maximum amount of AID that can be redeemed per block
    function setMaxRedeemPerBlock(uint256 maxAmount) external onlyRole(DEFAULT_ADMIN_ROLE) {
        MinterStorage storage $ = _getMinterStorage();
        $.maxRedeemPerBlock = maxAmount;
        emit MaxRedeemPerBlockSet(maxAmount);
    }

    // cool down on redemption
    /// @notice Initiate a redemption
    /// @param to The address to redeem AID from
    /// @param amount The amount of AID to redeem
    /// @return amount of asset redeemed
    function initiateRedeem(address to, uint128 amount) external returns (uint256) {
        MinterStorage storage $ = _getMinterStorage();
        // will pass if the user has no requets, or the previous request is completed
        if ($.redemptionRequests[msg.sender].status == REDEMPTION_STATUS_PENDING) {
            revert RedemptionAlreadyInitiated();
        }
        $.redemptionRequests[msg.sender] = RedemptionRequest({
            status: REDEMPTION_STATUS_PENDING,
            to: to,
            amount: amount,
            timestamp: uint128(block.timestamp)
        });
        return amount;
    }

    /// @notice Complete a redemption
    /// @dev only callable by the contract itself
    /// @return amount of asset redeemed
    function completeRedeem() external returns (uint256) {
        MinterStorage storage $ = _getMinterStorage();
        RedemptionRequest memory request = $.redemptionRequests[msg.sender];
        if (request.status != REDEMPTION_STATUS_PENDING) {
            revert RedemptionNotInitiated();
        }
        if (block.timestamp < request.timestamp + $.cooldownDuration) {
            revert RedemptionOnCooldown();
        }
        IAID($.aid).burn($.asset, request.to, request.amount); // AID.burn will transfer the asset to the user
        $.redemptionRequests[msg.sender].status = REDEMPTION_STATUS_COMPLETED;
        return request.amount;
    }

    /// @notice Set the cooldown duration for redemption
    /// @param duration The cooldown duration for redemption
    function setCooldownDuration(uint24 duration) external onlyRole(DEFAULT_ADMIN_ROLE) {
        MinterStorage storage $ = _getMinterStorage();
        $.cooldownDuration = duration;
        emit CooldownDurationSet(duration);
    }
}
