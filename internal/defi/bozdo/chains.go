package bozdo

import (
	"math/big"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

var ChainMap = map[v1.Network]*big.Int{
	v1.Network_ZKSYNCERA:   big.NewInt(324),
	v1.Network_ARBITRUM:    big.NewInt(42161),
	v1.Network_BinanaceBNB: big.NewInt(56),
	v1.Network_Base:        big.NewInt(8453),
	v1.Network_AVALANCHE:   big.NewInt(43114),
	v1.Network_Etherium:    big.NewInt(1),
	v1.Network_OPTIMISM:    big.NewInt(10),
	v1.Network_POLIGON:     big.NewInt(137),
}
