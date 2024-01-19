package stargate

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/_bridge/layerzero"
	"github.com/hardstylez72/cry/internal/defi/_bridge/stargate/abi/stg"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func (c *Bridge) BridgeSTG(ctx context.Context, req *defi.DefaultBridgeReq) (_ *bozdo.DefaultRes, _ error) {

	ca, ok := c.Cli.Cfg.TokenMap[v1.Token_STG]
	if !ok {
		return nil, errors.New("no stg address")
	}

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

	abi, err := stg.StgMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	pack, err := abi.Pack("sendTokens",
		destChainId,
		w.WalletAddr.Bytes(),
		req.Amount,
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
		[]byte{},
	)
	if err != nil {
		return nil, err
	}

	txData := &bozdo.TxData{
		Data:         pack,
		Value:        fee.Fee1,
		ContractAddr: ca,
		Details:      []bozdo.TxDetail{bozdo.NewLZFeeDetails(fee.Fee1, req.FromNetwork, req.FromToken)},
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
