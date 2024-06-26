package stargate

import (
	"context"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/_bridge/layerzero"
	"github.com/hardstylez72/cry/internal/defi/_bridge/stargate/abi/stg"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func (c *Bridge) BridgeSTG(ctx context.Context, req *defi.DefaultBridgeReq) (res *bozdo.DefaultRes, _ error) {

	ca, ok := c.Cli.Cfg.TokenMap[v1.Token_STG]
	if !ok {
		return nil, errors.New("no stg address")
	}

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
		return nil, errors.Wrap(err, "TokenLimitChecker1")
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

	fee.Fee1 = bozdo.BigIntSum(fee.Fee1, bozdo.Percent(fee.Fee1, 2))

	destChainId := layerzero.LayerZeroChainMap[req.ToNetwork]

	abi, err := stg.StgMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	ext := []byte{}
	if c.Cli.Network() == v1.Network_Fantom || c.Cli.Network() == v1.Network_Base {
		ext, err = hex.DecodeString("00010000000000000000000000000000000000000000000000000000000000014c08")
		if err != nil {
			return nil, err
		}
	}

	pack, err := abi.Pack("sendTokens",
		destChainId,
		w.WalletAddr.Bytes(),
		req.Amount,
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
		ext,
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
