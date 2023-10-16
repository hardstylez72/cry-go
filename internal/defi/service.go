package defi

import (
	"context"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/config"
	"github.com/hardstylez72/cry/internal/traderjoe"
	"github.com/pkg/errors"
)

type Dict struct {
	Stargate                 Stargate
	TestNetBridgeSwapAddress common.Address
	Merkly                   Merkly
	Pancake                  Pancake
	Across                   Across
}

type Across struct {
	RouterETH   common.Address `json:"router"`
	RouterToken common.Address `json:"routerToken"`
}

type Pancake struct {
	Router  common.Address `json:"router"`
	Factory common.Address `json:"factory"`
	Quoter  common.Address `json:"quoter"`
}

type Merkly struct {
	NFT common.Address
}

type SyncSwap struct {
	RouterSwap         common.Address
	ClassicPoolFactory common.Address
}

type Stargate struct {
	StargateRouterAddress    common.Address
	StargateRouterEthAddress common.Address
}

type EtheriumClient struct {
	Cli              *ethclient.Client
	Cfg              *Config
	RpcCli           *rpc.Client
	traderJoeService *traderjoe.Service
}

func (c *EtheriumClient) Network() v1.Network {
	return c.Cfg.Network
}

func (c *EtheriumClient) GetNetworkId() *big.Int {
	return c.Cfg.NetworkId
}

func (c *EtheriumClient) GetPublicKey(pk string, subType v1.ProfileSubType) (string, error) {
	t, err := NewWalletTransactor(pk)
	if err != nil {
		return "", err
	}
	return t.WalletAddrHR, nil
}

type EstimateL1Gas func(ctx context.Context, data []byte) (*big.Int, error)

type Config struct {
	Network       Network
	MainToken     Token
	MainNet       string
	TokenMap      map[Token]common.Address
	Dict          Dict
	Httpcli       *http.Client
	TxViewFn      func(s string) string
	NetworkId     *big.Int
	EstimateL1Gas EstimateL1Gas
}

func NewEVMClient(c *Config) (*EtheriumClient, error) {

	rpcClient, err := rpc.DialOptions(context.Background(), c.MainNet, rpc.WithHTTPClient(c.Httpcli))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ETH: "+c.MainNet)
	}

	traderJoe := ""
	if config.CFG != nil {
		traderJoe = config.CFG.HalperHost
	}

	return &EtheriumClient{
		Cli:              ethclient.NewClient(rpcClient),
		Cfg:              c,
		RpcCli:           rpcClient,
		traderJoeService: traderjoe.NewService(&traderjoe.Config{Host: traderJoe}),
	}, nil
}

func (c *EtheriumClient) ResoleGas(ctx context.Context, gas *bozdo.Gas, opt *bind.TransactOpts) *bind.TransactOpts {

	if !gas.RuleSet() {
		return opt
	}

	head, errHead := c.Cli.HeaderByNumber(ctx, nil)
	if errHead == nil && head.BaseFee != nil {
		opt.GasLimit = gas.GasLimit.Uint64()
		opt.GasFeeCap = &gas.GasPrice
	} else {
		opt.GasLimit = gas.GasLimit.Uint64()
		opt.GasPrice = &gas.GasPrice
	}

	return opt
}
