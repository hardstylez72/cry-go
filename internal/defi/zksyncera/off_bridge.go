package zksyncera

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/contracts/ethtoken"
	"github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
)

// транзакция через офф бридж https://explorer.zksync.io/tx/0x8e6f82d93aac1af142d3e3fd6e58d54b6df294109b5f3b364ba7710f2eea7c99
// транзакция через софт https://explorer.zksync.io/tx/0xf471fea71aa35a1fc0716b40a41aaf286d186df1b5220ce13d16fd4a4b53591b
func (c *Client) BridgeToEthereumNetwork(ctx context.Context, req *L1L2BridgeReq) (*L1L2BridgeRes, error) {

	w, wtx, err := c.Wallet(req.WalletPK)
	if err != nil {
		return nil, errors.Wrap(err, "zksync2.NewWallet")
	}

	ethTokenAbi, err := abi.JSON(strings.NewReader(ethtoken.IEthTokenMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load ethTokenAbi: %w", err)
	}

	data, err := ethTokenAbi.Pack("withdraw", w.GetAddress())
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw function: %w", err)
	}
	tx := CreateFunctionCallTransaction(
		w.Address(),
		utils.L2EthTokenAddress,
		big.NewInt(0),
		big.NewInt(0),
		req.Amount,
		data,
		nil, nil,
	)

	result := &L1L2BridgeRes{}
	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, wtx.Signer)
	if err != nil {
		return nil, errors.Wrap(err, "Make712Tx")
	}

	result.EstimatedGasCost = estimate

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.ClientL2.SendRawTransaction(ctx, raw)
	if err != nil {
		return nil, errors.Wrap(err, "rpcL2.SendRawTransaction")
	}

	return &L1L2BridgeRes{
		TxHash: c.NewTx(hash, defi.CodeContract, nil),
	}, nil
}

// сайт https://etherscan.io/tx/0x8d05dfa50a8be054b5c27d50f36e4076a1f8eda1185fd23e12eb365a71bf50a3
// код
func (c *Client) BridgeFromEthereumNetwork(ctx context.Context, req *L1L2BridgeReq) (*L1L2BridgeRes, error) {

	w, _, err := c.Wallet(req.WalletPK)
	if err != nil {
		return nil, errors.Wrap(err, "zksync2.NewWallet")
	}

	priority := big.NewInt(5e7) //1.5GWEI

	msg := accounts.DepositCallMsg{
		To:                w.Address(),
		Token:             utils.EthAddress,
		Amount:            req.Amount,
		OperatorTip:       nil,
		BridgeAddress:     nil,
		L2GasLimit:        nil,
		GasPerPubdataByte: utils.RequiredL1ToL2GasPerPubdataLimit,
		RefundRecipient:   w.Address(),
		CustomBridgeData:  nil,
		Value:             nil,
		Gas:               0,
		GasPrice:          nil,
		GasFeeCap:         nil,
		GasTipCap:         priority,
		AccessList:        nil,
	}

	// L2 gas
	to := w.Address()
	gasLimitL2, err := c.ClientL2.EstimateL1ToL2Execute(ctx, types.CallMsg{
		CallMsg: ethereum.CallMsg{
			From:  w.Address(),
			To:    &to,
			Value: req.Amount,
		},
		Meta: &types.Eip712Meta{
			GasPerPubdata: utils.NewBig(msg.GasPerPubdataByte.Int64()),
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "EstimateL1ToL2Execute")
	}

	gasLimitL2 *= 3

	l2GasPrice, err := c.ClientL2.SuggestGasPrice(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "SuggestGasPrice")
	}

	msg.L2GasLimit = big.NewInt(0).SetUint64(gasLimitL2)

	gasL2 := defi.MinerGasLegacy(l2GasPrice, gasLimitL2)

	result := &L1L2BridgeRes{}
	var gas uint64
	var gasPrice *big.Int
	if !req.Gas.RuleSet() {

		// L1 gas
		header, err := c.ClientL1.HeaderByNumber(ctx, nil)
		if err != nil {
			return nil, errors.Wrap(err, "SuggestGasPrice")
		}

		msg.GasFeeCap = bozdo.BigIntSum(header.BaseFee, bozdo.Percent(header.BaseFee, 50))

		gasPrice = big.NewInt(0).Add(priority, msg.GasFeeCap)
		gasLimitL1, err := w.EstimateGasDeposit(ctx, msg)
		if err != nil {
			return nil, errors.Wrap(err, "EstimateGasDeposit")
		}
		gas = gasLimitL1 * 8

	} else {

		gasPrice = big.NewInt(0).Sub(&req.Gas.GasPrice, priority)
		gas = req.Gas.GasLimit.Uint64()
	}

	totalGas := bozdo.BigIntSum(defi.MinerGas(gasPrice, priority, gas), gasL2)
	l1GasLimit := gas
	gasPriceSpecial := big.NewInt(0).Div(totalGas, big.NewInt(0).SetUint64(l1GasLimit))

	result.EstimatedGasCost = &bozdo.EstimatedGasCost{
		Type:        bozdo.TxTypeDynamic,
		Name:        "deposit",
		GasLimit:    big.NewInt(0).SetUint64(l1GasLimit),
		GasPrice:    gasPriceSpecial,
		TotalGasWei: totalGas,
		Details:     nil,
	}

	if req.EstimateOnly {
		return result, nil
	}

	msg.Gas = gas
	msg.L2GasLimit = big.NewInt(0).SetUint64(gasLimitL2)

	l1Gas := big.NewInt(0).Sub(totalGas, gasL2)
	gasPrice = big.NewInt(0).Div(l1Gas, big.NewInt(0).SetUint64(gas))
	gasPrice = big.NewInt(0).Sub(gasPrice, priority)

	req.Amount = bozdo.Percent(req.Amount, 99)

	opt := &accounts.TransactOpts{
		Nonce:     nil,
		Value:     bozdo.BigIntSum(req.Amount, gasL2),
		GasPrice:  nil,
		GasFeeCap: gasPrice,
		GasTipCap: priority,
		GasLimit:  gas,
		Context:   ctx,
	}

	// 86433205000000
	// 3378824906437992

	b, _ := c.ClientL1.BalanceAt(ctx, w.Address(), nil)

	println("v-am: " + big.NewInt(0).Sub(opt.Value, req.Amount).String())
	println("b-v: " + big.NewInt(0).Sub(b, opt.Value).String())
	println("b-tx-fee: " + big.NewInt(0).Sub(b, defi.MinerGas(opt.GasFeeCap, opt.GasTipCap, opt.GasLimit)).String())
	println("diff: " + big.NewInt(0).Sub(b, bozdo.BigIntSum(opt.Value, defi.MinerGas(opt.GasFeeCap, opt.GasTipCap, opt.GasLimit))).String())

	tx, err := w.Deposit(opt, msg.ToDepositTransaction())
	if err != nil {
		return nil, errors.Wrap(err, "Deposit")
	}

	result.TxHash = c.NewTx(tx.Hash(), defi.CodeContract, nil)

	return &L1L2BridgeRes{
		TxHash: &bozdo.Transaction{
			Code:    defi.CodeContract,
			Network: c.Cfg.Network,
			Id:      tx.Hash().String(),
			Url:     "https://etherscan.io/tx/" + tx.Hash().String(),
			Details: nil,
		},
	}, nil
}
