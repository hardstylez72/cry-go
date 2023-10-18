package odos

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type OdosMaker struct {
	CA       common.Address
	TokenMap map[v1.Token]common.Address
	CliHttp  *http.Client
	ChainId  *big.Int
	Addr     common.Address
	Network  v1.Network
}

func (c *OdosMaker) MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {

	quote, err := c.Quote(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "Quote")
	}

	prepared, err := c.Assemble(ctx, req, quote)
	if err != nil {
		return nil, errors.Wrap(err, "Assemble")
	}

	if len(prepared.OutputTokens) == 0 {
		return nil, errors.New("output token amount not found")
	}
	outString := prepared.OutputTokens[0].Amount
	out, ok := big.NewInt(0).SetString(outString, 10)
	if !ok {
		return nil, errors.New("invalid output amount: " + prepared.OutputTokens[0].Amount)
	}
	data := common.Hex2Bytes(strings.ReplaceAll(prepared.Transaction.Data, "0x", ""))
	v, ok := big.NewInt(0).SetString(prepared.Transaction.Value, 10)
	if !ok {
		return nil, errors.New("invalid value: " + prepared.Transaction.Value)
	}

	if len(prepared.InputTokens) == 0 {
		return nil, errors.New("input token amount not found")
	}
	inString := prepared.InputTokens[0].Amount

	value := v

	return &bozdo.TxData{
		Data:         data,
		Value:        value,
		ContractAddr: c.CA,
		Details:      bozdo.NewOdos(inString, outString, c.Network, req.FromToken, req.ToToken, quote.PercentDiff),
		Rate:         defi.CalcRate(req.FromToken, req.ToToken, req.Amount, out),
		Gas:          big.NewInt(int64(quote.GasEstimate)),
	}, nil
}

type QuoteReq struct {
	ChainId     int `json:"chainId"`
	InputTokens [1]struct {
		TokenAddress string `json:"tokenAddress"`
		Amount       string `json:"amount"`
	} `json:"inputTokens"`
	OutputTokens [1]struct {
		TokenAddress string `json:"tokenAddress"`
		Proportion   int    `json:"proportion"`
	} `json:"outputTokens"`
	UserAddr             string  `json:"userAddr"`
	SlippageLimitPercent float64 `json:"slippageLimitPercent"`
	ReferralCode         int     `json:"referralCode"`
	Compact              bool    `json:"compact"`
}

type QuoteRes struct {
	InTokens          []string    `json:"inTokens"`
	OutTokens         []string    `json:"outTokens"`
	InAmounts         []string    `json:"inAmounts"`
	OutAmounts        []string    `json:"outAmounts"`
	GasEstimate       float64     `json:"gasEstimate"`
	DataGasEstimate   int         `json:"dataGasEstimate"`
	GweiPerGas        float64     `json:"gweiPerGas"`
	GasEstimateValue  float64     `json:"gasEstimateValue"`
	InValues          []float64   `json:"inValues"`
	OutValues         []float64   `json:"outValues"`
	NetOutValue       float64     `json:"netOutValue"`
	PriceImpact       float64     `json:"priceImpact"`
	PercentDiff       float64     `json:"percentDiff"`
	PartnerFeePercent float64     `json:"partnerFeePercent"`
	PathId            string      `json:"pathId"`
	PathViz           interface{} `json:"pathViz"`
	BlockNumber       int         `json:"blockNumber"`
}

func (c *OdosMaker) Quote(ctx context.Context, in *defi.DefaultSwapReq) (*QuoteRes, error) {

	slippage, err := lib.StringToFloat(in.Slippage)
	if err != nil {
		return nil, err
	}

	fromToken, supported := c.TokenMap[in.FromToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(in.FromToken)
	}
	if in.FromToken == v1.Token_ETH {
		fromToken = defi.ZeroAddress
	}

	toToken, supported := c.TokenMap[in.ToToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(in.FromToken)
	}

	b := QuoteReq{
		ChainId: int(c.ChainId.Int64()),
		InputTokens: [1]struct {
			TokenAddress string `json:"tokenAddress"`
			Amount       string `json:"amount"`
		}{
			{
				TokenAddress: fromToken.String(),
				Amount:       in.Amount.String(),
			},
		},
		OutputTokens: [1]struct {
			TokenAddress string `json:"tokenAddress"`
			Proportion   int    `json:"proportion"`
		}{
			{
				TokenAddress: toToken.String(),
				Proportion:   1,
			},
		},
		UserAddr:             c.Addr.String(),
		SlippageLimitPercent: slippage,
		//ReferralCode:         2241664650,
		ReferralCode: 0,
		Compact:      true,
	}

	body, err := json.Marshal(&b)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.odos.xyz/sor/quote/v2", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.CliHttp.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		e, _ := io.ReadAll(res.Body)
		return nil, errors.New("invalid status core: " + strconv.Itoa(res.StatusCode) + " " + string(e))
	}

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var output QuoteRes
	if err := json.Unmarshal(response, &output); err != nil {
		return nil, err
	}

	return &output, nil
}

type AssembleReq struct {
	UserAddr string `json:"userAddr"`
	PathId   string `json:"pathId"`
	Simulate bool   `json:"simulate"`
}

func (c *OdosMaker) Assemble(ctx context.Context, in *defi.DefaultSwapReq, in2 *QuoteRes) (*AssembleRes, error) {

	b := AssembleReq{
		UserAddr: c.Addr.String(),
		PathId:   in2.PathId,
		Simulate: in.FromToken == v1.Token_ETH,
	}

	body, err := json.Marshal(&b)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.odos.xyz/sor/assemble", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.CliHttp.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		e, _ := io.ReadAll(res.Body)
		return nil, errors.New("invalid status core: " + strconv.Itoa(res.StatusCode) + " " + string(e))
	}

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var output AssembleRes
	if err := json.Unmarshal(response, &output); err != nil {
		return nil, err
	}

	return &output, nil
}

type AssembleRes struct {
	Deprecated       interface{} `json:"deprecated"`
	BlockNumber      int         `json:"blockNumber"`
	GasEstimate      int         `json:"gasEstimate"`
	GasEstimateValue float64     `json:"gasEstimateValue"`
	InputTokens      []struct {
		TokenAddress string `json:"tokenAddress"`
		Amount       string `json:"amount"`
	} `json:"inputTokens"`
	OutputTokens []struct {
		TokenAddress string `json:"tokenAddress"`
		Amount       string `json:"amount"`
	} `json:"outputTokens"`
	NetOutValue float64  `json:"netOutValue"`
	OutValues   []string `json:"outValues"`
	Transaction struct {
		Gas      int    `json:"gas"`
		GasPrice int    `json:"gasPrice"`
		Value    string `json:"value"`
		To       string `json:"to"`
		From     string `json:"from"`
		Data     string `json:"data"`
		Nonce    int    `json:"nonce"`
	} `json:"transaction"`
}
