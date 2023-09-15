package zksyncera

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

type odosMaker struct {
	CA      common.Address
	Cli     *Client
	CliHttp *http.Client
}

func (c *Client) OdosSwap(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.DefaultRes, error) {
	return c.GenericSwap(ctx, &odosMaker{
		CA:      common.HexToAddress("0x4bba932e9792a2b917d47830c93a9bc79320e4f7"),
		Cli:     c,
		CliHttp: &http.Client{},
	}, req)
}

func (c *odosMaker) MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {

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

	return &bozdo.TxData{
		Data:         data,
		Value:        v,
		ContractAddr: c.CA,
		Details:      bozdo.NewOdos(inString, outString, v1.Network_ZKSYNCERA, req.FromToken, req.ToToken, quote.PercentDiff),
		Rate:         CalcRate(req.FromToken, req.ToToken, req.Amount, out),
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

func (c *odosMaker) Quote(ctx context.Context, in *defi.DefaultSwapReq) (*QuoteRes, error) {

	wt, err := NewWalletTransactor(in.WalletPK, big.NewInt(ChainId))
	if err != nil {
		return nil, err
	}
	slippage, err := lib.StringToFloat(in.Slippage)
	if err != nil {
		return nil, err
	}

	fromToken, supported := c.Cli.Cfg.TokenMap[in.FromToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(in.FromToken)
	}
	toToken, supported := c.Cli.Cfg.TokenMap[in.ToToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(in.FromToken)
	}

	b := QuoteReq{
		ChainId: ChainId,
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
		UserAddr:             wt.WalletAddrHR,
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

func (c *odosMaker) Assemble(ctx context.Context, in *defi.DefaultSwapReq, in2 *QuoteRes) (*AssembleRes, error) {

	wt, err := NewWalletTransactor(in.WalletPK, big.NewInt(ChainId))
	if err != nil {
		return nil, err
	}

	b := AssembleReq{
		UserAddr: wt.WalletAddrHR,
		PathId:   in2.PathId,
		Simulate: false,
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
