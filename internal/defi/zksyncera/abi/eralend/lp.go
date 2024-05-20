// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eralend

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

// LpMetaData contains all meta data concerning the Lp contract.
var LpMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cashPrior\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"oldComptroller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"NewComptroller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"oldInterestRateModel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewMarketInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"RepayBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"benefactor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesReduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"compLikeDelegatee\",\"type\":\"address\"}],\"name\":\"_delegateCompLikeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying_\",\"type\":\"address\"},{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolSeizeShareMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEIP20NonStandardInterface\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LpABI is the input ABI used to generate the binding from.
// Deprecated: Use LpMetaData.ABI instead.
var LpABI = LpMetaData.ABI

// Lp is an auto generated Go binding around an Ethereum contract.
type Lp struct {
	LpCaller     // Read-only binding to the contract
	LpTransactor // Write-only binding to the contract
	LpFilterer   // Log filterer for contract events
}

// LpCaller is an auto generated read-only Go binding around an Ethereum contract.
type LpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LpSession struct {
	Contract     *Lp               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LpCallerSession struct {
	Contract *LpCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LpTransactorSession struct {
	Contract     *LpTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LpRaw is an auto generated low-level Go binding around an Ethereum contract.
type LpRaw struct {
	Contract *Lp // Generic contract binding to access the raw methods on
}

// LpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LpCallerRaw struct {
	Contract *LpCaller // Generic read-only contract binding to access the raw methods on
}

// LpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LpTransactorRaw struct {
	Contract *LpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLp creates a new instance of Lp, bound to a specific deployed contract.
func NewLp(address common.Address, backend bind.ContractBackend) (*Lp, error) {
	contract, err := bindLp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lp{LpCaller: LpCaller{contract: contract}, LpTransactor: LpTransactor{contract: contract}, LpFilterer: LpFilterer{contract: contract}}, nil
}

// NewLpCaller creates a new read-only instance of Lp, bound to a specific deployed contract.
func NewLpCaller(address common.Address, caller bind.ContractCaller) (*LpCaller, error) {
	contract, err := bindLp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LpCaller{contract: contract}, nil
}

// NewLpTransactor creates a new write-only instance of Lp, bound to a specific deployed contract.
func NewLpTransactor(address common.Address, transactor bind.ContractTransactor) (*LpTransactor, error) {
	contract, err := bindLp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LpTransactor{contract: contract}, nil
}

// NewLpFilterer creates a new log filterer instance of Lp, bound to a specific deployed contract.
func NewLpFilterer(address common.Address, filterer bind.ContractFilterer) (*LpFilterer, error) {
	contract, err := bindLp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LpFilterer{contract: contract}, nil
}

// bindLp binds a generic wrapper to an already deployed contract.
func bindLp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lp *LpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lp.Contract.LpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lp *LpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.Contract.LpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lp *LpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lp.Contract.LpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lp *LpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lp *LpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lp *LpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lp.Contract.contract.Transact(opts, method, params...)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Lp *LpCaller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "accrualBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Lp *LpSession) AccrualBlockNumber() (*big.Int, error) {
	return _Lp.Contract.AccrualBlockNumber(&_Lp.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Lp *LpCallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _Lp.Contract.AccrualBlockNumber(&_Lp.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Lp *LpCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Lp *LpSession) Admin() (common.Address, error) {
	return _Lp.Contract.Admin(&_Lp.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Lp *LpCallerSession) Admin() (common.Address, error) {
	return _Lp.Contract.Admin(&_Lp.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Lp *LpCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Lp *LpSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Lp.Contract.Allowance(&_Lp.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Lp *LpCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Lp.Contract.Allowance(&_Lp.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Lp *LpCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Lp *LpSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Lp.Contract.BalanceOf(&_Lp.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Lp *LpCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Lp.Contract.BalanceOf(&_Lp.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Lp *LpCaller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Lp *LpSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Lp.Contract.BorrowBalanceStored(&_Lp.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Lp *LpCallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Lp.Contract.BorrowBalanceStored(&_Lp.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Lp *LpCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Lp *LpSession) BorrowIndex() (*big.Int, error) {
	return _Lp.Contract.BorrowIndex(&_Lp.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Lp *LpCallerSession) BorrowIndex() (*big.Int, error) {
	return _Lp.Contract.BorrowIndex(&_Lp.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Lp *LpCaller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "borrowRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Lp *LpSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Lp.Contract.BorrowRatePerBlock(&_Lp.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Lp *LpCallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Lp.Contract.BorrowRatePerBlock(&_Lp.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Lp *LpCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Lp *LpSession) Comptroller() (common.Address, error) {
	return _Lp.Contract.Comptroller(&_Lp.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Lp *LpCallerSession) Comptroller() (common.Address, error) {
	return _Lp.Contract.Comptroller(&_Lp.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Lp *LpCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Lp *LpSession) Decimals() (uint8, error) {
	return _Lp.Contract.Decimals(&_Lp.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Lp *LpCallerSession) Decimals() (uint8, error) {
	return _Lp.Contract.Decimals(&_Lp.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Lp *LpCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Lp *LpSession) ExchangeRateStored() (*big.Int, error) {
	return _Lp.Contract.ExchangeRateStored(&_Lp.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Lp *LpCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _Lp.Contract.ExchangeRateStored(&_Lp.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Lp *LpCaller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getAccountSnapshot", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Lp *LpSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Lp.Contract.GetAccountSnapshot(&_Lp.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Lp *LpCallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Lp.Contract.GetAccountSnapshot(&_Lp.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Lp *LpCaller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Lp *LpSession) GetCash() (*big.Int, error) {
	return _Lp.Contract.GetCash(&_Lp.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Lp *LpCallerSession) GetCash() (*big.Int, error) {
	return _Lp.Contract.GetCash(&_Lp.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Lp *LpCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Lp *LpSession) InterestRateModel() (common.Address, error) {
	return _Lp.Contract.InterestRateModel(&_Lp.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Lp *LpCallerSession) InterestRateModel() (common.Address, error) {
	return _Lp.Contract.InterestRateModel(&_Lp.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_Lp *LpCaller) IsCToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "isCToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_Lp *LpSession) IsCToken() (bool, error) {
	return _Lp.Contract.IsCToken(&_Lp.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_Lp *LpCallerSession) IsCToken() (bool, error) {
	return _Lp.Contract.IsCToken(&_Lp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lp *LpCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lp *LpSession) Name() (string, error) {
	return _Lp.Contract.Name(&_Lp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lp *LpCallerSession) Name() (string, error) {
	return _Lp.Contract.Name(&_Lp.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Lp *LpCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Lp *LpSession) PendingAdmin() (common.Address, error) {
	return _Lp.Contract.PendingAdmin(&_Lp.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Lp *LpCallerSession) PendingAdmin() (common.Address, error) {
	return _Lp.Contract.PendingAdmin(&_Lp.CallOpts)
}

// ProtocolSeizeShareMantissa is a free data retrieval call binding the contract method 0x6752e702.
//
// Solidity: function protocolSeizeShareMantissa() view returns(uint256)
func (_Lp *LpCaller) ProtocolSeizeShareMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "protocolSeizeShareMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolSeizeShareMantissa is a free data retrieval call binding the contract method 0x6752e702.
//
// Solidity: function protocolSeizeShareMantissa() view returns(uint256)
func (_Lp *LpSession) ProtocolSeizeShareMantissa() (*big.Int, error) {
	return _Lp.Contract.ProtocolSeizeShareMantissa(&_Lp.CallOpts)
}

// ProtocolSeizeShareMantissa is a free data retrieval call binding the contract method 0x6752e702.
//
// Solidity: function protocolSeizeShareMantissa() view returns(uint256)
func (_Lp *LpCallerSession) ProtocolSeizeShareMantissa() (*big.Int, error) {
	return _Lp.Contract.ProtocolSeizeShareMantissa(&_Lp.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Lp *LpCaller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "reserveFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Lp *LpSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Lp.Contract.ReserveFactorMantissa(&_Lp.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Lp *LpCallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Lp.Contract.ReserveFactorMantissa(&_Lp.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Lp *LpCaller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Lp *LpSession) SupplyRatePerBlock() (*big.Int, error) {
	return _Lp.Contract.SupplyRatePerBlock(&_Lp.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Lp *LpCallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _Lp.Contract.SupplyRatePerBlock(&_Lp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lp *LpCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lp *LpSession) Symbol() (string, error) {
	return _Lp.Contract.Symbol(&_Lp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lp *LpCallerSession) Symbol() (string, error) {
	return _Lp.Contract.Symbol(&_Lp.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Lp *LpCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Lp *LpSession) TotalBorrows() (*big.Int, error) {
	return _Lp.Contract.TotalBorrows(&_Lp.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Lp *LpCallerSession) TotalBorrows() (*big.Int, error) {
	return _Lp.Contract.TotalBorrows(&_Lp.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Lp *LpCaller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Lp *LpSession) TotalReserves() (*big.Int, error) {
	return _Lp.Contract.TotalReserves(&_Lp.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Lp *LpCallerSession) TotalReserves() (*big.Int, error) {
	return _Lp.Contract.TotalReserves(&_Lp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lp *LpCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lp *LpSession) TotalSupply() (*big.Int, error) {
	return _Lp.Contract.TotalSupply(&_Lp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lp *LpCallerSession) TotalSupply() (*big.Int, error) {
	return _Lp.Contract.TotalSupply(&_Lp.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Lp *LpCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Lp *LpSession) Underlying() (common.Address, error) {
	return _Lp.Contract.Underlying(&_Lp.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Lp *LpCallerSession) Underlying() (common.Address, error) {
	return _Lp.Contract.Underlying(&_Lp.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Lp *LpTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Lp *LpSession) AcceptAdmin() (*types.Transaction, error) {
	return _Lp.Contract.AcceptAdmin(&_Lp.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Lp *LpTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Lp.Contract.AcceptAdmin(&_Lp.TransactOpts)
}

// DelegateCompLikeTo is a paid mutator transaction binding the contract method 0x7f1e06be.
//
// Solidity: function _delegateCompLikeTo(address compLikeDelegatee) returns()
func (_Lp *LpTransactor) DelegateCompLikeTo(opts *bind.TransactOpts, compLikeDelegatee common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "_delegateCompLikeTo", compLikeDelegatee)
}

// DelegateCompLikeTo is a paid mutator transaction binding the contract method 0x7f1e06be.
//
// Solidity: function _delegateCompLikeTo(address compLikeDelegatee) returns()
func (_Lp *LpSession) DelegateCompLikeTo(compLikeDelegatee common.Address) (*types.Transaction, error) {
	return _Lp.Contract.DelegateCompLikeTo(&_Lp.TransactOpts, compLikeDelegatee)
}

// DelegateCompLikeTo is a paid mutator transaction binding the contract method 0x7f1e06be.
//
// Solidity: function _delegateCompLikeTo(address compLikeDelegatee) returns()
func (_Lp *LpTransactorSession) DelegateCompLikeTo(compLikeDelegatee common.Address) (*types.Transaction, error) {
	return _Lp.Contract.DelegateCompLikeTo(&_Lp.TransactOpts, compLikeDelegatee)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Lp *LpTransactor) SetComptroller(opts *bind.TransactOpts, newComptroller common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "_setComptroller", newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Lp *LpSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Lp.Contract.SetComptroller(&_Lp.TransactOpts, newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Lp *LpTransactorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Lp.Contract.SetComptroller(&_Lp.TransactOpts, newComptroller)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Lp *LpTransactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Lp *LpSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Lp.Contract.SetInterestRateModel(&_Lp.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Lp *LpTransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Lp.Contract.SetInterestRateModel(&_Lp.TransactOpts, newInterestRateModel)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Lp *LpTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Lp *LpSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Lp.Contract.SetPendingAdmin(&_Lp.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Lp *LpTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Lp.Contract.SetPendingAdmin(&_Lp.TransactOpts, newPendingAdmin)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Lp *LpTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Lp *LpSession) AccrueInterest() (*types.Transaction, error) {
	return _Lp.Contract.AccrueInterest(&_Lp.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Lp *LpTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _Lp.Contract.AccrueInterest(&_Lp.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Lp *LpTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Lp *LpSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Approve(&_Lp.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Lp *LpTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Approve(&_Lp.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Lp *LpTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Lp *LpSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Lp.Contract.BalanceOfUnderlying(&_Lp.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Lp *LpTransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Lp.Contract.BalanceOfUnderlying(&_Lp.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Lp *LpTransactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Lp *LpSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Borrow(&_Lp.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Lp *LpTransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Borrow(&_Lp.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Lp *LpTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Lp *LpSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Lp.Contract.BorrowBalanceCurrent(&_Lp.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Lp *LpTransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Lp.Contract.BorrowBalanceCurrent(&_Lp.TransactOpts, account)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Lp *LpTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Lp *LpSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Lp.Contract.ExchangeRateCurrent(&_Lp.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Lp *LpTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Lp.Contract.ExchangeRateCurrent(&_Lp.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xafe65468.
//
// Solidity: function init(address underlying_, address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Lp *LpTransactor) Init(opts *bind.TransactOpts, underlying_ common.Address, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "init", underlying_, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Init is a paid mutator transaction binding the contract method 0xafe65468.
//
// Solidity: function init(address underlying_, address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Lp *LpSession) Init(underlying_ common.Address, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Lp.Contract.Init(&_Lp.TransactOpts, underlying_, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Init is a paid mutator transaction binding the contract method 0xafe65468.
//
// Solidity: function init(address underlying_, address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Lp *LpTransactorSession) Init(underlying_ common.Address, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Lp.Contract.Init(&_Lp.TransactOpts, underlying_, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Lp *LpTransactor) Initialize(opts *bind.TransactOpts, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "initialize", comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Lp *LpSession) Initialize(comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Lp.Contract.Initialize(&_Lp.TransactOpts, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Lp *LpTransactorSession) Initialize(comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Lp.Contract.Initialize(&_Lp.TransactOpts, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(uint256)
func (_Lp *LpTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(uint256)
func (_Lp *LpSession) Mint() (*types.Transaction, error) {
	return _Lp.Contract.Mint(&_Lp.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(uint256)
func (_Lp *LpTransactorSession) Mint() (*types.Transaction, error) {
	return _Lp.Contract.Mint(&_Lp.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Lp *LpTransactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Lp *LpSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Redeem(&_Lp.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Lp *LpTransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Redeem(&_Lp.TransactOpts, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Lp *LpTransactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Lp *LpSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.RedeemUnderlying(&_Lp.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Lp *LpTransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.RedeemUnderlying(&_Lp.TransactOpts, redeemAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Lp *LpTransactor) RepayBorrow(opts *bind.TransactOpts, repayAmount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "repayBorrow", repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Lp *LpSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.RepayBorrow(&_Lp.TransactOpts, repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Lp *LpTransactorSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.RepayBorrow(&_Lp.TransactOpts, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Lp *LpTransactor) RepayBorrowBehalf(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "repayBorrowBehalf", borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Lp *LpSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.RepayBorrowBehalf(&_Lp.TransactOpts, borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Lp *LpTransactorSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.RepayBorrowBehalf(&_Lp.TransactOpts, borrower, repayAmount)
}

// SweepToken is a paid mutator transaction binding the contract method 0x1be19560.
//
// Solidity: function sweepToken(address token) returns()
func (_Lp *LpTransactor) SweepToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "sweepToken", token)
}

// SweepToken is a paid mutator transaction binding the contract method 0x1be19560.
//
// Solidity: function sweepToken(address token) returns()
func (_Lp *LpSession) SweepToken(token common.Address) (*types.Transaction, error) {
	return _Lp.Contract.SweepToken(&_Lp.TransactOpts, token)
}

// SweepToken is a paid mutator transaction binding the contract method 0x1be19560.
//
// Solidity: function sweepToken(address token) returns()
func (_Lp *LpTransactorSession) SweepToken(token common.Address) (*types.Transaction, error) {
	return _Lp.Contract.SweepToken(&_Lp.TransactOpts, token)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Lp *LpTransactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Lp *LpSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Lp.Contract.TotalBorrowsCurrent(&_Lp.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Lp *LpTransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Lp.Contract.TotalBorrowsCurrent(&_Lp.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Lp *LpTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Lp *LpSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Transfer(&_Lp.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Lp *LpTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Transfer(&_Lp.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Lp *LpTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Lp *LpSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.TransferFrom(&_Lp.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Lp *LpTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.TransferFrom(&_Lp.TransactOpts, src, dst, amount)
}

// LpAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the Lp contract.
type LpAccrueInterestIterator struct {
	Event *LpAccrueInterest // Event containing the contract specifics and raw log

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
func (it *LpAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpAccrueInterest)
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
		it.Event = new(LpAccrueInterest)
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
func (it *LpAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpAccrueInterest represents a AccrueInterest event raised by the Lp contract.
type LpAccrueInterest struct {
	CashPrior           *big.Int
	InterestAccumulated *big.Int
	BorrowIndex         *big.Int
	TotalBorrows        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Lp *LpFilterer) FilterAccrueInterest(opts *bind.FilterOpts) (*LpAccrueInterestIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &LpAccrueInterestIterator{contract: _Lp.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Lp *LpFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *LpAccrueInterest) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpAccrueInterest)
				if err := _Lp.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
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

// ParseAccrueInterest is a log parse operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Lp *LpFilterer) ParseAccrueInterest(log types.Log) (*LpAccrueInterest, error) {
	event := new(LpAccrueInterest)
	if err := _Lp.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Lp contract.
type LpApprovalIterator struct {
	Event *LpApproval // Event containing the contract specifics and raw log

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
func (it *LpApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpApproval)
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
		it.Event = new(LpApproval)
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
func (it *LpApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpApproval represents a Approval event raised by the Lp contract.
type LpApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Lp *LpFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LpApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LpApprovalIterator{contract: _Lp.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Lp *LpFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LpApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpApproval)
				if err := _Lp.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Lp *LpFilterer) ParseApproval(log types.Log) (*LpApproval, error) {
	event := new(LpApproval)
	if err := _Lp.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Lp contract.
type LpBorrowIterator struct {
	Event *LpBorrow // Event containing the contract specifics and raw log

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
func (it *LpBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpBorrow)
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
		it.Event = new(LpBorrow)
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
func (it *LpBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpBorrow represents a Borrow event raised by the Lp contract.
type LpBorrow struct {
	Borrower       common.Address
	BorrowAmount   *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Lp *LpFilterer) FilterBorrow(opts *bind.FilterOpts) (*LpBorrowIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &LpBorrowIterator{contract: _Lp.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Lp *LpFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *LpBorrow) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpBorrow)
				if err := _Lp.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Lp *LpFilterer) ParseBorrow(log types.Log) (*LpBorrow, error) {
	event := new(LpBorrow)
	if err := _Lp.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Lp contract.
type LpFailureIterator struct {
	Event *LpFailure // Event containing the contract specifics and raw log

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
func (it *LpFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpFailure)
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
		it.Event = new(LpFailure)
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
func (it *LpFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpFailure represents a Failure event raised by the Lp contract.
type LpFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Lp *LpFilterer) FilterFailure(opts *bind.FilterOpts) (*LpFailureIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &LpFailureIterator{contract: _Lp.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Lp *LpFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *LpFailure) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpFailure)
				if err := _Lp.contract.UnpackLog(event, "Failure", log); err != nil {
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

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Lp *LpFilterer) ParseFailure(log types.Log) (*LpFailure, error) {
	event := new(LpFailure)
	if err := _Lp.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpLiquidateBorrowIterator is returned from FilterLiquidateBorrow and is used to iterate over the raw logs and unpacked data for LiquidateBorrow events raised by the Lp contract.
type LpLiquidateBorrowIterator struct {
	Event *LpLiquidateBorrow // Event containing the contract specifics and raw log

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
func (it *LpLiquidateBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpLiquidateBorrow)
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
		it.Event = new(LpLiquidateBorrow)
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
func (it *LpLiquidateBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpLiquidateBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpLiquidateBorrow represents a LiquidateBorrow event raised by the Lp contract.
type LpLiquidateBorrow struct {
	Liquidator       common.Address
	Borrower         common.Address
	RepayAmount      *big.Int
	CTokenCollateral common.Address
	SeizeTokens      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidateBorrow is a free log retrieval operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_Lp *LpFilterer) FilterLiquidateBorrow(opts *bind.FilterOpts) (*LpLiquidateBorrowIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return &LpLiquidateBorrowIterator{contract: _Lp.contract, event: "LiquidateBorrow", logs: logs, sub: sub}, nil
}

// WatchLiquidateBorrow is a free log subscription operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_Lp *LpFilterer) WatchLiquidateBorrow(opts *bind.WatchOpts, sink chan<- *LpLiquidateBorrow) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpLiquidateBorrow)
				if err := _Lp.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
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

// ParseLiquidateBorrow is a log parse operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_Lp *LpFilterer) ParseLiquidateBorrow(log types.Log) (*LpLiquidateBorrow, error) {
	event := new(LpLiquidateBorrow)
	if err := _Lp.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Lp contract.
type LpMintIterator struct {
	Event *LpMint // Event containing the contract specifics and raw log

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
func (it *LpMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpMint)
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
		it.Event = new(LpMint)
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
func (it *LpMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpMint represents a Mint event raised by the Lp contract.
type LpMint struct {
	Minter     common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Lp *LpFilterer) FilterMint(opts *bind.FilterOpts) (*LpMintIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &LpMintIterator{contract: _Lp.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Lp *LpFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *LpMint) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpMint)
				if err := _Lp.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Lp *LpFilterer) ParseMint(log types.Log) (*LpMint, error) {
	event := new(LpMint)
	if err := _Lp.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Lp contract.
type LpNewAdminIterator struct {
	Event *LpNewAdmin // Event containing the contract specifics and raw log

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
func (it *LpNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpNewAdmin)
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
		it.Event = new(LpNewAdmin)
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
func (it *LpNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpNewAdmin represents a NewAdmin event raised by the Lp contract.
type LpNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Lp *LpFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*LpNewAdminIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &LpNewAdminIterator{contract: _Lp.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Lp *LpFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *LpNewAdmin) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpNewAdmin)
				if err := _Lp.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Lp *LpFilterer) ParseNewAdmin(log types.Log) (*LpNewAdmin, error) {
	event := new(LpNewAdmin)
	if err := _Lp.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpNewComptrollerIterator is returned from FilterNewComptroller and is used to iterate over the raw logs and unpacked data for NewComptroller events raised by the Lp contract.
type LpNewComptrollerIterator struct {
	Event *LpNewComptroller // Event containing the contract specifics and raw log

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
func (it *LpNewComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpNewComptroller)
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
		it.Event = new(LpNewComptroller)
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
func (it *LpNewComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpNewComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpNewComptroller represents a NewComptroller event raised by the Lp contract.
type LpNewComptroller struct {
	OldComptroller common.Address
	NewComptroller common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewComptroller is a free log retrieval operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Lp *LpFilterer) FilterNewComptroller(opts *bind.FilterOpts) (*LpNewComptrollerIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return &LpNewComptrollerIterator{contract: _Lp.contract, event: "NewComptroller", logs: logs, sub: sub}, nil
}

// WatchNewComptroller is a free log subscription operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Lp *LpFilterer) WatchNewComptroller(opts *bind.WatchOpts, sink chan<- *LpNewComptroller) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpNewComptroller)
				if err := _Lp.contract.UnpackLog(event, "NewComptroller", log); err != nil {
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

// ParseNewComptroller is a log parse operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Lp *LpFilterer) ParseNewComptroller(log types.Log) (*LpNewComptroller, error) {
	event := new(LpNewComptroller)
	if err := _Lp.contract.UnpackLog(event, "NewComptroller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpNewMarketInterestRateModelIterator is returned from FilterNewMarketInterestRateModel and is used to iterate over the raw logs and unpacked data for NewMarketInterestRateModel events raised by the Lp contract.
type LpNewMarketInterestRateModelIterator struct {
	Event *LpNewMarketInterestRateModel // Event containing the contract specifics and raw log

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
func (it *LpNewMarketInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpNewMarketInterestRateModel)
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
		it.Event = new(LpNewMarketInterestRateModel)
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
func (it *LpNewMarketInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpNewMarketInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpNewMarketInterestRateModel represents a NewMarketInterestRateModel event raised by the Lp contract.
type LpNewMarketInterestRateModel struct {
	OldInterestRateModel common.Address
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewMarketInterestRateModel is a free log retrieval operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Lp *LpFilterer) FilterNewMarketInterestRateModel(opts *bind.FilterOpts) (*LpNewMarketInterestRateModelIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &LpNewMarketInterestRateModelIterator{contract: _Lp.contract, event: "NewMarketInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewMarketInterestRateModel is a free log subscription operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Lp *LpFilterer) WatchNewMarketInterestRateModel(opts *bind.WatchOpts, sink chan<- *LpNewMarketInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpNewMarketInterestRateModel)
				if err := _Lp.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
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

// ParseNewMarketInterestRateModel is a log parse operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Lp *LpFilterer) ParseNewMarketInterestRateModel(log types.Log) (*LpNewMarketInterestRateModel, error) {
	event := new(LpNewMarketInterestRateModel)
	if err := _Lp.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the Lp contract.
type LpNewPendingAdminIterator struct {
	Event *LpNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *LpNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpNewPendingAdmin)
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
		it.Event = new(LpNewPendingAdmin)
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
func (it *LpNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpNewPendingAdmin represents a NewPendingAdmin event raised by the Lp contract.
type LpNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Lp *LpFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*LpNewPendingAdminIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &LpNewPendingAdminIterator{contract: _Lp.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Lp *LpFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *LpNewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpNewPendingAdmin)
				if err := _Lp.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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

// ParseNewPendingAdmin is a log parse operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Lp *LpFilterer) ParseNewPendingAdmin(log types.Log) (*LpNewPendingAdmin, error) {
	event := new(LpNewPendingAdmin)
	if err := _Lp.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Lp contract.
type LpRedeemIterator struct {
	Event *LpRedeem // Event containing the contract specifics and raw log

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
func (it *LpRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpRedeem)
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
		it.Event = new(LpRedeem)
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
func (it *LpRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpRedeem represents a Redeem event raised by the Lp contract.
type LpRedeem struct {
	Redeemer     common.Address
	RedeemAmount *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Lp *LpFilterer) FilterRedeem(opts *bind.FilterOpts) (*LpRedeemIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &LpRedeemIterator{contract: _Lp.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Lp *LpFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *LpRedeem) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpRedeem)
				if err := _Lp.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Lp *LpFilterer) ParseRedeem(log types.Log) (*LpRedeem, error) {
	event := new(LpRedeem)
	if err := _Lp.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpRepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the Lp contract.
type LpRepayBorrowIterator struct {
	Event *LpRepayBorrow // Event containing the contract specifics and raw log

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
func (it *LpRepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpRepayBorrow)
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
		it.Event = new(LpRepayBorrow)
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
func (it *LpRepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpRepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpRepayBorrow represents a RepayBorrow event raised by the Lp contract.
type LpRepayBorrow struct {
	Payer          common.Address
	Borrower       common.Address
	RepayAmount    *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepayBorrow is a free log retrieval operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Lp *LpFilterer) FilterRepayBorrow(opts *bind.FilterOpts) (*LpRepayBorrowIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &LpRepayBorrowIterator{contract: _Lp.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Lp *LpFilterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *LpRepayBorrow) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpRepayBorrow)
				if err := _Lp.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
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

// ParseRepayBorrow is a log parse operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Lp *LpFilterer) ParseRepayBorrow(log types.Log) (*LpRepayBorrow, error) {
	event := new(LpRepayBorrow)
	if err := _Lp.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpReservesAddedIterator is returned from FilterReservesAdded and is used to iterate over the raw logs and unpacked data for ReservesAdded events raised by the Lp contract.
type LpReservesAddedIterator struct {
	Event *LpReservesAdded // Event containing the contract specifics and raw log

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
func (it *LpReservesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpReservesAdded)
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
		it.Event = new(LpReservesAdded)
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
func (it *LpReservesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpReservesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpReservesAdded represents a ReservesAdded event raised by the Lp contract.
type LpReservesAdded struct {
	Benefactor       common.Address
	AddAmount        *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesAdded is a free log retrieval operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Lp *LpFilterer) FilterReservesAdded(opts *bind.FilterOpts) (*LpReservesAddedIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return &LpReservesAddedIterator{contract: _Lp.contract, event: "ReservesAdded", logs: logs, sub: sub}, nil
}

// WatchReservesAdded is a free log subscription operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Lp *LpFilterer) WatchReservesAdded(opts *bind.WatchOpts, sink chan<- *LpReservesAdded) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpReservesAdded)
				if err := _Lp.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
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

// ParseReservesAdded is a log parse operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Lp *LpFilterer) ParseReservesAdded(log types.Log) (*LpReservesAdded, error) {
	event := new(LpReservesAdded)
	if err := _Lp.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpReservesReducedIterator is returned from FilterReservesReduced and is used to iterate over the raw logs and unpacked data for ReservesReduced events raised by the Lp contract.
type LpReservesReducedIterator struct {
	Event *LpReservesReduced // Event containing the contract specifics and raw log

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
func (it *LpReservesReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpReservesReduced)
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
		it.Event = new(LpReservesReduced)
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
func (it *LpReservesReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpReservesReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpReservesReduced represents a ReservesReduced event raised by the Lp contract.
type LpReservesReduced struct {
	Admin            common.Address
	ReduceAmount     *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesReduced is a free log retrieval operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Lp *LpFilterer) FilterReservesReduced(opts *bind.FilterOpts) (*LpReservesReducedIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return &LpReservesReducedIterator{contract: _Lp.contract, event: "ReservesReduced", logs: logs, sub: sub}, nil
}

// WatchReservesReduced is a free log subscription operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Lp *LpFilterer) WatchReservesReduced(opts *bind.WatchOpts, sink chan<- *LpReservesReduced) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpReservesReduced)
				if err := _Lp.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
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

// ParseReservesReduced is a log parse operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Lp *LpFilterer) ParseReservesReduced(log types.Log) (*LpReservesReduced, error) {
	event := new(LpReservesReduced)
	if err := _Lp.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Lp contract.
type LpTransferIterator struct {
	Event *LpTransfer // Event containing the contract specifics and raw log

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
func (it *LpTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpTransfer)
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
		it.Event = new(LpTransfer)
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
func (it *LpTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpTransfer represents a Transfer event raised by the Lp contract.
type LpTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Lp *LpFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LpTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LpTransferIterator{contract: _Lp.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Lp *LpFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LpTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpTransfer)
				if err := _Lp.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Lp *LpFilterer) ParseTransfer(log types.Log) (*LpTransfer, error) {
	event := new(LpTransfer)
	if err := _Lp.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
