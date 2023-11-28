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
	v1.Network_Linea:       big.NewInt(59144),
}

var NativeTokenMap = map[v1.Network]v1.Token{
	v1.Network_ZKSYNCERA:   v1.Token_ETH,
	v1.Network_ARBITRUM:    v1.Token_ETH,
	v1.Network_BinanaceBNB: v1.Token_BNB,
	v1.Network_Base:        v1.Token_ETH,
	v1.Network_AVALANCHE:   v1.Token_AVAX,
	v1.Network_Etherium:    v1.Token_ETH,
	v1.Network_OPTIMISM:    v1.Token_ETH,
	v1.Network_POLIGON:     v1.Token_MATIC,
	v1.Network_StarkNet:    v1.Token_ETH,
	v1.Network_Linea:       v1.Token_ETH,
}
