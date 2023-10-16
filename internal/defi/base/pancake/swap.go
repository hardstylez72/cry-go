package pancake

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/log"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Swapper struct {
	*defi.EtheriumClient
}

func (c *Swapper) MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {

	fromToken, ok := c.Cfg.TokenMap[req.FromToken]
	if !ok {
		return nil, defi.ErrTokenNotSupportedFn(req.FromToken)
	}

	toToken, ok := c.Cfg.TokenMap[req.ToToken]
	if !ok {
		return nil, defi.ErrTokenNotSupportedFn(req.ToToken)
	}

	_, err := c.Pool(ctx, fromToken, toToken)
	if err != nil {
		return nil, errors.Wrap(err, "pool failed")
	}

	if req.FromToken == v1.Token_ETH {
		return c.SwapToToken(ctx, req)
	}

	if req.ToToken == v1.Token_ETH {
		return c.SwapToETH(ctx, req)
	}

	return nil, errors.New("unsupported operation")
}

func (c *Swapper) Pool(ctx context.Context, from, to common.Address) (*common.Address, error) {
	fc, err := NewFactoryCaller(c.Cfg.Dict.Pancake.Factory, c.Cli)
	if err != nil {
		return nil, err
	}

	pool, err := fc.GetPool(&bind.CallOpts{Context: ctx}, from, to, big.NewInt(500))
	if err != nil {
		return nil, err
	}

	if pool.String() == defi.ZeroAddress.String() {
		return nil, errors.New("pool not found")
	}

	return &pool, nil
}

func (c *Swapper) GetAmountOut(ctx context.Context, fromToken, toToken common.Address, amount *big.Int) (*big.Int, error) {
	caller, err := NewQuoterCaller(c.Cfg.Dict.Pancake.Quoter, c.Cli)
	if err != nil {
		return nil, err
	}

	res, err := caller.QuoteExactInputSingle(&bind.CallOpts{Context: ctx}, IQuoterV2QuoteExactInputSingleParams{
		TokenIn:           fromToken,
		TokenOut:          toToken,
		AmountIn:          amount,
		Fee:               big.NewInt(500),
		SqrtPriceLimitX96: big.NewInt(0),
	})
	if err != nil {
		return nil, err
	}
	return res.AmountOut, nil
}
func (c *Swapper) SwapToETH(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {
	fromToken, ok := c.Cfg.TokenMap[req.FromToken]
	if !ok {
		return nil, defi.ErrTokenNotSupportedFn(req.FromToken)
	}

	toToken, ok := c.Cfg.TokenMap[req.ToToken]
	if !ok {
		return nil, defi.ErrTokenNotSupportedFn(req.ToToken)
	}

	value := big.NewInt(0)

	wt, err := defi.NewWalletTransactor(req.WalletPK)
	if err != nil {
		return nil, err
	}

	amOut, err := c.GetAmountOut(ctx, fromToken, toToken, req.Amount)
	if err != nil {
		return nil, err
	}

	amOut, err = defi.Slippage(amOut, req.Slippage)
	if err != nil {
		return nil, err
	}

	abi, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	swap, err := abi.Pack("exactInputSingle", IV3SwapRouterExactInputSingleParams{
		TokenIn:           fromToken,
		TokenOut:          toToken,
		Fee:               big.NewInt(500),
		Recipient:         common.HexToAddress("0x0000000000000000000000000000000000000002"),
		AmountIn:          req.Amount,
		AmountOutMinimum:  amOut,
		SqrtPriceLimitX96: big.NewInt(0),
	})
	if err != nil {
		return nil, err
	}

	unwrap, err := abi.Pack("unwrapWETH9", amOut, wt.WalletAddr)
	if err != nil {
		return nil, err
	}

	d := defi.DeadLine()

	multicall, err := abi.Pack("multicall0", d, [][]byte{swap, unwrap})
	if err != nil {
		return nil, err
	}

	if req.Debug {
		log.Log.Debug(zap.String("data", "0x"+common.Bytes2Hex(multicall)))
	}

	return &bozdo.TxData{
		Data:         multicall,
		Value:        value,
		ContractAddr: c.Cfg.Dict.Pancake.Router,
		Details:      nil,
		Rate:         defi.CalcRate(req.FromToken, req.ToToken, req.Amount, amOut),
		Gas:          nil,
	}, err

}
func (c *Swapper) SwapToToken(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {

	fromToken, ok := c.Cfg.TokenMap[req.FromToken]
	if !ok {
		return nil, defi.ErrTokenNotSupportedFn(req.FromToken)
	}

	toToken, ok := c.Cfg.TokenMap[req.ToToken]
	if !ok {
		return nil, defi.ErrTokenNotSupportedFn(req.ToToken)
	}

	value := req.Amount

	wt, err := defi.NewWalletTransactor(req.WalletPK)
	if err != nil {
		return nil, err
	}

	amOut, err := c.GetAmountOut(ctx, fromToken, toToken, req.Amount)
	if err != nil {
		return nil, err
	}

	abi, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	swap, err := abi.Pack("exactInputSingle", IV3SwapRouterExactInputSingleParams{
		TokenIn:           fromToken,
		TokenOut:          toToken,
		Fee:               big.NewInt(500),
		Recipient:         wt.WalletAddr,
		AmountIn:          req.Amount,
		AmountOutMinimum:  amOut,
		SqrtPriceLimitX96: big.NewInt(0),
	})
	if err != nil {
		return nil, err
	}

	d := defi.DeadLine()

	multicall, err := abi.Pack("multicall0", d, [][]byte{swap})
	if err != nil {
		return nil, err
	}

	if req.Debug {
		log.Log.Debug(zap.String("data", "0x"+common.Bytes2Hex(multicall)))
	}

	return &bozdo.TxData{
		Data:         multicall,
		Value:        value,
		ContractAddr: c.Cfg.Dict.Pancake.Router,
		Details:      nil,
		Rate:         defi.CalcRate(req.FromToken, req.ToToken, req.Amount, amOut),
		Gas:          nil,
	}, err
}
