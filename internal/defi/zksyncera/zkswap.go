package zksyncera

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/zkswaprouter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type zkSwapMaker struct {
	*Client
}

func (c *Client) ZkSwap(ctx context.Context, req *defi.DefaultSwapReq) (*defi.DefaultRes, error) {
	return c.GenericSwap(ctx, &zkSwapMaker{c}, req)
}

func (c *zkSwapMaker) MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (*defi.TxData, error) {

	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	w, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}
	call, err := zkswaprouter.NewStorageCaller(c.Cfg.ZkSwap.Router, c.Provider)
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

	abispacefirouter, err := zkswaprouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	if req.FromToken == v1.Token_ETH {
		amOut := amsOut[1]
		amOut = defi.Slippage(amOut, req.Slippage)
		data, err := abispacefirouter.Pack("swapExactETHForTokens", amOut, path, w.WalletAddr, DefaultDeadLine())
		if err != nil {
			return nil, err
		}

		return &defi.TxData{
			Data:         data,
			Value:        value,
			ContractAddr: c.Cfg.ZkSwap.Router,
		}, nil

	} else if req.ToToken == v1.Token_ETH {
		amOut := amsOut[1]
		amOut = defi.Slippage(amOut, req.Slippage)
		data, err := abispacefirouter.Pack("swapExactTokensForETH", req.Amount, amOut, path, w.WalletAddr, DefaultDeadLine())
		if err != nil {
			return nil, err
		}

		return &defi.TxData{
			Data:         data,
			Value:        value,
			ContractAddr: c.Cfg.ZkSwap.Router,
		}, nil
	}

	return nil, errors.New("unsupported input params")
}
