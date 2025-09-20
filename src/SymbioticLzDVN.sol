// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {DVN} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/dvn/DVN.sol";
import {ISettlement} from "@symbioticfi/relay-contracts/interfaces/modules/settlement/ISettlement.sol";
import {ReceiveUlnSymbiotic} from "./uln/ReceiveUlnSymbiotic.sol";


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
        bytes calldata _symbioticProof // Contains epoch and proof bytes
    ) external {
        // Step 1: Verify the proof against the Symbiotic Settlement contract
        (uint48 epoch, bytes memory proof) = abi.decode(_symbioticProof, (uint48, bytes));

        // Create a fixed-size message hash to be verified. This must match the message sent by the off-chain worker to the relay.
        bytes32 message = keccak256(abi.encode(_packetHeader, _payloadHash));

        ISettlement settlement = ISettlement(settlementContract);

        // Convert the bytes32 message hash to bytes memory to pass to the verification function.
        bytes memory messageBytes = new bytes(32);
        assembly {
            mstore(add(messageBytes, 32), message)
        }

        // Pass the fixed-size hash (as bytes) to the settlement contract for verification.
        bool success = settlement.verifyQuorumSigAt(
            messageBytes,
            settlement.getRequiredKeyTagFromValSetHeaderAt(epoch),
            settlement.getQuorumThresholdFromValSetHeaderAt(epoch),
            proof,
            epoch,
            new bytes(0)
        );
        require(success, "SymbioticLzDVN: symbiotic verification failed");

        // Step 2: If symbiotic verification is successful, submit the verification to the custom ReceiveUln library.
        ReceiveUlnSymbiotic(receiveUlnContract).verifyWithSymbiotic(
            _packetHeader,
            _payloadHash,
            _confirmations,
            address(this) // Pass this contract's address as the "signer"
        );
    }
}
