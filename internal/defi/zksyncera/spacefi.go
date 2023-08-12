package zksyncera

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/spacefirouter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type spaceFiMaker struct {
	*Client
}

// site ETH -> USDC https://explorer.zksync.io/tx/0xc9d9c33c9bdf7bcbafa45c313eb8b70c1d37f9797d06ce8e20f1fbf046a79ed5
// me ETH -> USDC https://explorer.zksync.io/tx/0xaddad23305e85e9c806621d274ce6ebbeeb4dfb633b415ea0a21cdb27520142b

// site USDC -> ETH https://explorer.zksync.io/tx/0x701ffd2a41613c5ecc6a1cb10ea19b742404a6b0bf3d88db82216d51eebf1b88
// me USDC -> ETH https://explorer.zksync.io/tx/0xee4a8f41ae9f32e69cb3205e88d2493ebedd9deac86b6cea9c247d896dafe590
func (c *Client) SpaceFiSwap(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.DefaultRes, error) {
	return c.GenericSwap(ctx, &spaceFiMaker{c}, req)
}

func (c *spaceFiMaker) MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {
	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	w, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}
	call, err := spacefirouter.NewStorageCaller(c.Cfg.SpaceFI.Router, c.Provider)
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

	path := []common.Address{from, to}
	amsOut, err := call.GetAmountsOut(&bind.CallOpts{Context: ctx}, req.Amount, path)
	if err != nil {
		return nil, err
	}
	if len(amsOut) != 2 {
		return nil, errors.New("errors getting rate")
	}

	abispacefirouter, err := spacefirouter.StorageMetaData.GetAbi()
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
		data, err := abispacefirouter.Pack("swapExactETHForTokens", amOut, path, w.WalletAddr, DefaultDeadLine())
		if err != nil {
			return nil, err
		}

		return &bozdo.TxData{
			Data:         data,
			Value:        value,
			ContractAddr: c.Cfg.SpaceFI.Router,
		}, nil

	} else if req.ToToken == v1.Token_ETH {
		amOut := amsOut[1]

		amSlip, err := defi.Slippage(amOut, req.Slippage)
		if err != nil {
			return nil, err
		}

		amOut = amSlip
		data, err := abispacefirouter.Pack("swapExactTokensForETH", req.Amount, amOut, path, w.WalletAddr, DefaultDeadLine())
		if err != nil {
			return nil, err
		}

		return &bozdo.TxData{
			Data:         data,
			Value:        value,
			ContractAddr: c.Cfg.SpaceFI.Router,
		}, nil
	}

	return nil, errors.New("unsupported input params")
}
