package starknet

import (
	"context"
	"net/http"
	"time"

	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/hardstylez72/cry/internal/defi/starknet/halper"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/config"
	"github.com/hardstylez72/cry/starknet.go/gateway"
	"github.com/hardstylez72/cry/starknet.go/rpc"
)

const MainnetRPC = "https://free-rpc.nethermind.io/mainnet-juno/v0_5"

type Client struct {
	GW          *gateway.Gateway
	GWP         *rpc.Provider
	NetworkId   string
	network     v1.Network
	cfg         *ClientConfig
	halper      halper.HalperService
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

	crpc, err := ethrpc.DialOptions(context.Background(),
		"https://starknet-mainnet.public.blastapi.io",
		ethrpc.WithHTTPClient(cfg.HttpCli))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	networkId, err := gw.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	gwp := rpc.NewProvider(crpc)

	host := "http://localhost:7272"
	if config.CFG != nil {
		host = config.CFG.HalperHost
	}

	return &Client{
		GW:        gw,
		GWP:       gwp,
		NetworkId: networkId,
		cfg:       cfg,
		halper: halper.NewService(&halper.Config{
			Host: host,
			RPC:  MainnetRPC,
		}),
		txViewFn: func(txId string) string {
			return "https://starkscan.co/tx/" + txId
		},
		NativeToken: v1.Token_ETH,
		network:     v1.Network_StarkNet,
		TokenMap: map[v1.Token]string{
			v1.Token_ETH:  "0x049d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7",
			v1.Token_USDC: "0x053c91253bc9682c04929ca02ed00b3e423f6710d2ee7e0d5ebb06f3ecf368a8",
			v1.Token_USDT: "0x68f5c6a61780768455de69077e07e89787839bf8166decfbf92b645209c0fb8",
		},
	}, nil
}

func (c *Client) TxViewFn(id string) string {
	return c.txViewFn(id)
}

func (c *Client) GetPublicKey(pk string, subType v1.ProfileSubType) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	res, err := c.halper.AccountPubKey(ctx, &halper.AccountPubKeyReq{PrivateKey: pk, SubType: subType.String()})
	if err != nil {
		return "", err
	}

	if len(res.PublicKey) < 64+2 {
		addZeros := len(res.PublicKey) - 64
		base := res.PublicKey[2:]

		for i := 0; i < addZeros; i++ {
			base = "0" + base
		}
		return "0x" + base, nil
	}

	return res.PublicKey, nil
}

func (c *Client) Network() v1.Network {
	return c.network
}
