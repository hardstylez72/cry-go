package starknet

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/starknet.go/felt"
	"github.com/hardstylez72/cry/starknet.go/rpc"
	"github.com/hardstylez72/cry/starknet.go/types"
	"github.com/hardstylez72/cry/starknet.go/utils"
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

	addrFelt, err := utils.HexToFelt(ca)
	if err != nil {
		return nil, err
	}

	selectorFelt := types.GetSelectorFromNameFelt("balanceOf")
	if err != nil {
		return nil, err
	}

	wFelt, err := utils.HexToFelt(req.WalletAddress)
	if err != nil {
		return nil, err
	}
	tx := rpc.FunctionCall{
		ContractAddress:    addrFelt,
		EntryPointSelector: selectorFelt,
		Calldata:           []*felt.Felt{wFelt},
	}

	res, err := c.GWP.Call(ctx, tx, rpc.BlockID{Tag: "latest"})
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

	f, _ := defi.WeiToToken(wei, req.Token).Float64()

	return &defi.GetBalanceRes{
		Req:           req,
		WEI:           wei,
		ETHER:         defi.WeiToToken(wei, req.Token),
		HumanReadable: defi.WeiToToken(wei, req.Token).String(),
		Float:         f,
	}
}
