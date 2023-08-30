package starknet

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/starknet/halper"
	"github.com/hardstylez72/cry/starknet.go/gateway"
)

func (c *Client) Mint(ctx context.Context, req *defi.SimpleReq) (*bozdo.DefaultRes, error) {

	res, err := c.halper.Mint(ctx, &halper.MintReq{
		BaseTx: halper.BaseTx{
			ChainRPC:     gateway.MAINNET_BASE,
			PrivateKey:   req.PK,
			Proxy:        c.cfg.Proxy,
			SubType:      req.SubType.String(),
			EstimateOnly: req.EstimateOnly,
		},
	})
	if err != nil {
		return nil, err
	}

	result := &bozdo.DefaultRes{}
	wei, ok := big.NewInt(0).SetString(res.EstimatedMaxFee, 10)
	if ok {
		result.ECost = &bozdo.EstimatedGasCost{
			Type:        bozdo.TxTypeStarkNet,
			Name:        "mint",
			GasLimit:    big.NewInt(0),
			GasPrice:    big.NewInt(0),
			TotalGasWei: wei,
			Details:     nil,
		}
	}

	if res.TxHash != nil {
		result.Tx = c.NewTx(bozdo.CodeMint, *res.TxHash, nil)
	}

	return result, nil
}
