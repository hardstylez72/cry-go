package binance

import (
	"context"

	"github.com/adshao/go-binance/v2"
	_ "github.com/adshao/go-binance/v2"
	"github.com/hardstylez72/cry/internal/lib"
	"github.com/hardstylez72/cry/internal/server/config"
	"github.com/hardstylez72/cry/internal/tests"
)

type Ws struct {
	cli *binance.Client
}

func NewWsClient() *Ws {
	client := binance.NewClient("", "")
	if config.CFG.Env == config.Local {
		client.HTTPClient = tests.GetConfig().Cli
	}

	return &Ws{client}
}

func (ws *Ws) GetPairRate(ctx context.Context, pair string) (float64, error) {
	res, err := ws.cli.NewAveragePriceService().Symbol(pair).Do(ctx)
	if err != nil {
		return 0, err
	}

	return lib.StringToFloat(res.Price)
}
