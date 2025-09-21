// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {DVN} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/dvn/DVN.sol";
import {ISettlement} from "@symbioticfi/relay-contracts/interfaces/modules/settlement/ISettlement.sol";
import {IReceiveUlnE2} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/interfaces/IReceiveUlnE2.sol";


// This contract extends the standard LayerZero DVN to use Symbiotic's settlement contract for verification,
// instead of the default multi-signature verification.
contract SymbioticLzDVN is DVN {
    address public immutable settlementContract;
    address public immutable receiveUlnContract;

    constructor(
        uint32 _localEid,
        address _priceFeed,
        address _settlementContract,
        address _receiveUlnContract,
        address[] memory _messageLibs,
        address[] memory _signers,
        uint8 _threshold,
        address[] memory _admins
    ) DVN(_localEid, _localEid, _messageLibs, _priceFeed, _signers, _threshold, _admins) {
        settlementContract = _settlementContract;
        receiveUlnContract = _receiveUlnContract;
    }

    // @notice Overrides the standard DVN's verification logic to use Symbiotic's on-chain settlement.
    // The DVN worker will call this function with the packet details and the Symbiotic proof.
    function verifyWithSymbiotic(
        bytes calldata _packetHeader,
        bytes32 _payloadHash,
        uint64 _confirmations,
        uint48 _epoch,
        bytes calldata _proof
    ) external {
        // Step 1: Verify the proof against the Symbiotic Settlement contract
        // The message to be verified must exactly match the message sent by the off-chain worker to the relay.
        // This is the keccak256 hash of the abi.encodePacked packet header and payload hash.
        bytes32 messageHash = keccak256(abi.encodePacked(_packetHeader, _payloadHash));

        ISettlement settlement = ISettlement(settlementContract);

        // The settlement contract's `verifyQuorumSigAt` function expects a `bytes` message,
        // so we pass the hash. `abi.encode` will convert the bytes32 to a 32-byte bytes array.
        bool success = settlement.verifyQuorumSigAt(
            abi.encode(messageHash),
            settlement.getRequiredKeyTagFromValSetHeaderAt(_epoch),
            settlement.getQuorumThresholdFromValSetHeaderAt(_epoch),
            _proof,
            _epoch,
            new bytes(0)
        );
        require(success, "SymbioticLzDVN: symbiotic verification failed");

        // Step 2: If symbiotic verification is successful, submit the verification to the custom ReceiveUln library.
        // The receiveUlnContract will verify that msg.sender is the trusted symbiotic DVN.
        IReceiveUlnE2(receiveUlnContract).verify(_packetHeader, _payloadHash, _confirmations);
    }
}
