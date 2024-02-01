// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2pass

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

// RefuelMetaData contains all meta data concerning the Refuel contract.
var RefuelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lzEndpoint_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasRefuelPrice_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"nativeForDst\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addressOnDst\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useZro\",\"type\":\"bool\"}],\"name\":\"estimateGasRefuelFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zroFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nativeForDst\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addressOnDst\",\"type\":\"address\"}],\"name\":\"gasRefuel\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasRefuelPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasRefuelPrice_\",\"type\":\"uint256\"}],\"name\":\"setGasRefuelPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RefuelABI is the input ABI used to generate the binding from.
// Deprecated: Use RefuelMetaData.ABI instead.
var RefuelABI = RefuelMetaData.ABI

// Refuel is an auto generated Go binding around an Ethereum contract.
type Refuel struct {
	RefuelCaller     // Read-only binding to the contract
	RefuelTransactor // Write-only binding to the contract
	RefuelFilterer   // Log filterer for contract events
}

// RefuelCaller is an auto generated read-only Go binding around an Ethereum contract.
type RefuelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefuelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RefuelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefuelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RefuelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefuelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RefuelSession struct {
	Contract     *Refuel           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RefuelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RefuelCallerSession struct {
	Contract *RefuelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RefuelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RefuelTransactorSession struct {
	Contract     *RefuelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RefuelRaw is an auto generated low-level Go binding around an Ethereum contract.
type RefuelRaw struct {
	Contract *Refuel // Generic contract binding to access the raw methods on
}

// RefuelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RefuelCallerRaw struct {
	Contract *RefuelCaller // Generic read-only contract binding to access the raw methods on
}

// RefuelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RefuelTransactorRaw struct {
	Contract *RefuelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRefuel creates a new instance of Refuel, bound to a specific deployed contract.
func NewRefuel(address common.Address, backend bind.ContractBackend) (*Refuel, error) {
	contract, err := bindRefuel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Refuel{RefuelCaller: RefuelCaller{contract: contract}, RefuelTransactor: RefuelTransactor{contract: contract}, RefuelFilterer: RefuelFilterer{contract: contract}}, nil
}

// NewRefuelCaller creates a new read-only instance of Refuel, bound to a specific deployed contract.
func NewRefuelCaller(address common.Address, caller bind.ContractCaller) (*RefuelCaller, error) {
	contract, err := bindRefuel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RefuelCaller{contract: contract}, nil
}

// NewRefuelTransactor creates a new write-only instance of Refuel, bound to a specific deployed contract.
func NewRefuelTransactor(address common.Address, transactor bind.ContractTransactor) (*RefuelTransactor, error) {
	contract, err := bindRefuel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RefuelTransactor{contract: contract}, nil
}

// NewRefuelFilterer creates a new log filterer instance of Refuel, bound to a specific deployed contract.
func NewRefuelFilterer(address common.Address, filterer bind.ContractFilterer) (*RefuelFilterer, error) {
	contract, err := bindRefuel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RefuelFilterer{contract: contract}, nil
}

// bindRefuel binds a generic wrapper to an already deployed contract.
func bindRefuel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RefuelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Refuel *RefuelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Refuel.Contract.RefuelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Refuel *RefuelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Refuel.Contract.RefuelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Refuel *RefuelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Refuel.Contract.RefuelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Refuel *RefuelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Refuel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Refuel *RefuelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Refuel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Refuel *RefuelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Refuel.Contract.contract.Transact(opts, method, params...)
}

// EstimateGasRefuelFee is a free data retrieval call binding the contract method 0x7a22cb76.
//
// Solidity: function estimateGasRefuelFee(uint16 dstChainId, uint256 nativeForDst, address addressOnDst, bool useZro) view returns(uint256 nativeFee, uint256 zroFee)
func (_Refuel *RefuelCaller) EstimateGasRefuelFee(opts *bind.CallOpts, dstChainId uint16, nativeForDst *big.Int, addressOnDst common.Address, useZro bool) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Refuel.contract.Call(opts, &out, "estimateGasRefuelFee", dstChainId, nativeForDst, addressOnDst, useZro)

	outstruct := new(struct {
		NativeFee *big.Int
		ZroFee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NativeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ZroFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateGasRefuelFee is a free data retrieval call binding the contract method 0x7a22cb76.
//
// Solidity: function estimateGasRefuelFee(uint16 dstChainId, uint256 nativeForDst, address addressOnDst, bool useZro) view returns(uint256 nativeFee, uint256 zroFee)
func (_Refuel *RefuelSession) EstimateGasRefuelFee(dstChainId uint16, nativeForDst *big.Int, addressOnDst common.Address, useZro bool) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Refuel.Contract.EstimateGasRefuelFee(&_Refuel.CallOpts, dstChainId, nativeForDst, addressOnDst, useZro)
}

// EstimateGasRefuelFee is a free data retrieval call binding the contract method 0x7a22cb76.
//
// Solidity: function estimateGasRefuelFee(uint16 dstChainId, uint256 nativeForDst, address addressOnDst, bool useZro) view returns(uint256 nativeFee, uint256 zroFee)
func (_Refuel *RefuelCallerSession) EstimateGasRefuelFee(dstChainId uint16, nativeForDst *big.Int, addressOnDst common.Address, useZro bool) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Refuel.Contract.EstimateGasRefuelFee(&_Refuel.CallOpts, dstChainId, nativeForDst, addressOnDst, useZro)
}

// GasRefuelPrice is a free data retrieval call binding the contract method 0xbbf98963.
//
// Solidity: function gasRefuelPrice() view returns(uint256)
func (_Refuel *RefuelCaller) GasRefuelPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Refuel.contract.Call(opts, &out, "gasRefuelPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasRefuelPrice is a free data retrieval call binding the contract method 0xbbf98963.
//
// Solidity: function gasRefuelPrice() view returns(uint256)
func (_Refuel *RefuelSession) GasRefuelPrice() (*big.Int, error) {
	return _Refuel.Contract.GasRefuelPrice(&_Refuel.CallOpts)
}

// GasRefuelPrice is a free data retrieval call binding the contract method 0xbbf98963.
//
// Solidity: function gasRefuelPrice() view returns(uint256)
func (_Refuel *RefuelCallerSession) GasRefuelPrice() (*big.Int, error) {
	return _Refuel.Contract.GasRefuelPrice(&_Refuel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Refuel *RefuelCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Refuel.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Refuel *RefuelSession) Owner() (common.Address, error) {
	return _Refuel.Contract.Owner(&_Refuel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Refuel *RefuelCallerSession) Owner() (common.Address, error) {
	return _Refuel.Contract.Owner(&_Refuel.CallOpts)
}

// GasRefuel is a paid mutator transaction binding the contract method 0x32accbb9.
//
// Solidity: function gasRefuel(uint16 dstChainId, address zroPaymentAddress, uint256 nativeForDst, address addressOnDst) payable returns()
func (_Refuel *RefuelTransactor) GasRefuel(opts *bind.TransactOpts, dstChainId uint16, zroPaymentAddress common.Address, nativeForDst *big.Int, addressOnDst common.Address) (*types.Transaction, error) {
	return _Refuel.contract.Transact(opts, "gasRefuel", dstChainId, zroPaymentAddress, nativeForDst, addressOnDst)
}

// GasRefuel is a paid mutator transaction binding the contract method 0x32accbb9.
//
// Solidity: function gasRefuel(uint16 dstChainId, address zroPaymentAddress, uint256 nativeForDst, address addressOnDst) payable returns()
func (_Refuel *RefuelSession) GasRefuel(dstChainId uint16, zroPaymentAddress common.Address, nativeForDst *big.Int, addressOnDst common.Address) (*types.Transaction, error) {
	return _Refuel.Contract.GasRefuel(&_Refuel.TransactOpts, dstChainId, zroPaymentAddress, nativeForDst, addressOnDst)
}

// GasRefuel is a paid mutator transaction binding the contract method 0x32accbb9.
//
// Solidity: function gasRefuel(uint16 dstChainId, address zroPaymentAddress, uint256 nativeForDst, address addressOnDst) payable returns()
func (_Refuel *RefuelTransactorSession) GasRefuel(dstChainId uint16, zroPaymentAddress common.Address, nativeForDst *big.Int, addressOnDst common.Address) (*types.Transaction, error) {
	return _Refuel.Contract.GasRefuel(&_Refuel.TransactOpts, dstChainId, zroPaymentAddress, nativeForDst, addressOnDst)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 , bytes , uint64 , bytes ) returns()
func (_Refuel *RefuelTransactor) LzReceive(opts *bind.TransactOpts, arg0 uint16, arg1 []byte, arg2 uint64, arg3 []byte) (*types.Transaction, error) {
	return _Refuel.contract.Transact(opts, "lzReceive", arg0, arg1, arg2, arg3)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 , bytes , uint64 , bytes ) returns()
func (_Refuel *RefuelSession) LzReceive(arg0 uint16, arg1 []byte, arg2 uint64, arg3 []byte) (*types.Transaction, error) {
	return _Refuel.Contract.LzReceive(&_Refuel.TransactOpts, arg0, arg1, arg2, arg3)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 , bytes , uint64 , bytes ) returns()
func (_Refuel *RefuelTransactorSession) LzReceive(arg0 uint16, arg1 []byte, arg2 uint64, arg3 []byte) (*types.Transaction, error) {
	return _Refuel.Contract.LzReceive(&_Refuel.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Refuel *RefuelTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Refuel.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Refuel *RefuelSession) RenounceOwnership() (*types.Transaction, error) {
	return _Refuel.Contract.RenounceOwnership(&_Refuel.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Refuel *RefuelTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Refuel.Contract.RenounceOwnership(&_Refuel.TransactOpts)
}

// SetGasRefuelPrice is a paid mutator transaction binding the contract method 0x7e93c69d.
//
// Solidity: function setGasRefuelPrice(uint256 gasRefuelPrice_) returns()
func (_Refuel *RefuelTransactor) SetGasRefuelPrice(opts *bind.TransactOpts, gasRefuelPrice_ *big.Int) (*types.Transaction, error) {
	return _Refuel.contract.Transact(opts, "setGasRefuelPrice", gasRefuelPrice_)
}

// SetGasRefuelPrice is a paid mutator transaction binding the contract method 0x7e93c69d.
//
// Solidity: function setGasRefuelPrice(uint256 gasRefuelPrice_) returns()
func (_Refuel *RefuelSession) SetGasRefuelPrice(gasRefuelPrice_ *big.Int) (*types.Transaction, error) {
	return _Refuel.Contract.SetGasRefuelPrice(&_Refuel.TransactOpts, gasRefuelPrice_)
}

// SetGasRefuelPrice is a paid mutator transaction binding the contract method 0x7e93c69d.
//
// Solidity: function setGasRefuelPrice(uint256 gasRefuelPrice_) returns()
func (_Refuel *RefuelTransactorSession) SetGasRefuelPrice(gasRefuelPrice_ *big.Int) (*types.Transaction, error) {
	return _Refuel.Contract.SetGasRefuelPrice(&_Refuel.TransactOpts, gasRefuelPrice_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Refuel *RefuelTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Refuel.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Refuel *RefuelSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Refuel.Contract.TransferOwnership(&_Refuel.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Refuel *RefuelTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Refuel.Contract.TransferOwnership(&_Refuel.TransactOpts, newOwner)
}

// RefuelOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Refuel contract.
type RefuelOwnershipTransferredIterator struct {
	Event *RefuelOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RefuelOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefuelOwnershipTransferred)
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
		it.Event = new(RefuelOwnershipTransferred)
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
func (it *RefuelOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefuelOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefuelOwnershipTransferred represents a OwnershipTransferred event raised by the Refuel contract.
type RefuelOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Refuel *RefuelFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RefuelOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Refuel.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RefuelOwnershipTransferredIterator{contract: _Refuel.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Refuel *RefuelFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RefuelOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Refuel.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefuelOwnershipTransferred)
				if err := _Refuel.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Refuel *RefuelFilterer) ParseOwnershipTransferred(log types.Log) (*RefuelOwnershipTransferred, error) {
	event := new(RefuelOwnershipTransferred)
	if err := _Refuel.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
