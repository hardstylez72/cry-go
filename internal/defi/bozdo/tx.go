package bozdo

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/hardstylez72/cry/internal/exchange/pub"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type TxDetailKey = string

const (
	TxDetailKeyProtocolFee         = "ProtocolFee"
	TxDetailKeyLayerZeroFee        = "LayerZeroFee"
	TxDetailKeyOptimismL1Fee       = "OptimismL1Fee"
	TxDetailKeyTokenBalanceBefore  = "TokenBalanceBefore"
	TxDetailKeyTokenBalanceAfter   = "TokenBalanceAfter"
	TxDetailKeyNativeBalanceBefore = "NativeBalanceBefore"
	TxDetailKeyNativeBalanceAfter  = "NativeBalanceAfter"
	TxDetailKeyTxFee               = "TxFee"
)

type TxDetail struct {
	Key   TxDetailKey
	Value string
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

func NewLZFeeDetails(s *big.Int, network v1.Network, token v1.Token) TxDetail {
	return TxDetail{
		Key:   TxDetailKeyLayerZeroFee,
		Value: CastUSD(s, network, token) + " USD",
	}
}

func NewOpimismFeeDetails(s *big.Int, network v1.Network, token v1.Token) TxDetail {
	return TxDetail{
		Key:   TxDetailKeyOptimismL1Fee,
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

func CastUSD(wei *big.Int, network v1.Network, token v1.Token) string {
	gasTokenPrice := pub.Price().ETH
	switch network {
	case v1.Network_ZKSYNCERA, v1.Network_ZKSYNCLITE,
		v1.Network_Etherium, v1.Network_ARBITRUM,
		v1.Network_OPTIMISM, v1.Network_GOERLIETH:
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
