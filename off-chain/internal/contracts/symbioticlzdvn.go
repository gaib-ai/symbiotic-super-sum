// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ExecuteParam is an auto generated low-level Go binding around an user-defined struct.
type ExecuteParam struct {
	Vid        uint32
	Target     common.Address
	CallData   []byte
	Expiration *big.Int
	Signatures []byte
}

// IDVNDstConfigParam is an auto generated low-level Go binding around an user-defined struct.
type IDVNDstConfigParam struct {
	DstEid         uint32
	Gas            uint64
	MultiplierBps  uint16
	FloorMarginUSD *big.Int
}

// ILayerZeroDVNAssignJobParam is an auto generated low-level Go binding around an user-defined struct.
type ILayerZeroDVNAssignJobParam struct {
	DstEid        uint32
	PacketHeader  []byte
	PayloadHash   [32]byte
	Confirmations uint64
	Sender        common.Address
}

// SymbioticLzDVNMetaData contains all meta data concerning the SymbioticLzDVN contract.
var SymbioticLzDVNMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_localEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_priceFeed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_settlementContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiveUlnContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_messageLibs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_admins\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowlistSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"assignJob\",\"inputs\":[{\"name\":\"_param\",\"type\":\"tuple\",\"internalType\":\"structILayerZeroDVN.AssignJobParam\",\"components\":[{\"name\":\"dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"packetHeader\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"payloadHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"_options\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"totalFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"assignJob\",\"inputs\":[{\"name\":\"_dstEid\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"totalFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"assignJob\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_cmd\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_options\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"defaultMultiplierBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dstConfig\",\"inputs\":[{\"name\":\"dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"gas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"multiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"floorMarginUSD\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"_params\",\"type\":\"tuple[]\",\"internalType\":\"structExecuteParam[]\",\"components\":[{\"name\":\"vid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signatures\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"_dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_options\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"_dstEid\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_cmd\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_options\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSigners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedOptionTypes\",\"inputs\":[{\"name\":\"_eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"_role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasAcl\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashCallData\",\"inputs\":[{\"name\":\"_vid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_expiration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isSigner\",\"inputs\":[{\"name\":\"_signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"localEidV2\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"priceFeed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumChangeAdmin\",\"inputs\":[{\"name\":\"_param\",\"type\":\"tuple\",\"internalType\":\"structExecuteParam\",\"components\":[{\"name\":\"vid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signatures\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"receiveUlnContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"_role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaultMultiplierBps\",\"inputs\":[{\"name\":\"_multiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDstConfig\",\"inputs\":[{\"name\":\"_params\",\"type\":\"tuple[]\",\"internalType\":\"structIDVN.DstConfigParam[]\",\"components\":[{\"name\":\"dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"multiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"floorMarginUSD\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPaused\",\"inputs\":[{\"name\":\"_paused\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceFeed\",\"inputs\":[{\"name\":\"_priceFeed\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setQuorum\",\"inputs\":[{\"name\":\"_quorum\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSigner\",\"inputs\":[{\"name\":\"_signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSupportedOptionTypes\",\"inputs\":[{\"name\":\"_eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_optionTypes\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWorkerFeeLib\",\"inputs\":[{\"name\":\"_workerFeeLib\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settlementContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"signerSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"signers\",\"inputs\":[{\"name\":\"_signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usedHashes\",\"inputs\":[{\"name\":\"executableHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"used\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifySignatures\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signatures\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumMultiSig.Errors\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyWithSymbiotic\",\"inputs\":[{\"name\":\"_packetHeader\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_payloadHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_symbioticProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vid\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFee\",\"inputs\":[{\"name\":\"_lib\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeFromUlnV2\",\"inputs\":[{\"name\":\"_lib\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"workerFeeLib\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecuteFailed\",\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HashAlreadyUsed\",\"inputs\":[{\"name\":\"param\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structExecuteParam\",\"components\":[{\"name\":\"vid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signatures\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_hash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDefaultMultiplierBps\",\"inputs\":[{\"name\":\"multiplierBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDstConfig\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIDVN.DstConfigParam[]\",\"components\":[{\"name\":\"dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"multiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"floorMarginUSD\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetPriceFeed\",\"inputs\":[{\"name\":\"priceFeed\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetSupportedOptionTypes\",\"inputs\":[{\"name\":\"dstEid\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"optionTypes\",\"type\":\"uint8[]\",\"indexed\":false,\"internalType\":\"uint8[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetWorkerLib\",\"inputs\":[{\"name\":\"workerLib\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateQuorum\",\"inputs\":[{\"name\":\"_quorum\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateSigner\",\"inputs\":[{\"name\":\"_signer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_active\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifierFeePaid\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifySignaturesFailed\",\"inputs\":[{\"name\":\"idx\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"lib\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DVN_DuplicatedHash\",\"inputs\":[{\"name\":\"executableHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"DVN_InstructionExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DVN_InvalidRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"DVN_InvalidSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DVN_InvalidTarget\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DVN_InvalidVid\",\"inputs\":[{\"name\":\"vid\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"DVN_OnlySelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MultiSig_InvalidSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MultiSig_OnlySigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MultiSig_QuorumIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MultiSig_SignersSizeIsLessThanQuorum\",\"inputs\":[{\"name\":\"signersSize\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorum\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MultiSig_StateAlreadySet\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"MultiSig_StateNotSet\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"MultiSig_UnorderedSigners\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Transfer_NativeFailed\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Transfer_ToAddressIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Worker_NotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Worker_OnlyMessageLib\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Worker_RoleRenouncingDisabled\",\"inputs\":[]}]",
}

// SymbioticLzDVNABI is the input ABI used to generate the binding from.
// Deprecated: Use SymbioticLzDVNMetaData.ABI instead.
var SymbioticLzDVNABI = SymbioticLzDVNMetaData.ABI

// SymbioticLzDVN is an auto generated Go binding around an Ethereum contract.
type SymbioticLzDVN struct {
	SymbioticLzDVNCaller     // Read-only binding to the contract
	SymbioticLzDVNTransactor // Write-only binding to the contract
	SymbioticLzDVNFilterer   // Log filterer for contract events
}

// SymbioticLzDVNCaller is an auto generated read-only Go binding around an Ethereum contract.
type SymbioticLzDVNCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SymbioticLzDVNTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SymbioticLzDVNTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SymbioticLzDVNFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SymbioticLzDVNFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SymbioticLzDVNSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SymbioticLzDVNSession struct {
	Contract     *SymbioticLzDVN   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SymbioticLzDVNCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SymbioticLzDVNCallerSession struct {
	Contract *SymbioticLzDVNCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SymbioticLzDVNTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SymbioticLzDVNTransactorSession struct {
	Contract     *SymbioticLzDVNTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SymbioticLzDVNRaw is an auto generated low-level Go binding around an Ethereum contract.
type SymbioticLzDVNRaw struct {
	Contract *SymbioticLzDVN // Generic contract binding to access the raw methods on
}

// SymbioticLzDVNCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SymbioticLzDVNCallerRaw struct {
	Contract *SymbioticLzDVNCaller // Generic read-only contract binding to access the raw methods on
}

// SymbioticLzDVNTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SymbioticLzDVNTransactorRaw struct {
	Contract *SymbioticLzDVNTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSymbioticLzDVN creates a new instance of SymbioticLzDVN, bound to a specific deployed contract.
func NewSymbioticLzDVN(address common.Address, backend bind.ContractBackend) (*SymbioticLzDVN, error) {
	contract, err := bindSymbioticLzDVN(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVN{SymbioticLzDVNCaller: SymbioticLzDVNCaller{contract: contract}, SymbioticLzDVNTransactor: SymbioticLzDVNTransactor{contract: contract}, SymbioticLzDVNFilterer: SymbioticLzDVNFilterer{contract: contract}}, nil
}

// NewSymbioticLzDVNCaller creates a new read-only instance of SymbioticLzDVN, bound to a specific deployed contract.
func NewSymbioticLzDVNCaller(address common.Address, caller bind.ContractCaller) (*SymbioticLzDVNCaller, error) {
	contract, err := bindSymbioticLzDVN(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNCaller{contract: contract}, nil
}

// NewSymbioticLzDVNTransactor creates a new write-only instance of SymbioticLzDVN, bound to a specific deployed contract.
func NewSymbioticLzDVNTransactor(address common.Address, transactor bind.ContractTransactor) (*SymbioticLzDVNTransactor, error) {
	contract, err := bindSymbioticLzDVN(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNTransactor{contract: contract}, nil
}

// NewSymbioticLzDVNFilterer creates a new log filterer instance of SymbioticLzDVN, bound to a specific deployed contract.
func NewSymbioticLzDVNFilterer(address common.Address, filterer bind.ContractFilterer) (*SymbioticLzDVNFilterer, error) {
	contract, err := bindSymbioticLzDVN(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNFilterer{contract: contract}, nil
}

// bindSymbioticLzDVN binds a generic wrapper to an already deployed contract.
func bindSymbioticLzDVN(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SymbioticLzDVNMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SymbioticLzDVN *SymbioticLzDVNRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SymbioticLzDVN.Contract.SymbioticLzDVNCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SymbioticLzDVN *SymbioticLzDVNRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SymbioticLzDVNTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SymbioticLzDVN *SymbioticLzDVNRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SymbioticLzDVNTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SymbioticLzDVN *SymbioticLzDVNCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SymbioticLzDVN.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SymbioticLzDVN *SymbioticLzDVNTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SymbioticLzDVN *SymbioticLzDVNTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SymbioticLzDVN *SymbioticLzDVNSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SymbioticLzDVN.Contract.DEFAULTADMINROLE(&_SymbioticLzDVN.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SymbioticLzDVN.Contract.DEFAULTADMINROLE(&_SymbioticLzDVN.CallOpts)
}

// AllowlistSize is a free data retrieval call binding the contract method 0xd2ae2104.
//
// Solidity: function allowlistSize() view returns(uint64)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) AllowlistSize(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "allowlistSize")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// AllowlistSize is a free data retrieval call binding the contract method 0xd2ae2104.
//
// Solidity: function allowlistSize() view returns(uint64)
func (_SymbioticLzDVN *SymbioticLzDVNSession) AllowlistSize() (uint64, error) {
	return _SymbioticLzDVN.Contract.AllowlistSize(&_SymbioticLzDVN.CallOpts)
}

// AllowlistSize is a free data retrieval call binding the contract method 0xd2ae2104.
//
// Solidity: function allowlistSize() view returns(uint64)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) AllowlistSize() (uint64, error) {
	return _SymbioticLzDVN.Contract.AllowlistSize(&_SymbioticLzDVN.CallOpts)
}

// DefaultMultiplierBps is a free data retrieval call binding the contract method 0x00bf2e80.
//
// Solidity: function defaultMultiplierBps() view returns(uint16)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) DefaultMultiplierBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "defaultMultiplierBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DefaultMultiplierBps is a free data retrieval call binding the contract method 0x00bf2e80.
//
// Solidity: function defaultMultiplierBps() view returns(uint16)
func (_SymbioticLzDVN *SymbioticLzDVNSession) DefaultMultiplierBps() (uint16, error) {
	return _SymbioticLzDVN.Contract.DefaultMultiplierBps(&_SymbioticLzDVN.CallOpts)
}

// DefaultMultiplierBps is a free data retrieval call binding the contract method 0x00bf2e80.
//
// Solidity: function defaultMultiplierBps() view returns(uint16)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) DefaultMultiplierBps() (uint16, error) {
	return _SymbioticLzDVN.Contract.DefaultMultiplierBps(&_SymbioticLzDVN.CallOpts)
}

// DstConfig is a free data retrieval call binding the contract method 0x9e944965.
//
// Solidity: function dstConfig(uint32 dstEid) view returns(uint64 gas, uint16 multiplierBps, uint128 floorMarginUSD)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) DstConfig(opts *bind.CallOpts, dstEid uint32) (struct {
	Gas            uint64
	MultiplierBps  uint16
	FloorMarginUSD *big.Int
}, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "dstConfig", dstEid)

	outstruct := new(struct {
		Gas            uint64
		MultiplierBps  uint16
		FloorMarginUSD *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Gas = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.MultiplierBps = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.FloorMarginUSD = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DstConfig is a free data retrieval call binding the contract method 0x9e944965.
//
// Solidity: function dstConfig(uint32 dstEid) view returns(uint64 gas, uint16 multiplierBps, uint128 floorMarginUSD)
func (_SymbioticLzDVN *SymbioticLzDVNSession) DstConfig(dstEid uint32) (struct {
	Gas            uint64
	MultiplierBps  uint16
	FloorMarginUSD *big.Int
}, error) {
	return _SymbioticLzDVN.Contract.DstConfig(&_SymbioticLzDVN.CallOpts, dstEid)
}

// DstConfig is a free data retrieval call binding the contract method 0x9e944965.
//
// Solidity: function dstConfig(uint32 dstEid) view returns(uint64 gas, uint16 multiplierBps, uint128 floorMarginUSD)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) DstConfig(dstEid uint32) (struct {
	Gas            uint64
	MultiplierBps  uint16
	FloorMarginUSD *big.Int
}, error) {
	return _SymbioticLzDVN.Contract.DstConfig(&_SymbioticLzDVN.CallOpts, dstEid)
}

// GetFee is a free data retrieval call binding the contract method 0x30bb3aac.
//
// Solidity: function getFee(uint32 _dstEid, uint64 _confirmations, address _sender, bytes _options) view returns(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) GetFee(opts *bind.CallOpts, _dstEid uint32, _confirmations uint64, _sender common.Address, _options []byte) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "getFee", _dstEid, _confirmations, _sender, _options)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0x30bb3aac.
//
// Solidity: function getFee(uint32 _dstEid, uint64 _confirmations, address _sender, bytes _options) view returns(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNSession) GetFee(_dstEid uint32, _confirmations uint64, _sender common.Address, _options []byte) (*big.Int, error) {
	return _SymbioticLzDVN.Contract.GetFee(&_SymbioticLzDVN.CallOpts, _dstEid, _confirmations, _sender, _options)
}

// GetFee is a free data retrieval call binding the contract method 0x30bb3aac.
//
// Solidity: function getFee(uint32 _dstEid, uint64 _confirmations, address _sender, bytes _options) view returns(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) GetFee(_dstEid uint32, _confirmations uint64, _sender common.Address, _options []byte) (*big.Int, error) {
	return _SymbioticLzDVN.Contract.GetFee(&_SymbioticLzDVN.CallOpts, _dstEid, _confirmations, _sender, _options)
}

// GetFee0 is a free data retrieval call binding the contract method 0x5553fb8e.
//
// Solidity: function getFee(uint16 _dstEid, uint16 , uint64 _confirmations, address _sender) view returns(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) GetFee0(opts *bind.CallOpts, _dstEid uint16, arg1 uint16, _confirmations uint64, _sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "getFee0", _dstEid, arg1, _confirmations, _sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee0 is a free data retrieval call binding the contract method 0x5553fb8e.
//
// Solidity: function getFee(uint16 _dstEid, uint16 , uint64 _confirmations, address _sender) view returns(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNSession) GetFee0(_dstEid uint16, arg1 uint16, _confirmations uint64, _sender common.Address) (*big.Int, error) {
	return _SymbioticLzDVN.Contract.GetFee0(&_SymbioticLzDVN.CallOpts, _dstEid, arg1, _confirmations, _sender)
}

// GetFee0 is a free data retrieval call binding the contract method 0x5553fb8e.
//
// Solidity: function getFee(uint16 _dstEid, uint16 , uint64 _confirmations, address _sender) view returns(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) GetFee0(_dstEid uint16, arg1 uint16, _confirmations uint64, _sender common.Address) (*big.Int, error) {
	return _SymbioticLzDVN.Contract.GetFee0(&_SymbioticLzDVN.CallOpts, _dstEid, arg1, _confirmations, _sender)
}

// GetFee1 is a free data retrieval call binding the contract method 0xfdb9b0f1.
//
// Solidity: function getFee(address _sender, bytes , bytes _cmd, bytes _options) view returns(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) GetFee1(opts *bind.CallOpts, _sender common.Address, arg1 []byte, _cmd []byte, _options []byte) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "getFee1", _sender, arg1, _cmd, _options)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee1 is a free data retrieval call binding the contract method 0xfdb9b0f1.
//
// Solidity: function getFee(address _sender, bytes , bytes _cmd, bytes _options) view returns(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNSession) GetFee1(_sender common.Address, arg1 []byte, _cmd []byte, _options []byte) (*big.Int, error) {
	return _SymbioticLzDVN.Contract.GetFee1(&_SymbioticLzDVN.CallOpts, _sender, arg1, _cmd, _options)
}

// GetFee1 is a free data retrieval call binding the contract method 0xfdb9b0f1.
//
// Solidity: function getFee(address _sender, bytes , bytes _cmd, bytes _options) view returns(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) GetFee1(_sender common.Address, arg1 []byte, _cmd []byte, _options []byte) (*big.Int, error) {
	return _SymbioticLzDVN.Contract.GetFee1(&_SymbioticLzDVN.CallOpts, _sender, arg1, _cmd, _options)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SymbioticLzDVN *SymbioticLzDVNSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SymbioticLzDVN.Contract.GetRoleAdmin(&_SymbioticLzDVN.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SymbioticLzDVN.Contract.GetRoleAdmin(&_SymbioticLzDVN.CallOpts, role)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_SymbioticLzDVN *SymbioticLzDVNCaller) GetSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "getSigners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_SymbioticLzDVN *SymbioticLzDVNSession) GetSigners() ([]common.Address, error) {
	return _SymbioticLzDVN.Contract.GetSigners(&_SymbioticLzDVN.CallOpts)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) GetSigners() ([]common.Address, error) {
	return _SymbioticLzDVN.Contract.GetSigners(&_SymbioticLzDVN.CallOpts)
}

// GetSupportedOptionTypes is a free data retrieval call binding the contract method 0x26e67a37.
//
// Solidity: function getSupportedOptionTypes(uint32 _eid) view returns(uint8[])
func (_SymbioticLzDVN *SymbioticLzDVNCaller) GetSupportedOptionTypes(opts *bind.CallOpts, _eid uint32) ([]uint8, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "getSupportedOptionTypes", _eid)

	if err != nil {
		return *new([]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint8)).(*[]uint8)

	return out0, err

}

// GetSupportedOptionTypes is a free data retrieval call binding the contract method 0x26e67a37.
//
// Solidity: function getSupportedOptionTypes(uint32 _eid) view returns(uint8[])
func (_SymbioticLzDVN *SymbioticLzDVNSession) GetSupportedOptionTypes(_eid uint32) ([]uint8, error) {
	return _SymbioticLzDVN.Contract.GetSupportedOptionTypes(&_SymbioticLzDVN.CallOpts, _eid)
}

// GetSupportedOptionTypes is a free data retrieval call binding the contract method 0x26e67a37.
//
// Solidity: function getSupportedOptionTypes(uint32 _eid) view returns(uint8[])
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) GetSupportedOptionTypes(_eid uint32) ([]uint8, error) {
	return _SymbioticLzDVN.Contract.GetSupportedOptionTypes(&_SymbioticLzDVN.CallOpts, _eid)
}

// HasAcl is a free data retrieval call binding the contract method 0x2de11376.
//
// Solidity: function hasAcl(address _sender) view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) HasAcl(opts *bind.CallOpts, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "hasAcl", _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAcl is a free data retrieval call binding the contract method 0x2de11376.
//
// Solidity: function hasAcl(address _sender) view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNSession) HasAcl(_sender common.Address) (bool, error) {
	return _SymbioticLzDVN.Contract.HasAcl(&_SymbioticLzDVN.CallOpts, _sender)
}

// HasAcl is a free data retrieval call binding the contract method 0x2de11376.
//
// Solidity: function hasAcl(address _sender) view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) HasAcl(_sender common.Address) (bool, error) {
	return _SymbioticLzDVN.Contract.HasAcl(&_SymbioticLzDVN.CallOpts, _sender)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SymbioticLzDVN.Contract.HasRole(&_SymbioticLzDVN.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SymbioticLzDVN.Contract.HasRole(&_SymbioticLzDVN.CallOpts, role, account)
}

// HashCallData is a free data retrieval call binding the contract method 0xf010cb23.
//
// Solidity: function hashCallData(uint32 _vid, address _target, bytes _callData, uint256 _expiration) pure returns(bytes32)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) HashCallData(opts *bind.CallOpts, _vid uint32, _target common.Address, _callData []byte, _expiration *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "hashCallData", _vid, _target, _callData, _expiration)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashCallData is a free data retrieval call binding the contract method 0xf010cb23.
//
// Solidity: function hashCallData(uint32 _vid, address _target, bytes _callData, uint256 _expiration) pure returns(bytes32)
func (_SymbioticLzDVN *SymbioticLzDVNSession) HashCallData(_vid uint32, _target common.Address, _callData []byte, _expiration *big.Int) ([32]byte, error) {
	return _SymbioticLzDVN.Contract.HashCallData(&_SymbioticLzDVN.CallOpts, _vid, _target, _callData, _expiration)
}

// HashCallData is a free data retrieval call binding the contract method 0xf010cb23.
//
// Solidity: function hashCallData(uint32 _vid, address _target, bytes _callData, uint256 _expiration) pure returns(bytes32)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) HashCallData(_vid uint32, _target common.Address, _callData []byte, _expiration *big.Int) ([32]byte, error) {
	return _SymbioticLzDVN.Contract.HashCallData(&_SymbioticLzDVN.CallOpts, _vid, _target, _callData, _expiration)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _signer) view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) IsSigner(opts *bind.CallOpts, _signer common.Address) (bool, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "isSigner", _signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _signer) view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNSession) IsSigner(_signer common.Address) (bool, error) {
	return _SymbioticLzDVN.Contract.IsSigner(&_SymbioticLzDVN.CallOpts, _signer)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _signer) view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) IsSigner(_signer common.Address) (bool, error) {
	return _SymbioticLzDVN.Contract.IsSigner(&_SymbioticLzDVN.CallOpts, _signer)
}

// LocalEidV2 is a free data retrieval call binding the contract method 0xe395eb5c.
//
// Solidity: function localEidV2() view returns(uint32)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) LocalEidV2(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "localEidV2")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalEidV2 is a free data retrieval call binding the contract method 0xe395eb5c.
//
// Solidity: function localEidV2() view returns(uint32)
func (_SymbioticLzDVN *SymbioticLzDVNSession) LocalEidV2() (uint32, error) {
	return _SymbioticLzDVN.Contract.LocalEidV2(&_SymbioticLzDVN.CallOpts)
}

// LocalEidV2 is a free data retrieval call binding the contract method 0xe395eb5c.
//
// Solidity: function localEidV2() view returns(uint32)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) LocalEidV2() (uint32, error) {
	return _SymbioticLzDVN.Contract.LocalEidV2(&_SymbioticLzDVN.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNSession) Paused() (bool, error) {
	return _SymbioticLzDVN.Contract.Paused(&_SymbioticLzDVN.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) Paused() (bool, error) {
	return _SymbioticLzDVN.Contract.Paused(&_SymbioticLzDVN.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNSession) PriceFeed() (common.Address, error) {
	return _SymbioticLzDVN.Contract.PriceFeed(&_SymbioticLzDVN.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) PriceFeed() (common.Address, error) {
	return _SymbioticLzDVN.Contract.PriceFeed(&_SymbioticLzDVN.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) Quorum(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "quorum")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_SymbioticLzDVN *SymbioticLzDVNSession) Quorum() (uint64, error) {
	return _SymbioticLzDVN.Contract.Quorum(&_SymbioticLzDVN.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) Quorum() (uint64, error) {
	return _SymbioticLzDVN.Contract.Quorum(&_SymbioticLzDVN.CallOpts)
}

// ReceiveUlnContract is a free data retrieval call binding the contract method 0x8aae6379.
//
// Solidity: function receiveUlnContract() view returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) ReceiveUlnContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "receiveUlnContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReceiveUlnContract is a free data retrieval call binding the contract method 0x8aae6379.
//
// Solidity: function receiveUlnContract() view returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNSession) ReceiveUlnContract() (common.Address, error) {
	return _SymbioticLzDVN.Contract.ReceiveUlnContract(&_SymbioticLzDVN.CallOpts)
}

// ReceiveUlnContract is a free data retrieval call binding the contract method 0x8aae6379.
//
// Solidity: function receiveUlnContract() view returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) ReceiveUlnContract() (common.Address, error) {
	return _SymbioticLzDVN.Contract.ReceiveUlnContract(&_SymbioticLzDVN.CallOpts)
}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_SymbioticLzDVN *SymbioticLzDVNCaller) RenounceRole(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) error {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "renounceRole", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) RenounceRole(arg0 [32]byte, arg1 common.Address) error {
	return _SymbioticLzDVN.Contract.RenounceRole(&_SymbioticLzDVN.CallOpts, arg0, arg1)
}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) RenounceRole(arg0 [32]byte, arg1 common.Address) error {
	return _SymbioticLzDVN.Contract.RenounceRole(&_SymbioticLzDVN.CallOpts, arg0, arg1)
}

// SettlementContract is a free data retrieval call binding the contract method 0xea42418b.
//
// Solidity: function settlementContract() view returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) SettlementContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "settlementContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SettlementContract is a free data retrieval call binding the contract method 0xea42418b.
//
// Solidity: function settlementContract() view returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNSession) SettlementContract() (common.Address, error) {
	return _SymbioticLzDVN.Contract.SettlementContract(&_SymbioticLzDVN.CallOpts)
}

// SettlementContract is a free data retrieval call binding the contract method 0xea42418b.
//
// Solidity: function settlementContract() view returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) SettlementContract() (common.Address, error) {
	return _SymbioticLzDVN.Contract.SettlementContract(&_SymbioticLzDVN.CallOpts)
}

// SignerSize is a free data retrieval call binding the contract method 0xfd62e750.
//
// Solidity: function signerSize() view returns(uint256)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) SignerSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "signerSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignerSize is a free data retrieval call binding the contract method 0xfd62e750.
//
// Solidity: function signerSize() view returns(uint256)
func (_SymbioticLzDVN *SymbioticLzDVNSession) SignerSize() (*big.Int, error) {
	return _SymbioticLzDVN.Contract.SignerSize(&_SymbioticLzDVN.CallOpts)
}

// SignerSize is a free data retrieval call binding the contract method 0xfd62e750.
//
// Solidity: function signerSize() view returns(uint256)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) SignerSize() (*big.Int, error) {
	return _SymbioticLzDVN.Contract.SignerSize(&_SymbioticLzDVN.CallOpts)
}

// Signers is a free data retrieval call binding the contract method 0x736c0d5b.
//
// Solidity: function signers(address _signer) view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) Signers(opts *bind.CallOpts, _signer common.Address) (bool, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "signers", _signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Signers is a free data retrieval call binding the contract method 0x736c0d5b.
//
// Solidity: function signers(address _signer) view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNSession) Signers(_signer common.Address) (bool, error) {
	return _SymbioticLzDVN.Contract.Signers(&_SymbioticLzDVN.CallOpts, _signer)
}

// Signers is a free data retrieval call binding the contract method 0x736c0d5b.
//
// Solidity: function signers(address _signer) view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) Signers(_signer common.Address) (bool, error) {
	return _SymbioticLzDVN.Contract.Signers(&_SymbioticLzDVN.CallOpts, _signer)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SymbioticLzDVN.Contract.SupportsInterface(&_SymbioticLzDVN.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SymbioticLzDVN.Contract.SupportsInterface(&_SymbioticLzDVN.CallOpts, interfaceId)
}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 executableHash) view returns(bool used)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) UsedHashes(opts *bind.CallOpts, executableHash [32]byte) (bool, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "usedHashes", executableHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 executableHash) view returns(bool used)
func (_SymbioticLzDVN *SymbioticLzDVNSession) UsedHashes(executableHash [32]byte) (bool, error) {
	return _SymbioticLzDVN.Contract.UsedHashes(&_SymbioticLzDVN.CallOpts, executableHash)
}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 executableHash) view returns(bool used)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) UsedHashes(executableHash [32]byte) (bool, error) {
	return _SymbioticLzDVN.Contract.UsedHashes(&_SymbioticLzDVN.CallOpts, executableHash)
}

// VerifySignatures is a free data retrieval call binding the contract method 0xc7a823e0.
//
// Solidity: function verifySignatures(bytes32 _hash, bytes _signatures) view returns(bool, uint8)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) VerifySignatures(opts *bind.CallOpts, _hash [32]byte, _signatures []byte) (bool, uint8, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "verifySignatures", _hash, _signatures)

	if err != nil {
		return *new(bool), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// VerifySignatures is a free data retrieval call binding the contract method 0xc7a823e0.
//
// Solidity: function verifySignatures(bytes32 _hash, bytes _signatures) view returns(bool, uint8)
func (_SymbioticLzDVN *SymbioticLzDVNSession) VerifySignatures(_hash [32]byte, _signatures []byte) (bool, uint8, error) {
	return _SymbioticLzDVN.Contract.VerifySignatures(&_SymbioticLzDVN.CallOpts, _hash, _signatures)
}

// VerifySignatures is a free data retrieval call binding the contract method 0xc7a823e0.
//
// Solidity: function verifySignatures(bytes32 _hash, bytes _signatures) view returns(bool, uint8)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) VerifySignatures(_hash [32]byte, _signatures []byte) (bool, uint8, error) {
	return _SymbioticLzDVN.Contract.VerifySignatures(&_SymbioticLzDVN.CallOpts, _hash, _signatures)
}

// Vid is a free data retrieval call binding the contract method 0xcf34c768.
//
// Solidity: function vid() view returns(uint32)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) Vid(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "vid")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Vid is a free data retrieval call binding the contract method 0xcf34c768.
//
// Solidity: function vid() view returns(uint32)
func (_SymbioticLzDVN *SymbioticLzDVNSession) Vid() (uint32, error) {
	return _SymbioticLzDVN.Contract.Vid(&_SymbioticLzDVN.CallOpts)
}

// Vid is a free data retrieval call binding the contract method 0xcf34c768.
//
// Solidity: function vid() view returns(uint32)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) Vid() (uint32, error) {
	return _SymbioticLzDVN.Contract.Vid(&_SymbioticLzDVN.CallOpts)
}

// WorkerFeeLib is a free data retrieval call binding the contract method 0xc416aa51.
//
// Solidity: function workerFeeLib() view returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) WorkerFeeLib(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "workerFeeLib")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WorkerFeeLib is a free data retrieval call binding the contract method 0xc416aa51.
//
// Solidity: function workerFeeLib() view returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNSession) WorkerFeeLib() (common.Address, error) {
	return _SymbioticLzDVN.Contract.WorkerFeeLib(&_SymbioticLzDVN.CallOpts)
}

// WorkerFeeLib is a free data retrieval call binding the contract method 0xc416aa51.
//
// Solidity: function workerFeeLib() view returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) WorkerFeeLib() (common.Address, error) {
	return _SymbioticLzDVN.Contract.WorkerFeeLib(&_SymbioticLzDVN.CallOpts)
}

// AssignJob is a paid mutator transaction binding the contract method 0x95d376d7.
//
// Solidity: function assignJob((uint32,bytes,bytes32,uint64,address) _param, bytes _options) payable returns(uint256 totalFee)
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) AssignJob(opts *bind.TransactOpts, _param ILayerZeroDVNAssignJobParam, _options []byte) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "assignJob", _param, _options)
}

// AssignJob is a paid mutator transaction binding the contract method 0x95d376d7.
//
// Solidity: function assignJob((uint32,bytes,bytes32,uint64,address) _param, bytes _options) payable returns(uint256 totalFee)
func (_SymbioticLzDVN *SymbioticLzDVNSession) AssignJob(_param ILayerZeroDVNAssignJobParam, _options []byte) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.AssignJob(&_SymbioticLzDVN.TransactOpts, _param, _options)
}

// AssignJob is a paid mutator transaction binding the contract method 0x95d376d7.
//
// Solidity: function assignJob((uint32,bytes,bytes32,uint64,address) _param, bytes _options) payable returns(uint256 totalFee)
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) AssignJob(_param ILayerZeroDVNAssignJobParam, _options []byte) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.AssignJob(&_SymbioticLzDVN.TransactOpts, _param, _options)
}

// AssignJob0 is a paid mutator transaction binding the contract method 0xc5e193cd.
//
// Solidity: function assignJob(uint16 _dstEid, uint16 , uint64 _confirmations, address _sender) returns(uint256 totalFee)
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) AssignJob0(opts *bind.TransactOpts, _dstEid uint16, arg1 uint16, _confirmations uint64, _sender common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "assignJob0", _dstEid, arg1, _confirmations, _sender)
}

// AssignJob0 is a paid mutator transaction binding the contract method 0xc5e193cd.
//
// Solidity: function assignJob(uint16 _dstEid, uint16 , uint64 _confirmations, address _sender) returns(uint256 totalFee)
func (_SymbioticLzDVN *SymbioticLzDVNSession) AssignJob0(_dstEid uint16, arg1 uint16, _confirmations uint64, _sender common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.AssignJob0(&_SymbioticLzDVN.TransactOpts, _dstEid, arg1, _confirmations, _sender)
}

// AssignJob0 is a paid mutator transaction binding the contract method 0xc5e193cd.
//
// Solidity: function assignJob(uint16 _dstEid, uint16 , uint64 _confirmations, address _sender) returns(uint256 totalFee)
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) AssignJob0(_dstEid uint16, arg1 uint16, _confirmations uint64, _sender common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.AssignJob0(&_SymbioticLzDVN.TransactOpts, _dstEid, arg1, _confirmations, _sender)
}

// AssignJob1 is a paid mutator transaction binding the contract method 0xf42ed2ed.
//
// Solidity: function assignJob(address _sender, bytes , bytes _cmd, bytes _options) payable returns(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) AssignJob1(opts *bind.TransactOpts, _sender common.Address, arg1 []byte, _cmd []byte, _options []byte) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "assignJob1", _sender, arg1, _cmd, _options)
}

// AssignJob1 is a paid mutator transaction binding the contract method 0xf42ed2ed.
//
// Solidity: function assignJob(address _sender, bytes , bytes _cmd, bytes _options) payable returns(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNSession) AssignJob1(_sender common.Address, arg1 []byte, _cmd []byte, _options []byte) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.AssignJob1(&_SymbioticLzDVN.TransactOpts, _sender, arg1, _cmd, _options)
}

// AssignJob1 is a paid mutator transaction binding the contract method 0xf42ed2ed.
//
// Solidity: function assignJob(address _sender, bytes , bytes _cmd, bytes _options) payable returns(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) AssignJob1(_sender common.Address, arg1 []byte, _cmd []byte, _options []byte) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.AssignJob1(&_SymbioticLzDVN.TransactOpts, _sender, arg1, _cmd, _options)
}

// Execute is a paid mutator transaction binding the contract method 0xb143044b.
//
// Solidity: function execute((uint32,address,bytes,uint256,bytes)[] _params) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) Execute(opts *bind.TransactOpts, _params []ExecuteParam) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "execute", _params)
}

// Execute is a paid mutator transaction binding the contract method 0xb143044b.
//
// Solidity: function execute((uint32,address,bytes,uint256,bytes)[] _params) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) Execute(_params []ExecuteParam) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.Execute(&_SymbioticLzDVN.TransactOpts, _params)
}

// Execute is a paid mutator transaction binding the contract method 0xb143044b.
//
// Solidity: function execute((uint32,address,bytes,uint256,bytes)[] _params) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) Execute(_params []ExecuteParam) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.Execute(&_SymbioticLzDVN.TransactOpts, _params)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 _role, address _account) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) GrantRole(opts *bind.TransactOpts, _role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "grantRole", _role, _account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 _role, address _account) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) GrantRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.GrantRole(&_SymbioticLzDVN.TransactOpts, _role, _account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 _role, address _account) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) GrantRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.GrantRole(&_SymbioticLzDVN.TransactOpts, _role, _account)
}

// QuorumChangeAdmin is a paid mutator transaction binding the contract method 0xf3b4ebd0.
//
// Solidity: function quorumChangeAdmin((uint32,address,bytes,uint256,bytes) _param) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) QuorumChangeAdmin(opts *bind.TransactOpts, _param ExecuteParam) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "quorumChangeAdmin", _param)
}

// QuorumChangeAdmin is a paid mutator transaction binding the contract method 0xf3b4ebd0.
//
// Solidity: function quorumChangeAdmin((uint32,address,bytes,uint256,bytes) _param) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) QuorumChangeAdmin(_param ExecuteParam) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.QuorumChangeAdmin(&_SymbioticLzDVN.TransactOpts, _param)
}

// QuorumChangeAdmin is a paid mutator transaction binding the contract method 0xf3b4ebd0.
//
// Solidity: function quorumChangeAdmin((uint32,address,bytes,uint256,bytes) _param) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) QuorumChangeAdmin(_param ExecuteParam) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.QuorumChangeAdmin(&_SymbioticLzDVN.TransactOpts, _param)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 _role, address _account) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) RevokeRole(opts *bind.TransactOpts, _role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "revokeRole", _role, _account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 _role, address _account) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) RevokeRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.RevokeRole(&_SymbioticLzDVN.TransactOpts, _role, _account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 _role, address _account) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) RevokeRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.RevokeRole(&_SymbioticLzDVN.TransactOpts, _role, _account)
}

// SetDefaultMultiplierBps is a paid mutator transaction binding the contract method 0xc358de0a.
//
// Solidity: function setDefaultMultiplierBps(uint16 _multiplierBps) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) SetDefaultMultiplierBps(opts *bind.TransactOpts, _multiplierBps uint16) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "setDefaultMultiplierBps", _multiplierBps)
}

// SetDefaultMultiplierBps is a paid mutator transaction binding the contract method 0xc358de0a.
//
// Solidity: function setDefaultMultiplierBps(uint16 _multiplierBps) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) SetDefaultMultiplierBps(_multiplierBps uint16) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetDefaultMultiplierBps(&_SymbioticLzDVN.TransactOpts, _multiplierBps)
}

// SetDefaultMultiplierBps is a paid mutator transaction binding the contract method 0xc358de0a.
//
// Solidity: function setDefaultMultiplierBps(uint16 _multiplierBps) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) SetDefaultMultiplierBps(_multiplierBps uint16) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetDefaultMultiplierBps(&_SymbioticLzDVN.TransactOpts, _multiplierBps)
}

// SetDstConfig is a paid mutator transaction binding the contract method 0x52d3b871.
//
// Solidity: function setDstConfig((uint32,uint64,uint16,uint128)[] _params) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) SetDstConfig(opts *bind.TransactOpts, _params []IDVNDstConfigParam) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "setDstConfig", _params)
}

// SetDstConfig is a paid mutator transaction binding the contract method 0x52d3b871.
//
// Solidity: function setDstConfig((uint32,uint64,uint16,uint128)[] _params) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) SetDstConfig(_params []IDVNDstConfigParam) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetDstConfig(&_SymbioticLzDVN.TransactOpts, _params)
}

// SetDstConfig is a paid mutator transaction binding the contract method 0x52d3b871.
//
// Solidity: function setDstConfig((uint32,uint64,uint16,uint128)[] _params) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) SetDstConfig(_params []IDVNDstConfigParam) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetDstConfig(&_SymbioticLzDVN.TransactOpts, _params)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) SetPaused(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "setPaused", _paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) SetPaused(_paused bool) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetPaused(&_SymbioticLzDVN.TransactOpts, _paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) SetPaused(_paused bool) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetPaused(&_SymbioticLzDVN.TransactOpts, _paused)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) SetPriceFeed(opts *bind.TransactOpts, _priceFeed common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "setPriceFeed", _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetPriceFeed(&_SymbioticLzDVN.TransactOpts, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetPriceFeed(&_SymbioticLzDVN.TransactOpts, _priceFeed)
}

// SetQuorum is a paid mutator transaction binding the contract method 0x8585c945.
//
// Solidity: function setQuorum(uint64 _quorum) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) SetQuorum(opts *bind.TransactOpts, _quorum uint64) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "setQuorum", _quorum)
}

// SetQuorum is a paid mutator transaction binding the contract method 0x8585c945.
//
// Solidity: function setQuorum(uint64 _quorum) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) SetQuorum(_quorum uint64) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetQuorum(&_SymbioticLzDVN.TransactOpts, _quorum)
}

// SetQuorum is a paid mutator transaction binding the contract method 0x8585c945.
//
// Solidity: function setQuorum(uint64 _quorum) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) SetQuorum(_quorum uint64) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetQuorum(&_SymbioticLzDVN.TransactOpts, _quorum)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address _signer, bool _active) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) SetSigner(opts *bind.TransactOpts, _signer common.Address, _active bool) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "setSigner", _signer, _active)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address _signer, bool _active) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) SetSigner(_signer common.Address, _active bool) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetSigner(&_SymbioticLzDVN.TransactOpts, _signer, _active)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address _signer, bool _active) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) SetSigner(_signer common.Address, _active bool) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetSigner(&_SymbioticLzDVN.TransactOpts, _signer, _active)
}

// SetSupportedOptionTypes is a paid mutator transaction binding the contract method 0xcd88b903.
//
// Solidity: function setSupportedOptionTypes(uint32 _eid, uint8[] _optionTypes) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) SetSupportedOptionTypes(opts *bind.TransactOpts, _eid uint32, _optionTypes []uint8) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "setSupportedOptionTypes", _eid, _optionTypes)
}

// SetSupportedOptionTypes is a paid mutator transaction binding the contract method 0xcd88b903.
//
// Solidity: function setSupportedOptionTypes(uint32 _eid, uint8[] _optionTypes) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) SetSupportedOptionTypes(_eid uint32, _optionTypes []uint8) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetSupportedOptionTypes(&_SymbioticLzDVN.TransactOpts, _eid, _optionTypes)
}

// SetSupportedOptionTypes is a paid mutator transaction binding the contract method 0xcd88b903.
//
// Solidity: function setSupportedOptionTypes(uint32 _eid, uint8[] _optionTypes) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) SetSupportedOptionTypes(_eid uint32, _optionTypes []uint8) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetSupportedOptionTypes(&_SymbioticLzDVN.TransactOpts, _eid, _optionTypes)
}

// SetWorkerFeeLib is a paid mutator transaction binding the contract method 0xc7b2370b.
//
// Solidity: function setWorkerFeeLib(address _workerFeeLib) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) SetWorkerFeeLib(opts *bind.TransactOpts, _workerFeeLib common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "setWorkerFeeLib", _workerFeeLib)
}

// SetWorkerFeeLib is a paid mutator transaction binding the contract method 0xc7b2370b.
//
// Solidity: function setWorkerFeeLib(address _workerFeeLib) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) SetWorkerFeeLib(_workerFeeLib common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetWorkerFeeLib(&_SymbioticLzDVN.TransactOpts, _workerFeeLib)
}

// SetWorkerFeeLib is a paid mutator transaction binding the contract method 0xc7b2370b.
//
// Solidity: function setWorkerFeeLib(address _workerFeeLib) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) SetWorkerFeeLib(_workerFeeLib common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetWorkerFeeLib(&_SymbioticLzDVN.TransactOpts, _workerFeeLib)
}

// VerifyWithSymbiotic is a paid mutator transaction binding the contract method 0x0b3e7f0c.
//
// Solidity: function verifyWithSymbiotic(bytes _packetHeader, bytes32 _payloadHash, uint64 _confirmations, bytes _symbioticProof) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) VerifyWithSymbiotic(opts *bind.TransactOpts, _packetHeader []byte, _payloadHash [32]byte, _confirmations uint64, _symbioticProof []byte) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "verifyWithSymbiotic", _packetHeader, _payloadHash, _confirmations, _symbioticProof)
}

// VerifyWithSymbiotic is a paid mutator transaction binding the contract method 0x0b3e7f0c.
//
// Solidity: function verifyWithSymbiotic(bytes _packetHeader, bytes32 _payloadHash, uint64 _confirmations, bytes _symbioticProof) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) VerifyWithSymbiotic(_packetHeader []byte, _payloadHash [32]byte, _confirmations uint64, _symbioticProof []byte) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.VerifyWithSymbiotic(&_SymbioticLzDVN.TransactOpts, _packetHeader, _payloadHash, _confirmations, _symbioticProof)
}

// VerifyWithSymbiotic is a paid mutator transaction binding the contract method 0x0b3e7f0c.
//
// Solidity: function verifyWithSymbiotic(bytes _packetHeader, bytes32 _payloadHash, uint64 _confirmations, bytes _symbioticProof) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) VerifyWithSymbiotic(_packetHeader []byte, _payloadHash [32]byte, _confirmations uint64, _symbioticProof []byte) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.VerifyWithSymbiotic(&_SymbioticLzDVN.TransactOpts, _packetHeader, _payloadHash, _confirmations, _symbioticProof)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1095b6d7.
//
// Solidity: function withdrawFee(address _lib, address _to, uint256 _amount) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) WithdrawFee(opts *bind.TransactOpts, _lib common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "withdrawFee", _lib, _to, _amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1095b6d7.
//
// Solidity: function withdrawFee(address _lib, address _to, uint256 _amount) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) WithdrawFee(_lib common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.WithdrawFee(&_SymbioticLzDVN.TransactOpts, _lib, _to, _amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1095b6d7.
//
// Solidity: function withdrawFee(address _lib, address _to, uint256 _amount) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) WithdrawFee(_lib common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.WithdrawFee(&_SymbioticLzDVN.TransactOpts, _lib, _to, _amount)
}

// WithdrawFeeFromUlnV2 is a paid mutator transaction binding the contract method 0xdafe0ccc.
//
// Solidity: function withdrawFeeFromUlnV2(address _lib, address _to, uint256 _amount) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) WithdrawFeeFromUlnV2(opts *bind.TransactOpts, _lib common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "withdrawFeeFromUlnV2", _lib, _to, _amount)
}

// WithdrawFeeFromUlnV2 is a paid mutator transaction binding the contract method 0xdafe0ccc.
//
// Solidity: function withdrawFeeFromUlnV2(address _lib, address _to, uint256 _amount) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) WithdrawFeeFromUlnV2(_lib common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.WithdrawFeeFromUlnV2(&_SymbioticLzDVN.TransactOpts, _lib, _to, _amount)
}

// WithdrawFeeFromUlnV2 is a paid mutator transaction binding the contract method 0xdafe0ccc.
//
// Solidity: function withdrawFeeFromUlnV2(address _lib, address _to, uint256 _amount) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) WithdrawFeeFromUlnV2(_lib common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.WithdrawFeeFromUlnV2(&_SymbioticLzDVN.TransactOpts, _lib, _to, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _to, uint256 _amount) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "withdrawToken", _token, _to, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _to, uint256 _amount) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) WithdrawToken(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.WithdrawToken(&_SymbioticLzDVN.TransactOpts, _token, _to, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _to, uint256 _amount) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) WithdrawToken(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.WithdrawToken(&_SymbioticLzDVN.TransactOpts, _token, _to, _amount)
}

// SymbioticLzDVNExecuteFailedIterator is returned from FilterExecuteFailed and is used to iterate over the raw logs and unpacked data for ExecuteFailed events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNExecuteFailedIterator struct {
	Event *SymbioticLzDVNExecuteFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNExecuteFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNExecuteFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNExecuteFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNExecuteFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNExecuteFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNExecuteFailed represents a ExecuteFailed event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNExecuteFailed struct {
	Index *big.Int
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExecuteFailed is a free log retrieval operation binding the contract event 0xdc8cdd96296241bbefda4a8e18ad2e0985a8da6495b34d409cfc4c886ee3b0cf.
//
// Solidity: event ExecuteFailed(uint256 _index, bytes _data)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterExecuteFailed(opts *bind.FilterOpts) (*SymbioticLzDVNExecuteFailedIterator, error) {

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "ExecuteFailed")
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNExecuteFailedIterator{contract: _SymbioticLzDVN.contract, event: "ExecuteFailed", logs: logs, sub: sub}, nil
}

// WatchExecuteFailed is a free log subscription operation binding the contract event 0xdc8cdd96296241bbefda4a8e18ad2e0985a8da6495b34d409cfc4c886ee3b0cf.
//
// Solidity: event ExecuteFailed(uint256 _index, bytes _data)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchExecuteFailed(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNExecuteFailed) (event.Subscription, error) {

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "ExecuteFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNExecuteFailed)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "ExecuteFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecuteFailed is a log parse operation binding the contract event 0xdc8cdd96296241bbefda4a8e18ad2e0985a8da6495b34d409cfc4c886ee3b0cf.
//
// Solidity: event ExecuteFailed(uint256 _index, bytes _data)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParseExecuteFailed(log types.Log) (*SymbioticLzDVNExecuteFailed, error) {
	event := new(SymbioticLzDVNExecuteFailed)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "ExecuteFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticLzDVNHashAlreadyUsedIterator is returned from FilterHashAlreadyUsed and is used to iterate over the raw logs and unpacked data for HashAlreadyUsed events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNHashAlreadyUsedIterator struct {
	Event *SymbioticLzDVNHashAlreadyUsed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNHashAlreadyUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNHashAlreadyUsed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNHashAlreadyUsed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNHashAlreadyUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNHashAlreadyUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNHashAlreadyUsed represents a HashAlreadyUsed event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNHashAlreadyUsed struct {
	Param ExecuteParam
	Hash  [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterHashAlreadyUsed is a free log retrieval operation binding the contract event 0x9bb9bddbdf537a2104255307230b323d7982f4512ee8e5bd15df62ddca50ab97.
//
// Solidity: event HashAlreadyUsed((uint32,address,bytes,uint256,bytes) param, bytes32 _hash)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterHashAlreadyUsed(opts *bind.FilterOpts) (*SymbioticLzDVNHashAlreadyUsedIterator, error) {

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "HashAlreadyUsed")
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNHashAlreadyUsedIterator{contract: _SymbioticLzDVN.contract, event: "HashAlreadyUsed", logs: logs, sub: sub}, nil
}

// WatchHashAlreadyUsed is a free log subscription operation binding the contract event 0x9bb9bddbdf537a2104255307230b323d7982f4512ee8e5bd15df62ddca50ab97.
//
// Solidity: event HashAlreadyUsed((uint32,address,bytes,uint256,bytes) param, bytes32 _hash)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchHashAlreadyUsed(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNHashAlreadyUsed) (event.Subscription, error) {

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "HashAlreadyUsed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNHashAlreadyUsed)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "HashAlreadyUsed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHashAlreadyUsed is a log parse operation binding the contract event 0x9bb9bddbdf537a2104255307230b323d7982f4512ee8e5bd15df62ddca50ab97.
//
// Solidity: event HashAlreadyUsed((uint32,address,bytes,uint256,bytes) param, bytes32 _hash)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParseHashAlreadyUsed(log types.Log) (*SymbioticLzDVNHashAlreadyUsed, error) {
	event := new(SymbioticLzDVNHashAlreadyUsed)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "HashAlreadyUsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticLzDVNPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNPausedIterator struct {
	Event *SymbioticLzDVNPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNPaused represents a Paused event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterPaused(opts *bind.FilterOpts) (*SymbioticLzDVNPausedIterator, error) {

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNPausedIterator{contract: _SymbioticLzDVN.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNPaused) (event.Subscription, error) {

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNPaused)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParsePaused(log types.Log) (*SymbioticLzDVNPaused, error) {
	event := new(SymbioticLzDVNPaused)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticLzDVNRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNRoleAdminChangedIterator struct {
	Event *SymbioticLzDVNRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNRoleAdminChanged represents a RoleAdminChanged event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SymbioticLzDVNRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNRoleAdminChangedIterator{contract: _SymbioticLzDVN.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNRoleAdminChanged)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParseRoleAdminChanged(log types.Log) (*SymbioticLzDVNRoleAdminChanged, error) {
	event := new(SymbioticLzDVNRoleAdminChanged)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticLzDVNRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNRoleGrantedIterator struct {
	Event *SymbioticLzDVNRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNRoleGranted represents a RoleGranted event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SymbioticLzDVNRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNRoleGrantedIterator{contract: _SymbioticLzDVN.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNRoleGranted)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParseRoleGranted(log types.Log) (*SymbioticLzDVNRoleGranted, error) {
	event := new(SymbioticLzDVNRoleGranted)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticLzDVNRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNRoleRevokedIterator struct {
	Event *SymbioticLzDVNRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNRoleRevoked represents a RoleRevoked event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SymbioticLzDVNRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNRoleRevokedIterator{contract: _SymbioticLzDVN.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNRoleRevoked)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParseRoleRevoked(log types.Log) (*SymbioticLzDVNRoleRevoked, error) {
	event := new(SymbioticLzDVNRoleRevoked)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticLzDVNSetDefaultMultiplierBpsIterator is returned from FilterSetDefaultMultiplierBps and is used to iterate over the raw logs and unpacked data for SetDefaultMultiplierBps events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNSetDefaultMultiplierBpsIterator struct {
	Event *SymbioticLzDVNSetDefaultMultiplierBps // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNSetDefaultMultiplierBpsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNSetDefaultMultiplierBps)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNSetDefaultMultiplierBps)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNSetDefaultMultiplierBpsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNSetDefaultMultiplierBpsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNSetDefaultMultiplierBps represents a SetDefaultMultiplierBps event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNSetDefaultMultiplierBps struct {
	MultiplierBps uint16
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetDefaultMultiplierBps is a free log retrieval operation binding the contract event 0x7af0ac740036ffb1c97b03697859d729e80a44ae5030543d64971c313565ab4d.
//
// Solidity: event SetDefaultMultiplierBps(uint16 multiplierBps)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterSetDefaultMultiplierBps(opts *bind.FilterOpts) (*SymbioticLzDVNSetDefaultMultiplierBpsIterator, error) {

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "SetDefaultMultiplierBps")
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNSetDefaultMultiplierBpsIterator{contract: _SymbioticLzDVN.contract, event: "SetDefaultMultiplierBps", logs: logs, sub: sub}, nil
}

// WatchSetDefaultMultiplierBps is a free log subscription operation binding the contract event 0x7af0ac740036ffb1c97b03697859d729e80a44ae5030543d64971c313565ab4d.
//
// Solidity: event SetDefaultMultiplierBps(uint16 multiplierBps)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchSetDefaultMultiplierBps(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNSetDefaultMultiplierBps) (event.Subscription, error) {

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "SetDefaultMultiplierBps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNSetDefaultMultiplierBps)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "SetDefaultMultiplierBps", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetDefaultMultiplierBps is a log parse operation binding the contract event 0x7af0ac740036ffb1c97b03697859d729e80a44ae5030543d64971c313565ab4d.
//
// Solidity: event SetDefaultMultiplierBps(uint16 multiplierBps)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParseSetDefaultMultiplierBps(log types.Log) (*SymbioticLzDVNSetDefaultMultiplierBps, error) {
	event := new(SymbioticLzDVNSetDefaultMultiplierBps)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "SetDefaultMultiplierBps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticLzDVNSetDstConfigIterator is returned from FilterSetDstConfig and is used to iterate over the raw logs and unpacked data for SetDstConfig events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNSetDstConfigIterator struct {
	Event *SymbioticLzDVNSetDstConfig // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNSetDstConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNSetDstConfig)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNSetDstConfig)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNSetDstConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNSetDstConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNSetDstConfig represents a SetDstConfig event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNSetDstConfig struct {
	Params []IDVNDstConfigParam
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetDstConfig is a free log retrieval operation binding the contract event 0x7dd21e42791b013d1929e86f0c59085e4fca24251f0f1aa81917b3b1611766e0.
//
// Solidity: event SetDstConfig((uint32,uint64,uint16,uint128)[] params)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterSetDstConfig(opts *bind.FilterOpts) (*SymbioticLzDVNSetDstConfigIterator, error) {

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "SetDstConfig")
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNSetDstConfigIterator{contract: _SymbioticLzDVN.contract, event: "SetDstConfig", logs: logs, sub: sub}, nil
}

// WatchSetDstConfig is a free log subscription operation binding the contract event 0x7dd21e42791b013d1929e86f0c59085e4fca24251f0f1aa81917b3b1611766e0.
//
// Solidity: event SetDstConfig((uint32,uint64,uint16,uint128)[] params)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchSetDstConfig(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNSetDstConfig) (event.Subscription, error) {

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "SetDstConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNSetDstConfig)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "SetDstConfig", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetDstConfig is a log parse operation binding the contract event 0x7dd21e42791b013d1929e86f0c59085e4fca24251f0f1aa81917b3b1611766e0.
//
// Solidity: event SetDstConfig((uint32,uint64,uint16,uint128)[] params)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParseSetDstConfig(log types.Log) (*SymbioticLzDVNSetDstConfig, error) {
	event := new(SymbioticLzDVNSetDstConfig)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "SetDstConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticLzDVNSetPriceFeedIterator is returned from FilterSetPriceFeed and is used to iterate over the raw logs and unpacked data for SetPriceFeed events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNSetPriceFeedIterator struct {
	Event *SymbioticLzDVNSetPriceFeed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNSetPriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNSetPriceFeed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNSetPriceFeed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNSetPriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNSetPriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNSetPriceFeed represents a SetPriceFeed event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNSetPriceFeed struct {
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetPriceFeed is a free log retrieval operation binding the contract event 0xf724a45d041687842411f2b977ef22ab8f43c8f1104f4592b42a00f9b34a643d.
//
// Solidity: event SetPriceFeed(address priceFeed)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterSetPriceFeed(opts *bind.FilterOpts) (*SymbioticLzDVNSetPriceFeedIterator, error) {

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "SetPriceFeed")
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNSetPriceFeedIterator{contract: _SymbioticLzDVN.contract, event: "SetPriceFeed", logs: logs, sub: sub}, nil
}

// WatchSetPriceFeed is a free log subscription operation binding the contract event 0xf724a45d041687842411f2b977ef22ab8f43c8f1104f4592b42a00f9b34a643d.
//
// Solidity: event SetPriceFeed(address priceFeed)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchSetPriceFeed(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNSetPriceFeed) (event.Subscription, error) {

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "SetPriceFeed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNSetPriceFeed)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "SetPriceFeed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPriceFeed is a log parse operation binding the contract event 0xf724a45d041687842411f2b977ef22ab8f43c8f1104f4592b42a00f9b34a643d.
//
// Solidity: event SetPriceFeed(address priceFeed)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParseSetPriceFeed(log types.Log) (*SymbioticLzDVNSetPriceFeed, error) {
	event := new(SymbioticLzDVNSetPriceFeed)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "SetPriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticLzDVNSetSupportedOptionTypesIterator is returned from FilterSetSupportedOptionTypes and is used to iterate over the raw logs and unpacked data for SetSupportedOptionTypes events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNSetSupportedOptionTypesIterator struct {
	Event *SymbioticLzDVNSetSupportedOptionTypes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNSetSupportedOptionTypesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNSetSupportedOptionTypes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNSetSupportedOptionTypes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNSetSupportedOptionTypesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNSetSupportedOptionTypesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNSetSupportedOptionTypes represents a SetSupportedOptionTypes event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNSetSupportedOptionTypes struct {
	DstEid      uint32
	OptionTypes []uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetSupportedOptionTypes is a free log retrieval operation binding the contract event 0x7a68270ca5c336167a62e36d8b7a73fecfabd6ce0ddc70be859bf4e2e3fb02d4.
//
// Solidity: event SetSupportedOptionTypes(uint32 dstEid, uint8[] optionTypes)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterSetSupportedOptionTypes(opts *bind.FilterOpts) (*SymbioticLzDVNSetSupportedOptionTypesIterator, error) {

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "SetSupportedOptionTypes")
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNSetSupportedOptionTypesIterator{contract: _SymbioticLzDVN.contract, event: "SetSupportedOptionTypes", logs: logs, sub: sub}, nil
}

// WatchSetSupportedOptionTypes is a free log subscription operation binding the contract event 0x7a68270ca5c336167a62e36d8b7a73fecfabd6ce0ddc70be859bf4e2e3fb02d4.
//
// Solidity: event SetSupportedOptionTypes(uint32 dstEid, uint8[] optionTypes)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchSetSupportedOptionTypes(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNSetSupportedOptionTypes) (event.Subscription, error) {

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "SetSupportedOptionTypes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNSetSupportedOptionTypes)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "SetSupportedOptionTypes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetSupportedOptionTypes is a log parse operation binding the contract event 0x7a68270ca5c336167a62e36d8b7a73fecfabd6ce0ddc70be859bf4e2e3fb02d4.
//
// Solidity: event SetSupportedOptionTypes(uint32 dstEid, uint8[] optionTypes)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParseSetSupportedOptionTypes(log types.Log) (*SymbioticLzDVNSetSupportedOptionTypes, error) {
	event := new(SymbioticLzDVNSetSupportedOptionTypes)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "SetSupportedOptionTypes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticLzDVNSetWorkerLibIterator is returned from FilterSetWorkerLib and is used to iterate over the raw logs and unpacked data for SetWorkerLib events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNSetWorkerLibIterator struct {
	Event *SymbioticLzDVNSetWorkerLib // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNSetWorkerLibIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNSetWorkerLib)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNSetWorkerLib)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNSetWorkerLibIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNSetWorkerLibIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNSetWorkerLib represents a SetWorkerLib event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNSetWorkerLib struct {
	WorkerLib common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetWorkerLib is a free log retrieval operation binding the contract event 0x1399be28223800f8669b3ba5f8721d9fc16fc4e8d0bbf98378791c8c5a3015e0.
//
// Solidity: event SetWorkerLib(address workerLib)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterSetWorkerLib(opts *bind.FilterOpts) (*SymbioticLzDVNSetWorkerLibIterator, error) {

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "SetWorkerLib")
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNSetWorkerLibIterator{contract: _SymbioticLzDVN.contract, event: "SetWorkerLib", logs: logs, sub: sub}, nil
}

// WatchSetWorkerLib is a free log subscription operation binding the contract event 0x1399be28223800f8669b3ba5f8721d9fc16fc4e8d0bbf98378791c8c5a3015e0.
//
// Solidity: event SetWorkerLib(address workerLib)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchSetWorkerLib(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNSetWorkerLib) (event.Subscription, error) {

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "SetWorkerLib")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNSetWorkerLib)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "SetWorkerLib", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetWorkerLib is a log parse operation binding the contract event 0x1399be28223800f8669b3ba5f8721d9fc16fc4e8d0bbf98378791c8c5a3015e0.
//
// Solidity: event SetWorkerLib(address workerLib)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParseSetWorkerLib(log types.Log) (*SymbioticLzDVNSetWorkerLib, error) {
	event := new(SymbioticLzDVNSetWorkerLib)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "SetWorkerLib", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticLzDVNUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNUnpausedIterator struct {
	Event *SymbioticLzDVNUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNUnpaused represents a Unpaused event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SymbioticLzDVNUnpausedIterator, error) {

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNUnpausedIterator{contract: _SymbioticLzDVN.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNUnpaused) (event.Subscription, error) {

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNUnpaused)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParseUnpaused(log types.Log) (*SymbioticLzDVNUnpaused, error) {
	event := new(SymbioticLzDVNUnpaused)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticLzDVNUpdateQuorumIterator is returned from FilterUpdateQuorum and is used to iterate over the raw logs and unpacked data for UpdateQuorum events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNUpdateQuorumIterator struct {
	Event *SymbioticLzDVNUpdateQuorum // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNUpdateQuorumIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNUpdateQuorum)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNUpdateQuorum)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNUpdateQuorumIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNUpdateQuorumIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNUpdateQuorum represents a UpdateQuorum event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNUpdateQuorum struct {
	Quorum uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateQuorum is a free log retrieval operation binding the contract event 0xb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a23249.
//
// Solidity: event UpdateQuorum(uint64 _quorum)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterUpdateQuorum(opts *bind.FilterOpts) (*SymbioticLzDVNUpdateQuorumIterator, error) {

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "UpdateQuorum")
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNUpdateQuorumIterator{contract: _SymbioticLzDVN.contract, event: "UpdateQuorum", logs: logs, sub: sub}, nil
}

// WatchUpdateQuorum is a free log subscription operation binding the contract event 0xb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a23249.
//
// Solidity: event UpdateQuorum(uint64 _quorum)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchUpdateQuorum(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNUpdateQuorum) (event.Subscription, error) {

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "UpdateQuorum")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNUpdateQuorum)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "UpdateQuorum", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateQuorum is a log parse operation binding the contract event 0xb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a23249.
//
// Solidity: event UpdateQuorum(uint64 _quorum)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParseUpdateQuorum(log types.Log) (*SymbioticLzDVNUpdateQuorum, error) {
	event := new(SymbioticLzDVNUpdateQuorum)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "UpdateQuorum", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticLzDVNUpdateSignerIterator is returned from FilterUpdateSigner and is used to iterate over the raw logs and unpacked data for UpdateSigner events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNUpdateSignerIterator struct {
	Event *SymbioticLzDVNUpdateSigner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNUpdateSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNUpdateSigner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNUpdateSigner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNUpdateSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNUpdateSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNUpdateSigner represents a UpdateSigner event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNUpdateSigner struct {
	Signer common.Address
	Active bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateSigner is a free log retrieval operation binding the contract event 0x863d338cad74814b108a06288ad5e0e80d56495e0332238b1d2cdcfa0ca8e5ce.
//
// Solidity: event UpdateSigner(address _signer, bool _active)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterUpdateSigner(opts *bind.FilterOpts) (*SymbioticLzDVNUpdateSignerIterator, error) {

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "UpdateSigner")
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNUpdateSignerIterator{contract: _SymbioticLzDVN.contract, event: "UpdateSigner", logs: logs, sub: sub}, nil
}

// WatchUpdateSigner is a free log subscription operation binding the contract event 0x863d338cad74814b108a06288ad5e0e80d56495e0332238b1d2cdcfa0ca8e5ce.
//
// Solidity: event UpdateSigner(address _signer, bool _active)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchUpdateSigner(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNUpdateSigner) (event.Subscription, error) {

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "UpdateSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNUpdateSigner)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "UpdateSigner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateSigner is a log parse operation binding the contract event 0x863d338cad74814b108a06288ad5e0e80d56495e0332238b1d2cdcfa0ca8e5ce.
//
// Solidity: event UpdateSigner(address _signer, bool _active)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParseUpdateSigner(log types.Log) (*SymbioticLzDVNUpdateSigner, error) {
	event := new(SymbioticLzDVNUpdateSigner)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "UpdateSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticLzDVNVerifierFeePaidIterator is returned from FilterVerifierFeePaid and is used to iterate over the raw logs and unpacked data for VerifierFeePaid events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNVerifierFeePaidIterator struct {
	Event *SymbioticLzDVNVerifierFeePaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNVerifierFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNVerifierFeePaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNVerifierFeePaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNVerifierFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNVerifierFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNVerifierFeePaid represents a VerifierFeePaid event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNVerifierFeePaid struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVerifierFeePaid is a free log retrieval operation binding the contract event 0x87e46b0a6199bc734632187269a103c05714ee0adae5b28f30723955724f37ef.
//
// Solidity: event VerifierFeePaid(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterVerifierFeePaid(opts *bind.FilterOpts) (*SymbioticLzDVNVerifierFeePaidIterator, error) {

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "VerifierFeePaid")
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNVerifierFeePaidIterator{contract: _SymbioticLzDVN.contract, event: "VerifierFeePaid", logs: logs, sub: sub}, nil
}

// WatchVerifierFeePaid is a free log subscription operation binding the contract event 0x87e46b0a6199bc734632187269a103c05714ee0adae5b28f30723955724f37ef.
//
// Solidity: event VerifierFeePaid(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchVerifierFeePaid(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNVerifierFeePaid) (event.Subscription, error) {

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "VerifierFeePaid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNVerifierFeePaid)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "VerifierFeePaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerifierFeePaid is a log parse operation binding the contract event 0x87e46b0a6199bc734632187269a103c05714ee0adae5b28f30723955724f37ef.
//
// Solidity: event VerifierFeePaid(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParseVerifierFeePaid(log types.Log) (*SymbioticLzDVNVerifierFeePaid, error) {
	event := new(SymbioticLzDVNVerifierFeePaid)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "VerifierFeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticLzDVNVerifySignaturesFailedIterator is returned from FilterVerifySignaturesFailed and is used to iterate over the raw logs and unpacked data for VerifySignaturesFailed events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNVerifySignaturesFailedIterator struct {
	Event *SymbioticLzDVNVerifySignaturesFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNVerifySignaturesFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNVerifySignaturesFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNVerifySignaturesFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNVerifySignaturesFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNVerifySignaturesFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNVerifySignaturesFailed represents a VerifySignaturesFailed event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNVerifySignaturesFailed struct {
	Idx *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVerifySignaturesFailed is a free log retrieval operation binding the contract event 0xd6d90193101048cc1b6edcdc2348f5acf7a4a4a97d3e7b668b74cb7602ab3ebc.
//
// Solidity: event VerifySignaturesFailed(uint256 idx)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterVerifySignaturesFailed(opts *bind.FilterOpts) (*SymbioticLzDVNVerifySignaturesFailedIterator, error) {

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "VerifySignaturesFailed")
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNVerifySignaturesFailedIterator{contract: _SymbioticLzDVN.contract, event: "VerifySignaturesFailed", logs: logs, sub: sub}, nil
}

// WatchVerifySignaturesFailed is a free log subscription operation binding the contract event 0xd6d90193101048cc1b6edcdc2348f5acf7a4a4a97d3e7b668b74cb7602ab3ebc.
//
// Solidity: event VerifySignaturesFailed(uint256 idx)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchVerifySignaturesFailed(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNVerifySignaturesFailed) (event.Subscription, error) {

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "VerifySignaturesFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNVerifySignaturesFailed)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "VerifySignaturesFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerifySignaturesFailed is a log parse operation binding the contract event 0xd6d90193101048cc1b6edcdc2348f5acf7a4a4a97d3e7b668b74cb7602ab3ebc.
//
// Solidity: event VerifySignaturesFailed(uint256 idx)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParseVerifySignaturesFailed(log types.Log) (*SymbioticLzDVNVerifySignaturesFailed, error) {
	event := new(SymbioticLzDVNVerifySignaturesFailed)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "VerifySignaturesFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticLzDVNWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SymbioticLzDVN contract.
type SymbioticLzDVNWithdrawIterator struct {
	Event *SymbioticLzDVNWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SymbioticLzDVNWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticLzDVNWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SymbioticLzDVNWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SymbioticLzDVNWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticLzDVNWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticLzDVNWithdraw represents a Withdraw event raised by the SymbioticLzDVN contract.
type SymbioticLzDVNWithdraw struct {
	Lib    common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address lib, address to, uint256 amount)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) FilterWithdraw(opts *bind.FilterOpts) (*SymbioticLzDVNWithdrawIterator, error) {

	logs, sub, err := _SymbioticLzDVN.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &SymbioticLzDVNWithdrawIterator{contract: _SymbioticLzDVN.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address lib, address to, uint256 amount)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SymbioticLzDVNWithdraw) (event.Subscription, error) {

	logs, sub, err := _SymbioticLzDVN.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticLzDVNWithdraw)
				if err := _SymbioticLzDVN.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address lib, address to, uint256 amount)
func (_SymbioticLzDVN *SymbioticLzDVNFilterer) ParseWithdraw(log types.Log) (*SymbioticLzDVNWithdraw, error) {
	event := new(SymbioticLzDVNWithdraw)
	if err := _SymbioticLzDVN.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
