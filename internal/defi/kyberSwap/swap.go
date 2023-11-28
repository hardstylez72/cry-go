package kyberswap

import (
	"context"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type KyberSwapMaker struct {
	TokenMap map[v1.Token]common.Address
	CliHttp  *http.Client
	ChainId  *big.Int
	Addr     common.Address
	Network  v1.Network
}

func (c *KyberSwapMaker) makeBaseURL(network v1.Network) string {
	switch network {
	case v1.Network_ZKSYNCERA:
		return "https://aggregator-api.kyberswap.com/zksync"
	default:
		return "unknown"
	}
}
func (c *KyberSwapMaker) MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (_ *bozdo.TxData, err error) {

	var quote *QuoteRes
	quoteCount := 0
	for {
		quote, err = c.Quote(ctx, req)
		if err != nil {
			if quoteCount > 5 {
				return nil, errors.Wrap(err, "Quote")
			}
			quoteCount++

			time.Sleep(time.Second * 5)
			continue
		}
		break
	}

	var prepared *BuildRes
	buildCount := 0
	for {
		prepared, err = c.Build(ctx, req, quote)
		if err != nil {
			if buildCount > 5 {
				return nil, errors.Wrap(err, "Build")
			}
			buildCount++

			time.Sleep(time.Second * 5)
			continue
		}
		break
	}

	if prepared.Code != 0 {
		return nil, errors.New("failed")
	}

	d := prepared.Data.Data[2:]
	data := common.Hex2Bytes(d)
	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		var ok bool
		value, ok = big.NewInt(0).SetString(prepared.Data.AmountIn, 10)
		if !ok {
			return nil, errors.New("invalid input amount")
		}
	}

	return &bozdo.TxData{
		Data:         data,
		Value:        value,
		ContractAddr: common.HexToAddress(prepared.Data.RouterAddress),
		Details:      bozdo.NewKyperSwap(prepared.Data.AmountInUsd, prepared.Data.AmountOutUsd, prepared.Data.OutputChange.Percent),
		//Rate: defi.CalcRate(req.FromToken, req.ToToken, req.Amount, out),
	}, nil
}
