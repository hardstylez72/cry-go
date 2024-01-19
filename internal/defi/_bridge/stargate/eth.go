package stargate

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/_bridge/layerzero"
	"github.com/hardstylez72/cry/internal/defi/_bridge/stargate/abi/routereth"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func (c *Bridge) BridgeETH(ctx context.Context, req *defi.DefaultBridgeReq) (_ *bozdo.DefaultRes, _ error) {
	details := []bozdo.TxDetail{}

	ca := c.Cli.Cfg.Dict.Stargate.StargateRouterEthAddress

	w, err := defi.NewWalletTransactor(req.PK)
	if err != nil {
		return nil, err
	}

	fee, err := c.GetStargateBridgeFee(ctx, &GetStargateBridgeFeeReq{
		ToChain: req.ToNetwork,
		Wallet:  w.WalletAddr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetStargateBridgeFee")
	}

	destChainId := layerzero.LayerZeroChainMap[req.ToNetwork]

	details = append(details, bozdo.NewLZFeeDetails(fee.Fee1, c.Cli.Cfg.Network, v1.Token_ETH))

	amSlip, err := defi.Slippage(req.Amount, req.Slippage)
	if err != nil {
		return nil, err
	}
	abi, err := routereth.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	pack, err := abi.Pack("SwapETH",
		destChainId,
		w.WalletAddr,
		w.WalletAddr.Bytes(),
		req.Amount,
		amSlip,
	)
	if err != nil {
		return nil, err
	}

	txData := &bozdo.TxData{
		Data:         pack,
		Value:        bozdo.BigIntSum(req.Amount, fee.Fee1),
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
