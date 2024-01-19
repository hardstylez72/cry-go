// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package core_dao

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

// LzLibCallParams is an auto generated low-level Go binding around an user-defined struct.


// BridgecoreMetaData contains all meta data concerning the Bridgecore contract.
var BridgecoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"name\":\"_endpoint\",\"internalType\":\"address\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"indexed\":false,\"name\":\"_srcChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"_srcAddress\",\"internalType\":\"bytes\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"_nonce\",\"internalType\":\"uint64\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"_payload\",\"internalType\":\"bytes\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"_reason\",\"internalType\":\"bytes\",\"type\":\"bytes\"}],\"name\":\"MessageFailed\",\"anonymous\":false,\"type\":\"event\"},{\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"internalType\":\"address\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"anonymous\":false,\"type\":\"event\"},{\"inputs\":[{\"indexed\":false,\"name\":\"localToken\",\"internalType\":\"address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"remoteChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"remoteToken\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"RegisterToken\",\"anonymous\":false,\"type\":\"event\"},{\"inputs\":[{\"indexed\":false,\"name\":\"_srcChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"_srcAddress\",\"internalType\":\"bytes\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"_nonce\",\"internalType\":\"uint64\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"_payloadHash\",\"internalType\":\"bytes32\",\"type\":\"bytes32\"}],\"name\":\"RetryMessageSuccess\",\"anonymous\":false,\"type\":\"event\"},{\"inputs\":[{\"indexed\":false,\"name\":\"_dstChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"_type\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"_minDstGas\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"name\":\"SetMinDstGas\",\"anonymous\":false,\"type\":\"event\"},{\"inputs\":[{\"indexed\":false,\"name\":\"precrime\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"SetPrecrime\",\"anonymous\":false,\"type\":\"event\"},{\"inputs\":[{\"indexed\":false,\"name\":\"_remoteChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"_path\",\"internalType\":\"bytes\",\"type\":\"bytes\"}],\"name\":\"SetTrustedRemote\",\"anonymous\":false,\"type\":\"event\"},{\"inputs\":[{\"indexed\":false,\"name\":\"_remoteChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"_remoteAddress\",\"internalType\":\"bytes\",\"type\":\"bytes\"}],\"name\":\"SetTrustedRemoteAddress\",\"anonymous\":false,\"type\":\"event\"},{\"inputs\":[{\"indexed\":false,\"name\":\"useCustomAdapterParams\",\"internalType\":\"bool\",\"type\":\"bool\"}],\"name\":\"SetUseCustomAdapterParams\",\"anonymous\":false,\"type\":\"event\"},{\"inputs\":[{\"indexed\":false,\"name\":\"withdrawalFeeBps\",\"internalType\":\"uint16\",\"type\":\"uint16\"}],\"name\":\"SetWithdrawalFeeBps\",\"anonymous\":false,\"type\":\"event\"},{\"inputs\":[{\"indexed\":false,\"name\":\"localToken\",\"internalType\":\"address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"remoteToken\",\"internalType\":\"address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"remoteChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"to\",\"internalType\":\"address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"name\":\"UnwrapToken\",\"anonymous\":false,\"type\":\"event\"},{\"inputs\":[{\"indexed\":false,\"name\":\"localToken\",\"internalType\":\"address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"remoteToken\",\"internalType\":\"address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"remoteChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"to\",\"internalType\":\"address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"name\":\"WrapToken\",\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"uint8\",\"type\":\"uint8\"}],\"inputs\":[],\"name\":\"PT_MINT\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"uint8\",\"type\":\"uint8\"}],\"inputs\":[],\"name\":\"PT_UNLOCK\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"uint16\",\"type\":\"uint16\"}],\"inputs\":[],\"name\":\"TOTAL_BPS\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"localToken\",\"internalType\":\"address\",\"type\":\"address\"},{\"name\":\"remoteChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"amount\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"to\",\"internalType\":\"address\",\"type\":\"address\"},{\"name\":\"unwrapWeth\",\"internalType\":\"bool\",\"type\":\"bool\"},{\"components\":[{\"name\":\"refundAddress\",\"internalType\":\"addresspayable\",\"type\":\"address\"},{\"name\":\"zroPaymentAddress\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"callParams\",\"internalType\":\"structLzLib.CallParams\",\"type\":\"tuple\"},{\"name\":\"adapterParams\",\"internalType\":\"bytes\",\"type\":\"bytes\"}],\"name\":\"bridge\",\"stateMutability\":\"payable\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"nativeFee\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"zroFee\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"inputs\":[{\"name\":\"remoteChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"useZro\",\"internalType\":\"bool\",\"type\":\"bool\"},{\"name\":\"adapterParams\",\"internalType\":\"bytes\",\"type\":\"bytes\"}],\"name\":\"estimateBridgeFee\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"bytes32\",\"type\":\"bytes32\"}],\"inputs\":[{\"name\":\"\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"\",\"internalType\":\"bytes\",\"type\":\"bytes\"},{\"name\":\"\",\"internalType\":\"uint64\",\"type\":\"uint64\"}],\"name\":\"failedMessages\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_srcChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"_srcAddress\",\"internalType\":\"bytes\",\"type\":\"bytes\"}],\"name\":\"forceResumeReceive\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"bytes\",\"type\":\"bytes\"}],\"inputs\":[{\"name\":\"_version\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"_chainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"\",\"internalType\":\"address\",\"type\":\"address\"},{\"name\":\"_configType\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"name\":\"getConfig\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"bytes\",\"type\":\"bytes\"}],\"inputs\":[{\"name\":\"_remoteChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"}],\"name\":\"getTrustedRemoteAddress\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"bool\",\"type\":\"bool\"}],\"inputs\":[{\"name\":\"_srcChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"_srcAddress\",\"internalType\":\"bytes\",\"type\":\"bytes\"}],\"name\":\"isTrustedRemote\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"address\",\"type\":\"address\"}],\"inputs\":[{\"name\":\"\",\"internalType\":\"address\",\"type\":\"address\"},{\"name\":\"\",\"internalType\":\"uint16\",\"type\":\"uint16\"}],\"name\":\"localToRemote\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"contractILayerZeroEndpoint\",\"type\":\"address\"}],\"inputs\":[],\"name\":\"lzEndpoint\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_srcChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"_srcAddress\",\"internalType\":\"bytes\",\"type\":\"bytes\"},{\"name\":\"_nonce\",\"internalType\":\"uint64\",\"type\":\"uint64\"},{\"name\":\"_payload\",\"internalType\":\"bytes\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"inputs\":[{\"name\":\"\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"\",\"internalType\":\"uint16\",\"type\":\"uint16\"}],\"name\":\"minDstGasLookup\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_srcChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"_srcAddress\",\"internalType\":\"bytes\",\"type\":\"bytes\"},{\"name\":\"_nonce\",\"internalType\":\"uint64\",\"type\":\"uint64\"},{\"name\":\"_payload\",\"internalType\":\"bytes\",\"type\":\"bytes\"}],\"name\":\"nonblockingLzReceive\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"address\",\"type\":\"address\"}],\"inputs\":[],\"name\":\"owner\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"address\",\"type\":\"address\"}],\"inputs\":[],\"name\":\"precrime\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"localToken\",\"internalType\":\"address\",\"type\":\"address\"},{\"name\":\"remoteChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"remoteToken\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"registerToken\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"address\",\"type\":\"address\"}],\"inputs\":[{\"name\":\"\",\"internalType\":\"address\",\"type\":\"address\"},{\"name\":\"\",\"internalType\":\"uint16\",\"type\":\"uint16\"}],\"name\":\"remoteToLocal\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[],\"name\":\"renounceOwnership\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_srcChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"_srcAddress\",\"internalType\":\"bytes\",\"type\":\"bytes\"},{\"name\":\"_nonce\",\"internalType\":\"uint64\",\"type\":\"uint64\"},{\"name\":\"_payload\",\"internalType\":\"bytes\",\"type\":\"bytes\"}],\"name\":\"retryMessage\",\"stateMutability\":\"payable\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_version\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"_chainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"_configType\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"_config\",\"internalType\":\"bytes\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_dstChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"_packetType\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"_minGas\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"name\":\"setMinDstGas\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_precrime\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"setPrecrime\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_version\",\"internalType\":\"uint16\",\"type\":\"uint16\"}],\"name\":\"setReceiveVersion\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_version\",\"internalType\":\"uint16\",\"type\":\"uint16\"}],\"name\":\"setSendVersion\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_srcChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"_path\",\"internalType\":\"bytes\",\"type\":\"bytes\"}],\"name\":\"setTrustedRemote\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_remoteChainId\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"_remoteAddress\",\"internalType\":\"bytes\",\"type\":\"bytes\"}],\"name\":\"setTrustedRemoteAddress\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_useCustomAdapterParams\",\"internalType\":\"bool\",\"type\":\"bool\"}],\"name\":\"setUseCustomAdapterParams\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"_withdrawalFeeBps\",\"internalType\":\"uint16\",\"type\":\"uint16\"}],\"name\":\"setWithdrawalFeeBps\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"inputs\":[{\"name\":\"\",\"internalType\":\"uint16\",\"type\":\"uint16\"},{\"name\":\"\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"totalValueLocked\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[],\"inputs\":[{\"name\":\"newOwner\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"bytes\",\"type\":\"bytes\"}],\"inputs\":[{\"name\":\"\",\"internalType\":\"uint16\",\"type\":\"uint16\"}],\"name\":\"trustedRemoteLookup\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"bool\",\"type\":\"bool\"}],\"inputs\":[],\"name\":\"useCustomAdapterParams\",\"stateMutability\":\"view\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"\",\"internalType\":\"uint16\",\"type\":\"uint16\"}],\"inputs\":[],\"name\":\"withdrawalFeeBps\",\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BridgecoreABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgecoreMetaData.ABI instead.
var BridgecoreABI = BridgecoreMetaData.ABI

// Bridgecore is an auto generated Go binding around an Ethereum contract.
type Bridgecore struct {
	BridgecoreCaller     // Read-only binding to the contract
	BridgecoreTransactor // Write-only binding to the contract
	BridgecoreFilterer   // Log filterer for contract events
}

// BridgecoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgecoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgecoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgecoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgecoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgecoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgecoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgecoreSession struct {
	Contract     *Bridgecore       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgecoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgecoreCallerSession struct {
	Contract *BridgecoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BridgecoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgecoreTransactorSession struct {
	Contract     *BridgecoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BridgecoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgecoreRaw struct {
	Contract *Bridgecore // Generic contract binding to access the raw methods on
}

// BridgecoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgecoreCallerRaw struct {
	Contract *BridgecoreCaller // Generic read-only contract binding to access the raw methods on
}

// BridgecoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgecoreTransactorRaw struct {
	Contract *BridgecoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgecore creates a new instance of Bridgecore, bound to a specific deployed contract.
func NewBridgecore(address common.Address, backend bind.ContractBackend) (*Bridgecore, error) {
	contract, err := bindBridgecore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridgecore{BridgecoreCaller: BridgecoreCaller{contract: contract}, BridgecoreTransactor: BridgecoreTransactor{contract: contract}, BridgecoreFilterer: BridgecoreFilterer{contract: contract}}, nil
}

// NewBridgecoreCaller creates a new read-only instance of Bridgecore, bound to a specific deployed contract.
func NewBridgecoreCaller(address common.Address, caller bind.ContractCaller) (*BridgecoreCaller, error) {
	contract, err := bindBridgecore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgecoreCaller{contract: contract}, nil
}

// NewBridgecoreTransactor creates a new write-only instance of Bridgecore, bound to a specific deployed contract.
func NewBridgecoreTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgecoreTransactor, error) {
	contract, err := bindBridgecore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgecoreTransactor{contract: contract}, nil
}

// NewBridgecoreFilterer creates a new log filterer instance of Bridgecore, bound to a specific deployed contract.
func NewBridgecoreFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgecoreFilterer, error) {
	contract, err := bindBridgecore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgecoreFilterer{contract: contract}, nil
}

// bindBridgecore binds a generic wrapper to an already deployed contract.
func bindBridgecore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgecoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgecore *BridgecoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgecore.Contract.BridgecoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgecore *BridgecoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgecore.Contract.BridgecoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgecore *BridgecoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgecore.Contract.BridgecoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgecore *BridgecoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgecore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgecore *BridgecoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgecore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgecore *BridgecoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgecore.Contract.contract.Transact(opts, method, params...)
}

// PTMINT is a free data retrieval call binding the contract method 0x46f6f9b5.
//
// Solidity: function PT_MINT() view returns(uint8)
func (_Bridgecore *BridgecoreCaller) PTMINT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "PT_MINT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PTMINT is a free data retrieval call binding the contract method 0x46f6f9b5.
//
// Solidity: function PT_MINT() view returns(uint8)
func (_Bridgecore *BridgecoreSession) PTMINT() (uint8, error) {
	return _Bridgecore.Contract.PTMINT(&_Bridgecore.CallOpts)
}

// PTMINT is a free data retrieval call binding the contract method 0x46f6f9b5.
//
// Solidity: function PT_MINT() view returns(uint8)
func (_Bridgecore *BridgecoreCallerSession) PTMINT() (uint8, error) {
	return _Bridgecore.Contract.PTMINT(&_Bridgecore.CallOpts)
}

// PTUNLOCK is a free data retrieval call binding the contract method 0x68ea28b0.
//
// Solidity: function PT_UNLOCK() view returns(uint8)
func (_Bridgecore *BridgecoreCaller) PTUNLOCK(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "PT_UNLOCK")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PTUNLOCK is a free data retrieval call binding the contract method 0x68ea28b0.
//
// Solidity: function PT_UNLOCK() view returns(uint8)
func (_Bridgecore *BridgecoreSession) PTUNLOCK() (uint8, error) {
	return _Bridgecore.Contract.PTUNLOCK(&_Bridgecore.CallOpts)
}

// PTUNLOCK is a free data retrieval call binding the contract method 0x68ea28b0.
//
// Solidity: function PT_UNLOCK() view returns(uint8)
func (_Bridgecore *BridgecoreCallerSession) PTUNLOCK() (uint8, error) {
	return _Bridgecore.Contract.PTUNLOCK(&_Bridgecore.CallOpts)
}

// TOTALBPS is a free data retrieval call binding the contract method 0xd3cd52bc.
//
// Solidity: function TOTAL_BPS() view returns(uint16)
func (_Bridgecore *BridgecoreCaller) TOTALBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "TOTAL_BPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TOTALBPS is a free data retrieval call binding the contract method 0xd3cd52bc.
//
// Solidity: function TOTAL_BPS() view returns(uint16)
func (_Bridgecore *BridgecoreSession) TOTALBPS() (uint16, error) {
	return _Bridgecore.Contract.TOTALBPS(&_Bridgecore.CallOpts)
}

// TOTALBPS is a free data retrieval call binding the contract method 0xd3cd52bc.
//
// Solidity: function TOTAL_BPS() view returns(uint16)
func (_Bridgecore *BridgecoreCallerSession) TOTALBPS() (uint16, error) {
	return _Bridgecore.Contract.TOTALBPS(&_Bridgecore.CallOpts)
}

// EstimateBridgeFee is a free data retrieval call binding the contract method 0x94723256.
//
// Solidity: function estimateBridgeFee(uint16 remoteChainId, bool useZro, bytes adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Bridgecore *BridgecoreCaller) EstimateBridgeFee(opts *bind.CallOpts, remoteChainId uint16, useZro bool, adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "estimateBridgeFee", remoteChainId, useZro, adapterParams)

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

// EstimateBridgeFee is a free data retrieval call binding the contract method 0x94723256.
//
// Solidity: function estimateBridgeFee(uint16 remoteChainId, bool useZro, bytes adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Bridgecore *BridgecoreSession) EstimateBridgeFee(remoteChainId uint16, useZro bool, adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Bridgecore.Contract.EstimateBridgeFee(&_Bridgecore.CallOpts, remoteChainId, useZro, adapterParams)
}

// EstimateBridgeFee is a free data retrieval call binding the contract method 0x94723256.
//
// Solidity: function estimateBridgeFee(uint16 remoteChainId, bool useZro, bytes adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Bridgecore *BridgecoreCallerSession) EstimateBridgeFee(remoteChainId uint16, useZro bool, adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Bridgecore.Contract.EstimateBridgeFee(&_Bridgecore.CallOpts, remoteChainId, useZro, adapterParams)
}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Bridgecore *BridgecoreCaller) FailedMessages(opts *bind.CallOpts, arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "failedMessages", arg0, arg1, arg2)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Bridgecore *BridgecoreSession) FailedMessages(arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	return _Bridgecore.Contract.FailedMessages(&_Bridgecore.CallOpts, arg0, arg1, arg2)
}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Bridgecore *BridgecoreCallerSession) FailedMessages(arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	return _Bridgecore.Contract.FailedMessages(&_Bridgecore.CallOpts, arg0, arg1, arg2)
}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Bridgecore *BridgecoreCaller) GetConfig(opts *bind.CallOpts, _version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "getConfig", _version, _chainId, arg2, _configType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Bridgecore *BridgecoreSession) GetConfig(_version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	return _Bridgecore.Contract.GetConfig(&_Bridgecore.CallOpts, _version, _chainId, arg2, _configType)
}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Bridgecore *BridgecoreCallerSession) GetConfig(_version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	return _Bridgecore.Contract.GetConfig(&_Bridgecore.CallOpts, _version, _chainId, arg2, _configType)
}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Bridgecore *BridgecoreCaller) GetTrustedRemoteAddress(opts *bind.CallOpts, _remoteChainId uint16) ([]byte, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "getTrustedRemoteAddress", _remoteChainId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Bridgecore *BridgecoreSession) GetTrustedRemoteAddress(_remoteChainId uint16) ([]byte, error) {
	return _Bridgecore.Contract.GetTrustedRemoteAddress(&_Bridgecore.CallOpts, _remoteChainId)
}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Bridgecore *BridgecoreCallerSession) GetTrustedRemoteAddress(_remoteChainId uint16) ([]byte, error) {
	return _Bridgecore.Contract.GetTrustedRemoteAddress(&_Bridgecore.CallOpts, _remoteChainId)
}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Bridgecore *BridgecoreCaller) IsTrustedRemote(opts *bind.CallOpts, _srcChainId uint16, _srcAddress []byte) (bool, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "isTrustedRemote", _srcChainId, _srcAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Bridgecore *BridgecoreSession) IsTrustedRemote(_srcChainId uint16, _srcAddress []byte) (bool, error) {
	return _Bridgecore.Contract.IsTrustedRemote(&_Bridgecore.CallOpts, _srcChainId, _srcAddress)
}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Bridgecore *BridgecoreCallerSession) IsTrustedRemote(_srcChainId uint16, _srcAddress []byte) (bool, error) {
	return _Bridgecore.Contract.IsTrustedRemote(&_Bridgecore.CallOpts, _srcChainId, _srcAddress)
}

// LocalToRemote is a free data retrieval call binding the contract method 0x083f61fe.
//
// Solidity: function localToRemote(address , uint16 ) view returns(address)
func (_Bridgecore *BridgecoreCaller) LocalToRemote(opts *bind.CallOpts, arg0 common.Address, arg1 uint16) (common.Address, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "localToRemote", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LocalToRemote is a free data retrieval call binding the contract method 0x083f61fe.
//
// Solidity: function localToRemote(address , uint16 ) view returns(address)
func (_Bridgecore *BridgecoreSession) LocalToRemote(arg0 common.Address, arg1 uint16) (common.Address, error) {
	return _Bridgecore.Contract.LocalToRemote(&_Bridgecore.CallOpts, arg0, arg1)
}

// LocalToRemote is a free data retrieval call binding the contract method 0x083f61fe.
//
// Solidity: function localToRemote(address , uint16 ) view returns(address)
func (_Bridgecore *BridgecoreCallerSession) LocalToRemote(arg0 common.Address, arg1 uint16) (common.Address, error) {
	return _Bridgecore.Contract.LocalToRemote(&_Bridgecore.CallOpts, arg0, arg1)
}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Bridgecore *BridgecoreCaller) LzEndpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "lzEndpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Bridgecore *BridgecoreSession) LzEndpoint() (common.Address, error) {
	return _Bridgecore.Contract.LzEndpoint(&_Bridgecore.CallOpts)
}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Bridgecore *BridgecoreCallerSession) LzEndpoint() (common.Address, error) {
	return _Bridgecore.Contract.LzEndpoint(&_Bridgecore.CallOpts)
}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Bridgecore *BridgecoreCaller) MinDstGasLookup(opts *bind.CallOpts, arg0 uint16, arg1 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "minDstGasLookup", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Bridgecore *BridgecoreSession) MinDstGasLookup(arg0 uint16, arg1 uint16) (*big.Int, error) {
	return _Bridgecore.Contract.MinDstGasLookup(&_Bridgecore.CallOpts, arg0, arg1)
}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Bridgecore *BridgecoreCallerSession) MinDstGasLookup(arg0 uint16, arg1 uint16) (*big.Int, error) {
	return _Bridgecore.Contract.MinDstGasLookup(&_Bridgecore.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridgecore *BridgecoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridgecore *BridgecoreSession) Owner() (common.Address, error) {
	return _Bridgecore.Contract.Owner(&_Bridgecore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridgecore *BridgecoreCallerSession) Owner() (common.Address, error) {
	return _Bridgecore.Contract.Owner(&_Bridgecore.CallOpts)
}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Bridgecore *BridgecoreCaller) Precrime(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "precrime")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Bridgecore *BridgecoreSession) Precrime() (common.Address, error) {
	return _Bridgecore.Contract.Precrime(&_Bridgecore.CallOpts)
}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Bridgecore *BridgecoreCallerSession) Precrime() (common.Address, error) {
	return _Bridgecore.Contract.Precrime(&_Bridgecore.CallOpts)
}

// RemoteToLocal is a free data retrieval call binding the contract method 0x9b51251e.
//
// Solidity: function remoteToLocal(address , uint16 ) view returns(address)
func (_Bridgecore *BridgecoreCaller) RemoteToLocal(opts *bind.CallOpts, arg0 common.Address, arg1 uint16) (common.Address, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "remoteToLocal", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemoteToLocal is a free data retrieval call binding the contract method 0x9b51251e.
//
// Solidity: function remoteToLocal(address , uint16 ) view returns(address)
func (_Bridgecore *BridgecoreSession) RemoteToLocal(arg0 common.Address, arg1 uint16) (common.Address, error) {
	return _Bridgecore.Contract.RemoteToLocal(&_Bridgecore.CallOpts, arg0, arg1)
}

// RemoteToLocal is a free data retrieval call binding the contract method 0x9b51251e.
//
// Solidity: function remoteToLocal(address , uint16 ) view returns(address)
func (_Bridgecore *BridgecoreCallerSession) RemoteToLocal(arg0 common.Address, arg1 uint16) (common.Address, error) {
	return _Bridgecore.Contract.RemoteToLocal(&_Bridgecore.CallOpts, arg0, arg1)
}

// TotalValueLocked is a free data retrieval call binding the contract method 0x9a973587.
//
// Solidity: function totalValueLocked(uint16 , address ) view returns(uint256)
func (_Bridgecore *BridgecoreCaller) TotalValueLocked(opts *bind.CallOpts, arg0 uint16, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "totalValueLocked", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalValueLocked is a free data retrieval call binding the contract method 0x9a973587.
//
// Solidity: function totalValueLocked(uint16 , address ) view returns(uint256)
func (_Bridgecore *BridgecoreSession) TotalValueLocked(arg0 uint16, arg1 common.Address) (*big.Int, error) {
	return _Bridgecore.Contract.TotalValueLocked(&_Bridgecore.CallOpts, arg0, arg1)
}

// TotalValueLocked is a free data retrieval call binding the contract method 0x9a973587.
//
// Solidity: function totalValueLocked(uint16 , address ) view returns(uint256)
func (_Bridgecore *BridgecoreCallerSession) TotalValueLocked(arg0 uint16, arg1 common.Address) (*big.Int, error) {
	return _Bridgecore.Contract.TotalValueLocked(&_Bridgecore.CallOpts, arg0, arg1)
}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Bridgecore *BridgecoreCaller) TrustedRemoteLookup(opts *bind.CallOpts, arg0 uint16) ([]byte, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "trustedRemoteLookup", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Bridgecore *BridgecoreSession) TrustedRemoteLookup(arg0 uint16) ([]byte, error) {
	return _Bridgecore.Contract.TrustedRemoteLookup(&_Bridgecore.CallOpts, arg0)
}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Bridgecore *BridgecoreCallerSession) TrustedRemoteLookup(arg0 uint16) ([]byte, error) {
	return _Bridgecore.Contract.TrustedRemoteLookup(&_Bridgecore.CallOpts, arg0)
}

// UseCustomAdapterParams is a free data retrieval call binding the contract method 0xed629c5c.
//
// Solidity: function useCustomAdapterParams() view returns(bool)
func (_Bridgecore *BridgecoreCaller) UseCustomAdapterParams(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "useCustomAdapterParams")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseCustomAdapterParams is a free data retrieval call binding the contract method 0xed629c5c.
//
// Solidity: function useCustomAdapterParams() view returns(bool)
func (_Bridgecore *BridgecoreSession) UseCustomAdapterParams() (bool, error) {
	return _Bridgecore.Contract.UseCustomAdapterParams(&_Bridgecore.CallOpts)
}

// UseCustomAdapterParams is a free data retrieval call binding the contract method 0xed629c5c.
//
// Solidity: function useCustomAdapterParams() view returns(bool)
func (_Bridgecore *BridgecoreCallerSession) UseCustomAdapterParams() (bool, error) {
	return _Bridgecore.Contract.UseCustomAdapterParams(&_Bridgecore.CallOpts)
}

// WithdrawalFeeBps is a free data retrieval call binding the contract method 0x04336bb3.
//
// Solidity: function withdrawalFeeBps() view returns(uint16)
func (_Bridgecore *BridgecoreCaller) WithdrawalFeeBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Bridgecore.contract.Call(opts, &out, "withdrawalFeeBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// WithdrawalFeeBps is a free data retrieval call binding the contract method 0x04336bb3.
//
// Solidity: function withdrawalFeeBps() view returns(uint16)
func (_Bridgecore *BridgecoreSession) WithdrawalFeeBps() (uint16, error) {
	return _Bridgecore.Contract.WithdrawalFeeBps(&_Bridgecore.CallOpts)
}

// WithdrawalFeeBps is a free data retrieval call binding the contract method 0x04336bb3.
//
// Solidity: function withdrawalFeeBps() view returns(uint16)
func (_Bridgecore *BridgecoreCallerSession) WithdrawalFeeBps() (uint16, error) {
	return _Bridgecore.Contract.WithdrawalFeeBps(&_Bridgecore.CallOpts)
}

// Bridge is a paid mutator transaction binding the contract method 0x879762e2.
//
// Solidity: function bridge(address localToken, uint16 remoteChainId, uint256 amount, address to, bool unwrapWeth, (address,address) callParams, bytes adapterParams) payable returns()
func (_Bridgecore *BridgecoreTransactor) Bridge(opts *bind.TransactOpts, localToken common.Address, remoteChainId uint16, amount *big.Int, to common.Address, unwrapWeth bool, callParams LzLibCallParams, adapterParams []byte) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "bridge", localToken, remoteChainId, amount, to, unwrapWeth, callParams, adapterParams)
}

// Bridge is a paid mutator transaction binding the contract method 0x879762e2.
//
// Solidity: function bridge(address localToken, uint16 remoteChainId, uint256 amount, address to, bool unwrapWeth, (address,address) callParams, bytes adapterParams) payable returns()
func (_Bridgecore *BridgecoreSession) Bridge(localToken common.Address, remoteChainId uint16, amount *big.Int, to common.Address, unwrapWeth bool, callParams LzLibCallParams, adapterParams []byte) (*types.Transaction, error) {
	return _Bridgecore.Contract.Bridge(&_Bridgecore.TransactOpts, localToken, remoteChainId, amount, to, unwrapWeth, callParams, adapterParams)
}

// Bridge is a paid mutator transaction binding the contract method 0x879762e2.
//
// Solidity: function bridge(address localToken, uint16 remoteChainId, uint256 amount, address to, bool unwrapWeth, (address,address) callParams, bytes adapterParams) payable returns()
func (_Bridgecore *BridgecoreTransactorSession) Bridge(localToken common.Address, remoteChainId uint16, amount *big.Int, to common.Address, unwrapWeth bool, callParams LzLibCallParams, adapterParams []byte) (*types.Transaction, error) {
	return _Bridgecore.Contract.Bridge(&_Bridgecore.TransactOpts, localToken, remoteChainId, amount, to, unwrapWeth, callParams, adapterParams)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Bridgecore *BridgecoreTransactor) ForceResumeReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "forceResumeReceive", _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Bridgecore *BridgecoreSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Bridgecore.Contract.ForceResumeReceive(&_Bridgecore.TransactOpts, _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Bridgecore *BridgecoreTransactorSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Bridgecore.Contract.ForceResumeReceive(&_Bridgecore.TransactOpts, _srcChainId, _srcAddress)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Bridgecore *BridgecoreTransactor) LzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "lzReceive", _srcChainId, _srcAddress, _nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Bridgecore *BridgecoreSession) LzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgecore.Contract.LzReceive(&_Bridgecore.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Bridgecore *BridgecoreTransactorSession) LzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgecore.Contract.LzReceive(&_Bridgecore.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Bridgecore *BridgecoreTransactor) NonblockingLzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "nonblockingLzReceive", _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Bridgecore *BridgecoreSession) NonblockingLzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgecore.Contract.NonblockingLzReceive(&_Bridgecore.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Bridgecore *BridgecoreTransactorSession) NonblockingLzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgecore.Contract.NonblockingLzReceive(&_Bridgecore.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xb7cb1f7f.
//
// Solidity: function registerToken(address localToken, uint16 remoteChainId, address remoteToken) returns()
func (_Bridgecore *BridgecoreTransactor) RegisterToken(opts *bind.TransactOpts, localToken common.Address, remoteChainId uint16, remoteToken common.Address) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "registerToken", localToken, remoteChainId, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xb7cb1f7f.
//
// Solidity: function registerToken(address localToken, uint16 remoteChainId, address remoteToken) returns()
func (_Bridgecore *BridgecoreSession) RegisterToken(localToken common.Address, remoteChainId uint16, remoteToken common.Address) (*types.Transaction, error) {
	return _Bridgecore.Contract.RegisterToken(&_Bridgecore.TransactOpts, localToken, remoteChainId, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xb7cb1f7f.
//
// Solidity: function registerToken(address localToken, uint16 remoteChainId, address remoteToken) returns()
func (_Bridgecore *BridgecoreTransactorSession) RegisterToken(localToken common.Address, remoteChainId uint16, remoteToken common.Address) (*types.Transaction, error) {
	return _Bridgecore.Contract.RegisterToken(&_Bridgecore.TransactOpts, localToken, remoteChainId, remoteToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridgecore *BridgecoreTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridgecore *BridgecoreSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridgecore.Contract.RenounceOwnership(&_Bridgecore.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridgecore *BridgecoreTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridgecore.Contract.RenounceOwnership(&_Bridgecore.TransactOpts)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Bridgecore *BridgecoreTransactor) RetryMessage(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "retryMessage", _srcChainId, _srcAddress, _nonce, _payload)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Bridgecore *BridgecoreSession) RetryMessage(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgecore.Contract.RetryMessage(&_Bridgecore.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Bridgecore *BridgecoreTransactorSession) RetryMessage(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgecore.Contract.RetryMessage(&_Bridgecore.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Bridgecore *BridgecoreTransactor) SetConfig(opts *bind.TransactOpts, _version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "setConfig", _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Bridgecore *BridgecoreSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetConfig(&_Bridgecore.TransactOpts, _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Bridgecore *BridgecoreTransactorSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetConfig(&_Bridgecore.TransactOpts, _version, _chainId, _configType, _config)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Bridgecore *BridgecoreTransactor) SetMinDstGas(opts *bind.TransactOpts, _dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "setMinDstGas", _dstChainId, _packetType, _minGas)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Bridgecore *BridgecoreSession) SetMinDstGas(_dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetMinDstGas(&_Bridgecore.TransactOpts, _dstChainId, _packetType, _minGas)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Bridgecore *BridgecoreTransactorSession) SetMinDstGas(_dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetMinDstGas(&_Bridgecore.TransactOpts, _dstChainId, _packetType, _minGas)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Bridgecore *BridgecoreTransactor) SetPrecrime(opts *bind.TransactOpts, _precrime common.Address) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "setPrecrime", _precrime)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Bridgecore *BridgecoreSession) SetPrecrime(_precrime common.Address) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetPrecrime(&_Bridgecore.TransactOpts, _precrime)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Bridgecore *BridgecoreTransactorSession) SetPrecrime(_precrime common.Address) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetPrecrime(&_Bridgecore.TransactOpts, _precrime)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Bridgecore *BridgecoreTransactor) SetReceiveVersion(opts *bind.TransactOpts, _version uint16) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "setReceiveVersion", _version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Bridgecore *BridgecoreSession) SetReceiveVersion(_version uint16) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetReceiveVersion(&_Bridgecore.TransactOpts, _version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Bridgecore *BridgecoreTransactorSession) SetReceiveVersion(_version uint16) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetReceiveVersion(&_Bridgecore.TransactOpts, _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Bridgecore *BridgecoreTransactor) SetSendVersion(opts *bind.TransactOpts, _version uint16) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "setSendVersion", _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Bridgecore *BridgecoreSession) SetSendVersion(_version uint16) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetSendVersion(&_Bridgecore.TransactOpts, _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Bridgecore *BridgecoreTransactorSession) SetSendVersion(_version uint16) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetSendVersion(&_Bridgecore.TransactOpts, _version)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _srcChainId, bytes _path) returns()
func (_Bridgecore *BridgecoreTransactor) SetTrustedRemote(opts *bind.TransactOpts, _srcChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "setTrustedRemote", _srcChainId, _path)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _srcChainId, bytes _path) returns()
func (_Bridgecore *BridgecoreSession) SetTrustedRemote(_srcChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetTrustedRemote(&_Bridgecore.TransactOpts, _srcChainId, _path)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _srcChainId, bytes _path) returns()
func (_Bridgecore *BridgecoreTransactorSession) SetTrustedRemote(_srcChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetTrustedRemote(&_Bridgecore.TransactOpts, _srcChainId, _path)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Bridgecore *BridgecoreTransactor) SetTrustedRemoteAddress(opts *bind.TransactOpts, _remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "setTrustedRemoteAddress", _remoteChainId, _remoteAddress)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Bridgecore *BridgecoreSession) SetTrustedRemoteAddress(_remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetTrustedRemoteAddress(&_Bridgecore.TransactOpts, _remoteChainId, _remoteAddress)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Bridgecore *BridgecoreTransactorSession) SetTrustedRemoteAddress(_remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetTrustedRemoteAddress(&_Bridgecore.TransactOpts, _remoteChainId, _remoteAddress)
}

// SetUseCustomAdapterParams is a paid mutator transaction binding the contract method 0xeab45d9c.
//
// Solidity: function setUseCustomAdapterParams(bool _useCustomAdapterParams) returns()
func (_Bridgecore *BridgecoreTransactor) SetUseCustomAdapterParams(opts *bind.TransactOpts, _useCustomAdapterParams bool) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "setUseCustomAdapterParams", _useCustomAdapterParams)
}

// SetUseCustomAdapterParams is a paid mutator transaction binding the contract method 0xeab45d9c.
//
// Solidity: function setUseCustomAdapterParams(bool _useCustomAdapterParams) returns()
func (_Bridgecore *BridgecoreSession) SetUseCustomAdapterParams(_useCustomAdapterParams bool) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetUseCustomAdapterParams(&_Bridgecore.TransactOpts, _useCustomAdapterParams)
}

// SetUseCustomAdapterParams is a paid mutator transaction binding the contract method 0xeab45d9c.
//
// Solidity: function setUseCustomAdapterParams(bool _useCustomAdapterParams) returns()
func (_Bridgecore *BridgecoreTransactorSession) SetUseCustomAdapterParams(_useCustomAdapterParams bool) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetUseCustomAdapterParams(&_Bridgecore.TransactOpts, _useCustomAdapterParams)
}

// SetWithdrawalFeeBps is a paid mutator transaction binding the contract method 0xe6e7dd00.
//
// Solidity: function setWithdrawalFeeBps(uint16 _withdrawalFeeBps) returns()
func (_Bridgecore *BridgecoreTransactor) SetWithdrawalFeeBps(opts *bind.TransactOpts, _withdrawalFeeBps uint16) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "setWithdrawalFeeBps", _withdrawalFeeBps)
}

// SetWithdrawalFeeBps is a paid mutator transaction binding the contract method 0xe6e7dd00.
//
// Solidity: function setWithdrawalFeeBps(uint16 _withdrawalFeeBps) returns()
func (_Bridgecore *BridgecoreSession) SetWithdrawalFeeBps(_withdrawalFeeBps uint16) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetWithdrawalFeeBps(&_Bridgecore.TransactOpts, _withdrawalFeeBps)
}

// SetWithdrawalFeeBps is a paid mutator transaction binding the contract method 0xe6e7dd00.
//
// Solidity: function setWithdrawalFeeBps(uint16 _withdrawalFeeBps) returns()
func (_Bridgecore *BridgecoreTransactorSession) SetWithdrawalFeeBps(_withdrawalFeeBps uint16) (*types.Transaction, error) {
	return _Bridgecore.Contract.SetWithdrawalFeeBps(&_Bridgecore.TransactOpts, _withdrawalFeeBps)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridgecore *BridgecoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridgecore.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridgecore *BridgecoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridgecore.Contract.TransferOwnership(&_Bridgecore.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridgecore *BridgecoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridgecore.Contract.TransferOwnership(&_Bridgecore.TransactOpts, newOwner)
}

// BridgecoreMessageFailedIterator is returned from FilterMessageFailed and is used to iterate over the raw logs and unpacked data for MessageFailed events raised by the Bridgecore contract.
type BridgecoreMessageFailedIterator struct {
	Event *BridgecoreMessageFailed // Event containing the contract specifics and raw log

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
func (it *BridgecoreMessageFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgecoreMessageFailed)
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
		it.Event = new(BridgecoreMessageFailed)
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
func (it *BridgecoreMessageFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgecoreMessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgecoreMessageFailed represents a MessageFailed event raised by the Bridgecore contract.
type BridgecoreMessageFailed struct {
	SrcChainId uint16
	SrcAddress []byte
	Nonce      uint64
	Payload    []byte
	Reason     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageFailed is a free log retrieval operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_Bridgecore *BridgecoreFilterer) FilterMessageFailed(opts *bind.FilterOpts) (*BridgecoreMessageFailedIterator, error) {

	logs, sub, err := _Bridgecore.contract.FilterLogs(opts, "MessageFailed")
	if err != nil {
		return nil, err
	}
	return &BridgecoreMessageFailedIterator{contract: _Bridgecore.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

// WatchMessageFailed is a free log subscription operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_Bridgecore *BridgecoreFilterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *BridgecoreMessageFailed) (event.Subscription, error) {

	logs, sub, err := _Bridgecore.contract.WatchLogs(opts, "MessageFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgecoreMessageFailed)
				if err := _Bridgecore.contract.UnpackLog(event, "MessageFailed", log); err != nil {
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

// ParseMessageFailed is a log parse operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_Bridgecore *BridgecoreFilterer) ParseMessageFailed(log types.Log) (*BridgecoreMessageFailed, error) {
	event := new(BridgecoreMessageFailed)
	if err := _Bridgecore.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgecoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridgecore contract.
type BridgecoreOwnershipTransferredIterator struct {
	Event *BridgecoreOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgecoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgecoreOwnershipTransferred)
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
		it.Event = new(BridgecoreOwnershipTransferred)
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
func (it *BridgecoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgecoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgecoreOwnershipTransferred represents a OwnershipTransferred event raised by the Bridgecore contract.
type BridgecoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridgecore *BridgecoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgecoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridgecore.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgecoreOwnershipTransferredIterator{contract: _Bridgecore.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridgecore *BridgecoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgecoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridgecore.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgecoreOwnershipTransferred)
				if err := _Bridgecore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bridgecore *BridgecoreFilterer) ParseOwnershipTransferred(log types.Log) (*BridgecoreOwnershipTransferred, error) {
	event := new(BridgecoreOwnershipTransferred)
	if err := _Bridgecore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgecoreRegisterTokenIterator is returned from FilterRegisterToken and is used to iterate over the raw logs and unpacked data for RegisterToken events raised by the Bridgecore contract.
type BridgecoreRegisterTokenIterator struct {
	Event *BridgecoreRegisterToken // Event containing the contract specifics and raw log

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
func (it *BridgecoreRegisterTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgecoreRegisterToken)
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
		it.Event = new(BridgecoreRegisterToken)
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
func (it *BridgecoreRegisterTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgecoreRegisterTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgecoreRegisterToken represents a RegisterToken event raised by the Bridgecore contract.
type BridgecoreRegisterToken struct {
	LocalToken    common.Address
	RemoteChainId uint16
	RemoteToken   common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRegisterToken is a free log retrieval operation binding the contract event 0x45e419139d22e1853a154f058e7b60373892104bf13079aa3a156a73f84f637d.
//
// Solidity: event RegisterToken(address localToken, uint16 remoteChainId, address remoteToken)
func (_Bridgecore *BridgecoreFilterer) FilterRegisterToken(opts *bind.FilterOpts) (*BridgecoreRegisterTokenIterator, error) {

	logs, sub, err := _Bridgecore.contract.FilterLogs(opts, "RegisterToken")
	if err != nil {
		return nil, err
	}
	return &BridgecoreRegisterTokenIterator{contract: _Bridgecore.contract, event: "RegisterToken", logs: logs, sub: sub}, nil
}

// WatchRegisterToken is a free log subscription operation binding the contract event 0x45e419139d22e1853a154f058e7b60373892104bf13079aa3a156a73f84f637d.
//
// Solidity: event RegisterToken(address localToken, uint16 remoteChainId, address remoteToken)
func (_Bridgecore *BridgecoreFilterer) WatchRegisterToken(opts *bind.WatchOpts, sink chan<- *BridgecoreRegisterToken) (event.Subscription, error) {

	logs, sub, err := _Bridgecore.contract.WatchLogs(opts, "RegisterToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgecoreRegisterToken)
				if err := _Bridgecore.contract.UnpackLog(event, "RegisterToken", log); err != nil {
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

// ParseRegisterToken is a log parse operation binding the contract event 0x45e419139d22e1853a154f058e7b60373892104bf13079aa3a156a73f84f637d.
//
// Solidity: event RegisterToken(address localToken, uint16 remoteChainId, address remoteToken)
func (_Bridgecore *BridgecoreFilterer) ParseRegisterToken(log types.Log) (*BridgecoreRegisterToken, error) {
	event := new(BridgecoreRegisterToken)
	if err := _Bridgecore.contract.UnpackLog(event, "RegisterToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgecoreRetryMessageSuccessIterator is returned from FilterRetryMessageSuccess and is used to iterate over the raw logs and unpacked data for RetryMessageSuccess events raised by the Bridgecore contract.
type BridgecoreRetryMessageSuccessIterator struct {
	Event *BridgecoreRetryMessageSuccess // Event containing the contract specifics and raw log

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
func (it *BridgecoreRetryMessageSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgecoreRetryMessageSuccess)
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
		it.Event = new(BridgecoreRetryMessageSuccess)
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
func (it *BridgecoreRetryMessageSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgecoreRetryMessageSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgecoreRetryMessageSuccess represents a RetryMessageSuccess event raised by the Bridgecore contract.
type BridgecoreRetryMessageSuccess struct {
	SrcChainId  uint16
	SrcAddress  []byte
	Nonce       uint64
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRetryMessageSuccess is a free log retrieval operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_Bridgecore *BridgecoreFilterer) FilterRetryMessageSuccess(opts *bind.FilterOpts) (*BridgecoreRetryMessageSuccessIterator, error) {

	logs, sub, err := _Bridgecore.contract.FilterLogs(opts, "RetryMessageSuccess")
	if err != nil {
		return nil, err
	}
	return &BridgecoreRetryMessageSuccessIterator{contract: _Bridgecore.contract, event: "RetryMessageSuccess", logs: logs, sub: sub}, nil
}

// WatchRetryMessageSuccess is a free log subscription operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_Bridgecore *BridgecoreFilterer) WatchRetryMessageSuccess(opts *bind.WatchOpts, sink chan<- *BridgecoreRetryMessageSuccess) (event.Subscription, error) {

	logs, sub, err := _Bridgecore.contract.WatchLogs(opts, "RetryMessageSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgecoreRetryMessageSuccess)
				if err := _Bridgecore.contract.UnpackLog(event, "RetryMessageSuccess", log); err != nil {
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

// ParseRetryMessageSuccess is a log parse operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_Bridgecore *BridgecoreFilterer) ParseRetryMessageSuccess(log types.Log) (*BridgecoreRetryMessageSuccess, error) {
	event := new(BridgecoreRetryMessageSuccess)
	if err := _Bridgecore.contract.UnpackLog(event, "RetryMessageSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgecoreSetMinDstGasIterator is returned from FilterSetMinDstGas and is used to iterate over the raw logs and unpacked data for SetMinDstGas events raised by the Bridgecore contract.
type BridgecoreSetMinDstGasIterator struct {
	Event *BridgecoreSetMinDstGas // Event containing the contract specifics and raw log

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
func (it *BridgecoreSetMinDstGasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgecoreSetMinDstGas)
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
		it.Event = new(BridgecoreSetMinDstGas)
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
func (it *BridgecoreSetMinDstGasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgecoreSetMinDstGasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgecoreSetMinDstGas represents a SetMinDstGas event raised by the Bridgecore contract.
type BridgecoreSetMinDstGas struct {
	DstChainId uint16
	Type       uint16
	MinDstGas  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetMinDstGas is a free log retrieval operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_Bridgecore *BridgecoreFilterer) FilterSetMinDstGas(opts *bind.FilterOpts) (*BridgecoreSetMinDstGasIterator, error) {

	logs, sub, err := _Bridgecore.contract.FilterLogs(opts, "SetMinDstGas")
	if err != nil {
		return nil, err
	}
	return &BridgecoreSetMinDstGasIterator{contract: _Bridgecore.contract, event: "SetMinDstGas", logs: logs, sub: sub}, nil
}

// WatchSetMinDstGas is a free log subscription operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_Bridgecore *BridgecoreFilterer) WatchSetMinDstGas(opts *bind.WatchOpts, sink chan<- *BridgecoreSetMinDstGas) (event.Subscription, error) {

	logs, sub, err := _Bridgecore.contract.WatchLogs(opts, "SetMinDstGas")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgecoreSetMinDstGas)
				if err := _Bridgecore.contract.UnpackLog(event, "SetMinDstGas", log); err != nil {
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

// ParseSetMinDstGas is a log parse operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_Bridgecore *BridgecoreFilterer) ParseSetMinDstGas(log types.Log) (*BridgecoreSetMinDstGas, error) {
	event := new(BridgecoreSetMinDstGas)
	if err := _Bridgecore.contract.UnpackLog(event, "SetMinDstGas", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgecoreSetPrecrimeIterator is returned from FilterSetPrecrime and is used to iterate over the raw logs and unpacked data for SetPrecrime events raised by the Bridgecore contract.
type BridgecoreSetPrecrimeIterator struct {
	Event *BridgecoreSetPrecrime // Event containing the contract specifics and raw log

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
func (it *BridgecoreSetPrecrimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgecoreSetPrecrime)
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
		it.Event = new(BridgecoreSetPrecrime)
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
func (it *BridgecoreSetPrecrimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgecoreSetPrecrimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgecoreSetPrecrime represents a SetPrecrime event raised by the Bridgecore contract.
type BridgecoreSetPrecrime struct {
	Precrime common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPrecrime is a free log retrieval operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_Bridgecore *BridgecoreFilterer) FilterSetPrecrime(opts *bind.FilterOpts) (*BridgecoreSetPrecrimeIterator, error) {

	logs, sub, err := _Bridgecore.contract.FilterLogs(opts, "SetPrecrime")
	if err != nil {
		return nil, err
	}
	return &BridgecoreSetPrecrimeIterator{contract: _Bridgecore.contract, event: "SetPrecrime", logs: logs, sub: sub}, nil
}

// WatchSetPrecrime is a free log subscription operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_Bridgecore *BridgecoreFilterer) WatchSetPrecrime(opts *bind.WatchOpts, sink chan<- *BridgecoreSetPrecrime) (event.Subscription, error) {

	logs, sub, err := _Bridgecore.contract.WatchLogs(opts, "SetPrecrime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgecoreSetPrecrime)
				if err := _Bridgecore.contract.UnpackLog(event, "SetPrecrime", log); err != nil {
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

// ParseSetPrecrime is a log parse operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_Bridgecore *BridgecoreFilterer) ParseSetPrecrime(log types.Log) (*BridgecoreSetPrecrime, error) {
	event := new(BridgecoreSetPrecrime)
	if err := _Bridgecore.contract.UnpackLog(event, "SetPrecrime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgecoreSetTrustedRemoteIterator is returned from FilterSetTrustedRemote and is used to iterate over the raw logs and unpacked data for SetTrustedRemote events raised by the Bridgecore contract.
type BridgecoreSetTrustedRemoteIterator struct {
	Event *BridgecoreSetTrustedRemote // Event containing the contract specifics and raw log

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
func (it *BridgecoreSetTrustedRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgecoreSetTrustedRemote)
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
		it.Event = new(BridgecoreSetTrustedRemote)
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
func (it *BridgecoreSetTrustedRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgecoreSetTrustedRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgecoreSetTrustedRemote represents a SetTrustedRemote event raised by the Bridgecore contract.
type BridgecoreSetTrustedRemote struct {
	RemoteChainId uint16
	Path          []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemote is a free log retrieval operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_Bridgecore *BridgecoreFilterer) FilterSetTrustedRemote(opts *bind.FilterOpts) (*BridgecoreSetTrustedRemoteIterator, error) {

	logs, sub, err := _Bridgecore.contract.FilterLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return &BridgecoreSetTrustedRemoteIterator{contract: _Bridgecore.contract, event: "SetTrustedRemote", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemote is a free log subscription operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_Bridgecore *BridgecoreFilterer) WatchSetTrustedRemote(opts *bind.WatchOpts, sink chan<- *BridgecoreSetTrustedRemote) (event.Subscription, error) {

	logs, sub, err := _Bridgecore.contract.WatchLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgecoreSetTrustedRemote)
				if err := _Bridgecore.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
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

// ParseSetTrustedRemote is a log parse operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_Bridgecore *BridgecoreFilterer) ParseSetTrustedRemote(log types.Log) (*BridgecoreSetTrustedRemote, error) {
	event := new(BridgecoreSetTrustedRemote)
	if err := _Bridgecore.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgecoreSetTrustedRemoteAddressIterator is returned from FilterSetTrustedRemoteAddress and is used to iterate over the raw logs and unpacked data for SetTrustedRemoteAddress events raised by the Bridgecore contract.
type BridgecoreSetTrustedRemoteAddressIterator struct {
	Event *BridgecoreSetTrustedRemoteAddress // Event containing the contract specifics and raw log

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
func (it *BridgecoreSetTrustedRemoteAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgecoreSetTrustedRemoteAddress)
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
		it.Event = new(BridgecoreSetTrustedRemoteAddress)
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
func (it *BridgecoreSetTrustedRemoteAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgecoreSetTrustedRemoteAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgecoreSetTrustedRemoteAddress represents a SetTrustedRemoteAddress event raised by the Bridgecore contract.
type BridgecoreSetTrustedRemoteAddress struct {
	RemoteChainId uint16
	RemoteAddress []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemoteAddress is a free log retrieval operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_Bridgecore *BridgecoreFilterer) FilterSetTrustedRemoteAddress(opts *bind.FilterOpts) (*BridgecoreSetTrustedRemoteAddressIterator, error) {

	logs, sub, err := _Bridgecore.contract.FilterLogs(opts, "SetTrustedRemoteAddress")
	if err != nil {
		return nil, err
	}
	return &BridgecoreSetTrustedRemoteAddressIterator{contract: _Bridgecore.contract, event: "SetTrustedRemoteAddress", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemoteAddress is a free log subscription operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_Bridgecore *BridgecoreFilterer) WatchSetTrustedRemoteAddress(opts *bind.WatchOpts, sink chan<- *BridgecoreSetTrustedRemoteAddress) (event.Subscription, error) {

	logs, sub, err := _Bridgecore.contract.WatchLogs(opts, "SetTrustedRemoteAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgecoreSetTrustedRemoteAddress)
				if err := _Bridgecore.contract.UnpackLog(event, "SetTrustedRemoteAddress", log); err != nil {
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

// ParseSetTrustedRemoteAddress is a log parse operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_Bridgecore *BridgecoreFilterer) ParseSetTrustedRemoteAddress(log types.Log) (*BridgecoreSetTrustedRemoteAddress, error) {
	event := new(BridgecoreSetTrustedRemoteAddress)
	if err := _Bridgecore.contract.UnpackLog(event, "SetTrustedRemoteAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgecoreSetUseCustomAdapterParamsIterator is returned from FilterSetUseCustomAdapterParams and is used to iterate over the raw logs and unpacked data for SetUseCustomAdapterParams events raised by the Bridgecore contract.
type BridgecoreSetUseCustomAdapterParamsIterator struct {
	Event *BridgecoreSetUseCustomAdapterParams // Event containing the contract specifics and raw log

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
func (it *BridgecoreSetUseCustomAdapterParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgecoreSetUseCustomAdapterParams)
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
		it.Event = new(BridgecoreSetUseCustomAdapterParams)
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
func (it *BridgecoreSetUseCustomAdapterParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgecoreSetUseCustomAdapterParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgecoreSetUseCustomAdapterParams represents a SetUseCustomAdapterParams event raised by the Bridgecore contract.
type BridgecoreSetUseCustomAdapterParams struct {
	UseCustomAdapterParams bool
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetUseCustomAdapterParams is a free log retrieval operation binding the contract event 0x1584ad594a70cbe1e6515592e1272a987d922b097ead875069cebe8b40c004a4.
//
// Solidity: event SetUseCustomAdapterParams(bool useCustomAdapterParams)
func (_Bridgecore *BridgecoreFilterer) FilterSetUseCustomAdapterParams(opts *bind.FilterOpts) (*BridgecoreSetUseCustomAdapterParamsIterator, error) {

	logs, sub, err := _Bridgecore.contract.FilterLogs(opts, "SetUseCustomAdapterParams")
	if err != nil {
		return nil, err
	}
	return &BridgecoreSetUseCustomAdapterParamsIterator{contract: _Bridgecore.contract, event: "SetUseCustomAdapterParams", logs: logs, sub: sub}, nil
}

// WatchSetUseCustomAdapterParams is a free log subscription operation binding the contract event 0x1584ad594a70cbe1e6515592e1272a987d922b097ead875069cebe8b40c004a4.
//
// Solidity: event SetUseCustomAdapterParams(bool useCustomAdapterParams)
func (_Bridgecore *BridgecoreFilterer) WatchSetUseCustomAdapterParams(opts *bind.WatchOpts, sink chan<- *BridgecoreSetUseCustomAdapterParams) (event.Subscription, error) {

	logs, sub, err := _Bridgecore.contract.WatchLogs(opts, "SetUseCustomAdapterParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgecoreSetUseCustomAdapterParams)
				if err := _Bridgecore.contract.UnpackLog(event, "SetUseCustomAdapterParams", log); err != nil {
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

// ParseSetUseCustomAdapterParams is a log parse operation binding the contract event 0x1584ad594a70cbe1e6515592e1272a987d922b097ead875069cebe8b40c004a4.
//
// Solidity: event SetUseCustomAdapterParams(bool useCustomAdapterParams)
func (_Bridgecore *BridgecoreFilterer) ParseSetUseCustomAdapterParams(log types.Log) (*BridgecoreSetUseCustomAdapterParams, error) {
	event := new(BridgecoreSetUseCustomAdapterParams)
	if err := _Bridgecore.contract.UnpackLog(event, "SetUseCustomAdapterParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgecoreSetWithdrawalFeeBpsIterator is returned from FilterSetWithdrawalFeeBps and is used to iterate over the raw logs and unpacked data for SetWithdrawalFeeBps events raised by the Bridgecore contract.
type BridgecoreSetWithdrawalFeeBpsIterator struct {
	Event *BridgecoreSetWithdrawalFeeBps // Event containing the contract specifics and raw log

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
func (it *BridgecoreSetWithdrawalFeeBpsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgecoreSetWithdrawalFeeBps)
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
		it.Event = new(BridgecoreSetWithdrawalFeeBps)
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
func (it *BridgecoreSetWithdrawalFeeBpsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgecoreSetWithdrawalFeeBpsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgecoreSetWithdrawalFeeBps represents a SetWithdrawalFeeBps event raised by the Bridgecore contract.
type BridgecoreSetWithdrawalFeeBps struct {
	WithdrawalFeeBps uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetWithdrawalFeeBps is a free log retrieval operation binding the contract event 0x6d0c3abfe7f8a420e34f51060ff6519e1b0f47249f7043e76abc83f61b9f99f7.
//
// Solidity: event SetWithdrawalFeeBps(uint16 withdrawalFeeBps)
func (_Bridgecore *BridgecoreFilterer) FilterSetWithdrawalFeeBps(opts *bind.FilterOpts) (*BridgecoreSetWithdrawalFeeBpsIterator, error) {

	logs, sub, err := _Bridgecore.contract.FilterLogs(opts, "SetWithdrawalFeeBps")
	if err != nil {
		return nil, err
	}
	return &BridgecoreSetWithdrawalFeeBpsIterator{contract: _Bridgecore.contract, event: "SetWithdrawalFeeBps", logs: logs, sub: sub}, nil
}

// WatchSetWithdrawalFeeBps is a free log subscription operation binding the contract event 0x6d0c3abfe7f8a420e34f51060ff6519e1b0f47249f7043e76abc83f61b9f99f7.
//
// Solidity: event SetWithdrawalFeeBps(uint16 withdrawalFeeBps)
func (_Bridgecore *BridgecoreFilterer) WatchSetWithdrawalFeeBps(opts *bind.WatchOpts, sink chan<- *BridgecoreSetWithdrawalFeeBps) (event.Subscription, error) {

	logs, sub, err := _Bridgecore.contract.WatchLogs(opts, "SetWithdrawalFeeBps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgecoreSetWithdrawalFeeBps)
				if err := _Bridgecore.contract.UnpackLog(event, "SetWithdrawalFeeBps", log); err != nil {
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

// ParseSetWithdrawalFeeBps is a log parse operation binding the contract event 0x6d0c3abfe7f8a420e34f51060ff6519e1b0f47249f7043e76abc83f61b9f99f7.
//
// Solidity: event SetWithdrawalFeeBps(uint16 withdrawalFeeBps)
func (_Bridgecore *BridgecoreFilterer) ParseSetWithdrawalFeeBps(log types.Log) (*BridgecoreSetWithdrawalFeeBps, error) {
	event := new(BridgecoreSetWithdrawalFeeBps)
	if err := _Bridgecore.contract.UnpackLog(event, "SetWithdrawalFeeBps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgecoreUnwrapTokenIterator is returned from FilterUnwrapToken and is used to iterate over the raw logs and unpacked data for UnwrapToken events raised by the Bridgecore contract.
type BridgecoreUnwrapTokenIterator struct {
	Event *BridgecoreUnwrapToken // Event containing the contract specifics and raw log

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
func (it *BridgecoreUnwrapTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgecoreUnwrapToken)
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
		it.Event = new(BridgecoreUnwrapToken)
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
func (it *BridgecoreUnwrapTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgecoreUnwrapTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgecoreUnwrapToken represents a UnwrapToken event raised by the Bridgecore contract.
type BridgecoreUnwrapToken struct {
	LocalToken    common.Address
	RemoteToken   common.Address
	RemoteChainId uint16
	To            common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnwrapToken is a free log retrieval operation binding the contract event 0x3b661011d9e0ff8f0dc432bac4ed79eabf70cf52596ed9de985810ef36689e9e.
//
// Solidity: event UnwrapToken(address localToken, address remoteToken, uint16 remoteChainId, address to, uint256 amount)
func (_Bridgecore *BridgecoreFilterer) FilterUnwrapToken(opts *bind.FilterOpts) (*BridgecoreUnwrapTokenIterator, error) {

	logs, sub, err := _Bridgecore.contract.FilterLogs(opts, "UnwrapToken")
	if err != nil {
		return nil, err
	}
	return &BridgecoreUnwrapTokenIterator{contract: _Bridgecore.contract, event: "UnwrapToken", logs: logs, sub: sub}, nil
}

// WatchUnwrapToken is a free log subscription operation binding the contract event 0x3b661011d9e0ff8f0dc432bac4ed79eabf70cf52596ed9de985810ef36689e9e.
//
// Solidity: event UnwrapToken(address localToken, address remoteToken, uint16 remoteChainId, address to, uint256 amount)
func (_Bridgecore *BridgecoreFilterer) WatchUnwrapToken(opts *bind.WatchOpts, sink chan<- *BridgecoreUnwrapToken) (event.Subscription, error) {

	logs, sub, err := _Bridgecore.contract.WatchLogs(opts, "UnwrapToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgecoreUnwrapToken)
				if err := _Bridgecore.contract.UnpackLog(event, "UnwrapToken", log); err != nil {
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

// ParseUnwrapToken is a log parse operation binding the contract event 0x3b661011d9e0ff8f0dc432bac4ed79eabf70cf52596ed9de985810ef36689e9e.
//
// Solidity: event UnwrapToken(address localToken, address remoteToken, uint16 remoteChainId, address to, uint256 amount)
func (_Bridgecore *BridgecoreFilterer) ParseUnwrapToken(log types.Log) (*BridgecoreUnwrapToken, error) {
	event := new(BridgecoreUnwrapToken)
	if err := _Bridgecore.contract.UnpackLog(event, "UnwrapToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgecoreWrapTokenIterator is returned from FilterWrapToken and is used to iterate over the raw logs and unpacked data for WrapToken events raised by the Bridgecore contract.
type BridgecoreWrapTokenIterator struct {
	Event *BridgecoreWrapToken // Event containing the contract specifics and raw log

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
func (it *BridgecoreWrapTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgecoreWrapToken)
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
		it.Event = new(BridgecoreWrapToken)
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
func (it *BridgecoreWrapTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgecoreWrapTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgecoreWrapToken represents a WrapToken event raised by the Bridgecore contract.
type BridgecoreWrapToken struct {
	LocalToken    common.Address
	RemoteToken   common.Address
	RemoteChainId uint16
	To            common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWrapToken is a free log retrieval operation binding the contract event 0xf865724e934515a839f76ebdd6a53df378816b384e8c10270332411676c48dd5.
//
// Solidity: event WrapToken(address localToken, address remoteToken, uint16 remoteChainId, address to, uint256 amount)
func (_Bridgecore *BridgecoreFilterer) FilterWrapToken(opts *bind.FilterOpts) (*BridgecoreWrapTokenIterator, error) {

	logs, sub, err := _Bridgecore.contract.FilterLogs(opts, "WrapToken")
	if err != nil {
		return nil, err
	}
	return &BridgecoreWrapTokenIterator{contract: _Bridgecore.contract, event: "WrapToken", logs: logs, sub: sub}, nil
}

// WatchWrapToken is a free log subscription operation binding the contract event 0xf865724e934515a839f76ebdd6a53df378816b384e8c10270332411676c48dd5.
//
// Solidity: event WrapToken(address localToken, address remoteToken, uint16 remoteChainId, address to, uint256 amount)
func (_Bridgecore *BridgecoreFilterer) WatchWrapToken(opts *bind.WatchOpts, sink chan<- *BridgecoreWrapToken) (event.Subscription, error) {

	logs, sub, err := _Bridgecore.contract.WatchLogs(opts, "WrapToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgecoreWrapToken)
				if err := _Bridgecore.contract.UnpackLog(event, "WrapToken", log); err != nil {
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

// ParseWrapToken is a log parse operation binding the contract event 0xf865724e934515a839f76ebdd6a53df378816b384e8c10270332411676c48dd5.
//
// Solidity: event WrapToken(address localToken, address remoteToken, uint16 remoteChainId, address to, uint256 amount)
func (_Bridgecore *BridgecoreFilterer) ParseWrapToken(log types.Log) (*BridgecoreWrapToken, error) {
	event := new(BridgecoreWrapToken)
	if err := _Bridgecore.contract.UnpackLog(event, "WrapToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
