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

// UlnConfig is an auto generated low-level Go binding around an user-defined struct.
type UlnConfig struct {
	Confirmations        uint64
	RequiredDVNCount     uint8
	OptionalDVNCount     uint8
	OptionalDVNThreshold uint8
	RequiredDVNs         []common.Address
	OptionalDVNs         []common.Address
}

// ReceiveUlnMetaData contains all meta data concerning the ReceiveUln contract.
var ReceiveUlnMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_remoteEid\",\"type\":\"uint32\"}],\"name\":\"getUlnConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"confirmations\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"requiredDVNCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"optionalDVNCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"optionalDVNThreshold\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"requiredDVNs\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"optionalDVNs\",\"type\":\"address[]\"}],\"internalType\":\"structUlnConfig\",\"name\":\"rtnConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ReceiveUlnABI is the input ABI used to generate the binding from.
// Deprecated: Use ReceiveUlnMetaData.ABI instead.
var ReceiveUlnABI = ReceiveUlnMetaData.ABI

// ReceiveUln is an auto generated Go binding around an Ethereum contract.
type ReceiveUln struct {
	ReceiveUlnCaller     // Read-only binding to the contract
	ReceiveUlnTransactor // Write-only binding to the contract
	ReceiveUlnFilterer   // Log filterer for contract events
}

// ReceiveUlnCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiveUlnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiveUlnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiveUlnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiveUlnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiveUlnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiveUlnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiveUlnSession struct {
	Contract     *ReceiveUln       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReceiveUlnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiveUlnCallerSession struct {
	Contract *ReceiveUlnCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ReceiveUlnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiveUlnTransactorSession struct {
	Contract     *ReceiveUlnTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ReceiveUlnRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiveUlnRaw struct {
	Contract *ReceiveUln // Generic contract binding to access the raw methods on
}

// ReceiveUlnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiveUlnCallerRaw struct {
	Contract *ReceiveUlnCaller // Generic read-only contract binding to access the raw methods on
}

// ReceiveUlnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiveUlnTransactorRaw struct {
	Contract *ReceiveUlnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiveUln creates a new instance of ReceiveUln, bound to a specific deployed contract.
func NewReceiveUln(address common.Address, backend bind.ContractBackend) (*ReceiveUln, error) {
	contract, err := bindReceiveUln(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReceiveUln{ReceiveUlnCaller: ReceiveUlnCaller{contract: contract}, ReceiveUlnTransactor: ReceiveUlnTransactor{contract: contract}, ReceiveUlnFilterer: ReceiveUlnFilterer{contract: contract}}, nil
}

// NewReceiveUlnCaller creates a new read-only instance of ReceiveUln, bound to a specific deployed contract.
func NewReceiveUlnCaller(address common.Address, caller bind.ContractCaller) (*ReceiveUlnCaller, error) {
	contract, err := bindReceiveUln(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiveUlnCaller{contract: contract}, nil
}

// NewReceiveUlnTransactor creates a new write-only instance of ReceiveUln, bound to a specific deployed contract.
func NewReceiveUlnTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiveUlnTransactor, error) {
	contract, err := bindReceiveUln(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiveUlnTransactor{contract: contract}, nil
}

// NewReceiveUlnFilterer creates a new log filterer instance of ReceiveUln, bound to a specific deployed contract.
func NewReceiveUlnFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiveUlnFilterer, error) {
	contract, err := bindReceiveUln(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiveUlnFilterer{contract: contract}, nil
}

// bindReceiveUln binds a generic wrapper to an already deployed contract.
func bindReceiveUln(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReceiveUlnMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiveUln *ReceiveUlnRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiveUln.Contract.ReceiveUlnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiveUln *ReceiveUlnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiveUln.Contract.ReceiveUlnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiveUln *ReceiveUlnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiveUln.Contract.ReceiveUlnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiveUln *ReceiveUlnCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiveUln.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiveUln *ReceiveUlnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiveUln.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiveUln *ReceiveUlnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiveUln.Contract.contract.Transact(opts, method, params...)
}

// GetUlnConfig is a free data retrieval call binding the contract method 0x43ea4fa9.
//
// Solidity: function getUlnConfig(address _oapp, uint32 _remoteEid) view returns((uint64,uint8,uint8,uint8,address[],address[]) rtnConfig)
func (_ReceiveUln *ReceiveUlnCaller) GetUlnConfig(opts *bind.CallOpts, _oapp common.Address, _remoteEid uint32) (UlnConfig, error) {
	var out []interface{}
	err := _ReceiveUln.contract.Call(opts, &out, "getUlnConfig", _oapp, _remoteEid)

	if err != nil {
		return *new(UlnConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(UlnConfig)).(*UlnConfig)

	return out0, err

}

// GetUlnConfig is a free data retrieval call binding the contract method 0x43ea4fa9.
//
// Solidity: function getUlnConfig(address _oapp, uint32 _remoteEid) view returns((uint64,uint8,uint8,uint8,address[],address[]) rtnConfig)
func (_ReceiveUln *ReceiveUlnSession) GetUlnConfig(_oapp common.Address, _remoteEid uint32) (UlnConfig, error) {
	return _ReceiveUln.Contract.GetUlnConfig(&_ReceiveUln.CallOpts, _oapp, _remoteEid)
}

// GetUlnConfig is a free data retrieval call binding the contract method 0x43ea4fa9.
//
// Solidity: function getUlnConfig(address _oapp, uint32 _remoteEid) view returns((uint64,uint8,uint8,uint8,address[],address[]) rtnConfig)
func (_ReceiveUln *ReceiveUlnCallerSession) GetUlnConfig(_oapp common.Address, _remoteEid uint32) (UlnConfig, error) {
	return _ReceiveUln.Contract.GetUlnConfig(&_ReceiveUln.CallOpts, _oapp, _remoteEid)
}
