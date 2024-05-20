package core

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
	MainNetURL = "https://rpc.ankr.com/core"
)

// docs https://arbiscan.io/tokens?p=1
var TokenAddress = map[defi.Token]common.Address{
	v1.Token_CORE: common.HexToAddress("0xfcdcab97c96ce3e690703e83729dfc853075ec89"),
	v1.Token_USDT: common.HexToAddress("0x900101d06A7426441Ae63e9AB3B9b0F63Be145F1"),
}

var Dict = defi.Dict{
	Stargate: defi.Stargate{
		StargateRouterAddress:    common.HexToAddress(""),
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
	return "https://scan.coredao.org/tx/" + txId
}

func NewClient(c *ClientConfig) (*Client, error) {

	config := &ClientConfig{
		HttpCli: &http.Client{},
	}
	if c != nil {
		config = c
	}

	ethcli, err := defi.NewEVMClient(&defi.Config{
		Network:   v1.Network_Core,
		MainToken: v1.Token_CORE,
		MainNet:   c.RPCEndpoint,
		TokenMap:  TokenAddress,
		Dict:      Dict,
		Httpcli:   config.HttpCli,
		TxViewFn:  TxViewer,
		NetworkId: bozdo.ChainMap[v1.Network_Core],
		EstimateL1Gas: func(ctx context.Context, data []byte) (*big.Int, error) {
			return big.NewInt(0), nil
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ethereum main: "+c.RPCEndpoint)
	}

	ethcli.London = defi.NeLondonReadyTx

	return &Client{defi: ethcli, NetworkId: bozdo.ChainMap[v1.Network_Core]}, nil
}
