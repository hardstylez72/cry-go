package private

import (
	"github.com/hardstylez72/cry/internal/exchange/okex/driver/events"
	"github.com/hardstylez72/cry/internal/exchange/okex/driver/models/account"
	"github.com/hardstylez72/cry/internal/exchange/okex/driver/models/trade"
)

type (
	Account struct {
		Arg      *events.Argument   `json:"arg"`
		Balances []*account.Balance `json:"data"`
	}
	Position struct {
		Arg       *events.Argument    `json:"arg"`
		Positions []*account.Position `json:"data"`
	}
	BalanceAndPosition struct {
		Arg                 *events.Argument              `json:"arg"`
		BalanceAndPositions []*account.BalanceAndPosition `json:"data"`
	}
	Order struct {
		Arg    *events.Argument `json:"arg"`
		Orders []*trade.Order   `json:"data"`
	}
)
