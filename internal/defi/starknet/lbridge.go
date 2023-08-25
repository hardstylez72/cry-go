package starknet

import (
	"context"
	"math/big"

	defi "github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/starknet/halper"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func (c *Client) LiquidityBridge(ctx context.Context, req *defi.LiquidityBridgeReq, taskType v1.TaskType) (*bozdo.DefaultRes, error) {
	if taskType != v1.TaskType_StarkNetBridge {
		return nil, errors.New("unsupported task type: " + taskType.String())
	}

	if req.From == v1.Network_Etherium && req.To == v1.Network_StarkNet {

		res, err := c.halper.LiquidityBridge(ctx, &halper.LiquidityBridgeReq{
			Proxy:        c.cfg.Proxy,
			PKEth:        req.PkFrom,
			PKStark:      req.PkTo,
			Percent:      req.Percent,
			EstimateOnly: req.EstimateOnly,
			SubType:      req.ToSubType.String(),
			FromNetwork:  req.From.String(),
			ToNetwork:    req.To.String(),
		})
		if err != nil {
			return nil, err
		}

		r := &bozdo.DefaultRes{}

		gas, ok := big.NewInt(0).SetString(res.Gas.Total, 10)
		if !ok {
			return nil, errors.New("invalid big int: " + res.Gas.Total)
		}

		limit, ok := big.NewInt(0).SetString(res.Gas.Limit, 10)
		if !ok {
			return nil, errors.New("invalid big int: " + res.Gas.Total)
		}

		price, ok := big.NewInt(0).SetString(res.Gas.Price, 10)
		if !ok {
			return nil, errors.New("invalid big int: " + res.Gas.Total)
		}

		if res.TxHash != nil {
			r.Tx = &bozdo.Transaction{
				Code:    bozdo.CodeBridge,
				Network: v1.Network_Etherium,
				Id:      *res.TxHash,
				Url:     "https://etherscan.io/tx/" + *res.TxHash,
				Details: []bozdo.TxDetail{bozdo.NewTxFee(gas, v1.Network_Etherium, v1.Token_ETH)},
			}
		}

		r.ECost = &bozdo.EstimatedGasCost{
			Type:        bozdo.TxTypeDynamic,
			Name:        "bridge",
			GasLimit:    limit,
			GasPrice:    price,
			TotalGasWei: gas,
			Details:     []bozdo.TxDetail{bozdo.NewTxFee(gas, v1.Network_Etherium, v1.Token_ETH)},
		}

		return r, nil
	}

	return nil, errors.New("usupported pair: " + req.From.String() + " -> " + req.To.String())
}
