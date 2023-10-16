package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/bridge/layerzero"
	"github.com/hardstylez72/cry/internal/defi/contracts/optimism_fee"
	"github.com/hardstylez72/cry/internal/defi/contracts/stargate/router"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type StargateBridgeSwapTokenReq = StargateBridgeSwapReq

type StargateBridgeSwapToken struct {
	Tx    *types.Transaction
	ECost *bozdo.EstimatedGasCost
}

func (c *EtheriumClient) StargateBridgeSwapToken(ctx context.Context, req *StargateBridgeSwapTokenReq) (*StargateBridgeSwapToken, error) {
	details := []bozdo.TxDetail{}
	contractAddr := c.Cfg.Dict.Stargate.StargateRouterAddress

	if err := req.Validate(c.Cfg.Network); err != nil {
		return nil, err
	}

	wt, err := NewWalletTransactor(req.WalletPk)
	if err != nil {
		return nil, err
	}
	tr, err := router.NewRouterTransactor(contractAddr, c.Cli)
	if err != nil {
		return nil, err
	}

	opt, err := bind.NewKeyedTransactorWithChainID(wt.PrivateKey, c.Cfg.NetworkId)
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

	fee.Fee1 = bozdo.BigIntSum(fee.Fee1, bozdo.Percent(fee.Fee1, layerzero.LayerZeroBoostPercent))

	opt.Value = fee.Fee1
	destChainId := layerzero.LayerZeroChainMap[req.DestChain]
	srcPoolId := PoolIdMap[c.Cfg.Network][req.FromToken]
	distPoolId := PoolIdMap[req.DestChain][req.ToToken]

	amSlip, err := Slippage(req.Quantity, req.Slippage)
	if err != nil {
		return nil, err
	}

	opt.NoSend = req.EstimateOnly

	l1Gasfee := big.NewInt(0)
	if c.Cfg.Network == v1.Network_OPTIMISM {
		optFeeCaller, err := optimism_fee.NewStorageCaller(common.HexToAddress("0x420000000000000000000000000000000000000F"), c.Cli)
		if err != nil {
			return nil, err
		}

		abi, err := router.RouterMetaData.GetAbi()
		if err != nil {
			return nil, err
		}
		data, err := abi.Pack("swap",
			destChainId,
			big.NewInt(srcPoolId),
			big.NewInt(distPoolId),
			wt.WalletAddr,
			req.Quantity,
			amSlip,
			router.IStargateRouterlzTxObj{
				DstGasForCall:   big.NewInt(0),
				DstNativeAmount: big.NewInt(0),
				DstNativeAddr:   common.HexToAddress("0x0000000000000000000000000000000000000001").Bytes(),
			},
			wt.WalletAddr.Bytes(),
			[]byte{},
		)
		if err != nil {
			return nil, err
		}
		o := &bind.CallOpts{}
		o.Context = ctx
		l1Gasfee, err = optFeeCaller.GetL1Fee(o, data)
		if err != nil {
			return nil, err
		}
		details = append(details, bozdo.NewL1FeeDetails(l1Gasfee, c.Cfg.Network, v1.Token_ETH))
	}

	details = append(details, bozdo.NewLZFeeDetails(fee.Fee1, c.Cfg.Network, req.FromToken))

	opt.Value = bozdo.BigIntSum(opt.Value, l1Gasfee)
	opt = c.ResoleGas(ctx, req.Gas, opt)

	// 38.677058
	tx, err := tr.Swap(
		opt,
		destChainId,
		big.NewInt(srcPoolId),
		big.NewInt(distPoolId),
		wt.WalletAddr,
		req.Quantity,
		amSlip,
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
		ECost: Estimate(tx, nil, "bridge", details),
	}, nil
}
