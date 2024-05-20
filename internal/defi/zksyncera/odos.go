package zksyncera

import (
	"context"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/_swap/odos"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
)

func (c *Client) OdosSwap(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.DefaultRes, error) {

	w, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	return c.GenericSwap(ctx, &odos.OdosMaker{
		CA:       common.HexToAddress("0x4bba932e9792a2b917d47830c93a9bc79320e4f7"),
		TokenMap: c.Cfg.TokenMap,
		CliHttp:  &http.Client{},
		ChainId:  c.NetworkId,
		Addr:     w.WalletAddr,
		Network:  c.Network(),
	}, req)
}
