// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {MintBurnOFTAdapter} from "@layerzerolabs/oft-evm/contracts/MintBurnOFTAdapter.sol";
import {IMintableBurnable} from "@layerzerolabs/oft-evm/contracts/interfaces/IMintableBurnable.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/// @title Aid OFT Adapter Contract
/// @dev inherits MintBurnOFTAdapter
contract AidOFTAdapter is MintBurnOFTAdapter {
    constructor(address _token, IMintableBurnable _minterBurner, address _lzEndpoint, address _owner)
        MintBurnOFTAdapter(_token, _minterBurner, _lzEndpoint, _owner)
        Ownable(_owner)
    {}
}
