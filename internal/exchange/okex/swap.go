package okex

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	okex "github.com/hardstylez72/cry/internal/exchange/okex/driver"
	"github.com/hardstylez72/cry/internal/exchange/okex/driver/requests/rest/funding"
	requests "github.com/hardstylez72/cry/internal/exchange/okex/driver/requests/rest/trade"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type SwapReq struct {
	From      v1.Token
	To        v1.Token
	AmPercent float64
}

func (s *Service) Swap(ctx context.Context, req *SwapReq) error {

	from := "ETH"
	to := "USDC"

	b, err := s.GetFundingBalance(ctx, from)
	if err != nil {
		return errors.Wrap(err, "GetFundingBalance")
	}

	am := b * req.AmPercent / 100

	if err := s.fromFundingToTrading(ctx, from, am); err != nil {
		return errors.Wrap(err, "fromFundingToTrading")
	}

	if err := s.PlaceOrder(ctx, req.From, req.To, am); err != nil {
		return errors.Wrap(err, "PlaceOrder")
	}

	if err := s.AllBalanceToFounding(ctx, to); err != nil {
		return errors.Wrap(err, "AllBalanceToTrading")
	}
	return nil
}

type PlaceOrderRes struct {
	InstId        string
	CustomOrderId string
}

func (s *Service) PlaceOrder(ctx context.Context, from, to v1.Token, am float64) (*PlaceOrderRes, error) {
	id := strconv.Itoa(int(uuid.New().ID()))

	inst, err := s.tradePair(from, to)
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
			Side:       okex.OrderSell,
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

func (s *Service) tradePair(from, to v1.Token) (string, error) {

	if (from == v1.Token_ETH && to == v1.Token_USDC) || (from == v1.Token_USDC && to == v1.Token_ETH) {
		return "ETH-USDC", nil
	}

	return "", errors.New("unsupported pair: " + from.String() + " " + to.String())
}

func (s *Service) WaitOrderConplete(ctx context.Context) {
	detail, err := s.cli.Rest.Trade.GetOrderDetail(ctx, requests.OrderDetails{
		InstID:  "",
		OrdID:   "",
		ClOrdID: "",
	})
	if err != nil {
		return err
	}
}

func (s *Service) AllBalanceToTrading(ctx context.Context, Ccy string) error {
	b, err := s.GetFundingBalance(ctx, Ccy)
	if err != nil {
		return errors.Wrap(err, "GetFundingBalance")
	}

	if err := s.fromFundingToTrading(ctx, Ccy, b); err != nil {
		return errors.Wrap(err, "fromFundingToTrading")
	}
	return nil
}

func (s *Service) fromFundingToTrading(ctx context.Context, Ccy string, Amount float64) error {
	_, err := s.cli.Rest.Funding.FundsTransfer(ctx, funding.FundsTransfer{
		Ccy:  Ccy,
		Amt:  Amount,
		Type: okex.TransferWithinAccount,
		From: okex.FundingAccount,
		To:   okex.SpotAccount,
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
		From: okex.SpotAccount,
		To:   okex.FundingAccount,
	})
	if err != nil {
		return err
	}

	return nil
}
