package starknet

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi/starknet/halper"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type ApproveReq struct {
	Token       v1.Token
	PK          string
	Amount      *big.Int
	SpenderAddr string
}

type ApproveRes struct {
	TxId *string
}

func (c *Client) Approve(ctx context.Context, req *ApproveReq) (*ApproveRes, error) {
	res, err := c.halper.Approve(ctx, &halper.ApproveReq{
		BaseTx: halper.BaseTx{
			ChainRPC:   MainnetRPC,
			PrivateKey: req.PK,
			Proxy:      c.cfg.Proxy,
		},
		Spender: req.SpenderAddr,
		Amount:  req.Amount.String(),
		Token:   req.Token.String(),
	})
	if err != nil {
		return nil, err
	}

	if res.Error != nil {
		return nil, errors.New("approve error: " + *res.Error)
	}

	return &ApproveRes{
		TxId: res.TxId,
	}, nil
}
