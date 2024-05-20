package starknet

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/starknet/halper"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type DeployAccountReq struct {
	PK      string
	SubType v1.ProfileSubType

	bozdo.BaseReq
}

func (c *Client) IsAccountDeployed(ctx context.Context, pk string, sub v1.ProfileSubType) (*bool, error) {
	res, err := c.halper.IsAccountDeployed(ctx, &halper.IsAccountDeployedReq{
		ChainRPC:   MainnetRPC,
		PrivateKey: pk,
		Proxy:      c.cfg.Proxy,
		SubType:    sub.String(),
	})
	if err != nil {
		return nil, err
	}
	return &res.Deployed, nil
}

type DeployAccountRes = bozdo.DefaultRes

// site https://starkscan.co/tx/0x6ea856e3209e04658d99729a8f485131b0944b64d5145b7924ae5a1f59f3b19
func (c *Client) DeployAccount(ctx context.Context, req *DeployAccountReq) (*DeployAccountRes, error) {

	res := &DeployAccountRes{}

	tx, err := c.halper.DeployAccount(ctx, &halper.DeployAccountReq{
		ChainRPC:     MainnetRPC,
		PrivateKey:   req.PK,
		Proxy:        c.cfg.Proxy,
		EstimateOnly: req.EstimateOnly,
		SubType:      req.SubType.String(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "DeployAccount error happened")
	}

	fee, ok := big.NewInt(0).SetString(tx.EstimatedMaxFee, 10)
	if !ok {
		return nil, errors.New("invalid fee value: " + tx.EstimatedMaxFee)
	}

	res.ECost = &bozdo.EstimatedGasCost{
		Type:        bozdo.TxTypeStarkNet,
		Name:        "deploy account",
		TotalGasWei: fee,
	}

	if tx.TxHash != nil && tx.ContractAddr != nil {
		res.Tx = c.NewTx(bozdo.CodeContract, *tx.TxHash, nil)
	}

	return res, err
}

func (c *Client) NewTx(code bozdo.TxCode, id string, details []bozdo.TxDetail) *bozdo.Transaction {
	return &bozdo.Transaction{
		Code:    code,
		Network: v1.Network_StarkNet,
		Id:      id,
		Url:     c.TxViewFn(id),
		Details: details,
	}
}
