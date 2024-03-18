package defi

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/contracts/optimism_fee"
	"github.com/hardstylez72/cry/internal/log"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"go.uber.org/zap"
)

type London func(ctx context.Context, c *EtheriumClient, opt *TxOpt, data *bozdo.TxData) (*bozdo.DefaultRes, error)

type TxOpt struct {
	NoSend       bool
	Pk           string
	Gas          *bozdo.Gas
	ExchangeRate *float64
	Debug        bool
	TaskType     v1.TaskType
}

func LondonReadyTx(ctx context.Context, c *EtheriumClient, opt *TxOpt, data *bozdo.TxData) (*bozdo.DefaultRes, error) {
	wt, err := NewWalletTransactor(opt.Pk)
	if err != nil {
		return nil, err
	}

	nonce, err := c.Cli.NonceAt(ctx, wt.WalletAddr, nil)
	if err != nil {
		return nil, err
	}

	dynamic := types.DynamicFeeTx{
		ChainID: c.Cfg.NetworkId,
		Nonce:   nonce,
		To:      &data.ContractAddr,
		Value:   data.Value,
		Data:    data.Data,
	}

	gasTipCap := big.NewInt(1e6)

	l1Fee, err := c.Cfg.EstimateL1Gas(ctx, data.Data)
	if err != nil {
		return nil, err
	}

	if opt.Debug {
		log.Log.Debug("0x" + hex.EncodeToString(data.Data))
	}

	if opt.Debug {
		log.Log.Debug("l1Fee "+opt.TaskType.String(), zap.String("l1Fee", WEIToEther(l1Fee).String()+" ETH"))
	}

	dynamic.Value = data.Value

	if opt.NoSend || !opt.Gas.RuleSet() {

		header, err := c.Cli.HeaderByNumber(ctx, nil)
		if err != nil {
			return nil, err
		}

		dynamic.GasTipCap = gasTipCap

		dynamic.GasFeeCap = bozdo.BigIntSum(header.BaseFee, header.BaseFee, dynamic.GasTipCap)

		gas, err := c.Cli.EstimateGas(ctx, TxToCallMsg(wt.WalletAddr, types.NewTx(&dynamic)))
		if err != nil {
			return nil, err
		}
		dynamic.Gas = gas

		if opt.Debug {
			log.Log.Debug(fmt.Sprintf("estimate: gas:%d feeCap:%s tipCap:%s",
				gas, dynamic.GasFeeCap.String(), dynamic.GasTipCap.String()))
		}
	}

	if opt.Gas.RuleSet() {
		dynamic.Gas = opt.Gas.GasLimit.Uint64()
		dynamic.GasFeeCap = big.NewInt(0).Add(&opt.Gas.GasPrice, gasTipCap)
		dynamic.GasTipCap = gasTipCap
	}

	r := &bozdo.DefaultRes{
		ECost: Estimate(types.NewTx(&dynamic), bozdo.BigIntSum(l1Fee, data.ExtraFee), "", nil),
	}

	r.ECost.Details = data.Details
	if l1Fee != nil && l1Fee.Cmp(big.NewInt(0)) != 0 {
		r.ECost.Details = append(r.ECost.Details, bozdo.NewProtocolFeeDetails(l1Fee, c.Network(), c.Cfg.MainToken))
	}

	if data.Rate != nil && opt.ExchangeRate != nil {
		r.ECost.Details = append(r.ECost.Details, bozdo.NewSwapRateRatio(*opt.ExchangeRate, *data.Rate))
	}

	if !opt.NoSend {

		if opt.Debug {
			log.Log.Debug(fmt.Sprintf("execute: gas:%d feeCap:%s tipCap:%s",
				dynamic.Gas, dynamic.GasFeeCap.String(), dynamic.GasTipCap.String()))
		}

		tx := types.NewTx(&dynamic)

		signer := types.NewLondonSigner(c.Cfg.NetworkId)

		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), wt.PrivateKey)
		if err != nil {
			return nil, err
		}

		tx, err = tx.WithSignature(signer, signature)
		if err != nil {
			return nil, err
		}

		if err := c.Cli.SendTransaction(ctx, tx); err != nil {
			return nil, err
		}

		r.Tx = c.NewTx(tx.Hash(), data.Code, r.ECost.Details)
	}

	return r, nil
}
func LondonReadyBoosty(ctx context.Context, c *EtheriumClient, opt *TxOpt, data *bozdo.TxData) (*bozdo.DefaultRes, error) {
	wt, err := NewWalletTransactor(opt.Pk)
	if err != nil {
		return nil, err
	}

	nonce, err := c.Cli.NonceAt(ctx, wt.WalletAddr, nil)
	if err != nil {
		return nil, err
	}

	dynamic := types.DynamicFeeTx{
		ChainID: c.Cfg.NetworkId,
		Nonce:   nonce,
		To:      &data.ContractAddr,
		Value:   data.Value,
		Data:    data.Data,
	}

	gasTipCap, err := c.Cli.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, err
	}

	gasTipCap = big.NewInt(0).Mul(gasTipCap, big.NewInt(1))

	l1Fee, err := c.Cfg.EstimateL1Gas(ctx, data.Data)
	if err != nil {
		return nil, err
	}

	if opt.Debug {
		log.Log.Debug("l1Fee "+opt.TaskType.String(), zap.String("l1Fee", WEIToEther(l1Fee).String()+" ETH"))
	}

	dynamic.Value = data.Value

	if opt.NoSend || !opt.Gas.RuleSet() {

		header, err := c.Cli.HeaderByNumber(ctx, nil)
		if err != nil {
			return nil, err
		}

		dynamic.GasTipCap = gasTipCap

		dynamic.GasFeeCap = bozdo.BigIntSum(header.BaseFee, header.BaseFee, dynamic.GasTipCap)

		gas, err := c.Cli.EstimateGas(ctx, TxToCallMsg(wt.WalletAddr, types.NewTx(&dynamic)))
		if err != nil {
			return nil, err
		}
		dynamic.Gas = gas

		if opt.Debug {
			log.Log.Debug(fmt.Sprintf("estimate: gas:%d feeCap:%s tipCap:%s",
				gas, dynamic.GasFeeCap.String(), dynamic.GasTipCap.String()))
		}
	}

	if opt.Gas.RuleSet() {
		dynamic.Gas = opt.Gas.GasLimit.Uint64()
		dynamic.GasFeeCap = &opt.Gas.GasPrice
		dynamic.GasTipCap = gasTipCap
	}

	r := &bozdo.DefaultRes{
		ECost: Estimate(types.NewTx(&dynamic), bozdo.BigIntSum(l1Fee, data.ExtraFee), "", nil),
	}

	r.ECost.Details = data.Details
	if l1Fee != nil && l1Fee.Cmp(big.NewInt(0)) != 0 {
		r.ECost.Details = append(r.ECost.Details, bozdo.NewProtocolFeeDetails(l1Fee, c.Network(), c.Cfg.MainToken))
	}

	if data.Rate != nil && opt.ExchangeRate != nil {
		r.ECost.Details = append(r.ECost.Details, bozdo.NewSwapRateRatio(*opt.ExchangeRate, *data.Rate))
	}

	if !opt.NoSend {

		if opt.Debug {
			log.Log.Debug(fmt.Sprintf("execute: gas:%d feeCap:%s tipCap:%s",
				dynamic.Gas, dynamic.GasFeeCap.String(), dynamic.GasTipCap.String()))
		}

		tx := types.NewTx(&dynamic)

		signer := types.NewLondonSigner(c.Cfg.NetworkId)

		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), wt.PrivateKey)
		if err != nil {
			return nil, err
		}

		tx, err = tx.WithSignature(signer, signature)
		if err != nil {
			return nil, err
		}

		if err := c.Cli.SendTransaction(ctx, tx); err != nil {
			return nil, err
		}

		r.Tx = c.NewTx(tx.Hash(), data.Code, r.ECost.Details)
	}

	return r, nil
}
func LondonNotReady(ctx context.Context, c *EtheriumClient, opt *TxOpt, data *bozdo.TxData) (*bozdo.DefaultRes, error) {
	wt, err := NewWalletTransactor(opt.Pk)
	if err != nil {
		return nil, err
	}

	nonce, err := c.Cli.NonceAt(ctx, wt.WalletAddr, nil)
	if err != nil {
		return nil, err
	}

	dynamic := types.LegacyTx{
		Nonce:    nonce,
		GasPrice: nil,
		Gas:      0,
		To:       &data.ContractAddr,
		Value:    data.Value,
		Data:     data.Data,
		V:        nil,
		R:        nil,
		S:        nil,
	}

	l1Fee, err := c.Cfg.EstimateL1Gas(ctx, data.Data)
	if err != nil {
		return nil, err
	}

	if opt.Debug {
		log.Log.Debug("l1Fee "+opt.TaskType.String(), zap.String("l1Fee", WEIToEther(l1Fee).String()+" ETH"))
	}

	dynamic.Value = data.Value

	if opt.NoSend || !opt.Gas.RuleSet() {

		gasPrice, err := c.Cli.SuggestGasPrice(ctx)
		if err != nil {
			return nil, err
		}

		dynamic.GasPrice = gasPrice

		gas, err := c.Cli.EstimateGas(ctx, TxToCallMsg(wt.WalletAddr, types.NewTx(&dynamic)))
		if err != nil {
			return nil, err
		}

		dynamic.Gas = gas

		if opt.Debug {
			log.Log.Debug(fmt.Sprintf("estimate: gas:%d feeCap:%s tipCap:%s",
				gas, dynamic.GasPrice.String(), dynamic.Gas))
		}
	}

	if opt.Gas.RuleSet() {
		dynamic.Gas = opt.Gas.GasLimit.Uint64()
		dynamic.GasPrice = &opt.Gas.GasPrice
	}

	r := &bozdo.DefaultRes{
		ECost: Estimate(types.NewTx(&dynamic), bozdo.BigIntSum(l1Fee, data.ExtraFee), "", nil),
	}

	r.ECost.Details = data.Details
	if l1Fee != nil && l1Fee.Cmp(big.NewInt(0)) != 0 {
		r.ECost.Details = append(r.ECost.Details, bozdo.NewProtocolFeeDetails(l1Fee, c.Network(), c.Cfg.MainToken))
	}

	if data.Rate != nil && opt.ExchangeRate != nil {
		r.ECost.Details = append(r.ECost.Details, bozdo.NewSwapRateRatio(*opt.ExchangeRate, *data.Rate))
	}

	if !opt.NoSend {

		if opt.Debug {
			log.Log.Debug(fmt.Sprintf("execute: gas:%d feeCap:%s tipCap:%s",
				dynamic.Gas, dynamic.GasPrice.String(), dynamic.GasPrice.String()))
		}

		tx := types.NewTx(&dynamic)

		signer := types.NewCancunSigner(c.Cfg.NetworkId)

		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), wt.PrivateKey)
		if err != nil {
			return nil, err
		}

		tx, err = tx.WithSignature(signer, signature)
		if err != nil {
			return nil, err
		}

		if err := c.Cli.SendTransaction(ctx, tx); err != nil {
			return nil, err
		}

		r.Tx = c.NewTx(tx.Hash(), data.Code, r.ECost.Details)
	}

	return r, nil
}

func TxToCallMsg(from common.Address, tx *types.Transaction) ethereum.CallMsg {
	return ethereum.CallMsg{
		From:      from,
		To:        tx.To(),
		Gas:       tx.Gas(),
		GasPrice:  tx.GasPrice(),
		GasFeeCap: tx.GasFeeCap(),
		GasTipCap: tx.GasTipCap(),
		Value:     tx.Value(),
		Data:      tx.Data(),
	}
}

func Gwei(wei *big.Int) *big.Float {
	return new(big.Float).Quo(big.NewFloat(0).SetInt(wei), big.NewFloat(1e9))
}

func (c *EtheriumClient) EstimateL1GasFee(ctx context.Context, data []byte) (*big.Int, error) {
	optFeeCaller, err := optimism_fee.NewStorageCaller(common.HexToAddress("0x420000000000000000000000000000000000000F"), c.Cli)
	if err != nil {
		return nil, err
	}

	o := &bind.CallOpts{}
	o.Context = ctx
	l1Gasfee, err := optFeeCaller.GetL1Fee(o, data)
	if err != nil {
		return nil, err
	}
	return bozdo.BigIntSum(l1Gasfee, bozdo.Percent(l1Gasfee, 70)), nil
}

func TxGas(tx *types.Transaction) string {

	s := fmt.Sprintf("gas:%d gasPrice:%s GWEI maxFee:%s GWEI tipCap:%s GWEI",
		tx.Gas(), Gwei(tx.GasPrice()).String(), Gwei(tx.GasFeeCap()).String(), Gwei(tx.GasTipCap()).String())
	return s
}

func NeLondonReadyTx(ctx context.Context, c *EtheriumClient, opt *TxOpt, data *bozdo.TxData) (*bozdo.DefaultRes, error) {
	wt, err := NewWalletTransactor(opt.Pk)
	if err != nil {
		return nil, err
	}

	nonce, err := c.Cli.NonceAt(ctx, wt.WalletAddr, nil)
	if err != nil {
		return nil, err
	}

	legacyTx := types.LegacyTx{
		Nonce: nonce,
		To:    &data.ContractAddr,
		Value: data.Value,
		Data:  data.Data,
	}

	l1Fee, err := c.Cfg.EstimateL1Gas(ctx, data.Data)
	if err != nil {
		return nil, err
	}

	if opt.Debug {
		log.Log.Debug("l1Fee "+opt.TaskType.String(), zap.String("l1Fee", WEIToEther(l1Fee).String()+" ETH"))
	}

	legacyTx.Value = data.Value

	if opt.NoSend || !opt.Gas.RuleSet() {

		gp, err := c.Cli.SuggestGasPrice(ctx)
		if err != nil {
			return nil, err
		}

		legacyTx.GasPrice = gp

		gas, err := c.Cli.EstimateGas(ctx, TxToCallMsg(wt.WalletAddr, types.NewTx(&legacyTx)))
		if err != nil {
			return nil, err
		}
		legacyTx.Gas = gas

		if opt.Debug {
			log.Log.Debug(fmt.Sprintf("estimate: gas:%d feeCap:%s tipCap:%s",
				gas, legacyTx.Gas, legacyTx.GasPrice.String()))
		}
	}

	if opt.Gas.RuleSet() {
		legacyTx.Gas = opt.Gas.GasLimit.Uint64()
		legacyTx.GasPrice = &opt.Gas.GasPrice
	}

	r := &bozdo.DefaultRes{
		ECost: Estimate(types.NewTx(&legacyTx), bozdo.BigIntSum(l1Fee, data.ExtraFee), "", nil),
	}

	r.ECost.Details = data.Details
	if l1Fee != nil && l1Fee.Cmp(big.NewInt(0)) != 0 {
		r.ECost.Details = append(r.ECost.Details, bozdo.NewProtocolFeeDetails(l1Fee, c.Network(), c.Cfg.MainToken))
	}

	if data.Rate != nil && opt.ExchangeRate != nil {
		r.ECost.Details = append(r.ECost.Details, bozdo.NewSwapRateRatio(*opt.ExchangeRate, *data.Rate))
	}

	if !opt.NoSend {

		if opt.Debug {
			log.Log.Debug(fmt.Sprintf("execute: gas:%d gasPrice::%s",
				legacyTx.Gas, legacyTx.GasPrice.String()))
		}

		signer := types.NewCancunSigner(c.Cfg.NetworkId)

		tx, err := types.SignNewTx(wt.PrivateKey, signer, &legacyTx)
		if err != nil {
			return nil, err
		}

		addr, err := signer.Sender(tx)
		if err != nil {
			return nil, err
		}

		println(addr.String())
		if err := c.Cli.SendTransaction(ctx, tx); err != nil {
			return nil, err
		}

		r.Tx = c.NewTx(tx.Hash(), data.Code, r.ECost.Details)
	}

	return r, nil
}

func Estimate(tx *types.Transaction, extraFee *big.Int, name string, details []bozdo.TxDetail) *bozdo.EstimatedGasCost {
	gasLimit := new(big.Int).SetUint64(tx.Gas())

	if tx.Type() == types.DynamicFeeTxType {
		fee := bozdo.BigIntSum(tx.GasFeeCap())

		return &bozdo.EstimatedGasCost{
			Name:        "bridge",
			GasLimit:    gasLimit,
			GasPrice:    fee,
			TotalGasWei: new(big.Int).Mul(bozdo.BigIntSum(fee), gasLimit),
			Details:     details,
			ExtraFee:    extraFee,
		}
	}

	return &bozdo.EstimatedGasCost{
		Type:        bozdo.TxTypeLegacy,
		Name:        name,
		GasLimit:    gasLimit,
		GasPrice:    tx.GasPrice(),
		TotalGasWei: new(big.Int).Mul(gasLimit, tx.GasPrice()),
		Details:     details,
		ExtraFee:    extraFee,
	}
}
