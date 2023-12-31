package defi

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type DefaultBridgeReq struct {
	FromNetwork  v1.Network
	ToNetwork    v1.Network
	PK           string
	Amount       *big.Int
	FromToken    Token
	ToToken      Token
	Gas          *bozdo.Gas
	Slippage     SlippagePercent
	EstimateOnly bool
	Debug        bool
	SubType      v1.ProfileSubType
}

type Bridger interface {
	Networker
	Bridge(ctx context.Context, req *DefaultBridgeReq, taskType v1.TaskType) (*bozdo.DefaultRes, error)
	WaitForConfirm(ctx context.Context, txId string, taskType v1.TaskType, receiver string) error
}
