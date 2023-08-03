package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hardstylez72/cry/internal/defi/contracts/stargate/router"
	"github.com/pkg/errors"
)

type StargateBridgeSwapTokenReq = StargateBridgeSwapReq

type StargateBridgeSwapToken struct {
	Tx    *types.Transaction
	ECost *EstimatedGasCost
}

func (c *EtheriumClient) StargateBridgeSwapToken(ctx context.Context, req *StargateBridgeSwapTokenReq) (*StargateBridgeSwapToken, error) {

	if err := req.Validate(c.Cfg.Network); err != nil {
		return nil, err
	}

	wt, err := NewWalletTransactor(req.WalletPk)
	if err != nil {
		return nil, err
	}
	tr, err := router.NewRouterTransactor(c.Cfg.Dict.Stargate.StargateRouterAddress, c.Cli)
	if err != nil {
		return nil, err
	}

	chainID, err := c.GetNetworkId(ctx)
	if err != nil {
		return nil, err
	}

	opt, err := bind.NewKeyedTransactorWithChainID(wt.PrivateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	opt.Context = ctx

	fee, err := c.GetStargateBridgeFee(ctx, &GetStargateBridgeFeeReq{
		ToChain: req.DestChain,
		Wallet:  wt.WalletAddr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetStargateBridgeFee")
	}

	opt.Value = BigIntSum(fee.Fee1, Percent(fee.Fee1, 2))
	destChainId := ChainIdMap[req.DestChain]
	srcPoolId := PoolIdMap[c.Cfg.Network][req.FromToken]
	distPoolId := PoolIdMap[req.DestChain][req.ToToken]
	min := Slippage(req.Quantity, SlippagePercent05)
	opt.NoSend = req.EstimateOnly

	opt = c.ResoleGas(ctx, req.Gas, opt)

	// 38.677058
	tx, err := tr.Swap(
		opt,
		destChainId,
		big.NewInt(srcPoolId),
		big.NewInt(distPoolId),
		wt.WalletAddr,
		req.Quantity,
		min,
		router.IStargateRouterlzTxObj{
			DstGasForCall:   big.NewInt(0),
			DstNativeAmount: big.NewInt(0),
			DstNativeAddr:   common.HexToAddress("0x0000000000000000000000000000000000000001").Bytes(),
		},
		wt.WalletAddr.Bytes(),
		[]byte{},
	)
	if err != nil {
		return nil, errors.Wrap(err, "tr.Swap")
	}

	return &StargateBridgeSwapToken{
		Tx:    tx,
		ECost: Estimate(tx),
	}, nil
}
