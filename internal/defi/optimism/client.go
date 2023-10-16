package optimism

import (
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

// docs https://community.optimism.io/docs/useful-tools/networks/

const (
	MainNetURL = "https://rpc.ankr.com/optimism"
)

// https://arbiscan.io/tokens?p=1
var TokenAddress = map[defi.Token]common.Address{
	v1.Token_USDT:        common.HexToAddress("0x94b008aa00579c1307b0ef2c499ad98a8ce58e58"),
	v1.Token_STG:         common.HexToAddress("0x296F55F8Fb28E498B858d0BcDA06D955B2Cb3f97"),
	v1.Token_USDCBridged: common.HexToAddress("0x7F5c764cBc14f9669B88837ca1490cCa17c31607"),
	v1.Token_USDC:        common.HexToAddress("0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85"),
	v1.Token_veSTG:       common.HexToAddress("0xB0D502E938ed5f4df2E681fE6E419ff29631d62b"),
	v1.Token_ETH:         common.HexToAddress("0x0000000000000000000000000000000000000000"),
}

var Dict = defi.Dict{
	Stargate: defi.Stargate{
		StargateRouterAddress:    common.HexToAddress("0xB0D502E938ed5f4df2E681fE6E419ff29631d62b"),
		StargateRouterEthAddress: common.HexToAddress("0xB49c4e680174E331CB0A7fF3Ab58afC9738d5F8b"),
	},
	TestNetBridgeSwapAddress: common.HexToAddress("0x0A9f824C05A74F577A536A8A0c673183a872Dff4"),
	Across: defi.Across{
		RouterETH:   common.HexToAddress("0x269727F088F16E1Aea52Cf5a97B1CD41DAA3f02D"),
		RouterToken: common.HexToAddress("0x6f26Bf09B1C792e3228e5467807a900A503c0281"),
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
	return "https://optimistic.etherscan.io/tx/" + txId
}

func NewClient(c *ClientConfig) (*Client, error) {
	config := &ClientConfig{
		HttpCli: &http.Client{},
	}
	if c != nil {
		config = c
	}

	ethcli, err := defi.NewEVMClient(&defi.Config{
		Network:   v1.Network_OPTIMISM,
		MainToken: v1.Token_ETH,
		MainNet:   c.RPCEndpoint,
		TokenMap:  TokenAddress,
		Dict:      Dict,
		Httpcli:   config.HttpCli,
		TxViewFn:  TxViewer,
		NetworkId: bozdo.ChainMap[v1.Network_OPTIMISM],
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ethereum net: "+c.RPCEndpoint)
	}
	ethcli.Cfg.EstimateL1Gas = ethcli.EstimateL1GasFee

	return &Client{
		defi:      ethcli,
		NetworkId: bozdo.ChainMap[v1.Network_OPTIMISM],
	}, nil
}
