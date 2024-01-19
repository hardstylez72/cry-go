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

// MerklyrefuelMetaData contains all meta data concerning the Merklyrefuel contract.
var MerklyrefuelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lzEndpoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_reason\",\"type\":\"bytes\"}],\"name\":\"MessageFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"RetryMessageSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_type\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minDstGas\",\"type\":\"uint256\"}],\"name\":\"SetMinDstGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"precrime\",\"type\":\"address\"}],\"name\":\"SetPrecrime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_path\",\"type\":\"bytes\"}],\"name\":\"SetTrustedRemote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_remoteAddress\",\"type\":\"bytes\"}],\"name\":\"SetTrustedRemoteAddress\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_PAYLOAD_SIZE_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NO_EXTRA_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PT_SEND\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"bridgeGas\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"estimateSendFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zroFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"failedMessages\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"}],\"name\":\"forceResumeReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_configType\",\"type\":\"uint256\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"}],\"name\":\"getTrustedRemoteAddress\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"}],\"name\":\"isTrustedRemote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lzEndpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"minDstGasLookup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"nonblockingLzReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"payloadSizeLimitLookup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precrime\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"retryMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_configType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_config\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_packetType\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minGas\",\"type\":\"uint256\"}],\"name\":\"setMinDstGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"name\":\"setPayloadSizeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_precrime\",\"type\":\"address\"}],\"name\":\"setPrecrime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"}],\"name\":\"setReceiveVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"}],\"name\":\"setSendVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_path\",\"type\":\"bytes\"}],\"name\":\"setTrustedRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_remoteAddress\",\"type\":\"bytes\"}],\"name\":\"setTrustedRemoteAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"trustedRemoteLookup\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useCustomAdapterParams\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// MerklyrefuelABI is the input ABI used to generate the binding from.
// Deprecated: Use MerklyrefuelMetaData.ABI instead.
var MerklyrefuelABI = MerklyrefuelMetaData.ABI

// Merklyrefuel is an auto generated Go binding around an Ethereum contract.
type Merklyrefuel struct {
	MerklyrefuelCaller     // Read-only binding to the contract
	MerklyrefuelTransactor // Write-only binding to the contract
	MerklyrefuelFilterer   // Log filterer for contract events
}

// MerklyrefuelCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerklyrefuelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerklyrefuelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerklyrefuelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerklyrefuelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerklyrefuelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerklyrefuelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerklyrefuelSession struct {
	Contract     *Merklyrefuel     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerklyrefuelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerklyrefuelCallerSession struct {
	Contract *MerklyrefuelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MerklyrefuelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerklyrefuelTransactorSession struct {
	Contract     *MerklyrefuelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MerklyrefuelRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerklyrefuelRaw struct {
	Contract *Merklyrefuel // Generic contract binding to access the raw methods on
}

// MerklyrefuelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerklyrefuelCallerRaw struct {
	Contract *MerklyrefuelCaller // Generic read-only contract binding to access the raw methods on
}

// MerklyrefuelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerklyrefuelTransactorRaw struct {
	Contract *MerklyrefuelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerklyrefuel creates a new instance of Merklyrefuel, bound to a specific deployed contract.
func NewMerklyrefuel(address common.Address, backend bind.ContractBackend) (*Merklyrefuel, error) {
	contract, err := bindMerklyrefuel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Merklyrefuel{MerklyrefuelCaller: MerklyrefuelCaller{contract: contract}, MerklyrefuelTransactor: MerklyrefuelTransactor{contract: contract}, MerklyrefuelFilterer: MerklyrefuelFilterer{contract: contract}}, nil
}

// NewMerklyrefuelCaller creates a new read-only instance of Merklyrefuel, bound to a specific deployed contract.
func NewMerklyrefuelCaller(address common.Address, caller bind.ContractCaller) (*MerklyrefuelCaller, error) {
	contract, err := bindMerklyrefuel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerklyrefuelCaller{contract: contract}, nil
}

// NewMerklyrefuelTransactor creates a new write-only instance of Merklyrefuel, bound to a specific deployed contract.
func NewMerklyrefuelTransactor(address common.Address, transactor bind.ContractTransactor) (*MerklyrefuelTransactor, error) {
	contract, err := bindMerklyrefuel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerklyrefuelTransactor{contract: contract}, nil
}

// NewMerklyrefuelFilterer creates a new log filterer instance of Merklyrefuel, bound to a specific deployed contract.
func NewMerklyrefuelFilterer(address common.Address, filterer bind.ContractFilterer) (*MerklyrefuelFilterer, error) {
	contract, err := bindMerklyrefuel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerklyrefuelFilterer{contract: contract}, nil
}

// bindMerklyrefuel binds a generic wrapper to an already deployed contract.
func bindMerklyrefuel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MerklyrefuelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merklyrefuel *MerklyrefuelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Merklyrefuel.Contract.MerklyrefuelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merklyrefuel *MerklyrefuelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.MerklyrefuelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merklyrefuel *MerklyrefuelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.MerklyrefuelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merklyrefuel *MerklyrefuelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Merklyrefuel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merklyrefuel *MerklyrefuelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merklyrefuel *MerklyrefuelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTPAYLOADSIZELIMIT is a free data retrieval call binding the contract method 0xc4461834.
//
// Solidity: function DEFAULT_PAYLOAD_SIZE_LIMIT() view returns(uint256)
func (_Merklyrefuel *MerklyrefuelCaller) DEFAULTPAYLOADSIZELIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Merklyrefuel.contract.Call(opts, &out, "DEFAULT_PAYLOAD_SIZE_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTPAYLOADSIZELIMIT is a free data retrieval call binding the contract method 0xc4461834.
//
// Solidity: function DEFAULT_PAYLOAD_SIZE_LIMIT() view returns(uint256)
func (_Merklyrefuel *MerklyrefuelSession) DEFAULTPAYLOADSIZELIMIT() (*big.Int, error) {
	return _Merklyrefuel.Contract.DEFAULTPAYLOADSIZELIMIT(&_Merklyrefuel.CallOpts)
}

// DEFAULTPAYLOADSIZELIMIT is a free data retrieval call binding the contract method 0xc4461834.
//
// Solidity: function DEFAULT_PAYLOAD_SIZE_LIMIT() view returns(uint256)
func (_Merklyrefuel *MerklyrefuelCallerSession) DEFAULTPAYLOADSIZELIMIT() (*big.Int, error) {
	return _Merklyrefuel.Contract.DEFAULTPAYLOADSIZELIMIT(&_Merklyrefuel.CallOpts)
}

// NOEXTRAGAS is a free data retrieval call binding the contract method 0x44770515.
//
// Solidity: function NO_EXTRA_GAS() view returns(uint256)
func (_Merklyrefuel *MerklyrefuelCaller) NOEXTRAGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Merklyrefuel.contract.Call(opts, &out, "NO_EXTRA_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NOEXTRAGAS is a free data retrieval call binding the contract method 0x44770515.
//
// Solidity: function NO_EXTRA_GAS() view returns(uint256)
func (_Merklyrefuel *MerklyrefuelSession) NOEXTRAGAS() (*big.Int, error) {
	return _Merklyrefuel.Contract.NOEXTRAGAS(&_Merklyrefuel.CallOpts)
}

// NOEXTRAGAS is a free data retrieval call binding the contract method 0x44770515.
//
// Solidity: function NO_EXTRA_GAS() view returns(uint256)
func (_Merklyrefuel *MerklyrefuelCallerSession) NOEXTRAGAS() (*big.Int, error) {
	return _Merklyrefuel.Contract.NOEXTRAGAS(&_Merklyrefuel.CallOpts)
}

// PTSEND is a free data retrieval call binding the contract method 0x4c42899a.
//
// Solidity: function PT_SEND() view returns(uint16)
func (_Merklyrefuel *MerklyrefuelCaller) PTSEND(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Merklyrefuel.contract.Call(opts, &out, "PT_SEND")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PTSEND is a free data retrieval call binding the contract method 0x4c42899a.
//
// Solidity: function PT_SEND() view returns(uint16)
func (_Merklyrefuel *MerklyrefuelSession) PTSEND() (uint16, error) {
	return _Merklyrefuel.Contract.PTSEND(&_Merklyrefuel.CallOpts)
}

// PTSEND is a free data retrieval call binding the contract method 0x4c42899a.
//
// Solidity: function PT_SEND() view returns(uint16)
func (_Merklyrefuel *MerklyrefuelCallerSession) PTSEND() (uint16, error) {
	return _Merklyrefuel.Contract.PTSEND(&_Merklyrefuel.CallOpts)
}

// EstimateSendFee is a free data retrieval call binding the contract method 0x5a749876.
//
// Solidity: function estimateSendFee(uint16 _dstChainId, bytes payload, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Merklyrefuel *MerklyrefuelCaller) EstimateSendFee(opts *bind.CallOpts, _dstChainId uint16, payload []byte, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Merklyrefuel.contract.Call(opts, &out, "estimateSendFee", _dstChainId, payload, _adapterParams)

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

// EstimateSendFee is a free data retrieval call binding the contract method 0x5a749876.
//
// Solidity: function estimateSendFee(uint16 _dstChainId, bytes payload, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Merklyrefuel *MerklyrefuelSession) EstimateSendFee(_dstChainId uint16, payload []byte, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Merklyrefuel.Contract.EstimateSendFee(&_Merklyrefuel.CallOpts, _dstChainId, payload, _adapterParams)
}

// EstimateSendFee is a free data retrieval call binding the contract method 0x5a749876.
//
// Solidity: function estimateSendFee(uint16 _dstChainId, bytes payload, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Merklyrefuel *MerklyrefuelCallerSession) EstimateSendFee(_dstChainId uint16, payload []byte, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Merklyrefuel.Contract.EstimateSendFee(&_Merklyrefuel.CallOpts, _dstChainId, payload, _adapterParams)
}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Merklyrefuel *MerklyrefuelCaller) FailedMessages(opts *bind.CallOpts, arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Merklyrefuel.contract.Call(opts, &out, "failedMessages", arg0, arg1, arg2)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Merklyrefuel *MerklyrefuelSession) FailedMessages(arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	return _Merklyrefuel.Contract.FailedMessages(&_Merklyrefuel.CallOpts, arg0, arg1, arg2)
}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Merklyrefuel *MerklyrefuelCallerSession) FailedMessages(arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	return _Merklyrefuel.Contract.FailedMessages(&_Merklyrefuel.CallOpts, arg0, arg1, arg2)
}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Merklyrefuel *MerklyrefuelCaller) GetConfig(opts *bind.CallOpts, _version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Merklyrefuel.contract.Call(opts, &out, "getConfig", _version, _chainId, arg2, _configType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Merklyrefuel *MerklyrefuelSession) GetConfig(_version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	return _Merklyrefuel.Contract.GetConfig(&_Merklyrefuel.CallOpts, _version, _chainId, arg2, _configType)
}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Merklyrefuel *MerklyrefuelCallerSession) GetConfig(_version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	return _Merklyrefuel.Contract.GetConfig(&_Merklyrefuel.CallOpts, _version, _chainId, arg2, _configType)
}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Merklyrefuel *MerklyrefuelCaller) GetTrustedRemoteAddress(opts *bind.CallOpts, _remoteChainId uint16) ([]byte, error) {
	var out []interface{}
	err := _Merklyrefuel.contract.Call(opts, &out, "getTrustedRemoteAddress", _remoteChainId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Merklyrefuel *MerklyrefuelSession) GetTrustedRemoteAddress(_remoteChainId uint16) ([]byte, error) {
	return _Merklyrefuel.Contract.GetTrustedRemoteAddress(&_Merklyrefuel.CallOpts, _remoteChainId)
}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Merklyrefuel *MerklyrefuelCallerSession) GetTrustedRemoteAddress(_remoteChainId uint16) ([]byte, error) {
	return _Merklyrefuel.Contract.GetTrustedRemoteAddress(&_Merklyrefuel.CallOpts, _remoteChainId)
}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Merklyrefuel *MerklyrefuelCaller) IsTrustedRemote(opts *bind.CallOpts, _srcChainId uint16, _srcAddress []byte) (bool, error) {
	var out []interface{}
	err := _Merklyrefuel.contract.Call(opts, &out, "isTrustedRemote", _srcChainId, _srcAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Merklyrefuel *MerklyrefuelSession) IsTrustedRemote(_srcChainId uint16, _srcAddress []byte) (bool, error) {
	return _Merklyrefuel.Contract.IsTrustedRemote(&_Merklyrefuel.CallOpts, _srcChainId, _srcAddress)
}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Merklyrefuel *MerklyrefuelCallerSession) IsTrustedRemote(_srcChainId uint16, _srcAddress []byte) (bool, error) {
	return _Merklyrefuel.Contract.IsTrustedRemote(&_Merklyrefuel.CallOpts, _srcChainId, _srcAddress)
}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Merklyrefuel *MerklyrefuelCaller) LzEndpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Merklyrefuel.contract.Call(opts, &out, "lzEndpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Merklyrefuel *MerklyrefuelSession) LzEndpoint() (common.Address, error) {
	return _Merklyrefuel.Contract.LzEndpoint(&_Merklyrefuel.CallOpts)
}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Merklyrefuel *MerklyrefuelCallerSession) LzEndpoint() (common.Address, error) {
	return _Merklyrefuel.Contract.LzEndpoint(&_Merklyrefuel.CallOpts)
}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Merklyrefuel *MerklyrefuelCaller) MinDstGasLookup(opts *bind.CallOpts, arg0 uint16, arg1 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Merklyrefuel.contract.Call(opts, &out, "minDstGasLookup", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Merklyrefuel *MerklyrefuelSession) MinDstGasLookup(arg0 uint16, arg1 uint16) (*big.Int, error) {
	return _Merklyrefuel.Contract.MinDstGasLookup(&_Merklyrefuel.CallOpts, arg0, arg1)
}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Merklyrefuel *MerklyrefuelCallerSession) MinDstGasLookup(arg0 uint16, arg1 uint16) (*big.Int, error) {
	return _Merklyrefuel.Contract.MinDstGasLookup(&_Merklyrefuel.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Merklyrefuel *MerklyrefuelCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Merklyrefuel.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Merklyrefuel *MerklyrefuelSession) Owner() (common.Address, error) {
	return _Merklyrefuel.Contract.Owner(&_Merklyrefuel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Merklyrefuel *MerklyrefuelCallerSession) Owner() (common.Address, error) {
	return _Merklyrefuel.Contract.Owner(&_Merklyrefuel.CallOpts)
}

// PayloadSizeLimitLookup is a free data retrieval call binding the contract method 0x3f1f4fa4.
//
// Solidity: function payloadSizeLimitLookup(uint16 ) view returns(uint256)
func (_Merklyrefuel *MerklyrefuelCaller) PayloadSizeLimitLookup(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Merklyrefuel.contract.Call(opts, &out, "payloadSizeLimitLookup", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayloadSizeLimitLookup is a free data retrieval call binding the contract method 0x3f1f4fa4.
//
// Solidity: function payloadSizeLimitLookup(uint16 ) view returns(uint256)
func (_Merklyrefuel *MerklyrefuelSession) PayloadSizeLimitLookup(arg0 uint16) (*big.Int, error) {
	return _Merklyrefuel.Contract.PayloadSizeLimitLookup(&_Merklyrefuel.CallOpts, arg0)
}

// PayloadSizeLimitLookup is a free data retrieval call binding the contract method 0x3f1f4fa4.
//
// Solidity: function payloadSizeLimitLookup(uint16 ) view returns(uint256)
func (_Merklyrefuel *MerklyrefuelCallerSession) PayloadSizeLimitLookup(arg0 uint16) (*big.Int, error) {
	return _Merklyrefuel.Contract.PayloadSizeLimitLookup(&_Merklyrefuel.CallOpts, arg0)
}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Merklyrefuel *MerklyrefuelCaller) Precrime(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Merklyrefuel.contract.Call(opts, &out, "precrime")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Merklyrefuel *MerklyrefuelSession) Precrime() (common.Address, error) {
	return _Merklyrefuel.Contract.Precrime(&_Merklyrefuel.CallOpts)
}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Merklyrefuel *MerklyrefuelCallerSession) Precrime() (common.Address, error) {
	return _Merklyrefuel.Contract.Precrime(&_Merklyrefuel.CallOpts)
}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Merklyrefuel *MerklyrefuelCaller) TrustedRemoteLookup(opts *bind.CallOpts, arg0 uint16) ([]byte, error) {
	var out []interface{}
	err := _Merklyrefuel.contract.Call(opts, &out, "trustedRemoteLookup", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Merklyrefuel *MerklyrefuelSession) TrustedRemoteLookup(arg0 uint16) ([]byte, error) {
	return _Merklyrefuel.Contract.TrustedRemoteLookup(&_Merklyrefuel.CallOpts, arg0)
}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Merklyrefuel *MerklyrefuelCallerSession) TrustedRemoteLookup(arg0 uint16) ([]byte, error) {
	return _Merklyrefuel.Contract.TrustedRemoteLookup(&_Merklyrefuel.CallOpts, arg0)
}

// UseCustomAdapterParams is a free data retrieval call binding the contract method 0xed629c5c.
//
// Solidity: function useCustomAdapterParams() view returns(bool)
func (_Merklyrefuel *MerklyrefuelCaller) UseCustomAdapterParams(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Merklyrefuel.contract.Call(opts, &out, "useCustomAdapterParams")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseCustomAdapterParams is a free data retrieval call binding the contract method 0xed629c5c.
//
// Solidity: function useCustomAdapterParams() view returns(bool)
func (_Merklyrefuel *MerklyrefuelSession) UseCustomAdapterParams() (bool, error) {
	return _Merklyrefuel.Contract.UseCustomAdapterParams(&_Merklyrefuel.CallOpts)
}

// UseCustomAdapterParams is a free data retrieval call binding the contract method 0xed629c5c.
//
// Solidity: function useCustomAdapterParams() view returns(bool)
func (_Merklyrefuel *MerklyrefuelCallerSession) UseCustomAdapterParams() (bool, error) {
	return _Merklyrefuel.Contract.UseCustomAdapterParams(&_Merklyrefuel.CallOpts)
}

// BridgeGas is a paid mutator transaction binding the contract method 0x126928c4.
//
// Solidity: function bridgeGas(uint16 _dstChainId, bytes _toAddress, bytes _adapterParams) payable returns()
func (_Merklyrefuel *MerklyrefuelTransactor) BridgeGas(opts *bind.TransactOpts, _dstChainId uint16, _toAddress []byte, _adapterParams []byte) (*types.Transaction, error) {
	return _Merklyrefuel.contract.Transact(opts, "bridgeGas", _dstChainId, _toAddress, _adapterParams)
}

// BridgeGas is a paid mutator transaction binding the contract method 0x126928c4.
//
// Solidity: function bridgeGas(uint16 _dstChainId, bytes _toAddress, bytes _adapterParams) payable returns()
func (_Merklyrefuel *MerklyrefuelSession) BridgeGas(_dstChainId uint16, _toAddress []byte, _adapterParams []byte) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.BridgeGas(&_Merklyrefuel.TransactOpts, _dstChainId, _toAddress, _adapterParams)
}

// BridgeGas is a paid mutator transaction binding the contract method 0x126928c4.
//
// Solidity: function bridgeGas(uint16 _dstChainId, bytes _toAddress, bytes _adapterParams) payable returns()
func (_Merklyrefuel *MerklyrefuelTransactorSession) BridgeGas(_dstChainId uint16, _toAddress []byte, _adapterParams []byte) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.BridgeGas(&_Merklyrefuel.TransactOpts, _dstChainId, _toAddress, _adapterParams)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Merklyrefuel *MerklyrefuelTransactor) ForceResumeReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Merklyrefuel.contract.Transact(opts, "forceResumeReceive", _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Merklyrefuel *MerklyrefuelSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.ForceResumeReceive(&_Merklyrefuel.TransactOpts, _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Merklyrefuel *MerklyrefuelTransactorSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.ForceResumeReceive(&_Merklyrefuel.TransactOpts, _srcChainId, _srcAddress)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Merklyrefuel *MerklyrefuelTransactor) LzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Merklyrefuel.contract.Transact(opts, "lzReceive", _srcChainId, _srcAddress, _nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Merklyrefuel *MerklyrefuelSession) LzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.LzReceive(&_Merklyrefuel.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Merklyrefuel *MerklyrefuelTransactorSession) LzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.LzReceive(&_Merklyrefuel.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Merklyrefuel *MerklyrefuelTransactor) NonblockingLzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Merklyrefuel.contract.Transact(opts, "nonblockingLzReceive", _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Merklyrefuel *MerklyrefuelSession) NonblockingLzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.NonblockingLzReceive(&_Merklyrefuel.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Merklyrefuel *MerklyrefuelTransactorSession) NonblockingLzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.NonblockingLzReceive(&_Merklyrefuel.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Merklyrefuel *MerklyrefuelTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merklyrefuel.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Merklyrefuel *MerklyrefuelSession) RenounceOwnership() (*types.Transaction, error) {
	return _Merklyrefuel.Contract.RenounceOwnership(&_Merklyrefuel.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Merklyrefuel *MerklyrefuelTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Merklyrefuel.Contract.RenounceOwnership(&_Merklyrefuel.TransactOpts)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Merklyrefuel *MerklyrefuelTransactor) RetryMessage(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Merklyrefuel.contract.Transact(opts, "retryMessage", _srcChainId, _srcAddress, _nonce, _payload)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Merklyrefuel *MerklyrefuelSession) RetryMessage(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.RetryMessage(&_Merklyrefuel.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Merklyrefuel *MerklyrefuelTransactorSession) RetryMessage(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.RetryMessage(&_Merklyrefuel.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Merklyrefuel *MerklyrefuelTransactor) SetConfig(opts *bind.TransactOpts, _version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Merklyrefuel.contract.Transact(opts, "setConfig", _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Merklyrefuel *MerklyrefuelSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.SetConfig(&_Merklyrefuel.TransactOpts, _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Merklyrefuel *MerklyrefuelTransactorSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.SetConfig(&_Merklyrefuel.TransactOpts, _version, _chainId, _configType, _config)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Merklyrefuel *MerklyrefuelTransactor) SetMinDstGas(opts *bind.TransactOpts, _dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Merklyrefuel.contract.Transact(opts, "setMinDstGas", _dstChainId, _packetType, _minGas)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Merklyrefuel *MerklyrefuelSession) SetMinDstGas(_dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.SetMinDstGas(&_Merklyrefuel.TransactOpts, _dstChainId, _packetType, _minGas)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Merklyrefuel *MerklyrefuelTransactorSession) SetMinDstGas(_dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.SetMinDstGas(&_Merklyrefuel.TransactOpts, _dstChainId, _packetType, _minGas)
}

// SetPayloadSizeLimit is a paid mutator transaction binding the contract method 0x0df37483.
//
// Solidity: function setPayloadSizeLimit(uint16 _dstChainId, uint256 _size) returns()
func (_Merklyrefuel *MerklyrefuelTransactor) SetPayloadSizeLimit(opts *bind.TransactOpts, _dstChainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _Merklyrefuel.contract.Transact(opts, "setPayloadSizeLimit", _dstChainId, _size)
}

// SetPayloadSizeLimit is a paid mutator transaction binding the contract method 0x0df37483.
//
// Solidity: function setPayloadSizeLimit(uint16 _dstChainId, uint256 _size) returns()
func (_Merklyrefuel *MerklyrefuelSession) SetPayloadSizeLimit(_dstChainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.SetPayloadSizeLimit(&_Merklyrefuel.TransactOpts, _dstChainId, _size)
}

// SetPayloadSizeLimit is a paid mutator transaction binding the contract method 0x0df37483.
//
// Solidity: function setPayloadSizeLimit(uint16 _dstChainId, uint256 _size) returns()
func (_Merklyrefuel *MerklyrefuelTransactorSession) SetPayloadSizeLimit(_dstChainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.SetPayloadSizeLimit(&_Merklyrefuel.TransactOpts, _dstChainId, _size)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Merklyrefuel *MerklyrefuelTransactor) SetPrecrime(opts *bind.TransactOpts, _precrime common.Address) (*types.Transaction, error) {
	return _Merklyrefuel.contract.Transact(opts, "setPrecrime", _precrime)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Merklyrefuel *MerklyrefuelSession) SetPrecrime(_precrime common.Address) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.SetPrecrime(&_Merklyrefuel.TransactOpts, _precrime)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Merklyrefuel *MerklyrefuelTransactorSession) SetPrecrime(_precrime common.Address) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.SetPrecrime(&_Merklyrefuel.TransactOpts, _precrime)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Merklyrefuel *MerklyrefuelTransactor) SetReceiveVersion(opts *bind.TransactOpts, _version uint16) (*types.Transaction, error) {
	return _Merklyrefuel.contract.Transact(opts, "setReceiveVersion", _version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Merklyrefuel *MerklyrefuelSession) SetReceiveVersion(_version uint16) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.SetReceiveVersion(&_Merklyrefuel.TransactOpts, _version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Merklyrefuel *MerklyrefuelTransactorSession) SetReceiveVersion(_version uint16) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.SetReceiveVersion(&_Merklyrefuel.TransactOpts, _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Merklyrefuel *MerklyrefuelTransactor) SetSendVersion(opts *bind.TransactOpts, _version uint16) (*types.Transaction, error) {
	return _Merklyrefuel.contract.Transact(opts, "setSendVersion", _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Merklyrefuel *MerklyrefuelSession) SetSendVersion(_version uint16) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.SetSendVersion(&_Merklyrefuel.TransactOpts, _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Merklyrefuel *MerklyrefuelTransactorSession) SetSendVersion(_version uint16) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.SetSendVersion(&_Merklyrefuel.TransactOpts, _version)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _remoteChainId, bytes _path) returns()
func (_Merklyrefuel *MerklyrefuelTransactor) SetTrustedRemote(opts *bind.TransactOpts, _remoteChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Merklyrefuel.contract.Transact(opts, "setTrustedRemote", _remoteChainId, _path)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _remoteChainId, bytes _path) returns()
func (_Merklyrefuel *MerklyrefuelSession) SetTrustedRemote(_remoteChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.SetTrustedRemote(&_Merklyrefuel.TransactOpts, _remoteChainId, _path)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _remoteChainId, bytes _path) returns()
func (_Merklyrefuel *MerklyrefuelTransactorSession) SetTrustedRemote(_remoteChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.SetTrustedRemote(&_Merklyrefuel.TransactOpts, _remoteChainId, _path)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Merklyrefuel *MerklyrefuelTransactor) SetTrustedRemoteAddress(opts *bind.TransactOpts, _remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Merklyrefuel.contract.Transact(opts, "setTrustedRemoteAddress", _remoteChainId, _remoteAddress)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Merklyrefuel *MerklyrefuelSession) SetTrustedRemoteAddress(_remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.SetTrustedRemoteAddress(&_Merklyrefuel.TransactOpts, _remoteChainId, _remoteAddress)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Merklyrefuel *MerklyrefuelTransactorSession) SetTrustedRemoteAddress(_remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.SetTrustedRemoteAddress(&_Merklyrefuel.TransactOpts, _remoteChainId, _remoteAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Merklyrefuel *MerklyrefuelTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Merklyrefuel.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Merklyrefuel *MerklyrefuelSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.TransferOwnership(&_Merklyrefuel.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Merklyrefuel *MerklyrefuelTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Merklyrefuel.Contract.TransferOwnership(&_Merklyrefuel.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Merklyrefuel *MerklyrefuelTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merklyrefuel.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Merklyrefuel *MerklyrefuelSession) Withdraw() (*types.Transaction, error) {
	return _Merklyrefuel.Contract.Withdraw(&_Merklyrefuel.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Merklyrefuel *MerklyrefuelTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Merklyrefuel.Contract.Withdraw(&_Merklyrefuel.TransactOpts)
}

// MerklyrefuelMessageFailedIterator is returned from FilterMessageFailed and is used to iterate over the raw logs and unpacked data for MessageFailed events raised by the Merklyrefuel contract.
type MerklyrefuelMessageFailedIterator struct {
	Event *MerklyrefuelMessageFailed // Event containing the contract specifics and raw log

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
func (it *MerklyrefuelMessageFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerklyrefuelMessageFailed)
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
		it.Event = new(MerklyrefuelMessageFailed)
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
func (it *MerklyrefuelMessageFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerklyrefuelMessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerklyrefuelMessageFailed represents a MessageFailed event raised by the Merklyrefuel contract.
type MerklyrefuelMessageFailed struct {
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
func (_Merklyrefuel *MerklyrefuelFilterer) FilterMessageFailed(opts *bind.FilterOpts) (*MerklyrefuelMessageFailedIterator, error) {

	logs, sub, err := _Merklyrefuel.contract.FilterLogs(opts, "MessageFailed")
	if err != nil {
		return nil, err
	}
	return &MerklyrefuelMessageFailedIterator{contract: _Merklyrefuel.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

// WatchMessageFailed is a free log subscription operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_Merklyrefuel *MerklyrefuelFilterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *MerklyrefuelMessageFailed) (event.Subscription, error) {

	logs, sub, err := _Merklyrefuel.contract.WatchLogs(opts, "MessageFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerklyrefuelMessageFailed)
				if err := _Merklyrefuel.contract.UnpackLog(event, "MessageFailed", log); err != nil {
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
func (_Merklyrefuel *MerklyrefuelFilterer) ParseMessageFailed(log types.Log) (*MerklyrefuelMessageFailed, error) {
	event := new(MerklyrefuelMessageFailed)
	if err := _Merklyrefuel.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerklyrefuelOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Merklyrefuel contract.
type MerklyrefuelOwnershipTransferredIterator struct {
	Event *MerklyrefuelOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MerklyrefuelOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerklyrefuelOwnershipTransferred)
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
		it.Event = new(MerklyrefuelOwnershipTransferred)
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
func (it *MerklyrefuelOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerklyrefuelOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerklyrefuelOwnershipTransferred represents a OwnershipTransferred event raised by the Merklyrefuel contract.
type MerklyrefuelOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Merklyrefuel *MerklyrefuelFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MerklyrefuelOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Merklyrefuel.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MerklyrefuelOwnershipTransferredIterator{contract: _Merklyrefuel.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Merklyrefuel *MerklyrefuelFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MerklyrefuelOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Merklyrefuel.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerklyrefuelOwnershipTransferred)
				if err := _Merklyrefuel.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Merklyrefuel *MerklyrefuelFilterer) ParseOwnershipTransferred(log types.Log) (*MerklyrefuelOwnershipTransferred, error) {
	event := new(MerklyrefuelOwnershipTransferred)
	if err := _Merklyrefuel.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerklyrefuelRetryMessageSuccessIterator is returned from FilterRetryMessageSuccess and is used to iterate over the raw logs and unpacked data for RetryMessageSuccess events raised by the Merklyrefuel contract.
type MerklyrefuelRetryMessageSuccessIterator struct {
	Event *MerklyrefuelRetryMessageSuccess // Event containing the contract specifics and raw log

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
func (it *MerklyrefuelRetryMessageSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerklyrefuelRetryMessageSuccess)
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
		it.Event = new(MerklyrefuelRetryMessageSuccess)
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
func (it *MerklyrefuelRetryMessageSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerklyrefuelRetryMessageSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerklyrefuelRetryMessageSuccess represents a RetryMessageSuccess event raised by the Merklyrefuel contract.
type MerklyrefuelRetryMessageSuccess struct {
	SrcChainId  uint16
	SrcAddress  []byte
	Nonce       uint64
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRetryMessageSuccess is a free log retrieval operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_Merklyrefuel *MerklyrefuelFilterer) FilterRetryMessageSuccess(opts *bind.FilterOpts) (*MerklyrefuelRetryMessageSuccessIterator, error) {

	logs, sub, err := _Merklyrefuel.contract.FilterLogs(opts, "RetryMessageSuccess")
	if err != nil {
		return nil, err
	}
	return &MerklyrefuelRetryMessageSuccessIterator{contract: _Merklyrefuel.contract, event: "RetryMessageSuccess", logs: logs, sub: sub}, nil
}

// WatchRetryMessageSuccess is a free log subscription operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_Merklyrefuel *MerklyrefuelFilterer) WatchRetryMessageSuccess(opts *bind.WatchOpts, sink chan<- *MerklyrefuelRetryMessageSuccess) (event.Subscription, error) {

	logs, sub, err := _Merklyrefuel.contract.WatchLogs(opts, "RetryMessageSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerklyrefuelRetryMessageSuccess)
				if err := _Merklyrefuel.contract.UnpackLog(event, "RetryMessageSuccess", log); err != nil {
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
func (_Merklyrefuel *MerklyrefuelFilterer) ParseRetryMessageSuccess(log types.Log) (*MerklyrefuelRetryMessageSuccess, error) {
	event := new(MerklyrefuelRetryMessageSuccess)
	if err := _Merklyrefuel.contract.UnpackLog(event, "RetryMessageSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerklyrefuelSetMinDstGasIterator is returned from FilterSetMinDstGas and is used to iterate over the raw logs and unpacked data for SetMinDstGas events raised by the Merklyrefuel contract.
type MerklyrefuelSetMinDstGasIterator struct {
	Event *MerklyrefuelSetMinDstGas // Event containing the contract specifics and raw log

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
func (it *MerklyrefuelSetMinDstGasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerklyrefuelSetMinDstGas)
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
		it.Event = new(MerklyrefuelSetMinDstGas)
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
func (it *MerklyrefuelSetMinDstGasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerklyrefuelSetMinDstGasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerklyrefuelSetMinDstGas represents a SetMinDstGas event raised by the Merklyrefuel contract.
type MerklyrefuelSetMinDstGas struct {
	DstChainId uint16
	Type       uint16
	MinDstGas  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetMinDstGas is a free log retrieval operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_Merklyrefuel *MerklyrefuelFilterer) FilterSetMinDstGas(opts *bind.FilterOpts) (*MerklyrefuelSetMinDstGasIterator, error) {

	logs, sub, err := _Merklyrefuel.contract.FilterLogs(opts, "SetMinDstGas")
	if err != nil {
		return nil, err
	}
	return &MerklyrefuelSetMinDstGasIterator{contract: _Merklyrefuel.contract, event: "SetMinDstGas", logs: logs, sub: sub}, nil
}

// WatchSetMinDstGas is a free log subscription operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_Merklyrefuel *MerklyrefuelFilterer) WatchSetMinDstGas(opts *bind.WatchOpts, sink chan<- *MerklyrefuelSetMinDstGas) (event.Subscription, error) {

	logs, sub, err := _Merklyrefuel.contract.WatchLogs(opts, "SetMinDstGas")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerklyrefuelSetMinDstGas)
				if err := _Merklyrefuel.contract.UnpackLog(event, "SetMinDstGas", log); err != nil {
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
func (_Merklyrefuel *MerklyrefuelFilterer) ParseSetMinDstGas(log types.Log) (*MerklyrefuelSetMinDstGas, error) {
	event := new(MerklyrefuelSetMinDstGas)
	if err := _Merklyrefuel.contract.UnpackLog(event, "SetMinDstGas", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerklyrefuelSetPrecrimeIterator is returned from FilterSetPrecrime and is used to iterate over the raw logs and unpacked data for SetPrecrime events raised by the Merklyrefuel contract.
type MerklyrefuelSetPrecrimeIterator struct {
	Event *MerklyrefuelSetPrecrime // Event containing the contract specifics and raw log

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
func (it *MerklyrefuelSetPrecrimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerklyrefuelSetPrecrime)
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
		it.Event = new(MerklyrefuelSetPrecrime)
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
func (it *MerklyrefuelSetPrecrimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerklyrefuelSetPrecrimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerklyrefuelSetPrecrime represents a SetPrecrime event raised by the Merklyrefuel contract.
type MerklyrefuelSetPrecrime struct {
	Precrime common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPrecrime is a free log retrieval operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_Merklyrefuel *MerklyrefuelFilterer) FilterSetPrecrime(opts *bind.FilterOpts) (*MerklyrefuelSetPrecrimeIterator, error) {

	logs, sub, err := _Merklyrefuel.contract.FilterLogs(opts, "SetPrecrime")
	if err != nil {
		return nil, err
	}
	return &MerklyrefuelSetPrecrimeIterator{contract: _Merklyrefuel.contract, event: "SetPrecrime", logs: logs, sub: sub}, nil
}

// WatchSetPrecrime is a free log subscription operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_Merklyrefuel *MerklyrefuelFilterer) WatchSetPrecrime(opts *bind.WatchOpts, sink chan<- *MerklyrefuelSetPrecrime) (event.Subscription, error) {

	logs, sub, err := _Merklyrefuel.contract.WatchLogs(opts, "SetPrecrime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerklyrefuelSetPrecrime)
				if err := _Merklyrefuel.contract.UnpackLog(event, "SetPrecrime", log); err != nil {
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
func (_Merklyrefuel *MerklyrefuelFilterer) ParseSetPrecrime(log types.Log) (*MerklyrefuelSetPrecrime, error) {
	event := new(MerklyrefuelSetPrecrime)
	if err := _Merklyrefuel.contract.UnpackLog(event, "SetPrecrime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerklyrefuelSetTrustedRemoteIterator is returned from FilterSetTrustedRemote and is used to iterate over the raw logs and unpacked data for SetTrustedRemote events raised by the Merklyrefuel contract.
type MerklyrefuelSetTrustedRemoteIterator struct {
	Event *MerklyrefuelSetTrustedRemote // Event containing the contract specifics and raw log

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
func (it *MerklyrefuelSetTrustedRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerklyrefuelSetTrustedRemote)
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
		it.Event = new(MerklyrefuelSetTrustedRemote)
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
func (it *MerklyrefuelSetTrustedRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerklyrefuelSetTrustedRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerklyrefuelSetTrustedRemote represents a SetTrustedRemote event raised by the Merklyrefuel contract.
type MerklyrefuelSetTrustedRemote struct {
	RemoteChainId uint16
	Path          []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemote is a free log retrieval operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_Merklyrefuel *MerklyrefuelFilterer) FilterSetTrustedRemote(opts *bind.FilterOpts) (*MerklyrefuelSetTrustedRemoteIterator, error) {

	logs, sub, err := _Merklyrefuel.contract.FilterLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return &MerklyrefuelSetTrustedRemoteIterator{contract: _Merklyrefuel.contract, event: "SetTrustedRemote", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemote is a free log subscription operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_Merklyrefuel *MerklyrefuelFilterer) WatchSetTrustedRemote(opts *bind.WatchOpts, sink chan<- *MerklyrefuelSetTrustedRemote) (event.Subscription, error) {

	logs, sub, err := _Merklyrefuel.contract.WatchLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerklyrefuelSetTrustedRemote)
				if err := _Merklyrefuel.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
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
func (_Merklyrefuel *MerklyrefuelFilterer) ParseSetTrustedRemote(log types.Log) (*MerklyrefuelSetTrustedRemote, error) {
	event := new(MerklyrefuelSetTrustedRemote)
	if err := _Merklyrefuel.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerklyrefuelSetTrustedRemoteAddressIterator is returned from FilterSetTrustedRemoteAddress and is used to iterate over the raw logs and unpacked data for SetTrustedRemoteAddress events raised by the Merklyrefuel contract.
type MerklyrefuelSetTrustedRemoteAddressIterator struct {
	Event *MerklyrefuelSetTrustedRemoteAddress // Event containing the contract specifics and raw log

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
func (it *MerklyrefuelSetTrustedRemoteAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerklyrefuelSetTrustedRemoteAddress)
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
		it.Event = new(MerklyrefuelSetTrustedRemoteAddress)
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
func (it *MerklyrefuelSetTrustedRemoteAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerklyrefuelSetTrustedRemoteAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerklyrefuelSetTrustedRemoteAddress represents a SetTrustedRemoteAddress event raised by the Merklyrefuel contract.
type MerklyrefuelSetTrustedRemoteAddress struct {
	RemoteChainId uint16
	RemoteAddress []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemoteAddress is a free log retrieval operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_Merklyrefuel *MerklyrefuelFilterer) FilterSetTrustedRemoteAddress(opts *bind.FilterOpts) (*MerklyrefuelSetTrustedRemoteAddressIterator, error) {

	logs, sub, err := _Merklyrefuel.contract.FilterLogs(opts, "SetTrustedRemoteAddress")
	if err != nil {
		return nil, err
	}
	return &MerklyrefuelSetTrustedRemoteAddressIterator{contract: _Merklyrefuel.contract, event: "SetTrustedRemoteAddress", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemoteAddress is a free log subscription operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_Merklyrefuel *MerklyrefuelFilterer) WatchSetTrustedRemoteAddress(opts *bind.WatchOpts, sink chan<- *MerklyrefuelSetTrustedRemoteAddress) (event.Subscription, error) {

	logs, sub, err := _Merklyrefuel.contract.WatchLogs(opts, "SetTrustedRemoteAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerklyrefuelSetTrustedRemoteAddress)
				if err := _Merklyrefuel.contract.UnpackLog(event, "SetTrustedRemoteAddress", log); err != nil {
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
func (_Merklyrefuel *MerklyrefuelFilterer) ParseSetTrustedRemoteAddress(log types.Log) (*MerklyrefuelSetTrustedRemoteAddress, error) {
	event := new(MerklyrefuelSetTrustedRemoteAddress)
	if err := _Merklyrefuel.contract.UnpackLog(event, "SetTrustedRemoteAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
