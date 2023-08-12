package zksyncera

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/velocorerouter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type velocoreMaker struct {
	*Client
}

// site ETH -> USDC
// me ETH -> USDC https://explorer.zksync.io/tx/0xd93d5099e7f0193778351365fae505e713e89acde63e9a0e229b27bf3c13d8b7

// site USDC -> ETH
// me USDC -> ETH https://explorer.zksync.io/tx/0xaa2dcfa8280071efa6fa4c7264902bbd5d589d2a9ab3b5b6e313cf903a0e9e18
func (c *Client) VelocoreSwap(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.DefaultRes, error) {
	return c.GenericSwap(ctx, &velocoreMaker{c}, req)
}

func (c velocoreMaker) MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {

	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	w, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}
	call, err := velocorerouter.NewStorageCaller(c.Cfg.Velocore.Router, c.Provider)
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

	path := []velocorerouter.Routerroute{{
		From:   from,
		To:     to,
		Stable: false,
	}}
	amsOut, err := call.GetAmountsOut(&bind.CallOpts{Context: ctx}, req.Amount, path)
	if err != nil {
		return nil, err
	}
	if len(amsOut) != 2 {
		return nil, errors.New("errors getting rate")
	}

	abiSource, err := velocorerouter.StorageMetaData.GetAbi()
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
		data, err := abiSource.Pack("swapExactETHForTokens", amOut, path, w.WalletAddr, DefaultDeadLine())
		if err != nil {
			return nil, err
		}

		return &bozdo.TxData{
			Data:         data,
			Value:        value,
			ContractAddr: c.Cfg.Velocore.Router,
		}, nil

	} else if req.ToToken == v1.Token_ETH {
		amOut := amsOut[1]
		amSlip, err := defi.Slippage(amOut, req.Slippage)
		if err != nil {
			return nil, err
		}
		amOut = amSlip
		data, err := abiSource.Pack("swapExactTokensForETH", req.Amount, amOut, path, w.WalletAddr, DefaultDeadLine())
		if err != nil {
			return nil, err
		}

		return &bozdo.TxData{
			Data:         data,
			Value:        value,
			ContractAddr: c.Cfg.Velocore.Router,
		}, nil
	}

	return nil, errors.New("unsupported input params")
}
