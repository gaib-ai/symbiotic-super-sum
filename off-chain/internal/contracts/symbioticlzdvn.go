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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_localEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_priceFeed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_settlementContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiveUlnContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"assignJob\",\"inputs\":[{\"name\":\"_param\",\"type\":\"tuple\",\"internalType\":\"structILayerZeroDVN.AssignJobParam\",\"components\":[{\"name\":\"dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"packetHeader\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"payloadHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"_options\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"assignJob\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"defaultMultiplierBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dstConfig\",\"inputs\":[{\"name\":\"_dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"gas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"multiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"floorMarginUSD\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dstConfigLookup\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"gas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"multiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"floorMarginUSD\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeLib\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"_dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_options\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"nativeFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedOptionTypes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"localEid\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"priceFeed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveUlnContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDefaultMultiplierBps\",\"inputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDstConfig\",\"inputs\":[{\"name\":\"_params\",\"type\":\"tuple[]\",\"internalType\":\"structIDVN.DstConfigParam[]\",\"components\":[{\"name\":\"dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"multiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"floorMarginUSD\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceFeed\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSupportedOptionTypes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWorkerFeeLib\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"settlementContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyWithSymbiotic\",\"inputs\":[{\"name\":\"_packetHeader\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_payloadHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_symbioticProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"worker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"SetDefaultMultiplierBps\",\"inputs\":[{\"name\":\"multiplierBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDstConfig\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIDVN.DstConfigParam[]\",\"components\":[{\"name\":\"dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"multiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"floorMarginUSD\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetPriceFeed\",\"inputs\":[{\"name\":\"priceFeed\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetSupportedOptionTypes\",\"inputs\":[{\"name\":\"dstEid\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"optionTypes\",\"type\":\"uint8[]\",\"indexed\":false,\"internalType\":\"uint8[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetWorkerLib\",\"inputs\":[{\"name\":\"workerLib\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"lib\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"Worker_NotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Worker_OnlyMessageLib\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Worker_RoleRenouncingDisabled\",\"inputs\":[]}]",
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
// Solidity: function dstConfig(uint32 _dstEid) view returns(uint64 gas, uint16 multiplierBps, uint128 floorMarginUSD)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) DstConfig(opts *bind.CallOpts, _dstEid uint32) (struct {
	Gas            uint64
	MultiplierBps  uint16
	FloorMarginUSD *big.Int
}, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "dstConfig", _dstEid)

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
// Solidity: function dstConfig(uint32 _dstEid) view returns(uint64 gas, uint16 multiplierBps, uint128 floorMarginUSD)
func (_SymbioticLzDVN *SymbioticLzDVNSession) DstConfig(_dstEid uint32) (struct {
	Gas            uint64
	MultiplierBps  uint16
	FloorMarginUSD *big.Int
}, error) {
	return _SymbioticLzDVN.Contract.DstConfig(&_SymbioticLzDVN.CallOpts, _dstEid)
}

// DstConfig is a free data retrieval call binding the contract method 0x9e944965.
//
// Solidity: function dstConfig(uint32 _dstEid) view returns(uint64 gas, uint16 multiplierBps, uint128 floorMarginUSD)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) DstConfig(_dstEid uint32) (struct {
	Gas            uint64
	MultiplierBps  uint16
	FloorMarginUSD *big.Int
}, error) {
	return _SymbioticLzDVN.Contract.DstConfig(&_SymbioticLzDVN.CallOpts, _dstEid)
}

// DstConfigLookup is a free data retrieval call binding the contract method 0x10187719.
//
// Solidity: function dstConfigLookup(uint32 ) view returns(uint64 gas, uint16 multiplierBps, uint128 floorMarginUSD)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) DstConfigLookup(opts *bind.CallOpts, arg0 uint32) (struct {
	Gas            uint64
	MultiplierBps  uint16
	FloorMarginUSD *big.Int
}, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "dstConfigLookup", arg0)

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

// DstConfigLookup is a free data retrieval call binding the contract method 0x10187719.
//
// Solidity: function dstConfigLookup(uint32 ) view returns(uint64 gas, uint16 multiplierBps, uint128 floorMarginUSD)
func (_SymbioticLzDVN *SymbioticLzDVNSession) DstConfigLookup(arg0 uint32) (struct {
	Gas            uint64
	MultiplierBps  uint16
	FloorMarginUSD *big.Int
}, error) {
	return _SymbioticLzDVN.Contract.DstConfigLookup(&_SymbioticLzDVN.CallOpts, arg0)
}

// DstConfigLookup is a free data retrieval call binding the contract method 0x10187719.
//
// Solidity: function dstConfigLookup(uint32 ) view returns(uint64 gas, uint16 multiplierBps, uint128 floorMarginUSD)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) DstConfigLookup(arg0 uint32) (struct {
	Gas            uint64
	MultiplierBps  uint16
	FloorMarginUSD *big.Int
}, error) {
	return _SymbioticLzDVN.Contract.DstConfigLookup(&_SymbioticLzDVN.CallOpts, arg0)
}

// FeeLib is a free data retrieval call binding the contract method 0xbb37d00c.
//
// Solidity: function feeLib() pure returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) FeeLib(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "feeLib")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeLib is a free data retrieval call binding the contract method 0xbb37d00c.
//
// Solidity: function feeLib() pure returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNSession) FeeLib() (common.Address, error) {
	return _SymbioticLzDVN.Contract.FeeLib(&_SymbioticLzDVN.CallOpts)
}

// FeeLib is a free data retrieval call binding the contract method 0xbb37d00c.
//
// Solidity: function feeLib() pure returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) FeeLib() (common.Address, error) {
	return _SymbioticLzDVN.Contract.FeeLib(&_SymbioticLzDVN.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0x30bb3aac.
//
// Solidity: function getFee(uint32 _dstEid, uint64 _confirmations, address _sender, bytes _options) view returns(uint256 nativeFee)
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
// Solidity: function getFee(uint32 _dstEid, uint64 _confirmations, address _sender, bytes _options) view returns(uint256 nativeFee)
func (_SymbioticLzDVN *SymbioticLzDVNSession) GetFee(_dstEid uint32, _confirmations uint64, _sender common.Address, _options []byte) (*big.Int, error) {
	return _SymbioticLzDVN.Contract.GetFee(&_SymbioticLzDVN.CallOpts, _dstEid, _confirmations, _sender, _options)
}

// GetFee is a free data retrieval call binding the contract method 0x30bb3aac.
//
// Solidity: function getFee(uint32 _dstEid, uint64 _confirmations, address _sender, bytes _options) view returns(uint256 nativeFee)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) GetFee(_dstEid uint32, _confirmations uint64, _sender common.Address, _options []byte) (*big.Int, error) {
	return _SymbioticLzDVN.Contract.GetFee(&_SymbioticLzDVN.CallOpts, _dstEid, _confirmations, _sender, _options)
}

// GetFee0 is a free data retrieval call binding the contract method 0xfdb9b0f1.
//
// Solidity: function getFee(address , bytes , bytes , bytes ) view returns(uint256)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) GetFee0(opts *bind.CallOpts, arg0 common.Address, arg1 []byte, arg2 []byte, arg3 []byte) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "getFee0", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee0 is a free data retrieval call binding the contract method 0xfdb9b0f1.
//
// Solidity: function getFee(address , bytes , bytes , bytes ) view returns(uint256)
func (_SymbioticLzDVN *SymbioticLzDVNSession) GetFee0(arg0 common.Address, arg1 []byte, arg2 []byte, arg3 []byte) (*big.Int, error) {
	return _SymbioticLzDVN.Contract.GetFee0(&_SymbioticLzDVN.CallOpts, arg0, arg1, arg2, arg3)
}

// GetFee0 is a free data retrieval call binding the contract method 0xfdb9b0f1.
//
// Solidity: function getFee(address , bytes , bytes , bytes ) view returns(uint256)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) GetFee0(arg0 common.Address, arg1 []byte, arg2 []byte, arg3 []byte) (*big.Int, error) {
	return _SymbioticLzDVN.Contract.GetFee0(&_SymbioticLzDVN.CallOpts, arg0, arg1, arg2, arg3)
}

// GetSupportedOptionTypes is a free data retrieval call binding the contract method 0x26e67a37.
//
// Solidity: function getSupportedOptionTypes(uint32 ) view returns(uint8[])
func (_SymbioticLzDVN *SymbioticLzDVNCaller) GetSupportedOptionTypes(opts *bind.CallOpts, arg0 uint32) ([]uint8, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "getSupportedOptionTypes", arg0)

	if err != nil {
		return *new([]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint8)).(*[]uint8)

	return out0, err

}

// GetSupportedOptionTypes is a free data retrieval call binding the contract method 0x26e67a37.
//
// Solidity: function getSupportedOptionTypes(uint32 ) view returns(uint8[])
func (_SymbioticLzDVN *SymbioticLzDVNSession) GetSupportedOptionTypes(arg0 uint32) ([]uint8, error) {
	return _SymbioticLzDVN.Contract.GetSupportedOptionTypes(&_SymbioticLzDVN.CallOpts, arg0)
}

// GetSupportedOptionTypes is a free data retrieval call binding the contract method 0x26e67a37.
//
// Solidity: function getSupportedOptionTypes(uint32 ) view returns(uint8[])
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) GetSupportedOptionTypes(arg0 uint32) ([]uint8, error) {
	return _SymbioticLzDVN.Contract.GetSupportedOptionTypes(&_SymbioticLzDVN.CallOpts, arg0)
}

// LocalEid is a free data retrieval call binding the contract method 0x72607537.
//
// Solidity: function localEid() view returns(uint32)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) LocalEid(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "localEid")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalEid is a free data retrieval call binding the contract method 0x72607537.
//
// Solidity: function localEid() view returns(uint32)
func (_SymbioticLzDVN *SymbioticLzDVNSession) LocalEid() (uint32, error) {
	return _SymbioticLzDVN.Contract.LocalEid(&_SymbioticLzDVN.CallOpts)
}

// LocalEid is a free data retrieval call binding the contract method 0x72607537.
//
// Solidity: function localEid() view returns(uint32)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) LocalEid() (uint32, error) {
	return _SymbioticLzDVN.Contract.LocalEid(&_SymbioticLzDVN.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNSession) Owner() (common.Address, error) {
	return _SymbioticLzDVN.Contract.Owner(&_SymbioticLzDVN.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) Owner() (common.Address, error) {
	return _SymbioticLzDVN.Contract.Owner(&_SymbioticLzDVN.CallOpts)
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

// SetWorkerFeeLib is a free data retrieval call binding the contract method 0xc7b2370b.
//
// Solidity: function setWorkerFeeLib(address ) pure returns()
func (_SymbioticLzDVN *SymbioticLzDVNCaller) SetWorkerFeeLib(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "setWorkerFeeLib", arg0)

	if err != nil {
		return err
	}

	return err

}

// SetWorkerFeeLib is a free data retrieval call binding the contract method 0xc7b2370b.
//
// Solidity: function setWorkerFeeLib(address ) pure returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) SetWorkerFeeLib(arg0 common.Address) error {
	return _SymbioticLzDVN.Contract.SetWorkerFeeLib(&_SymbioticLzDVN.CallOpts, arg0)
}

// SetWorkerFeeLib is a free data retrieval call binding the contract method 0xc7b2370b.
//
// Solidity: function setWorkerFeeLib(address ) pure returns()
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) SetWorkerFeeLib(arg0 common.Address) error {
	return _SymbioticLzDVN.Contract.SetWorkerFeeLib(&_SymbioticLzDVN.CallOpts, arg0)
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

// Worker is a free data retrieval call binding the contract method 0x4d547ada.
//
// Solidity: function worker() pure returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) Worker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "worker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Worker is a free data retrieval call binding the contract method 0x4d547ada.
//
// Solidity: function worker() pure returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNSession) Worker() (common.Address, error) {
	return _SymbioticLzDVN.Contract.Worker(&_SymbioticLzDVN.CallOpts)
}

// Worker is a free data retrieval call binding the contract method 0x4d547ada.
//
// Solidity: function worker() pure returns(address)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) Worker() (common.Address, error) {
	return _SymbioticLzDVN.Contract.Worker(&_SymbioticLzDVN.CallOpts)
}

// AssignJob is a paid mutator transaction binding the contract method 0x95d376d7.
//
// Solidity: function assignJob((uint32,bytes,bytes32,uint64,address) _param, bytes _options) payable returns(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) AssignJob(opts *bind.TransactOpts, _param ILayerZeroDVNAssignJobParam, _options []byte) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "assignJob", _param, _options)
}

// AssignJob is a paid mutator transaction binding the contract method 0x95d376d7.
//
// Solidity: function assignJob((uint32,bytes,bytes32,uint64,address) _param, bytes _options) payable returns(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNSession) AssignJob(_param ILayerZeroDVNAssignJobParam, _options []byte) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.AssignJob(&_SymbioticLzDVN.TransactOpts, _param, _options)
}

// AssignJob is a paid mutator transaction binding the contract method 0x95d376d7.
//
// Solidity: function assignJob((uint32,bytes,bytes32,uint64,address) _param, bytes _options) payable returns(uint256 fee)
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) AssignJob(_param ILayerZeroDVNAssignJobParam, _options []byte) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.AssignJob(&_SymbioticLzDVN.TransactOpts, _param, _options)
}

// AssignJob0 is a paid mutator transaction binding the contract method 0xf42ed2ed.
//
// Solidity: function assignJob(address , bytes , bytes , bytes ) payable returns(uint256)
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) AssignJob0(opts *bind.TransactOpts, arg0 common.Address, arg1 []byte, arg2 []byte, arg3 []byte) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "assignJob0", arg0, arg1, arg2, arg3)
}

// AssignJob0 is a paid mutator transaction binding the contract method 0xf42ed2ed.
//
// Solidity: function assignJob(address , bytes , bytes , bytes ) payable returns(uint256)
func (_SymbioticLzDVN *SymbioticLzDVNSession) AssignJob0(arg0 common.Address, arg1 []byte, arg2 []byte, arg3 []byte) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.AssignJob0(&_SymbioticLzDVN.TransactOpts, arg0, arg1, arg2, arg3)
}

// AssignJob0 is a paid mutator transaction binding the contract method 0xf42ed2ed.
//
// Solidity: function assignJob(address , bytes , bytes , bytes ) payable returns(uint256)
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) AssignJob0(arg0 common.Address, arg1 []byte, arg2 []byte, arg3 []byte) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.AssignJob0(&_SymbioticLzDVN.TransactOpts, arg0, arg1, arg2, arg3)
}

// SetDefaultMultiplierBps is a paid mutator transaction binding the contract method 0xc358de0a.
//
// Solidity: function setDefaultMultiplierBps(uint16 ) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) SetDefaultMultiplierBps(opts *bind.TransactOpts, arg0 uint16) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "setDefaultMultiplierBps", arg0)
}

// SetDefaultMultiplierBps is a paid mutator transaction binding the contract method 0xc358de0a.
//
// Solidity: function setDefaultMultiplierBps(uint16 ) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) SetDefaultMultiplierBps(arg0 uint16) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetDefaultMultiplierBps(&_SymbioticLzDVN.TransactOpts, arg0)
}

// SetDefaultMultiplierBps is a paid mutator transaction binding the contract method 0xc358de0a.
//
// Solidity: function setDefaultMultiplierBps(uint16 ) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) SetDefaultMultiplierBps(arg0 uint16) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetDefaultMultiplierBps(&_SymbioticLzDVN.TransactOpts, arg0)
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

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address ) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) SetPriceFeed(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "setPriceFeed", arg0)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address ) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) SetPriceFeed(arg0 common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetPriceFeed(&_SymbioticLzDVN.TransactOpts, arg0)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address ) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) SetPriceFeed(arg0 common.Address) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetPriceFeed(&_SymbioticLzDVN.TransactOpts, arg0)
}

// SetSupportedOptionTypes is a paid mutator transaction binding the contract method 0xcd88b903.
//
// Solidity: function setSupportedOptionTypes(uint32 , uint8[] ) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) SetSupportedOptionTypes(opts *bind.TransactOpts, arg0 uint32, arg1 []uint8) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "setSupportedOptionTypes", arg0, arg1)
}

// SetSupportedOptionTypes is a paid mutator transaction binding the contract method 0xcd88b903.
//
// Solidity: function setSupportedOptionTypes(uint32 , uint8[] ) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) SetSupportedOptionTypes(arg0 uint32, arg1 []uint8) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetSupportedOptionTypes(&_SymbioticLzDVN.TransactOpts, arg0, arg1)
}

// SetSupportedOptionTypes is a paid mutator transaction binding the contract method 0xcd88b903.
//
// Solidity: function setSupportedOptionTypes(uint32 , uint8[] ) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) SetSupportedOptionTypes(arg0 uint32, arg1 []uint8) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.SetSupportedOptionTypes(&_SymbioticLzDVN.TransactOpts, arg0, arg1)
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
// Solidity: function withdrawFee(address , address , uint256 ) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactor) WithdrawFee(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _SymbioticLzDVN.contract.Transact(opts, "withdrawFee", arg0, arg1, arg2)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1095b6d7.
//
// Solidity: function withdrawFee(address , address , uint256 ) returns()
func (_SymbioticLzDVN *SymbioticLzDVNSession) WithdrawFee(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.WithdrawFee(&_SymbioticLzDVN.TransactOpts, arg0, arg1, arg2)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1095b6d7.
//
// Solidity: function withdrawFee(address , address , uint256 ) returns()
func (_SymbioticLzDVN *SymbioticLzDVNTransactorSession) WithdrawFee(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _SymbioticLzDVN.Contract.WithdrawFee(&_SymbioticLzDVN.TransactOpts, arg0, arg1, arg2)
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
