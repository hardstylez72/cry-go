package starknet

import (
	"context"
	"net/http"
	"time"

	"github.com/dontpanicdao/caigo/gateway"
)

const MainnetRPC = gateway.MAINNET_BASE

type Client struct {
	GW        *gateway.Gateway
	GWP       *gateway.GatewayProvider
	NetworkId string
	cfg       *ClientConfig
}

type ClientConfig struct {
	HttpCli     *http.Client
	RPCEndpoint string
}

func NewClient(cfg *ClientConfig) (*Client, error) {

	gw := gateway.NewClient(
		gateway.WithChain(cfg.RPCEndpoint),
		gateway.WithHttpClient(*cfg.HttpCli),
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	networkId, err := gw.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	gwp := gateway.NewProvider(
		gateway.WithChain(cfg.RPCEndpoint),
		gateway.WithHttpClient(*cfg.HttpCli),
	)

	return &Client{
		GW:        gw,
		NetworkId: networkId,
		cfg:       cfg,
		GWP:       gwp,
	}, nil
}
