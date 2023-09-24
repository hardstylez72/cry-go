package exchange

import (
	"context"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type WithdrawId = string

type WithdrawRequest struct {
	ToAddress string
	Amount    string
	Network   string
	Token     string
}

type WithdrawResponse struct {
	WithdrawId WithdrawId
}

type Withdrawer interface {
	GetExchangeWithdrawOptions(ctx context.Context, req *v1.GetExchangeWithdrawOptionsRequest) (*v1.GetExchangeWithdrawOptionsResponse, error)
	Withdraw(ctx context.Context, req *WithdrawRequest) (*WithdrawResponse, error)
	WaitConfirm(ctx context.Context, id WithdrawId) (*string, error)
	Ping(ctx context.Context) error
	GetFundingBalance(ctx context.Context, coin string) (float64, error)
	WithdrawStatus(ctx context.Context, withdrawId string) (string, error)
}

type SwapReq struct {
	From      v1.Token
	To        v1.Token
	AmPercent float64
}

type SwapRes struct {
	Pair    string
	TradeId string
}

type Swapper interface {
	Before(ctx context.Context, r *SwapReq) error
	Swap(ctx context.Context, r *SwapReq) (*SwapRes, error)
	WaitSwapComplete(ctx context.Context, s *SwapRes) error
	After(ctx context.Context, r *SwapReq) error
}
