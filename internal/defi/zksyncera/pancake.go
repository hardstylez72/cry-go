package zksyncera

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/pancakeswap"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type pancakeMaker struct {
	CA  common.Address
	Cli *Client
}

func (c *Client) PancakeSwap(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.DefaultRes, error) {
	return c.GenericSwap(ctx, &pancakeMaker{
		CA:  common.HexToAddress("0x5aEaF2883FBf30f3D62471154eDa3C0c1b05942d"),
		Cli: c,
	}, req)
}

func (c *pancakeMaker) MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {
	call, err := pancakeswap.NewRouterCaller(c.CA, c.Cli.ClientL2)
	if err != nil {
		return nil, err
	}

	fromToken, supported := c.Cli.Cfg.TokenMap[req.FromToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(req.FromToken)
	}
	toToken, supported := c.Cli.Cfg.TokenMap[req.ToToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(req.FromToken)
	}

	path := []common.Address{
		fromToken,
		toToken,
	}

	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	w, err := NewWalletTransactor(req.WalletPK, c.Cli.NetworkId)
	if err != nil {
		return nil, err
	}

	opt := &bind.CallOpts{Context: ctx}
	amounts, err := call.GetAmountsOut(opt, req.Amount, path)
	if err != nil {
		return nil, err
	}
	amountOut := amounts[1]

	to := w.WalletAddr

	deadline := DefaultDeadLine()

	a, err := pancakeswap.RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	if req.FromToken == v1.Token_ETH {
		//var amountOutMin *big.Int //++
		//var path []common.Address //+
		//var to common.Address //+
		//var deadline *big.Int //+

		amountOutMin, err := defi.Slippage(amountOut, req.Slippage)
		if err != nil {
			return nil, err
		}

		data, err := a.Pack("swapExactETHForTokens", amountOutMin, path, to, deadline)
		if err != nil {
			return nil, err
		}
		return &bozdo.TxData{
			Data:         data,
			Value:        value,
			ContractAddr: c.CA,
			Details:      nil,
			Rate:         CalcRate(req.FromToken, req.ToToken, req.Amount, amountOut),
		}, nil
	}

	if req.ToToken == v1.Token_ETH {
		//var amountIn *big.Int
		var amountOutMin *big.Int
		//var path []common.Address //+
		//var to common.Address //+
		//var deadline *big.Int //+

		amountIn := req.Amount

		amountOutMin, err := defi.Slippage(amountOut, req.Slippage)
		if err != nil {
			return nil, err
		}
		
		data, err := a.Pack("swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
		if err != nil {
			return nil, err
		}
		return &bozdo.TxData{
			Data:         data,
			Value:        value,
			ContractAddr: c.CA,
			Details:      nil,
			Rate:         CalcRate(req.FromToken, req.ToToken, req.Amount, amountOut),
		}, nil
	}

	return nil, errors.New(fmt.Sprintf("unsupported pair: [%s -> %s]", req.FromToken.String(), req.ToToken.String()))
}
