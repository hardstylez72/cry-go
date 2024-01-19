package woofi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type Client struct {
	CA       common.Address
	caller   bind.ContractCaller
	tokenMap map[v1.Token]common.Address
}

func NewSwapper(ca common.Address, caller bind.ContractCaller, tokenMap map[v1.Token]common.Address) *Client {
	return &Client{
		CA:       ca,
		caller:   caller,
		tokenMap: tokenMap,
	}
}

func (c *Client) MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {
	switch req.Network {
	case v1.Network_Base:
		return c.makeSwapDefault(ctx, req)
	}
	return nil, errors.New("network is not supported")
}

func (c *Client) makeSwapDefault(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {

	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	w, err := defi.NewWalletTransactor(req.WalletPK)
	if err != nil {
		return nil, errors.Wrap(err, "defi.NewWalletTransactor")
	}

	caller, err := NewRouterCaller(c.CA, c.caller)
	if err != nil {
		return nil, errors.Wrap(err, "NewRouterCaller")
	}

	fromToken, supported := c.tokenMap[req.FromToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(req.FromToken)
	}

	toToken, supported := c.tokenMap[req.ToToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(req.FromToken)
	}

	if req.FromToken == v1.Token_ETH {
		fromToken = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
	}

	if req.ToToken == v1.Token_ETH {
		toToken = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
	}

	amOut, err := caller.QuerySwap(&bind.CallOpts{Context: ctx}, fromToken, toToken, req.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "QuerySwap")
	}
	amMin, err := defi.Slippage(amOut, req.Slippage)
	if err != nil {
		return nil, errors.Wrap(err, "defi.Slippage")
	}

	var fromAmount *big.Int
	var minToAmount *big.Int
	var to common.Address
	var rebateTo common.Address

	fromAmount = req.Amount
	minToAmount = amMin
	to = w.WalletAddr
	rebateTo = w.WalletAddr
	abi, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "RouterMetaData.GetAbi")
	}

	data, err := abi.Pack("swap", fromToken, toToken, fromAmount, minToAmount, to, rebateTo)
	if err != nil {
		return nil, errors.Wrap(err, "abi.Pack")
	}

	return &bozdo.TxData{
		Data:         data,
		Value:        value,
		ContractAddr: c.CA,
		Details:      nil,
		Rate:         defi.CalcRate(req.FromToken, req.ToToken, req.Amount, amMin),
		Gas:          nil,
		ExtraFee:     nil,
	}, nil
}
