package starknet

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/starknet/halper"
	"github.com/hardstylez72/cry/starknet.go/gateway"
)

func (c *Client) Swap10k(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.DefaultRes, error) {

	res, err := c.halper.Swap10k(ctx, c.CastHalperSwapReq(req))
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
func (c *Client) CastHalperSwapReq(req *defi.DefaultSwapReq) *halper.DefaultSwapReq {
	result := &halper.DefaultSwapReq{
		Proxy:        c.cfg.Proxy,
		ChainRPC:     gateway.MAINNET_BASE,
		PrivateKey:   req.WalletPK,
		FromToken:    req.FromToken.String(),
		ToToken:      req.ToToken.String(),
		Amount:       req.Amount.String(),
		EstimateOnly: req.EstimateOnly,
	}

	if req.Gas.RuleSet() {
		result.Fee = req.Gas.TotalGas.String()
	}

	return result
}
