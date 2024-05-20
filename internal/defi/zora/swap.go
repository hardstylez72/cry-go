package zora

import (
	"math/big"
)

func Gwei(wei *big.Int) *big.Float {
	return new(big.Float).Quo(big.NewFloat(0).SetInt(wei), big.NewFloat(1e9))
}
