package zksyncera

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/pkg/errors"
)

type TxSwapMaker interface {
	MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error)
}

type TxBridgeMaker interface {
	MakeBridgeTx(ctx context.Context, req *defi.DefaultBridgeReq) (*bozdo.TxData, error)
}

func (c *Client) GenericSwap(ctx context.Context, maker TxSwapMaker, req *defi.DefaultSwapReq) (*bozdo.DefaultRes, error) {
	result := &bozdo.DefaultRes{}

	transactor, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	txData, err := maker.MakeSwapTx(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "makeSpaceFiSwapData")
	}

	tokenLimitChecker, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
		Token:       req.FromToken,
		WalletPK:    req.WalletPK,
		Amount:      req.Amount,
		SpenderAddr: txData.ContractAddr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "TokenLimitChecker")
	}
	result.ApproveTx = tokenLimitChecker.ApproveTx

	tx := CreateFunctionCallTransaction(
		transactor.WalletAddr,
		txData.ContractAddr,
		big.NewInt(0),
		txData.Gas,
		txData.Value,
		txData.Data,
		nil, nil,
	)

	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, transactor.Signer)
	if err != nil {
		return nil, errors.Wrap(err, "Make712Tx")
	}

	result.ECost = estimate
	result.ECost.Details = append(result.ECost.Details, txData.Details...)
	if txData.Rate != nil && req.ExchangeRate != nil {
		result.ECost.Details = append(result.ECost.Details, bozdo.NewSwapRateRatio(*req.ExchangeRate, *txData.Rate))
	}

	if req.EstimateOnly {
		return result, nil
	}
	result.TxDetails = result.ECost.Details

	hash, err := c.ClientL2.SendRawTransaction(ctx, raw)
	if err != nil {
		return nil, errors.Wrap(err, "rpcL2.SendRawTransaction")
	}

	result.Tx = c.NewTx(hash, defi.CodeContract, txData.Details)

	return result, nil
}
func (c *Client) GenericBridge(ctx context.Context, maker TxBridgeMaker, req *defi.DefaultBridgeReq) (*bozdo.DefaultRes, error) {
	result := &bozdo.DefaultRes{}

	transactor, err := NewWalletTransactor(req.PK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	txData, err := maker.MakeBridgeTx(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "makeSpaceFiSwapData")
	}

	tokenLimitChecker, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
		Token:       req.FromToken,
		WalletPK:    req.PK,
		Amount:      req.Amount,
		SpenderAddr: txData.ContractAddr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "TokenLimitChecker")
	}
	result.ApproveTx = tokenLimitChecker.ApproveTx

	tx := CreateFunctionCallTransaction(
		transactor.WalletAddr,
		txData.ContractAddr,
		nil,
		nil,
		txData.Value,
		txData.Data,
		nil, nil,
	)

	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, transactor.Signer)
	if err != nil {
		return nil, errors.Wrap(err, "Make712Tx")
	}

	result.ECost = estimate

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.ClientL2.SendRawTransaction(ctx, raw)
	if err != nil {
		return nil, errors.Wrap(err, "rpcL2.SendRawTransaction")
	}

	result.Tx = c.NewTx(hash, defi.CodeContract, txData.Details)

	return result, nil
}
