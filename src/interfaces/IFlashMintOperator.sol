// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

interface IFlashMintOperator {
    /// @notice Executes the flash mint operation
    /// @param amount The amount of tokens to flash mint
    /// @param fee The fee to pay for the flash mint
    /// @param initiator The address that initiated the flash mint
    /// @param data The data to send with the flash mint
    /// @return success Whether the operation was successful
    function executeOperation(uint256 amount, uint256 fee, address initiator, bytes calldata data)
        external
        returns (bool);
}
