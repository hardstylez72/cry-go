package traderjoe

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type GetSwapDataReq struct {
	FromToken v1.Token
	ToToken   v1.Token
	ChainRPC  string
	Amount    *big.Int
	Recipient common.Address
}

type GetSwapDataRes struct {
	ContractAddr string `json:"contractAddr"`
	Value        string `json:"value"`
	Data         string `json:"data"`
}

type TraderJoe interface {
	GetSwapData(ctx context.Context, req *GetSwapDataReq) (*GetSwapDataRes, error)
}
