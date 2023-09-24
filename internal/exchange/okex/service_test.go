package okex

import (
	"context"
	"strconv"
	"testing"

	"github.com/google/uuid"
	okex "github.com/hardstylez72/cry/internal/exchange/okex/driver"
	requests "github.com/hardstylez72/cry/internal/exchange/okex/driver/requests/rest/trade"
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

	id := strconv.Itoa(int(uuid.New().ID()))
	println("id: " + id)
	cid := strconv.Itoa(int(uuid.New().ID()))
	println("cid: " + id)

	res, err := s.cli.Rest.Trade.PlaceOrder(ctx, []requests.PlaceOrder{
		{
			InstID:     "ETH-USDT",
			ClOrdID:    cid,   //клиентский
			ReduceOnly: false, // -
			Sz:         0.01,
			TdMode:     okex.TradeCashMode,
			Side:       okex.OrderSell,
			OrdType:    okex.OrderMarket,
			TgtCcy:     "", // quote_ccy for buy, base_ccy
		},
	})

	o, err := s.cli.Rest.Trade.GetOrderDetail(requests.OrderDetails{
		InstID:  "ETH-USDT",
		ClOrdID: res.PlaceOrders[0].ClOrdID,
	})
	println(len(o.Orders))
	//opt, err := s.GetExchangeWithdrawOptions(context.Background(), &v1.GetExchangeWithdrawOptionsRequest{})
	//assert.NoError(t, err)
	//assert.NotNil(t, opt)
	//
	//for i := range opt.Tokens {
	//	l := opt.Tokens[i]
	//	if l.Token == "USDT" {
	//		println(1)
	//	}
	//}
	//
	//wd, err := s.Withdraw(context.Background(), &exchange.WithdrawRequest{
	//	ToAddress: "0x4A6e7c137a6691D55693CA3Bc7E5C698d9d43815",
	//	Amount:    "10",
	//	Network:   "USDT-Arbitrum one",
	//	Token:     "USDT",
	//})
	//assert.NoError(t, err)
	//
	//txId, err := s.WaitConfirm(context.Background(), wd.WithdrawId)
	//assert.NoError(t, err)
	//assert.NotNil(t, txId)
}
