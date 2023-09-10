package pub

import (
	"context"
	"sync"
	"time"

	"github.com/hardstylez72/cry/internal/exchange/binance"
	"github.com/hardstylez72/cry/internal/lib"
	"github.com/hardstylez72/cry/internal/log"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type Listener struct {
	cli *binance.Ws
}

var p = Pairs{
	AVAX:  14,
	BNB:   250,
	ETH:   1900,
	MATIC: 0.7,
}

type Pairs struct {
	AVAX  float64
	BNB   float64
	ETH   float64
	MATIC float64
}

func Price() *Pairs {
	return &p
}

func NewIListener() *Listener {
	return &Listener{
		cli: binance.NewWsClient(),
	}
}

func (l *Listener) Listen(ctx context.Context) {

	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		for {

			d, err := l.GetData()
			if err != nil {
				log.Log.Error("pub.Listen.GetData", err)
			} else {
				p = *d
			}

			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
			}
		}
	}()

}

func (l *Listener) GetData() (*Pairs, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	bnb, err := l.cli.GetPairRate(ctx, "BNBUSDT")
	if err != nil {
		return nil, err
	}

	eth, err := l.cli.GetPairRate(ctx, "ETHUSDT")
	if err != nil {
		return nil, err
	}

	avax, err := l.cli.GetPairRate(ctx, "AVAXUSDT")
	if err != nil {
		return nil, err
	}

	matic, err := l.cli.GetPairRate(ctx, "MATICUSDT")
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		rate, err := l.cli.GetPairRate(ctx, "ETHUSDC")
		if err == nil {
			exchangeRateMap.Set(key(v1.Token_ETH, v1.Token_USDC), rate)
		}
	}()

	go func() {
		defer wg.Done()
		rate, err := l.cli.GetPairRate(ctx, "ETHUSDT")
		if err == nil {
			exchangeRateMap.Set(key(v1.Token_ETH, v1.Token_USDT), rate)
		}
	}()

	go func() {
		defer wg.Done()
		rate, err := l.cli.GetPairRate(ctx, "USDCUSDT")
		if err == nil {
			exchangeRateMap.Set(key(v1.Token_USDC, v1.Token_USDT), rate)
		}
	}()

	wg.Wait()

	return &Pairs{
		AVAX:  avax,
		BNB:   bnb,
		ETH:   eth,
		MATIC: matic,
	}, nil
}

func key(t1, t2 v1.Token) string {
	return t1.String() + t2.String()
}

var exchangeRateMap = lib.NewMap[string, float64]()

func GetExchangeRate(t1, t2 v1.Token) *float64 {
	rate, ok := exchangeRateMap.Get(t1.String() + t2.String())
	if ok {
		if rate < 1 {
			rate = 1 / rate
		}

		return &rate
	}

	rate, ok = exchangeRateMap.Get(t2.String() + t1.String())
	if ok {
		if rate < 1 {
			rate = 1 / rate
		}
		return &rate
	}
	return nil
}
