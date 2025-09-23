// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockERC20 is ERC20 {
    constructor(string memory name, string memory symbol) ERC20(name, symbol) {
        _mint(msg.sender, 1_000_000_000_000_000_000_000_000);
    }

    /// @dev Public mint function added for testing and deployment scripts.
    /// The base OpenZeppelin v5 ERC20 contract only has an internal _mint function.
    function mint(address to, uint256 amount) public {
        _mint(to, amount);
    }
}
