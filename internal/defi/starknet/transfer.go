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

func (c *Client) Transfer(ctx context.Context, r *defi.TransferReq) (*defi.TransferRes, error) {

	res, err := c.halper.Transfer(ctx, &halper.TransferReq{
		BaseTx: c.makeHalperBase(r.Pk, r.PSubType, r.EstimateOnly, r.Gas),
		ToAddr: r.ToAddr,
		Token:  r.Token.String(),
		Amount: r.Amount.String(),
	})
	if err != nil {
		return nil, err
	}

	result := &bozdo.DefaultRes{}

	if res.TxHash != nil {
		result.Tx = c.NewTx(bozdo.CodeTransfer, *res.TxHash, nil)
	}

	if res.EstimatedMaxFee != "" {

		t, ok := big.NewInt(0).SetString(res.EstimatedMaxFee, 10)
		if ok {
			result.ECost = &bozdo.EstimatedGasCost{
				Type:        bozdo.TxTypeStarkNet,
				Name:        "transfer",
				GasLimit:    big.NewInt(0),
				GasPrice:    big.NewInt(0),
				TotalGasWei: t,
				Details:     nil,
			}
		}
	}

	return &defi.TransferRes{
		Tx:    result.Tx,
		ECost: result.ECost,
	}, nil

}

func (c *Client) makeHalperBase(pk string, subType v1.ProfileSubType, estimateOnly bool, gas *bozdo.Gas) halper.BaseTx {
	out := halper.BaseTx{
		ChainRPC:     gateway.MAINNET_BASE,
		PrivateKey:   pk,
		Proxy:        c.cfg.Proxy,
		SubType:      subType.String(),
		EstimateOnly: estimateOnly,
	}

	if gas != nil {
		out.MaxFee = gas.TotalGas.String()
	}

	return out
}
