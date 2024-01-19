package stargate

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/_bridge/layerzero"
	"github.com/hardstylez72/cry/internal/defi/_bridge/stargate/abi/router"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func (c *Bridge) BridgeToken(ctx context.Context, req *defi.DefaultBridgeReq) (res *bozdo.DefaultRes, _ error) {
	details := []bozdo.TxDetail{}
	ca := c.Cli.Cfg.Dict.Stargate.StargateRouterAddress

	w, err := defi.NewWalletTransactor(req.PK)
	if err != nil {
		return nil, err
	}

	limitTx, err := c.Cli.TokenLimitChecker(ctx, &defi.TokenLimitCheckerReq{
		Token:       req.FromToken,
		Wallet:      w,
		Amount:      req.Amount,
		SpenderAddr: ca,
	})
	if err != nil {
		return nil, errors.Wrap(err, "TokenLimitChecker")
	}

	defer func() {
		if limitTx != nil && limitTx.LimitExtended {
			res.ApproveTx = c.Cli.NewTx(limitTx.ApproveTx.Hash(), defi.CodeApprove, nil)
		}
	}()

	fee, err := c.GetStargateBridgeFee(ctx, &GetStargateBridgeFeeReq{
		ToChain: req.ToNetwork,
		Wallet:  w.WalletAddr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetStargateBridgeFee")
	}

	destChainId := layerzero.LayerZeroChainMap[req.ToNetwork]
	srcPoolId := PoolIdMap[c.Cli.Cfg.Network][req.FromToken]
	distPoolId := PoolIdMap[req.ToNetwork][req.ToToken]

	amSlip, err := defi.Slippage(req.Amount, req.Slippage)
	if err != nil {
		return nil, err
	}

	details = append(details, bozdo.NewLZFeeDetails(fee.Fee1, req.FromNetwork, req.FromToken))

	abi, err := router.RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	data, err := abi.Pack("swap",
		destChainId,
		big.NewInt(srcPoolId),
		big.NewInt(distPoolId),
		w.WalletAddr,
		req.Amount,
		amSlip,
		router.IStargateRouterlzTxObj{
			DstGasForCall:   big.NewInt(0),
			DstNativeAmount: big.NewInt(0),
			DstNativeAddr:   common.HexToAddress("0x0000000000000000000000000000000000000001").Bytes(),
		},
		w.WalletAddr.Bytes(),
		[]byte{},
	)
	if err != nil {
		return nil, err
	}

	txData := &bozdo.TxData{
		Data:         data,
		Value:        fee.Fee1,
		ContractAddr: ca,
		Details:      details,
		Code:         bozdo.CodeBridge,
	}

	opt := &defi.TxOpt{
		NoSend:   req.EstimateOnly,
		Pk:       req.PK,
		Gas:      req.Gas,
		Debug:    req.Debug,
		TaskType: v1.TaskType_StargateBridge,
	}

	return c.Cli.London(ctx, c.Cli, opt, txData)
}
