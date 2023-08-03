package defi

import (
	"context"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
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

var ZEROADDR = common.HexToAddress("0x0000000000000000000000000000000000000000")

func NewEVMClient(c *ClientConfig) (*EtheriumClient, error) {

	rpcClient, err := rpc.DialOptions(context.Background(), c.MainNet, rpc.WithHTTPClient(c.Httpcli))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ETH: "+c.MainNet)
	}

	return &EtheriumClient{
		Cli:              ethclient.NewClient(rpcClient),
		Cfg:              c,
		RpcCli:           rpcClient,
		traderJoeService: traderjoe.NewService(&traderjoe.Config{Host: config.CFG.HalperHost}),
	}, nil
}

func (c *EtheriumClient) ResoleGas(ctx context.Context, gas *Gas, opt *bind.TransactOpts) *bind.TransactOpts {

	if !gas.RuleSet() {
		head, errHead := c.Cli.HeaderByNumber(ctx, nil)
		if errHead == nil && head.BaseFee != nil {
			opt.GasFeeCap = big.NewInt(0).Mul(head.BaseFee, big.NewInt(3))
		}
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
