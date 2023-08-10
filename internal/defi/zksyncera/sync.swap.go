package zksyncera

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/syncswaprouter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type syncSwapMaker struct {
	*Client
}

func (c *Client) SyncSwap(ctx context.Context, req *defi.DefaultSwapReq) (*defi.DefaultRes, error) {
	return c.GenericSwap(ctx, &syncSwapMaker{c}, req)
}

func (c *syncSwapMaker) MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {

	transactor, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	fee, err := c.getSyncSwapFee(ctx, &getSyncSwapFeeReq{
		Network:   req.Network,
		FromToken: req.FromToken,
		ToToken:   req.ToToken,
		Wallet:    transactor.WalletAddr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetSyncSwapFee")
	}

	amOut, err := c.getAmountOut(ctx, &getAmountOutReq{
		Addr:      transactor.WalletAddr,
		FromToken: req.FromToken,
		ToToken:   req.ToToken,
		Amount:    req.Amount,
	})
	if err != nil {
		return nil, errors.Wrap(err, "getAmountOut")
	}

	min, err := defi.Slippage(amOut, req.Slippage)
	if err != nil {
		return nil, err
	}

	structThing, err := abi.NewType("tuple", "struct thing", []abi.ArgumentMarshaling{
		{Name: "a", Type: "address"},
		{Name: "b", Type: "address"},
		{Name: "c", Type: "uint8"},
	})
	if err != nil {
		return nil, errors.Wrap(err, "abi.NewType")
	}

	record := struct {
		A common.Address
		B common.Address
		C uint8
	}{
		A: c.Cfg.TokenMap[req.FromToken],
		B: transactor.WalletAddr,
		C: 1,
	}

	args := abi.Arguments{
		{Type: structThing, Name: "param_one"},
	}
	data, err := args.Pack(&record)
	if err != nil {
		return nil, errors.Wrap(err, "args.Pack(&record)")
	}

	TokenIn := c.Cfg.TokenMap[req.FromToken]

	value := fee.Fee

	if req.FromToken == v1.Token_ETH {
		TokenIn = common.HexToAddress("0x0000000000000000000000000000000000000000")
		value = new(big.Int).Add(fee.Fee, req.Amount)
	}

	path := []syncswaprouter.IRouterSwapPath{
		{
			Steps: []syncswaprouter.IRouterSwapStep{
				{
					Pool:         fee.PoolAddr,                                                      // +
					Data:         data,                                                              //+
					Callback:     common.HexToAddress("0x0000000000000000000000000000000000000000"), //+
					CallbackData: []byte("0x"),                                                      //+
				},
			},
			TokenIn:  TokenIn,    //+
			AmountIn: req.Amount, //+
		},
	}
	ded := time.Now().Add(time.Minute).UnixMilli()

	syncswaprouterabi, err := syncswaprouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	data, err = syncswaprouterabi.Pack("swap", path, min, big.NewInt(ded))
	if err != nil {
		return nil, err
	}

	return &bozdo.TxData{
		Data:         data,
		Value:        value,
		ContractAddr: c.Cfg.SyncSwap.RouterSwap,
	}, nil
}
