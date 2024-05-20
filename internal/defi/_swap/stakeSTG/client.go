package stakestg

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type Client struct {
	Cli *defi.EtheriumClient
	CA  common.Address
}

func NewClient(Cli *defi.EtheriumClient, CA common.Address) *Client {
	return &Client{
		Cli: Cli,
		CA:  CA,
	}
}

func (c *Client) LP(ctx context.Context, req *defi.LPReq) (*defi.LPRes, error) {

	w, err := defi.NewWalletTransactor(req.PK)
	if err != nil {
		return nil, err
	}

	approves := make([]bozdo.Transaction, 0)
	for _, token := range req.Tokens {

		limitTx, err := c.Cli.TokenLimitChecker(ctx, &defi.TokenLimitCheckerReq{
			Token:       token,
			Wallet:      w,
			Amount:      req.Amount,
			SpenderAddr: c.CA,
		})
		if err != nil {
			return nil, errors.Wrap(err, "TokenLimitChecker")
		}

		if limitTx != nil && limitTx.LimitExtended {
			approves = append(approves, *c.Cli.NewTx(limitTx.ApproveTx.Hash(), defi.CodeApprove, nil))
		}
	}

	tx, err := c.MakeTx(ctx, req)
	if err != nil {
		return nil, err
	}

	opt := &defi.TxOpt{
		NoSend:   req.EstimateOnly,
		Pk:       req.PK,
		Gas:      req.Gas,
		Debug:    req.Debug,
		TaskType: v1.TaskType_StakeSTG,
	}

	res, err := c.Cli.London(ctx, c.Cli, opt, tx)
	if err != nil {
		return nil, err
	}
	return &defi.LPRes{
		Tx:        res.Tx,
		Approves:  approves,
		ECost:     res.ECost,
		TxDetails: res.TxDetails,
	}, nil
}

func (c *Client) MakeTx(ctx context.Context, req *defi.LPReq) (*bozdo.TxData, error) {

	var _value *big.Int //+
	var _unlock_time *big.Int

	_unlock_time = big.NewInt(time.Now().AddDate(2, 11, 0).Unix())
	_value = req.Amount

	abi, err := StakerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	pack, err := abi.Pack("create_lock", _value, _unlock_time)
	if err != nil {
		return nil, err
	}

	return &bozdo.TxData{
		Data:         pack,
		Value:        big.NewInt(0),
		ContractAddr: c.CA,
		Details:      nil,
		Rate:         nil,
		Gas:          nil,
		ExtraFee:     nil,
		Code:         bozdo.CodeLP,
	}, nil
}
