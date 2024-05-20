package zksyncera

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
)

type LP interface {
	GetSyncSwapPool(ctx context.Context, req *GetSyncSwapPoolReq) (*common.Address, error)
	SyncSwapLiquidity(ctx context.Context, req *SyncSwapLiquidityReq) (*SyncSwapLiquidityRes, error)
	SyncSwapLPBalance(ctx context.Context, pool, user common.Address) (*big.Int, error)
	defi.Networker
}
