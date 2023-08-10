package zksyncera

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/ezkaliburrouter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type ezkaliburMaker struct {
	*Client
}

func (c *Client) EzkaliburSwap(ctx context.Context, req *defi.DefaultSwapReq) (*defi.DefaultRes, error) {
	return c.GenericSwap(ctx, &ezkaliburMaker{c}, req)
}

func (c *ezkaliburMaker) MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {

	w, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}
	call, err := ezkaliburrouter.NewStorageCaller(c.Cfg.Ezkalibur.Router, c.Provider)
	if err != nil {
		return nil, err
	}
	from, supported := c.Cfg.TokenMap[req.FromToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(req.FromToken)
	}
	to, supported := c.Cfg.TokenMap[req.ToToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(req.FromToken)
	}

	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	path := []common.Address{from, to}
	amsOut, err := call.GetAmountsOut(&bind.CallOpts{Context: ctx}, req.Amount, path)
	if err != nil {
		return nil, err
	}
	if len(amsOut) != 2 {
		return nil, errors.New("errors getting rate")
	}

	abiEzkaliburrouter, err := ezkaliburrouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	if req.FromToken == v1.Token_ETH {
		amOut := amsOut[1]
		amSlip, err := defi.Slippage(amOut, req.Slippage)
		if err != nil {
			return nil, err
		}
		amOut = amSlip
		data, err := abiEzkaliburrouter.Pack("swapExactETHForTokensSupportingFeeOnTransferTokens", amOut, path, w.WalletAddr, w.WalletAddr, DefaultDeadLine())
		if err != nil {
			return nil, err
		}
		return &bozdo.TxData{
			Data:         data,
			Value:        value,
			ContractAddr: c.Cfg.Ezkalibur.Router,
		}, nil

	} else if req.ToToken == v1.Token_ETH {
		amOut := amsOut[1]
		amSlip, err := defi.Slippage(amOut, req.Slippage)
		if err != nil {
			return nil, err
		}
		amOut = amSlip
		data, err := abiEzkaliburrouter.Pack("swapExactTokensForETHSupportingFeeOnTransferTokens", req.Amount, amOut, path, w.WalletAddr, w.WalletAddr, DefaultDeadLine())
		if err != nil {
			return nil, err
		}
		return &bozdo.TxData{
			Data:         data,
			Value:        value,
			ContractAddr: c.Cfg.Ezkalibur.Router,
		}, nil
	}

	return nil, errors.New("unsupported input params")
}
