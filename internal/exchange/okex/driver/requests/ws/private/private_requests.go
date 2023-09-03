package private

import "github.com/hardstylez72/cry/internal/exchange/okex/driver"

type (
	Account struct {
		Ccy string `json:"ccy,omitempty"`
	}
	Position struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		InstType okex.InstrumentType `json:"instType"`
	}
	Order struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		InstType okex.InstrumentType `json:"instType"`
	}
	AlgoOrder struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		InstType okex.InstrumentType `json:"instType"`
	}
)
