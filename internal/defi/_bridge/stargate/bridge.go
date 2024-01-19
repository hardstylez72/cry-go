package stargate

import (
	"context"
	"net/http"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/lzscan"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type Bridge struct {
	Cli     *defi.EtheriumClient
	HttpCli *http.Client
}

func NewBridge(cli *defi.EtheriumClient) *Bridge {
	return &Bridge{
		Cli:     cli,
		HttpCli: &http.Client{},
	}
}

func (a *Bridge) Bridge(ctx context.Context, req *defi.DefaultBridgeReq) (*bozdo.DefaultRes, error) {

	switch req.FromToken {
	case v1.Token_ETH:
		return a.BridgeETH(ctx, req)
	case v1.Token_STG:
		return a.BridgeSTG(ctx, req)
	default:
		return a.BridgeToken(ctx, req)
	}
}

func (a *Bridge) WaitForConfirm(ctx context.Context, txId string, taskType v1.TaskType, receiver string) error {
	s := lzscan.NewService()

	_, err := s.WaitConfirm(ctx, txId)
	return err
}
