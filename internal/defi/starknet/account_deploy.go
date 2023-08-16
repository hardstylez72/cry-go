package starknet

import (
	"context"
	"fmt"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/starknet/halper"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/starknet.go/gateway"
	"github.com/hardstylez72/cry/starknet.go/types"
	"github.com/pkg/errors"
)

type DeployAccountReq struct {
	PK string

	bozdo.BaseReq
}

func (c *Client) IsAccountDeployed(ctx context.Context, pk string) (*bool, error) {
	res, err := c.halper.IsAccountDeployed(ctx, &halper.IsAccountDeployedReq{
		ChainRPC:   MainnetRPC,
		PrivateKey: pk,
		Proxy:      c.cfg.Proxy,
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

func (c *Client) WaitTxComplete(ctx context.Context, tx string) error {
	acceptedOnL2 := false
	var receipt *gateway.TransactionReceipt
	var err error
	fmt.Println("Polling until transaction is accepted on L2...")
	for !acceptedOnL2 {
		_, receipt, err = c.GW.WaitForTransaction(ctx, tx, 5, 10)
		if err != nil {
			return fmt.Errorf("Transaction Failure (%s): can't poll to desired status: %s", tx, err.Error())
		}
		if receipt.Status == types.TransactionAcceptedOnL2 {
			acceptedOnL2 = true
		}
	}
	return nil
}
