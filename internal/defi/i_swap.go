package defi

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type DefaultSwapReq struct {
	Network      v1.Network
	Amount       *big.Int
	FromToken    v1.Token
	ToToken      v1.Token
	WalletPK     string
	EstimateOnly bool
	Gas          *bozdo.Gas
	Slippage     SlippagePercent
	Debug        bool
	SubType      v1.ProfileSubType
	ExchangeRate *float64
}

type Swapper interface {
	Networker
	Swap(ctx context.Context, req *DefaultSwapReq, taskType v1.TaskType) (*bozdo.DefaultRes, error)
}

func IsStablePool(a, b v1.Token) bool {
	m := map[v1.Token]bool{
		v1.Token_USDT: true,
		v1.Token_USDC: true,
		v1.Token_USDp: true,
		v1.Token_BUSD: true,
		v1.Token_LSD:  true,
		v1.Token_LUSD: true,
	}

	_, aok := m[a]
	_, bok := m[b]

	return aok && bok
}
