package kyberswap

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/lib"
	"github.com/pkg/errors"
)

type RouteSummary struct {
	TokenIn                      string `json:"tokenIn"`
	AmountIn                     string `json:"amountIn"`
	AmountInUsd                  string `json:"amountInUsd"`
	TokenInMarketPriceAvailable  bool   `json:"tokenInMarketPriceAvailable"`
	TokenOut                     string `json:"tokenOut"`
	AmountOut                    string `json:"amountOut"`
	AmountOutUsd                 string `json:"amountOutUsd"`
	TokenOutMarketPriceAvailable bool   `json:"tokenOutMarketPriceAvailable"`
	Gas                          string `json:"gas"`
	GasPrice                     string `json:"gasPrice"`
	GasUsd                       string `json:"gasUsd"`
	ExtraFee                     struct {
		FeeAmount   string `json:"feeAmount"`
		ChargeFeeBy string `json:"chargeFeeBy"`
		IsInBps     bool   `json:"isInBps"`
		FeeReceiver string `json:"feeReceiver"`
	} `json:"extraFee"`
	Route [][]struct {
		Pool              string `json:"pool"`
		TokenIn           string `json:"tokenIn"`
		TokenOut          string `json:"tokenOut"`
		LimitReturnAmount string `json:"limitReturnAmount"`
		SwapAmount        string `json:"swapAmount"`
		AmountOut         string `json:"amountOut"`
		Exchange          string `json:"exchange"`
		PoolLength        int    `json:"poolLength"`
		PoolType          string `json:"poolType"`
		PoolExtra         struct {
			LimitPoint int `json:"limitPoint"`
		} `json:"poolExtra"`
		Extra struct {
		} `json:"extra"`
	} `json:"route"`
}

type BuildReq struct {
	RouteSummary      RouteSummary `json:"routeSummary"`
	Deadline          int          `json:"deadline"`
	SlippageTolerance int          `json:"slippageTolerance"`
	Sender            string       `json:"sender"`
	Recipient         string       `json:"recipient"`
	Source            string       `json:"source"`
	SkipSimulateTx    bool         `json:"skipSimulateTx"`
}

type BuildRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		AmountIn     string `json:"amountIn"`
		AmountInUsd  string `json:"amountInUsd"`
		AmountOut    string `json:"amountOut"`
		AmountOutUsd string `json:"amountOutUsd"`
		Gas          string `json:"gas"`
		GasUsd       string `json:"gasUsd"`
		OutputChange struct {
			Amount  string `json:"amount"`
			Percent int    `json:"percent"`
			Level   int    `json:"level"`
		} `json:"outputChange"`
		Data          string `json:"data"`
		RouterAddress string `json:"routerAddress"`
	} `json:"data"`
	RequestId string `json:"requestId"`
}

func (c *KyberSwapMaker) Build(ctx context.Context, in *defi.DefaultSwapReq, in2 *QuoteRes) (*BuildRes, error) {

	w, err := defi.NewWalletTransactor(in.WalletPK)
	if err != nil {
		return nil, err
	}

	f, err := lib.StringToFloat(in.Slippage)
	if err != nil {
		return nil, errors.Wrap(err, "slippage")
	}
	tolerance := int(f * 10)

	b := BuildReq{
		RouteSummary:      in2.Data.RouteSummary,
		Deadline:          int(defi.DeadLine().Int64()),
		SlippageTolerance: tolerance,
		Sender:            w.WalletAddrHR,
		Recipient:         w.WalletAddrHR,
		Source:            in2.Data.RouteSummary.Route[0][0].Exchange,
		SkipSimulateTx:    false,
	}

	body, err := json.Marshal(&b)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.makeBaseURL(in.Network)+"/api/v1/route/build", bytes.NewBuffer(body))
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

	var output BuildRes
	if err := json.Unmarshal(response, &output); err != nil {
		return nil, err
	}

	return &output, nil
}
