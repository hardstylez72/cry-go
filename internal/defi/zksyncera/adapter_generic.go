package zksyncera

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/utils"
)

type TxSwapMaker interface {
	MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (*defi.TxData, error)
}

type TxBridgeMaker interface {
	MakeBridgeTx(ctx context.Context, req *defi.DefaultBridgeReq) (*defi.TxData, error)
}

func (c *Client) GenericSwap(ctx context.Context, maker TxSwapMaker, req *defi.DefaultSwapReq) (*defi.DefaultRes, error) {
	result := &defi.DefaultRes{}

	transactor, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	tokenLimitChecker, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
		Token:       req.FromToken,
		WalletPK:    req.WalletPK,
		Amount:      req.Amount,
		SpenderAddr: c.Cfg.SpaceFI.Router,
	})
	if err != nil {
		return nil, errors.Wrap(err, "TokenLimitChecker")
	}
	result.ApproveTx = tokenLimitChecker.ApproveTx

	txData, err := maker.MakeSwapTx(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "makeSpaceFiSwapData")
	}

	tx := utils.CreateFunctionCallTransaction(
		transactor.WalletAddr,
		txData.ContractAddr,
		big.NewInt(0),
		big.NewInt(0),
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

	hash, err := c.Provider.SendRawTransaction(raw)
	if err != nil {
		return nil, errors.Wrap(err, "Provider.SendRawTransaction")
	}

	result.Tx = c.NewTx(hash, defi.CodeContract)

	return result, nil
}
func (c *Client) GenericBridge(ctx context.Context, maker TxBridgeMaker, req *defi.DefaultBridgeReq) (*defi.DefaultRes, error) {
	result := &defi.DefaultRes{}

	transactor, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	tokenLimitChecker, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
		Token:       req.FromToken,
		WalletPK:    req.WalletPK,
		Amount:      req.Amount,
		SpenderAddr: c.Cfg.SpaceFI.Router,
	})
	if err != nil {
		return nil, errors.Wrap(err, "TokenLimitChecker")
	}
	result.ApproveTx = tokenLimitChecker.ApproveTx

	txData, err := maker.MakeBridgeTx(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "makeSpaceFiSwapData")
	}

	tx := utils.CreateFunctionCallTransaction(
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

	hash, err := c.Provider.SendRawTransaction(raw)
	if err != nil {
		return nil, errors.Wrap(err, "Provider.SendRawTransaction")
	}

	result.Tx = c.NewTx(hash, defi.CodeContract)

	return result, nil
}
