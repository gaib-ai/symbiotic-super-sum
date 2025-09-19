// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {MintBurnOFTAdapter} from "@layerzerolabs/oft-evm/contracts/MintBurnOFTAdapter.sol";
import {IMintableBurnable} from "@layerzerolabs/oft-evm/contracts/interfaces/IMintableBurnable.sol";

/// @title Aid OFT Adapter Contract
/// @dev inherits MintBurnOFTAdapter
contract AidOFTAdapter is MintBurnOFTAdapter {
    constructor(address _token, IMintableBurnable _minterBurner, address _lzEndpoint, address _owner)
        MintBurnOFTAdapter(_token, _minterBurner, _lzEndpoint, _owner)
    // REMOVED: Ownable(_owner)
    {
        // In OZ v4, owner is set to msg.sender by default.
        // If _owner is different, we need to transfer ownership.
        if (_owner != msg.sender) {
            transferOwnership(_owner);
        }
    }
}
