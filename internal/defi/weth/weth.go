package weth

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type Client struct {
	WethAddr common.Address
}

func (c *Client) MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {
	if req.FromToken == v1.Token_ETH {
		return c.WrapETH(ctx, req)
	}
	return c.UnWrapETH(ctx, req)
}

func (c *Client) WrapETH(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {

	wethabi, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	data, err := wethabi.Pack("deposit")
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw function: %w", err)
	}

	return &bozdo.TxData{
		Data:         data,
		Value:        req.Amount,
		ContractAddr: c.WethAddr,
		Details:      nil,
		Rate:         nil,
		Gas:          nil,
		ExtraFee:     nil,
		Code:         "",
	}, nil
}

func (c *Client) UnWrapETH(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {

	wethabi, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	data, err := wethabi.Pack("withdraw", req.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw function: %w", err)
	}

	return &bozdo.TxData{
		Data:         data,
		Value:        big.NewInt(0),
		ContractAddr: c.WethAddr,
		Details:      nil,
		Rate:         nil,
		Gas:          nil,
		ExtraFee:     nil,
		Code:         "",
	}, nil

}
