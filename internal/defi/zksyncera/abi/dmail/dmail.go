// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dmail

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

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"name\":null,\"type\":\"constructor\",\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"gas\":null,\"_isFragment\":true},{\"name\":\"AdminChanged\",\"anonymous\":false,\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true}],\"type\":\"event\",\"_isFragment\":true},{\"name\":\"Approval\",\"anonymous\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"uint256\",\"_isParamType\":true}],\"type\":\"event\",\"_isFragment\":true},{\"name\":\"BeaconUpgraded\",\"anonymous\":false,\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true}],\"type\":\"event\",\"_isFragment\":true},{\"name\":\"Initialized\",\"anonymous\":false,\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"uint8\",\"_isParamType\":true}],\"type\":\"event\",\"_isFragment\":true},{\"name\":\"Message\",\"anonymous\":false,\"inputs\":[{\"name\":\"subject\",\"type\":\"string\",\"indexed\":true,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"string\",\"_isParamType\":true},{\"name\":\"from_address\",\"type\":\"address\",\"indexed\":true,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"to\",\"type\":\"string\",\"indexed\":true,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"string\",\"_isParamType\":true}],\"type\":\"event\",\"_isFragment\":true},{\"name\":\"OwnershipTransferred\",\"anonymous\":false,\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true}],\"type\":\"event\",\"_isFragment\":true},{\"name\":\"Paused\",\"anonymous\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true}],\"type\":\"event\",\"_isFragment\":true},{\"name\":\"Transfer\",\"anonymous\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"uint256\",\"_isParamType\":true}],\"type\":\"event\",\"_isFragment\":true},{\"name\":\"Unpaused\",\"anonymous\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true}],\"type\":\"event\",\"_isFragment\":true},{\"name\":\"Upgraded\",\"anonymous\":false,\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true}],\"type\":\"event\",\"_isFragment\":true},{\"type\":\"function\",\"name\":\"allowance\",\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true}],\"outputs\":[{\"name\":null,\"type\":\"uint256\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"uint256\",\"_isParamType\":true}],\"payable\":false,\"stateMutability\":\"view\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"approve\",\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"uint256\",\"_isParamType\":true}],\"outputs\":[{\"name\":null,\"type\":\"bool\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"bool\",\"_isParamType\":true}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"balanceOf\",\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true}],\"outputs\":[{\"name\":null,\"type\":\"uint256\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"uint256\",\"_isParamType\":true}],\"payable\":false,\"stateMutability\":\"view\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"burn\",\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"uint256\",\"_isParamType\":true}],\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"burnFrom\",\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"uint256\",\"_isParamType\":true}],\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"decimals\",\"constant\":true,\"inputs\":[],\"outputs\":[{\"name\":null,\"type\":\"uint8\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"uint8\",\"_isParamType\":true}],\"payable\":false,\"stateMutability\":\"view\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"uint256\",\"_isParamType\":true}],\"outputs\":[{\"name\":null,\"type\":\"bool\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"bool\",\"_isParamType\":true}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"addedValue\",\"type\":\"uint256\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"uint256\",\"_isParamType\":true}],\"outputs\":[{\"name\":null,\"type\":\"bool\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"bool\",\"_isParamType\":true}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"initialize\",\"constant\":false,\"inputs\":[],\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"mint\",\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"uint256\",\"_isParamType\":true}],\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"name\",\"constant\":true,\"inputs\":[],\"outputs\":[{\"name\":null,\"type\":\"string\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"string\",\"_isParamType\":true}],\"payable\":false,\"stateMutability\":\"view\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"owner\",\"constant\":true,\"inputs\":[],\"outputs\":[{\"name\":null,\"type\":\"address\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true}],\"payable\":false,\"stateMutability\":\"view\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"pause\",\"constant\":false,\"inputs\":[],\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"paused\",\"constant\":true,\"inputs\":[],\"outputs\":[{\"name\":null,\"type\":\"bool\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"bool\",\"_isParamType\":true}],\"payable\":false,\"stateMutability\":\"view\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"constant\":true,\"inputs\":[],\"outputs\":[{\"name\":null,\"type\":\"bytes32\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"bytes32\",\"_isParamType\":true}],\"payable\":false,\"stateMutability\":\"view\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"constant\":false,\"inputs\":[],\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"send_mail\",\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"string\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"string\",\"_isParamType\":true},{\"name\":\"subject\",\"type\":\"string\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"string\",\"_isParamType\":true}],\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"symbol\",\"constant\":true,\"inputs\":[],\"outputs\":[{\"name\":null,\"type\":\"string\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"string\",\"_isParamType\":true}],\"payable\":false,\"stateMutability\":\"view\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"totalSupply\",\"constant\":true,\"inputs\":[],\"outputs\":[{\"name\":null,\"type\":\"uint256\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"uint256\",\"_isParamType\":true}],\"payable\":false,\"stateMutability\":\"view\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"transfer\",\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"uint256\",\"_isParamType\":true}],\"outputs\":[{\"name\":null,\"type\":\"bool\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"bool\",\"_isParamType\":true}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"transferFrom\",\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"to\",\"type\":\"address\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"uint256\",\"_isParamType\":true}],\"outputs\":[{\"name\":null,\"type\":\"bool\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"bool\",\"_isParamType\":true}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"transferOwnership\",\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true}],\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"unpause\",\"constant\":false,\"inputs\":[],\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"upgradeTo\",\"constant\":false,\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true}],\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"gas\":null,\"_isFragment\":true},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"constant\":false,\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"address\",\"_isParamType\":true},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":null,\"components\":null,\"arrayLength\":null,\"arrayChildren\":null,\"baseType\":\"bytes\",\"_isParamType\":true}],\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"gas\":null,\"_isFragment\":true}]",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Storage *StorageCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Storage *StorageSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Storage.Contract.Allowance(&_Storage.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Storage *StorageCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Storage.Contract.Allowance(&_Storage.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Storage *StorageCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Storage *StorageSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Storage.Contract.BalanceOf(&_Storage.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Storage *StorageCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Storage.Contract.BalanceOf(&_Storage.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Storage *StorageCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Storage *StorageSession) Decimals() (uint8, error) {
	return _Storage.Contract.Decimals(&_Storage.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Storage *StorageCallerSession) Decimals() (uint8, error) {
	return _Storage.Contract.Decimals(&_Storage.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Storage *StorageCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Storage *StorageSession) Name() (string, error) {
	return _Storage.Contract.Name(&_Storage.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Storage *StorageCallerSession) Name() (string, error) {
	return _Storage.Contract.Name(&_Storage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageCallerSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Storage *StorageCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Storage *StorageSession) Paused() (bool, error) {
	return _Storage.Contract.Paused(&_Storage.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Storage *StorageCallerSession) Paused() (bool, error) {
	return _Storage.Contract.Paused(&_Storage.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Storage *StorageCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Storage *StorageSession) ProxiableUUID() ([32]byte, error) {
	return _Storage.Contract.ProxiableUUID(&_Storage.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Storage *StorageCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Storage.Contract.ProxiableUUID(&_Storage.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Storage *StorageCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Storage *StorageSession) Symbol() (string, error) {
	return _Storage.Contract.Symbol(&_Storage.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Storage *StorageCallerSession) Symbol() (string, error) {
	return _Storage.Contract.Symbol(&_Storage.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Storage *StorageCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Storage *StorageSession) TotalSupply() (*big.Int, error) {
	return _Storage.Contract.TotalSupply(&_Storage.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Storage *StorageCallerSession) TotalSupply() (*big.Int, error) {
	return _Storage.Contract.TotalSupply(&_Storage.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Storage *StorageTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Storage *StorageSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Approve(&_Storage.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Storage *StorageTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Approve(&_Storage.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Storage *StorageTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Storage *StorageSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Burn(&_Storage.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Storage *StorageTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Burn(&_Storage.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Storage *StorageTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Storage *StorageSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.BurnFrom(&_Storage.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Storage *StorageTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.BurnFrom(&_Storage.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Storage *StorageTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Storage *StorageSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DecreaseAllowance(&_Storage.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Storage *StorageTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DecreaseAllowance(&_Storage.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Storage *StorageTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Storage *StorageSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.IncreaseAllowance(&_Storage.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Storage *StorageTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.IncreaseAllowance(&_Storage.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Storage *StorageTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Storage *StorageSession) Initialize() (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Storage *StorageTransactorSession) Initialize() (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Storage *StorageTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Storage *StorageSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Mint(&_Storage.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Storage *StorageTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Mint(&_Storage.TransactOpts, to, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Storage *StorageTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Storage *StorageSession) Pause() (*types.Transaction, error) {
	return _Storage.Contract.Pause(&_Storage.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Storage *StorageTransactorSession) Pause() (*types.Transaction, error) {
	return _Storage.Contract.Pause(&_Storage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _Storage.Contract.RenounceOwnership(&_Storage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Storage.Contract.RenounceOwnership(&_Storage.TransactOpts)
}

// SendMail is a paid mutator transaction binding the contract method 0x5b7d7482.
//
// Solidity: function send_mail(string to, string subject) returns()
func (_Storage *StorageTransactor) SendMail(opts *bind.TransactOpts, to string, subject string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "send_mail", to, subject)
}

// SendMail is a paid mutator transaction binding the contract method 0x5b7d7482.
//
// Solidity: function send_mail(string to, string subject) returns()
func (_Storage *StorageSession) SendMail(to string, subject string) (*types.Transaction, error) {
	return _Storage.Contract.SendMail(&_Storage.TransactOpts, to, subject)
}

// SendMail is a paid mutator transaction binding the contract method 0x5b7d7482.
//
// Solidity: function send_mail(string to, string subject) returns()
func (_Storage *StorageTransactorSession) SendMail(to string, subject string) (*types.Transaction, error) {
	return _Storage.Contract.SendMail(&_Storage.TransactOpts, to, subject)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Storage *StorageTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Storage *StorageSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Transfer(&_Storage.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Storage *StorageTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Transfer(&_Storage.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Storage *StorageTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Storage *StorageSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TransferFrom(&_Storage.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Storage *StorageTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TransferFrom(&_Storage.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwnership(&_Storage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwnership(&_Storage.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Storage *StorageTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Storage *StorageSession) Unpause() (*types.Transaction, error) {
	return _Storage.Contract.Unpause(&_Storage.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Storage *StorageTransactorSession) Unpause() (*types.Transaction, error) {
	return _Storage.Contract.Unpause(&_Storage.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Storage *StorageTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Storage *StorageSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeTo(&_Storage.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Storage *StorageTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeTo(&_Storage.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Storage *StorageTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Storage *StorageSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeToAndCall(&_Storage.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Storage *StorageTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeToAndCall(&_Storage.TransactOpts, newImplementation, data)
}

// StorageAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Storage contract.
type StorageAdminChangedIterator struct {
	Event *StorageAdminChanged // Event containing the contract specifics and raw log

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
func (it *StorageAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageAdminChanged)
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
		it.Event = new(StorageAdminChanged)
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
func (it *StorageAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageAdminChanged represents a AdminChanged event raised by the Storage contract.
type StorageAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Storage *StorageFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*StorageAdminChangedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &StorageAdminChangedIterator{contract: _Storage.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Storage *StorageFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *StorageAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageAdminChanged)
				if err := _Storage.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Storage *StorageFilterer) ParseAdminChanged(log types.Log) (*StorageAdminChanged, error) {
	event := new(StorageAdminChanged)
	if err := _Storage.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Storage contract.
type StorageApprovalIterator struct {
	Event *StorageApproval // Event containing the contract specifics and raw log

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
func (it *StorageApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageApproval)
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
		it.Event = new(StorageApproval)
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
func (it *StorageApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageApproval represents a Approval event raised by the Storage contract.
type StorageApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Storage *StorageFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StorageApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StorageApprovalIterator{contract: _Storage.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Storage *StorageFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StorageApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageApproval)
				if err := _Storage.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Storage *StorageFilterer) ParseApproval(log types.Log) (*StorageApproval, error) {
	event := new(StorageApproval)
	if err := _Storage.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Storage contract.
type StorageBeaconUpgradedIterator struct {
	Event *StorageBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *StorageBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageBeaconUpgraded)
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
		it.Event = new(StorageBeaconUpgraded)
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
func (it *StorageBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageBeaconUpgraded represents a BeaconUpgraded event raised by the Storage contract.
type StorageBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Storage *StorageFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*StorageBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &StorageBeaconUpgradedIterator{contract: _Storage.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Storage *StorageFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *StorageBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageBeaconUpgraded)
				if err := _Storage.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Storage *StorageFilterer) ParseBeaconUpgraded(log types.Log) (*StorageBeaconUpgraded, error) {
	event := new(StorageBeaconUpgraded)
	if err := _Storage.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Storage contract.
type StorageInitializedIterator struct {
	Event *StorageInitialized // Event containing the contract specifics and raw log

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
func (it *StorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageInitialized)
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
		it.Event = new(StorageInitialized)
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
func (it *StorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageInitialized represents a Initialized event raised by the Storage contract.
type StorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Storage *StorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*StorageInitializedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StorageInitializedIterator{contract: _Storage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Storage *StorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StorageInitialized) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageInitialized)
				if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Storage *StorageFilterer) ParseInitialized(log types.Log) (*StorageInitialized, error) {
	event := new(StorageInitialized)
	if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageMessageIterator is returned from FilterMessage and is used to iterate over the raw logs and unpacked data for Message events raised by the Storage contract.
type StorageMessageIterator struct {
	Event *StorageMessage // Event containing the contract specifics and raw log

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
func (it *StorageMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageMessage)
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
		it.Event = new(StorageMessage)
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
func (it *StorageMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageMessage represents a Message event raised by the Storage contract.
type StorageMessage struct {
	Subject     common.Hash
	FromAddress common.Address
	To          common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessage is a free log retrieval operation binding the contract event 0xf26bfd49b39c52efaf04ee7f21ca2fdc73c680fada92ab7a8f1ea37b350bcf8c.
//
// Solidity: event Message(string indexed subject, address indexed from_address, string indexed to)
func (_Storage *StorageFilterer) FilterMessage(opts *bind.FilterOpts, subject []string, from_address []common.Address, to []string) (*StorageMessageIterator, error) {

	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var from_addressRule []interface{}
	for _, from_addressItem := range from_address {
		from_addressRule = append(from_addressRule, from_addressItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Message", subjectRule, from_addressRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StorageMessageIterator{contract: _Storage.contract, event: "Message", logs: logs, sub: sub}, nil
}

// WatchMessage is a free log subscription operation binding the contract event 0xf26bfd49b39c52efaf04ee7f21ca2fdc73c680fada92ab7a8f1ea37b350bcf8c.
//
// Solidity: event Message(string indexed subject, address indexed from_address, string indexed to)
func (_Storage *StorageFilterer) WatchMessage(opts *bind.WatchOpts, sink chan<- *StorageMessage, subject []string, from_address []common.Address, to []string) (event.Subscription, error) {

	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var from_addressRule []interface{}
	for _, from_addressItem := range from_address {
		from_addressRule = append(from_addressRule, from_addressItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Message", subjectRule, from_addressRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageMessage)
				if err := _Storage.contract.UnpackLog(event, "Message", log); err != nil {
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

// ParseMessage is a log parse operation binding the contract event 0xf26bfd49b39c52efaf04ee7f21ca2fdc73c680fada92ab7a8f1ea37b350bcf8c.
//
// Solidity: event Message(string indexed subject, address indexed from_address, string indexed to)
func (_Storage *StorageFilterer) ParseMessage(log types.Log) (*StorageMessage, error) {
	event := new(StorageMessage)
	if err := _Storage.contract.UnpackLog(event, "Message", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Storage contract.
type StorageOwnershipTransferredIterator struct {
	Event *StorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageOwnershipTransferred)
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
		it.Event = new(StorageOwnershipTransferred)
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
func (it *StorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageOwnershipTransferred represents a OwnershipTransferred event raised by the Storage contract.
type StorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StorageOwnershipTransferredIterator{contract: _Storage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageOwnershipTransferred)
				if err := _Storage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Storage *StorageFilterer) ParseOwnershipTransferred(log types.Log) (*StorageOwnershipTransferred, error) {
	event := new(StorageOwnershipTransferred)
	if err := _Storage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoragePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Storage contract.
type StoragePausedIterator struct {
	Event *StoragePaused // Event containing the contract specifics and raw log

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
func (it *StoragePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoragePaused)
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
		it.Event = new(StoragePaused)
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
func (it *StoragePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoragePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoragePaused represents a Paused event raised by the Storage contract.
type StoragePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Storage *StorageFilterer) FilterPaused(opts *bind.FilterOpts) (*StoragePausedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StoragePausedIterator{contract: _Storage.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Storage *StorageFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StoragePaused) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoragePaused)
				if err := _Storage.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Storage *StorageFilterer) ParsePaused(log types.Log) (*StoragePaused, error) {
	event := new(StoragePaused)
	if err := _Storage.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Storage contract.
type StorageTransferIterator struct {
	Event *StorageTransfer // Event containing the contract specifics and raw log

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
func (it *StorageTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageTransfer)
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
		it.Event = new(StorageTransfer)
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
func (it *StorageTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageTransfer represents a Transfer event raised by the Storage contract.
type StorageTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Storage *StorageFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StorageTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StorageTransferIterator{contract: _Storage.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Storage *StorageFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StorageTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageTransfer)
				if err := _Storage.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Storage *StorageFilterer) ParseTransfer(log types.Log) (*StorageTransfer, error) {
	event := new(StorageTransfer)
	if err := _Storage.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Storage contract.
type StorageUnpausedIterator struct {
	Event *StorageUnpaused // Event containing the contract specifics and raw log

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
func (it *StorageUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageUnpaused)
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
		it.Event = new(StorageUnpaused)
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
func (it *StorageUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageUnpaused represents a Unpaused event raised by the Storage contract.
type StorageUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Storage *StorageFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StorageUnpausedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StorageUnpausedIterator{contract: _Storage.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Storage *StorageFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StorageUnpaused) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageUnpaused)
				if err := _Storage.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Storage *StorageFilterer) ParseUnpaused(log types.Log) (*StorageUnpaused, error) {
	event := new(StorageUnpaused)
	if err := _Storage.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Storage contract.
type StorageUpgradedIterator struct {
	Event *StorageUpgraded // Event containing the contract specifics and raw log

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
func (it *StorageUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageUpgraded)
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
		it.Event = new(StorageUpgraded)
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
func (it *StorageUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageUpgraded represents a Upgraded event raised by the Storage contract.
type StorageUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Storage *StorageFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StorageUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StorageUpgradedIterator{contract: _Storage.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Storage *StorageFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StorageUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageUpgraded)
				if err := _Storage.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Storage *StorageFilterer) ParseUpgraded(log types.Log) (*StorageUpgraded, error) {
	event := new(StorageUpgraded)
	if err := _Storage.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
