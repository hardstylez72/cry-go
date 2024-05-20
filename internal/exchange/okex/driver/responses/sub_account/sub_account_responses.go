package sub_account

import (
	"github.com/hardstylez72/cry/internal/exchange/okex/driver/models/account"
	models "github.com/hardstylez72/cry/internal/exchange/okex/driver/models/subaccount"
	"github.com/hardstylez72/cry/internal/exchange/okex/driver/responses"
)

type (
	ViewList struct {
		responses.Basic
		SubAccounts []*models.SubAccount `json:"data,omitempty"`
	}
	APIKey struct {
		responses.Basic
		APIKeys []*models.APIKey `json:"data,omitempty"`
	}
	GetBalance struct {
		responses.Basic
		Balances []*account.Balance `json:"data,omitempty"`
	}
	HistoryTransfer struct {
		responses.Basic
		HistoryTransfers []*models.HistoryTransfer `json:"data,omitempty"`
	}
	ManageTransfer struct {
		responses.Basic
		Transfers []*models.Transfer `json:"data,omitempty"`
	}
)
