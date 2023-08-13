package starknet

import (
	"context"
	"math/big"

	"github.com/dontpanicdao/caigo/types"
	"github.com/hardstylez72/cry/internal/defi"
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

	ca, ok := c.TokenMap[req.Token]
	if !ok {
		return nil, defi.ErrTokenNotSupportedFn(req.Token)
	}

	tx := types.FunctionCall{
		ContractAddress:    types.HexToHash(ca),
		EntryPointSelector: "balanceOf",
		Calldata:           []string{req.WalletAddress},
	}

	res, err := c.GW.Call(ctx, tx, "")
	if err != nil {
		return nil, err
	}

	var balance *big.Int
	if len(res) == 2 {
		balance = types.HexToBN(res[0])
	}

	return c.getBalance(balance, req), nil
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

	// todo:///
	return &defi.GetBalanceRes{
		Req:           req,
		WEI:           wei,
		ETHER:         nil,
		HumanReadable: "",
		Float:         0,
	}
}