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
type LzLibCallParams struct {
	RefundAddress     common.Address
	ZroPaymentAddress common.Address
}

// BridgebnbMetaData contains all meta data concerning the Bridgebnb contract.
var BridgebnbMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_endpoint\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_reason\",\"type\":\"bytes\"}],\"name\":\"MessageFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceiveToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegisterToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"RetryMessageSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SendToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_type\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minDstGas\",\"type\":\"uint256\"}],\"name\":\"SetMinDstGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"precrime\",\"type\":\"address\"}],\"name\":\"SetPrecrime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"remoteChainId\",\"type\":\"uint16\"}],\"name\":\"SetRemoteChainId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_path\",\"type\":\"bytes\"}],\"name\":\"SetTrustedRemote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_remoteAddress\",\"type\":\"bytes\"}],\"name\":\"SetTrustedRemoteAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useCustomAdapterParams\",\"type\":\"bool\"}],\"name\":\"SetUseCustomAdapterParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawFee\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"LDtoSDConversionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PT_MINT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PT_UNLOCK\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"accruedFeeLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zroPaymentAddress\",\"type\":\"address\"}],\"internalType\":\"structLzLib.CallParams\",\"name\":\"callParams\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"adapterParams\",\"type\":\"bytes\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zroPaymentAddress\",\"type\":\"address\"}],\"internalType\":\"structLzLib.CallParams\",\"name\":\"callParams\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"adapterParams\",\"type\":\"bytes\"}],\"name\":\"bridgeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"useZro\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"adapterParams\",\"type\":\"bytes\"}],\"name\":\"estimateBridgeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zroFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"failedMessages\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"}],\"name\":\"forceResumeReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_configType\",\"type\":\"uint256\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"}],\"name\":\"getTrustedRemoteAddress\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"}],\"name\":\"isTrustedRemote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lzEndpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"minDstGasLookup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"nonblockingLzReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precrime\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sharedDecimals\",\"type\":\"uint8\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remoteChainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"retryMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_configType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_config\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_packetType\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minGas\",\"type\":\"uint256\"}],\"name\":\"setMinDstGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_precrime\",\"type\":\"address\"}],\"name\":\"setPrecrime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"}],\"name\":\"setReceiveVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"}],\"name\":\"setRemoteChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"}],\"name\":\"setSendVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_path\",\"type\":\"bytes\"}],\"name\":\"setTrustedRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_remoteAddress\",\"type\":\"bytes\"}],\"name\":\"setTrustedRemoteAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_useCustomAdapterParams\",\"type\":\"bool\"}],\"name\":\"setUseCustomAdapterParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalValueLockedSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"trustedRemoteLookup\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useCustomAdapterParams\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BridgebnbABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgebnbMetaData.ABI instead.
var BridgebnbABI = BridgebnbMetaData.ABI

// Bridgebnb is an auto generated Go binding around an Ethereum contract.
type Bridgebnb struct {
	BridgebnbCaller     // Read-only binding to the contract
	BridgebnbTransactor // Write-only binding to the contract
	BridgebnbFilterer   // Log filterer for contract events
}

// BridgebnbCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgebnbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgebnbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgebnbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgebnbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgebnbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgebnbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgebnbSession struct {
	Contract     *Bridgebnb        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgebnbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgebnbCallerSession struct {
	Contract *BridgebnbCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BridgebnbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgebnbTransactorSession struct {
	Contract     *BridgebnbTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BridgebnbRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgebnbRaw struct {
	Contract *Bridgebnb // Generic contract binding to access the raw methods on
}

// BridgebnbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgebnbCallerRaw struct {
	Contract *BridgebnbCaller // Generic read-only contract binding to access the raw methods on
}

// BridgebnbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgebnbTransactorRaw struct {
	Contract *BridgebnbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgebnb creates a new instance of Bridgebnb, bound to a specific deployed contract.
func NewBridgebnb(address common.Address, backend bind.ContractBackend) (*Bridgebnb, error) {
	contract, err := bindBridgebnb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridgebnb{BridgebnbCaller: BridgebnbCaller{contract: contract}, BridgebnbTransactor: BridgebnbTransactor{contract: contract}, BridgebnbFilterer: BridgebnbFilterer{contract: contract}}, nil
}

// NewBridgebnbCaller creates a new read-only instance of Bridgebnb, bound to a specific deployed contract.
func NewBridgebnbCaller(address common.Address, caller bind.ContractCaller) (*BridgebnbCaller, error) {
	contract, err := bindBridgebnb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgebnbCaller{contract: contract}, nil
}

// NewBridgebnbTransactor creates a new write-only instance of Bridgebnb, bound to a specific deployed contract.
func NewBridgebnbTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgebnbTransactor, error) {
	contract, err := bindBridgebnb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgebnbTransactor{contract: contract}, nil
}

// NewBridgebnbFilterer creates a new log filterer instance of Bridgebnb, bound to a specific deployed contract.
func NewBridgebnbFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgebnbFilterer, error) {
	contract, err := bindBridgebnb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgebnbFilterer{contract: contract}, nil
}

// bindBridgebnb binds a generic wrapper to an already deployed contract.
func bindBridgebnb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgebnbMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgebnb *BridgebnbRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgebnb.Contract.BridgebnbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgebnb *BridgebnbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgebnb.Contract.BridgebnbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgebnb *BridgebnbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgebnb.Contract.BridgebnbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgebnb *BridgebnbCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgebnb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgebnb *BridgebnbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgebnb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgebnb *BridgebnbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgebnb.Contract.contract.Transact(opts, method, params...)
}

// LDtoSDConversionRate is a free data retrieval call binding the contract method 0xe823553e.
//
// Solidity: function LDtoSDConversionRate(address ) view returns(uint256)
func (_Bridgebnb *BridgebnbCaller) LDtoSDConversionRate(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "LDtoSDConversionRate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LDtoSDConversionRate is a free data retrieval call binding the contract method 0xe823553e.
//
// Solidity: function LDtoSDConversionRate(address ) view returns(uint256)
func (_Bridgebnb *BridgebnbSession) LDtoSDConversionRate(arg0 common.Address) (*big.Int, error) {
	return _Bridgebnb.Contract.LDtoSDConversionRate(&_Bridgebnb.CallOpts, arg0)
}

// LDtoSDConversionRate is a free data retrieval call binding the contract method 0xe823553e.
//
// Solidity: function LDtoSDConversionRate(address ) view returns(uint256)
func (_Bridgebnb *BridgebnbCallerSession) LDtoSDConversionRate(arg0 common.Address) (*big.Int, error) {
	return _Bridgebnb.Contract.LDtoSDConversionRate(&_Bridgebnb.CallOpts, arg0)
}

// PTMINT is a free data retrieval call binding the contract method 0x46f6f9b5.
//
// Solidity: function PT_MINT() view returns(uint8)
func (_Bridgebnb *BridgebnbCaller) PTMINT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "PT_MINT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PTMINT is a free data retrieval call binding the contract method 0x46f6f9b5.
//
// Solidity: function PT_MINT() view returns(uint8)
func (_Bridgebnb *BridgebnbSession) PTMINT() (uint8, error) {
	return _Bridgebnb.Contract.PTMINT(&_Bridgebnb.CallOpts)
}

// PTMINT is a free data retrieval call binding the contract method 0x46f6f9b5.
//
// Solidity: function PT_MINT() view returns(uint8)
func (_Bridgebnb *BridgebnbCallerSession) PTMINT() (uint8, error) {
	return _Bridgebnb.Contract.PTMINT(&_Bridgebnb.CallOpts)
}

// PTUNLOCK is a free data retrieval call binding the contract method 0x68ea28b0.
//
// Solidity: function PT_UNLOCK() view returns(uint8)
func (_Bridgebnb *BridgebnbCaller) PTUNLOCK(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "PT_UNLOCK")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PTUNLOCK is a free data retrieval call binding the contract method 0x68ea28b0.
//
// Solidity: function PT_UNLOCK() view returns(uint8)
func (_Bridgebnb *BridgebnbSession) PTUNLOCK() (uint8, error) {
	return _Bridgebnb.Contract.PTUNLOCK(&_Bridgebnb.CallOpts)
}

// PTUNLOCK is a free data retrieval call binding the contract method 0x68ea28b0.
//
// Solidity: function PT_UNLOCK() view returns(uint8)
func (_Bridgebnb *BridgebnbCallerSession) PTUNLOCK() (uint8, error) {
	return _Bridgebnb.Contract.PTUNLOCK(&_Bridgebnb.CallOpts)
}

// AccruedFeeLD is a free data retrieval call binding the contract method 0xa2f27ae0.
//
// Solidity: function accruedFeeLD(address token) view returns(uint256)
func (_Bridgebnb *BridgebnbCaller) AccruedFeeLD(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "accruedFeeLD", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedFeeLD is a free data retrieval call binding the contract method 0xa2f27ae0.
//
// Solidity: function accruedFeeLD(address token) view returns(uint256)
func (_Bridgebnb *BridgebnbSession) AccruedFeeLD(token common.Address) (*big.Int, error) {
	return _Bridgebnb.Contract.AccruedFeeLD(&_Bridgebnb.CallOpts, token)
}

// AccruedFeeLD is a free data retrieval call binding the contract method 0xa2f27ae0.
//
// Solidity: function accruedFeeLD(address token) view returns(uint256)
func (_Bridgebnb *BridgebnbCallerSession) AccruedFeeLD(token common.Address) (*big.Int, error) {
	return _Bridgebnb.Contract.AccruedFeeLD(&_Bridgebnb.CallOpts, token)
}

// EstimateBridgeFee is a free data retrieval call binding the contract method 0x20211678.
//
// Solidity: function estimateBridgeFee(bool useZro, bytes adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Bridgebnb *BridgebnbCaller) EstimateBridgeFee(opts *bind.CallOpts, useZro bool, adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "estimateBridgeFee", useZro, adapterParams)

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

// EstimateBridgeFee is a free data retrieval call binding the contract method 0x20211678.
//
// Solidity: function estimateBridgeFee(bool useZro, bytes adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Bridgebnb *BridgebnbSession) EstimateBridgeFee(useZro bool, adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Bridgebnb.Contract.EstimateBridgeFee(&_Bridgebnb.CallOpts, useZro, adapterParams)
}

// EstimateBridgeFee is a free data retrieval call binding the contract method 0x20211678.
//
// Solidity: function estimateBridgeFee(bool useZro, bytes adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Bridgebnb *BridgebnbCallerSession) EstimateBridgeFee(useZro bool, adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Bridgebnb.Contract.EstimateBridgeFee(&_Bridgebnb.CallOpts, useZro, adapterParams)
}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Bridgebnb *BridgebnbCaller) FailedMessages(opts *bind.CallOpts, arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "failedMessages", arg0, arg1, arg2)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Bridgebnb *BridgebnbSession) FailedMessages(arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	return _Bridgebnb.Contract.FailedMessages(&_Bridgebnb.CallOpts, arg0, arg1, arg2)
}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Bridgebnb *BridgebnbCallerSession) FailedMessages(arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	return _Bridgebnb.Contract.FailedMessages(&_Bridgebnb.CallOpts, arg0, arg1, arg2)
}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Bridgebnb *BridgebnbCaller) GetConfig(opts *bind.CallOpts, _version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "getConfig", _version, _chainId, arg2, _configType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Bridgebnb *BridgebnbSession) GetConfig(_version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	return _Bridgebnb.Contract.GetConfig(&_Bridgebnb.CallOpts, _version, _chainId, arg2, _configType)
}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Bridgebnb *BridgebnbCallerSession) GetConfig(_version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	return _Bridgebnb.Contract.GetConfig(&_Bridgebnb.CallOpts, _version, _chainId, arg2, _configType)
}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Bridgebnb *BridgebnbCaller) GetTrustedRemoteAddress(opts *bind.CallOpts, _remoteChainId uint16) ([]byte, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "getTrustedRemoteAddress", _remoteChainId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Bridgebnb *BridgebnbSession) GetTrustedRemoteAddress(_remoteChainId uint16) ([]byte, error) {
	return _Bridgebnb.Contract.GetTrustedRemoteAddress(&_Bridgebnb.CallOpts, _remoteChainId)
}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Bridgebnb *BridgebnbCallerSession) GetTrustedRemoteAddress(_remoteChainId uint16) ([]byte, error) {
	return _Bridgebnb.Contract.GetTrustedRemoteAddress(&_Bridgebnb.CallOpts, _remoteChainId)
}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Bridgebnb *BridgebnbCaller) IsTrustedRemote(opts *bind.CallOpts, _srcChainId uint16, _srcAddress []byte) (bool, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "isTrustedRemote", _srcChainId, _srcAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Bridgebnb *BridgebnbSession) IsTrustedRemote(_srcChainId uint16, _srcAddress []byte) (bool, error) {
	return _Bridgebnb.Contract.IsTrustedRemote(&_Bridgebnb.CallOpts, _srcChainId, _srcAddress)
}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Bridgebnb *BridgebnbCallerSession) IsTrustedRemote(_srcChainId uint16, _srcAddress []byte) (bool, error) {
	return _Bridgebnb.Contract.IsTrustedRemote(&_Bridgebnb.CallOpts, _srcChainId, _srcAddress)
}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Bridgebnb *BridgebnbCaller) LzEndpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "lzEndpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Bridgebnb *BridgebnbSession) LzEndpoint() (common.Address, error) {
	return _Bridgebnb.Contract.LzEndpoint(&_Bridgebnb.CallOpts)
}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Bridgebnb *BridgebnbCallerSession) LzEndpoint() (common.Address, error) {
	return _Bridgebnb.Contract.LzEndpoint(&_Bridgebnb.CallOpts)
}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Bridgebnb *BridgebnbCaller) MinDstGasLookup(opts *bind.CallOpts, arg0 uint16, arg1 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "minDstGasLookup", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Bridgebnb *BridgebnbSession) MinDstGasLookup(arg0 uint16, arg1 uint16) (*big.Int, error) {
	return _Bridgebnb.Contract.MinDstGasLookup(&_Bridgebnb.CallOpts, arg0, arg1)
}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Bridgebnb *BridgebnbCallerSession) MinDstGasLookup(arg0 uint16, arg1 uint16) (*big.Int, error) {
	return _Bridgebnb.Contract.MinDstGasLookup(&_Bridgebnb.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridgebnb *BridgebnbCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridgebnb *BridgebnbSession) Owner() (common.Address, error) {
	return _Bridgebnb.Contract.Owner(&_Bridgebnb.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridgebnb *BridgebnbCallerSession) Owner() (common.Address, error) {
	return _Bridgebnb.Contract.Owner(&_Bridgebnb.CallOpts)
}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Bridgebnb *BridgebnbCaller) Precrime(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "precrime")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Bridgebnb *BridgebnbSession) Precrime() (common.Address, error) {
	return _Bridgebnb.Contract.Precrime(&_Bridgebnb.CallOpts)
}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Bridgebnb *BridgebnbCallerSession) Precrime() (common.Address, error) {
	return _Bridgebnb.Contract.Precrime(&_Bridgebnb.CallOpts)
}

// RemoteChainId is a free data retrieval call binding the contract method 0xe9518196.
//
// Solidity: function remoteChainId() view returns(uint16)
func (_Bridgebnb *BridgebnbCaller) RemoteChainId(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "remoteChainId")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RemoteChainId is a free data retrieval call binding the contract method 0xe9518196.
//
// Solidity: function remoteChainId() view returns(uint16)
func (_Bridgebnb *BridgebnbSession) RemoteChainId() (uint16, error) {
	return _Bridgebnb.Contract.RemoteChainId(&_Bridgebnb.CallOpts)
}

// RemoteChainId is a free data retrieval call binding the contract method 0xe9518196.
//
// Solidity: function remoteChainId() view returns(uint16)
func (_Bridgebnb *BridgebnbCallerSession) RemoteChainId() (uint16, error) {
	return _Bridgebnb.Contract.RemoteChainId(&_Bridgebnb.CallOpts)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool)
func (_Bridgebnb *BridgebnbCaller) SupportedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "supportedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool)
func (_Bridgebnb *BridgebnbSession) SupportedTokens(arg0 common.Address) (bool, error) {
	return _Bridgebnb.Contract.SupportedTokens(&_Bridgebnb.CallOpts, arg0)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool)
func (_Bridgebnb *BridgebnbCallerSession) SupportedTokens(arg0 common.Address) (bool, error) {
	return _Bridgebnb.Contract.SupportedTokens(&_Bridgebnb.CallOpts, arg0)
}

// TotalValueLockedSD is a free data retrieval call binding the contract method 0x2d09c4ed.
//
// Solidity: function totalValueLockedSD(address ) view returns(uint256)
func (_Bridgebnb *BridgebnbCaller) TotalValueLockedSD(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "totalValueLockedSD", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalValueLockedSD is a free data retrieval call binding the contract method 0x2d09c4ed.
//
// Solidity: function totalValueLockedSD(address ) view returns(uint256)
func (_Bridgebnb *BridgebnbSession) TotalValueLockedSD(arg0 common.Address) (*big.Int, error) {
	return _Bridgebnb.Contract.TotalValueLockedSD(&_Bridgebnb.CallOpts, arg0)
}

// TotalValueLockedSD is a free data retrieval call binding the contract method 0x2d09c4ed.
//
// Solidity: function totalValueLockedSD(address ) view returns(uint256)
func (_Bridgebnb *BridgebnbCallerSession) TotalValueLockedSD(arg0 common.Address) (*big.Int, error) {
	return _Bridgebnb.Contract.TotalValueLockedSD(&_Bridgebnb.CallOpts, arg0)
}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Bridgebnb *BridgebnbCaller) TrustedRemoteLookup(opts *bind.CallOpts, arg0 uint16) ([]byte, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "trustedRemoteLookup", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Bridgebnb *BridgebnbSession) TrustedRemoteLookup(arg0 uint16) ([]byte, error) {
	return _Bridgebnb.Contract.TrustedRemoteLookup(&_Bridgebnb.CallOpts, arg0)
}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Bridgebnb *BridgebnbCallerSession) TrustedRemoteLookup(arg0 uint16) ([]byte, error) {
	return _Bridgebnb.Contract.TrustedRemoteLookup(&_Bridgebnb.CallOpts, arg0)
}

// UseCustomAdapterParams is a free data retrieval call binding the contract method 0xed629c5c.
//
// Solidity: function useCustomAdapterParams() view returns(bool)
func (_Bridgebnb *BridgebnbCaller) UseCustomAdapterParams(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "useCustomAdapterParams")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseCustomAdapterParams is a free data retrieval call binding the contract method 0xed629c5c.
//
// Solidity: function useCustomAdapterParams() view returns(bool)
func (_Bridgebnb *BridgebnbSession) UseCustomAdapterParams() (bool, error) {
	return _Bridgebnb.Contract.UseCustomAdapterParams(&_Bridgebnb.CallOpts)
}

// UseCustomAdapterParams is a free data retrieval call binding the contract method 0xed629c5c.
//
// Solidity: function useCustomAdapterParams() view returns(bool)
func (_Bridgebnb *BridgebnbCallerSession) UseCustomAdapterParams() (bool, error) {
	return _Bridgebnb.Contract.UseCustomAdapterParams(&_Bridgebnb.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Bridgebnb *BridgebnbCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgebnb.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Bridgebnb *BridgebnbSession) Weth() (common.Address, error) {
	return _Bridgebnb.Contract.Weth(&_Bridgebnb.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Bridgebnb *BridgebnbCallerSession) Weth() (common.Address, error) {
	return _Bridgebnb.Contract.Weth(&_Bridgebnb.CallOpts)
}

// Bridge is a paid mutator transaction binding the contract method 0xfe359a0d.
//
// Solidity: function bridge(address token, uint256 amountLD, address to, (address,address) callParams, bytes adapterParams) payable returns()
func (_Bridgebnb *BridgebnbTransactor) Bridge(opts *bind.TransactOpts, token common.Address, amountLD *big.Int, to common.Address, callParams LzLibCallParams, adapterParams []byte) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "bridge", token, amountLD, to, callParams, adapterParams)
}

// Bridge is a paid mutator transaction binding the contract method 0xfe359a0d.
//
// Solidity: function bridge(address token, uint256 amountLD, address to, (address,address) callParams, bytes adapterParams) payable returns()
func (_Bridgebnb *BridgebnbSession) Bridge(token common.Address, amountLD *big.Int, to common.Address, callParams LzLibCallParams, adapterParams []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.Bridge(&_Bridgebnb.TransactOpts, token, amountLD, to, callParams, adapterParams)
}

// Bridge is a paid mutator transaction binding the contract method 0xfe359a0d.
//
// Solidity: function bridge(address token, uint256 amountLD, address to, (address,address) callParams, bytes adapterParams) payable returns()
func (_Bridgebnb *BridgebnbTransactorSession) Bridge(token common.Address, amountLD *big.Int, to common.Address, callParams LzLibCallParams, adapterParams []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.Bridge(&_Bridgebnb.TransactOpts, token, amountLD, to, callParams, adapterParams)
}

// BridgeETH is a paid mutator transaction binding the contract method 0xde7aaff4.
//
// Solidity: function bridgeETH(uint256 amountLD, address to, (address,address) callParams, bytes adapterParams) payable returns()
func (_Bridgebnb *BridgebnbTransactor) BridgeETH(opts *bind.TransactOpts, amountLD *big.Int, to common.Address, callParams LzLibCallParams, adapterParams []byte) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "bridgeETH", amountLD, to, callParams, adapterParams)
}

// BridgeETH is a paid mutator transaction binding the contract method 0xde7aaff4.
//
// Solidity: function bridgeETH(uint256 amountLD, address to, (address,address) callParams, bytes adapterParams) payable returns()
func (_Bridgebnb *BridgebnbSession) BridgeETH(amountLD *big.Int, to common.Address, callParams LzLibCallParams, adapterParams []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.BridgeETH(&_Bridgebnb.TransactOpts, amountLD, to, callParams, adapterParams)
}

// BridgeETH is a paid mutator transaction binding the contract method 0xde7aaff4.
//
// Solidity: function bridgeETH(uint256 amountLD, address to, (address,address) callParams, bytes adapterParams) payable returns()
func (_Bridgebnb *BridgebnbTransactorSession) BridgeETH(amountLD *big.Int, to common.Address, callParams LzLibCallParams, adapterParams []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.BridgeETH(&_Bridgebnb.TransactOpts, amountLD, to, callParams, adapterParams)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Bridgebnb *BridgebnbTransactor) ForceResumeReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "forceResumeReceive", _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Bridgebnb *BridgebnbSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.ForceResumeReceive(&_Bridgebnb.TransactOpts, _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Bridgebnb *BridgebnbTransactorSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.ForceResumeReceive(&_Bridgebnb.TransactOpts, _srcChainId, _srcAddress)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Bridgebnb *BridgebnbTransactor) LzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "lzReceive", _srcChainId, _srcAddress, _nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Bridgebnb *BridgebnbSession) LzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.LzReceive(&_Bridgebnb.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Bridgebnb *BridgebnbTransactorSession) LzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.LzReceive(&_Bridgebnb.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Bridgebnb *BridgebnbTransactor) NonblockingLzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "nonblockingLzReceive", _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Bridgebnb *BridgebnbSession) NonblockingLzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.NonblockingLzReceive(&_Bridgebnb.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Bridgebnb *BridgebnbTransactorSession) NonblockingLzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.NonblockingLzReceive(&_Bridgebnb.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x5a4967e5.
//
// Solidity: function registerToken(address token, uint8 sharedDecimals) returns()
func (_Bridgebnb *BridgebnbTransactor) RegisterToken(opts *bind.TransactOpts, token common.Address, sharedDecimals uint8) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "registerToken", token, sharedDecimals)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x5a4967e5.
//
// Solidity: function registerToken(address token, uint8 sharedDecimals) returns()
func (_Bridgebnb *BridgebnbSession) RegisterToken(token common.Address, sharedDecimals uint8) (*types.Transaction, error) {
	return _Bridgebnb.Contract.RegisterToken(&_Bridgebnb.TransactOpts, token, sharedDecimals)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x5a4967e5.
//
// Solidity: function registerToken(address token, uint8 sharedDecimals) returns()
func (_Bridgebnb *BridgebnbTransactorSession) RegisterToken(token common.Address, sharedDecimals uint8) (*types.Transaction, error) {
	return _Bridgebnb.Contract.RegisterToken(&_Bridgebnb.TransactOpts, token, sharedDecimals)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridgebnb *BridgebnbTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridgebnb *BridgebnbSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridgebnb.Contract.RenounceOwnership(&_Bridgebnb.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridgebnb *BridgebnbTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridgebnb.Contract.RenounceOwnership(&_Bridgebnb.TransactOpts)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Bridgebnb *BridgebnbTransactor) RetryMessage(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "retryMessage", _srcChainId, _srcAddress, _nonce, _payload)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Bridgebnb *BridgebnbSession) RetryMessage(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.RetryMessage(&_Bridgebnb.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Bridgebnb *BridgebnbTransactorSession) RetryMessage(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.RetryMessage(&_Bridgebnb.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Bridgebnb *BridgebnbTransactor) SetConfig(opts *bind.TransactOpts, _version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "setConfig", _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Bridgebnb *BridgebnbSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetConfig(&_Bridgebnb.TransactOpts, _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Bridgebnb *BridgebnbTransactorSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetConfig(&_Bridgebnb.TransactOpts, _version, _chainId, _configType, _config)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Bridgebnb *BridgebnbTransactor) SetMinDstGas(opts *bind.TransactOpts, _dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "setMinDstGas", _dstChainId, _packetType, _minGas)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Bridgebnb *BridgebnbSession) SetMinDstGas(_dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetMinDstGas(&_Bridgebnb.TransactOpts, _dstChainId, _packetType, _minGas)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Bridgebnb *BridgebnbTransactorSession) SetMinDstGas(_dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetMinDstGas(&_Bridgebnb.TransactOpts, _dstChainId, _packetType, _minGas)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Bridgebnb *BridgebnbTransactor) SetPrecrime(opts *bind.TransactOpts, _precrime common.Address) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "setPrecrime", _precrime)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Bridgebnb *BridgebnbSession) SetPrecrime(_precrime common.Address) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetPrecrime(&_Bridgebnb.TransactOpts, _precrime)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Bridgebnb *BridgebnbTransactorSession) SetPrecrime(_precrime common.Address) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetPrecrime(&_Bridgebnb.TransactOpts, _precrime)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Bridgebnb *BridgebnbTransactor) SetReceiveVersion(opts *bind.TransactOpts, _version uint16) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "setReceiveVersion", _version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Bridgebnb *BridgebnbSession) SetReceiveVersion(_version uint16) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetReceiveVersion(&_Bridgebnb.TransactOpts, _version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Bridgebnb *BridgebnbTransactorSession) SetReceiveVersion(_version uint16) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetReceiveVersion(&_Bridgebnb.TransactOpts, _version)
}

// SetRemoteChainId is a paid mutator transaction binding the contract method 0x38db1ebc.
//
// Solidity: function setRemoteChainId(uint16 _remoteChainId) returns()
func (_Bridgebnb *BridgebnbTransactor) SetRemoteChainId(opts *bind.TransactOpts, _remoteChainId uint16) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "setRemoteChainId", _remoteChainId)
}

// SetRemoteChainId is a paid mutator transaction binding the contract method 0x38db1ebc.
//
// Solidity: function setRemoteChainId(uint16 _remoteChainId) returns()
func (_Bridgebnb *BridgebnbSession) SetRemoteChainId(_remoteChainId uint16) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetRemoteChainId(&_Bridgebnb.TransactOpts, _remoteChainId)
}

// SetRemoteChainId is a paid mutator transaction binding the contract method 0x38db1ebc.
//
// Solidity: function setRemoteChainId(uint16 _remoteChainId) returns()
func (_Bridgebnb *BridgebnbTransactorSession) SetRemoteChainId(_remoteChainId uint16) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetRemoteChainId(&_Bridgebnb.TransactOpts, _remoteChainId)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Bridgebnb *BridgebnbTransactor) SetSendVersion(opts *bind.TransactOpts, _version uint16) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "setSendVersion", _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Bridgebnb *BridgebnbSession) SetSendVersion(_version uint16) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetSendVersion(&_Bridgebnb.TransactOpts, _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Bridgebnb *BridgebnbTransactorSession) SetSendVersion(_version uint16) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetSendVersion(&_Bridgebnb.TransactOpts, _version)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _srcChainId, bytes _path) returns()
func (_Bridgebnb *BridgebnbTransactor) SetTrustedRemote(opts *bind.TransactOpts, _srcChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "setTrustedRemote", _srcChainId, _path)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _srcChainId, bytes _path) returns()
func (_Bridgebnb *BridgebnbSession) SetTrustedRemote(_srcChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetTrustedRemote(&_Bridgebnb.TransactOpts, _srcChainId, _path)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _srcChainId, bytes _path) returns()
func (_Bridgebnb *BridgebnbTransactorSession) SetTrustedRemote(_srcChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetTrustedRemote(&_Bridgebnb.TransactOpts, _srcChainId, _path)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Bridgebnb *BridgebnbTransactor) SetTrustedRemoteAddress(opts *bind.TransactOpts, _remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "setTrustedRemoteAddress", _remoteChainId, _remoteAddress)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Bridgebnb *BridgebnbSession) SetTrustedRemoteAddress(_remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetTrustedRemoteAddress(&_Bridgebnb.TransactOpts, _remoteChainId, _remoteAddress)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Bridgebnb *BridgebnbTransactorSession) SetTrustedRemoteAddress(_remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetTrustedRemoteAddress(&_Bridgebnb.TransactOpts, _remoteChainId, _remoteAddress)
}

// SetUseCustomAdapterParams is a paid mutator transaction binding the contract method 0xeab45d9c.
//
// Solidity: function setUseCustomAdapterParams(bool _useCustomAdapterParams) returns()
func (_Bridgebnb *BridgebnbTransactor) SetUseCustomAdapterParams(opts *bind.TransactOpts, _useCustomAdapterParams bool) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "setUseCustomAdapterParams", _useCustomAdapterParams)
}

// SetUseCustomAdapterParams is a paid mutator transaction binding the contract method 0xeab45d9c.
//
// Solidity: function setUseCustomAdapterParams(bool _useCustomAdapterParams) returns()
func (_Bridgebnb *BridgebnbSession) SetUseCustomAdapterParams(_useCustomAdapterParams bool) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetUseCustomAdapterParams(&_Bridgebnb.TransactOpts, _useCustomAdapterParams)
}

// SetUseCustomAdapterParams is a paid mutator transaction binding the contract method 0xeab45d9c.
//
// Solidity: function setUseCustomAdapterParams(bool _useCustomAdapterParams) returns()
func (_Bridgebnb *BridgebnbTransactorSession) SetUseCustomAdapterParams(_useCustomAdapterParams bool) (*types.Transaction, error) {
	return _Bridgebnb.Contract.SetUseCustomAdapterParams(&_Bridgebnb.TransactOpts, _useCustomAdapterParams)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridgebnb *BridgebnbTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridgebnb *BridgebnbSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridgebnb.Contract.TransferOwnership(&_Bridgebnb.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridgebnb *BridgebnbTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridgebnb.Contract.TransferOwnership(&_Bridgebnb.TransactOpts, newOwner)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1095b6d7.
//
// Solidity: function withdrawFee(address token, address to, uint256 amountLD) returns()
func (_Bridgebnb *BridgebnbTransactor) WithdrawFee(opts *bind.TransactOpts, token common.Address, to common.Address, amountLD *big.Int) (*types.Transaction, error) {
	return _Bridgebnb.contract.Transact(opts, "withdrawFee", token, to, amountLD)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1095b6d7.
//
// Solidity: function withdrawFee(address token, address to, uint256 amountLD) returns()
func (_Bridgebnb *BridgebnbSession) WithdrawFee(token common.Address, to common.Address, amountLD *big.Int) (*types.Transaction, error) {
	return _Bridgebnb.Contract.WithdrawFee(&_Bridgebnb.TransactOpts, token, to, amountLD)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1095b6d7.
//
// Solidity: function withdrawFee(address token, address to, uint256 amountLD) returns()
func (_Bridgebnb *BridgebnbTransactorSession) WithdrawFee(token common.Address, to common.Address, amountLD *big.Int) (*types.Transaction, error) {
	return _Bridgebnb.Contract.WithdrawFee(&_Bridgebnb.TransactOpts, token, to, amountLD)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridgebnb *BridgebnbTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgebnb.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridgebnb *BridgebnbSession) Receive() (*types.Transaction, error) {
	return _Bridgebnb.Contract.Receive(&_Bridgebnb.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridgebnb *BridgebnbTransactorSession) Receive() (*types.Transaction, error) {
	return _Bridgebnb.Contract.Receive(&_Bridgebnb.TransactOpts)
}

// BridgebnbMessageFailedIterator is returned from FilterMessageFailed and is used to iterate over the raw logs and unpacked data for MessageFailed events raised by the Bridgebnb contract.
type BridgebnbMessageFailedIterator struct {
	Event *BridgebnbMessageFailed // Event containing the contract specifics and raw log

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
func (it *BridgebnbMessageFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgebnbMessageFailed)
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
		it.Event = new(BridgebnbMessageFailed)
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
func (it *BridgebnbMessageFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgebnbMessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgebnbMessageFailed represents a MessageFailed event raised by the Bridgebnb contract.
type BridgebnbMessageFailed struct {
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
func (_Bridgebnb *BridgebnbFilterer) FilterMessageFailed(opts *bind.FilterOpts) (*BridgebnbMessageFailedIterator, error) {

	logs, sub, err := _Bridgebnb.contract.FilterLogs(opts, "MessageFailed")
	if err != nil {
		return nil, err
	}
	return &BridgebnbMessageFailedIterator{contract: _Bridgebnb.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

// WatchMessageFailed is a free log subscription operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_Bridgebnb *BridgebnbFilterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *BridgebnbMessageFailed) (event.Subscription, error) {

	logs, sub, err := _Bridgebnb.contract.WatchLogs(opts, "MessageFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgebnbMessageFailed)
				if err := _Bridgebnb.contract.UnpackLog(event, "MessageFailed", log); err != nil {
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
func (_Bridgebnb *BridgebnbFilterer) ParseMessageFailed(log types.Log) (*BridgebnbMessageFailed, error) {
	event := new(BridgebnbMessageFailed)
	if err := _Bridgebnb.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgebnbOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridgebnb contract.
type BridgebnbOwnershipTransferredIterator struct {
	Event *BridgebnbOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgebnbOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgebnbOwnershipTransferred)
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
		it.Event = new(BridgebnbOwnershipTransferred)
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
func (it *BridgebnbOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgebnbOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgebnbOwnershipTransferred represents a OwnershipTransferred event raised by the Bridgebnb contract.
type BridgebnbOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridgebnb *BridgebnbFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgebnbOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridgebnb.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgebnbOwnershipTransferredIterator{contract: _Bridgebnb.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridgebnb *BridgebnbFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgebnbOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridgebnb.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgebnbOwnershipTransferred)
				if err := _Bridgebnb.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bridgebnb *BridgebnbFilterer) ParseOwnershipTransferred(log types.Log) (*BridgebnbOwnershipTransferred, error) {
	event := new(BridgebnbOwnershipTransferred)
	if err := _Bridgebnb.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgebnbReceiveTokenIterator is returned from FilterReceiveToken and is used to iterate over the raw logs and unpacked data for ReceiveToken events raised by the Bridgebnb contract.
type BridgebnbReceiveTokenIterator struct {
	Event *BridgebnbReceiveToken // Event containing the contract specifics and raw log

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
func (it *BridgebnbReceiveTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgebnbReceiveToken)
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
		it.Event = new(BridgebnbReceiveToken)
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
func (it *BridgebnbReceiveTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgebnbReceiveTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgebnbReceiveToken represents a ReceiveToken event raised by the Bridgebnb contract.
type BridgebnbReceiveToken struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceiveToken is a free log retrieval operation binding the contract event 0x5e3da8fba24af91505c66214c9e629ba712ce2c1b8c318f14f7024fdcba544a8.
//
// Solidity: event ReceiveToken(address token, address to, uint256 amount)
func (_Bridgebnb *BridgebnbFilterer) FilterReceiveToken(opts *bind.FilterOpts) (*BridgebnbReceiveTokenIterator, error) {

	logs, sub, err := _Bridgebnb.contract.FilterLogs(opts, "ReceiveToken")
	if err != nil {
		return nil, err
	}
	return &BridgebnbReceiveTokenIterator{contract: _Bridgebnb.contract, event: "ReceiveToken", logs: logs, sub: sub}, nil
}

// WatchReceiveToken is a free log subscription operation binding the contract event 0x5e3da8fba24af91505c66214c9e629ba712ce2c1b8c318f14f7024fdcba544a8.
//
// Solidity: event ReceiveToken(address token, address to, uint256 amount)
func (_Bridgebnb *BridgebnbFilterer) WatchReceiveToken(opts *bind.WatchOpts, sink chan<- *BridgebnbReceiveToken) (event.Subscription, error) {

	logs, sub, err := _Bridgebnb.contract.WatchLogs(opts, "ReceiveToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgebnbReceiveToken)
				if err := _Bridgebnb.contract.UnpackLog(event, "ReceiveToken", log); err != nil {
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

// ParseReceiveToken is a log parse operation binding the contract event 0x5e3da8fba24af91505c66214c9e629ba712ce2c1b8c318f14f7024fdcba544a8.
//
// Solidity: event ReceiveToken(address token, address to, uint256 amount)
func (_Bridgebnb *BridgebnbFilterer) ParseReceiveToken(log types.Log) (*BridgebnbReceiveToken, error) {
	event := new(BridgebnbReceiveToken)
	if err := _Bridgebnb.contract.UnpackLog(event, "ReceiveToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgebnbRegisterTokenIterator is returned from FilterRegisterToken and is used to iterate over the raw logs and unpacked data for RegisterToken events raised by the Bridgebnb contract.
type BridgebnbRegisterTokenIterator struct {
	Event *BridgebnbRegisterToken // Event containing the contract specifics and raw log

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
func (it *BridgebnbRegisterTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgebnbRegisterToken)
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
		it.Event = new(BridgebnbRegisterToken)
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
func (it *BridgebnbRegisterTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgebnbRegisterTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgebnbRegisterToken represents a RegisterToken event raised by the Bridgebnb contract.
type BridgebnbRegisterToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRegisterToken is a free log retrieval operation binding the contract event 0xf7fe8023cb2e36bde1d59a88ac5763a8c11be6d25e6819f71bb7e23e5bf0dc16.
//
// Solidity: event RegisterToken(address token)
func (_Bridgebnb *BridgebnbFilterer) FilterRegisterToken(opts *bind.FilterOpts) (*BridgebnbRegisterTokenIterator, error) {

	logs, sub, err := _Bridgebnb.contract.FilterLogs(opts, "RegisterToken")
	if err != nil {
		return nil, err
	}
	return &BridgebnbRegisterTokenIterator{contract: _Bridgebnb.contract, event: "RegisterToken", logs: logs, sub: sub}, nil
}

// WatchRegisterToken is a free log subscription operation binding the contract event 0xf7fe8023cb2e36bde1d59a88ac5763a8c11be6d25e6819f71bb7e23e5bf0dc16.
//
// Solidity: event RegisterToken(address token)
func (_Bridgebnb *BridgebnbFilterer) WatchRegisterToken(opts *bind.WatchOpts, sink chan<- *BridgebnbRegisterToken) (event.Subscription, error) {

	logs, sub, err := _Bridgebnb.contract.WatchLogs(opts, "RegisterToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgebnbRegisterToken)
				if err := _Bridgebnb.contract.UnpackLog(event, "RegisterToken", log); err != nil {
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

// ParseRegisterToken is a log parse operation binding the contract event 0xf7fe8023cb2e36bde1d59a88ac5763a8c11be6d25e6819f71bb7e23e5bf0dc16.
//
// Solidity: event RegisterToken(address token)
func (_Bridgebnb *BridgebnbFilterer) ParseRegisterToken(log types.Log) (*BridgebnbRegisterToken, error) {
	event := new(BridgebnbRegisterToken)
	if err := _Bridgebnb.contract.UnpackLog(event, "RegisterToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgebnbRetryMessageSuccessIterator is returned from FilterRetryMessageSuccess and is used to iterate over the raw logs and unpacked data for RetryMessageSuccess events raised by the Bridgebnb contract.
type BridgebnbRetryMessageSuccessIterator struct {
	Event *BridgebnbRetryMessageSuccess // Event containing the contract specifics and raw log

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
func (it *BridgebnbRetryMessageSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgebnbRetryMessageSuccess)
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
		it.Event = new(BridgebnbRetryMessageSuccess)
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
func (it *BridgebnbRetryMessageSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgebnbRetryMessageSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgebnbRetryMessageSuccess represents a RetryMessageSuccess event raised by the Bridgebnb contract.
type BridgebnbRetryMessageSuccess struct {
	SrcChainId  uint16
	SrcAddress  []byte
	Nonce       uint64
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRetryMessageSuccess is a free log retrieval operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_Bridgebnb *BridgebnbFilterer) FilterRetryMessageSuccess(opts *bind.FilterOpts) (*BridgebnbRetryMessageSuccessIterator, error) {

	logs, sub, err := _Bridgebnb.contract.FilterLogs(opts, "RetryMessageSuccess")
	if err != nil {
		return nil, err
	}
	return &BridgebnbRetryMessageSuccessIterator{contract: _Bridgebnb.contract, event: "RetryMessageSuccess", logs: logs, sub: sub}, nil
}

// WatchRetryMessageSuccess is a free log subscription operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_Bridgebnb *BridgebnbFilterer) WatchRetryMessageSuccess(opts *bind.WatchOpts, sink chan<- *BridgebnbRetryMessageSuccess) (event.Subscription, error) {

	logs, sub, err := _Bridgebnb.contract.WatchLogs(opts, "RetryMessageSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgebnbRetryMessageSuccess)
				if err := _Bridgebnb.contract.UnpackLog(event, "RetryMessageSuccess", log); err != nil {
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
func (_Bridgebnb *BridgebnbFilterer) ParseRetryMessageSuccess(log types.Log) (*BridgebnbRetryMessageSuccess, error) {
	event := new(BridgebnbRetryMessageSuccess)
	if err := _Bridgebnb.contract.UnpackLog(event, "RetryMessageSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgebnbSendTokenIterator is returned from FilterSendToken and is used to iterate over the raw logs and unpacked data for SendToken events raised by the Bridgebnb contract.
type BridgebnbSendTokenIterator struct {
	Event *BridgebnbSendToken // Event containing the contract specifics and raw log

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
func (it *BridgebnbSendTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgebnbSendToken)
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
		it.Event = new(BridgebnbSendToken)
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
func (it *BridgebnbSendTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgebnbSendTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgebnbSendToken represents a SendToken event raised by the Bridgebnb contract.
type BridgebnbSendToken struct {
	Token  common.Address
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSendToken is a free log retrieval operation binding the contract event 0x49b9b5358c9580b3e6c5ee10b8b260c1e64bede87cb8a212e9e20a0b7dc20e5a.
//
// Solidity: event SendToken(address token, address from, address to, uint256 amount)
func (_Bridgebnb *BridgebnbFilterer) FilterSendToken(opts *bind.FilterOpts) (*BridgebnbSendTokenIterator, error) {

	logs, sub, err := _Bridgebnb.contract.FilterLogs(opts, "SendToken")
	if err != nil {
		return nil, err
	}
	return &BridgebnbSendTokenIterator{contract: _Bridgebnb.contract, event: "SendToken", logs: logs, sub: sub}, nil
}

// WatchSendToken is a free log subscription operation binding the contract event 0x49b9b5358c9580b3e6c5ee10b8b260c1e64bede87cb8a212e9e20a0b7dc20e5a.
//
// Solidity: event SendToken(address token, address from, address to, uint256 amount)
func (_Bridgebnb *BridgebnbFilterer) WatchSendToken(opts *bind.WatchOpts, sink chan<- *BridgebnbSendToken) (event.Subscription, error) {

	logs, sub, err := _Bridgebnb.contract.WatchLogs(opts, "SendToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgebnbSendToken)
				if err := _Bridgebnb.contract.UnpackLog(event, "SendToken", log); err != nil {
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

// ParseSendToken is a log parse operation binding the contract event 0x49b9b5358c9580b3e6c5ee10b8b260c1e64bede87cb8a212e9e20a0b7dc20e5a.
//
// Solidity: event SendToken(address token, address from, address to, uint256 amount)
func (_Bridgebnb *BridgebnbFilterer) ParseSendToken(log types.Log) (*BridgebnbSendToken, error) {
	event := new(BridgebnbSendToken)
	if err := _Bridgebnb.contract.UnpackLog(event, "SendToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgebnbSetMinDstGasIterator is returned from FilterSetMinDstGas and is used to iterate over the raw logs and unpacked data for SetMinDstGas events raised by the Bridgebnb contract.
type BridgebnbSetMinDstGasIterator struct {
	Event *BridgebnbSetMinDstGas // Event containing the contract specifics and raw log

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
func (it *BridgebnbSetMinDstGasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgebnbSetMinDstGas)
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
		it.Event = new(BridgebnbSetMinDstGas)
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
func (it *BridgebnbSetMinDstGasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgebnbSetMinDstGasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgebnbSetMinDstGas represents a SetMinDstGas event raised by the Bridgebnb contract.
type BridgebnbSetMinDstGas struct {
	DstChainId uint16
	Type       uint16
	MinDstGas  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetMinDstGas is a free log retrieval operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_Bridgebnb *BridgebnbFilterer) FilterSetMinDstGas(opts *bind.FilterOpts) (*BridgebnbSetMinDstGasIterator, error) {

	logs, sub, err := _Bridgebnb.contract.FilterLogs(opts, "SetMinDstGas")
	if err != nil {
		return nil, err
	}
	return &BridgebnbSetMinDstGasIterator{contract: _Bridgebnb.contract, event: "SetMinDstGas", logs: logs, sub: sub}, nil
}

// WatchSetMinDstGas is a free log subscription operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_Bridgebnb *BridgebnbFilterer) WatchSetMinDstGas(opts *bind.WatchOpts, sink chan<- *BridgebnbSetMinDstGas) (event.Subscription, error) {

	logs, sub, err := _Bridgebnb.contract.WatchLogs(opts, "SetMinDstGas")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgebnbSetMinDstGas)
				if err := _Bridgebnb.contract.UnpackLog(event, "SetMinDstGas", log); err != nil {
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
func (_Bridgebnb *BridgebnbFilterer) ParseSetMinDstGas(log types.Log) (*BridgebnbSetMinDstGas, error) {
	event := new(BridgebnbSetMinDstGas)
	if err := _Bridgebnb.contract.UnpackLog(event, "SetMinDstGas", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgebnbSetPrecrimeIterator is returned from FilterSetPrecrime and is used to iterate over the raw logs and unpacked data for SetPrecrime events raised by the Bridgebnb contract.
type BridgebnbSetPrecrimeIterator struct {
	Event *BridgebnbSetPrecrime // Event containing the contract specifics and raw log

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
func (it *BridgebnbSetPrecrimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgebnbSetPrecrime)
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
		it.Event = new(BridgebnbSetPrecrime)
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
func (it *BridgebnbSetPrecrimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgebnbSetPrecrimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgebnbSetPrecrime represents a SetPrecrime event raised by the Bridgebnb contract.
type BridgebnbSetPrecrime struct {
	Precrime common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPrecrime is a free log retrieval operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_Bridgebnb *BridgebnbFilterer) FilterSetPrecrime(opts *bind.FilterOpts) (*BridgebnbSetPrecrimeIterator, error) {

	logs, sub, err := _Bridgebnb.contract.FilterLogs(opts, "SetPrecrime")
	if err != nil {
		return nil, err
	}
	return &BridgebnbSetPrecrimeIterator{contract: _Bridgebnb.contract, event: "SetPrecrime", logs: logs, sub: sub}, nil
}

// WatchSetPrecrime is a free log subscription operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_Bridgebnb *BridgebnbFilterer) WatchSetPrecrime(opts *bind.WatchOpts, sink chan<- *BridgebnbSetPrecrime) (event.Subscription, error) {

	logs, sub, err := _Bridgebnb.contract.WatchLogs(opts, "SetPrecrime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgebnbSetPrecrime)
				if err := _Bridgebnb.contract.UnpackLog(event, "SetPrecrime", log); err != nil {
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
func (_Bridgebnb *BridgebnbFilterer) ParseSetPrecrime(log types.Log) (*BridgebnbSetPrecrime, error) {
	event := new(BridgebnbSetPrecrime)
	if err := _Bridgebnb.contract.UnpackLog(event, "SetPrecrime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgebnbSetRemoteChainIdIterator is returned from FilterSetRemoteChainId and is used to iterate over the raw logs and unpacked data for SetRemoteChainId events raised by the Bridgebnb contract.
type BridgebnbSetRemoteChainIdIterator struct {
	Event *BridgebnbSetRemoteChainId // Event containing the contract specifics and raw log

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
func (it *BridgebnbSetRemoteChainIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgebnbSetRemoteChainId)
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
		it.Event = new(BridgebnbSetRemoteChainId)
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
func (it *BridgebnbSetRemoteChainIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgebnbSetRemoteChainIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgebnbSetRemoteChainId represents a SetRemoteChainId event raised by the Bridgebnb contract.
type BridgebnbSetRemoteChainId struct {
	RemoteChainId uint16
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetRemoteChainId is a free log retrieval operation binding the contract event 0xe8df78a276e2b718a366328e9120b436ea83832fbeede026392fed933e3ffa5b.
//
// Solidity: event SetRemoteChainId(uint16 remoteChainId)
func (_Bridgebnb *BridgebnbFilterer) FilterSetRemoteChainId(opts *bind.FilterOpts) (*BridgebnbSetRemoteChainIdIterator, error) {

	logs, sub, err := _Bridgebnb.contract.FilterLogs(opts, "SetRemoteChainId")
	if err != nil {
		return nil, err
	}
	return &BridgebnbSetRemoteChainIdIterator{contract: _Bridgebnb.contract, event: "SetRemoteChainId", logs: logs, sub: sub}, nil
}

// WatchSetRemoteChainId is a free log subscription operation binding the contract event 0xe8df78a276e2b718a366328e9120b436ea83832fbeede026392fed933e3ffa5b.
//
// Solidity: event SetRemoteChainId(uint16 remoteChainId)
func (_Bridgebnb *BridgebnbFilterer) WatchSetRemoteChainId(opts *bind.WatchOpts, sink chan<- *BridgebnbSetRemoteChainId) (event.Subscription, error) {

	logs, sub, err := _Bridgebnb.contract.WatchLogs(opts, "SetRemoteChainId")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgebnbSetRemoteChainId)
				if err := _Bridgebnb.contract.UnpackLog(event, "SetRemoteChainId", log); err != nil {
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

// ParseSetRemoteChainId is a log parse operation binding the contract event 0xe8df78a276e2b718a366328e9120b436ea83832fbeede026392fed933e3ffa5b.
//
// Solidity: event SetRemoteChainId(uint16 remoteChainId)
func (_Bridgebnb *BridgebnbFilterer) ParseSetRemoteChainId(log types.Log) (*BridgebnbSetRemoteChainId, error) {
	event := new(BridgebnbSetRemoteChainId)
	if err := _Bridgebnb.contract.UnpackLog(event, "SetRemoteChainId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgebnbSetTrustedRemoteIterator is returned from FilterSetTrustedRemote and is used to iterate over the raw logs and unpacked data for SetTrustedRemote events raised by the Bridgebnb contract.
type BridgebnbSetTrustedRemoteIterator struct {
	Event *BridgebnbSetTrustedRemote // Event containing the contract specifics and raw log

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
func (it *BridgebnbSetTrustedRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgebnbSetTrustedRemote)
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
		it.Event = new(BridgebnbSetTrustedRemote)
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
func (it *BridgebnbSetTrustedRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgebnbSetTrustedRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgebnbSetTrustedRemote represents a SetTrustedRemote event raised by the Bridgebnb contract.
type BridgebnbSetTrustedRemote struct {
	RemoteChainId uint16
	Path          []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemote is a free log retrieval operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_Bridgebnb *BridgebnbFilterer) FilterSetTrustedRemote(opts *bind.FilterOpts) (*BridgebnbSetTrustedRemoteIterator, error) {

	logs, sub, err := _Bridgebnb.contract.FilterLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return &BridgebnbSetTrustedRemoteIterator{contract: _Bridgebnb.contract, event: "SetTrustedRemote", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemote is a free log subscription operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_Bridgebnb *BridgebnbFilterer) WatchSetTrustedRemote(opts *bind.WatchOpts, sink chan<- *BridgebnbSetTrustedRemote) (event.Subscription, error) {

	logs, sub, err := _Bridgebnb.contract.WatchLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgebnbSetTrustedRemote)
				if err := _Bridgebnb.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
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
func (_Bridgebnb *BridgebnbFilterer) ParseSetTrustedRemote(log types.Log) (*BridgebnbSetTrustedRemote, error) {
	event := new(BridgebnbSetTrustedRemote)
	if err := _Bridgebnb.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgebnbSetTrustedRemoteAddressIterator is returned from FilterSetTrustedRemoteAddress and is used to iterate over the raw logs and unpacked data for SetTrustedRemoteAddress events raised by the Bridgebnb contract.
type BridgebnbSetTrustedRemoteAddressIterator struct {
	Event *BridgebnbSetTrustedRemoteAddress // Event containing the contract specifics and raw log

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
func (it *BridgebnbSetTrustedRemoteAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgebnbSetTrustedRemoteAddress)
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
		it.Event = new(BridgebnbSetTrustedRemoteAddress)
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
func (it *BridgebnbSetTrustedRemoteAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgebnbSetTrustedRemoteAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgebnbSetTrustedRemoteAddress represents a SetTrustedRemoteAddress event raised by the Bridgebnb contract.
type BridgebnbSetTrustedRemoteAddress struct {
	RemoteChainId uint16
	RemoteAddress []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemoteAddress is a free log retrieval operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_Bridgebnb *BridgebnbFilterer) FilterSetTrustedRemoteAddress(opts *bind.FilterOpts) (*BridgebnbSetTrustedRemoteAddressIterator, error) {

	logs, sub, err := _Bridgebnb.contract.FilterLogs(opts, "SetTrustedRemoteAddress")
	if err != nil {
		return nil, err
	}
	return &BridgebnbSetTrustedRemoteAddressIterator{contract: _Bridgebnb.contract, event: "SetTrustedRemoteAddress", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemoteAddress is a free log subscription operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_Bridgebnb *BridgebnbFilterer) WatchSetTrustedRemoteAddress(opts *bind.WatchOpts, sink chan<- *BridgebnbSetTrustedRemoteAddress) (event.Subscription, error) {

	logs, sub, err := _Bridgebnb.contract.WatchLogs(opts, "SetTrustedRemoteAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgebnbSetTrustedRemoteAddress)
				if err := _Bridgebnb.contract.UnpackLog(event, "SetTrustedRemoteAddress", log); err != nil {
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
func (_Bridgebnb *BridgebnbFilterer) ParseSetTrustedRemoteAddress(log types.Log) (*BridgebnbSetTrustedRemoteAddress, error) {
	event := new(BridgebnbSetTrustedRemoteAddress)
	if err := _Bridgebnb.contract.UnpackLog(event, "SetTrustedRemoteAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgebnbSetUseCustomAdapterParamsIterator is returned from FilterSetUseCustomAdapterParams and is used to iterate over the raw logs and unpacked data for SetUseCustomAdapterParams events raised by the Bridgebnb contract.
type BridgebnbSetUseCustomAdapterParamsIterator struct {
	Event *BridgebnbSetUseCustomAdapterParams // Event containing the contract specifics and raw log

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
func (it *BridgebnbSetUseCustomAdapterParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgebnbSetUseCustomAdapterParams)
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
		it.Event = new(BridgebnbSetUseCustomAdapterParams)
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
func (it *BridgebnbSetUseCustomAdapterParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgebnbSetUseCustomAdapterParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgebnbSetUseCustomAdapterParams represents a SetUseCustomAdapterParams event raised by the Bridgebnb contract.
type BridgebnbSetUseCustomAdapterParams struct {
	UseCustomAdapterParams bool
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetUseCustomAdapterParams is a free log retrieval operation binding the contract event 0x1584ad594a70cbe1e6515592e1272a987d922b097ead875069cebe8b40c004a4.
//
// Solidity: event SetUseCustomAdapterParams(bool useCustomAdapterParams)
func (_Bridgebnb *BridgebnbFilterer) FilterSetUseCustomAdapterParams(opts *bind.FilterOpts) (*BridgebnbSetUseCustomAdapterParamsIterator, error) {

	logs, sub, err := _Bridgebnb.contract.FilterLogs(opts, "SetUseCustomAdapterParams")
	if err != nil {
		return nil, err
	}
	return &BridgebnbSetUseCustomAdapterParamsIterator{contract: _Bridgebnb.contract, event: "SetUseCustomAdapterParams", logs: logs, sub: sub}, nil
}

// WatchSetUseCustomAdapterParams is a free log subscription operation binding the contract event 0x1584ad594a70cbe1e6515592e1272a987d922b097ead875069cebe8b40c004a4.
//
// Solidity: event SetUseCustomAdapterParams(bool useCustomAdapterParams)
func (_Bridgebnb *BridgebnbFilterer) WatchSetUseCustomAdapterParams(opts *bind.WatchOpts, sink chan<- *BridgebnbSetUseCustomAdapterParams) (event.Subscription, error) {

	logs, sub, err := _Bridgebnb.contract.WatchLogs(opts, "SetUseCustomAdapterParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgebnbSetUseCustomAdapterParams)
				if err := _Bridgebnb.contract.UnpackLog(event, "SetUseCustomAdapterParams", log); err != nil {
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
func (_Bridgebnb *BridgebnbFilterer) ParseSetUseCustomAdapterParams(log types.Log) (*BridgebnbSetUseCustomAdapterParams, error) {
	event := new(BridgebnbSetUseCustomAdapterParams)
	if err := _Bridgebnb.contract.UnpackLog(event, "SetUseCustomAdapterParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgebnbWithdrawFeeIterator is returned from FilterWithdrawFee and is used to iterate over the raw logs and unpacked data for WithdrawFee events raised by the Bridgebnb contract.
type BridgebnbWithdrawFeeIterator struct {
	Event *BridgebnbWithdrawFee // Event containing the contract specifics and raw log

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
func (it *BridgebnbWithdrawFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgebnbWithdrawFee)
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
		it.Event = new(BridgebnbWithdrawFee)
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
func (it *BridgebnbWithdrawFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgebnbWithdrawFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgebnbWithdrawFee represents a WithdrawFee event raised by the Bridgebnb contract.
type BridgebnbWithdrawFee struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFee is a free log retrieval operation binding the contract event 0xf15a0a3784dea9b4fe33bc98e2450745e262d310237b2868ea8ef56967ff3ecb.
//
// Solidity: event WithdrawFee(address indexed token, address to, uint256 amount)
func (_Bridgebnb *BridgebnbFilterer) FilterWithdrawFee(opts *bind.FilterOpts, token []common.Address) (*BridgebnbWithdrawFeeIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridgebnb.contract.FilterLogs(opts, "WithdrawFee", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgebnbWithdrawFeeIterator{contract: _Bridgebnb.contract, event: "WithdrawFee", logs: logs, sub: sub}, nil
}

// WatchWithdrawFee is a free log subscription operation binding the contract event 0xf15a0a3784dea9b4fe33bc98e2450745e262d310237b2868ea8ef56967ff3ecb.
//
// Solidity: event WithdrawFee(address indexed token, address to, uint256 amount)
func (_Bridgebnb *BridgebnbFilterer) WatchWithdrawFee(opts *bind.WatchOpts, sink chan<- *BridgebnbWithdrawFee, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridgebnb.contract.WatchLogs(opts, "WithdrawFee", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgebnbWithdrawFee)
				if err := _Bridgebnb.contract.UnpackLog(event, "WithdrawFee", log); err != nil {
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

// ParseWithdrawFee is a log parse operation binding the contract event 0xf15a0a3784dea9b4fe33bc98e2450745e262d310237b2868ea8ef56967ff3ecb.
//
// Solidity: event WithdrawFee(address indexed token, address to, uint256 amount)
func (_Bridgebnb *BridgebnbFilterer) ParseWithdrawFee(log types.Log) (*BridgebnbWithdrawFee, error) {
	event := new(BridgebnbWithdrawFee)
	if err := _Bridgebnb.contract.UnpackLog(event, "WithdrawFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
