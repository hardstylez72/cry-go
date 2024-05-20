package bnb

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

	//MainNetURL https://docs.polygonscan.com/getting-started/endpoint-urls
	//MainNetURL = "https://bsc-dataseed.binance.org"
	MainNetURL = "https://rpc.ankr.com/bsc"
)

var TokenAddress = map[defi.Token]common.Address{
	v1.Token_USDT: common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
	v1.Token_STG:  common.HexToAddress("0xB0D502E938ed5f4df2E681fE6E419ff29631d62b"),
	v1.Token_USDC: common.HexToAddress("0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"),
	v1.Token_BNB:  common.HexToAddress("0x0000000000000000000000000000000000000000"),
	v1.Token_MAV:  common.HexToAddress("0xd691d9a68C887BDF34DA8c36f63487333ACfD103"),
	v1.Token_ZK:   common.HexToAddress("0xc71b5f631354be6853efe9c3ab6b9590f8302e81"),
}

var Dict = defi.Dict{
	Stargate: defi.Stargate{
		StargateRouterAddress:    common.HexToAddress("0x4a364f8c717cAAD9A442737Eb7b8A55cc6cf18D8"),
		StargateRouterEthAddress: common.HexToAddress("not supported"),
	},
	Merkly: defi.Merkly{
		NFT: common.HexToAddress("0xFDc9018aF0E37AbF89233554C937eB5068127080"),
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

func NewClient(c *ClientConfig) (*Client, error) {

	config := &ClientConfig{
		HttpCli: &http.Client{},
	}
	if c != nil {
		config = c
	}

	ethcli, err := defi.NewEVMClient(&defi.Config{
		Network:   v1.Network_BinanaceBNB,
		MainToken: v1.Token_BNB,
		MainNet:   c.RPCEndpoint,
		TokenMap:  TokenAddress,
		Dict:      Dict,
		Httpcli:   config.HttpCli,
		TxViewFn: func(txId string) string {
			return "https://bscscan.com/tx/" + txId
		},
		NetworkId: bozdo.ChainMap[v1.Network_BinanaceBNB],
		EstimateL1Gas: func(ctx context.Context, data []byte) (*big.Int, error) {
			return big.NewInt(0), nil
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ethereum net: "+c.RPCEndpoint)
	}
	ethcli.London = defi.NeLondonReadyTx

	return &Client{
		defi:      ethcli,
		NetworkId: bozdo.ChainMap[v1.Network_BinanaceBNB],
	}, nil
}
