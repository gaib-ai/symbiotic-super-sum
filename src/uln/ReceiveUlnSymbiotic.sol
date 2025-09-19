// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {PacketV1Codec} from "@layerzerolabs/lz-evm-protocol-v2/contracts/messagelib/libs/PacketV1Codec.sol";
import {SetConfigParam} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/IMessageLibManager.sol";
import {ILayerZeroEndpointV2, Origin} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";
import {IReceiveUlnE2} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/interfaces/IReceiveUlnE2.sol";
import {ReceiveUlnBase, Verification} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/ReceiveUlnBase.sol";
import {ReceiveLibBaseE2} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/ReceiveLibBaseE2.sol";
import {UlnConfig} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/UlnBase.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract ReceiveUlnSymbiotic is IReceiveUlnE2, ReceiveUlnBase, ReceiveLibBaseE2 {
    using PacketV1Codec for bytes;

    uint32 internal constant CONFIG_TYPE_ULN = 2;

    error LZ_ULN_InvalidConfigType(uint32 configType);
    error ReceiveUlnSymbiotic_InvalidDVN();

    constructor(address _endpoint) ReceiveLibBaseE2(_endpoint) Ownable(msg.sender) {}

    function supportsInterface(bytes4 _interfaceId) public view override returns (bool) {
        return _interfaceId == type(IReceiveUlnE2).interfaceId || super.supportsInterface(_interfaceId);
    }

    // ============================ OnlyEndpoint ===================================

    function setConfig(address _oapp, SetConfigParam[] calldata _params) external override onlyEndpoint {
        for (uint256 i = 0; i < _params.length; i++) {
            SetConfigParam calldata param = _params[i];
            _assertSupportedEid(param.eid);
            if (param.configType == CONFIG_TYPE_ULN) {
                _setUlnConfig(param.eid, _oapp, abi.decode(param.config, (UlnConfig)));
            } else {
                revert LZ_ULN_InvalidConfigType(param.configType);
            }
        }
    }

    // ============================ External ===================================

    function commitVerification(bytes calldata _packetHeader, bytes32 _payloadHash) external {
        _assertHeader(_packetHeader, localEid);

        address receiver = _packetHeader.receiverB20();
        uint32 srcEid = _packetHeader.srcEid();

        UlnConfig memory config = getUlnConfig(receiver, srcEid);
        _verifyAndReclaimStorage(config, keccak256(_packetHeader), _payloadHash);

        Origin memory origin = Origin(srcEid, _packetHeader.sender(), _packetHeader.nonce());
        ILayerZeroEndpointV2(endpoint).verify(origin, receiver, _payloadHash);
    }

    // This is the custom verification function called by our SymbioticLzDVN contract
    function verifyWithSymbiotic(
        bytes calldata _packetHeader,
        bytes32 _payloadHash,
        uint64 _confirmations,
        address _dvnContract // The address of the SymbioticLzDVN contract that performed symbiotic verification
    ) external {
        // In the UlnConfig, the `requiredDVNs` list must contain the `_dvnContract` address.
        // This ensures that only a trusted DVN contract that has access to Symbiotic's settlement layer can perform verification.
        UlnConfig memory config = getUlnConfig(_packetHeader.receiverB20(), _packetHeader.srcEid());
        bool isValidDVN = false;
        for (uint i = 0; i < config.requiredDVNCount; i++) {
            if (config.requiredDVNs[i] == _dvnContract) {
                isValidDVN = true;
                break;
            }
        }
        if (!isValidDVN) revert ReceiveUlnSymbiotic_InvalidDVN();
        
        // Use the internal _verify function, which records the verification in hashLookup.
        // Critical: The "verifier" recorded will be the DVN contract's address itself.
        _verify(_packetHeader, _payloadHash, _confirmations, _dvnContract);
    }

    // This standard `verify` function is disabled to force verification through the Symbiotic workflow.
    function verify(bytes calldata, bytes32, uint64) external pure {
        revert("ReceiveUlnSymbiotic: use verifyWithSymbiotic");
    }

    // ============================ Internal ===================================
    
    // Override the internal verify function to accept the dvn address as a parameter
    function _verify(bytes calldata _packetHeader, bytes32 _payloadHash, uint64 _confirmations, address _dvn) internal {
        hashLookup[keccak256(_packetHeader)][_payloadHash][_dvn] = Verification(true, _confirmations);
        emit PayloadVerified(_dvn, _packetHeader, _confirmations, _payloadHash);
    }

    // ============================ View ===================================

    function getConfig(uint32 _eid, address _oapp, uint32 _configType) external view override returns (bytes memory) {
        if (_configType == CONFIG_TYPE_ULN) {
            return abi.encode(getUlnConfig(_oapp, _eid));
        } else {
            revert LZ_ULN_InvalidConfigType(_configType);
        }
    }

    function isSupportedEid(uint32 _eid) external view override returns (bool) {
        return _isSupportedEid(_eid);
    }

    function version() external pure override returns (uint64 major, uint8 minor, uint8 endpointVersion) {
        return (3, 0, 2); // Mimic ULN302 versioning
    }
}
