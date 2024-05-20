package defi

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hardstylez72/cry/internal/defi/contracts/erc_20"
	"github.com/pkg/errors"
)

type AllowedReq struct {
	Token       Token
	WalletAddr  common.Address
	SpenderAddr common.Address
}

type AllowedRes struct {
	Allowance *big.Int
}

func (c *EtheriumClient) TokenAllowed(ctx context.Context, req *AllowedReq) (*AllowedRes, error) {
	return c.tokenAllowed(ctx, req)
}
func (c *EtheriumClient) tokenAllowed(ctx context.Context, req *AllowedReq) (*AllowedRes, error) {

	addr, ok := c.Cfg.TokenMap[req.Token]
	if !ok {
		return nil, ErrTokenNotSupportedFn(req.Token)
	}

	caller, err := erc_20.NewStorageCaller(addr, c.Cli)
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

type AllowedReqV2 struct {
	TokenAddr   common.Address
	WalletAddr  common.Address
	SpenderAddr common.Address
}

func (c *EtheriumClient) tokenAllowedV2(ctx context.Context, req *AllowedReqV2) (*AllowedRes, error) {

	addr := req.TokenAddr

	caller, err := erc_20.NewStorageCaller(addr, c.Cli)
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
	Token       Token
	Wallet      *WalletTransactor
	Amount      *big.Int
	SpenderAddr common.Address
}

type ApproveRes struct {
	Tx     *types.Transaction
	TxHash common.Hash
}

type TokenLimitCheckerReq struct {
	Token       Token
	Wallet      *WalletTransactor
	Amount      *big.Int
	SpenderAddr common.Address
}

type TokenLimitCheckerRes struct {
	LimitExtended bool
	ApproveTx     *types.Transaction
}

func (c *EtheriumClient) Approve(ctx context.Context, req *Approve) (*TokenLimitCheckerRes, error) {
	r := &TokenLimitCheckerRes{
		LimitExtended: false,
		ApproveTx:     nil,
	}

	if req.TokenAddr.String() == c.Cfg.TokenMap[c.Cfg.MainToken].String() {
		return r, nil
	}

	w, err := NewWalletTransactor(req.PK)
	if err != nil {
		return nil, errors.Wrap(err, "NewWalletTransactor")
	}
	allowed, err := c.tokenAllowedV2(ctx, &AllowedReqV2{
		TokenAddr:   req.TokenAddr,
		WalletAddr:  w.WalletAddr,
		SpenderAddr: req.SpenderAddr,
	})
	if err != nil {
		return nil, err
	}

	if req.Amount.Cmp(allowed.Allowance) == 1 {

		b, err := c.Balance(ctx, &BalanceReq{
			WalletAddress: w.WalletAddrHR,
			TokenAddr:     req.TokenAddr,
		})
		if err != nil {
			return nil, err
		}

		tx, err := c.tokenApproveV2(ctx, &ApproveReqV2{
			TokenAddr:   req.TokenAddr,
			PK:          req.PK,
			Amount:      b,
			SpenderAddr: req.SpenderAddr,
		})
		if err != nil {
			return nil, err
		}
		r.LimitExtended = true
		r.ApproveTx = tx.Tx

		time.Sleep(time.Second * 5)
		_ = c.WaitTxComplete(ctx, tx.Tx.Hash().String())
	}

	return r, nil
}

type Approve struct {
	TokenAddr   common.Address
	PK          string
	Amount      *big.Int
	SpenderAddr common.Address
}

func (c *EtheriumClient) TokenLimitChecker(ctx context.Context, req *TokenLimitCheckerReq) (*TokenLimitCheckerRes, error) {
	r := &TokenLimitCheckerRes{
		LimitExtended: false,
		ApproveTx:     nil,
	}

	if req.Token == c.GetNetworkToken() {
		return r, nil
	}

	allowed, err := c.TokenAllowed(ctx, &AllowedReq{
		Token:       req.Token,
		WalletAddr:  req.Wallet.WalletAddr,
		SpenderAddr: req.SpenderAddr,
	})
	if err != nil {
		return nil, err
	}

	if req.Amount.Cmp(allowed.Allowance) == 1 {

		b, err := c.GetBalance(ctx, &GetBalanceReq{
			WalletAddress: req.Wallet.WalletAddrHR,
			Token:         req.Token,
		})
		if err != nil {
			return nil, err
		}

		tx, err := c.TokenApprove(ctx, &ApproveReq{
			Token:       req.Token,
			Wallet:      req.Wallet,
			Amount:      b.WEI,
			SpenderAddr: req.SpenderAddr,
		})
		if err != nil {
			return nil, err
		}
		r.LimitExtended = true
		r.ApproveTx = tx.Tx

		time.Sleep(time.Second * 5)
		_ = c.WaitTxComplete(ctx, tx.Tx.Hash().String())
	}

	return r, nil
}

func (c *EtheriumClient) TokenApprove(ctx context.Context, req *ApproveReq) (*ApproveRes, error) {
	return c.tokenApprove(ctx, req)
}

func (r *ApproveReq) Validate(c *Config) error {

	if r.Wallet == nil {
		return errors.New("empty wallet")
	}

	_, ok := c.TokenMap[r.Token]
	if !ok {
		return ErrTokenNotSupportedFn(r.Token)
	}

	if r.Amount.Cmp(big.NewInt(0)) == 0 {
		return errors.New("zero Amount")
	}

	return nil
}

func (c *EtheriumClient) tokenApprove(ctx context.Context, req *ApproveReq) (*ApproveRes, error) {

	if err := req.Validate(c.Cfg); err != nil {
		return nil, err
	}

	addr := c.Cfg.TokenMap[req.Token]

	caller, err := erc_20.NewStorageTransactor(addr, c.Cli)
	if err != nil {
		return nil, errors.Wrap(err, "stg.NewStgCaller")
	}

	chainID, err := c.Cli.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	opt, err := bind.NewKeyedTransactorWithChainID(req.Wallet.PrivateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	opt.Context = ctx

	tx, err := caller.Approve(opt, req.SpenderAddr, req.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "caller.Allowance")
	}

	return &ApproveRes{Tx: tx}, nil
}

type ApproveReqV2 struct {
	TokenAddr   common.Address
	PK          string
	Amount      *big.Int
	SpenderAddr common.Address
}

func (c *EtheriumClient) tokenApproveV2(ctx context.Context, req *ApproveReqV2) (*ApproveRes, error) {

	addr := req.TokenAddr

	w, err := NewWalletTransactor(req.PK)
	if err != nil {
		return nil, err
	}

	caller, err := erc_20.NewStorageTransactor(addr, c.Cli)
	if err != nil {
		return nil, errors.Wrap(err, "stg.NewStgCaller")
	}

	chainID, err := c.Cli.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	opt, err := bind.NewKeyedTransactorWithChainID(w.PrivateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	opt.Context = ctx

	tx, err := caller.Approve(opt, req.SpenderAddr, req.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "caller.Allowance")
	}

	return &ApproveRes{Tx: tx}, nil
}
