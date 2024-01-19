package kyberswap

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/hardstylez72/cry/internal/defi"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type QuoteReq struct {
	tokenIn    string
	tokenOut   string
	amountIn   string
	saveGas    bool
	gasInclude bool
}

type QuoteRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		RouteSummary  RouteSummary `json:"routeSummary"`
		RouterAddress string       `json:"routerAddress"`
	} `json:"data"`
	RequestId string `json:"requestId"`
}

func (c *KyberSwapMaker) Quote(ctx context.Context, in *defi.DefaultSwapReq) (*QuoteRes, error) {

	fromToken, supported := c.TokenMap[in.FromToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(in.FromToken)
	}
	if in.FromToken == v1.Token_ETH {
		fromToken = defi.NativeTokenAddress
	}

	toToken, supported := c.TokenMap[in.ToToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(in.FromToken)
	}
	if in.ToToken == v1.Token_ETH {
		toToken = defi.NativeTokenAddress
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.makeBaseURL(in.Network)+"/api/v1/routes", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Set("tokenIn", fromToken.String())
	q.Set("tokenOut", toToken.String())
	q.Set("amountIn", in.Amount.String())
	q.Set("saveGas", "false")
	q.Set("gasInclude", "true")
	req.URL.RawQuery = q.Encode()

	res, err := c.CliHttp.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		e, _ := io.ReadAll(res.Body)
		return nil, errors.New("invalid status code: " + strconv.Itoa(res.StatusCode) + " " + string(e))
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
