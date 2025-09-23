// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {IMintableBurnable} from "@layerzerolabs/oft-evm/contracts/interfaces/IMintableBurnable.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

interface IAID {
    function bridgeMint(address to, uint256 amount) external;
    function bridgeBurn(address from, uint256 amount) external;
}

/// @title Aid OFT Mint Burner Contract
/// @dev inherits Ownable
contract AidOFTMintBurner is Ownable, IMintableBurnable {
    address public immutable aid;
    address public adapter;

    error NotAdapter();

    constructor(address _aid) {
        aid = _aid;
    }

    modifier onlyAdapter() {
        if (msg.sender != adapter) revert NotAdapter();
        _;
    }

    function setAdapter(address _adapter) external onlyOwner {
        adapter = _adapter;
    }

    function mint(address to, uint256 amount) external onlyAdapter returns (bool) {
        IAID(aid).bridgeMint(to, amount);
        return true;
    }

    function burn(address from, uint256 amount) external onlyAdapter returns (bool) {
        IAID(aid).bridgeBurn(from, amount);
        return true;
    }
}
