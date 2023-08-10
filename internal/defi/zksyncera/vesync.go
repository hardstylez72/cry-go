package zksyncera

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/vesyncrouter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type veSyncMaker struct {
	*Client
}

func (c *Client) VeSyncSwap(ctx context.Context, req *defi.DefaultSwapReq) (*defi.DefaultRes, error) {
	return c.GenericSwap(ctx, &veSyncMaker{c}, req)
}

func (c *veSyncMaker) MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {

	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	w, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}
	call, err := vesyncrouter.NewStorageCaller(c.Cfg.VeSync.Router, c.Provider)
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

	path := []vesyncrouter.Routerroute{
		{
			From:   from,
			To:     to,
			Stable: false,
		},
	}
	amsOut, err := call.GetAmountsOut(&bind.CallOpts{Context: ctx}, req.Amount, path)
	if err != nil {
		return nil, err
	}
	if len(amsOut) != 2 {
		return nil, errors.New("errors getting rate")
	}

	abiVeSyncrouter, err := vesyncrouter.StorageMetaData.GetAbi()
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
		data, err := abiVeSyncrouter.Pack("swapExactETHForTokens", amOut, path, w.WalletAddr, DefaultDeadLine())
		if err != nil {
			return nil, err
		}

		return &bozdo.TxData{
			Data:         data,
			Value:        value,
			ContractAddr: c.Cfg.VeSync.Router,
		}, nil

	} else if req.ToToken == v1.Token_ETH {
		amOut := amsOut[1]
		amSlip, err := defi.Slippage(amOut, req.Slippage)
		if err != nil {
			return nil, err
		}
		amOut = amSlip
		data, err := abiVeSyncrouter.Pack("swapExactTokensForETH", req.Amount, amOut, path, w.WalletAddr, DefaultDeadLine())
		if err != nil {
			return nil, err
		}

		return &bozdo.TxData{
			Data:         data,
			Value:        value,
			ContractAddr: c.Cfg.VeSync.Router,
		}, nil
	}

	return nil, errors.New("unsupported input params")
}
