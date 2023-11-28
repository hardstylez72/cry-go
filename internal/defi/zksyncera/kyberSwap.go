package zksyncera

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	kyberswap "github.com/hardstylez72/cry/internal/defi/kyberSwap"
	"github.com/hardstylez72/cry/internal/socks5"
)

func (c *Client) KeyberSwap(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.DefaultRes, error) {

	w, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	p, err := socks5.NewSock5ProxyString("", "")
	if err != nil {
		return nil, err
	}

	return c.GenericSwap(ctx, &kyberswap.KyberSwapMaker{
		TokenMap: c.Cfg.TokenMap,
		CliHttp:  p.Cli,
		ChainId:  c.NetworkId,
		Addr:     w.WalletAddr,
		Network:  c.Network(),
	}, req)
}
