package bozdo

import (
	"math"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/hardstylez72/cry/internal/exchange/pub"
	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type TxDetailKey = string

const (
	TxDetailKeyProtocolFee  = "ProtocolFee"
	TxDetailKeyL1Fee        = "l1Fee"
	TxDetailKeyLayerZeroFee = "LayerZeroFee"

	TxDetailKeyTokenBalanceBefore  = "TokenBalanceBefore"
	TxDetailKeyTokenBalanceAfter   = "TokenBalanceAfter"
	TxDetailKeyNativeBalanceBefore = "NativeBalanceBefore"
	TxDetailKeyNativeBalanceAfter  = "NativeBalanceAfter"

	TxDetailKeyTxFee           = "TxFee"
	TxDetailEaxchnageRateRatio = "RateRatio"
	TxDetailAcrossRelayFee     = "AcrossRelayFee"
)

type TxDetail struct {
	Key   TxDetailKey `json:"key"`
	Value string      `json:"value"`
}

type TxCode = string

const (
	CodeApprove  TxCode = "approve"
	CodeContract TxCode = "contract"
	CodeTransfer TxCode = "transfer"
	CodeSwap     TxCode = "swap"
	CodeBridge   TxCode = "bridge"
	CodeDmail    TxCode = "dmail"
	CodeMint     TxCode = "mint"
	CodeLP       TxCode = "lp"
)

type Transaction struct {
	Code    TxCode
	Network v1.Network
	Id      string
	Url     string
	Details []TxDetail
}

type DefaultRes struct {
	Tx        *Transaction
	ApproveTx *Transaction
	ECost     *EstimatedGasCost
	TxDetails []TxDetail
}

type TxData struct {
	Data         []byte
	Value        *big.Int
	ContractAddr common.Address
	Details      []TxDetail
	Rate         *float64
	Gas          *big.Int
	ExtraFee     *big.Int
	Code         TxCode
}

type BaseReq struct {
	EstimateOnly bool
	Gas          *Gas
	Debug        bool
}

func (g *Gas) RuleSet() bool {
	return g != nil
}

type TxType = string

const TxTypeLegacy = "legacy"
const TxTypeDynamic = "legacy"
const TxTypeStarkNet = "starkNet"
const TxTypeOptimism = "optimism"

type EstimatedGasCost struct {
	Type        TxType
	Name        string
	GasLimit    *big.Int
	GasPrice    *big.Int
	TotalGasWei *big.Int
	Details     []TxDetail
	ExtraFee    *big.Int
}

type Gas struct {
	Network       v1.Network
	GasMultiplier float64
	MaxGas        big.Int

	GasBeforeMultiplier big.Int
	GasLimit            big.Int
	GasPrice            big.Int
	TotalGas            big.Int
}

func NewKyperSwap(in, out string, priceImpact int) []TxDetail {
	return []TxDetail{
		{
			Key:   "KyperSwapIn",
			Value: in + " USD",
		},
		{
			Key:   "KyperSwapOut",
			Value: out + " USD",
		},
		{
			Key:   "KyperSwapPriceImpact",
			Value: strconv.Itoa(priceImpact) + " %",
		},
	}
}

func NewOdos(in, out string, network v1.Network, tokenIn, tokenOut v1.Token, priceImpact float64) []TxDetail {
	return []TxDetail{
		{
			Key:   "odosIn",
			Value: in + " " + tokenIn.String(),
		},
		{
			Key:   "odosOut",
			Value: out + " " + tokenOut.String(),
		},
		{
			Key:   "odosPriceImpact",
			Value: lib.FloatToString(priceImpact) + " %",
		},
	}
}

func NewTxFee(s *big.Int, network v1.Network, token v1.Token) TxDetail {
	return TxDetail{
		Key:   TxDetailKeyTxFee,
		Value: CastUSD(s, network, token) + " USD",
	}
}

func NewProtocolFeeDetails(s *big.Int, network v1.Network, token v1.Token) TxDetail {
	return TxDetail{
		Key:   TxDetailKeyProtocolFee,
		Value: CastUSD(s, network, token) + " USD",
	}
}

func NewL1FeeDetails(s *big.Int, network v1.Network, token v1.Token) TxDetail {
	return TxDetail{
		Key:   TxDetailKeyL1Fee,
		Value: CastUSD(s, network, token) + " USD",
	}
}

func NewAcrossDetails(relayFee string, token v1.Token, network v1.Network) TxDetail {

	b, ok := big.NewInt(0).SetString(relayFee, 10)
	if !ok {
		return TxDetail{
			Key:   TxDetailAcrossRelayFee,
			Value: "- USD",
		}
	}

	return TxDetail{
		Key:   TxDetailAcrossRelayFee,
		Value: CastUSD(b, network, token) + " USD",
	}
}

func NewLZFeeDetails(s *big.Int, network v1.Network, token v1.Token) TxDetail {
	return TxDetail{
		Key:   TxDetailKeyLayerZeroFee,
		Value: CastUSD(s, network, token) + " USD",
	}
}

func NewNativeBalanceBefore(s *big.Int, network v1.Network, token v1.Token) TxDetail {
	return TxDetail{
		Key:   TxDetailKeyNativeBalanceBefore,
		Value: CastUSD(s, network, token) + " USD",
	}
}

func NewNativeBalanceAfter(s *big.Int, network v1.Network, token v1.Token) TxDetail {
	return TxDetail{
		Key:   TxDetailKeyNativeBalanceAfter,
		Value: CastUSD(s, network, token) + " USD",
	}
}

func NewTokenBalanceAfter(s *big.Int, network v1.Network, token v1.Token) TxDetail {
	return TxDetail{
		Key:   TxDetailKeyTokenBalanceAfter,
		Value: CastUSD(s, network, token) + " USD",
	}
}

func NewTokenBalanceBefore(s *big.Int, network v1.Network, token v1.Token) TxDetail {
	return TxDetail{
		Key:   TxDetailKeyTokenBalanceBefore,
		Value: CastUSD(s, network, token) + " USD",
	}
}

func NewSwapRateRatio(exchange, pool float64) TxDetail {
	ratio := math.Abs(1-exchange/pool) * 100

	return TxDetail{
		Key:   TxDetailEaxchnageRateRatio,
		Value: lib.FloatToString(ratio) + " %",
	}
}

func CastUSD(wei *big.Int, network v1.Network, token v1.Token) string {
	var gasTokenPrice float64

	if network == v1.Network_BinanaceBNB {
		gasTokenPrice = pub.Price().BNB
		amUsd := EthToUsd(WEIToEther(wei), gasTokenPrice)
		return amUsd.String()
	}

	if network == v1.Network_Core {
		gasTokenPrice = pub.Price().CORE
		amUsd := EthToUsd(WEIToEther(wei), gasTokenPrice)
		return amUsd.String()
	}

	if network == v1.Network_POLIGON {
		gasTokenPrice = pub.Price().MATIC
		amUsd := EthToUsd(WEIToEther(wei), gasTokenPrice)
		return amUsd.String()
	}

	switch token {
	case v1.Token_USDC, v1.Token_USDT, v1.Token_USDCBridged:
		gasTokenPrice = 1
		amEth := WeiToToken(wei, token)
		return amEth.String()
	case v1.Token_ETH:
		switch network {
		case v1.Network_ZKSYNCERA, v1.Network_ZKSYNCLITE,
			v1.Network_Etherium, v1.Network_ARBITRUM,
			v1.Network_OPTIMISM, v1.Network_GOERLIETH, v1.Network_Base:
			gasTokenPrice = pub.Price().ETH
		case v1.Network_POLIGON:
			gasTokenPrice = pub.Price().MATIC
		case v1.Network_BinanaceBNB:
			gasTokenPrice = pub.Price().BNB
		case v1.Network_AVALANCHE:
			gasTokenPrice = pub.Price().AVAX
		}
		amEth := WEIToEther(wei)
		amUsd := EthToUsd(amEth, gasTokenPrice)
		return amUsd.String()
	}

	return "-1"
}

func WeiToToken(wei *big.Int, token v1.Token) *big.Float {
	switch token {
	case v1.Token_USDT, v1.Token_USDC, v1.Token_USDCBridged:
		return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether*1e-12))
	default:
		return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
	}
}

func WEIToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
}
func EthToUsd(e *big.Float, price float64) *big.Float {
	return new(big.Float).Mul(e, new(big.Float).SetFloat64(price))
}

func Percent(value *big.Int, percent int) *big.Int {
	f := math.Round(float64(percent))
	b1p := new(big.Int).Div(value, new(big.Int).SetInt64(100))
	return new(big.Int).Mul(b1p, new(big.Int).SetInt64(int64(f)))
}

func BigIntSum(values ...*big.Int) *big.Int {

	result := big.NewInt(0)

	for _, v := range values {
		if v == nil {
			continue
		}
		result = new(big.Int).Add(v, result)
	}

	return result

}
