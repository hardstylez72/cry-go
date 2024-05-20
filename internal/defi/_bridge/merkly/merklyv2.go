// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package merkly

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

// Origin is an auto generated low-level Go binding around an user-defined struct.
type Origin struct {
	SrcEid uint32
	Sender [32]byte
	Nonce  uint64
}

// Merklyrefuelv2MetaData contains all meta data concerning the Merklyrefuelv2 contract.
var Merklyrefuelv2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_endpoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEndpointCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LzTokenUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"NoPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"NotEnoughNative\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"OnlyEndpoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"OnlyPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"name\":\"PeerSet\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"}],\"name\":\"allowInitializePath\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_dstEid\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"}],\"name\":\"bridgeGas\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpointV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nextNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oAppVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"senderVersion\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"receiverVersion\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"peers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_dstEid\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"_payInLzToken\",\"type\":\"bool\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"setDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"setPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// Merklyrefuelv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Merklyrefuelv2MetaData.ABI instead.
var Merklyrefuelv2ABI = Merklyrefuelv2MetaData.ABI

// Merklyrefuelv2 is an auto generated Go binding around an Ethereum contract.
type Merklyrefuelv2 struct {
	Merklyrefuelv2Caller     // Read-only binding to the contract
	Merklyrefuelv2Transactor // Write-only binding to the contract
	Merklyrefuelv2Filterer   // Log filterer for contract events
}

// Merklyrefuelv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Merklyrefuelv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Merklyrefuelv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Merklyrefuelv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Merklyrefuelv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Merklyrefuelv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Merklyrefuelv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Merklyrefuelv2Session struct {
	Contract     *Merklyrefuelv2   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Merklyrefuelv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Merklyrefuelv2CallerSession struct {
	Contract *Merklyrefuelv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Merklyrefuelv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Merklyrefuelv2TransactorSession struct {
	Contract     *Merklyrefuelv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Merklyrefuelv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Merklyrefuelv2Raw struct {
	Contract *Merklyrefuelv2 // Generic contract binding to access the raw methods on
}

// Merklyrefuelv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Merklyrefuelv2CallerRaw struct {
	Contract *Merklyrefuelv2Caller // Generic read-only contract binding to access the raw methods on
}

// Merklyrefuelv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Merklyrefuelv2TransactorRaw struct {
	Contract *Merklyrefuelv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMerklyrefuelv2 creates a new instance of Merklyrefuelv2, bound to a specific deployed contract.
func NewMerklyrefuelv2(address common.Address, backend bind.ContractBackend) (*Merklyrefuelv2, error) {
	contract, err := bindMerklyrefuelv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Merklyrefuelv2{Merklyrefuelv2Caller: Merklyrefuelv2Caller{contract: contract}, Merklyrefuelv2Transactor: Merklyrefuelv2Transactor{contract: contract}, Merklyrefuelv2Filterer: Merklyrefuelv2Filterer{contract: contract}}, nil
}

// NewMerklyrefuelv2Caller creates a new read-only instance of Merklyrefuelv2, bound to a specific deployed contract.
func NewMerklyrefuelv2Caller(address common.Address, caller bind.ContractCaller) (*Merklyrefuelv2Caller, error) {
	contract, err := bindMerklyrefuelv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Merklyrefuelv2Caller{contract: contract}, nil
}

// NewMerklyrefuelv2Transactor creates a new write-only instance of Merklyrefuelv2, bound to a specific deployed contract.
func NewMerklyrefuelv2Transactor(address common.Address, transactor bind.ContractTransactor) (*Merklyrefuelv2Transactor, error) {
	contract, err := bindMerklyrefuelv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Merklyrefuelv2Transactor{contract: contract}, nil
}

// NewMerklyrefuelv2Filterer creates a new log filterer instance of Merklyrefuelv2, bound to a specific deployed contract.
func NewMerklyrefuelv2Filterer(address common.Address, filterer bind.ContractFilterer) (*Merklyrefuelv2Filterer, error) {
	contract, err := bindMerklyrefuelv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Merklyrefuelv2Filterer{contract: contract}, nil
}

// bindMerklyrefuelv2 binds a generic wrapper to an already deployed contract.
func bindMerklyrefuelv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Merklyrefuelv2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merklyrefuelv2 *Merklyrefuelv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Merklyrefuelv2.Contract.Merklyrefuelv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merklyrefuelv2 *Merklyrefuelv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.Merklyrefuelv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merklyrefuelv2 *Merklyrefuelv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.Merklyrefuelv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merklyrefuelv2 *Merklyrefuelv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Merklyrefuelv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merklyrefuelv2 *Merklyrefuelv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merklyrefuelv2 *Merklyrefuelv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.contract.Transact(opts, method, params...)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Merklyrefuelv2 *Merklyrefuelv2Caller) AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error) {
	var out []interface{}
	err := _Merklyrefuelv2.contract.Call(opts, &out, "allowInitializePath", origin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Merklyrefuelv2 *Merklyrefuelv2Session) AllowInitializePath(origin Origin) (bool, error) {
	return _Merklyrefuelv2.Contract.AllowInitializePath(&_Merklyrefuelv2.CallOpts, origin)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Merklyrefuelv2 *Merklyrefuelv2CallerSession) AllowInitializePath(origin Origin) (bool, error) {
	return _Merklyrefuelv2.Contract.AllowInitializePath(&_Merklyrefuelv2.CallOpts, origin)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Merklyrefuelv2 *Merklyrefuelv2Caller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Merklyrefuelv2.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Merklyrefuelv2 *Merklyrefuelv2Session) Endpoint() (common.Address, error) {
	return _Merklyrefuelv2.Contract.Endpoint(&_Merklyrefuelv2.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Merklyrefuelv2 *Merklyrefuelv2CallerSession) Endpoint() (common.Address, error) {
	return _Merklyrefuelv2.Contract.Endpoint(&_Merklyrefuelv2.CallOpts)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Merklyrefuelv2 *Merklyrefuelv2Caller) NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error) {
	var out []interface{}
	err := _Merklyrefuelv2.contract.Call(opts, &out, "nextNonce", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Merklyrefuelv2 *Merklyrefuelv2Session) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _Merklyrefuelv2.Contract.NextNonce(&_Merklyrefuelv2.CallOpts, arg0, arg1)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Merklyrefuelv2 *Merklyrefuelv2CallerSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _Merklyrefuelv2.Contract.NextNonce(&_Merklyrefuelv2.CallOpts, arg0, arg1)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Merklyrefuelv2 *Merklyrefuelv2Caller) OAppVersion(opts *bind.CallOpts) (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	var out []interface{}
	err := _Merklyrefuelv2.contract.Call(opts, &out, "oAppVersion")

	outstruct := new(struct {
		SenderVersion   uint64
		ReceiverVersion uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SenderVersion = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ReceiverVersion = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Merklyrefuelv2 *Merklyrefuelv2Session) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _Merklyrefuelv2.Contract.OAppVersion(&_Merklyrefuelv2.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Merklyrefuelv2 *Merklyrefuelv2CallerSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _Merklyrefuelv2.Contract.OAppVersion(&_Merklyrefuelv2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Merklyrefuelv2 *Merklyrefuelv2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Merklyrefuelv2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Merklyrefuelv2 *Merklyrefuelv2Session) Owner() (common.Address, error) {
	return _Merklyrefuelv2.Contract.Owner(&_Merklyrefuelv2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Merklyrefuelv2 *Merklyrefuelv2CallerSession) Owner() (common.Address, error) {
	return _Merklyrefuelv2.Contract.Owner(&_Merklyrefuelv2.CallOpts)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Merklyrefuelv2 *Merklyrefuelv2Caller) Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error) {
	var out []interface{}
	err := _Merklyrefuelv2.contract.Call(opts, &out, "peers", eid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Merklyrefuelv2 *Merklyrefuelv2Session) Peers(eid uint32) ([32]byte, error) {
	return _Merklyrefuelv2.Contract.Peers(&_Merklyrefuelv2.CallOpts, eid)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Merklyrefuelv2 *Merklyrefuelv2CallerSession) Peers(eid uint32) ([32]byte, error) {
	return _Merklyrefuelv2.Contract.Peers(&_Merklyrefuelv2.CallOpts, eid)
}

// Quote is a free data retrieval call binding the contract method 0xf77e5dd3.
//
// Solidity: function quote(uint32 _dstEid, string _message, bytes _options, bool _payInLzToken) view returns(uint256 nativeFee, uint256 lzTokenFee)
func (_Merklyrefuelv2 *Merklyrefuelv2Caller) Quote(opts *bind.CallOpts, _dstEid uint32, _message string, _options []byte, _payInLzToken bool) (struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}, error) {
	var out []interface{}
	err := _Merklyrefuelv2.contract.Call(opts, &out, "quote", _dstEid, _message, _options, _payInLzToken)

	outstruct := new(struct {
		NativeFee  *big.Int
		LzTokenFee *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NativeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LzTokenFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Quote is a free data retrieval call binding the contract method 0xf77e5dd3.
//
// Solidity: function quote(uint32 _dstEid, string _message, bytes _options, bool _payInLzToken) view returns(uint256 nativeFee, uint256 lzTokenFee)
func (_Merklyrefuelv2 *Merklyrefuelv2Session) Quote(_dstEid uint32, _message string, _options []byte, _payInLzToken bool) (struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}, error) {
	return _Merklyrefuelv2.Contract.Quote(&_Merklyrefuelv2.CallOpts, _dstEid, _message, _options, _payInLzToken)
}

// Quote is a free data retrieval call binding the contract method 0xf77e5dd3.
//
// Solidity: function quote(uint32 _dstEid, string _message, bytes _options, bool _payInLzToken) view returns(uint256 nativeFee, uint256 lzTokenFee)
func (_Merklyrefuelv2 *Merklyrefuelv2CallerSession) Quote(_dstEid uint32, _message string, _options []byte, _payInLzToken bool) (struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}, error) {
	return _Merklyrefuelv2.Contract.Quote(&_Merklyrefuelv2.CallOpts, _dstEid, _message, _options, _payInLzToken)
}

// BridgeGas is a paid mutator transaction binding the contract method 0x52b94e4d.
//
// Solidity: function bridgeGas(uint32 _dstEid, string _message, bytes _options) payable returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Transactor) BridgeGas(opts *bind.TransactOpts, _dstEid uint32, _message string, _options []byte) (*types.Transaction, error) {
	return _Merklyrefuelv2.contract.Transact(opts, "bridgeGas", _dstEid, _message, _options)
}

// BridgeGas is a paid mutator transaction binding the contract method 0x52b94e4d.
//
// Solidity: function bridgeGas(uint32 _dstEid, string _message, bytes _options) payable returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Session) BridgeGas(_dstEid uint32, _message string, _options []byte) (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.BridgeGas(&_Merklyrefuelv2.TransactOpts, _dstEid, _message, _options)
}

// BridgeGas is a paid mutator transaction binding the contract method 0x52b94e4d.
//
// Solidity: function bridgeGas(uint32 _dstEid, string _message, bytes _options) payable returns()
func (_Merklyrefuelv2 *Merklyrefuelv2TransactorSession) BridgeGas(_dstEid uint32, _message string, _options []byte) (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.BridgeGas(&_Merklyrefuelv2.TransactOpts, _dstEid, _message, _options)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Transactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Merklyrefuelv2.contract.Transact(opts, "lzReceive", _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Session) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.LzReceive(&_Merklyrefuelv2.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Merklyrefuelv2 *Merklyrefuelv2TransactorSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.LzReceive(&_Merklyrefuelv2.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merklyrefuelv2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Session) RenounceOwnership() (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.RenounceOwnership(&_Merklyrefuelv2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Merklyrefuelv2 *Merklyrefuelv2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.RenounceOwnership(&_Merklyrefuelv2.TransactOpts)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Transactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _Merklyrefuelv2.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Session) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.SetDelegate(&_Merklyrefuelv2.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Merklyrefuelv2 *Merklyrefuelv2TransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.SetDelegate(&_Merklyrefuelv2.TransactOpts, _delegate)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Transactor) SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Merklyrefuelv2.contract.Transact(opts, "setPeer", _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Session) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.SetPeer(&_Merklyrefuelv2.TransactOpts, _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Merklyrefuelv2 *Merklyrefuelv2TransactorSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.SetPeer(&_Merklyrefuelv2.TransactOpts, _eid, _peer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Merklyrefuelv2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.TransferOwnership(&_Merklyrefuelv2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Merklyrefuelv2 *Merklyrefuelv2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.TransferOwnership(&_Merklyrefuelv2.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merklyrefuelv2.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Session) Withdraw() (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.Withdraw(&_Merklyrefuelv2.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Merklyrefuelv2 *Merklyrefuelv2TransactorSession) Withdraw() (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.Withdraw(&_Merklyrefuelv2.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Merklyrefuelv2.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.Fallback(&_Merklyrefuelv2.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Merklyrefuelv2 *Merklyrefuelv2TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.Fallback(&_Merklyrefuelv2.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merklyrefuelv2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Merklyrefuelv2 *Merklyrefuelv2Session) Receive() (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.Receive(&_Merklyrefuelv2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Merklyrefuelv2 *Merklyrefuelv2TransactorSession) Receive() (*types.Transaction, error) {
	return _Merklyrefuelv2.Contract.Receive(&_Merklyrefuelv2.TransactOpts)
}

// Merklyrefuelv2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Merklyrefuelv2 contract.
type Merklyrefuelv2OwnershipTransferredIterator struct {
	Event *Merklyrefuelv2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Merklyrefuelv2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Merklyrefuelv2OwnershipTransferred)
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
		it.Event = new(Merklyrefuelv2OwnershipTransferred)
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
func (it *Merklyrefuelv2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Merklyrefuelv2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Merklyrefuelv2OwnershipTransferred represents a OwnershipTransferred event raised by the Merklyrefuelv2 contract.
type Merklyrefuelv2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Merklyrefuelv2 *Merklyrefuelv2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Merklyrefuelv2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Merklyrefuelv2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Merklyrefuelv2OwnershipTransferredIterator{contract: _Merklyrefuelv2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Merklyrefuelv2 *Merklyrefuelv2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Merklyrefuelv2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Merklyrefuelv2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Merklyrefuelv2OwnershipTransferred)
				if err := _Merklyrefuelv2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Merklyrefuelv2 *Merklyrefuelv2Filterer) ParseOwnershipTransferred(log types.Log) (*Merklyrefuelv2OwnershipTransferred, error) {
	event := new(Merklyrefuelv2OwnershipTransferred)
	if err := _Merklyrefuelv2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Merklyrefuelv2PeerSetIterator is returned from FilterPeerSet and is used to iterate over the raw logs and unpacked data for PeerSet events raised by the Merklyrefuelv2 contract.
type Merklyrefuelv2PeerSetIterator struct {
	Event *Merklyrefuelv2PeerSet // Event containing the contract specifics and raw log

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
func (it *Merklyrefuelv2PeerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Merklyrefuelv2PeerSet)
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
		it.Event = new(Merklyrefuelv2PeerSet)
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
func (it *Merklyrefuelv2PeerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Merklyrefuelv2PeerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Merklyrefuelv2PeerSet represents a PeerSet event raised by the Merklyrefuelv2 contract.
type Merklyrefuelv2PeerSet struct {
	Eid  uint32
	Peer [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPeerSet is a free log retrieval operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Merklyrefuelv2 *Merklyrefuelv2Filterer) FilterPeerSet(opts *bind.FilterOpts) (*Merklyrefuelv2PeerSetIterator, error) {

	logs, sub, err := _Merklyrefuelv2.contract.FilterLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return &Merklyrefuelv2PeerSetIterator{contract: _Merklyrefuelv2.contract, event: "PeerSet", logs: logs, sub: sub}, nil
}

// WatchPeerSet is a free log subscription operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Merklyrefuelv2 *Merklyrefuelv2Filterer) WatchPeerSet(opts *bind.WatchOpts, sink chan<- *Merklyrefuelv2PeerSet) (event.Subscription, error) {

	logs, sub, err := _Merklyrefuelv2.contract.WatchLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Merklyrefuelv2PeerSet)
				if err := _Merklyrefuelv2.contract.UnpackLog(event, "PeerSet", log); err != nil {
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

// ParsePeerSet is a log parse operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Merklyrefuelv2 *Merklyrefuelv2Filterer) ParsePeerSet(log types.Log) (*Merklyrefuelv2PeerSet, error) {
	event := new(Merklyrefuelv2PeerSet)
	if err := _Merklyrefuelv2.contract.UnpackLog(event, "PeerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
