package zksyncera

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/weth"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

func (c *Client) WETH(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.DefaultRes, error) {
	return c.GenericSwap(ctx, &weth.Client{
		WethAddr: c.Cfg.TokenMap[v1.Token_WETH],
	}, req)
}
