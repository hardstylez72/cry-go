package poligon

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
	//MainNetURL = "https://rpc.ankr.com/polygon"
	MainNetURL = "https://polygon-rpc.com/"
)

// docs https://arbiscan.io/tokens?p=1
var TokenAddress = map[defi.Token]common.Address{
	v1.Token_USDT:  common.HexToAddress("0xc2132d05d31c914a87c6611c10748aeb04b58e8f"),
	v1.Token_STG:   common.HexToAddress("0x2F6F07CDcf3588944Bf4C42aC74ff24bF56e7590"),
	v1.Token_USDC:  common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"),
	v1.Token_veSTG: common.HexToAddress("0xb0d502e938ed5f4df2e681fe6e419ff29631d62b"),
	v1.Token_MATIC: common.HexToAddress("0x0000000000000000000000000000000000000000"),
}

var Dict = defi.Dict{
	Stargate: defi.Stargate{
		StargateRouterAddress:    common.HexToAddress("0x45A01E4e04F14f7A4a6702c74187c5F6222033cd"),
		StargateRouterEthAddress: common.HexToAddress(""),
	},
	Merkly: defi.Merkly{
		NFT: common.HexToAddress("0xa184998eC58dc1dA77a1F9f1e361541257A50CF4"),
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
	return "https://polygonscan.com/tx/" + txId
}

func NewClient(c *ClientConfig) (*Client, error) {

	config := &ClientConfig{
		HttpCli: &http.Client{},
	}
	if c != nil {
		config = c
	}

	ethcli, err := defi.NewEVMClient(&defi.Config{
		Network:   v1.Network_POLIGON,
		MainToken: v1.Token_MATIC,
		MainNet:   c.RPCEndpoint,
		TokenMap:  TokenAddress,
		Dict:      Dict,
		Httpcli:   config.HttpCli,
		TxViewFn:  TxViewer,
		NetworkId: bozdo.ChainMap[v1.Network_POLIGON],
		EstimateL1Gas: func(ctx context.Context, data []byte) (*big.Int, error) {
			return big.NewInt(0), nil
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ethereum main: "+c.RPCEndpoint)
	}
	ethcli.London = defi.LondonReadyBoosty

	return &Client{defi: ethcli, NetworkId: bozdo.ChainMap[v1.Network_POLIGON]}, nil
}
