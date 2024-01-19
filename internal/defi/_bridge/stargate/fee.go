package stargate

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi/_bridge/layerzero"
	"github.com/hardstylez72/cry/internal/defi/_bridge/stargate/abi/router"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type GetStargateBridgeFeeReq struct {
	ToChain v1.Network
	Wallet  common.Address
	Retry   int
}

type GetStargateBridgeFeeRes struct {
	Fee1 *big.Int
	Fee2 *big.Int
}

func (c *Bridge) GetStargateBridgeFee(ctx context.Context, req *GetStargateBridgeFeeReq) (*GetStargateBridgeFeeRes, error) {
	trx, err := router.NewRouterCaller(c.Cli.Cfg.Dict.Stargate.StargateRouterAddress, c.Cli.Cli)
	if err != nil {
		return nil, err
	}

	toAddress := req.Wallet.Bytes()
	payload := common.HexToAddress("0").Bytes()
	distChain, ok := layerzero.LayerZeroChainMap[req.ToChain]
	if !ok {
		return nil, errors.New("invalid chain: " + string(req.ToChain))
	}

	opt := &bind.CallOpts{
		Context: ctx,
	}
	fee1, fee2, err := trx.QuoteLayerZeroFee(opt, distChain, layerzero.TypeFuncSwap, toAddress, payload, router.IStargateRouterlzTxObj{
		DstGasForCall:   big.NewInt(0),
		DstNativeAmount: big.NewInt(0),
		DstNativeAddr:   []byte{},
	})
	if err != nil {
		return nil, err
	}

	return &GetStargateBridgeFeeRes{
		Fee1: fee1,
		Fee2: fee2,
	}, nil
}
