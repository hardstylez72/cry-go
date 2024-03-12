package base

import (
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

const (
	MainNetURL = "https://rpc.ankr.com/base"
)

// docs https://arbiscan.io/tokens?p=1
var TokenAddress = map[defi.Token]common.Address{
	v1.Token_ETH:         common.HexToAddress("0x0000000000000000000000000000000000000000"),
	v1.Token_WETH:        common.HexToAddress("0x4200000000000000000000000000000000000006"),
	v1.Token_USDC:        common.HexToAddress("0x833589fcd6edb6e08f4c7c32d4f71b54bda02913"),
	v1.Token_USDCBridged: common.HexToAddress("0xd9aaec86b65d86f6a7b5b1b0c42ffa531710b6ca"),
	v1.Token_STG:         common.HexToAddress("0xE3B53AF74a4BF62Ae5511055290838050bf764Df"),
}

var Dict = defi.Dict{
	Stargate: defi.Stargate{
		StargateRouterAddress:    common.HexToAddress("0x45f1A95A4D3f3836523F5c83673c797f4d4d263B"),
		StargateRouterEthAddress: common.Address{},
	},
	TestNetBridgeSwapAddress: common.HexToAddress(""),
	Merkly: defi.Merkly{
		NFT: common.HexToAddress("0xF882c982a95F4D3e8187eFE12713835406d11840"),
	},
	Pancake: defi.Pancake{
		Router:  SpecMap["PancakeRouter"].Addr,
		Factory: SpecMap["PancakeFactory"].Addr,
		Quoter:  SpecMap["PancakeQuoter"].Addr,
	},
	Across: defi.Across{
		RouterETH:   common.HexToAddress("0x269727F088F16E1Aea52Cf5a97B1CD41DAA3f02D"),
		RouterToken: common.HexToAddress("0x09aea4b2242abC8bb4BB78D537A67a245A7bEC64"),
	},
	Odos: defi.Odos{
		Router: common.HexToAddress("0x19cEeAd7105607Cd444F5ad10dd51356436095a1"),
	},
	Woofi: defi.Woofi{
		Router: common.HexToAddress("0x27425e9fb6a9a625e8484cfd9620851d1fa322e5"),
	},
	Aave: defi.Aave{
		LP: common.HexToAddress("0xA238Dd80C259a72e81d7e4664a9801593F98d1c5"),
	},
	MintFun: defi.MintFun{
		Minter: common.HexToAddress("0x69b69cc6e9f99c62a003fd9e143c126504d49dc2"),
	},
	Zerius: defi.Zerius{
		Minter: common.HexToAddress("0x178608fFe2Cca5d36f3Fc6e69426c4D3A5A74A41"),
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
	return "https://basescan.org/tx/" + txId
}

func NewClient(c *ClientConfig) (*Client, error) {

	config := &ClientConfig{
		HttpCli: &http.Client{},
	}
	if c != nil {
		config = c
	}

	ethcli, err := defi.NewEVMClient(&defi.Config{
		Network:   v1.Network_Base,
		MainToken: v1.Token_ETH,
		MainNet:   c.RPCEndpoint,
		TokenMap:  TokenAddress,
		Dict:      Dict,
		Httpcli:   config.HttpCli,
		TxViewFn:  TxViewer,
		NetworkId: bozdo.ChainMap[v1.Network_Base],
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ethereum main: "+c.RPCEndpoint)
	}

	ethcli.Cfg.EstimateL1Gas = ethcli.EstimateL1GasFee
	ethcli.London = defi.LondonReadyTx

	return &Client{
		defi:      ethcli,
		NetworkId: bozdo.ChainMap[v1.Network_Base],
	}, nil
}

var SpecMap = map[string]defi.Spec{
	"PancakeRouter": {
		Addr:     common.HexToAddress("0x678Aa4bF4E210cf2166753e054d5b7c31cc7fa86"),
		TaskType: v1.TaskType_PancakeSwap,
	},
	"PancakeFactory": {
		Addr:     common.HexToAddress("0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865"),
		TaskType: v1.TaskType_PancakeSwap,
	},
	"PancakeQuoter": {
		Addr:     common.HexToAddress("0xB048Bbc1Ee6b733FFfCFb9e9CeF7375518e25997"),
		TaskType: v1.TaskType_PancakeSwap,
	},
}
