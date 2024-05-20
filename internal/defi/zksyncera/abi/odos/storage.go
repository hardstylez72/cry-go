// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package odos

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

// OdosRouterV2inputTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type OdosRouterV2inputTokenInfo struct {
	TokenAddress common.Address
	AmountIn     *big.Int
	Receiver     common.Address
}

// OdosRouterV2outputTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type OdosRouterV2outputTokenInfo struct {
	TokenAddress  common.Address
	RelativeValue *big.Int
	Receiver      common.Address
}

// OdosRouterV2permit2Info is an auto generated low-level Go binding around an user-defined struct.
type OdosRouterV2permit2Info struct {
	ContractAddress common.Address
	Nonce           *big.Int
	Deadline        *big.Int
	Signature       []byte
}

// OdosRouterV2swapTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type OdosRouterV2swapTokenInfo struct {
	InputToken     common.Address
	InputAmount    *big.Int
	InputReceiver  common.Address
	OutputToken    common.Address
	OutputQuote    *big.Int
	OutputMin      *big.Int
	OutputReceiver common.Address
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"slippage\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"referralCode\",\"type\":\"uint32\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokensIn\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokensOut\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"referralCode\",\"type\":\"uint32\"}],\"name\":\"SwapMulti\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_DENOM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REFERRAL_WITH_FEE_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addressList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"referralLookup\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"referralFee\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_referralCode\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_referralFee\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"registerReferralCode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_swapMultiFee\",\"type\":\"uint256\"}],\"name\":\"setSwapMultiFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"outputQuote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputReceiver\",\"type\":\"address\"}],\"internalType\":\"structOdosRouterV2.swapTokenInfo\",\"name\":\"tokenInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"pathDefinition\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"referralCode\",\"type\":\"uint32\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapCompact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structOdosRouterV2.inputTokenInfo[]\",\"name\":\"inputs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"relativeValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structOdosRouterV2.outputTokenInfo[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"valueOutMin\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pathDefinition\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"referralCode\",\"type\":\"uint32\"}],\"name\":\"swapMulti\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapMultiCompact\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapMultiFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOdosRouterV2.permit2Info\",\"name\":\"permit2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structOdosRouterV2.inputTokenInfo[]\",\"name\":\"inputs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"relativeValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structOdosRouterV2.outputTokenInfo[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"valueOutMin\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pathDefinition\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"referralCode\",\"type\":\"uint32\"}],\"name\":\"swapMultiPermit2\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOdosRouterV2.permit2Info\",\"name\":\"permit2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"outputQuote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputReceiver\",\"type\":\"address\"}],\"internalType\":\"structOdosRouterV2.swapTokenInfo\",\"name\":\"tokenInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"pathDefinition\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"referralCode\",\"type\":\"uint32\"}],\"name\":\"swapPermit2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structOdosRouterV2.inputTokenInfo[]\",\"name\":\"inputs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"relativeValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structOdosRouterV2.outputTokenInfo[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"valueOutMin\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pathDefinition\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"swapRouterFunds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"transferRouterFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"writeAddressList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// FEEDENOM is a free data retrieval call binding the contract method 0x4886c675.
//
// Solidity: function FEE_DENOM() view returns(uint256)
func (_Storage *StorageCaller) FEEDENOM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "FEE_DENOM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEDENOM is a free data retrieval call binding the contract method 0x4886c675.
//
// Solidity: function FEE_DENOM() view returns(uint256)
func (_Storage *StorageSession) FEEDENOM() (*big.Int, error) {
	return _Storage.Contract.FEEDENOM(&_Storage.CallOpts)
}

// FEEDENOM is a free data retrieval call binding the contract method 0x4886c675.
//
// Solidity: function FEE_DENOM() view returns(uint256)
func (_Storage *StorageCallerSession) FEEDENOM() (*big.Int, error) {
	return _Storage.Contract.FEEDENOM(&_Storage.CallOpts)
}

// REFERRALWITHFEETHRESHOLD is a free data retrieval call binding the contract method 0x6c082c13.
//
// Solidity: function REFERRAL_WITH_FEE_THRESHOLD() view returns(uint256)
func (_Storage *StorageCaller) REFERRALWITHFEETHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "REFERRAL_WITH_FEE_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REFERRALWITHFEETHRESHOLD is a free data retrieval call binding the contract method 0x6c082c13.
//
// Solidity: function REFERRAL_WITH_FEE_THRESHOLD() view returns(uint256)
func (_Storage *StorageSession) REFERRALWITHFEETHRESHOLD() (*big.Int, error) {
	return _Storage.Contract.REFERRALWITHFEETHRESHOLD(&_Storage.CallOpts)
}

// REFERRALWITHFEETHRESHOLD is a free data retrieval call binding the contract method 0x6c082c13.
//
// Solidity: function REFERRAL_WITH_FEE_THRESHOLD() view returns(uint256)
func (_Storage *StorageCallerSession) REFERRALWITHFEETHRESHOLD() (*big.Int, error) {
	return _Storage.Contract.REFERRALWITHFEETHRESHOLD(&_Storage.CallOpts)
}

// AddressList is a free data retrieval call binding the contract method 0xb810fb43.
//
// Solidity: function addressList(uint256 ) view returns(address)
func (_Storage *StorageCaller) AddressList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "addressList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressList is a free data retrieval call binding the contract method 0xb810fb43.
//
// Solidity: function addressList(uint256 ) view returns(address)
func (_Storage *StorageSession) AddressList(arg0 *big.Int) (common.Address, error) {
	return _Storage.Contract.AddressList(&_Storage.CallOpts, arg0)
}

// AddressList is a free data retrieval call binding the contract method 0xb810fb43.
//
// Solidity: function addressList(uint256 ) view returns(address)
func (_Storage *StorageCallerSession) AddressList(arg0 *big.Int) (common.Address, error) {
	return _Storage.Contract.AddressList(&_Storage.CallOpts, arg0)
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

// ReferralLookup is a free data retrieval call binding the contract method 0xf827065e.
//
// Solidity: function referralLookup(uint32 ) view returns(uint64 referralFee, address beneficiary, bool registered)
func (_Storage *StorageCaller) ReferralLookup(opts *bind.CallOpts, arg0 uint32) (struct {
	ReferralFee uint64
	Beneficiary common.Address
	Registered  bool
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "referralLookup", arg0)

	outstruct := new(struct {
		ReferralFee uint64
		Beneficiary common.Address
		Registered  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReferralFee = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Beneficiary = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Registered = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// ReferralLookup is a free data retrieval call binding the contract method 0xf827065e.
//
// Solidity: function referralLookup(uint32 ) view returns(uint64 referralFee, address beneficiary, bool registered)
func (_Storage *StorageSession) ReferralLookup(arg0 uint32) (struct {
	ReferralFee uint64
	Beneficiary common.Address
	Registered  bool
}, error) {
	return _Storage.Contract.ReferralLookup(&_Storage.CallOpts, arg0)
}

// ReferralLookup is a free data retrieval call binding the contract method 0xf827065e.
//
// Solidity: function referralLookup(uint32 ) view returns(uint64 referralFee, address beneficiary, bool registered)
func (_Storage *StorageCallerSession) ReferralLookup(arg0 uint32) (struct {
	ReferralFee uint64
	Beneficiary common.Address
	Registered  bool
}, error) {
	return _Storage.Contract.ReferralLookup(&_Storage.CallOpts, arg0)
}

// SwapMultiFee is a free data retrieval call binding the contract method 0xe7d3fc60.
//
// Solidity: function swapMultiFee() view returns(uint256)
func (_Storage *StorageCaller) SwapMultiFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "swapMultiFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapMultiFee is a free data retrieval call binding the contract method 0xe7d3fc60.
//
// Solidity: function swapMultiFee() view returns(uint256)
func (_Storage *StorageSession) SwapMultiFee() (*big.Int, error) {
	return _Storage.Contract.SwapMultiFee(&_Storage.CallOpts)
}

// SwapMultiFee is a free data retrieval call binding the contract method 0xe7d3fc60.
//
// Solidity: function swapMultiFee() view returns(uint256)
func (_Storage *StorageCallerSession) SwapMultiFee() (*big.Int, error) {
	return _Storage.Contract.SwapMultiFee(&_Storage.CallOpts)
}

// RegisterReferralCode is a paid mutator transaction binding the contract method 0xe10895f9.
//
// Solidity: function registerReferralCode(uint32 _referralCode, uint64 _referralFee, address _beneficiary) returns()
func (_Storage *StorageTransactor) RegisterReferralCode(opts *bind.TransactOpts, _referralCode uint32, _referralFee uint64, _beneficiary common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "registerReferralCode", _referralCode, _referralFee, _beneficiary)
}

// RegisterReferralCode is a paid mutator transaction binding the contract method 0xe10895f9.
//
// Solidity: function registerReferralCode(uint32 _referralCode, uint64 _referralFee, address _beneficiary) returns()
func (_Storage *StorageSession) RegisterReferralCode(_referralCode uint32, _referralFee uint64, _beneficiary common.Address) (*types.Transaction, error) {
	return _Storage.Contract.RegisterReferralCode(&_Storage.TransactOpts, _referralCode, _referralFee, _beneficiary)
}

// RegisterReferralCode is a paid mutator transaction binding the contract method 0xe10895f9.
//
// Solidity: function registerReferralCode(uint32 _referralCode, uint64 _referralFee, address _beneficiary) returns()
func (_Storage *StorageTransactorSession) RegisterReferralCode(_referralCode uint32, _referralFee uint64, _beneficiary common.Address) (*types.Transaction, error) {
	return _Storage.Contract.RegisterReferralCode(&_Storage.TransactOpts, _referralCode, _referralFee, _beneficiary)
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

// SetSwapMultiFee is a paid mutator transaction binding the contract method 0x9286b93d.
//
// Solidity: function setSwapMultiFee(uint256 _swapMultiFee) returns()
func (_Storage *StorageTransactor) SetSwapMultiFee(opts *bind.TransactOpts, _swapMultiFee *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setSwapMultiFee", _swapMultiFee)
}

// SetSwapMultiFee is a paid mutator transaction binding the contract method 0x9286b93d.
//
// Solidity: function setSwapMultiFee(uint256 _swapMultiFee) returns()
func (_Storage *StorageSession) SetSwapMultiFee(_swapMultiFee *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetSwapMultiFee(&_Storage.TransactOpts, _swapMultiFee)
}

// SetSwapMultiFee is a paid mutator transaction binding the contract method 0x9286b93d.
//
// Solidity: function setSwapMultiFee(uint256 _swapMultiFee) returns()
func (_Storage *StorageTransactorSession) SetSwapMultiFee(_swapMultiFee *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetSwapMultiFee(&_Storage.TransactOpts, _swapMultiFee)
}

// Swap is a paid mutator transaction binding the contract method 0x3b635ce4.
//
// Solidity: function swap((address,uint256,address,address,uint256,uint256,address) tokenInfo, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256 amountOut)
func (_Storage *StorageTransactor) Swap(opts *bind.TransactOpts, tokenInfo OdosRouterV2swapTokenInfo, pathDefinition []byte, executor common.Address, referralCode uint32) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swap", tokenInfo, pathDefinition, executor, referralCode)
}

// Swap is a paid mutator transaction binding the contract method 0x3b635ce4.
//
// Solidity: function swap((address,uint256,address,address,uint256,uint256,address) tokenInfo, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256 amountOut)
func (_Storage *StorageSession) Swap(tokenInfo OdosRouterV2swapTokenInfo, pathDefinition []byte, executor common.Address, referralCode uint32) (*types.Transaction, error) {
	return _Storage.Contract.Swap(&_Storage.TransactOpts, tokenInfo, pathDefinition, executor, referralCode)
}

// Swap is a paid mutator transaction binding the contract method 0x3b635ce4.
//
// Solidity: function swap((address,uint256,address,address,uint256,uint256,address) tokenInfo, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256 amountOut)
func (_Storage *StorageTransactorSession) Swap(tokenInfo OdosRouterV2swapTokenInfo, pathDefinition []byte, executor common.Address, referralCode uint32) (*types.Transaction, error) {
	return _Storage.Contract.Swap(&_Storage.TransactOpts, tokenInfo, pathDefinition, executor, referralCode)
}

// SwapCompact is a paid mutator transaction binding the contract method 0x83bd37f9.
//
// Solidity: function swapCompact() payable returns(uint256)
func (_Storage *StorageTransactor) SwapCompact(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapCompact")
}

// SwapCompact is a paid mutator transaction binding the contract method 0x83bd37f9.
//
// Solidity: function swapCompact() payable returns(uint256)
func (_Storage *StorageSession) SwapCompact() (*types.Transaction, error) {
	return _Storage.Contract.SwapCompact(&_Storage.TransactOpts)
}

// SwapCompact is a paid mutator transaction binding the contract method 0x83bd37f9.
//
// Solidity: function swapCompact() payable returns(uint256)
func (_Storage *StorageTransactorSession) SwapCompact() (*types.Transaction, error) {
	return _Storage.Contract.SwapCompact(&_Storage.TransactOpts)
}

// SwapMulti is a paid mutator transaction binding the contract method 0x7bf2d6d4.
//
// Solidity: function swapMulti((address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256[] amountsOut)
func (_Storage *StorageTransactor) SwapMulti(opts *bind.TransactOpts, inputs []OdosRouterV2inputTokenInfo, outputs []OdosRouterV2outputTokenInfo, valueOutMin *big.Int, pathDefinition []byte, executor common.Address, referralCode uint32) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapMulti", inputs, outputs, valueOutMin, pathDefinition, executor, referralCode)
}

// SwapMulti is a paid mutator transaction binding the contract method 0x7bf2d6d4.
//
// Solidity: function swapMulti((address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256[] amountsOut)
func (_Storage *StorageSession) SwapMulti(inputs []OdosRouterV2inputTokenInfo, outputs []OdosRouterV2outputTokenInfo, valueOutMin *big.Int, pathDefinition []byte, executor common.Address, referralCode uint32) (*types.Transaction, error) {
	return _Storage.Contract.SwapMulti(&_Storage.TransactOpts, inputs, outputs, valueOutMin, pathDefinition, executor, referralCode)
}

// SwapMulti is a paid mutator transaction binding the contract method 0x7bf2d6d4.
//
// Solidity: function swapMulti((address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256[] amountsOut)
func (_Storage *StorageTransactorSession) SwapMulti(inputs []OdosRouterV2inputTokenInfo, outputs []OdosRouterV2outputTokenInfo, valueOutMin *big.Int, pathDefinition []byte, executor common.Address, referralCode uint32) (*types.Transaction, error) {
	return _Storage.Contract.SwapMulti(&_Storage.TransactOpts, inputs, outputs, valueOutMin, pathDefinition, executor, referralCode)
}

// SwapMultiCompact is a paid mutator transaction binding the contract method 0x84a7f3dd.
//
// Solidity: function swapMultiCompact() payable returns(uint256[] amountsOut)
func (_Storage *StorageTransactor) SwapMultiCompact(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapMultiCompact")
}

// SwapMultiCompact is a paid mutator transaction binding the contract method 0x84a7f3dd.
//
// Solidity: function swapMultiCompact() payable returns(uint256[] amountsOut)
func (_Storage *StorageSession) SwapMultiCompact() (*types.Transaction, error) {
	return _Storage.Contract.SwapMultiCompact(&_Storage.TransactOpts)
}

// SwapMultiCompact is a paid mutator transaction binding the contract method 0x84a7f3dd.
//
// Solidity: function swapMultiCompact() payable returns(uint256[] amountsOut)
func (_Storage *StorageTransactorSession) SwapMultiCompact() (*types.Transaction, error) {
	return _Storage.Contract.SwapMultiCompact(&_Storage.TransactOpts)
}

// SwapMultiPermit2 is a paid mutator transaction binding the contract method 0x080c25b3.
//
// Solidity: function swapMultiPermit2((address,uint256,uint256,bytes) permit2, (address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256[] amountsOut)
func (_Storage *StorageTransactor) SwapMultiPermit2(opts *bind.TransactOpts, permit2 OdosRouterV2permit2Info, inputs []OdosRouterV2inputTokenInfo, outputs []OdosRouterV2outputTokenInfo, valueOutMin *big.Int, pathDefinition []byte, executor common.Address, referralCode uint32) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapMultiPermit2", permit2, inputs, outputs, valueOutMin, pathDefinition, executor, referralCode)
}

// SwapMultiPermit2 is a paid mutator transaction binding the contract method 0x080c25b3.
//
// Solidity: function swapMultiPermit2((address,uint256,uint256,bytes) permit2, (address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256[] amountsOut)
func (_Storage *StorageSession) SwapMultiPermit2(permit2 OdosRouterV2permit2Info, inputs []OdosRouterV2inputTokenInfo, outputs []OdosRouterV2outputTokenInfo, valueOutMin *big.Int, pathDefinition []byte, executor common.Address, referralCode uint32) (*types.Transaction, error) {
	return _Storage.Contract.SwapMultiPermit2(&_Storage.TransactOpts, permit2, inputs, outputs, valueOutMin, pathDefinition, executor, referralCode)
}

// SwapMultiPermit2 is a paid mutator transaction binding the contract method 0x080c25b3.
//
// Solidity: function swapMultiPermit2((address,uint256,uint256,bytes) permit2, (address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256[] amountsOut)
func (_Storage *StorageTransactorSession) SwapMultiPermit2(permit2 OdosRouterV2permit2Info, inputs []OdosRouterV2inputTokenInfo, outputs []OdosRouterV2outputTokenInfo, valueOutMin *big.Int, pathDefinition []byte, executor common.Address, referralCode uint32) (*types.Transaction, error) {
	return _Storage.Contract.SwapMultiPermit2(&_Storage.TransactOpts, permit2, inputs, outputs, valueOutMin, pathDefinition, executor, referralCode)
}

// SwapPermit2 is a paid mutator transaction binding the contract method 0x87b621b5.
//
// Solidity: function swapPermit2((address,uint256,uint256,bytes) permit2, (address,uint256,address,address,uint256,uint256,address) tokenInfo, bytes pathDefinition, address executor, uint32 referralCode) returns(uint256 amountOut)
func (_Storage *StorageTransactor) SwapPermit2(opts *bind.TransactOpts, permit2 OdosRouterV2permit2Info, tokenInfo OdosRouterV2swapTokenInfo, pathDefinition []byte, executor common.Address, referralCode uint32) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapPermit2", permit2, tokenInfo, pathDefinition, executor, referralCode)
}

// SwapPermit2 is a paid mutator transaction binding the contract method 0x87b621b5.
//
// Solidity: function swapPermit2((address,uint256,uint256,bytes) permit2, (address,uint256,address,address,uint256,uint256,address) tokenInfo, bytes pathDefinition, address executor, uint32 referralCode) returns(uint256 amountOut)
func (_Storage *StorageSession) SwapPermit2(permit2 OdosRouterV2permit2Info, tokenInfo OdosRouterV2swapTokenInfo, pathDefinition []byte, executor common.Address, referralCode uint32) (*types.Transaction, error) {
	return _Storage.Contract.SwapPermit2(&_Storage.TransactOpts, permit2, tokenInfo, pathDefinition, executor, referralCode)
}

// SwapPermit2 is a paid mutator transaction binding the contract method 0x87b621b5.
//
// Solidity: function swapPermit2((address,uint256,uint256,bytes) permit2, (address,uint256,address,address,uint256,uint256,address) tokenInfo, bytes pathDefinition, address executor, uint32 referralCode) returns(uint256 amountOut)
func (_Storage *StorageTransactorSession) SwapPermit2(permit2 OdosRouterV2permit2Info, tokenInfo OdosRouterV2swapTokenInfo, pathDefinition []byte, executor common.Address, referralCode uint32) (*types.Transaction, error) {
	return _Storage.Contract.SwapPermit2(&_Storage.TransactOpts, permit2, tokenInfo, pathDefinition, executor, referralCode)
}

// SwapRouterFunds is a paid mutator transaction binding the contract method 0x28be42f4.
//
// Solidity: function swapRouterFunds((address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor) returns(uint256[] amountsOut)
func (_Storage *StorageTransactor) SwapRouterFunds(opts *bind.TransactOpts, inputs []OdosRouterV2inputTokenInfo, outputs []OdosRouterV2outputTokenInfo, valueOutMin *big.Int, pathDefinition []byte, executor common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapRouterFunds", inputs, outputs, valueOutMin, pathDefinition, executor)
}

// SwapRouterFunds is a paid mutator transaction binding the contract method 0x28be42f4.
//
// Solidity: function swapRouterFunds((address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor) returns(uint256[] amountsOut)
func (_Storage *StorageSession) SwapRouterFunds(inputs []OdosRouterV2inputTokenInfo, outputs []OdosRouterV2outputTokenInfo, valueOutMin *big.Int, pathDefinition []byte, executor common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SwapRouterFunds(&_Storage.TransactOpts, inputs, outputs, valueOutMin, pathDefinition, executor)
}

// SwapRouterFunds is a paid mutator transaction binding the contract method 0x28be42f4.
//
// Solidity: function swapRouterFunds((address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor) returns(uint256[] amountsOut)
func (_Storage *StorageTransactorSession) SwapRouterFunds(inputs []OdosRouterV2inputTokenInfo, outputs []OdosRouterV2outputTokenInfo, valueOutMin *big.Int, pathDefinition []byte, executor common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SwapRouterFunds(&_Storage.TransactOpts, inputs, outputs, valueOutMin, pathDefinition, executor)
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

// TransferRouterFunds is a paid mutator transaction binding the contract method 0x174da621.
//
// Solidity: function transferRouterFunds(address[] tokens, uint256[] amounts, address dest) returns()
func (_Storage *StorageTransactor) TransferRouterFunds(opts *bind.TransactOpts, tokens []common.Address, amounts []*big.Int, dest common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transferRouterFunds", tokens, amounts, dest)
}

// TransferRouterFunds is a paid mutator transaction binding the contract method 0x174da621.
//
// Solidity: function transferRouterFunds(address[] tokens, uint256[] amounts, address dest) returns()
func (_Storage *StorageSession) TransferRouterFunds(tokens []common.Address, amounts []*big.Int, dest common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferRouterFunds(&_Storage.TransactOpts, tokens, amounts, dest)
}

// TransferRouterFunds is a paid mutator transaction binding the contract method 0x174da621.
//
// Solidity: function transferRouterFunds(address[] tokens, uint256[] amounts, address dest) returns()
func (_Storage *StorageTransactorSession) TransferRouterFunds(tokens []common.Address, amounts []*big.Int, dest common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferRouterFunds(&_Storage.TransactOpts, tokens, amounts, dest)
}

// WriteAddressList is a paid mutator transaction binding the contract method 0x3596f9a2.
//
// Solidity: function writeAddressList(address[] addresses) returns()
func (_Storage *StorageTransactor) WriteAddressList(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "writeAddressList", addresses)
}

// WriteAddressList is a paid mutator transaction binding the contract method 0x3596f9a2.
//
// Solidity: function writeAddressList(address[] addresses) returns()
func (_Storage *StorageSession) WriteAddressList(addresses []common.Address) (*types.Transaction, error) {
	return _Storage.Contract.WriteAddressList(&_Storage.TransactOpts, addresses)
}

// WriteAddressList is a paid mutator transaction binding the contract method 0x3596f9a2.
//
// Solidity: function writeAddressList(address[] addresses) returns()
func (_Storage *StorageTransactorSession) WriteAddressList(addresses []common.Address) (*types.Transaction, error) {
	return _Storage.Contract.WriteAddressList(&_Storage.TransactOpts, addresses)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Storage *StorageTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Storage *StorageSession) Receive() (*types.Transaction, error) {
	return _Storage.Contract.Receive(&_Storage.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Storage *StorageTransactorSession) Receive() (*types.Transaction, error) {
	return _Storage.Contract.Receive(&_Storage.TransactOpts)
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

// StorageSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Storage contract.
type StorageSwapIterator struct {
	Event *StorageSwap // Event containing the contract specifics and raw log

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
func (it *StorageSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSwap)
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
		it.Event = new(StorageSwap)
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
func (it *StorageSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSwap represents a Swap event raised by the Storage contract.
type StorageSwap struct {
	Sender       common.Address
	InputAmount  *big.Int
	InputToken   common.Address
	AmountOut    *big.Int
	OutputToken  common.Address
	Slippage     *big.Int
	ReferralCode uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x823eaf01002d7353fbcadb2ea3305cc46fa35d799cb0914846d185ac06f8ad05.
//
// Solidity: event Swap(address sender, uint256 inputAmount, address inputToken, uint256 amountOut, address outputToken, int256 slippage, uint32 referralCode)
func (_Storage *StorageFilterer) FilterSwap(opts *bind.FilterOpts) (*StorageSwapIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &StorageSwapIterator{contract: _Storage.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x823eaf01002d7353fbcadb2ea3305cc46fa35d799cb0914846d185ac06f8ad05.
//
// Solidity: event Swap(address sender, uint256 inputAmount, address inputToken, uint256 amountOut, address outputToken, int256 slippage, uint32 referralCode)
func (_Storage *StorageFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *StorageSwap) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSwap)
				if err := _Storage.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x823eaf01002d7353fbcadb2ea3305cc46fa35d799cb0914846d185ac06f8ad05.
//
// Solidity: event Swap(address sender, uint256 inputAmount, address inputToken, uint256 amountOut, address outputToken, int256 slippage, uint32 referralCode)
func (_Storage *StorageFilterer) ParseSwap(log types.Log) (*StorageSwap, error) {
	event := new(StorageSwap)
	if err := _Storage.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSwapMultiIterator is returned from FilterSwapMulti and is used to iterate over the raw logs and unpacked data for SwapMulti events raised by the Storage contract.
type StorageSwapMultiIterator struct {
	Event *StorageSwapMulti // Event containing the contract specifics and raw log

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
func (it *StorageSwapMultiIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSwapMulti)
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
		it.Event = new(StorageSwapMulti)
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
func (it *StorageSwapMultiIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSwapMultiIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSwapMulti represents a SwapMulti event raised by the Storage contract.
type StorageSwapMulti struct {
	Sender       common.Address
	AmountsIn    []*big.Int
	TokensIn     []common.Address
	AmountsOut   []*big.Int
	TokensOut    []common.Address
	ReferralCode uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwapMulti is a free log retrieval operation binding the contract event 0x7d7fb03518253ae01913536628b78d6d82e63e19b943aab5f4948356021259be.
//
// Solidity: event SwapMulti(address sender, uint256[] amountsIn, address[] tokensIn, uint256[] amountsOut, address[] tokensOut, uint32 referralCode)
func (_Storage *StorageFilterer) FilterSwapMulti(opts *bind.FilterOpts) (*StorageSwapMultiIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "SwapMulti")
	if err != nil {
		return nil, err
	}
	return &StorageSwapMultiIterator{contract: _Storage.contract, event: "SwapMulti", logs: logs, sub: sub}, nil
}

// WatchSwapMulti is a free log subscription operation binding the contract event 0x7d7fb03518253ae01913536628b78d6d82e63e19b943aab5f4948356021259be.
//
// Solidity: event SwapMulti(address sender, uint256[] amountsIn, address[] tokensIn, uint256[] amountsOut, address[] tokensOut, uint32 referralCode)
func (_Storage *StorageFilterer) WatchSwapMulti(opts *bind.WatchOpts, sink chan<- *StorageSwapMulti) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "SwapMulti")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSwapMulti)
				if err := _Storage.contract.UnpackLog(event, "SwapMulti", log); err != nil {
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

// ParseSwapMulti is a log parse operation binding the contract event 0x7d7fb03518253ae01913536628b78d6d82e63e19b943aab5f4948356021259be.
//
// Solidity: event SwapMulti(address sender, uint256[] amountsIn, address[] tokensIn, uint256[] amountsOut, address[] tokensOut, uint32 referralCode)
func (_Storage *StorageFilterer) ParseSwapMulti(log types.Log) (*StorageSwapMulti, error) {
	event := new(StorageSwapMulti)
	if err := _Storage.contract.UnpackLog(event, "SwapMulti", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
