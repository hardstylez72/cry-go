package starknet

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/starknet/halper"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/starknet.go/gateway"
)

func (c *Client) Swap10kRouterAddress() string {
	return "0x07a6f98c03379b9513ca84cca1373ff452a7462a3b61598f0af5bb27ad7f76d1"
}
func (c *Client) SithSwapRouterAddress() string {
	return "0x028c858a586fa12123a1ccb337a0a3b369281f91ea00544d0c086524b759f627"
}
func (c *Client) JediSwapRouterAddress() string {
	return "0x041fd22b238fa21cfcf5dd45a8548974d8263b3a531a60388411c5e230f97023"
}
func (c *Client) Swap(ctx context.Context, req *defi.DefaultSwapReq, platform v1.TaskType) (*bozdo.DefaultRes, error) {

	res, err := c.halper.Swap(ctx, c.CastHalperSwapReq(req, platform))
	if err != nil {
		return nil, err
	}

	return c.CastHalperSwapRes(res), nil
}

func (c *Client) CastHalperSwapRes(res *halper.DefaultSwapRes) *bozdo.DefaultRes {
	result := &bozdo.DefaultRes{}

	if res.ApproveTx != nil {
		result.ApproveTx = c.NewTx(bozdo.CodeApprove, *res.ApproveTx, nil)
	}

	if res.SwapTx != nil {
		result.Tx = c.NewTx(bozdo.CodeSwap, *res.SwapTx, nil)
	}

	if res.Fee != nil {

		t, ok := big.NewInt(0).SetString(*res.Fee, 10)
		if ok {
			result.ECost = &bozdo.EstimatedGasCost{
				Type:        bozdo.TxTypeStarkNet,
				Name:        "swap",
				GasLimit:    big.NewInt(0),
				GasPrice:    big.NewInt(0),
				TotalGasWei: t,
				Details:     nil,
			}
		}
	}

	return result
}
func (c *Client) CastHalperSwapReq(req *defi.DefaultSwapReq, platform v1.TaskType) *halper.DefaultSwapReq {
	result := &halper.DefaultSwapReq{
		Proxy:        c.cfg.Proxy,
		ChainRPC:     gateway.MAINNET_BASE,
		PrivateKey:   req.WalletPK,
		FromToken:    req.FromToken.String(),
		ToToken:      req.ToToken.String(),
		Amount:       req.Amount.String(),
		EstimateOnly: req.EstimateOnly,
		Slippage:     req.Slippage,
		Platform:     platform.String(),
	}

	if req.Gas.RuleSet() {
		result.Fee = req.Gas.TotalGas.String()
	}

	return result
}
