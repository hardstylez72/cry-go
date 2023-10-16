package starknet

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/starknet/halper"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

func (c *Client) GetNetworkToken() v1.Token {
	return c.NativeToken
}

func (c *Client) GetNetworkId() *big.Int {
	// todo:...
	return big.NewInt(0)
}

func (c *Client) GetBalance(ctx context.Context, req *defi.GetBalanceReq) (*defi.GetBalanceRes, error) {

	res, err := c.halper.Balance(ctx, &halper.BalanceReq{
		ChainRPC: c.cfg.RPCEndpoint,
		Proxy:    c.cfg.Proxy,
		Pub:      req.WalletAddress,
		Token:    req.Token.String(),
	})
	if err != nil {
		return nil, err
	}

	b, ok := big.NewInt(0).SetString(res.WEI, 10)
	if !ok {
		b = big.NewInt(0)
	}

	return c.getBalance(b, req), nil
}

func (c *Client) getBalance(wei *big.Int, req *defi.GetBalanceReq) *defi.GetBalanceRes {

	if wei == nil {
		return &defi.GetBalanceRes{
			Req:           req,
			WEI:           big.NewInt(0),
			ETHER:         big.NewFloat(0),
			HumanReadable: "0",
			Float:         0,
		}
	}

	if req.Token == v1.Token_ETH {

		f, _ := defi.WEIToEther(wei).Float64()

		return &defi.GetBalanceRes{
			Req:           req,
			WEI:           wei,
			ETHER:         defi.WEIToEther(wei),
			HumanReadable: defi.WEIToEther(wei).String(),
			Float:         f,
		}
	}

	f, _ := defi.WeiToToken(wei, req.Token).Float64()

	return &defi.GetBalanceRes{
		Req:           req,
		WEI:           wei,
		ETHER:         defi.WeiToToken(wei, req.Token),
		HumanReadable: defi.WeiToToken(wei, req.Token).String(),
		Float:         f,
	}
}
