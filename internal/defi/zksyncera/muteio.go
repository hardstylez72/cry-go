package zksyncera

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	muteiorouter "github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/muteio"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func (c *Client) MuteIOSwap(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.DefaultRes, error) {
	if req.FromToken == v1.Token_ETH || req.FromToken == v1.Token_WETH {
		return c.muteIoSwapFromEth(ctx, req)
	}

	res, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
		Token:       req.FromToken,
		WalletPK:    req.WalletPK,
		Amount:      req.Amount,
		SpenderAddr: c.Cfg.Muteio.RouterSwap,
	})
	if err != nil {
		return nil, err
	}

	r, err := c.muteIoSwapToEth(ctx, req)
	if err != nil {
		return nil, err
	}

	if res.ApproveTx != nil {
		r.ApproveTx = res.ApproveTx
	}
	return r, nil
}

// off https://explorer.zksync.io/tx/0x74dc1805388938923ffbfdbf0b021656f7bc1aca19a6508a64ffccf2632c7c94
// my
func (c *Client) muteIoSwapToEth(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.DefaultRes, error) {

	w, wtx, err := c.Wallet(req.WalletPK)
	if err != nil {
		return nil, errors.Wrap(err, "newWalletTransactor")
	}

	muteiorouterabi, err := muteiorouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	amountOut, _, _, err := c.MuteIoPairFee(ctx, req)
	if err != nil {
		return nil, err
	}

	amountOutMin, err := defi.Slippage(amountOut, req.Slippage)
	if err != nil {
		return nil, err
	}

	min := req.Amount
	path := []common.Address{c.Cfg.TokenMap[req.FromToken], c.Cfg.Weth}
	to := w.Address()
	deadline := new(big.Int).SetInt64(time.Now().Add(time.Second * 20).Unix())
	stableMap := []bool{false, false}

	data, err := muteiorouterabi.Pack("swapExactTokensForETHSupportingFeeOnTransferTokens", min, amountOutMin, path, to, deadline, stableMap)
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw function: %w", err)
	}

	tx := CreateFunctionCallTransaction(
		w.Address(),
		c.Cfg.Muteio.RouterSwap,
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		data,
		nil, nil,
	)

	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, wtx.Signer)
	if err != nil {
		return nil, fmt.Errorf("Make712Tx: %w", err)
	}
	result := &bozdo.DefaultRes{}
	result.ECost = estimate

	rate := CalcRate(req.FromToken, req.ToToken, req.Amount, amountOut)
	if rate != nil && req.ExchangeRate != nil {
		result.ECost.Details = append(result.ECost.Details, bozdo.NewSwapRateRatio(*req.ExchangeRate, *rate))
	}

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.ClientL2.SendRawTransaction(ctx, raw)
	if err != nil {
		return nil, errors.Wrap(err, "rpcL2.SendRawTransaction")
	}
	result.Tx = c.NewTx(hash, defi.CodeContract, nil)
	result.TxDetails = result.ECost.Details
	return result, nil
}

// off https://explorer.zksync.io/tx/0xf5aa782e37711f8ebdfe19608ce24d5d0219103264d39f20d04738e677756c7d
// my https://explorer.zksync.io/tx/0x5f86de07362b4de327a59f6ae866ba98bb4ddd65e73d19cfb1baeca5fc702292
func (c *Client) muteIoSwapFromEth(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.DefaultRes, error) {
	wtx, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, errors.Wrap(err, "newWalletTransactor")
	}

	w, wtx, err := c.Wallet(req.WalletPK)
	if err != nil {
		return nil, errors.Wrap(err, "newWalletTransactor")
	}

	muteiorouterabi, err := muteiorouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	amountOut, _, fee, err := c.MuteIoPairFee(ctx, req)
	if err != nil {
		return nil, err
	}

	amountOutMin, err := defi.Slippage(amountOut, req.Slippage)
	if err != nil {
		return nil, err
	}

	path := []common.Address{c.Cfg.Weth, c.Cfg.TokenMap[req.ToToken]}
	to := w.Address()
	deadline := new(big.Int).SetInt64(time.Now().Add(time.Second * 20).Unix())
	stableMap := []bool{false, false}

	data, err := muteiorouterabi.Pack("swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline, stableMap)
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw function: %w", err)
	}

	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	value = bozdo.BigIntSum(value, fee)

	tx := CreateFunctionCallTransaction(
		w.Address(),
		c.Cfg.Muteio.RouterSwap,
		big.NewInt(0),
		big.NewInt(0),
		req.Amount,
		data,
		nil, nil,
	)

	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, wtx.Signer)
	if err != nil {
		return nil, fmt.Errorf("Make712Tx: %w", err)
	}
	result := &bozdo.DefaultRes{}
	result.ECost = estimate

	result.ECost = estimate
	rate := CalcRate(req.FromToken, req.ToToken, req.Amount, amountOut)
	if rate != nil && req.ExchangeRate != nil {
		result.ECost.Details = append(result.ECost.Details, bozdo.NewSwapRateRatio(*req.ExchangeRate, *rate))
	}

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.ClientL2.SendRawTransaction(ctx, raw)
	if err != nil {
		return nil, errors.Wrap(err, "rpcL2.SendRawTransaction")
	}
	result.Tx = c.NewTx(hash, defi.CodeContract, nil)
	result.TxDetails = estimate.Details
	return result, nil
}

func (c *Client) MuteIoPairFee(ctx context.Context, req *defi.DefaultSwapReq) (AmountOut *big.Int, Stable bool, Fee *big.Int, err error) {
	caller, err := muteiorouter.NewStorageCaller(c.Cfg.Muteio.RouterSwap, c.ClientL2)
	if err != nil {
		return nil, false, nil, err
	}
	opt := &bind.CallOpts{}
	opt.Context = ctx

	amOut, err := caller.GetAmountOut(opt, req.Amount, c.Cfg.TokenMap[req.FromToken], c.Cfg.TokenMap[req.ToToken])
	if err != nil {
		return nil, false, nil, err
	}

	return amOut.AmountOut, amOut.Stable, amOut.Fee, nil
}
