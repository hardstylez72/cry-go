package zksyncera

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/contracts/erc_20"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type TokenLimitCheckerReq struct {
	Token       v1.Token
	WalletPK    string
	Amount      *big.Int
	SpenderAddr common.Address
	Gas         *bozdo.Gas
}

type TokenLimitCheckerRes struct {
	ApproveTx *bozdo.Transaction
}

func (c *Client) TokenLimitChecker(ctx context.Context, req *TokenLimitCheckerReq) (*TokenLimitCheckerRes, error) {
	r := &TokenLimitCheckerRes{
		ApproveTx: nil,
	}

	if req.Token == c.Cfg.MainToken {
		return &TokenLimitCheckerRes{
			ApproveTx: nil,
		}, nil
	}

	tx, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}
	allowed, err := c.Allowance(ctx, &AllowedReq{
		Token:       req.Token,
		WalletAddr:  tx.WalletAddr,
		SpenderAddr: req.SpenderAddr,
	})
	if err != nil {
		return nil, err
	}

	if req.Amount.Cmp(allowed.Allowance) == 1 {

		b, err := c.GetBalance(ctx, &defi.GetBalanceReq{
			WalletAddress: tx.WalletAddrHR,
			Token:         req.Token,
		})
		if err != nil {
			return nil, err
		}

		tx, err := c.TokenApprove(ctx, &ApproveReq{
			Token:       req.Token,
			Wallet:      tx,
			Amount:      b.WEI,
			SpenderAddr: req.SpenderAddr,
		})
		if err != nil {
			return nil, err
		}

		if err := c.WaitTxComplete(ctx, tx.TxHash.String()); err != nil {
			return nil, err
		}

		r.ApproveTx = c.NewTx(tx.TxHash, defi.CodeApprove, nil)
	}

	return r, nil
}

type AllowedReq struct {
	Token       v1.Token
	WalletAddr  common.Address
	SpenderAddr common.Address
	Retry       int
}

type AllowedRes struct {
	Allowance *big.Int
}

func (c *Client) Allowance(ctx context.Context, req *AllowedReq) (*AllowedRes, error) {
	addr, ok := c.Cfg.TokenMap[req.Token]
	if !ok {
		return nil, defi.ErrTokenNotSupportedFn(req.Token)
	}

	caller, err := erc_20.NewStorageCaller(addr, c.ClientL2)
	if err != nil {
		return nil, errors.Wrap(err, "stg.NewStgCaller")
	}

	opt := &bind.CallOpts{Context: ctx}
	allowance, err := caller.Allowance(opt, req.WalletAddr, req.SpenderAddr)
	if err != nil {
		return nil, errors.Wrap(err, "caller.Allowance")
	}

	return &AllowedRes{Allowance: allowance}, nil
}

type ApproveReq struct {
	Token       v1.Token
	Wallet      *WalletTransactor
	Amount      *big.Int
	SpenderAddr common.Address
	Retry       int
}

func (r *ApproveReq) Validate(tm map[v1.Token]common.Address) error {

	if r.Wallet == nil {
		return errors.New("empty wallet")
	}

	_, ok := tm[r.Token]
	if !ok {
		return defi.ErrTokenNotSupportedFn(r.Token)
	}

	if r.Amount.Cmp(big.NewInt(0)) == 0 {
		return errors.New("zero Amount")
	}

	return nil
}

type ApproveRes struct {
	TxHash common.Hash
}

func (c *Client) TokenApprove(ctx context.Context, req *ApproveReq) (*ApproveRes, error) {

	if err := req.Validate(c.Cfg.TokenMap); err != nil {
		return nil, err
	}

	addr := c.Cfg.TokenMap[req.Token]

	abi, err := erc_20.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	data, err := abi.Pack("approve", req.SpenderAddr, req.Amount)
	if err != nil {
		return nil, err
	}

	w, wtx, err := c.Wallet(req.Wallet.PK)
	if err != nil {
		return nil, err
	}

	call := CreateFunctionCallTransaction(w.Address(), addr, nil, big.NewInt(0), nil, data, nil, nil)

	tx, _, err := c.Make712Tx(ctx, call, nil, wtx.Signer)
	if err != nil {
		return nil, err
	}

	hash, err := c.ClientL2.SendRawTransaction(ctx, tx)
	if err != nil {
		return nil, errors.Wrap(err, "caller.Allowance")
	}

	return &ApproveRes{TxHash: hash}, nil
}
