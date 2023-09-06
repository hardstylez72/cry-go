package zksyncera

import (
	"context"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/contracts/maverickrouter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type maverickFiMaker struct {
	*Client
}

// eth -> usdc
// 0xc04b8d59000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000004a6e7c137a6691d55693ca3bc7e5c698d9d438150000000000000000000000000000000000000000000000000000000064c11132000000000000000000000000000000000000000000000000001f33bb6f50acd30000000000000000000000000000000000000000000000000000000000f5d994000000000000000000000000000000000000000000000000000000000000003c5aea5775959fbc2557cc8789bc1bf90a239d9a9141c8cf74c27554a8972d3bf3d2bd4a14d8b604ab3355df6d4c9c3035724fd0e3914de96a5a83aaf400000000

// site https://explorer.zksync.io/tx/0x614d510f3ba8b394efbbb7712c5bc94e98239be5e79267ce6c45205ddddc8552 ETH -> USDC
// me https://explorer.zksync.io/tx/0x044fcd4eaf54ddd686589f216f506d9c310ed8af77d8139f2ef4bb5ef45643e5 USDC -> ETH
func (c *Client) MaverickSwap(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.DefaultRes, error) {
	return c.GenericSwap(ctx, &maverickFiMaker{c}, req)
}

func (c *maverickFiMaker) MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {

	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	w, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	fromTokenAddr, supported := c.Cfg.TokenMap[req.FromToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(req.FromToken)
	}

	toTokenAddr, supported := c.Cfg.TokenMap[req.ToToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(req.ToToken)
	}

	am := defi.WeiToToken(req.Amount, req.FromToken)

	d, err := getMaverickSwapData(ctx, fromTokenAddr, toTokenAddr, am.String(), "0.1")
	if err != nil {
		return nil, errors.Wrap(err, "getMaverickSwapData")
	}

	amOut := defi.TokenAmountFloatToWEI(d.OutputAmount, req.ToToken)

	amSlip, err := defi.Slippage(amOut, req.Slippage)
	if err != nil {
		return nil, err
	}
	amOut = amSlip

	pathb := []byte{}
	pathb = append(pathb, common.HexToAddress(d.Path[0]).Bytes()...)
	pathb = append(pathb, common.HexToAddress(d.Path[1]).Bytes()...)
	pathb = append(pathb, common.HexToAddress(d.Path[2]).Bytes()...)
	//path := common.BytesToHash(pathb)

	// (USDC -> ETH) recipient = ZERO
	// unwrapWETH9 + am 0, recipient - wallet
	// ETH -> LUSD refundETH
	recipient := w.WalletAddr
	if req.ToToken == v1.Token_ETH {
		recipient = ZEROADDR
	}

	params := maverickrouter.IRouterExactInputParams{
		Path:             pathb,
		Recipient:        recipient,
		Deadline:         DefaultDeadLine(),
		AmountIn:         req.Amount,
		AmountOutMinimum: amOut,
	}

	constractabi, err := maverickrouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	p1, err := constractabi.Pack("exactInput", params)
	if err != nil {
		return nil, err
	}

	var p2 []byte
	if req.ToToken == v1.Token_ETH {
		p2, err = constractabi.Pack("unwrapWETH9", big.NewInt(0), w.WalletAddr)
		if err != nil {
			return nil, err
		}
	}

	if req.FromToken == v1.Token_ETH {
		p2, err = constractabi.Pack("refundETH")
		if err != nil {
			return nil, err
		}
	}

	if len(p2) == 0 {

		data, err := constractabi.Pack("multicall", [][]byte{p1})
		if err != nil {
			return nil, err
		}
		return &bozdo.TxData{
			Data:         data,
			Value:        value,
			ContractAddr: c.Cfg.Maverick.Router,
			Rate:         CalcRate(req.FromToken, req.ToToken, req.Amount, amOut),
		}, nil
	}

	data, err := constractabi.Pack("multicall", [][]byte{p1, p2})
	if err != nil {
		return nil, err
	}
	return &bozdo.TxData{
		Data:         data,
		Value:        value,
		ContractAddr: c.Cfg.Maverick.Router,
		Rate:         CalcRate(req.FromToken, req.ToToken, req.Amount, amOut),
	}, nil
}

func getMaverickSwapData(ctx context.Context, fromToken, toToken common.Address, amountIn string, slippage string) (*getMaverickSwapDataRes, error) {
	cli := &http.Client{}
	zksyncChainId := "324"
	base := strings.Join([]string{
		"https://api.mav.xyz/api/swapRoutes?chainId=", zksyncChainId,
		"&inTokenAddress=", fromToken.String(),
		"&outTokenAddress=", toToken.String(),
		"&amount=", amountIn,
		"&gasPrice=0.25&slippage=", slippage,
	}, "")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base, nil)
	if err != nil {
		return nil, err
	}

	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		b, _ := io.ReadAll(res.Body)
		return nil, errors.New(string(b))
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result getMaverickSwapDataRes
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	if result.Message != "Path successfully found." {
		return nil, errors.New(result.Message)
	}

	return &result, nil
}

type getMaverickSwapDataRes struct {
	Path         [3]string `json:"path"`
	OutputAmount float64   `json:"outputAmount"`
	Message      string    `json:"message"`
}
