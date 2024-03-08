package zksyncera

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/zksyncera/abi/eralend"
	"github.com/hardstylez72/cry/internal/log"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type EraLend struct {
	cli *Client
}

func NewEraLendEraLend(cli *Client) *EraLend {
	return &EraLend{
		cli: cli,
	}
}

func (c *EraLend) LP(ctx context.Context, req *defi.LPReq) (*defi.LPRes, error) {
	return c.lp(ctx, req)
}

func (c *EraLend) lp(ctx context.Context, req *defi.LPReq) (*defi.LPRes, error) {

	if req.Add {
		return c.Supply(ctx, req)
	} else {
		return c.Withdraw(ctx, req)
	}
}

func (c *EraLend) Supply(ctx context.Context, req *defi.LPReq) (*defi.LPRes, error) {

	result := &defi.LPRes{}

	if len(req.Tokens) != 1 {
		return nil, errors.New("invalid input tokens...")
	}
	t := req.Tokens[0]
	tokenAddr := common.HexToAddress("")
	if t == v1.Token_ETH {
		tokenAddr = common.HexToAddress("0x22D8b71599e14F20a49a397b88c1C878c86F5579")
	}

	abi, err := eralend.LpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	pack, err := abi.Pack("mint")
	if err != nil {
		return nil, err
	}

	if req.Debug {
		log.Log.Debug(zap.String("data", "0x"+common.Bytes2Hex(pack)))
	}

	_, w, err := c.cli.Wallet(req.PK)
	if err != nil {
		return nil, errors.Wrap(err, "defi.NewWalletTransactor")
	}

	tx := CreateFunctionCallTransaction(
		w.WalletAddr,
		tokenAddr,
		big.NewInt(0),
		big.NewInt(0),
		req.Amount,
		pack,
		nil, nil,
	)

	raw, estimate, err := c.cli.Make712Tx(ctx, tx, req.Gas, w.Signer)
	if err != nil {
		return nil, fmt.Errorf("Make712Tx: %w", err)
	}

	result.ECost = estimate

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.cli.ClientL2.SendRawTransaction(ctx, raw)
	if err != nil {
		return nil, errors.Wrap(err, "rpcL2.SendRawTransaction")
	}

	result.Tx = c.cli.NewTx(hash, defi.CodeContract, nil)

	return result, nil
}

func (c *EraLend) Withdraw(ctx context.Context, req *defi.LPReq) (*defi.LPRes, error) {
	result := &defi.LPRes{}

	_, w, err := c.cli.Wallet(req.PK)
	if err != nil {
		return nil, errors.Wrap(err, "defi.NewWalletTransactor")
	}

	t := req.Tokens[0]
	tokenAddr := common.HexToAddress("")
	if t == v1.Token_ETH {
		tokenAddr = common.HexToAddress("0x22D8b71599e14F20a49a397b88c1C878c86F5579")
	}

	call, err := eralend.NewLpCaller(tokenAddr, c.cli.ClientL2)
	if err != nil {
		return nil, err
	}

	//supply, err := call.TotalSupply(nil)
	//if err != nil {
	//	return nil, err
	//}
	//
	//reserves, err := call.TotalReserves(nil)
	//if err != nil {
	//	return nil, err
	//}

	balance, err := call.BalanceOf(nil, w.WalletAddr)
	if err != nil {
		return nil, err
	}

	//gg := big.NewInt(0).Div(supply, reserves)
	//println(gg)
	// 201158463

	abi, err := eralend.LpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	pack, err := abi.Pack("redeem", balance)
	if err != nil {
		return nil, err
	}

	if req.Debug {
		log.Log.Debug(zap.String("data", "0x"+common.Bytes2Hex(pack)))
	}

	tx := CreateFunctionCallTransaction(
		w.WalletAddr,
		tokenAddr,
		big.NewInt(0),
		big.NewInt(0),
		nil,
		pack,
		nil, nil,
	)

	raw, estimate, err := c.cli.Make712Tx(ctx, tx, req.Gas, w.Signer)
	if err != nil {
		return nil, fmt.Errorf("Make712Tx: %w", err)
	}

	result.ECost = estimate

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.cli.ClientL2.SendRawTransaction(ctx, raw)
	if err != nil {
		return nil, errors.Wrap(err, "rpcL2.SendRawTransaction")
	}

	result.Tx = c.cli.NewTx(hash, defi.CodeContract, nil)

	return result, nil

}
