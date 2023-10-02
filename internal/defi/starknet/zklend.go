package starknet

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/starknet/halper"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func (c *Client) LP(ctx context.Context, req *defi.LPReq, taskType v1.TaskType) (*defi.LPRes, error) {

	switch taskType {
	case v1.TaskType_ZkLandLP:
		return c.zklendLP(ctx, req)
	default:
		return nil, errors.New("unsupported task type: " + taskType.String())
	}
}

func (c *Client) zklendLP(ctx context.Context, req *defi.LPReq) (*defi.LPRes, error) {

	op := "deposit"
	if !req.Add {
		op = "withdraw"
	}

	if len(req.Tokens) == 0 {
		return nil, errors.New("tokens len is zero")
	}

	r := &halper.ZkLendReq{
		BaseTx: c.makeHalperBase(req.PK, req.PSubType, req.EstimateOnly, req.Gas),
		Token:  req.Tokens[0].String(),
		Amount: req.Amount.String(),
		Op:     op,
	}

	res, err := c.halper.ZkLend(ctx, r)
	if err != nil {
		return nil, err
	}

	out := &defi.LPRes{
		Tx:        nil,
		ECost:     nil,
		TxDetails: nil,
	}

	if res.TxHash != nil {
		out.Tx = c.NewTx(bozdo.CodeLP, *res.TxHash, nil)
	}

	if res.EstimatedMaxFee != "" {
		t, ok := big.NewInt(0).SetString(res.EstimatedMaxFee, 10)
		if ok {
			out.ECost = &bozdo.EstimatedGasCost{
				Type:        bozdo.TxTypeStarkNet,
				Name:        "lp",
				GasLimit:    big.NewInt(0),
				GasPrice:    big.NewInt(0),
				TotalGasWei: t,
				Details:     nil,
			}
		}
	}

	return out, nil
}
