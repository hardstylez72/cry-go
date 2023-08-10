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
	"github.com/hardstylez72/cry/internal/server/config"
	"github.com/hardstylez72/cry/internal/traderjoe"
	"github.com/pkg/errors"
)

const RetryMax = 1

type Dict struct {
	Stargate                 Stargate
	TestNetBridgeSwapAddress common.Address
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
	Cfg              *ClientConfig
	RpcCli           *rpc.Client
	traderJoeService *traderjoe.Service
}

type ClientConfig struct {
	Network   Network
	MainToken Token
	MainNet   string
	TokenMap  map[Token]common.Address
	Dict      *Dict
	Httpcli   *http.Client
	TxViewFn  func(s string) string
	networkId *big.Int
}

func NewEVMClient(c *ClientConfig) (*EtheriumClient, error) {

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
