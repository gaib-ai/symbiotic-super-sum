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

// SetConfigParam is an auto generated low-level Go binding around an user-defined struct.
// type SetConfigParam struct {
// 	Eid        uint32
// 	ConfigType uint32
// 	Config     []byte
// }

// SetDefaultUlnConfigParam is an auto generated low-level Go binding around an user-defined struct.
type SetDefaultUlnConfigParam struct {
	Eid    uint32
	Config UlnConfig
}

// UlnConfig is an auto generated low-level Go binding around an user-defined struct.
type UlnConfig struct {
	Confirmations        uint64
	RequiredDVNCount     uint8
	OptionalDVNCount     uint8
	OptionalDVNThreshold uint8
	RequiredDVNs         []common.Address
	OptionalDVNs         []common.Address
}

// ReceiveUln302MetaData contains all meta data concerning the ReceiveUln302 contract.
var ReceiveUln302MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_endpoint\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"assertHeader\",\"inputs\":[{\"name\":\"_packetHeader\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_localEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"commitVerification\",\"inputs\":[{\"name\":\"_packetHeader\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_payloadHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAppUlnConfig\",\"inputs\":[{\"name\":\"_oapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_remoteEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structUlnConfig\",\"components\":[{\"name\":\"confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requiredDVNCount\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"optionalDVNCount\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"optionalDVNThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"requiredDVNs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalDVNs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConfig\",\"inputs\":[{\"name\":\"_eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_oapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_configType\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUlnConfig\",\"inputs\":[{\"name\":\"_oapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_remoteEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"rtnConfig\",\"type\":\"tuple\",\"internalType\":\"structUlnConfig\",\"components\":[{\"name\":\"confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requiredDVNCount\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"optionalDVNCount\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"optionalDVNThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"requiredDVNs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalDVNs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashLookup\",\"inputs\":[{\"name\":\"headerHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payloadHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"dvn\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedEid\",\"inputs\":[{\"name\":\"_eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messageLibType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumMessageLibType\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setConfig\",\"inputs\":[{\"name\":\"_oapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_params\",\"type\":\"tuple[]\",\"internalType\":\"structSetConfigParam[]\",\"components\":[{\"name\":\"eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configType\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaultUlnConfigs\",\"inputs\":[{\"name\":\"_params\",\"type\":\"tuple[]\",\"internalType\":\"structSetDefaultUlnConfigParam[]\",\"components\":[{\"name\":\"eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structUlnConfig\",\"components\":[{\"name\":\"confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requiredDVNCount\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"optionalDVNCount\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"optionalDVNThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"requiredDVNs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalDVNs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifiable\",\"inputs\":[{\"name\":\"_config\",\"type\":\"tuple\",\"internalType\":\"structUlnConfig\",\"components\":[{\"name\":\"confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requiredDVNCount\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"optionalDVNCount\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"optionalDVNThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"requiredDVNs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalDVNs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_headerHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_payloadHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"_packetHeader\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_payloadHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"major\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minor\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"endpointVersion\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"DefaultUlnConfigsSet\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structSetDefaultUlnConfigParam[]\",\"components\":[{\"name\":\"eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structUlnConfig\",\"components\":[{\"name\":\"confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requiredDVNCount\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"optionalDVNCount\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"optionalDVNThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"requiredDVNs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalDVNs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PayloadVerified\",\"inputs\":[{\"name\":\"dvn\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"header\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"confirmations\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"proofHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UlnConfigSet\",\"inputs\":[{\"name\":\"oapp\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"eid\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structUlnConfig\",\"components\":[{\"name\":\"confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requiredDVNCount\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"optionalDVNCount\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"optionalDVNThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"requiredDVNs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalDVNs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"LZ_MessageLib_OnlyEndpoint\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LZ_ULN_AtLeastOneDVN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LZ_ULN_InvalidConfigType\",\"inputs\":[{\"name\":\"configType\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"LZ_ULN_InvalidConfirmations\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LZ_ULN_InvalidEid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LZ_ULN_InvalidOptionalDVNCount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LZ_ULN_InvalidOptionalDVNThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LZ_ULN_InvalidPacketHeader\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LZ_ULN_InvalidPacketVersion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LZ_ULN_InvalidRequiredDVNCount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LZ_ULN_Unsorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LZ_ULN_UnsupportedEid\",\"inputs\":[{\"name\":\"eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"LZ_ULN_Verifying\",\"inputs\":[]}]",
}

// ReceiveUln302ABI is the input ABI used to generate the binding from.
// Deprecated: Use ReceiveUln302MetaData.ABI instead.
var ReceiveUln302ABI = ReceiveUln302MetaData.ABI

// ReceiveUln302 is an auto generated Go binding around an Ethereum contract.
type ReceiveUln302 struct {
	ReceiveUln302Caller     // Read-only binding to the contract
	ReceiveUln302Transactor // Write-only binding to the contract
	ReceiveUln302Filterer   // Log filterer for contract events
}

// ReceiveUln302Caller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiveUln302Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiveUln302Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiveUln302Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiveUln302Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiveUln302Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiveUln302Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiveUln302Session struct {
	Contract     *ReceiveUln302    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReceiveUln302CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiveUln302CallerSession struct {
	Contract *ReceiveUln302Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ReceiveUln302TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiveUln302TransactorSession struct {
	Contract     *ReceiveUln302Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ReceiveUln302Raw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiveUln302Raw struct {
	Contract *ReceiveUln302 // Generic contract binding to access the raw methods on
}

// ReceiveUln302CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiveUln302CallerRaw struct {
	Contract *ReceiveUln302Caller // Generic read-only contract binding to access the raw methods on
}

// ReceiveUln302TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiveUln302TransactorRaw struct {
	Contract *ReceiveUln302Transactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiveUln302 creates a new instance of ReceiveUln302, bound to a specific deployed contract.
func NewReceiveUln302(address common.Address, backend bind.ContractBackend) (*ReceiveUln302, error) {
	contract, err := bindReceiveUln302(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReceiveUln302{ReceiveUln302Caller: ReceiveUln302Caller{contract: contract}, ReceiveUln302Transactor: ReceiveUln302Transactor{contract: contract}, ReceiveUln302Filterer: ReceiveUln302Filterer{contract: contract}}, nil
}

// NewReceiveUln302Caller creates a new read-only instance of ReceiveUln302, bound to a specific deployed contract.
func NewReceiveUln302Caller(address common.Address, caller bind.ContractCaller) (*ReceiveUln302Caller, error) {
	contract, err := bindReceiveUln302(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiveUln302Caller{contract: contract}, nil
}

// NewReceiveUln302Transactor creates a new write-only instance of ReceiveUln302, bound to a specific deployed contract.
func NewReceiveUln302Transactor(address common.Address, transactor bind.ContractTransactor) (*ReceiveUln302Transactor, error) {
	contract, err := bindReceiveUln302(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiveUln302Transactor{contract: contract}, nil
}

// NewReceiveUln302Filterer creates a new log filterer instance of ReceiveUln302, bound to a specific deployed contract.
func NewReceiveUln302Filterer(address common.Address, filterer bind.ContractFilterer) (*ReceiveUln302Filterer, error) {
	contract, err := bindReceiveUln302(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiveUln302Filterer{contract: contract}, nil
}

// bindReceiveUln302 binds a generic wrapper to an already deployed contract.
func bindReceiveUln302(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReceiveUln302MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiveUln302 *ReceiveUln302Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiveUln302.Contract.ReceiveUln302Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiveUln302 *ReceiveUln302Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiveUln302.Contract.ReceiveUln302Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiveUln302 *ReceiveUln302Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiveUln302.Contract.ReceiveUln302Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiveUln302 *ReceiveUln302CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiveUln302.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiveUln302 *ReceiveUln302TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiveUln302.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiveUln302 *ReceiveUln302TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiveUln302.Contract.contract.Transact(opts, method, params...)
}

// AssertHeader is a free data retrieval call binding the contract method 0xc40ff835.
//
// Solidity: function assertHeader(bytes _packetHeader, uint32 _localEid) pure returns()
func (_ReceiveUln302 *ReceiveUln302Caller) AssertHeader(opts *bind.CallOpts, _packetHeader []byte, _localEid uint32) error {
	var out []interface{}
	err := _ReceiveUln302.contract.Call(opts, &out, "assertHeader", _packetHeader, _localEid)

	if err != nil {
		return err
	}

	return err

}

// AssertHeader is a free data retrieval call binding the contract method 0xc40ff835.
//
// Solidity: function assertHeader(bytes _packetHeader, uint32 _localEid) pure returns()
func (_ReceiveUln302 *ReceiveUln302Session) AssertHeader(_packetHeader []byte, _localEid uint32) error {
	return _ReceiveUln302.Contract.AssertHeader(&_ReceiveUln302.CallOpts, _packetHeader, _localEid)
}

// AssertHeader is a free data retrieval call binding the contract method 0xc40ff835.
//
// Solidity: function assertHeader(bytes _packetHeader, uint32 _localEid) pure returns()
func (_ReceiveUln302 *ReceiveUln302CallerSession) AssertHeader(_packetHeader []byte, _localEid uint32) error {
	return _ReceiveUln302.Contract.AssertHeader(&_ReceiveUln302.CallOpts, _packetHeader, _localEid)
}

// GetAppUlnConfig is a free data retrieval call binding the contract method 0x39e3f938.
//
// Solidity: function getAppUlnConfig(address _oapp, uint32 _remoteEid) view returns((uint64,uint8,uint8,uint8,address[],address[]))
func (_ReceiveUln302 *ReceiveUln302Caller) GetAppUlnConfig(opts *bind.CallOpts, _oapp common.Address, _remoteEid uint32) (UlnConfig, error) {
	var out []interface{}
	err := _ReceiveUln302.contract.Call(opts, &out, "getAppUlnConfig", _oapp, _remoteEid)

	if err != nil {
		return *new(UlnConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(UlnConfig)).(*UlnConfig)

	return out0, err

}

// GetAppUlnConfig is a free data retrieval call binding the contract method 0x39e3f938.
//
// Solidity: function getAppUlnConfig(address _oapp, uint32 _remoteEid) view returns((uint64,uint8,uint8,uint8,address[],address[]))
func (_ReceiveUln302 *ReceiveUln302Session) GetAppUlnConfig(_oapp common.Address, _remoteEid uint32) (UlnConfig, error) {
	return _ReceiveUln302.Contract.GetAppUlnConfig(&_ReceiveUln302.CallOpts, _oapp, _remoteEid)
}

// GetAppUlnConfig is a free data retrieval call binding the contract method 0x39e3f938.
//
// Solidity: function getAppUlnConfig(address _oapp, uint32 _remoteEid) view returns((uint64,uint8,uint8,uint8,address[],address[]))
func (_ReceiveUln302 *ReceiveUln302CallerSession) GetAppUlnConfig(_oapp common.Address, _remoteEid uint32) (UlnConfig, error) {
	return _ReceiveUln302.Contract.GetAppUlnConfig(&_ReceiveUln302.CallOpts, _oapp, _remoteEid)
}

// GetConfig is a free data retrieval call binding the contract method 0x9c33abf7.
//
// Solidity: function getConfig(uint32 _eid, address _oapp, uint32 _configType) view returns(bytes)
func (_ReceiveUln302 *ReceiveUln302Caller) GetConfig(opts *bind.CallOpts, _eid uint32, _oapp common.Address, _configType uint32) ([]byte, error) {
	var out []interface{}
	err := _ReceiveUln302.contract.Call(opts, &out, "getConfig", _eid, _oapp, _configType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0x9c33abf7.
//
// Solidity: function getConfig(uint32 _eid, address _oapp, uint32 _configType) view returns(bytes)
func (_ReceiveUln302 *ReceiveUln302Session) GetConfig(_eid uint32, _oapp common.Address, _configType uint32) ([]byte, error) {
	return _ReceiveUln302.Contract.GetConfig(&_ReceiveUln302.CallOpts, _eid, _oapp, _configType)
}

// GetConfig is a free data retrieval call binding the contract method 0x9c33abf7.
//
// Solidity: function getConfig(uint32 _eid, address _oapp, uint32 _configType) view returns(bytes)
func (_ReceiveUln302 *ReceiveUln302CallerSession) GetConfig(_eid uint32, _oapp common.Address, _configType uint32) ([]byte, error) {
	return _ReceiveUln302.Contract.GetConfig(&_ReceiveUln302.CallOpts, _eid, _oapp, _configType)
}

// GetUlnConfig is a free data retrieval call binding the contract method 0x43ea4fa9.
//
// Solidity: function getUlnConfig(address _oapp, uint32 _remoteEid) view returns((uint64,uint8,uint8,uint8,address[],address[]) rtnConfig)
func (_ReceiveUln302 *ReceiveUln302Caller) GetUlnConfig(opts *bind.CallOpts, _oapp common.Address, _remoteEid uint32) (UlnConfig, error) {
	var out []interface{}
	err := _ReceiveUln302.contract.Call(opts, &out, "getUlnConfig", _oapp, _remoteEid)

	if err != nil {
		return *new(UlnConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(UlnConfig)).(*UlnConfig)

	return out0, err

}

// GetUlnConfig is a free data retrieval call binding the contract method 0x43ea4fa9.
//
// Solidity: function getUlnConfig(address _oapp, uint32 _remoteEid) view returns((uint64,uint8,uint8,uint8,address[],address[]) rtnConfig)
func (_ReceiveUln302 *ReceiveUln302Session) GetUlnConfig(_oapp common.Address, _remoteEid uint32) (UlnConfig, error) {
	return _ReceiveUln302.Contract.GetUlnConfig(&_ReceiveUln302.CallOpts, _oapp, _remoteEid)
}

// GetUlnConfig is a free data retrieval call binding the contract method 0x43ea4fa9.
//
// Solidity: function getUlnConfig(address _oapp, uint32 _remoteEid) view returns((uint64,uint8,uint8,uint8,address[],address[]) rtnConfig)
func (_ReceiveUln302 *ReceiveUln302CallerSession) GetUlnConfig(_oapp common.Address, _remoteEid uint32) (UlnConfig, error) {
	return _ReceiveUln302.Contract.GetUlnConfig(&_ReceiveUln302.CallOpts, _oapp, _remoteEid)
}

// HashLookup is a free data retrieval call binding the contract method 0x3c782a52.
//
// Solidity: function hashLookup(bytes32 headerHash, bytes32 payloadHash, address dvn) view returns(bool submitted, uint64 confirmations)
func (_ReceiveUln302 *ReceiveUln302Caller) HashLookup(opts *bind.CallOpts, headerHash [32]byte, payloadHash [32]byte, dvn common.Address) (struct {
	Submitted     bool
	Confirmations uint64
}, error) {
	var out []interface{}
	err := _ReceiveUln302.contract.Call(opts, &out, "hashLookup", headerHash, payloadHash, dvn)

	outstruct := new(struct {
		Submitted     bool
		Confirmations uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Submitted = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Confirmations = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// HashLookup is a free data retrieval call binding the contract method 0x3c782a52.
//
// Solidity: function hashLookup(bytes32 headerHash, bytes32 payloadHash, address dvn) view returns(bool submitted, uint64 confirmations)
func (_ReceiveUln302 *ReceiveUln302Session) HashLookup(headerHash [32]byte, payloadHash [32]byte, dvn common.Address) (struct {
	Submitted     bool
	Confirmations uint64
}, error) {
	return _ReceiveUln302.Contract.HashLookup(&_ReceiveUln302.CallOpts, headerHash, payloadHash, dvn)
}

// HashLookup is a free data retrieval call binding the contract method 0x3c782a52.
//
// Solidity: function hashLookup(bytes32 headerHash, bytes32 payloadHash, address dvn) view returns(bool submitted, uint64 confirmations)
func (_ReceiveUln302 *ReceiveUln302CallerSession) HashLookup(headerHash [32]byte, payloadHash [32]byte, dvn common.Address) (struct {
	Submitted     bool
	Confirmations uint64
}, error) {
	return _ReceiveUln302.Contract.HashLookup(&_ReceiveUln302.CallOpts, headerHash, payloadHash, dvn)
}

// IsSupportedEid is a free data retrieval call binding the contract method 0x6750cd4c.
//
// Solidity: function isSupportedEid(uint32 _eid) view returns(bool)
func (_ReceiveUln302 *ReceiveUln302Caller) IsSupportedEid(opts *bind.CallOpts, _eid uint32) (bool, error) {
	var out []interface{}
	err := _ReceiveUln302.contract.Call(opts, &out, "isSupportedEid", _eid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedEid is a free data retrieval call binding the contract method 0x6750cd4c.
//
// Solidity: function isSupportedEid(uint32 _eid) view returns(bool)
func (_ReceiveUln302 *ReceiveUln302Session) IsSupportedEid(_eid uint32) (bool, error) {
	return _ReceiveUln302.Contract.IsSupportedEid(&_ReceiveUln302.CallOpts, _eid)
}

// IsSupportedEid is a free data retrieval call binding the contract method 0x6750cd4c.
//
// Solidity: function isSupportedEid(uint32 _eid) view returns(bool)
func (_ReceiveUln302 *ReceiveUln302CallerSession) IsSupportedEid(_eid uint32) (bool, error) {
	return _ReceiveUln302.Contract.IsSupportedEid(&_ReceiveUln302.CallOpts, _eid)
}

// MessageLibType is a free data retrieval call binding the contract method 0x1881d94d.
//
// Solidity: function messageLibType() pure returns(uint8)
func (_ReceiveUln302 *ReceiveUln302Caller) MessageLibType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ReceiveUln302.contract.Call(opts, &out, "messageLibType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MessageLibType is a free data retrieval call binding the contract method 0x1881d94d.
//
// Solidity: function messageLibType() pure returns(uint8)
func (_ReceiveUln302 *ReceiveUln302Session) MessageLibType() (uint8, error) {
	return _ReceiveUln302.Contract.MessageLibType(&_ReceiveUln302.CallOpts)
}

// MessageLibType is a free data retrieval call binding the contract method 0x1881d94d.
//
// Solidity: function messageLibType() pure returns(uint8)
func (_ReceiveUln302 *ReceiveUln302CallerSession) MessageLibType() (uint8, error) {
	return _ReceiveUln302.Contract.MessageLibType(&_ReceiveUln302.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReceiveUln302 *ReceiveUln302Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReceiveUln302.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReceiveUln302 *ReceiveUln302Session) Owner() (common.Address, error) {
	return _ReceiveUln302.Contract.Owner(&_ReceiveUln302.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReceiveUln302 *ReceiveUln302CallerSession) Owner() (common.Address, error) {
	return _ReceiveUln302.Contract.Owner(&_ReceiveUln302.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ReceiveUln302 *ReceiveUln302Caller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ReceiveUln302.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ReceiveUln302 *ReceiveUln302Session) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ReceiveUln302.Contract.SupportsInterface(&_ReceiveUln302.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ReceiveUln302 *ReceiveUln302CallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ReceiveUln302.Contract.SupportsInterface(&_ReceiveUln302.CallOpts, _interfaceId)
}

// Verifiable is a free data retrieval call binding the contract method 0xe084d952.
//
// Solidity: function verifiable((uint64,uint8,uint8,uint8,address[],address[]) _config, bytes32 _headerHash, bytes32 _payloadHash) view returns(bool)
func (_ReceiveUln302 *ReceiveUln302Caller) Verifiable(opts *bind.CallOpts, _config UlnConfig, _headerHash [32]byte, _payloadHash [32]byte) (bool, error) {
	var out []interface{}
	err := _ReceiveUln302.contract.Call(opts, &out, "verifiable", _config, _headerHash, _payloadHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verifiable is a free data retrieval call binding the contract method 0xe084d952.
//
// Solidity: function verifiable((uint64,uint8,uint8,uint8,address[],address[]) _config, bytes32 _headerHash, bytes32 _payloadHash) view returns(bool)
func (_ReceiveUln302 *ReceiveUln302Session) Verifiable(_config UlnConfig, _headerHash [32]byte, _payloadHash [32]byte) (bool, error) {
	return _ReceiveUln302.Contract.Verifiable(&_ReceiveUln302.CallOpts, _config, _headerHash, _payloadHash)
}

// Verifiable is a free data retrieval call binding the contract method 0xe084d952.
//
// Solidity: function verifiable((uint64,uint8,uint8,uint8,address[],address[]) _config, bytes32 _headerHash, bytes32 _payloadHash) view returns(bool)
func (_ReceiveUln302 *ReceiveUln302CallerSession) Verifiable(_config UlnConfig, _headerHash [32]byte, _payloadHash [32]byte) (bool, error) {
	return _ReceiveUln302.Contract.Verifiable(&_ReceiveUln302.CallOpts, _config, _headerHash, _payloadHash)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint64 major, uint8 minor, uint8 endpointVersion)
func (_ReceiveUln302 *ReceiveUln302Caller) Version(opts *bind.CallOpts) (struct {
	Major           uint64
	Minor           uint8
	EndpointVersion uint8
}, error) {
	var out []interface{}
	err := _ReceiveUln302.contract.Call(opts, &out, "version")

	outstruct := new(struct {
		Major           uint64
		Minor           uint8
		EndpointVersion uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Major = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Minor = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.EndpointVersion = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint64 major, uint8 minor, uint8 endpointVersion)
func (_ReceiveUln302 *ReceiveUln302Session) Version() (struct {
	Major           uint64
	Minor           uint8
	EndpointVersion uint8
}, error) {
	return _ReceiveUln302.Contract.Version(&_ReceiveUln302.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint64 major, uint8 minor, uint8 endpointVersion)
func (_ReceiveUln302 *ReceiveUln302CallerSession) Version() (struct {
	Major           uint64
	Minor           uint8
	EndpointVersion uint8
}, error) {
	return _ReceiveUln302.Contract.Version(&_ReceiveUln302.CallOpts)
}

// CommitVerification is a paid mutator transaction binding the contract method 0x0894edf1.
//
// Solidity: function commitVerification(bytes _packetHeader, bytes32 _payloadHash) returns()
func (_ReceiveUln302 *ReceiveUln302Transactor) CommitVerification(opts *bind.TransactOpts, _packetHeader []byte, _payloadHash [32]byte) (*types.Transaction, error) {
	return _ReceiveUln302.contract.Transact(opts, "commitVerification", _packetHeader, _payloadHash)
}

// CommitVerification is a paid mutator transaction binding the contract method 0x0894edf1.
//
// Solidity: function commitVerification(bytes _packetHeader, bytes32 _payloadHash) returns()
func (_ReceiveUln302 *ReceiveUln302Session) CommitVerification(_packetHeader []byte, _payloadHash [32]byte) (*types.Transaction, error) {
	return _ReceiveUln302.Contract.CommitVerification(&_ReceiveUln302.TransactOpts, _packetHeader, _payloadHash)
}

// CommitVerification is a paid mutator transaction binding the contract method 0x0894edf1.
//
// Solidity: function commitVerification(bytes _packetHeader, bytes32 _payloadHash) returns()
func (_ReceiveUln302 *ReceiveUln302TransactorSession) CommitVerification(_packetHeader []byte, _payloadHash [32]byte) (*types.Transaction, error) {
	return _ReceiveUln302.Contract.CommitVerification(&_ReceiveUln302.TransactOpts, _packetHeader, _payloadHash)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReceiveUln302 *ReceiveUln302Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiveUln302.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReceiveUln302 *ReceiveUln302Session) RenounceOwnership() (*types.Transaction, error) {
	return _ReceiveUln302.Contract.RenounceOwnership(&_ReceiveUln302.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReceiveUln302 *ReceiveUln302TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ReceiveUln302.Contract.RenounceOwnership(&_ReceiveUln302.TransactOpts)
}

// SetConfig is a paid mutator transaction binding the contract method 0x20efd722.
//
// Solidity: function setConfig(address _oapp, (uint32,uint32,bytes)[] _params) returns()
func (_ReceiveUln302 *ReceiveUln302Transactor) SetConfig(opts *bind.TransactOpts, _oapp common.Address, _params []SetConfigParam) (*types.Transaction, error) {
	return _ReceiveUln302.contract.Transact(opts, "setConfig", _oapp, _params)
}

// SetConfig is a paid mutator transaction binding the contract method 0x20efd722.
//
// Solidity: function setConfig(address _oapp, (uint32,uint32,bytes)[] _params) returns()
func (_ReceiveUln302 *ReceiveUln302Session) SetConfig(_oapp common.Address, _params []SetConfigParam) (*types.Transaction, error) {
	return _ReceiveUln302.Contract.SetConfig(&_ReceiveUln302.TransactOpts, _oapp, _params)
}

// SetConfig is a paid mutator transaction binding the contract method 0x20efd722.
//
// Solidity: function setConfig(address _oapp, (uint32,uint32,bytes)[] _params) returns()
func (_ReceiveUln302 *ReceiveUln302TransactorSession) SetConfig(_oapp common.Address, _params []SetConfigParam) (*types.Transaction, error) {
	return _ReceiveUln302.Contract.SetConfig(&_ReceiveUln302.TransactOpts, _oapp, _params)
}

// SetDefaultUlnConfigs is a paid mutator transaction binding the contract method 0x29460b0b.
//
// Solidity: function setDefaultUlnConfigs((uint32,(uint64,uint8,uint8,uint8,address[],address[]))[] _params) returns()
func (_ReceiveUln302 *ReceiveUln302Transactor) SetDefaultUlnConfigs(opts *bind.TransactOpts, _params []SetDefaultUlnConfigParam) (*types.Transaction, error) {
	return _ReceiveUln302.contract.Transact(opts, "setDefaultUlnConfigs", _params)
}

// SetDefaultUlnConfigs is a paid mutator transaction binding the contract method 0x29460b0b.
//
// Solidity: function setDefaultUlnConfigs((uint32,(uint64,uint8,uint8,uint8,address[],address[]))[] _params) returns()
func (_ReceiveUln302 *ReceiveUln302Session) SetDefaultUlnConfigs(_params []SetDefaultUlnConfigParam) (*types.Transaction, error) {
	return _ReceiveUln302.Contract.SetDefaultUlnConfigs(&_ReceiveUln302.TransactOpts, _params)
}

// SetDefaultUlnConfigs is a paid mutator transaction binding the contract method 0x29460b0b.
//
// Solidity: function setDefaultUlnConfigs((uint32,(uint64,uint8,uint8,uint8,address[],address[]))[] _params) returns()
func (_ReceiveUln302 *ReceiveUln302TransactorSession) SetDefaultUlnConfigs(_params []SetDefaultUlnConfigParam) (*types.Transaction, error) {
	return _ReceiveUln302.Contract.SetDefaultUlnConfigs(&_ReceiveUln302.TransactOpts, _params)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReceiveUln302 *ReceiveUln302Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ReceiveUln302.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReceiveUln302 *ReceiveUln302Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReceiveUln302.Contract.TransferOwnership(&_ReceiveUln302.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReceiveUln302 *ReceiveUln302TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReceiveUln302.Contract.TransferOwnership(&_ReceiveUln302.TransactOpts, newOwner)
}

// Verify is a paid mutator transaction binding the contract method 0x0223536e.
//
// Solidity: function verify(bytes _packetHeader, bytes32 _payloadHash, uint64 _confirmations) returns()
func (_ReceiveUln302 *ReceiveUln302Transactor) Verify(opts *bind.TransactOpts, _packetHeader []byte, _payloadHash [32]byte, _confirmations uint64) (*types.Transaction, error) {
	return _ReceiveUln302.contract.Transact(opts, "verify", _packetHeader, _payloadHash, _confirmations)
}

// Verify is a paid mutator transaction binding the contract method 0x0223536e.
//
// Solidity: function verify(bytes _packetHeader, bytes32 _payloadHash, uint64 _confirmations) returns()
func (_ReceiveUln302 *ReceiveUln302Session) Verify(_packetHeader []byte, _payloadHash [32]byte, _confirmations uint64) (*types.Transaction, error) {
	return _ReceiveUln302.Contract.Verify(&_ReceiveUln302.TransactOpts, _packetHeader, _payloadHash, _confirmations)
}

// Verify is a paid mutator transaction binding the contract method 0x0223536e.
//
// Solidity: function verify(bytes _packetHeader, bytes32 _payloadHash, uint64 _confirmations) returns()
func (_ReceiveUln302 *ReceiveUln302TransactorSession) Verify(_packetHeader []byte, _payloadHash [32]byte, _confirmations uint64) (*types.Transaction, error) {
	return _ReceiveUln302.Contract.Verify(&_ReceiveUln302.TransactOpts, _packetHeader, _payloadHash, _confirmations)
}

// ReceiveUln302DefaultUlnConfigsSetIterator is returned from FilterDefaultUlnConfigsSet and is used to iterate over the raw logs and unpacked data for DefaultUlnConfigsSet events raised by the ReceiveUln302 contract.
type ReceiveUln302DefaultUlnConfigsSetIterator struct {
	Event *ReceiveUln302DefaultUlnConfigsSet // Event containing the contract specifics and raw log

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
func (it *ReceiveUln302DefaultUlnConfigsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiveUln302DefaultUlnConfigsSet)
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
		it.Event = new(ReceiveUln302DefaultUlnConfigsSet)
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
func (it *ReceiveUln302DefaultUlnConfigsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiveUln302DefaultUlnConfigsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiveUln302DefaultUlnConfigsSet represents a DefaultUlnConfigsSet event raised by the ReceiveUln302 contract.
type ReceiveUln302DefaultUlnConfigsSet struct {
	Params []SetDefaultUlnConfigParam
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDefaultUlnConfigsSet is a free log retrieval operation binding the contract event 0xaaf3aaa0c11056e86ac56eb653e25b005ca1a7d4dcd21ba24647f7ab63f3b560.
//
// Solidity: event DefaultUlnConfigsSet((uint32,(uint64,uint8,uint8,uint8,address[],address[]))[] params)
func (_ReceiveUln302 *ReceiveUln302Filterer) FilterDefaultUlnConfigsSet(opts *bind.FilterOpts) (*ReceiveUln302DefaultUlnConfigsSetIterator, error) {

	logs, sub, err := _ReceiveUln302.contract.FilterLogs(opts, "DefaultUlnConfigsSet")
	if err != nil {
		return nil, err
	}
	return &ReceiveUln302DefaultUlnConfigsSetIterator{contract: _ReceiveUln302.contract, event: "DefaultUlnConfigsSet", logs: logs, sub: sub}, nil
}

// WatchDefaultUlnConfigsSet is a free log subscription operation binding the contract event 0xaaf3aaa0c11056e86ac56eb653e25b005ca1a7d4dcd21ba24647f7ab63f3b560.
//
// Solidity: event DefaultUlnConfigsSet((uint32,(uint64,uint8,uint8,uint8,address[],address[]))[] params)
func (_ReceiveUln302 *ReceiveUln302Filterer) WatchDefaultUlnConfigsSet(opts *bind.WatchOpts, sink chan<- *ReceiveUln302DefaultUlnConfigsSet) (event.Subscription, error) {

	logs, sub, err := _ReceiveUln302.contract.WatchLogs(opts, "DefaultUlnConfigsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiveUln302DefaultUlnConfigsSet)
				if err := _ReceiveUln302.contract.UnpackLog(event, "DefaultUlnConfigsSet", log); err != nil {
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

// ParseDefaultUlnConfigsSet is a log parse operation binding the contract event 0xaaf3aaa0c11056e86ac56eb653e25b005ca1a7d4dcd21ba24647f7ab63f3b560.
//
// Solidity: event DefaultUlnConfigsSet((uint32,(uint64,uint8,uint8,uint8,address[],address[]))[] params)
func (_ReceiveUln302 *ReceiveUln302Filterer) ParseDefaultUlnConfigsSet(log types.Log) (*ReceiveUln302DefaultUlnConfigsSet, error) {
	event := new(ReceiveUln302DefaultUlnConfigsSet)
	if err := _ReceiveUln302.contract.UnpackLog(event, "DefaultUlnConfigsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiveUln302OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ReceiveUln302 contract.
type ReceiveUln302OwnershipTransferredIterator struct {
	Event *ReceiveUln302OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ReceiveUln302OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiveUln302OwnershipTransferred)
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
		it.Event = new(ReceiveUln302OwnershipTransferred)
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
func (it *ReceiveUln302OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiveUln302OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiveUln302OwnershipTransferred represents a OwnershipTransferred event raised by the ReceiveUln302 contract.
type ReceiveUln302OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ReceiveUln302 *ReceiveUln302Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ReceiveUln302OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ReceiveUln302.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ReceiveUln302OwnershipTransferredIterator{contract: _ReceiveUln302.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ReceiveUln302 *ReceiveUln302Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ReceiveUln302OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ReceiveUln302.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiveUln302OwnershipTransferred)
				if err := _ReceiveUln302.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ReceiveUln302 *ReceiveUln302Filterer) ParseOwnershipTransferred(log types.Log) (*ReceiveUln302OwnershipTransferred, error) {
	event := new(ReceiveUln302OwnershipTransferred)
	if err := _ReceiveUln302.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiveUln302PayloadVerifiedIterator is returned from FilterPayloadVerified and is used to iterate over the raw logs and unpacked data for PayloadVerified events raised by the ReceiveUln302 contract.
type ReceiveUln302PayloadVerifiedIterator struct {
	Event *ReceiveUln302PayloadVerified // Event containing the contract specifics and raw log

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
func (it *ReceiveUln302PayloadVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiveUln302PayloadVerified)
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
		it.Event = new(ReceiveUln302PayloadVerified)
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
func (it *ReceiveUln302PayloadVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiveUln302PayloadVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiveUln302PayloadVerified represents a PayloadVerified event raised by the ReceiveUln302 contract.
type ReceiveUln302PayloadVerified struct {
	Dvn           common.Address
	Header        []byte
	Confirmations *big.Int
	ProofHash     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPayloadVerified is a free log retrieval operation binding the contract event 0x2cb0eed7538baeae4c6fde038c0fd0384d27de0dd55a228c65847bda6aa1ab56.
//
// Solidity: event PayloadVerified(address dvn, bytes header, uint256 confirmations, bytes32 proofHash)
func (_ReceiveUln302 *ReceiveUln302Filterer) FilterPayloadVerified(opts *bind.FilterOpts) (*ReceiveUln302PayloadVerifiedIterator, error) {

	logs, sub, err := _ReceiveUln302.contract.FilterLogs(opts, "PayloadVerified")
	if err != nil {
		return nil, err
	}
	return &ReceiveUln302PayloadVerifiedIterator{contract: _ReceiveUln302.contract, event: "PayloadVerified", logs: logs, sub: sub}, nil
}

// WatchPayloadVerified is a free log subscription operation binding the contract event 0x2cb0eed7538baeae4c6fde038c0fd0384d27de0dd55a228c65847bda6aa1ab56.
//
// Solidity: event PayloadVerified(address dvn, bytes header, uint256 confirmations, bytes32 proofHash)
func (_ReceiveUln302 *ReceiveUln302Filterer) WatchPayloadVerified(opts *bind.WatchOpts, sink chan<- *ReceiveUln302PayloadVerified) (event.Subscription, error) {

	logs, sub, err := _ReceiveUln302.contract.WatchLogs(opts, "PayloadVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiveUln302PayloadVerified)
				if err := _ReceiveUln302.contract.UnpackLog(event, "PayloadVerified", log); err != nil {
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

// ParsePayloadVerified is a log parse operation binding the contract event 0x2cb0eed7538baeae4c6fde038c0fd0384d27de0dd55a228c65847bda6aa1ab56.
//
// Solidity: event PayloadVerified(address dvn, bytes header, uint256 confirmations, bytes32 proofHash)
func (_ReceiveUln302 *ReceiveUln302Filterer) ParsePayloadVerified(log types.Log) (*ReceiveUln302PayloadVerified, error) {
	event := new(ReceiveUln302PayloadVerified)
	if err := _ReceiveUln302.contract.UnpackLog(event, "PayloadVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiveUln302UlnConfigSetIterator is returned from FilterUlnConfigSet and is used to iterate over the raw logs and unpacked data for UlnConfigSet events raised by the ReceiveUln302 contract.
type ReceiveUln302UlnConfigSetIterator struct {
	Event *ReceiveUln302UlnConfigSet // Event containing the contract specifics and raw log

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
func (it *ReceiveUln302UlnConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiveUln302UlnConfigSet)
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
		it.Event = new(ReceiveUln302UlnConfigSet)
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
func (it *ReceiveUln302UlnConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiveUln302UlnConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiveUln302UlnConfigSet represents a UlnConfigSet event raised by the ReceiveUln302 contract.
type ReceiveUln302UlnConfigSet struct {
	Oapp   common.Address
	Eid    uint32
	Config UlnConfig
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUlnConfigSet is a free log retrieval operation binding the contract event 0x82118522aa536ac0e96cc5c689407ae42b89d592aa133890a01f1509842f5081.
//
// Solidity: event UlnConfigSet(address oapp, uint32 eid, (uint64,uint8,uint8,uint8,address[],address[]) config)
func (_ReceiveUln302 *ReceiveUln302Filterer) FilterUlnConfigSet(opts *bind.FilterOpts) (*ReceiveUln302UlnConfigSetIterator, error) {

	logs, sub, err := _ReceiveUln302.contract.FilterLogs(opts, "UlnConfigSet")
	if err != nil {
		return nil, err
	}
	return &ReceiveUln302UlnConfigSetIterator{contract: _ReceiveUln302.contract, event: "UlnConfigSet", logs: logs, sub: sub}, nil
}

// WatchUlnConfigSet is a free log subscription operation binding the contract event 0x82118522aa536ac0e96cc5c689407ae42b89d592aa133890a01f1509842f5081.
//
// Solidity: event UlnConfigSet(address oapp, uint32 eid, (uint64,uint8,uint8,uint8,address[],address[]) config)
func (_ReceiveUln302 *ReceiveUln302Filterer) WatchUlnConfigSet(opts *bind.WatchOpts, sink chan<- *ReceiveUln302UlnConfigSet) (event.Subscription, error) {

	logs, sub, err := _ReceiveUln302.contract.WatchLogs(opts, "UlnConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiveUln302UlnConfigSet)
				if err := _ReceiveUln302.contract.UnpackLog(event, "UlnConfigSet", log); err != nil {
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

// ParseUlnConfigSet is a log parse operation binding the contract event 0x82118522aa536ac0e96cc5c689407ae42b89d592aa133890a01f1509842f5081.
//
// Solidity: event UlnConfigSet(address oapp, uint32 eid, (uint64,uint8,uint8,uint8,address[],address[]) config)
func (_ReceiveUln302 *ReceiveUln302Filterer) ParseUlnConfigSet(log types.Log) (*ReceiveUln302UlnConfigSet, error) {
	event := new(ReceiveUln302UlnConfigSet)
	if err := _ReceiveUln302.contract.UnpackLog(event, "UlnConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
