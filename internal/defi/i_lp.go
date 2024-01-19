package defi

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type LPReq struct {
	Amount *big.Int
	Tokens []v1.Token

	PK  string
	Add bool

	EstimateOnly bool
	Gas          *bozdo.Gas
	Debug        bool

	PSubType v1.ProfileSubType
	Network  v1.Network
}

type LPRes struct {
	Tx        *bozdo.Transaction
	Approves  []bozdo.Transaction
	ECost     *bozdo.EstimatedGasCost
	TxDetails []bozdo.TxDetail
}

type LP interface {
	LP(ctx context.Context, req *LPReq, taskType v1.TaskType) (*LPRes, error)
	Networker
}
