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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localEid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_settlementContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiveUlnContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"multiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint128\",\"name\":\"floorMarginUSD\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structIDVN.DstConfigParam[]\",\"name\":\"params\",\"type\":\"tuple[]\"}],\"name\":\"SetDstConfig\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"packetHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"payloadHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"confirmations\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"internalType\":\"structILayerZeroDVN.AssignJobParam\",\"name\":\"_param\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"}],\"name\":\"assignJob\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_dstEid\",\"type\":\"uint32\"}],\"name\":\"dstConfig\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"multiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint128\",\"name\":\"floorMarginUSD\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"dstConfigLookup\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"multiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint128\",\"name\":\"floorMarginUSD\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeLib\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_dstEid\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localEid\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveUlnContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settlementContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"multiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint128\",\"name\":\"floorMarginUSD\",\"type\":\"uint128\"}],\"internalType\":\"structIDVN.DstConfigParam[]\",\"name\":\"_params\",\"type\":\"tuple[]\"}],\"name\":\"setDstConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setWorkerFeeLib\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_packetHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_confirmations\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_symbioticProof\",\"type\":\"bytes\"}],\"name\":\"verifyWithSymbiotic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"worker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
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
// Solidity: function getFee(uint32 _dstEid, uint64 , address , bytes ) view returns(uint256 nativeFee)
func (_SymbioticLzDVN *SymbioticLzDVNCaller) GetFee(opts *bind.CallOpts, _dstEid uint32, arg1 uint64, arg2 common.Address, arg3 []byte) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticLzDVN.contract.Call(opts, &out, "getFee", _dstEid, arg1, arg2, arg3)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0x30bb3aac.
//
// Solidity: function getFee(uint32 _dstEid, uint64 , address , bytes ) view returns(uint256 nativeFee)
func (_SymbioticLzDVN *SymbioticLzDVNSession) GetFee(_dstEid uint32, arg1 uint64, arg2 common.Address, arg3 []byte) (*big.Int, error) {
	return _SymbioticLzDVN.Contract.GetFee(&_SymbioticLzDVN.CallOpts, _dstEid, arg1, arg2, arg3)
}

// GetFee is a free data retrieval call binding the contract method 0x30bb3aac.
//
// Solidity: function getFee(uint32 _dstEid, uint64 , address , bytes ) view returns(uint256 nativeFee)
func (_SymbioticLzDVN *SymbioticLzDVNCallerSession) GetFee(_dstEid uint32, arg1 uint64, arg2 common.Address, arg3 []byte) (*big.Int, error) {
	return _SymbioticLzDVN.Contract.GetFee(&_SymbioticLzDVN.CallOpts, _dstEid, arg1, arg2, arg3)
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
