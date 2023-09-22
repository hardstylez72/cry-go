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

func (c *Client) Swap(ctx context.Context, req *defi.DefaultSwapReq, platform v1.TaskType) (*bozdo.DefaultRes, error) {

	res, err := c.halper.Swap(ctx, c.castHalperSwapReq(req, platform))
	if err != nil {
		return nil, err
	}

	result := c.CastHalperSwapRes(res)
	if req.ExchangeRate != nil {
		result.TxDetails = append(result.TxDetails, bozdo.NewSwapRateRatio(*req.ExchangeRate, res.Rate))
	}

	result.ECost.Details = result.TxDetails
	return result, nil
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
func (c *Client) castHalperSwapReq(req *defi.DefaultSwapReq, platform v1.TaskType) *halper.DefaultSwapReq {
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
		SubType:      req.SubType.String(),
	}

	if req.Gas.RuleSet() {
		result.Fee = req.Gas.TotalGas.String()
	}

	return result
}
