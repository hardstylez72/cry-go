package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type BridgeNFTReq struct {
	WalletPK    string
	FromNetwork v1.Network
	ToNetwork   v1.Network
	NFTId       *big.Int
	bozdo.BaseReq
}

type Minter interface {
	Mint(ctx context.Context, req *SimpleReq, taskType v1.TaskType) (*bozdo.DefaultRes, error)
	Networker
}

type MintAndBridgeNFT interface {
	Networker
	Minter
	BridgeNft(ctx context.Context, req *BridgeNFTReq, taskType v1.TaskType) (*bozdo.DefaultRes, error)
	GetNFTId(ctx context.Context, txHash common.Hash, owner common.Address) (*big.Int, error)
}
