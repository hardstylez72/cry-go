// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancake

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

// IQuoterV2QuoteExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuoterV2QuoteExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	AmountIn          *big.Int
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IQuoterV2QuoteExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuoterV2QuoteExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Amount            *big.Int
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
}

// QuoterMetaData contains all meta data concerning the Quoter contract.
var QuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_deployer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"pancakeV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"quoteExactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// QuoterABI is the input ABI used to generate the binding from.
// Deprecated: Use QuoterMetaData.ABI instead.
var QuoterABI = QuoterMetaData.ABI

// Quoter is an auto generated Go binding around an Ethereum contract.
type Quoter struct {
	QuoterCaller     // Read-only binding to the contract
	QuoterTransactor // Write-only binding to the contract
	QuoterFilterer   // Log filterer for contract events
}

// QuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuoterSession struct {
	Contract     *Quoter           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuoterCallerSession struct {
	Contract *QuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// QuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuoterTransactorSession struct {
	Contract     *QuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuoterRaw struct {
	Contract *Quoter // Generic contract binding to access the raw methods on
}

// QuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuoterCallerRaw struct {
	Contract *QuoterCaller // Generic read-only contract binding to access the raw methods on
}

// QuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuoterTransactorRaw struct {
	Contract *QuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuoter creates a new instance of Quoter, bound to a specific deployed contract.
func NewQuoter(address common.Address, backend bind.ContractBackend) (*Quoter, error) {
	contract, err := bindQuoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Quoter{QuoterCaller: QuoterCaller{contract: contract}, QuoterTransactor: QuoterTransactor{contract: contract}, QuoterFilterer: QuoterFilterer{contract: contract}}, nil
}

// NewQuoterCaller creates a new read-only instance of Quoter, bound to a specific deployed contract.
func NewQuoterCaller(address common.Address, caller bind.ContractCaller) (*QuoterCaller, error) {
	contract, err := bindQuoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuoterCaller{contract: contract}, nil
}

// NewQuoterTransactor creates a new write-only instance of Quoter, bound to a specific deployed contract.
func NewQuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*QuoterTransactor, error) {
	contract, err := bindQuoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuoterTransactor{contract: contract}, nil
}

// NewQuoterFilterer creates a new log filterer instance of Quoter, bound to a specific deployed contract.
func NewQuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*QuoterFilterer, error) {
	contract, err := bindQuoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuoterFilterer{contract: contract}, nil
}

// bindQuoter binds a generic wrapper to an already deployed contract.
func bindQuoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quoter *QuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quoter.Contract.QuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quoter *QuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quoter.Contract.QuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quoter *QuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quoter.Contract.QuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quoter *QuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quoter *QuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quoter *QuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quoter.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Quoter *QuoterCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Quoter *QuoterSession) WETH9() (common.Address, error) {
	return _Quoter.Contract.WETH9(&_Quoter.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Quoter *QuoterCallerSession) WETH9() (common.Address, error) {
	return _Quoter.Contract.WETH9(&_Quoter.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Quoter *QuoterCaller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Quoter *QuoterSession) Deployer() (common.Address, error) {
	return _Quoter.Contract.Deployer(&_Quoter.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Quoter *QuoterCallerSession) Deployer() (common.Address, error) {
	return _Quoter.Contract.Deployer(&_Quoter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Quoter *QuoterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Quoter *QuoterSession) Factory() (common.Address, error) {
	return _Quoter.Contract.Factory(&_Quoter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Quoter *QuoterCallerSession) Factory() (common.Address, error) {
	return _Quoter.Contract.Factory(&_Quoter.CallOpts)
}

// PancakeV3SwapCallback is a free data retrieval call binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_Quoter *QuoterCaller) PancakeV3SwapCallback(opts *bind.CallOpts, amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "pancakeV3SwapCallback", amount0Delta, amount1Delta, path)

	if err != nil {
		return err
	}

	return err

}

// PancakeV3SwapCallback is a free data retrieval call binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_Quoter *QuoterSession) PancakeV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _Quoter.Contract.PancakeV3SwapCallback(&_Quoter.CallOpts, amount0Delta, amount1Delta, path)
}

// PancakeV3SwapCallback is a free data retrieval call binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_Quoter *QuoterCallerSession) PancakeV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _Quoter.Contract.PancakeV3SwapCallback(&_Quoter.CallOpts, amount0Delta, amount1Delta, path)
}

// QuoteExactInput is a free data retrieval call binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) view returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Quoter *QuoterCaller) QuoteExactInput(opts *bind.CallOpts, path []byte, amountIn *big.Int) (struct {
	AmountOut                   *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "quoteExactInput", path, amountIn)

	outstruct := new(struct {
		AmountOut                   *big.Int
		SqrtPriceX96AfterList       []*big.Int
		InitializedTicksCrossedList []uint32
		GasEstimate                 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SqrtPriceX96AfterList = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.InitializedTicksCrossedList = *abi.ConvertType(out[2], new([]uint32)).(*[]uint32)
	outstruct.GasEstimate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuoteExactInput is a free data retrieval call binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) view returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Quoter *QuoterSession) QuoteExactInput(path []byte, amountIn *big.Int) (struct {
	AmountOut                   *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}, error) {
	return _Quoter.Contract.QuoteExactInput(&_Quoter.CallOpts, path, amountIn)
}

// QuoteExactInput is a free data retrieval call binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) view returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Quoter *QuoterCallerSession) QuoteExactInput(path []byte, amountIn *big.Int) (struct {
	AmountOut                   *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}, error) {
	return _Quoter.Contract.QuoteExactInput(&_Quoter.CallOpts, path, amountIn)
}

// QuoteExactInputSingle is a free data retrieval call binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) view returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Quoter *QuoterCaller) QuoteExactInputSingle(opts *bind.CallOpts, params IQuoterV2QuoteExactInputSingleParams) (struct {
	AmountOut               *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "quoteExactInputSingle", params)

	outstruct := new(struct {
		AmountOut               *big.Int
		SqrtPriceX96After       *big.Int
		InitializedTicksCrossed uint32
		GasEstimate             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SqrtPriceX96After = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InitializedTicksCrossed = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.GasEstimate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuoteExactInputSingle is a free data retrieval call binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) view returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Quoter *QuoterSession) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (struct {
	AmountOut               *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}, error) {
	return _Quoter.Contract.QuoteExactInputSingle(&_Quoter.CallOpts, params)
}

// QuoteExactInputSingle is a free data retrieval call binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) view returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Quoter *QuoterCallerSession) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (struct {
	AmountOut               *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}, error) {
	return _Quoter.Contract.QuoteExactInputSingle(&_Quoter.CallOpts, params)
}

// QuoteExactOutputSingle is a free data retrieval call binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) view returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Quoter *QuoterCaller) QuoteExactOutputSingle(opts *bind.CallOpts, params IQuoterV2QuoteExactOutputSingleParams) (struct {
	AmountIn                *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "quoteExactOutputSingle", params)

	outstruct := new(struct {
		AmountIn                *big.Int
		SqrtPriceX96After       *big.Int
		InitializedTicksCrossed uint32
		GasEstimate             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountIn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SqrtPriceX96After = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InitializedTicksCrossed = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.GasEstimate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuoteExactOutputSingle is a free data retrieval call binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) view returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Quoter *QuoterSession) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (struct {
	AmountIn                *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}, error) {
	return _Quoter.Contract.QuoteExactOutputSingle(&_Quoter.CallOpts, params)
}

// QuoteExactOutputSingle is a free data retrieval call binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) view returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Quoter *QuoterCallerSession) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (struct {
	AmountIn                *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}, error) {
	return _Quoter.Contract.QuoteExactOutputSingle(&_Quoter.CallOpts, params)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Quoter *QuoterTransactor) QuoteExactOutput(opts *bind.TransactOpts, path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "quoteExactOutput", path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Quoter *QuoterSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.QuoteExactOutput(&_Quoter.TransactOpts, path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Quoter *QuoterTransactorSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.QuoteExactOutput(&_Quoter.TransactOpts, path, amountOut)
}
