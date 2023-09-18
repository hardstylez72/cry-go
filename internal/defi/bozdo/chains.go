package bozdo

import (
	"math/big"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

var ChainMap = map[v1.Network]*big.Int{
	v1.Network_ZKSYNCERA: big.NewInt(324),
	v1.Network_ARBITRUM:  big.NewInt(42161),
}
