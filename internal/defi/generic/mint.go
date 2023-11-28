package generic

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/nft/zerius"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type minter struct {
	cli *defi.EtheriumClient
}

func NewMinter(cli *defi.EtheriumClient) *minter {
	return &minter{
		cli: cli,
	}
}

func (c *minter) Mint(ctx context.Context, req *defi.SimpleReq, taskType v1.TaskType) (*bozdo.DefaultRes, error) {

	var txData *bozdo.TxData
	var err error

	switch taskType {
	case v1.TaskType_MerklyMintAndBridgeNFT, v1.TaskType_MintMerkly:
		return c.cli.MintMerkly(ctx, req)
	case v1.TaskType_MintZerius:
		cli := zerius.NewClient(c.cli)
		txData, err = cli.MakeMintTx(ctx, req)
	default:
		return nil, errors.New("not supported")
	}
	if err != nil {
		return nil, errors.Wrap(err, "defi.Mint")
	}

	opt := &defi.TxOpt{
		NoSend:   req.EstimateOnly,
		Pk:       req.PK,
		Gas:      req.Gas,
		Debug:    req.Debug,
		TaskType: taskType,
	}

	return c.cli.London(ctx, c.cli, opt, txData)
}
