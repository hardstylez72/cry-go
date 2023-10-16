package arbitrum

import (
	"context"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

const (
	MainNetURL = "https://rpc.ankr.com/arbitrum"
)

// docs https://arbiscan.io/tokens?p=1
var TokenAddress = map[defi.Token]common.Address{
	v1.Token_USDT: common.HexToAddress("0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9"),
	v1.Token_STG:  common.HexToAddress("0x6694340fc020c5e6b96567843da2df01b2ce1eb6"),

	v1.Token_USDC:        common.HexToAddress("0xaf88d065e77c8cc2239327c5edb3a432268e5831"),
	v1.Token_USDCBridged: common.HexToAddress("0xff970a61a04b1ca14834a43f5de4533ebddb5cc8"),

	v1.Token_veSTG: common.HexToAddress("0xb0d502e938ed5f4df2e681fe6e419ff29631d62b"),
	v1.Token_ETH:   common.HexToAddress("0x0000000000000000000000000000000000000000"),
}

var Dict = defi.Dict{
	Stargate: defi.Stargate{
		StargateRouterAddress:    common.HexToAddress("0x53Bf833A5d6c4ddA888F69c22C88C9f356a41614"),
		StargateRouterEthAddress: common.HexToAddress("0xbf22f0f184bCcbeA268dF387a49fF5238dD23E40"),
	},
	TestNetBridgeSwapAddress: common.HexToAddress("0x0A9f824C05A74F577A536A8A0c673183a872Dff4"),
	Merkly: defi.Merkly{
		NFT: common.HexToAddress("0xAa58e77238f0E4A565343a89A79b4aDDD744d649"),
	},
	Across: defi.Across{
		RouterETH:   common.HexToAddress("0x269727F088F16E1Aea52Cf5a97B1CD41DAA3f02D"),
		RouterToken: common.HexToAddress("0xe35e9842fceaCA96570B734083f4a58e8F7C5f2A"),
	},
}

type Client struct {
	defi      *defi.EtheriumClient
	NetworkId *big.Int
}

type ClientConfig struct {
	HttpCli     *http.Client
	RPCEndpoint string
}

func TxViewer(txId string) string {
	return "https://arbiscan.io/tx/" + txId
}

func NewClient(c *ClientConfig) (*Client, error) {

	config := &ClientConfig{
		HttpCli: &http.Client{},
	}
	if c != nil {
		config = c
	}

	ethcli, err := defi.NewEVMClient(&defi.Config{
		Network:   v1.Network_ARBITRUM,
		MainToken: v1.Token_ETH,
		MainNet:   c.RPCEndpoint,
		TokenMap:  TokenAddress,
		Dict:      Dict,
		Httpcli:   config.HttpCli,
		TxViewFn:  TxViewer,
		NetworkId: bozdo.ChainMap[v1.Network_ARBITRUM],
		EstimateL1Gas: func(ctx context.Context, data []byte) (*big.Int, error) {
			return big.NewInt(0), nil
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ethereum main: "+c.RPCEndpoint)
	}

	return &Client{
		defi:      ethcli,
		NetworkId: bozdo.ChainMap[v1.Network_ARBITRUM],
	}, nil
}
