package aave

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/log"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Client struct {
	ca  common.Address
	cli *defi.EtheriumClient
}

func NewClient(ca common.Address, cli *defi.EtheriumClient) *Client {
	return &Client{
		ca:  ca,
		cli: cli,
	}
}

func (c *Client) LP(ctx context.Context, req *defi.LPReq) (*defi.LPRes, error) {
	return c.lp(ctx, req)
}

func (c *Client) lp(ctx context.Context, req *defi.LPReq) (*defi.LPRes, error) {
	if req.Add {
		return c.Supply(ctx, req)
	} else {
		return c.Withdraw(ctx, req)
	}
}

func (c *Client) Supply(ctx context.Context, req *defi.LPReq) (*defi.LPRes, error) {

	out := &defi.LPRes{}

	if len(req.Tokens) != 1 {
		return nil, errors.New("invalid input tokens...")
	}
	a, ok := c.cli.Cfg.TokenMap[req.Tokens[0]]
	if !ok {
		return nil, defi.ErrTokenNotSupportedFn(req.Tokens[0])
	}

	ap, err := c.cli.Approve(ctx, &defi.Approve{
		TokenAddr:   a,
		PK:          req.PK,
		Amount:      req.Amount,
		SpenderAddr: c.ca,
	})
	if err != nil {
		return nil, errors.Wrap(err, "c.cli.Approve")
	}
	if ap.LimitExtended {
		out.Approves = []bozdo.Transaction{*c.cli.NewTx(ap.ApproveTx.Hash(), bozdo.CodeApprove, nil)}
	}

	abi, err := LpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	var asset common.Address
	var amount *big.Int
	var onBehalfOf common.Address
	var referralCode uint16

	w, err := defi.NewWalletTransactor(req.PK)
	if err != nil {
		return nil, errors.Wrap(err, "defi.NewWalletTransactor")
	}
	asset = a
	amount = req.Amount
	onBehalfOf = w.WalletAddr
	referralCode = 0

	pack, err := abi.Pack("supply", asset, amount, onBehalfOf, referralCode)
	if err != nil {
		return nil, err
	}

	if req.Debug {
		log.Log.Debug(zap.String("data", "0x"+common.Bytes2Hex(pack)))
	}

	data := &bozdo.TxData{
		Data:         pack,
		Value:        nil,
		ContractAddr: c.ca,
		Details:      nil,
		Rate:         nil,
		Gas:          nil,
		ExtraFee:     nil,
		Code:         bozdo.CodeLP,
	}

	opt := &defi.TxOpt{
		NoSend:       req.EstimateOnly,
		Pk:           req.PK,
		Gas:          req.Gas,
		ExchangeRate: nil,
		Debug:        false,
		TaskType:     v1.TaskType_AaveLP,
	}
	tx, err := c.cli.London(ctx, c.cli, opt, data)
	if err != nil {
		return nil, err
	}

	out.Tx = tx.Tx
	out.ECost = tx.ECost
	out.TxDetails = tx.TxDetails

	return out, nil
}

func (c *Client) Withdraw(ctx context.Context, req *defi.LPReq) (*defi.LPRes, error) {
	out := &defi.LPRes{}

	if len(req.Tokens) != 1 {
		return nil, errors.New("invalid input tokens...")
	}
	a, ok := c.cli.Cfg.TokenMap[req.Tokens[0]]
	if !ok {
		return nil, defi.ErrTokenNotSupportedFn(req.Tokens[0])
	}

	var asset common.Address
	var amount *big.Int
	var to common.Address

	w, err := defi.NewWalletTransactor(req.PK)
	if err != nil {
		return nil, errors.Wrap(err, "defi.NewWalletTransactor")
	}

	to = w.WalletAddr
	amount, _ = big.NewInt(0).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
	asset = a

	abi, err := LpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	pack, err := abi.Pack("withdraw", asset, amount, to)
	if err != nil {
		return nil, err
	}

	data := &bozdo.TxData{
		Data:         pack,
		ContractAddr: c.ca,
		Code:         bozdo.CodeLP,
	}

	opt := &defi.TxOpt{
		NoSend:       req.EstimateOnly,
		Pk:           req.PK,
		Gas:          req.Gas,
		ExchangeRate: nil,
		Debug:        false,
		TaskType:     v1.TaskType_AaveLP,
	}
	tx, err := c.cli.London(ctx, c.cli, opt, data)
	if err != nil {
		return nil, err
	}

	out.Tx = tx.Tx
	out.ECost = tx.ECost
	out.TxDetails = tx.TxDetails

	return out, nil

}
