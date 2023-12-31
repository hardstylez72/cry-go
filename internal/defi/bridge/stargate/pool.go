package stargate

import (
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

var (
	PoolIdMap = map[v1.Network]map[v1.Token]int64{
		v1.Network_ARBITRUM: {
			v1.Token_USDCBridged: 1,
			v1.Token_USDT:        2,
			//"FRAX":    7,
			v1.Token_ETH: 13,
			//"LUSD":    15,
			//"MAI":     16,
		},
		v1.Network_OPTIMISM: {
			v1.Token_USDCBridged: 1,
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
