package scanner

import (
	"context"
	"math/big"
	"sort"
	"time"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/exchange/pub"
	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Transaction struct {
	To string
	TS int64
}

type Transfer struct {
	From   string
	Token  string
	To     string
	Amount *big.Int
}

type Scanner interface {
	Transactions(ctx context.Context, add string) ([]Transaction, error)
	Transfers(ctx context.Context, add string) ([]Transfer, error)
	SpecMap() map[string]v1.TaskType
	Token(a string) (*v1.Token, bool)
	Service(a string) (*Service, bool)
}

type Service struct {
	Name string
	Url  string
}

type address = string
type tokenAddress = string

type WAGG struct {
	TotalValue  string
	Interaction map[address]*Interaction
	Txs         []Transaction
	ActivePeriods
}

type Interaction struct {
	Addr   string
	Amount map[address]*big.Int
	Count  int64
}

func Stat(ctx context.Context, s Scanner, add string) (*v1.Stat, error) {

	transfers, err := s.Transfers(ctx, add)
	if err != nil {
		return nil, errors.Wrap(err, "Transfers")
	}

	transfers = filterTransfers(transfers, add)

	txs, err := s.Transactions(ctx, add)
	if err != nil {
		return nil, errors.Wrap(err, "Transactions")
	}

	sort.Slice(txs, func(i, j int) bool {
		return txs[i].TS > txs[j].TS
	})

	ap := GetActivePeriods(txs)

	Interaction := getInteractions(transfers)
	totalAmountUsd := 0.0

	interactions := make([]*v1.Interaction, 0)
	for _, el := range Interaction {

		if el.Addr == add {
			continue
		}

		ams := make([]*v1.InteractionAmount, 0)

		interactionAmount := 0.0
		for token, am := range el.Amount {

			var t v1.Token
			usd := ""
			tt, ok := s.Token(token)
			if ok {
				t = *tt
				usd, ok = CastTokenAmountToUSD(*tt, am)
				if !ok {
					continue
				}

				f, _ := lib.StringToFloat(usd)
				interactionAmount += f
			}

			ams = append(ams, &v1.InteractionAmount{
				Token:     &t,
				TokenUrl:  "https://explorer.zksync.io/address/" + token,
				AmountWei: am.String(),
				AmountUsd: &usd,
			})
		}

		totalAmountUsd += interactionAmount

		service := Service{
			Name: "noname",
			Url:  "",
		}
		v, ok := s.Service(el.Addr)
		if ok {
			service = *v
		}

		interactions = append(interactions, &v1.Interaction{
			To:          el.Addr,
			ToUrl:       "https://explorer.zksync.io/address/" + el.Addr,
			ServiceName: service.Name,
			ServiceUrl:  &service.Url,
			Amounts:     ams,
			Txs:         el.Count,
			AmountUsd:   interactionAmount,
		})

	}

	var la *timestamppb.Timestamp
	if ap.LastActivity != nil {
		la = timestamppb.New(time.Unix(*ap.LastActivity, 0))
	}

	interactions = groupByServiceName(interactions)

	return &v1.Stat{
		ActiveDays:     ap.Days,
		ActiveMonths:   ap.Months,
		LastActivity:   la,
		Interactions:   interactions,
		TotalUsdAmount: totalAmountUsd,
		TxCount:        int64(len(txs)),
		UniqAddress:    int64(len(interactions)),
	}, nil
}

func filterTransfers(in []Transfer, addr string) []Transfer {

	out := make([]Transfer, 0)
	for _, el := range in {
		if el.From != addr {
			continue
		}

		if el.To == addr {
			continue
		}

		out = append(out, el)
	}
	return out
}

func groupByServiceName(in []*v1.Interaction) []*v1.Interaction {
	out := make([]*v1.Interaction, 0)

	noname := make([]*v1.Interaction, 0)

	m := make(map[string]*v1.Interaction)
	for _, i := range in {
		if i.ServiceName == "noname" {
			noname = append(noname, i)
			continue
		}
		_, ok := m[i.ServiceName]
		if !ok {
			m[i.ServiceName] = i
		} else {
			m[i.ServiceName].Txs += i.Txs
			m[i.ServiceName].AmountUsd += i.AmountUsd
		}
	}

	for _, i := range m {
		out = append(out, i)
	}

	out = append(out, noname...)
	return out
}

func CastTokenAmountToUSD(token v1.Token, am *big.Int) (string, bool) {

	switch token {
	case v1.Token_LUSD, v1.Token_BUSD:
		out := defi.EthToUsd(defi.WEIToEther(am), 1).String()
		return out, true
	case v1.Token_USDC, v1.Token_USDT, v1.Token_USDCBridged:
		out := big.NewInt(0).Div(am, big.NewInt(1e6)).String()
		return out, true
	case v1.Token_ETH, v1.Token_WETH:
		price := pub.Price().ETH
		out := defi.EthToUsd(defi.WEIToEther(am), price).String()
		return out, true
	default:
		return "", false
	}
}

func getInteractions(in []Transfer) map[address]*Interaction {
	m := make(map[address]*Interaction)

	for _, el := range in {
		_, ok := m[el.To]
		if ok {
			m[el.To].Count++
			_, ex := m[el.To].Amount[el.Token]
			if ex {
				m[el.To].Amount[el.Token] = big.NewInt(0).Add(m[el.To].Amount[el.Token], el.Amount)
			} else {
				m[el.To].Amount[el.Token] = el.Amount
			}
		} else {

			m[el.To] = &Interaction{
				Addr: el.To,
				Amount: map[address]*big.Int{
					el.Token: el.Amount,
				},
				Count: 1,
			}
		}
	}

	return m
}

type ActivePeriods struct {
	LastActivity *int64
	Days         *int64
	Months       *int64
}

func GetActivePeriods(txs []Transaction) ActivePeriods {

	var LastActivity *int64
	var days *int64
	var months *int64

	if len(txs) > 0 {
		LastActivity = &txs[len(txs)-1].TS
		dayMap := map[int]bool{}
		monthMap := map[time.Month]bool{}
		for _, el := range txs {
			t := time.Unix(el.TS, 0)
			dayMap[t.Day()] = true
			monthMap[t.Month()] = true
		}

		tmp := int64(len(dayMap))
		days = &tmp

		tmp1 := int64(len(monthMap))
		months = &tmp1
	}

	return ActivePeriods{
		LastActivity: LastActivity,
		Days:         days,
		Months:       months,
	}
}
