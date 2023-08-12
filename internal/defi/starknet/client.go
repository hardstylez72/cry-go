package starknet

import (
	"context"
	"net/http"
	"time"

	"github.com/dontpanicdao/caigo/gateway"
	"github.com/hardstylez72/cry/internal/defi/starknet/halper"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/config"
)

const MainnetRPC = gateway.MAINNET_BASE

type Client struct {
	GW        *gateway.Gateway
	GWP       *gateway.GatewayProvider
	NetworkId string
	cfg       *ClientConfig
	halper.HalperService
	txViewFn    func(s string) string
	NativeToken v1.Token
	TokenMap    map[v1.Token]string
}

type ClientConfig struct {
	HttpCli     *http.Client
	RPCEndpoint string
	Proxy       string
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
		GWP:       gwp,
		NetworkId: networkId,
		cfg:       cfg,
		HalperService: halper.NewService(&halper.Config{
			Host: config.CFG.HalperHost,
			RPC:  MainnetRPC,
		}),
		txViewFn: func(txId string) string {
			return "https://starkscan.co/tx/" + txId
		},
		NativeToken: v1.Token_ETH,
		TokenMap: map[v1.Token]string{
			v1.Token_ETH: "0x049d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7",
		},
	}, nil
}

func (c *Client) TxViewFn(id string) string {
	return c.txViewFn(id)
}

func (c *Client) GetPublicKey(pk string) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	res, err := c.HalperService.AccountPubKey(ctx, &halper.AccountPubKeyReq{PrivateKey: pk})
	if err != nil {
		return ""
	}

	if len(res.PublicKey) < 64+2 {
		addZeros := len(res.PublicKey) - 64
		base := res.PublicKey[2:]

		for i := 0; i < addZeros; i++ {
			base = "0" + base
		}
		return "0x" + base
	}

	return res.PublicKey
}
