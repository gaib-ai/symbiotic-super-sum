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

// MessagingReceipt is an auto generated low-level Go binding around an user-defined struct.
type MessagingReceipt struct {
	Guid  [32]byte
	Nonce uint64
	Fee   *big.Int
}

// Origin is an auto generated low-level Go binding around an user-defined struct.
type Origin struct {
	SrcEid uint32
	Sender [32]byte
	Nonce  uint64
}

// EndpointV2MetaData contains all meta data concerning the EndpointV2 contract.
var EndpointV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_lzToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Endpoint_AlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oapp\",\"type\":\"address\"}],\"name\":\"Endpoint_DefaultLibraryAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Endpoint_InvalidAccounting\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Endpoint_InvalidBlockReason\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"Endpoint_InvalidEndpointId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Endpoint_InvalidNativePayload\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"Endpoint_InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Endpoint_InvalidRelayer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Endpoint_InvalidSendLib\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Endpoint_InvalidTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Endpoint_InvalidUln\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"}],\"name\":\"Endpoint_MsgValueNotMatchFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oapp\",\"type\":\"address\"}],\"name\":\"Endpoint_NoDefaultSendLibrary\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"}],\"name\":\"Endpoint_PacketAlreadyDelivered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"}],\"name\":\"Endpoint_PacketNotVerified\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"}],\"name\":\"Endpoint_PacketNotVerifiedOrSent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Endpoint_PathAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Endpoint_PathNotSet\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"}],\"name\":\"Endpoint_SameNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Endpoint_Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oapp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"uln\",\"type\":\"address\"}],\"name\":\"DefaultReceiveLibrarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oapp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sendLib\",\"type\":\"address\"}],\"name\":\"DefaultSendLibrarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeCollector\",\"type\":\"address\"}],\"name\":\"LzTokenFeeCollectorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCompositionFee\",\"type\":\"address\"}],\"name\":\"LzTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"uln\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NativeFeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRelayer\",\"type\":\"address\"}],\"name\":\"NewRelayer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"PacketBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"}],\"name\":\"PacketDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sendLib\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"}],\"name\":\"PacketFeeTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"packetHeader\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"PacketSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"payloadHash\",\"type\":\"bytes32\"}],\"name\":\"PacketVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oapp\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteEid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldReceiveLib\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newReceiveLib\",\"type\":\"address\"}],\"name\":\"ReceiveLibrarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oapp\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteEid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timeout\",\"type\":\"uint64\"}],\"name\":\"ReceiveLibraryTimeoutSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oapp\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteEid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldSendLib\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSendLib\",\"type\":\"address\"}],\"name\":\"SendLibrarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nativeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"TokenFeeRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"uln\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenFeeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"TreasuryUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCK_REASON_BOUNCED\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_REASON_DUPLICATE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_VERSION\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eid\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_packet\",\"type\":\"bytes\"}],\"name\":\"getNativeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_remoteEid\",\"type\":\"uint32\"}],\"name\":\"getReceiveLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiveLib\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDefault\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_remoteEid\",\"type\":\"uint32\"}],\"name\":\"getSendLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sendLib\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDefault\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"}],\"name\":\"inboundNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"}],\"name\":\"isDefaultReceiveLibrary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"}],\"name\":\"isDefaultSendLibrary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isRegisteredLibrary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"}],\"name\":\"lazyInboundNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lzToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lzTokenFeeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"messageComposed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"composer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nativePayload\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"outboundNonce\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"packetDelivered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"packetSent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"packetVerified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uln\",\"type\":\"address\"}],\"name\":\"removeDefaultReceiveLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"}],\"name\":\"removeDefaultSendLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uln\",\"type\":\"address\"}],\"name\":\"setDefaultReceiveLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sendLib\",\"type\":\"address\"}],\"name\":\"setDefaultSendLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newCollector\",\"type\":\"address\"}],\"name\":\"setLzTokenFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_remoteEid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_newReceiveLib\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_gracePeriod\",\"type\":\"uint64\"}],\"name\":\"setReceiveLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_remoteEid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_newSendLib\",\"type\":\"address\"}],\"name\":\"setSendLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_composer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"}],\"name\":\"skip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EndpointV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use EndpointV2MetaData.ABI instead.
var EndpointV2ABI = EndpointV2MetaData.ABI

// EndpointV2 is an auto generated Go binding around an Ethereum contract.
type EndpointV2 struct {
	EndpointV2Caller     // Read-only binding to the contract
	EndpointV2Transactor // Write-only binding to the contract
	EndpointV2Filterer   // Log filterer for contract events
}

// EndpointV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type EndpointV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EndpointV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type EndpointV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EndpointV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EndpointV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EndpointV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EndpointV2Session struct {
	Contract     *EndpointV2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EndpointV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EndpointV2CallerSession struct {
	Contract *EndpointV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EndpointV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EndpointV2TransactorSession struct {
	Contract     *EndpointV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EndpointV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type EndpointV2Raw struct {
	Contract *EndpointV2 // Generic contract binding to access the raw methods on
}

// EndpointV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EndpointV2CallerRaw struct {
	Contract *EndpointV2Caller // Generic read-only contract binding to access the raw methods on
}

// EndpointV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EndpointV2TransactorRaw struct {
	Contract *EndpointV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewEndpointV2 creates a new instance of EndpointV2, bound to a specific deployed contract.
func NewEndpointV2(address common.Address, backend bind.ContractBackend) (*EndpointV2, error) {
	contract, err := bindEndpointV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EndpointV2{EndpointV2Caller: EndpointV2Caller{contract: contract}, EndpointV2Transactor: EndpointV2Transactor{contract: contract}, EndpointV2Filterer: EndpointV2Filterer{contract: contract}}, nil
}

// NewEndpointV2Caller creates a new read-only instance of EndpointV2, bound to a specific deployed contract.
func NewEndpointV2Caller(address common.Address, caller bind.ContractCaller) (*EndpointV2Caller, error) {
	contract, err := bindEndpointV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EndpointV2Caller{contract: contract}, nil
}

// NewEndpointV2Transactor creates a new write-only instance of EndpointV2, bound to a specific deployed contract.
func NewEndpointV2Transactor(address common.Address, transactor bind.ContractTransactor) (*EndpointV2Transactor, error) {
	contract, err := bindEndpointV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EndpointV2Transactor{contract: contract}, nil
}

// NewEndpointV2Filterer creates a new log filterer instance of EndpointV2, bound to a specific deployed contract.
func NewEndpointV2Filterer(address common.Address, filterer bind.ContractFilterer) (*EndpointV2Filterer, error) {
	contract, err := bindEndpointV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EndpointV2Filterer{contract: contract}, nil
}

// bindEndpointV2 binds a generic wrapper to an already deployed contract.
func bindEndpointV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EndpointV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EndpointV2 *EndpointV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EndpointV2.Contract.EndpointV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EndpointV2 *EndpointV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EndpointV2.Contract.EndpointV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EndpointV2 *EndpointV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EndpointV2.Contract.EndpointV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EndpointV2 *EndpointV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EndpointV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EndpointV2 *EndpointV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EndpointV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EndpointV2 *EndpointV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EndpointV2.Contract.contract.Transact(opts, method, params...)
}

// BLOCKREASONBOUNCED is a free data retrieval call binding the contract method 0xa46b26d9.
//
// Solidity: function BLOCK_REASON_BOUNCED() view returns(bytes32)
func (_EndpointV2 *EndpointV2Caller) BLOCKREASONBOUNCED(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "BLOCK_REASON_BOUNCED")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKREASONBOUNCED is a free data retrieval call binding the contract method 0xa46b26d9.
//
// Solidity: function BLOCK_REASON_BOUNCED() view returns(bytes32)
func (_EndpointV2 *EndpointV2Session) BLOCKREASONBOUNCED() ([32]byte, error) {
	return _EndpointV2.Contract.BLOCKREASONBOUNCED(&_EndpointV2.CallOpts)
}

// BLOCKREASONBOUNCED is a free data retrieval call binding the contract method 0xa46b26d9.
//
// Solidity: function BLOCK_REASON_BOUNCED() view returns(bytes32)
func (_EndpointV2 *EndpointV2CallerSession) BLOCKREASONBOUNCED() ([32]byte, error) {
	return _EndpointV2.Contract.BLOCKREASONBOUNCED(&_EndpointV2.CallOpts)
}

// BLOCKREASONDUPLICATE is a free data retrieval call binding the contract method 0x8169fd1d.
//
// Solidity: function BLOCK_REASON_DUPLICATE() view returns(bytes32)
func (_EndpointV2 *EndpointV2Caller) BLOCKREASONDUPLICATE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "BLOCK_REASON_DUPLICATE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKREASONDUPLICATE is a free data retrieval call binding the contract method 0x8169fd1d.
//
// Solidity: function BLOCK_REASON_DUPLICATE() view returns(bytes32)
func (_EndpointV2 *EndpointV2Session) BLOCKREASONDUPLICATE() ([32]byte, error) {
	return _EndpointV2.Contract.BLOCKREASONDUPLICATE(&_EndpointV2.CallOpts)
}

// BLOCKREASONDUPLICATE is a free data retrieval call binding the contract method 0x8169fd1d.
//
// Solidity: function BLOCK_REASON_DUPLICATE() view returns(bytes32)
func (_EndpointV2 *EndpointV2CallerSession) BLOCKREASONDUPLICATE() ([32]byte, error) {
	return _EndpointV2.Contract.BLOCKREASONDUPLICATE(&_EndpointV2.CallOpts)
}

// DEFAULTVERSION is a free data retrieval call binding the contract method 0x24ba3f2c.
//
// Solidity: function DEFAULT_VERSION() view returns(uint16)
func (_EndpointV2 *EndpointV2Caller) DEFAULTVERSION(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "DEFAULT_VERSION")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DEFAULTVERSION is a free data retrieval call binding the contract method 0x24ba3f2c.
//
// Solidity: function DEFAULT_VERSION() view returns(uint16)
func (_EndpointV2 *EndpointV2Session) DEFAULTVERSION() (uint16, error) {
	return _EndpointV2.Contract.DEFAULTVERSION(&_EndpointV2.CallOpts)
}

// DEFAULTVERSION is a free data retrieval call binding the contract method 0x24ba3f2c.
//
// Solidity: function DEFAULT_VERSION() view returns(uint16)
func (_EndpointV2 *EndpointV2CallerSession) DEFAULTVERSION() (uint16, error) {
	return _EndpointV2.Contract.DEFAULTVERSION(&_EndpointV2.CallOpts)
}

// Eid is a free data retrieval call binding the contract method 0x416ecebf.
//
// Solidity: function eid() view returns(uint32)
func (_EndpointV2 *EndpointV2Caller) Eid(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "eid")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Eid is a free data retrieval call binding the contract method 0x416ecebf.
//
// Solidity: function eid() view returns(uint32)
func (_EndpointV2 *EndpointV2Session) Eid() (uint32, error) {
	return _EndpointV2.Contract.Eid(&_EndpointV2.CallOpts)
}

// Eid is a free data retrieval call binding the contract method 0x416ecebf.
//
// Solidity: function eid() view returns(uint32)
func (_EndpointV2 *EndpointV2CallerSession) Eid() (uint32, error) {
	return _EndpointV2.Contract.Eid(&_EndpointV2.CallOpts)
}

// GetNativeFee is a free data retrieval call binding the contract method 0xbce4dd56.
//
// Solidity: function getNativeFee(bytes _packet) view returns(uint256)
func (_EndpointV2 *EndpointV2Caller) GetNativeFee(opts *bind.CallOpts, _packet []byte) (*big.Int, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "getNativeFee", _packet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNativeFee is a free data retrieval call binding the contract method 0xbce4dd56.
//
// Solidity: function getNativeFee(bytes _packet) view returns(uint256)
func (_EndpointV2 *EndpointV2Session) GetNativeFee(_packet []byte) (*big.Int, error) {
	return _EndpointV2.Contract.GetNativeFee(&_EndpointV2.CallOpts, _packet)
}

// GetNativeFee is a free data retrieval call binding the contract method 0xbce4dd56.
//
// Solidity: function getNativeFee(bytes _packet) view returns(uint256)
func (_EndpointV2 *EndpointV2CallerSession) GetNativeFee(_packet []byte) (*big.Int, error) {
	return _EndpointV2.Contract.GetNativeFee(&_EndpointV2.CallOpts, _packet)
}

// GetReceiveLibrary is a free data retrieval call binding the contract method 0x402f8468.
//
// Solidity: function getReceiveLibrary(address _oapp, uint32 _remoteEid) view returns(address receiveLib, bool isDefault)
func (_EndpointV2 *EndpointV2Caller) GetReceiveLibrary(opts *bind.CallOpts, _oapp common.Address, _remoteEid uint32) (struct {
	ReceiveLib common.Address
	IsDefault  bool
}, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "getReceiveLibrary", _oapp, _remoteEid)

	outstruct := new(struct {
		ReceiveLib common.Address
		IsDefault  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReceiveLib = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsDefault = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetReceiveLibrary is a free data retrieval call binding the contract method 0x402f8468.
//
// Solidity: function getReceiveLibrary(address _oapp, uint32 _remoteEid) view returns(address receiveLib, bool isDefault)
func (_EndpointV2 *EndpointV2Session) GetReceiveLibrary(_oapp common.Address, _remoteEid uint32) (struct {
	ReceiveLib common.Address
	IsDefault  bool
}, error) {
	return _EndpointV2.Contract.GetReceiveLibrary(&_EndpointV2.CallOpts, _oapp, _remoteEid)
}

// GetReceiveLibrary is a free data retrieval call binding the contract method 0x402f8468.
//
// Solidity: function getReceiveLibrary(address _oapp, uint32 _remoteEid) view returns(address receiveLib, bool isDefault)
func (_EndpointV2 *EndpointV2CallerSession) GetReceiveLibrary(_oapp common.Address, _remoteEid uint32) (struct {
	ReceiveLib common.Address
	IsDefault  bool
}, error) {
	return _EndpointV2.Contract.GetReceiveLibrary(&_EndpointV2.CallOpts, _oapp, _remoteEid)
}

// GetSendLibrary is a free data retrieval call binding the contract method 0xb96a277f.
//
// Solidity: function getSendLibrary(address _oapp, uint32 _remoteEid) view returns(address sendLib, bool isDefault)
func (_EndpointV2 *EndpointV2Caller) GetSendLibrary(opts *bind.CallOpts, _oapp common.Address, _remoteEid uint32) (struct {
	SendLib   common.Address
	IsDefault bool
}, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "getSendLibrary", _oapp, _remoteEid)

	outstruct := new(struct {
		SendLib   common.Address
		IsDefault bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SendLib = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsDefault = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetSendLibrary is a free data retrieval call binding the contract method 0xb96a277f.
//
// Solidity: function getSendLibrary(address _oapp, uint32 _remoteEid) view returns(address sendLib, bool isDefault)
func (_EndpointV2 *EndpointV2Session) GetSendLibrary(_oapp common.Address, _remoteEid uint32) (struct {
	SendLib   common.Address
	IsDefault bool
}, error) {
	return _EndpointV2.Contract.GetSendLibrary(&_EndpointV2.CallOpts, _oapp, _remoteEid)
}

// GetSendLibrary is a free data retrieval call binding the contract method 0xb96a277f.
//
// Solidity: function getSendLibrary(address _oapp, uint32 _remoteEid) view returns(address sendLib, bool isDefault)
func (_EndpointV2 *EndpointV2CallerSession) GetSendLibrary(_oapp common.Address, _remoteEid uint32) (struct {
	SendLib   common.Address
	IsDefault bool
}, error) {
	return _EndpointV2.Contract.GetSendLibrary(&_EndpointV2.CallOpts, _oapp, _remoteEid)
}

// InboundNonce is a free data retrieval call binding the contract method 0x11a83f17.
//
// Solidity: function inboundNonce((uint32,bytes32,uint64) _origin) view returns(uint64)
func (_EndpointV2 *EndpointV2Caller) InboundNonce(opts *bind.CallOpts, _origin Origin) (uint64, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "inboundNonce", _origin)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InboundNonce is a free data retrieval call binding the contract method 0x11a83f17.
//
// Solidity: function inboundNonce((uint32,bytes32,uint64) _origin) view returns(uint64)
func (_EndpointV2 *EndpointV2Session) InboundNonce(_origin Origin) (uint64, error) {
	return _EndpointV2.Contract.InboundNonce(&_EndpointV2.CallOpts, _origin)
}

// InboundNonce is a free data retrieval call binding the contract method 0x11a83f17.
//
// Solidity: function inboundNonce((uint32,bytes32,uint64) _origin) view returns(uint64)
func (_EndpointV2 *EndpointV2CallerSession) InboundNonce(_origin Origin) (uint64, error) {
	return _EndpointV2.Contract.InboundNonce(&_EndpointV2.CallOpts, _origin)
}

// IsDefaultReceiveLibrary is a free data retrieval call binding the contract method 0xc0033efe.
//
// Solidity: function isDefaultReceiveLibrary(address _oapp) view returns(bool)
func (_EndpointV2 *EndpointV2Caller) IsDefaultReceiveLibrary(opts *bind.CallOpts, _oapp common.Address) (bool, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "isDefaultReceiveLibrary", _oapp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDefaultReceiveLibrary is a free data retrieval call binding the contract method 0xc0033efe.
//
// Solidity: function isDefaultReceiveLibrary(address _oapp) view returns(bool)
func (_EndpointV2 *EndpointV2Session) IsDefaultReceiveLibrary(_oapp common.Address) (bool, error) {
	return _EndpointV2.Contract.IsDefaultReceiveLibrary(&_EndpointV2.CallOpts, _oapp)
}

// IsDefaultReceiveLibrary is a free data retrieval call binding the contract method 0xc0033efe.
//
// Solidity: function isDefaultReceiveLibrary(address _oapp) view returns(bool)
func (_EndpointV2 *EndpointV2CallerSession) IsDefaultReceiveLibrary(_oapp common.Address) (bool, error) {
	return _EndpointV2.Contract.IsDefaultReceiveLibrary(&_EndpointV2.CallOpts, _oapp)
}

// IsDefaultSendLibrary is a free data retrieval call binding the contract method 0xe9aa2d16.
//
// Solidity: function isDefaultSendLibrary(address _oapp) view returns(bool)
func (_EndpointV2 *EndpointV2Caller) IsDefaultSendLibrary(opts *bind.CallOpts, _oapp common.Address) (bool, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "isDefaultSendLibrary", _oapp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDefaultSendLibrary is a free data retrieval call binding the contract method 0xe9aa2d16.
//
// Solidity: function isDefaultSendLibrary(address _oapp) view returns(bool)
func (_EndpointV2 *EndpointV2Session) IsDefaultSendLibrary(_oapp common.Address) (bool, error) {
	return _EndpointV2.Contract.IsDefaultSendLibrary(&_EndpointV2.CallOpts, _oapp)
}

// IsDefaultSendLibrary is a free data retrieval call binding the contract method 0xe9aa2d16.
//
// Solidity: function isDefaultSendLibrary(address _oapp) view returns(bool)
func (_EndpointV2 *EndpointV2CallerSession) IsDefaultSendLibrary(_oapp common.Address) (bool, error) {
	return _EndpointV2.Contract.IsDefaultSendLibrary(&_EndpointV2.CallOpts, _oapp)
}

// IsRegisteredLibrary is a free data retrieval call binding the contract method 0xdc706a62.
//
// Solidity: function isRegisteredLibrary(address _addr) view returns(bool)
func (_EndpointV2 *EndpointV2Caller) IsRegisteredLibrary(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "isRegisteredLibrary", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegisteredLibrary is a free data retrieval call binding the contract method 0xdc706a62.
//
// Solidity: function isRegisteredLibrary(address _addr) view returns(bool)
func (_EndpointV2 *EndpointV2Session) IsRegisteredLibrary(_addr common.Address) (bool, error) {
	return _EndpointV2.Contract.IsRegisteredLibrary(&_EndpointV2.CallOpts, _addr)
}

// IsRegisteredLibrary is a free data retrieval call binding the contract method 0xdc706a62.
//
// Solidity: function isRegisteredLibrary(address _addr) view returns(bool)
func (_EndpointV2 *EndpointV2CallerSession) IsRegisteredLibrary(_addr common.Address) (bool, error) {
	return _EndpointV2.Contract.IsRegisteredLibrary(&_EndpointV2.CallOpts, _addr)
}

// LazyInboundNonce is a free data retrieval call binding the contract method 0x5f13d4fa.
//
// Solidity: function lazyInboundNonce((uint32,bytes32,uint64) _origin) view returns(uint64)
func (_EndpointV2 *EndpointV2Caller) LazyInboundNonce(opts *bind.CallOpts, _origin Origin) (uint64, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "lazyInboundNonce", _origin)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LazyInboundNonce is a free data retrieval call binding the contract method 0x5f13d4fa.
//
// Solidity: function lazyInboundNonce((uint32,bytes32,uint64) _origin) view returns(uint64)
func (_EndpointV2 *EndpointV2Session) LazyInboundNonce(_origin Origin) (uint64, error) {
	return _EndpointV2.Contract.LazyInboundNonce(&_EndpointV2.CallOpts, _origin)
}

// LazyInboundNonce is a free data retrieval call binding the contract method 0x5f13d4fa.
//
// Solidity: function lazyInboundNonce((uint32,bytes32,uint64) _origin) view returns(uint64)
func (_EndpointV2 *EndpointV2CallerSession) LazyInboundNonce(_origin Origin) (uint64, error) {
	return _EndpointV2.Contract.LazyInboundNonce(&_EndpointV2.CallOpts, _origin)
}

// LzToken is a free data retrieval call binding the contract method 0xe4fe1d94.
//
// Solidity: function lzToken() view returns(address)
func (_EndpointV2 *EndpointV2Caller) LzToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "lzToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LzToken is a free data retrieval call binding the contract method 0xe4fe1d94.
//
// Solidity: function lzToken() view returns(address)
func (_EndpointV2 *EndpointV2Session) LzToken() (common.Address, error) {
	return _EndpointV2.Contract.LzToken(&_EndpointV2.CallOpts)
}

// LzToken is a free data retrieval call binding the contract method 0xe4fe1d94.
//
// Solidity: function lzToken() view returns(address)
func (_EndpointV2 *EndpointV2CallerSession) LzToken() (common.Address, error) {
	return _EndpointV2.Contract.LzToken(&_EndpointV2.CallOpts)
}

// LzTokenFeeCollector is a free data retrieval call binding the contract method 0xd0edba47.
//
// Solidity: function lzTokenFeeCollector() view returns(address)
func (_EndpointV2 *EndpointV2Caller) LzTokenFeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "lzTokenFeeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LzTokenFeeCollector is a free data retrieval call binding the contract method 0xd0edba47.
//
// Solidity: function lzTokenFeeCollector() view returns(address)
func (_EndpointV2 *EndpointV2Session) LzTokenFeeCollector() (common.Address, error) {
	return _EndpointV2.Contract.LzTokenFeeCollector(&_EndpointV2.CallOpts)
}

// LzTokenFeeCollector is a free data retrieval call binding the contract method 0xd0edba47.
//
// Solidity: function lzTokenFeeCollector() view returns(address)
func (_EndpointV2 *EndpointV2CallerSession) LzTokenFeeCollector() (common.Address, error) {
	return _EndpointV2.Contract.LzTokenFeeCollector(&_EndpointV2.CallOpts)
}

// MessageComposed is a free data retrieval call binding the contract method 0x29266437.
//
// Solidity: function messageComposed(bytes32 ) view returns(address receiver, bytes32 guid, address composer, uint256 index)
func (_EndpointV2 *EndpointV2Caller) MessageComposed(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Receiver common.Address
	Guid     [32]byte
	Composer common.Address
	Index    *big.Int
}, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "messageComposed", arg0)

	outstruct := new(struct {
		Receiver common.Address
		Guid     [32]byte
		Composer common.Address
		Index    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Guid = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Composer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Index = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MessageComposed is a free data retrieval call binding the contract method 0x29266437.
//
// Solidity: function messageComposed(bytes32 ) view returns(address receiver, bytes32 guid, address composer, uint256 index)
func (_EndpointV2 *EndpointV2Session) MessageComposed(arg0 [32]byte) (struct {
	Receiver common.Address
	Guid     [32]byte
	Composer common.Address
	Index    *big.Int
}, error) {
	return _EndpointV2.Contract.MessageComposed(&_EndpointV2.CallOpts, arg0)
}

// MessageComposed is a free data retrieval call binding the contract method 0x29266437.
//
// Solidity: function messageComposed(bytes32 ) view returns(address receiver, bytes32 guid, address composer, uint256 index)
func (_EndpointV2 *EndpointV2CallerSession) MessageComposed(arg0 [32]byte) (struct {
	Receiver common.Address
	Guid     [32]byte
	Composer common.Address
	Index    *big.Int
}, error) {
	return _EndpointV2.Contract.MessageComposed(&_EndpointV2.CallOpts, arg0)
}

// NativePayload is a free data retrieval call binding the contract method 0xd0382e95.
//
// Solidity: function nativePayload(bytes32 ) view returns(bytes)
func (_EndpointV2 *EndpointV2Caller) NativePayload(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "nativePayload", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// NativePayload is a free data retrieval call binding the contract method 0xd0382e95.
//
// Solidity: function nativePayload(bytes32 ) view returns(bytes)
func (_EndpointV2 *EndpointV2Session) NativePayload(arg0 [32]byte) ([]byte, error) {
	return _EndpointV2.Contract.NativePayload(&_EndpointV2.CallOpts, arg0)
}

// NativePayload is a free data retrieval call binding the contract method 0xd0382e95.
//
// Solidity: function nativePayload(bytes32 ) view returns(bytes)
func (_EndpointV2 *EndpointV2CallerSession) NativePayload(arg0 [32]byte) ([]byte, error) {
	return _EndpointV2.Contract.NativePayload(&_EndpointV2.CallOpts, arg0)
}

// OutboundNonce is a free data retrieval call binding the contract method 0x511d54eb.
//
// Solidity: function outboundNonce(address , uint64 ) view returns(bytes32)
func (_EndpointV2 *EndpointV2Caller) OutboundNonce(opts *bind.CallOpts, arg0 common.Address, arg1 uint64) ([32]byte, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "outboundNonce", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OutboundNonce is a free data retrieval call binding the contract method 0x511d54eb.
//
// Solidity: function outboundNonce(address , uint64 ) view returns(bytes32)
func (_EndpointV2 *EndpointV2Session) OutboundNonce(arg0 common.Address, arg1 uint64) ([32]byte, error) {
	return _EndpointV2.Contract.OutboundNonce(&_EndpointV2.CallOpts, arg0, arg1)
}

// OutboundNonce is a free data retrieval call binding the contract method 0x511d54eb.
//
// Solidity: function outboundNonce(address , uint64 ) view returns(bytes32)
func (_EndpointV2 *EndpointV2CallerSession) OutboundNonce(arg0 common.Address, arg1 uint64) ([32]byte, error) {
	return _EndpointV2.Contract.OutboundNonce(&_EndpointV2.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EndpointV2 *EndpointV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EndpointV2 *EndpointV2Session) Owner() (common.Address, error) {
	return _EndpointV2.Contract.Owner(&_EndpointV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EndpointV2 *EndpointV2CallerSession) Owner() (common.Address, error) {
	return _EndpointV2.Contract.Owner(&_EndpointV2.CallOpts)
}

// PacketDelivered is a free data retrieval call binding the contract method 0x4eb9e16b.
//
// Solidity: function packetDelivered((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) view returns(bool)
func (_EndpointV2 *EndpointV2Caller) PacketDelivered(opts *bind.CallOpts, _origin Origin, _receiver common.Address, _payloadHash [32]byte) (bool, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "packetDelivered", _origin, _receiver, _payloadHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PacketDelivered is a free data retrieval call binding the contract method 0x4eb9e16b.
//
// Solidity: function packetDelivered((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) view returns(bool)
func (_EndpointV2 *EndpointV2Session) PacketDelivered(_origin Origin, _receiver common.Address, _payloadHash [32]byte) (bool, error) {
	return _EndpointV2.Contract.PacketDelivered(&_EndpointV2.CallOpts, _origin, _receiver, _payloadHash)
}

// PacketDelivered is a free data retrieval call binding the contract method 0x4eb9e16b.
//
// Solidity: function packetDelivered((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) view returns(bool)
func (_EndpointV2 *EndpointV2CallerSession) PacketDelivered(_origin Origin, _receiver common.Address, _payloadHash [32]byte) (bool, error) {
	return _EndpointV2.Contract.PacketDelivered(&_EndpointV2.CallOpts, _origin, _receiver, _payloadHash)
}

// PacketSent is a free data retrieval call binding the contract method 0xb4a0c460.
//
// Solidity: function packetSent((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_EndpointV2 *EndpointV2Caller) PacketSent(opts *bind.CallOpts, _origin Origin, _receiver common.Address) (bool, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "packetSent", _origin, _receiver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PacketSent is a free data retrieval call binding the contract method 0xb4a0c460.
//
// Solidity: function packetSent((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_EndpointV2 *EndpointV2Session) PacketSent(_origin Origin, _receiver common.Address) (bool, error) {
	return _EndpointV2.Contract.PacketSent(&_EndpointV2.CallOpts, _origin, _receiver)
}

// PacketSent is a free data retrieval call binding the contract method 0xb4a0c460.
//
// Solidity: function packetSent((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_EndpointV2 *EndpointV2CallerSession) PacketSent(_origin Origin, _receiver common.Address) (bool, error) {
	return _EndpointV2.Contract.PacketSent(&_EndpointV2.CallOpts, _origin, _receiver)
}

// PacketVerified is a free data retrieval call binding the contract method 0x2ef342a7.
//
// Solidity: function packetVerified((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) view returns(bool)
func (_EndpointV2 *EndpointV2Caller) PacketVerified(opts *bind.CallOpts, _origin Origin, _receiver common.Address, _payloadHash [32]byte) (bool, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "packetVerified", _origin, _receiver, _payloadHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PacketVerified is a free data retrieval call binding the contract method 0x2ef342a7.
//
// Solidity: function packetVerified((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) view returns(bool)
func (_EndpointV2 *EndpointV2Session) PacketVerified(_origin Origin, _receiver common.Address, _payloadHash [32]byte) (bool, error) {
	return _EndpointV2.Contract.PacketVerified(&_EndpointV2.CallOpts, _origin, _receiver, _payloadHash)
}

// PacketVerified is a free data retrieval call binding the contract method 0x2ef342a7.
//
// Solidity: function packetVerified((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) view returns(bool)
func (_EndpointV2 *EndpointV2CallerSession) PacketVerified(_origin Origin, _receiver common.Address, _payloadHash [32]byte) (bool, error) {
	return _EndpointV2.Contract.PacketVerified(&_EndpointV2.CallOpts, _origin, _receiver, _payloadHash)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EndpointV2 *EndpointV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EndpointV2 *EndpointV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EndpointV2.Contract.SupportsInterface(&_EndpointV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EndpointV2 *EndpointV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EndpointV2.Contract.SupportsInterface(&_EndpointV2.CallOpts, interfaceId)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_EndpointV2 *EndpointV2Caller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_EndpointV2 *EndpointV2Session) Treasury() (common.Address, error) {
	return _EndpointV2.Contract.Treasury(&_EndpointV2.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_EndpointV2 *EndpointV2CallerSession) Treasury() (common.Address, error) {
	return _EndpointV2.Contract.Treasury(&_EndpointV2.CallOpts)
}

// RemoveDefaultReceiveLibrary is a paid mutator transaction binding the contract method 0x8faf8d1b.
//
// Solidity: function removeDefaultReceiveLibrary(address _oapp, address _uln) returns()
func (_EndpointV2 *EndpointV2Transactor) RemoveDefaultReceiveLibrary(opts *bind.TransactOpts, _oapp common.Address, _uln common.Address) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "removeDefaultReceiveLibrary", _oapp, _uln)
}

// RemoveDefaultReceiveLibrary is a paid mutator transaction binding the contract method 0x8faf8d1b.
//
// Solidity: function removeDefaultReceiveLibrary(address _oapp, address _uln) returns()
func (_EndpointV2 *EndpointV2Session) RemoveDefaultReceiveLibrary(_oapp common.Address, _uln common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.RemoveDefaultReceiveLibrary(&_EndpointV2.TransactOpts, _oapp, _uln)
}

// RemoveDefaultReceiveLibrary is a paid mutator transaction binding the contract method 0x8faf8d1b.
//
// Solidity: function removeDefaultReceiveLibrary(address _oapp, address _uln) returns()
func (_EndpointV2 *EndpointV2TransactorSession) RemoveDefaultReceiveLibrary(_oapp common.Address, _uln common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.RemoveDefaultReceiveLibrary(&_EndpointV2.TransactOpts, _oapp, _uln)
}

// RemoveDefaultSendLibrary is a paid mutator transaction binding the contract method 0x7d1924c7.
//
// Solidity: function removeDefaultSendLibrary(address _oapp) returns()
func (_EndpointV2 *EndpointV2Transactor) RemoveDefaultSendLibrary(opts *bind.TransactOpts, _oapp common.Address) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "removeDefaultSendLibrary", _oapp)
}

// RemoveDefaultSendLibrary is a paid mutator transaction binding the contract method 0x7d1924c7.
//
// Solidity: function removeDefaultSendLibrary(address _oapp) returns()
func (_EndpointV2 *EndpointV2Session) RemoveDefaultSendLibrary(_oapp common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.RemoveDefaultSendLibrary(&_EndpointV2.TransactOpts, _oapp)
}

// RemoveDefaultSendLibrary is a paid mutator transaction binding the contract method 0x7d1924c7.
//
// Solidity: function removeDefaultSendLibrary(address _oapp) returns()
func (_EndpointV2 *EndpointV2TransactorSession) RemoveDefaultSendLibrary(_oapp common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.RemoveDefaultSendLibrary(&_EndpointV2.TransactOpts, _oapp)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EndpointV2 *EndpointV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EndpointV2 *EndpointV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _EndpointV2.Contract.RenounceOwnership(&_EndpointV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EndpointV2 *EndpointV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EndpointV2.Contract.RenounceOwnership(&_EndpointV2.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x7f0a3bf9.
//
// Solidity: function send(bytes _message, bytes _options) payable returns((bytes32,uint64,uint256) receipt)
func (_EndpointV2 *EndpointV2Transactor) Send(opts *bind.TransactOpts, _message []byte, _options []byte) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "send", _message, _options)
}

// Send is a paid mutator transaction binding the contract method 0x7f0a3bf9.
//
// Solidity: function send(bytes _message, bytes _options) payable returns((bytes32,uint64,uint256) receipt)
func (_EndpointV2 *EndpointV2Session) Send(_message []byte, _options []byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.Send(&_EndpointV2.TransactOpts, _message, _options)
}

// Send is a paid mutator transaction binding the contract method 0x7f0a3bf9.
//
// Solidity: function send(bytes _message, bytes _options) payable returns((bytes32,uint64,uint256) receipt)
func (_EndpointV2 *EndpointV2TransactorSession) Send(_message []byte, _options []byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.Send(&_EndpointV2.TransactOpts, _message, _options)
}

// SetDefaultReceiveLibrary is a paid mutator transaction binding the contract method 0xd86e58c1.
//
// Solidity: function setDefaultReceiveLibrary(address _oapp, address _uln) returns()
func (_EndpointV2 *EndpointV2Transactor) SetDefaultReceiveLibrary(opts *bind.TransactOpts, _oapp common.Address, _uln common.Address) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "setDefaultReceiveLibrary", _oapp, _uln)
}

// SetDefaultReceiveLibrary is a paid mutator transaction binding the contract method 0xd86e58c1.
//
// Solidity: function setDefaultReceiveLibrary(address _oapp, address _uln) returns()
func (_EndpointV2 *EndpointV2Session) SetDefaultReceiveLibrary(_oapp common.Address, _uln common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetDefaultReceiveLibrary(&_EndpointV2.TransactOpts, _oapp, _uln)
}

// SetDefaultReceiveLibrary is a paid mutator transaction binding the contract method 0xd86e58c1.
//
// Solidity: function setDefaultReceiveLibrary(address _oapp, address _uln) returns()
func (_EndpointV2 *EndpointV2TransactorSession) SetDefaultReceiveLibrary(_oapp common.Address, _uln common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetDefaultReceiveLibrary(&_EndpointV2.TransactOpts, _oapp, _uln)
}

// SetDefaultSendLibrary is a paid mutator transaction binding the contract method 0xb7afb04b.
//
// Solidity: function setDefaultSendLibrary(address _oapp, address _sendLib) returns()
func (_EndpointV2 *EndpointV2Transactor) SetDefaultSendLibrary(opts *bind.TransactOpts, _oapp common.Address, _sendLib common.Address) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "setDefaultSendLibrary", _oapp, _sendLib)
}

// SetDefaultSendLibrary is a paid mutator transaction binding the contract method 0xb7afb04b.
//
// Solidity: function setDefaultSendLibrary(address _oapp, address _sendLib) returns()
func (_EndpointV2 *EndpointV2Session) SetDefaultSendLibrary(_oapp common.Address, _sendLib common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetDefaultSendLibrary(&_EndpointV2.TransactOpts, _oapp, _sendLib)
}

// SetDefaultSendLibrary is a paid mutator transaction binding the contract method 0xb7afb04b.
//
// Solidity: function setDefaultSendLibrary(address _oapp, address _sendLib) returns()
func (_EndpointV2 *EndpointV2TransactorSession) SetDefaultSendLibrary(_oapp common.Address, _sendLib common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetDefaultSendLibrary(&_EndpointV2.TransactOpts, _oapp, _sendLib)
}

// SetLzTokenFeeCollector is a paid mutator transaction binding the contract method 0xa8f8cf92.
//
// Solidity: function setLzTokenFeeCollector(address _newCollector) returns()
func (_EndpointV2 *EndpointV2Transactor) SetLzTokenFeeCollector(opts *bind.TransactOpts, _newCollector common.Address) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "setLzTokenFeeCollector", _newCollector)
}

// SetLzTokenFeeCollector is a paid mutator transaction binding the contract method 0xa8f8cf92.
//
// Solidity: function setLzTokenFeeCollector(address _newCollector) returns()
func (_EndpointV2 *EndpointV2Session) SetLzTokenFeeCollector(_newCollector common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetLzTokenFeeCollector(&_EndpointV2.TransactOpts, _newCollector)
}

// SetLzTokenFeeCollector is a paid mutator transaction binding the contract method 0xa8f8cf92.
//
// Solidity: function setLzTokenFeeCollector(address _newCollector) returns()
func (_EndpointV2 *EndpointV2TransactorSession) SetLzTokenFeeCollector(_newCollector common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetLzTokenFeeCollector(&_EndpointV2.TransactOpts, _newCollector)
}

// SetReceiveLibrary is a paid mutator transaction binding the contract method 0x2a64b005.
//
// Solidity: function setReceiveLibrary(address _oapp, uint32 _remoteEid, address _newReceiveLib, uint64 _gracePeriod) returns()
func (_EndpointV2 *EndpointV2Transactor) SetReceiveLibrary(opts *bind.TransactOpts, _oapp common.Address, _remoteEid uint32, _newReceiveLib common.Address, _gracePeriod uint64) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "setReceiveLibrary", _oapp, _remoteEid, _newReceiveLib, _gracePeriod)
}

// SetReceiveLibrary is a paid mutator transaction binding the contract method 0x2a64b005.
//
// Solidity: function setReceiveLibrary(address _oapp, uint32 _remoteEid, address _newReceiveLib, uint64 _gracePeriod) returns()
func (_EndpointV2 *EndpointV2Session) SetReceiveLibrary(_oapp common.Address, _remoteEid uint32, _newReceiveLib common.Address, _gracePeriod uint64) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetReceiveLibrary(&_EndpointV2.TransactOpts, _oapp, _remoteEid, _newReceiveLib, _gracePeriod)
}

// SetReceiveLibrary is a paid mutator transaction binding the contract method 0x2a64b005.
//
// Solidity: function setReceiveLibrary(address _oapp, uint32 _remoteEid, address _newReceiveLib, uint64 _gracePeriod) returns()
func (_EndpointV2 *EndpointV2TransactorSession) SetReceiveLibrary(_oapp common.Address, _remoteEid uint32, _newReceiveLib common.Address, _gracePeriod uint64) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetReceiveLibrary(&_EndpointV2.TransactOpts, _oapp, _remoteEid, _newReceiveLib, _gracePeriod)
}

// SetSendLibrary is a paid mutator transaction binding the contract method 0x9535ff30.
//
// Solidity: function setSendLibrary(address _oapp, uint32 _remoteEid, address _newSendLib) returns()
func (_EndpointV2 *EndpointV2Transactor) SetSendLibrary(opts *bind.TransactOpts, _oapp common.Address, _remoteEid uint32, _newSendLib common.Address) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "setSendLibrary", _oapp, _remoteEid, _newSendLib)
}

// SetSendLibrary is a paid mutator transaction binding the contract method 0x9535ff30.
//
// Solidity: function setSendLibrary(address _oapp, uint32 _remoteEid, address _newSendLib) returns()
func (_EndpointV2 *EndpointV2Session) SetSendLibrary(_oapp common.Address, _remoteEid uint32, _newSendLib common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetSendLibrary(&_EndpointV2.TransactOpts, _oapp, _remoteEid, _newSendLib)
}

// SetSendLibrary is a paid mutator transaction binding the contract method 0x9535ff30.
//
// Solidity: function setSendLibrary(address _oapp, uint32 _remoteEid, address _newSendLib) returns()
func (_EndpointV2 *EndpointV2TransactorSession) SetSendLibrary(_oapp common.Address, _remoteEid uint32, _newSendLib common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetSendLibrary(&_EndpointV2.TransactOpts, _oapp, _remoteEid, _newSendLib)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _newTreasury) returns()
func (_EndpointV2 *EndpointV2Transactor) SetTreasury(opts *bind.TransactOpts, _newTreasury common.Address) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "setTreasury", _newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _newTreasury) returns()
func (_EndpointV2 *EndpointV2Session) SetTreasury(_newTreasury common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetTreasury(&_EndpointV2.TransactOpts, _newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _newTreasury) returns()
func (_EndpointV2 *EndpointV2TransactorSession) SetTreasury(_newTreasury common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetTreasury(&_EndpointV2.TransactOpts, _newTreasury)
}

// Skip is a paid mutator transaction binding the contract method 0x20702a5c.
//
// Solidity: function skip((uint32,bytes32,uint64) _origin, address _receiver, bytes _message, address _composer, bytes32 _guid) returns()
func (_EndpointV2 *EndpointV2Transactor) Skip(opts *bind.TransactOpts, _origin Origin, _receiver common.Address, _message []byte, _composer common.Address, _guid [32]byte) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "skip", _origin, _receiver, _message, _composer, _guid)
}

// Skip is a paid mutator transaction binding the contract method 0x20702a5c.
//
// Solidity: function skip((uint32,bytes32,uint64) _origin, address _receiver, bytes _message, address _composer, bytes32 _guid) returns()
func (_EndpointV2 *EndpointV2Session) Skip(_origin Origin, _receiver common.Address, _message []byte, _composer common.Address, _guid [32]byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.Skip(&_EndpointV2.TransactOpts, _origin, _receiver, _message, _composer, _guid)
}

// Skip is a paid mutator transaction binding the contract method 0x20702a5c.
//
// Solidity: function skip((uint32,bytes32,uint64) _origin, address _receiver, bytes _message, address _composer, bytes32 _guid) returns()
func (_EndpointV2 *EndpointV2TransactorSession) Skip(_origin Origin, _receiver common.Address, _message []byte, _composer common.Address, _guid [32]byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.Skip(&_EndpointV2.TransactOpts, _origin, _receiver, _message, _composer, _guid)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EndpointV2 *EndpointV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EndpointV2 *EndpointV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.TransferOwnership(&_EndpointV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EndpointV2 *EndpointV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.TransferOwnership(&_EndpointV2.TransactOpts, newOwner)
}

// Verify is a paid mutator transaction binding the contract method 0xa825d747.
//
// Solidity: function verify((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) returns()
func (_EndpointV2 *EndpointV2Transactor) Verify(opts *bind.TransactOpts, _origin Origin, _receiver common.Address, _payloadHash [32]byte) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "verify", _origin, _receiver, _payloadHash)
}

// Verify is a paid mutator transaction binding the contract method 0xa825d747.
//
// Solidity: function verify((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) returns()
func (_EndpointV2 *EndpointV2Session) Verify(_origin Origin, _receiver common.Address, _payloadHash [32]byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.Verify(&_EndpointV2.TransactOpts, _origin, _receiver, _payloadHash)
}

// Verify is a paid mutator transaction binding the contract method 0xa825d747.
//
// Solidity: function verify((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) returns()
func (_EndpointV2 *EndpointV2TransactorSession) Verify(_origin Origin, _receiver common.Address, _payloadHash [32]byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.Verify(&_EndpointV2.TransactOpts, _origin, _receiver, _payloadHash)
}

// EndpointV2DefaultReceiveLibrarySetIterator is returned from FilterDefaultReceiveLibrarySet and is used to iterate over the raw logs and unpacked data for DefaultReceiveLibrarySet events raised by the EndpointV2 contract.
type EndpointV2DefaultReceiveLibrarySetIterator struct {
	Event *EndpointV2DefaultReceiveLibrarySet // Event containing the contract specifics and raw log

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
func (it *EndpointV2DefaultReceiveLibrarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2DefaultReceiveLibrarySet)
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
		it.Event = new(EndpointV2DefaultReceiveLibrarySet)
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
func (it *EndpointV2DefaultReceiveLibrarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2DefaultReceiveLibrarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2DefaultReceiveLibrarySet represents a DefaultReceiveLibrarySet event raised by the EndpointV2 contract.
type EndpointV2DefaultReceiveLibrarySet struct {
	Oapp common.Address
	Uln  common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDefaultReceiveLibrarySet is a free log retrieval operation binding the contract event 0x75d57c5f698e5d43789fc0f3dfa60b2828a82b3b47326514a8fd7675ef9850d0.
//
// Solidity: event DefaultReceiveLibrarySet(address indexed oapp, address uln)
func (_EndpointV2 *EndpointV2Filterer) FilterDefaultReceiveLibrarySet(opts *bind.FilterOpts, oapp []common.Address) (*EndpointV2DefaultReceiveLibrarySetIterator, error) {

	var oappRule []interface{}
	for _, oappItem := range oapp {
		oappRule = append(oappRule, oappItem)
	}

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "DefaultReceiveLibrarySet", oappRule)
	if err != nil {
		return nil, err
	}
	return &EndpointV2DefaultReceiveLibrarySetIterator{contract: _EndpointV2.contract, event: "DefaultReceiveLibrarySet", logs: logs, sub: sub}, nil
}

// WatchDefaultReceiveLibrarySet is a free log subscription operation binding the contract event 0x75d57c5f698e5d43789fc0f3dfa60b2828a82b3b47326514a8fd7675ef9850d0.
//
// Solidity: event DefaultReceiveLibrarySet(address indexed oapp, address uln)
func (_EndpointV2 *EndpointV2Filterer) WatchDefaultReceiveLibrarySet(opts *bind.WatchOpts, sink chan<- *EndpointV2DefaultReceiveLibrarySet, oapp []common.Address) (event.Subscription, error) {

	var oappRule []interface{}
	for _, oappItem := range oapp {
		oappRule = append(oappRule, oappItem)
	}

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "DefaultReceiveLibrarySet", oappRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2DefaultReceiveLibrarySet)
				if err := _EndpointV2.contract.UnpackLog(event, "DefaultReceiveLibrarySet", log); err != nil {
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

// ParseDefaultReceiveLibrarySet is a log parse operation binding the contract event 0x75d57c5f698e5d43789fc0f3dfa60b2828a82b3b47326514a8fd7675ef9850d0.
//
// Solidity: event DefaultReceiveLibrarySet(address indexed oapp, address uln)
func (_EndpointV2 *EndpointV2Filterer) ParseDefaultReceiveLibrarySet(log types.Log) (*EndpointV2DefaultReceiveLibrarySet, error) {
	event := new(EndpointV2DefaultReceiveLibrarySet)
	if err := _EndpointV2.contract.UnpackLog(event, "DefaultReceiveLibrarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2DefaultSendLibrarySetIterator is returned from FilterDefaultSendLibrarySet and is used to iterate over the raw logs and unpacked data for DefaultSendLibrarySet events raised by the EndpointV2 contract.
type EndpointV2DefaultSendLibrarySetIterator struct {
	Event *EndpointV2DefaultSendLibrarySet // Event containing the contract specifics and raw log

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
func (it *EndpointV2DefaultSendLibrarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2DefaultSendLibrarySet)
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
		it.Event = new(EndpointV2DefaultSendLibrarySet)
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
func (it *EndpointV2DefaultSendLibrarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2DefaultSendLibrarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2DefaultSendLibrarySet represents a DefaultSendLibrarySet event raised by the EndpointV2 contract.
type EndpointV2DefaultSendLibrarySet struct {
	Oapp    common.Address
	SendLib common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDefaultSendLibrarySet is a free log retrieval operation binding the contract event 0x6e09482bf7892d31b5b1925098416d80f0237f96d029346d3d1fd87592ebd0c7.
//
// Solidity: event DefaultSendLibrarySet(address indexed oapp, address sendLib)
func (_EndpointV2 *EndpointV2Filterer) FilterDefaultSendLibrarySet(opts *bind.FilterOpts, oapp []common.Address) (*EndpointV2DefaultSendLibrarySetIterator, error) {

	var oappRule []interface{}
	for _, oappItem := range oapp {
		oappRule = append(oappRule, oappItem)
	}

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "DefaultSendLibrarySet", oappRule)
	if err != nil {
		return nil, err
	}
	return &EndpointV2DefaultSendLibrarySetIterator{contract: _EndpointV2.contract, event: "DefaultSendLibrarySet", logs: logs, sub: sub}, nil
}

// WatchDefaultSendLibrarySet is a free log subscription operation binding the contract event 0x6e09482bf7892d31b5b1925098416d80f0237f96d029346d3d1fd87592ebd0c7.
//
// Solidity: event DefaultSendLibrarySet(address indexed oapp, address sendLib)
func (_EndpointV2 *EndpointV2Filterer) WatchDefaultSendLibrarySet(opts *bind.WatchOpts, sink chan<- *EndpointV2DefaultSendLibrarySet, oapp []common.Address) (event.Subscription, error) {

	var oappRule []interface{}
	for _, oappItem := range oapp {
		oappRule = append(oappRule, oappItem)
	}

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "DefaultSendLibrarySet", oappRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2DefaultSendLibrarySet)
				if err := _EndpointV2.contract.UnpackLog(event, "DefaultSendLibrarySet", log); err != nil {
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

// ParseDefaultSendLibrarySet is a log parse operation binding the contract event 0x6e09482bf7892d31b5b1925098416d80f0237f96d029346d3d1fd87592ebd0c7.
//
// Solidity: event DefaultSendLibrarySet(address indexed oapp, address sendLib)
func (_EndpointV2 *EndpointV2Filterer) ParseDefaultSendLibrarySet(log types.Log) (*EndpointV2DefaultSendLibrarySet, error) {
	event := new(EndpointV2DefaultSendLibrarySet)
	if err := _EndpointV2.contract.UnpackLog(event, "DefaultSendLibrarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EndpointV2 contract.
type EndpointV2InitializedIterator struct {
	Event *EndpointV2Initialized // Event containing the contract specifics and raw log

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
func (it *EndpointV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2Initialized)
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
		it.Event = new(EndpointV2Initialized)
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
func (it *EndpointV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2Initialized represents a Initialized event raised by the EndpointV2 contract.
type EndpointV2Initialized struct {
	Version uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc80eaabd798ec84cc42c71a5226e2dff95376cf50c32a53de2a2732b8d293e31.
//
// Solidity: event Initialized(uint16 version)
func (_EndpointV2 *EndpointV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*EndpointV2InitializedIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EndpointV2InitializedIterator{contract: _EndpointV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc80eaabd798ec84cc42c71a5226e2dff95376cf50c32a53de2a2732b8d293e31.
//
// Solidity: event Initialized(uint16 version)
func (_EndpointV2 *EndpointV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EndpointV2Initialized) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2Initialized)
				if err := _EndpointV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc80eaabd798ec84cc42c71a5226e2dff95376cf50c32a53de2a2732b8d293e31.
//
// Solidity: event Initialized(uint16 version)
func (_EndpointV2 *EndpointV2Filterer) ParseInitialized(log types.Log) (*EndpointV2Initialized, error) {
	event := new(EndpointV2Initialized)
	if err := _EndpointV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2LzTokenFeeCollectorUpdatedIterator is returned from FilterLzTokenFeeCollectorUpdated and is used to iterate over the raw logs and unpacked data for LzTokenFeeCollectorUpdated events raised by the EndpointV2 contract.
type EndpointV2LzTokenFeeCollectorUpdatedIterator struct {
	Event *EndpointV2LzTokenFeeCollectorUpdated // Event containing the contract specifics and raw log

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
func (it *EndpointV2LzTokenFeeCollectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2LzTokenFeeCollectorUpdated)
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
		it.Event = new(EndpointV2LzTokenFeeCollectorUpdated)
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
func (it *EndpointV2LzTokenFeeCollectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2LzTokenFeeCollectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2LzTokenFeeCollectorUpdated represents a LzTokenFeeCollectorUpdated event raised by the EndpointV2 contract.
type EndpointV2LzTokenFeeCollectorUpdated struct {
	NewFeeCollector common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLzTokenFeeCollectorUpdated is a free log retrieval operation binding the contract event 0xe2e17c3336b07514e0f8adde7f0a5c10095238d7f5982ad5c64ae9cc1d3dd084.
//
// Solidity: event LzTokenFeeCollectorUpdated(address indexed newFeeCollector)
func (_EndpointV2 *EndpointV2Filterer) FilterLzTokenFeeCollectorUpdated(opts *bind.FilterOpts, newFeeCollector []common.Address) (*EndpointV2LzTokenFeeCollectorUpdatedIterator, error) {

	var newFeeCollectorRule []interface{}
	for _, newFeeCollectorItem := range newFeeCollector {
		newFeeCollectorRule = append(newFeeCollectorRule, newFeeCollectorItem)
	}

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "LzTokenFeeCollectorUpdated", newFeeCollectorRule)
	if err != nil {
		return nil, err
	}
	return &EndpointV2LzTokenFeeCollectorUpdatedIterator{contract: _EndpointV2.contract, event: "LzTokenFeeCollectorUpdated", logs: logs, sub: sub}, nil
}

// WatchLzTokenFeeCollectorUpdated is a free log subscription operation binding the contract event 0xe2e17c3336b07514e0f8adde7f0a5c10095238d7f5982ad5c64ae9cc1d3dd084.
//
// Solidity: event LzTokenFeeCollectorUpdated(address indexed newFeeCollector)
func (_EndpointV2 *EndpointV2Filterer) WatchLzTokenFeeCollectorUpdated(opts *bind.WatchOpts, sink chan<- *EndpointV2LzTokenFeeCollectorUpdated, newFeeCollector []common.Address) (event.Subscription, error) {

	var newFeeCollectorRule []interface{}
	for _, newFeeCollectorItem := range newFeeCollector {
		newFeeCollectorRule = append(newFeeCollectorRule, newFeeCollectorItem)
	}

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "LzTokenFeeCollectorUpdated", newFeeCollectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2LzTokenFeeCollectorUpdated)
				if err := _EndpointV2.contract.UnpackLog(event, "LzTokenFeeCollectorUpdated", log); err != nil {
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

// ParseLzTokenFeeCollectorUpdated is a log parse operation binding the contract event 0xe2e17c3336b07514e0f8adde7f0a5c10095238d7f5982ad5c64ae9cc1d3dd084.
//
// Solidity: event LzTokenFeeCollectorUpdated(address indexed newFeeCollector)
func (_EndpointV2 *EndpointV2Filterer) ParseLzTokenFeeCollectorUpdated(log types.Log) (*EndpointV2LzTokenFeeCollectorUpdated, error) {
	event := new(EndpointV2LzTokenFeeCollectorUpdated)
	if err := _EndpointV2.contract.UnpackLog(event, "LzTokenFeeCollectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2LzTokenUpdatedIterator is returned from FilterLzTokenUpdated and is used to iterate over the raw logs and unpacked data for LzTokenUpdated events raised by the EndpointV2 contract.
type EndpointV2LzTokenUpdatedIterator struct {
	Event *EndpointV2LzTokenUpdated // Event containing the contract specifics and raw log

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
func (it *EndpointV2LzTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2LzTokenUpdated)
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
		it.Event = new(EndpointV2LzTokenUpdated)
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
func (it *EndpointV2LzTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2LzTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2LzTokenUpdated represents a LzTokenUpdated event raised by the EndpointV2 contract.
type EndpointV2LzTokenUpdated struct {
	NewCompositionFee common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLzTokenUpdated is a free log retrieval operation binding the contract event 0xef29788352270a67ade1b3b0ab5aa38c7ae32f1aa4a3eeda6a00184fb985f298.
//
// Solidity: event LzTokenUpdated(address indexed newCompositionFee)
func (_EndpointV2 *EndpointV2Filterer) FilterLzTokenUpdated(opts *bind.FilterOpts, newCompositionFee []common.Address) (*EndpointV2LzTokenUpdatedIterator, error) {

	var newCompositionFeeRule []interface{}
	for _, newCompositionFeeItem := range newCompositionFee {
		newCompositionFeeRule = append(newCompositionFeeRule, newCompositionFeeItem)
	}

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "LzTokenUpdated", newCompositionFeeRule)
	if err != nil {
		return nil, err
	}
	return &EndpointV2LzTokenUpdatedIterator{contract: _EndpointV2.contract, event: "LzTokenUpdated", logs: logs, sub: sub}, nil
}

// WatchLzTokenUpdated is a free log subscription operation binding the contract event 0xef29788352270a67ade1b3b0ab5aa38c7ae32f1aa4a3eeda6a00184fb985f298.
//
// Solidity: event LzTokenUpdated(address indexed newCompositionFee)
func (_EndpointV2 *EndpointV2Filterer) WatchLzTokenUpdated(opts *bind.WatchOpts, sink chan<- *EndpointV2LzTokenUpdated, newCompositionFee []common.Address) (event.Subscription, error) {

	var newCompositionFeeRule []interface{}
	for _, newCompositionFeeItem := range newCompositionFee {
		newCompositionFeeRule = append(newCompositionFeeRule, newCompositionFeeItem)
	}

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "LzTokenUpdated", newCompositionFeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2LzTokenUpdated)
				if err := _EndpointV2.contract.UnpackLog(event, "LzTokenUpdated", log); err != nil {
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

// ParseLzTokenUpdated is a log parse operation binding the contract event 0xef29788352270a67ade1b3b0ab5aa38c7ae32f1aa4a3eeda6a00184fb985f298.
//
// Solidity: event LzTokenUpdated(address indexed newCompositionFee)
func (_EndpointV2 *EndpointV2Filterer) ParseLzTokenUpdated(log types.Log) (*EndpointV2LzTokenUpdated, error) {
	event := new(EndpointV2LzTokenUpdated)
	if err := _EndpointV2.contract.UnpackLog(event, "LzTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2NativeFeePaidIterator is returned from FilterNativeFeePaid and is used to iterate over the raw logs and unpacked data for NativeFeePaid events raised by the EndpointV2 contract.
type EndpointV2NativeFeePaidIterator struct {
	Event *EndpointV2NativeFeePaid // Event containing the contract specifics and raw log

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
func (it *EndpointV2NativeFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2NativeFeePaid)
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
		it.Event = new(EndpointV2NativeFeePaid)
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
func (it *EndpointV2NativeFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2NativeFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2NativeFeePaid represents a NativeFeePaid event raised by the EndpointV2 contract.
type EndpointV2NativeFeePaid struct {
	Uln    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNativeFeePaid is a free log retrieval operation binding the contract event 0xfae01433a7b6c07798414842e23859da937c7f1a5bbafb464c123b07760b5e45.
//
// Solidity: event NativeFeePaid(address indexed uln, uint256 amount)
func (_EndpointV2 *EndpointV2Filterer) FilterNativeFeePaid(opts *bind.FilterOpts, uln []common.Address) (*EndpointV2NativeFeePaidIterator, error) {

	var ulnRule []interface{}
	for _, ulnItem := range uln {
		ulnRule = append(ulnRule, ulnItem)
	}

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "NativeFeePaid", ulnRule)
	if err != nil {
		return nil, err
	}
	return &EndpointV2NativeFeePaidIterator{contract: _EndpointV2.contract, event: "NativeFeePaid", logs: logs, sub: sub}, nil
}

// WatchNativeFeePaid is a free log subscription operation binding the contract event 0xfae01433a7b6c07798414842e23859da937c7f1a5bbafb464c123b07760b5e45.
//
// Solidity: event NativeFeePaid(address indexed uln, uint256 amount)
func (_EndpointV2 *EndpointV2Filterer) WatchNativeFeePaid(opts *bind.WatchOpts, sink chan<- *EndpointV2NativeFeePaid, uln []common.Address) (event.Subscription, error) {

	var ulnRule []interface{}
	for _, ulnItem := range uln {
		ulnRule = append(ulnRule, ulnItem)
	}

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "NativeFeePaid", ulnRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2NativeFeePaid)
				if err := _EndpointV2.contract.UnpackLog(event, "NativeFeePaid", log); err != nil {
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

// ParseNativeFeePaid is a log parse operation binding the contract event 0xfae01433a7b6c07798414842e23859da937c7f1a5bbafb464c123b07760b5e45.
//
// Solidity: event NativeFeePaid(address indexed uln, uint256 amount)
func (_EndpointV2 *EndpointV2Filterer) ParseNativeFeePaid(log types.Log) (*EndpointV2NativeFeePaid, error) {
	event := new(EndpointV2NativeFeePaid)
	if err := _EndpointV2.contract.UnpackLog(event, "NativeFeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2NewRelayerIterator is returned from FilterNewRelayer and is used to iterate over the raw logs and unpacked data for NewRelayer events raised by the EndpointV2 contract.
type EndpointV2NewRelayerIterator struct {
	Event *EndpointV2NewRelayer // Event containing the contract specifics and raw log

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
func (it *EndpointV2NewRelayerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2NewRelayer)
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
		it.Event = new(EndpointV2NewRelayer)
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
func (it *EndpointV2NewRelayerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2NewRelayerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2NewRelayer represents a NewRelayer event raised by the EndpointV2 contract.
type EndpointV2NewRelayer struct {
	NewRelayer common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewRelayer is a free log retrieval operation binding the contract event 0x580142b724d5ec9ac79a79c1a74837f611bfc6a32d918bf700d0b8f8bff90d5b.
//
// Solidity: event NewRelayer(address newRelayer)
func (_EndpointV2 *EndpointV2Filterer) FilterNewRelayer(opts *bind.FilterOpts) (*EndpointV2NewRelayerIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "NewRelayer")
	if err != nil {
		return nil, err
	}
	return &EndpointV2NewRelayerIterator{contract: _EndpointV2.contract, event: "NewRelayer", logs: logs, sub: sub}, nil
}

// WatchNewRelayer is a free log subscription operation binding the contract event 0x580142b724d5ec9ac79a79c1a74837f611bfc6a32d918bf700d0b8f8bff90d5b.
//
// Solidity: event NewRelayer(address newRelayer)
func (_EndpointV2 *EndpointV2Filterer) WatchNewRelayer(opts *bind.WatchOpts, sink chan<- *EndpointV2NewRelayer) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "NewRelayer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2NewRelayer)
				if err := _EndpointV2.contract.UnpackLog(event, "NewRelayer", log); err != nil {
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

// ParseNewRelayer is a log parse operation binding the contract event 0x580142b724d5ec9ac79a79c1a74837f611bfc6a32d918bf700d0b8f8bff90d5b.
//
// Solidity: event NewRelayer(address newRelayer)
func (_EndpointV2 *EndpointV2Filterer) ParseNewRelayer(log types.Log) (*EndpointV2NewRelayer, error) {
	event := new(EndpointV2NewRelayer)
	if err := _EndpointV2.contract.UnpackLog(event, "NewRelayer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EndpointV2 contract.
type EndpointV2OwnershipTransferredIterator struct {
	Event *EndpointV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EndpointV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2OwnershipTransferred)
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
		it.Event = new(EndpointV2OwnershipTransferred)
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
func (it *EndpointV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2OwnershipTransferred represents a OwnershipTransferred event raised by the EndpointV2 contract.
type EndpointV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EndpointV2 *EndpointV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EndpointV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EndpointV2OwnershipTransferredIterator{contract: _EndpointV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EndpointV2 *EndpointV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EndpointV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2OwnershipTransferred)
				if err := _EndpointV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EndpointV2 *EndpointV2Filterer) ParseOwnershipTransferred(log types.Log) (*EndpointV2OwnershipTransferred, error) {
	event := new(EndpointV2OwnershipTransferred)
	if err := _EndpointV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2PacketBurntIterator is returned from FilterPacketBurnt and is used to iterate over the raw logs and unpacked data for PacketBurnt events raised by the EndpointV2 contract.
type EndpointV2PacketBurntIterator struct {
	Event *EndpointV2PacketBurnt // Event containing the contract specifics and raw log

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
func (it *EndpointV2PacketBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2PacketBurnt)
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
		it.Event = new(EndpointV2PacketBurnt)
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
func (it *EndpointV2PacketBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2PacketBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2PacketBurnt represents a PacketBurnt event raised by the EndpointV2 contract.
type EndpointV2PacketBurnt struct {
	Guid    [32]byte
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPacketBurnt is a free log retrieval operation binding the contract event 0xe5b79e71d459391825ef6937987b7c0b32d8fac7d7aa80700d5f5c16fdbe4e33.
//
// Solidity: event PacketBurnt(bytes32 guid, bool success)
func (_EndpointV2 *EndpointV2Filterer) FilterPacketBurnt(opts *bind.FilterOpts) (*EndpointV2PacketBurntIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "PacketBurnt")
	if err != nil {
		return nil, err
	}
	return &EndpointV2PacketBurntIterator{contract: _EndpointV2.contract, event: "PacketBurnt", logs: logs, sub: sub}, nil
}

// WatchPacketBurnt is a free log subscription operation binding the contract event 0xe5b79e71d459391825ef6937987b7c0b32d8fac7d7aa80700d5f5c16fdbe4e33.
//
// Solidity: event PacketBurnt(bytes32 guid, bool success)
func (_EndpointV2 *EndpointV2Filterer) WatchPacketBurnt(opts *bind.WatchOpts, sink chan<- *EndpointV2PacketBurnt) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "PacketBurnt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2PacketBurnt)
				if err := _EndpointV2.contract.UnpackLog(event, "PacketBurnt", log); err != nil {
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

// ParsePacketBurnt is a log parse operation binding the contract event 0xe5b79e71d459391825ef6937987b7c0b32d8fac7d7aa80700d5f5c16fdbe4e33.
//
// Solidity: event PacketBurnt(bytes32 guid, bool success)
func (_EndpointV2 *EndpointV2Filterer) ParsePacketBurnt(log types.Log) (*EndpointV2PacketBurnt, error) {
	event := new(EndpointV2PacketBurnt)
	if err := _EndpointV2.contract.UnpackLog(event, "PacketBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2PacketDeliveredIterator is returned from FilterPacketDelivered and is used to iterate over the raw logs and unpacked data for PacketDelivered events raised by the EndpointV2 contract.
type EndpointV2PacketDeliveredIterator struct {
	Event *EndpointV2PacketDelivered // Event containing the contract specifics and raw log

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
func (it *EndpointV2PacketDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2PacketDelivered)
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
		it.Event = new(EndpointV2PacketDelivered)
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
func (it *EndpointV2PacketDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2PacketDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2PacketDelivered represents a PacketDelivered event raised by the EndpointV2 contract.
type EndpointV2PacketDelivered struct {
	Guid [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPacketDelivered is a free log retrieval operation binding the contract event 0x1d88581d0a3d78a17f0a642defe5d0612655d98f4427c3d8ea86fbbf4e8bfb05.
//
// Solidity: event PacketDelivered(bytes32 guid)
func (_EndpointV2 *EndpointV2Filterer) FilterPacketDelivered(opts *bind.FilterOpts) (*EndpointV2PacketDeliveredIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "PacketDelivered")
	if err != nil {
		return nil, err
	}
	return &EndpointV2PacketDeliveredIterator{contract: _EndpointV2.contract, event: "PacketDelivered", logs: logs, sub: sub}, nil
}

// WatchPacketDelivered is a free log subscription operation binding the contract event 0x1d88581d0a3d78a17f0a642defe5d0612655d98f4427c3d8ea86fbbf4e8bfb05.
//
// Solidity: event PacketDelivered(bytes32 guid)
func (_EndpointV2 *EndpointV2Filterer) WatchPacketDelivered(opts *bind.WatchOpts, sink chan<- *EndpointV2PacketDelivered) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "PacketDelivered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2PacketDelivered)
				if err := _EndpointV2.contract.UnpackLog(event, "PacketDelivered", log); err != nil {
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

// ParsePacketDelivered is a log parse operation binding the contract event 0x1d88581d0a3d78a17f0a642defe5d0612655d98f4427c3d8ea86fbbf4e8bfb05.
//
// Solidity: event PacketDelivered(bytes32 guid)
func (_EndpointV2 *EndpointV2Filterer) ParsePacketDelivered(log types.Log) (*EndpointV2PacketDelivered, error) {
	event := new(EndpointV2PacketDelivered)
	if err := _EndpointV2.contract.UnpackLog(event, "PacketDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2PacketFeeTransferredIterator is returned from FilterPacketFeeTransferred and is used to iterate over the raw logs and unpacked data for PacketFeeTransferred events raised by the EndpointV2 contract.
type EndpointV2PacketFeeTransferredIterator struct {
	Event *EndpointV2PacketFeeTransferred // Event containing the contract specifics and raw log

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
func (it *EndpointV2PacketFeeTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2PacketFeeTransferred)
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
		it.Event = new(EndpointV2PacketFeeTransferred)
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
func (it *EndpointV2PacketFeeTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2PacketFeeTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2PacketFeeTransferred represents a PacketFeeTransferred event raised by the EndpointV2 contract.
type EndpointV2PacketFeeTransferred struct {
	SendLib   common.Address
	Guid      [32]byte
	NativeFee *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPacketFeeTransferred is a free log retrieval operation binding the contract event 0x036bdd083028fc45c48b9ea92828a44f8d3739856bac464b3f594822b81c6286.
//
// Solidity: event PacketFeeTransferred(address indexed sendLib, bytes32 indexed guid, uint256 nativeFee)
func (_EndpointV2 *EndpointV2Filterer) FilterPacketFeeTransferred(opts *bind.FilterOpts, sendLib []common.Address, guid [][32]byte) (*EndpointV2PacketFeeTransferredIterator, error) {

	var sendLibRule []interface{}
	for _, sendLibItem := range sendLib {
		sendLibRule = append(sendLibRule, sendLibItem)
	}
	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "PacketFeeTransferred", sendLibRule, guidRule)
	if err != nil {
		return nil, err
	}
	return &EndpointV2PacketFeeTransferredIterator{contract: _EndpointV2.contract, event: "PacketFeeTransferred", logs: logs, sub: sub}, nil
}

// WatchPacketFeeTransferred is a free log subscription operation binding the contract event 0x036bdd083028fc45c48b9ea92828a44f8d3739856bac464b3f594822b81c6286.
//
// Solidity: event PacketFeeTransferred(address indexed sendLib, bytes32 indexed guid, uint256 nativeFee)
func (_EndpointV2 *EndpointV2Filterer) WatchPacketFeeTransferred(opts *bind.WatchOpts, sink chan<- *EndpointV2PacketFeeTransferred, sendLib []common.Address, guid [][32]byte) (event.Subscription, error) {

	var sendLibRule []interface{}
	for _, sendLibItem := range sendLib {
		sendLibRule = append(sendLibRule, sendLibItem)
	}
	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "PacketFeeTransferred", sendLibRule, guidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2PacketFeeTransferred)
				if err := _EndpointV2.contract.UnpackLog(event, "PacketFeeTransferred", log); err != nil {
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

// ParsePacketFeeTransferred is a log parse operation binding the contract event 0x036bdd083028fc45c48b9ea92828a44f8d3739856bac464b3f594822b81c6286.
//
// Solidity: event PacketFeeTransferred(address indexed sendLib, bytes32 indexed guid, uint256 nativeFee)
func (_EndpointV2 *EndpointV2Filterer) ParsePacketFeeTransferred(log types.Log) (*EndpointV2PacketFeeTransferred, error) {
	event := new(EndpointV2PacketFeeTransferred)
	if err := _EndpointV2.contract.UnpackLog(event, "PacketFeeTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2PacketSentIterator is returned from FilterPacketSent and is used to iterate over the raw logs and unpacked data for PacketSent events raised by the EndpointV2 contract.
type EndpointV2PacketSentIterator struct {
	Event *EndpointV2PacketSent // Event containing the contract specifics and raw log

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
func (it *EndpointV2PacketSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2PacketSent)
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
		it.Event = new(EndpointV2PacketSent)
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
func (it *EndpointV2PacketSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2PacketSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2PacketSent represents a PacketSent event raised by the EndpointV2 contract.
type EndpointV2PacketSent struct {
	PacketHeader []byte
	Payload      []byte
	Options      []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPacketSent is a free log retrieval operation binding the contract event 0x46641819ee54cd5f17e30766afaa5c13418264067e9dec5ea84557ec978977b2.
//
// Solidity: event PacketSent(bytes packetHeader, bytes payload, bytes options)
func (_EndpointV2 *EndpointV2Filterer) FilterPacketSent(opts *bind.FilterOpts) (*EndpointV2PacketSentIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "PacketSent")
	if err != nil {
		return nil, err
	}
	return &EndpointV2PacketSentIterator{contract: _EndpointV2.contract, event: "PacketSent", logs: logs, sub: sub}, nil
}

// WatchPacketSent is a free log subscription operation binding the contract event 0x46641819ee54cd5f17e30766afaa5c13418264067e9dec5ea84557ec978977b2.
//
// Solidity: event PacketSent(bytes packetHeader, bytes payload, bytes options)
func (_EndpointV2 *EndpointV2Filterer) WatchPacketSent(opts *bind.WatchOpts, sink chan<- *EndpointV2PacketSent) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "PacketSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2PacketSent)
				if err := _EndpointV2.contract.UnpackLog(event, "PacketSent", log); err != nil {
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

// ParsePacketSent is a log parse operation binding the contract event 0x46641819ee54cd5f17e30766afaa5c13418264067e9dec5ea84557ec978977b2.
//
// Solidity: event PacketSent(bytes packetHeader, bytes payload, bytes options)
func (_EndpointV2 *EndpointV2Filterer) ParsePacketSent(log types.Log) (*EndpointV2PacketSent, error) {
	event := new(EndpointV2PacketSent)
	if err := _EndpointV2.contract.UnpackLog(event, "PacketSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2PacketVerifiedIterator is returned from FilterPacketVerified and is used to iterate over the raw logs and unpacked data for PacketVerified events raised by the EndpointV2 contract.
type EndpointV2PacketVerifiedIterator struct {
	Event *EndpointV2PacketVerified // Event containing the contract specifics and raw log

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
func (it *EndpointV2PacketVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2PacketVerified)
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
		it.Event = new(EndpointV2PacketVerified)
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
func (it *EndpointV2PacketVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2PacketVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2PacketVerified represents a PacketVerified event raised by the EndpointV2 contract.
type EndpointV2PacketVerified struct {
	Origin      Origin
	Receiver    common.Address
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPacketVerified is a free log retrieval operation binding the contract event 0x0d87345f3d1c929caba93e1c3821b54ff3512e12b66aa3cfe54b6bcbc17e59b4.
//
// Solidity: event PacketVerified((uint32,bytes32,uint64) origin, address receiver, bytes32 payloadHash)
func (_EndpointV2 *EndpointV2Filterer) FilterPacketVerified(opts *bind.FilterOpts) (*EndpointV2PacketVerifiedIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "PacketVerified")
	if err != nil {
		return nil, err
	}
	return &EndpointV2PacketVerifiedIterator{contract: _EndpointV2.contract, event: "PacketVerified", logs: logs, sub: sub}, nil
}

// WatchPacketVerified is a free log subscription operation binding the contract event 0x0d87345f3d1c929caba93e1c3821b54ff3512e12b66aa3cfe54b6bcbc17e59b4.
//
// Solidity: event PacketVerified((uint32,bytes32,uint64) origin, address receiver, bytes32 payloadHash)
func (_EndpointV2 *EndpointV2Filterer) WatchPacketVerified(opts *bind.WatchOpts, sink chan<- *EndpointV2PacketVerified) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "PacketVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2PacketVerified)
				if err := _EndpointV2.contract.UnpackLog(event, "PacketVerified", log); err != nil {
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

// ParsePacketVerified is a log parse operation binding the contract event 0x0d87345f3d1c929caba93e1c3821b54ff3512e12b66aa3cfe54b6bcbc17e59b4.
//
// Solidity: event PacketVerified((uint32,bytes32,uint64) origin, address receiver, bytes32 payloadHash)
func (_EndpointV2 *EndpointV2Filterer) ParsePacketVerified(log types.Log) (*EndpointV2PacketVerified, error) {
	event := new(EndpointV2PacketVerified)
	if err := _EndpointV2.contract.UnpackLog(event, "PacketVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2ReceiveLibrarySetIterator is returned from FilterReceiveLibrarySet and is used to iterate over the raw logs and unpacked data for ReceiveLibrarySet events raised by the EndpointV2 contract.
type EndpointV2ReceiveLibrarySetIterator struct {
	Event *EndpointV2ReceiveLibrarySet // Event containing the contract specifics and raw log

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
func (it *EndpointV2ReceiveLibrarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2ReceiveLibrarySet)
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
		it.Event = new(EndpointV2ReceiveLibrarySet)
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
func (it *EndpointV2ReceiveLibrarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2ReceiveLibrarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2ReceiveLibrarySet represents a ReceiveLibrarySet event raised by the EndpointV2 contract.
type EndpointV2ReceiveLibrarySet struct {
	Oapp          common.Address
	RemoteEid     uint32
	OldReceiveLib common.Address
	NewReceiveLib common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceiveLibrarySet is a free log retrieval operation binding the contract event 0xc0833c35bb1d0beadca36bed54c7098819e109542a6d233d33c7c2039c8ec9aa.
//
// Solidity: event ReceiveLibrarySet(address indexed oapp, uint32 indexed remoteEid, address oldReceiveLib, address newReceiveLib)
func (_EndpointV2 *EndpointV2Filterer) FilterReceiveLibrarySet(opts *bind.FilterOpts, oapp []common.Address, remoteEid []uint32) (*EndpointV2ReceiveLibrarySetIterator, error) {

	var oappRule []interface{}
	for _, oappItem := range oapp {
		oappRule = append(oappRule, oappItem)
	}
	var remoteEidRule []interface{}
	for _, remoteEidItem := range remoteEid {
		remoteEidRule = append(remoteEidRule, remoteEidItem)
	}

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "ReceiveLibrarySet", oappRule, remoteEidRule)
	if err != nil {
		return nil, err
	}
	return &EndpointV2ReceiveLibrarySetIterator{contract: _EndpointV2.contract, event: "ReceiveLibrarySet", logs: logs, sub: sub}, nil
}

// WatchReceiveLibrarySet is a free log subscription operation binding the contract event 0xc0833c35bb1d0beadca36bed54c7098819e109542a6d233d33c7c2039c8ec9aa.
//
// Solidity: event ReceiveLibrarySet(address indexed oapp, uint32 indexed remoteEid, address oldReceiveLib, address newReceiveLib)
func (_EndpointV2 *EndpointV2Filterer) WatchReceiveLibrarySet(opts *bind.WatchOpts, sink chan<- *EndpointV2ReceiveLibrarySet, oapp []common.Address, remoteEid []uint32) (event.Subscription, error) {

	var oappRule []interface{}
	for _, oappItem := range oapp {
		oappRule = append(oappRule, oappItem)
	}
	var remoteEidRule []interface{}
	for _, remoteEidItem := range remoteEid {
		remoteEidRule = append(remoteEidRule, remoteEidItem)
	}

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "ReceiveLibrarySet", oappRule, remoteEidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2ReceiveLibrarySet)
				if err := _EndpointV2.contract.UnpackLog(event, "ReceiveLibrarySet", log); err != nil {
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

// ParseReceiveLibrarySet is a log parse operation binding the contract event 0xc0833c35bb1d0beadca36bed54c7098819e109542a6d233d33c7c2039c8ec9aa.
//
// Solidity: event ReceiveLibrarySet(address indexed oapp, uint32 indexed remoteEid, address oldReceiveLib, address newReceiveLib)
func (_EndpointV2 *EndpointV2Filterer) ParseReceiveLibrarySet(log types.Log) (*EndpointV2ReceiveLibrarySet, error) {
	event := new(EndpointV2ReceiveLibrarySet)
	if err := _EndpointV2.contract.UnpackLog(event, "ReceiveLibrarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2ReceiveLibraryTimeoutSetIterator is returned from FilterReceiveLibraryTimeoutSet and is used to iterate over the raw logs and unpacked data for ReceiveLibraryTimeoutSet events raised by the EndpointV2 contract.
type EndpointV2ReceiveLibraryTimeoutSetIterator struct {
	Event *EndpointV2ReceiveLibraryTimeoutSet // Event containing the contract specifics and raw log

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
func (it *EndpointV2ReceiveLibraryTimeoutSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2ReceiveLibraryTimeoutSet)
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
		it.Event = new(EndpointV2ReceiveLibraryTimeoutSet)
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
func (it *EndpointV2ReceiveLibraryTimeoutSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2ReceiveLibraryTimeoutSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2ReceiveLibraryTimeoutSet represents a ReceiveLibraryTimeoutSet event raised by the EndpointV2 contract.
type EndpointV2ReceiveLibraryTimeoutSet struct {
	Oapp      common.Address
	RemoteEid uint32
	Timeout   uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReceiveLibraryTimeoutSet is a free log retrieval operation binding the contract event 0xf18663bc25fb7dbe12713e0b6098ff4cac24a0c1e1f47d5028885c4537688a33.
//
// Solidity: event ReceiveLibraryTimeoutSet(address indexed oapp, uint32 indexed remoteEid, uint64 timeout)
func (_EndpointV2 *EndpointV2Filterer) FilterReceiveLibraryTimeoutSet(opts *bind.FilterOpts, oapp []common.Address, remoteEid []uint32) (*EndpointV2ReceiveLibraryTimeoutSetIterator, error) {

	var oappRule []interface{}
	for _, oappItem := range oapp {
		oappRule = append(oappRule, oappItem)
	}
	var remoteEidRule []interface{}
	for _, remoteEidItem := range remoteEid {
		remoteEidRule = append(remoteEidRule, remoteEidItem)
	}

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "ReceiveLibraryTimeoutSet", oappRule, remoteEidRule)
	if err != nil {
		return nil, err
	}
	return &EndpointV2ReceiveLibraryTimeoutSetIterator{contract: _EndpointV2.contract, event: "ReceiveLibraryTimeoutSet", logs: logs, sub: sub}, nil
}

// WatchReceiveLibraryTimeoutSet is a free log subscription operation binding the contract event 0xf18663bc25fb7dbe12713e0b6098ff4cac24a0c1e1f47d5028885c4537688a33.
//
// Solidity: event ReceiveLibraryTimeoutSet(address indexed oapp, uint32 indexed remoteEid, uint64 timeout)
func (_EndpointV2 *EndpointV2Filterer) WatchReceiveLibraryTimeoutSet(opts *bind.WatchOpts, sink chan<- *EndpointV2ReceiveLibraryTimeoutSet, oapp []common.Address, remoteEid []uint32) (event.Subscription, error) {

	var oappRule []interface{}
	for _, oappItem := range oapp {
		oappRule = append(oappRule, oappItem)
	}
	var remoteEidRule []interface{}
	for _, remoteEidItem := range remoteEid {
		remoteEidRule = append(remoteEidRule, remoteEidItem)
	}

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "ReceiveLibraryTimeoutSet", oappRule, remoteEidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2ReceiveLibraryTimeoutSet)
				if err := _EndpointV2.contract.UnpackLog(event, "ReceiveLibraryTimeoutSet", log); err != nil {
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

// ParseReceiveLibraryTimeoutSet is a log parse operation binding the contract event 0xf18663bc25fb7dbe12713e0b6098ff4cac24a0c1e1f47d5028885c4537688a33.
//
// Solidity: event ReceiveLibraryTimeoutSet(address indexed oapp, uint32 indexed remoteEid, uint64 timeout)
func (_EndpointV2 *EndpointV2Filterer) ParseReceiveLibraryTimeoutSet(log types.Log) (*EndpointV2ReceiveLibraryTimeoutSet, error) {
	event := new(EndpointV2ReceiveLibraryTimeoutSet)
	if err := _EndpointV2.contract.UnpackLog(event, "ReceiveLibraryTimeoutSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2SendLibrarySetIterator is returned from FilterSendLibrarySet and is used to iterate over the raw logs and unpacked data for SendLibrarySet events raised by the EndpointV2 contract.
type EndpointV2SendLibrarySetIterator struct {
	Event *EndpointV2SendLibrarySet // Event containing the contract specifics and raw log

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
func (it *EndpointV2SendLibrarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2SendLibrarySet)
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
		it.Event = new(EndpointV2SendLibrarySet)
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
func (it *EndpointV2SendLibrarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2SendLibrarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2SendLibrarySet represents a SendLibrarySet event raised by the EndpointV2 contract.
type EndpointV2SendLibrarySet struct {
	Oapp       common.Address
	RemoteEid  uint32
	OldSendLib common.Address
	NewSendLib common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSendLibrarySet is a free log retrieval operation binding the contract event 0xa6aa125db2867f639108fa58cae0d45cd4df33fc3bf43501973817b163f5a985.
//
// Solidity: event SendLibrarySet(address indexed oapp, uint32 indexed remoteEid, address oldSendLib, address newSendLib)
func (_EndpointV2 *EndpointV2Filterer) FilterSendLibrarySet(opts *bind.FilterOpts, oapp []common.Address, remoteEid []uint32) (*EndpointV2SendLibrarySetIterator, error) {

	var oappRule []interface{}
	for _, oappItem := range oapp {
		oappRule = append(oappRule, oappItem)
	}
	var remoteEidRule []interface{}
	for _, remoteEidItem := range remoteEid {
		remoteEidRule = append(remoteEidRule, remoteEidItem)
	}

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "SendLibrarySet", oappRule, remoteEidRule)
	if err != nil {
		return nil, err
	}
	return &EndpointV2SendLibrarySetIterator{contract: _EndpointV2.contract, event: "SendLibrarySet", logs: logs, sub: sub}, nil
}

// WatchSendLibrarySet is a free log subscription operation binding the contract event 0xa6aa125db2867f639108fa58cae0d45cd4df33fc3bf43501973817b163f5a985.
//
// Solidity: event SendLibrarySet(address indexed oapp, uint32 indexed remoteEid, address oldSendLib, address newSendLib)
func (_EndpointV2 *EndpointV2Filterer) WatchSendLibrarySet(opts *bind.WatchOpts, sink chan<- *EndpointV2SendLibrarySet, oapp []common.Address, remoteEid []uint32) (event.Subscription, error) {

	var oappRule []interface{}
	for _, oappItem := range oapp {
		oappRule = append(oappRule, oappItem)
	}
	var remoteEidRule []interface{}
	for _, remoteEidItem := range remoteEid {
		remoteEidRule = append(remoteEidRule, remoteEidItem)
	}

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "SendLibrarySet", oappRule, remoteEidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2SendLibrarySet)
				if err := _EndpointV2.contract.UnpackLog(event, "SendLibrarySet", log); err != nil {
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

// ParseSendLibrarySet is a log parse operation binding the contract event 0xa6aa125db2867f639108fa58cae0d45cd4df33fc3bf43501973817b163f5a985.
//
// Solidity: event SendLibrarySet(address indexed oapp, uint32 indexed remoteEid, address oldSendLib, address newSendLib)
func (_EndpointV2 *EndpointV2Filterer) ParseSendLibrarySet(log types.Log) (*EndpointV2SendLibrarySet, error) {
	event := new(EndpointV2SendLibrarySet)
	if err := _EndpointV2.contract.UnpackLog(event, "SendLibrarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2TokenFeeRefundedIterator is returned from FilterTokenFeeRefunded and is used to iterate over the raw logs and unpacked data for TokenFeeRefunded events raised by the EndpointV2 contract.
type EndpointV2TokenFeeRefundedIterator struct {
	Event *EndpointV2TokenFeeRefunded // Event containing the contract specifics and raw log

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
func (it *EndpointV2TokenFeeRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2TokenFeeRefunded)
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
		it.Event = new(EndpointV2TokenFeeRefunded)
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
func (it *EndpointV2TokenFeeRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2TokenFeeRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2TokenFeeRefunded represents a TokenFeeRefunded event raised by the EndpointV2 contract.
type EndpointV2TokenFeeRefunded struct {
	Guid          [32]byte
	RefundAddress common.Address
	NativeAmount  *big.Int
	Token         common.Address
	TokenAmount   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenFeeRefunded is a free log retrieval operation binding the contract event 0xd2c64a832daedd397cc212ef65075784a4878c38fed3eea32ebd33a63805f9b0.
//
// Solidity: event TokenFeeRefunded(bytes32 guid, address refundAddress, uint256 nativeAmount, address token, uint256 tokenAmount)
func (_EndpointV2 *EndpointV2Filterer) FilterTokenFeeRefunded(opts *bind.FilterOpts) (*EndpointV2TokenFeeRefundedIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "TokenFeeRefunded")
	if err != nil {
		return nil, err
	}
	return &EndpointV2TokenFeeRefundedIterator{contract: _EndpointV2.contract, event: "TokenFeeRefunded", logs: logs, sub: sub}, nil
}

// WatchTokenFeeRefunded is a free log subscription operation binding the contract event 0xd2c64a832daedd397cc212ef65075784a4878c38fed3eea32ebd33a63805f9b0.
//
// Solidity: event TokenFeeRefunded(bytes32 guid, address refundAddress, uint256 nativeAmount, address token, uint256 tokenAmount)
func (_EndpointV2 *EndpointV2Filterer) WatchTokenFeeRefunded(opts *bind.WatchOpts, sink chan<- *EndpointV2TokenFeeRefunded) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "TokenFeeRefunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2TokenFeeRefunded)
				if err := _EndpointV2.contract.UnpackLog(event, "TokenFeeRefunded", log); err != nil {
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

// ParseTokenFeeRefunded is a log parse operation binding the contract event 0xd2c64a832daedd397cc212ef65075784a4878c38fed3eea32ebd33a63805f9b0.
//
// Solidity: event TokenFeeRefunded(bytes32 guid, address refundAddress, uint256 nativeAmount, address token, uint256 tokenAmount)
func (_EndpointV2 *EndpointV2Filterer) ParseTokenFeeRefunded(log types.Log) (*EndpointV2TokenFeeRefunded, error) {
	event := new(EndpointV2TokenFeeRefunded)
	if err := _EndpointV2.contract.UnpackLog(event, "TokenFeeRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2TokenFeeWithdrawnIterator is returned from FilterTokenFeeWithdrawn and is used to iterate over the raw logs and unpacked data for TokenFeeWithdrawn events raised by the EndpointV2 contract.
type EndpointV2TokenFeeWithdrawnIterator struct {
	Event *EndpointV2TokenFeeWithdrawn // Event containing the contract specifics and raw log

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
func (it *EndpointV2TokenFeeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2TokenFeeWithdrawn)
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
		it.Event = new(EndpointV2TokenFeeWithdrawn)
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
func (it *EndpointV2TokenFeeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2TokenFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2TokenFeeWithdrawn represents a TokenFeeWithdrawn event raised by the EndpointV2 contract.
type EndpointV2TokenFeeWithdrawn struct {
	Uln    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenFeeWithdrawn is a free log retrieval operation binding the contract event 0x8918ef8c3162ebbc94a323bf1b3b40a5ab34b6abc68c2258a4593e3b027299b0.
//
// Solidity: event TokenFeeWithdrawn(address indexed uln, uint256 amount)
func (_EndpointV2 *EndpointV2Filterer) FilterTokenFeeWithdrawn(opts *bind.FilterOpts, uln []common.Address) (*EndpointV2TokenFeeWithdrawnIterator, error) {

	var ulnRule []interface{}
	for _, ulnItem := range uln {
		ulnRule = append(ulnRule, ulnItem)
	}

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "TokenFeeWithdrawn", ulnRule)
	if err != nil {
		return nil, err
	}
	return &EndpointV2TokenFeeWithdrawnIterator{contract: _EndpointV2.contract, event: "TokenFeeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenFeeWithdrawn is a free log subscription operation binding the contract event 0x8918ef8c3162ebbc94a323bf1b3b40a5ab34b6abc68c2258a4593e3b027299b0.
//
// Solidity: event TokenFeeWithdrawn(address indexed uln, uint256 amount)
func (_EndpointV2 *EndpointV2Filterer) WatchTokenFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *EndpointV2TokenFeeWithdrawn, uln []common.Address) (event.Subscription, error) {

	var ulnRule []interface{}
	for _, ulnItem := range uln {
		ulnRule = append(ulnRule, ulnItem)
	}

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "TokenFeeWithdrawn", ulnRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2TokenFeeWithdrawn)
				if err := _EndpointV2.contract.UnpackLog(event, "TokenFeeWithdrawn", log); err != nil {
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

// ParseTokenFeeWithdrawn is a log parse operation binding the contract event 0x8918ef8c3162ebbc94a323bf1b3b40a5ab34b6abc68c2258a4593e3b027299b0.
//
// Solidity: event TokenFeeWithdrawn(address indexed uln, uint256 amount)
func (_EndpointV2 *EndpointV2Filterer) ParseTokenFeeWithdrawn(log types.Log) (*EndpointV2TokenFeeWithdrawn, error) {
	event := new(EndpointV2TokenFeeWithdrawn)
	if err := _EndpointV2.contract.UnpackLog(event, "TokenFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2TreasuryUpdatedIterator is returned from FilterTreasuryUpdated and is used to iterate over the raw logs and unpacked data for TreasuryUpdated events raised by the EndpointV2 contract.
type EndpointV2TreasuryUpdatedIterator struct {
	Event *EndpointV2TreasuryUpdated // Event containing the contract specifics and raw log

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
func (it *EndpointV2TreasuryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2TreasuryUpdated)
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
		it.Event = new(EndpointV2TreasuryUpdated)
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
func (it *EndpointV2TreasuryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2TreasuryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2TreasuryUpdated represents a TreasuryUpdated event raised by the EndpointV2 contract.
type EndpointV2TreasuryUpdated struct {
	NewTreasury common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTreasuryUpdated is a free log retrieval operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address newTreasury)
func (_EndpointV2 *EndpointV2Filterer) FilterTreasuryUpdated(opts *bind.FilterOpts) (*EndpointV2TreasuryUpdatedIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return &EndpointV2TreasuryUpdatedIterator{contract: _EndpointV2.contract, event: "TreasuryUpdated", logs: logs, sub: sub}, nil
}

// WatchTreasuryUpdated is a free log subscription operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address newTreasury)
func (_EndpointV2 *EndpointV2Filterer) WatchTreasuryUpdated(opts *bind.WatchOpts, sink chan<- *EndpointV2TreasuryUpdated) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2TreasuryUpdated)
				if err := _EndpointV2.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
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

// ParseTreasuryUpdated is a log parse operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address newTreasury)
func (_EndpointV2 *EndpointV2Filterer) ParseTreasuryUpdated(log types.Log) (*EndpointV2TreasuryUpdated, error) {
	event := new(EndpointV2TreasuryUpdated)
	if err := _EndpointV2.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
