package defi

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/bridge/layerzero"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

var (
	PoolIdMap = map[v1.Network]map[Token]int64{
		v1.Network_ARBITRUM: {
			v1.Token_USDC: 1,
			v1.Token_USDT: 2,
			//"FRAX":    7,
			v1.Token_ETH: 13,
			//"LUSD":    15,
			//"MAI":     16,
		},
		v1.Network_OPTIMISM: {
			v1.Token_USDC: 1,
			//"DAI":     3,
			//"FRAX":    7,
			v1.Token_ETH: 13,
			//"sUSD":    14,
			//"LUSD":    15,
			//"MAI":     16,
		},
		v1.Network_BinanaceBNB: {
			v1.Token_USDT: 2,
			//"BUSD":    5,
			//"USDD":    11,
			//"MAI":     16,
			//"METIS":   17,
		},
		v1.Network_Etherium: {
			v1.Token_USDC: 1,
			v1.Token_USDT: 2,
			v1.Token_ETH:  13,
			//DAI: 3
			//FRAX: 7
			//USDD: 11
			//sUSD: 14
			//LUSD: 15
			//MAI: 16
			//METIS: 17
			//metis.USDT: 19
		},
		v1.Network_POLIGON: {
			v1.Token_USDC: 1,
			v1.Token_USDT: 2,
			//DAI: 3
			//MAI: 16
		},
		v1.Network_AVALANCHE: {
			v1.Token_USDC: 1,
			v1.Token_USDT: 2,
			//FRAX: 7
			//MAI: 16
			//metis.USDT: 19
		},
	}
)

type (
	StargateBridgeChainName string
)

type GasLimit struct {
	TotalGas *big.Int
}

type StargateBridgeSwapReq struct {
	DestChain    v1.Network
	WalletPk     string
	Quantity     *big.Int
	FromToken    Token
	ToToken      Token
	Gas          *bozdo.Gas
	EstimateOnly bool
	Slippage     SlippagePercent
}

func (r *StargateBridgeSwapReq) Validate(currentChain v1.Network) error {
	if r == nil {
		return errors.New("empty input")
	}

	if r.Quantity.Cmp(big.NewInt(0)) == 0 {
		return errors.New("quantity 0")
	}

	if r.WalletPk == "" {
		return errors.New("empty wallet")
	}

	if r.DestChain == currentChain {
		return errors.New("invalid chain, same chain")
	}

	_, ok := layerzero.LayerZeroChainMap[r.DestChain]
	if !ok {
		return errors.New("invalid chain: " + string(r.DestChain))
	}

	_, ok = PoolIdMap[r.DestChain][r.ToToken]
	if !ok {
		return errors.New("invalid dest chain token: " + string(r.ToToken))
	}

	_, ok = PoolIdMap[currentChain][r.FromToken]
	if !ok {
		return errors.New("invalid current chain token: " + string(r.FromToken))
	}

	return nil
}

// https://stargateprotocol.gitbook.io/stargate/developers/how-to-swap
func (c *EtheriumClient) StargateBridgeSwap(ctx context.Context, req *DefaultBridgeReq) (*bozdo.DefaultRes, error) {

	wallet, err := newWalletTransactor(req.PK)
	if err != nil {
		return nil, err
	}

	r := &bozdo.DefaultRes{}

	switch req.FromToken {
	case v1.Token_ETH:
		res, err := c.StargateBridgeSwapEth(ctx, &StargateBridgeSwapEthReq{
			DestChain:    req.ToNetwork,
			Wallet:       wallet,
			Quantity:     req.Amount,
			Gas:          req.Gas,
			EstimateOnly: req.EstimateOnly,
			Slippage:     req.Slippage,
		})
		if err != nil {
			return nil, errors.Wrap(err, "stargateBridgeSwapEth")
		}
		if res.Tx != nil {
			r.Tx = c.NewTx(res.Tx.Hash(), CodeContract, res.ECost.Details)
		}
		r.ECost = res.ECost
	case v1.Token_MAV:
		return c.StargateBridgeSwapMAV(ctx, req)
	case v1.Token_STG:
		limitTx, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
			Token:       req.FromToken,
			Wallet:      wallet,
			Amount:      req.Amount,
			SpenderAddr: c.Cfg.TokenMap[v1.Token_STG],
		})
		if err != nil {
			return nil, errors.Wrap(err, "TokenLimitChecker")
		}
		if limitTx.LimitExtended {
			r.ApproveTx = c.NewTx(limitTx.ApproveTx.Hash(), CodeApprove, nil)
		}

		res, err := c.StargateBridgeSwapSTG(ctx, &StargateBridgeSwapSTGReq{
			DestChain:    req.ToNetwork,
			Wallet:       wallet,
			Quantity:     req.Amount,
			Gas:          req.Gas,
			EstimateOnly: req.EstimateOnly,
		})
		if err != nil {
			return nil, errors.Wrap(err, "stargateBridgeSwapSTG")
		}

		if res.Tx != nil {
			r.Tx = c.NewTx(res.Tx.Hash(), CodeContract, res.ECost.Details)
		}
		r.ECost = res.ECost
	default:
		limitTx, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
			Token:       req.FromToken,
			Wallet:      wallet,
			Amount:      req.Amount,
			SpenderAddr: c.Cfg.Dict.Stargate.StargateRouterAddress,
		})
		if err != nil {
			return nil, errors.Wrap(err, "TokenLimitChecker")
		}
		if limitTx.LimitExtended {
			r.ApproveTx = c.NewTx(limitTx.ApproveTx.Hash(), CodeApprove, nil)
		}

		rr := &StargateBridgeSwapTokenReq{
			DestChain:    req.ToNetwork,
			WalletPk:     req.PK,
			Quantity:     req.Amount,
			FromToken:    req.FromToken,
			ToToken:      req.ToToken,
			Gas:          req.Gas,
			EstimateOnly: req.EstimateOnly,
			Slippage:     req.Slippage,
		}
		res, err := c.StargateBridgeSwapToken(ctx, rr)
		if err != nil {
			return nil, errors.Wrap(err, "stargateBridgeSwap")
		}
		if res.Tx != nil {
			r.Tx = c.NewTx(res.Tx.Hash(), CodeContract, res.ECost.Details)
		}
		r.ECost = res.ECost
	}

	return r, nil
}
