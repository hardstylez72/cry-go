package okex

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/hardstylez72/cry/internal/exchange"
	okex "github.com/hardstylez72/cry/internal/exchange/okex/driver"
	"github.com/hardstylez72/cry/internal/exchange/okex/driver/requests/rest/funding"
	requests "github.com/hardstylez72/cry/internal/exchange/okex/driver/requests/rest/trade"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func token(in v1.Token) (string, error) {
	switch in {
	case v1.Token_ETH:
		return "ETH", nil
	case v1.Token_USDC:
		return "USDC", nil
	case v1.Token_USDT:
		return "USDT", nil
	default:
		return "", errors.New("unsupported token: " + in.String())
	}
}

func (s *Service) tradePair(from, to v1.Token) (string, okex.OrderSide, error) {

	if (from == v1.Token_ETH && to == v1.Token_USDC) || (from == v1.Token_USDC && to == v1.Token_ETH) {

		var side = okex.OrderBuy
		if from == v1.Token_ETH && to == v1.Token_USDC {
			side = okex.OrderSell
		}

		return "ETH-USDC", side, nil
	}

	if (from == v1.Token_ETH && to == v1.Token_USDT) || (from == v1.Token_USDT && to == v1.Token_ETH) {
		var side = okex.OrderBuy
		if from == v1.Token_ETH && to == v1.Token_USDT {
			side = okex.OrderSell
		}
		return "ETH-USDT", side, nil
	}

	return "", "", errors.New("unsupported pair: " + from.String() + " " + to.String())
}

func (s *Service) Before(ctx context.Context, req *exchange.SwapReq) error {

	token, err := token(req.From)
	if err != nil {
		return errors.Wrap(err, "token")
	}

	b, err := s.GetFundingBalance(ctx, token)
	if err != nil {
		return errors.Wrap(err, "GetFundingBalance")
	}

	am := b
	if req.AmPercent != 100 {
		am = b * req.AmPercent / 100
	}

	if err := s.fromFundingToTrading(ctx, token, am); err != nil {
		return errors.Wrap(err, "fromFundingToTrading")
	}

	return nil
}
func (s *Service) Swap(ctx context.Context, req *exchange.SwapReq) (*exchange.SwapRes, error) {

	token, err := token(req.From)
	if err != nil {
		return nil, errors.Wrap(err, "token")
	}

	b, err := s.GetTradingBalance(ctx, token)
	if err != nil {
		return nil, errors.Wrap(err, "GetTradingBalance")
	}

	res, err := s.PlaceOrder(ctx, req.From, req.To, b)
	if err != nil {
		return nil, errors.Wrap(err, "PlaceOrder")
	}

	return &exchange.SwapRes{
		Pair:    res.InstId,
		TradeId: res.CustomOrderId,
	}, nil
}
func (s *Service) WaitSwapComplete(ctx context.Context, res *exchange.SwapRes) error {
	detail, err := s.cli.Rest.Trade.GetOrderDetail(ctx, requests.OrderDetails{
		InstID:  res.Pair,
		OrdID:   "",
		ClOrdID: res.TradeId,
	})
	if err != nil {
		return err
	}

	for _, el := range detail.Orders {
		if el.ClOrdID == res.TradeId {
			if el.State == okex.OrderFilled {
				return nil
			}
		}
	}

	return errors.New("order is not found")
}
func (s *Service) After(ctx context.Context, req *exchange.SwapReq) error {

	token, err := token(req.To)
	if err != nil {
		return errors.Wrap(err, "token")
	}

	if err := s.AllBalanceToFounding(ctx, token); err != nil {
		return errors.Wrap(err, "AllBalanceToFounding")
	}
	return nil
}

type PlaceOrderRes struct {
	InstId        string
	CustomOrderId string
}

func (s *Service) PlaceOrder(ctx context.Context, from, to v1.Token, am float64) (*PlaceOrderRes, error) {
	id := strconv.Itoa(int(uuid.New().ID()))

	inst, side, err := s.tradePair(from, to)
	if err != nil {
		return nil, errors.Wrap(err, "tradePair")
	}

	res, err := s.cli.Rest.Trade.PlaceOrder(ctx, []requests.PlaceOrder{
		{
			InstID:     inst,
			ClOrdID:    id,    //клиентский
			ReduceOnly: false, // -
			Sz:         am,
			TdMode:     okex.TradeCashMode,
			Side:       side,
			OrdType:    okex.OrderMarket,
			TgtCcy:     "", // quote_ccy for buy, base_ccy
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "Rest.Trade.PlaceOrder")
	}

	if res.Code != 0 {
		return nil, errors.New("bug")
	}

	return &PlaceOrderRes{
		InstId:        inst,
		CustomOrderId: id,
	}, nil

}

func (s *Service) fromFundingToTrading(ctx context.Context, Ccy string, Amount float64) error {

	_, err := s.cli.Rest.Funding.FundsTransfer(ctx, funding.FundsTransfer{
		Ccy:  Ccy,
		Amt:  Amount,
		Type: okex.TransferWithinAccount,
		From: okex.FundingAccount,
		To:   18,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AllBalanceToFounding(ctx context.Context, Ccy string) error {
	b, err := s.GetTradingBalance(ctx, Ccy)
	if err != nil {
		return errors.Wrap(err, "GetTradingBalance")
	}

	if err := s.fromTradingToFunding(ctx, Ccy, b); err != nil {
		return errors.Wrap(err, "fromTradingToFunding")
	}
	return nil
}

func (s *Service) fromTradingToFunding(ctx context.Context, Ccy string, Amount float64) error {
	_, err := s.cli.Rest.Funding.FundsTransfer(ctx, funding.FundsTransfer{
		Ccy:  Ccy,
		Amt:  Amount,
		Type: okex.TransferWithinAccount,
		From: 18,
		To:   okex.FundingAccount,
	})
	if err != nil {
		return err
	}

	return nil
}
