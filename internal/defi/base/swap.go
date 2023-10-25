package base

import (
	"context"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/base/pancake"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/odos"
	"github.com/hardstylez72/cry/internal/defi/woofi"
	"github.com/hardstylez72/cry/internal/log"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"go.uber.org/zap"
)

func (c *Client) Swap(ctx context.Context, req *defi.DefaultSwapReq, taskType v1.TaskType) (*bozdo.DefaultRes, error) {

	wt, err := defi.NewWalletTransactor(req.WalletPK)
	if err != nil {
		return nil, err
	}

	token, ok := c.defi.Cfg.TokenMap[req.FromToken]
	if !ok {
		return nil, defi.ErrTokenNotSupportedFn(req.FromToken)
	}
	var ca common.Address
	switch taskType {
	case v1.TaskType_PancakeSwap:
		ca = c.defi.Cfg.Dict.Pancake.Router
	case v1.TaskType_OdosSwap:
		ca = c.defi.Cfg.Dict.Odos.Router
	case v1.TaskType_WoofiSwap:
		ca = c.defi.Cfg.Dict.Woofi.Router
	}

	approve, err := c.defi.Approve(ctx, &defi.Approve{
		TokenAddr:   token,
		PK:          req.WalletPK,
		Amount:      req.Amount,
		SpenderAddr: ca,
	})
	if err != nil {
		return nil, err
	}

	var data *bozdo.TxData
	switch taskType {
	case v1.TaskType_PancakeSwap:
		cli := &pancake.Swapper{EtheriumClient: c.defi}
		data, err = cli.MakeSwapTx(ctx, req)
	case v1.TaskType_OdosSwap:
		cli := &odos.OdosMaker{
			CA:       c.defi.Cfg.Dict.Odos.Router,
			TokenMap: c.defi.Cfg.TokenMap,
			CliHttp:  &http.Client{},
			ChainId:  c.NetworkId,
			Network:  c.Network(),
			Addr:     wt.WalletAddr,
		}
		data, err = cli.MakeSwapTx(ctx, req)
	case v1.TaskType_WoofiSwap:
		cli := woofi.NewSwapper(c.defi.Cfg.Dict.Woofi.Router, c.defi.Cli, c.defi.Cfg.TokenMap)
		data, err = cli.MakeSwapTx(ctx, req)
	}
	if err != nil {
		return nil, err
	}

	opt := &TxOpt{
		NoSend:       req.EstimateOnly,
		Pk:           req.WalletPK,
		Gas:          req.Gas,
		ExchangeRate: req.ExchangeRate,
		Debug:        req.Debug,
		TaskType:     taskType,
	}

	out, err := c.LondonReadyTx(ctx, opt, data)
	if err != nil {
		return nil, err
	}
	if approve.ApproveTx != nil {
		out.ApproveTx = c.defi.NewTx(approve.ApproveTx.Hash(), bozdo.CodeApprove, nil)
	}

	return out, nil
}

type TxOpt struct {
	NoSend       bool
	Pk           string
	Gas          *bozdo.Gas
	ExchangeRate *float64
	Debug        bool
	TaskType     v1.TaskType
}

func (c *Client) LondonReadyTx(ctx context.Context, opt *TxOpt, data *bozdo.TxData) (*bozdo.DefaultRes, error) {
	wt, err := defi.NewWalletTransactor(opt.Pk)
	if err != nil {
		return nil, err
	}

	nonce, err := c.defi.Cli.NonceAt(ctx, wt.WalletAddr, nil)
	if err != nil {
		return nil, err
	}

	dynamic := types.DynamicFeeTx{
		ChainID: c.NetworkId,
		Nonce:   nonce,
		To:      &data.ContractAddr,
		Value:   data.Value,
		Data:    data.Data,
	}

	gasTipCap := big.NewInt(1e6)

	if opt.NoSend || !opt.Gas.RuleSet() {

		header, err := c.defi.Cli.HeaderByNumber(ctx, nil)
		if err != nil {
			return nil, err
		}

		dynamic.GasTipCap = gasTipCap

		dynamic.GasFeeCap = bozdo.BigIntSum(header.BaseFee, bozdo.Percent(header.BaseFee, 50), dynamic.GasTipCap)

		gas, err := c.defi.Cli.EstimateGas(ctx, txToCallMsg(wt.WalletAddr, types.NewTx(&dynamic)))
		if err != nil {
			return nil, err
		}
		dynamic.Gas = gas
	}

	if opt.Gas.RuleSet() {
		dynamic.Gas = opt.Gas.GasLimit.Uint64()
		dynamic.GasFeeCap = &opt.Gas.GasPrice
		dynamic.GasTipCap = gasTipCap
	}

	l1Fee, err := c.EstimateL1GasFee(ctx, data.Data)
	if err != nil {
		return nil, err
	}

	if opt.Debug {
		log.Log.Debug("l1Fee "+opt.TaskType.String(), zap.String("l1Fee", defi.WEIToEther(l1Fee).String()+" ETH"))
	}

	dynamic.Value = bozdo.BigIntSum(dynamic.Value)

	r := &bozdo.DefaultRes{
		ECost: defi.Estimate(types.NewTx(&dynamic), nil, "", data.Details),
	}

	r.ECost.Details = append(r.ECost.Details, bozdo.NewProtocolFeeDetails(l1Fee, v1.Network_Base, c.defi.Cfg.MainToken))

	if data.Rate != nil && opt.ExchangeRate != nil {
		r.ECost.Details = append(r.ECost.Details, bozdo.NewSwapRateRatio(*opt.ExchangeRate, *data.Rate))
	}

	if !opt.NoSend {

		tx := types.NewTx(&dynamic)

		if opt.Debug {
			log.Log.Debug("swap tx gas "+opt.TaskType.String(), zap.String("gas", TxGas(tx)))
		}

		signer := types.NewLondonSigner(c.NetworkId)

		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), wt.PrivateKey)
		if err != nil {
			return nil, err
		}

		tx, err = tx.WithSignature(signer, signature)
		if err != nil {
			return nil, err
		}

		if err := c.defi.Cli.SendTransaction(ctx, tx); err != nil {
			return nil, err
		}

		r.Tx = c.defi.NewTx(tx.Hash(), bozdo.CodeSwap, r.ECost.Details)
	}

	return r, nil
}

func txToCallMsg(from common.Address, tx *types.Transaction) ethereum.CallMsg {
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
