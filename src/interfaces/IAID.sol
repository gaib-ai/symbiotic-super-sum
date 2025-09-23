// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

interface IAID {
    function mint(address asset, address to, uint256 amount) external;
    function burn(address asset, address to, uint256 amount) external;
}
