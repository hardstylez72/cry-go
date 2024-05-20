package fantom

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
	MainNetURL = "https://rpc.ankr.com/fantom"
)

// docs https://arbiscan.io/tokens?p=1
var TokenAddress = map[defi.Token]common.Address{
	v1.Token_STG: common.HexToAddress("0x2F6F07CDcf3588944Bf4C42aC74ff24bF56e7590"),
	v1.Token_FTM: common.HexToAddress("0x0000000000000000000000000000000000000000"),
}

var Dict = defi.Dict{
	Stargate: defi.Stargate{
		StargateRouterAddress:    common.HexToAddress("0xAf5191B0De278C7286d6C7CC6ab6BB8A73bA2Cd6"),
		StargateRouterEthAddress: common.HexToAddress(""),
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
	return "https://ftmscan.com/tx/" + txId
}

func NewClient(c *ClientConfig) (*Client, error) {

	config := &ClientConfig{
		HttpCli: &http.Client{},
	}
	if c != nil {
		config = c
	}

	ethcli, err := defi.NewEVMClient(&defi.Config{
		Network:   v1.Network_Fantom,
		MainToken: v1.Token_FTM,
		MainNet:   c.RPCEndpoint,
		TokenMap:  TokenAddress,
		Dict:      Dict,
		Httpcli:   config.HttpCli,
		TxViewFn:  TxViewer,
		NetworkId: bozdo.ChainMap[v1.Network_Fantom],
		EstimateL1Gas: func(ctx context.Context, data []byte) (*big.Int, error) {
			return big.NewInt(0), nil
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ethereum main: "+c.RPCEndpoint)
	}
	ethcli.London = defi.LondonNotReady

	return &Client{defi: ethcli, NetworkId: bozdo.ChainMap[v1.Network_Fantom]}, nil
}
