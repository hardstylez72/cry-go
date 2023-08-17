package zksyncera

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/contracts/ethtoken"
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

	msg := accounts.DepositCallMsg{
		To:                w.Address(),
		Token:             utils.EthAddress,
		Amount:            req.Amount,
		OperatorTip:       nil,
		BridgeAddress:     nil,
		L2GasLimit:        nil,
		GasPerPubdataByte: nil,
		RefundRecipient:   w.Address(),
		CustomBridgeData:  nil,
		Value:             req.Amount,
		Gas:               0,
		GasPrice:          nil,
		GasFeeCap:         nil,
		GasTipCap:         nil,
		AccessList:        nil,
	}

	result := &L1L2BridgeRes{}
	var gas uint64
	var gasPrice *big.Int
	if !req.Gas.RuleSet() {

		header, err := c.ClientL1.HeaderByNumber(ctx, nil)
		if err != nil {
			return nil, errors.Wrap(err, "SuggestGasPrice")
		}

		msg.GasFeeCap = bozdo.BigIntSum(header.BaseFee, bozdo.Percent(header.BaseFee, 50))

		gasPrice = msg.GasFeeCap
		gas, err = w.EstimateGasDeposit(ctx, msg)
		if err != nil {
			return nil, errors.Wrap(err, "EstimateGasDeposit")
		}

		gas *= 5

	} else {
		gasPrice = &req.Gas.GasPrice
		gas = req.Gas.GasLimit.Uint64()
	}

	result.EstimatedGasCost = &bozdo.EstimatedGasCost{
		Type:        bozdo.TxTypeLegacy,
		Name:        "deposit",
		GasLimit:    big.NewInt(0).SetUint64(gas),
		GasPrice:    gasPrice,
		TotalGasWei: defi.MinerGasLegacy(gasPrice, gas),
		Details:     nil,
	}

	if req.EstimateOnly {
		return result, nil
	}

	msg.GasPrice = gasPrice
	msg.Gas = gas

	tx, err := w.Deposit(nil, msg.ToDepositTransaction())
	if err != nil {
		return nil, errors.Wrap(err, "Deposit")
	}

	result.TxHash = c.NewTx(tx.Hash(), defi.CodeContract, nil)

	return &L1L2BridgeRes{
		TxHash: c.NewTx(tx.Hash(), defi.CodeContract, nil),
	}, nil
}
