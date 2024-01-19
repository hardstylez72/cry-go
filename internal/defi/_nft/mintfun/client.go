package mintfun

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type Client struct {
	ca  common.Address
	cli *defi.EtheriumClient
}

func NewClient(ca common.Address, cli *defi.EtheriumClient) *Client {
	return &Client{
		ca:  ca,
		cli: cli,
	}
}

func (c *Client) Mint(ctx context.Context, req *defi.SimpleReq) (*bozdo.DefaultRes, error) {
	//https://basescan.org/tx/0xa1213a189a3532070cbf76e9c656c297a4ebd456c849a0fd36b5c94037341e57

	data := &bozdo.TxData{
		Data:         common.Hex2Bytes("6a6278420000000000000000000000004a6e7c137a6691d55693ca3bc7e5c698d9d438150021fb3f"),
		ContractAddr: c.ca,
		Code:         bozdo.CodeMint,
	}

	opt := &defi.TxOpt{
		NoSend:   req.EstimateOnly,
		Pk:       req.PK,
		Gas:      req.Gas,
		Debug:    req.Debug,
		TaskType: v1.TaskType_MintFun,
	}
	tx, err := c.cli.London(ctx, c.cli, opt, data)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
