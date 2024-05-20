package okex

import (
	"context"
	"testing"

	"github.com/hardstylez72/cry/internal/exchange"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	c := tests.GetConfig()

	s, err := NewService(context.Background(), &Config{
		ApiKey:     c.OkexApiKey,
		SecretKey:  c.OkexSecret,
		PassPhrase: c.OkexPassPhrase,
		HttpClient: c.Cli,
	})
	assert.NoError(t, err)
	ctx := context.Background()

	req := &exchange.SwapReq{
		From:      v1.Token_ETH,
		To:        v1.Token_USDC,
		AmPercent: 10,
	}

	//err = s.Before(ctx, req)
	//assert.NoError(t, err)

	res, err := s.Swap(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	err = s.WaitSwapComplete(ctx, &exchange.SwapRes{
		Pair:    res.Pair,
		TradeId: res.TradeId,
	})
	assert.NoError(t, err)

	err = s.After(ctx, req)
	assert.NoError(t, err)

}
