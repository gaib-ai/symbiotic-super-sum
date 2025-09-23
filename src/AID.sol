// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {ERC20PermitUpgradeable} from
    "@openzeppelin-v5/contracts-upgradeable/token/ERC20/extensions/ERC20PermitUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin-v5/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {AccessControlUpgradeable} from "@openzeppelin-v5/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin-v5/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin-v5/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";

import {IERC20} from "@openzeppelin-v5/contracts/token/ERC20/IERC20.sol";

import {IFlashMintOperator} from "./interfaces/IFlashMintOperator.sol";

/// @title AID Token Contract
/// @dev 1 instance for 1 chain, inherits ERC20Permit, AccessControl, Pausable, and UUPS upgradeable contracts
///      Admin can add/remove minters and deny list guards using OpenZeppelin AccessControl
/// @notice This contract implements an upgradeable ERC20 token with minting, burning, flash minting,
///         and deny list functionality for enhanced security and control
contract AID is UUPSUpgradeable, ERC20PermitUpgradeable, AccessControlUpgradeable, PausableUpgradeable, ReentrancyGuardUpgradeable {
    // keccak256(abi.encode(uint256(keccak256("gaib.storage.AID")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant AID_STORAGE_LOCATION = 0x1a593d1faecac64eb23535236ee89bd031a16d8788dbfcfe3830e7f3a466c300;

    /// @notice Role identifier for minters who can mint and burn tokens
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    /// @notice Role identifier for deny list guards who can manage the deny list
    bytes32 public constant DENY_LIST_GUARD_ROLE = keccak256("DENY_LIST_GUARD_ROLE");

    /// @notice Role identifier for denied addresses
    bytes32 private constant DENIED_ADDRESS_ROLE = keccak256("DENIED_ADDRESS_ROLE");

    /// @notice Role identifier for bridge minters who can mint tokens
    bytes32 public constant BRIDGE_MINTER_ROLE = keccak256("BRIDGE_MINTER_ROLE");

    /// @custom:storage-location erc7201:gaib.storage.AID
    struct AIDStorage {
        /// @notice Flash mint fee rate
        uint16 flashMintFeeRateBips; // 10000 = 100%
    }

    /// @notice Event emitted when addresses are added to the deny list
    /// @param accounts Array of addresses added to deny list
    event AddedToDenyList(address[] accounts);

    /// @notice Event emitted when addresses are removed from the deny list
    /// @param accounts Array of addresses removed from deny list
    event RemovedFromDenyList(address[] accounts);

    /// @notice Event emitted when flash mint is executed
    /// @param to Address that received the flash minted tokens
    /// @param amount Amount of tokens flash minted
    event FlashMinted(address indexed to, uint256 amount);

    /// @notice Custom error for when an address is denied
    /// @param account The denied address
    error AddressDenied(address account);

    /// @notice Custom error for when an execution fails
    /// @param result The result of the failed execution
    error ExecutionFailed(bytes result);

    /// @notice Custom error for when flash mint fails
    error FlashMintFailed();

    /// @notice Gets the AID storage
    /// @return $ The AID storage
    function _getAIDStorage() private pure returns (AIDStorage storage $) {
        assembly {
            $.slot := AID_STORAGE_LOCATION
        }
    }

    /// @notice Initializes the AID token contract
    /// @param _name The name of the token
    /// @param _symbol The symbol of the token
    /// @param _admin The address that will be granted the DEFAULT_ADMIN_ROLE
    /// @dev This function can only be called once during contract deployment
    function initialize(string memory _name, string memory _symbol, address _admin, uint16 flashMintFeeRateBips)
        public
        initializer
    {
        __ERC20_init(_name, _symbol);
        __ERC20Permit_init(_name);
        __AccessControl_init();
        _grantRole(DEFAULT_ADMIN_ROLE, _admin);
        __Pausable_init();
        __UUPSUpgradeable_init();
        AIDStorage storage $ = _getAIDStorage();
        $.flashMintFeeRateBips = flashMintFeeRateBips;
    }

    /// @notice Authorizes contract upgrades
    /// @param newImplementation The address of the new implementation contract
    /// @dev Only accounts with DEFAULT_ADMIN_ROLE can authorize upgrades
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) {}

    /// @notice Mints AID tokens to a given address
    /// @param asset The address of the asset to mint AID tokens from
    /// @param to The address to mint AID tokens to
    /// @param amount The amount of AID tokens to mint
    /// @dev Only accounts with MINTER_ROLE can mint tokens. Assuming MINTER can be trusted
    /// @dev Reverts if the contract is paused
    function mint(address asset, address to, uint256 amount) public onlyRole(MINTER_ROLE) whenNotPaused {
        IERC20(asset).transferFrom(msg.sender, address(this), amount);
        _mint(to, amount);
    }

    /// @notice Burns AID tokens from a given address
    /// @param asset The address of the asset to burn AID tokens to
    /// @param to The address to transfer the asset to
    /// @param amount The amount of AID tokens to burn
    /// @dev Only accounts with MINTER_ROLE can burn tokens
    /// @dev Reverts if the contract is paused
    function burn(address asset, address to, uint256 amount) public onlyRole(MINTER_ROLE) whenNotPaused {
        _burn(msg.sender, amount);
        IERC20(asset).transfer(to, amount);
    }

    /// @notice Mints AID tokens to a given address upon receiving a message from bridge
    /// @param to The address to mint AID tokens to
    /// @param amount The amount of AID tokens to mint
    /// @dev Only accounts with BRIDGE_MINTER_ROLE can mint tokens
    /// @dev Reverts if the contract is paused
    function bridgeMint(address to, uint256 amount) public onlyRole(BRIDGE_MINTER_ROLE) whenNotPaused {
        _mint(to, amount);
    }

    /// @notice Burns AID tokens from a given address upon receiving a message from bridge
    /// @param from The address to burn AID tokens from
    /// @param amount The amount of AID tokens to burn
    /// @dev Only accounts with BRIDGE_MINTER_ROLE can burn tokens
    /// @dev Reverts if the contract is paused
    function bridgeBurn(address from, uint256 amount) public onlyRole(BRIDGE_MINTER_ROLE) whenNotPaused {
        _burn(from, amount);
    }

    /// @notice Flash mints AID tokens to a given address
    /// @param to The address to flash mint AID tokens to
    /// @param amount The amount of AID tokens to flash mint
    /// @param data The data to send with the flash mint
    /// @dev Mints tokens to the specified address and expects them to be burned by the end of the call
    /// @dev Reverts if the contract is paused or if the address is denied
    /// @dev Reverts if the flash mint callback fails (tokens not burned by end of call)
    function flashMint(address to, uint256 amount, bytes calldata data) public whenNotPaused nonReentrant {
        _mint(to, amount);
        uint256 fee = _getFlashMintFee(amount);
        bool success = IFlashMintOperator(to).executeOperation(amount, fee, msg.sender, data);
        if (!success) {
            revert FlashMintFailed();
        }
        _burn(to, amount + fee);
        emit FlashMinted(to, amount);
    }

    /// @notice Adds addresses to the deny list
    /// @param accounts Array of addresses to add to the deny list
    /// @dev Only accounts with DENY_LIST_GUARD_ROLE can add addresses to deny list
    function addToDenyList(address[] calldata accounts) public onlyRole(DENY_LIST_GUARD_ROLE) {
        for (uint256 i = 0; i < accounts.length; i++) {
            _grantRole(DENIED_ADDRESS_ROLE, accounts[i]);
        }
        emit AddedToDenyList(accounts);
    }

    /// @notice Removes addresses from the deny list
    /// @param accounts Array of addresses to remove from the deny list
    /// @dev Only accounts with DENY_LIST_GUARD_ROLE can remove addresses from deny list
    function removeFromDenyList(address[] calldata accounts) public onlyRole(DENY_LIST_GUARD_ROLE) {
        for (uint256 i = 0; i < accounts.length; i++) {
            _revokeRole(DENIED_ADDRESS_ROLE, accounts[i]);
        }
        emit RemovedFromDenyList(accounts);
    }

    /// @notice Pauses all token operations
    /// @dev Only accounts with DEFAULT_ADMIN_ROLE can pause the contract
    function pause() public onlyRole(DEFAULT_ADMIN_ROLE) {
        _pause();
    }

    /// @notice Unpauses all token operations
    /// @dev Only accounts with DEFAULT_ADMIN_ROLE can unpause the contract
    function unpause() public onlyRole(DEFAULT_ADMIN_ROLE) {
        _unpause();
    }

    /// @notice Override of _update to check deny list
    /// @param from The address transferring tokens from
    /// @param to The address transferring tokens to
    /// @param amount The amount of tokens being transferred
    /// @dev Reverts if the contract is paused or if either address is on the deny list
    function _update(address from, address to, uint256 amount) internal override {
        if (hasRole(DENIED_ADDRESS_ROLE, from)) {
            revert AddressDenied(from);
        } else if (hasRole(DENIED_ADDRESS_ROLE, to)) {
            revert AddressDenied(to);
        }

        super._update(from, to, amount);
    }

    /// @param to The address to execute the call to
    /// @param value The value to send with the call
    /// @param data The data to send with the call
    /// @dev Only accounts with DEFAULT_ADMIN_ROLE can execute calls
    /// @dev Reverts if the call fails
    function execute(address to, uint256 value, bytes calldata data)
        public
        payable
        onlyRole(DEFAULT_ADMIN_ROLE)
        returns (bytes memory)
    {
        (bool success, bytes memory result) = to.call{value: value}(data);
        if (!success) {
            revert ExecutionFailed(result);
        }
        return result;
    }

    /// @notice Sets the flash mint fee rate
    /// @param rate The flash mint fee rate
    /// @dev Only accounts with DEFAULT_ADMIN_ROLE can set the flash mint fee rate
    function setFlashMintFeeRate(uint16 rate) public onlyRole(DEFAULT_ADMIN_ROLE) {
        AIDStorage storage $ = _getAIDStorage();
        $.flashMintFeeRateBips = rate;
    }

    /// @notice Gets the flash mint fee for a given amount
    /// @param amount The amount of tokens to flash mint
    /// @return fee The flash mint fee
    function _getFlashMintFee(uint256 amount) internal view returns (uint256) {
        AIDStorage storage $ = _getAIDStorage();
        return (amount * $.flashMintFeeRateBips) / 10000;
    }
}
