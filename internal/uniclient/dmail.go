package uniclient

import (
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/starknet"
	"github.com/hardstylez72/cry/internal/defi/zksyncera"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/socks5"
	"github.com/pkg/errors"
)

func NewDmailClient(network v1.Network, c *BaseClientConfig) (defi.DmailSender, error) {

	proxy, err := socks5.NewSock5ProxyString(c.ProxyString, c.UserAgentHeader)
	if err != nil {
		return nil, err
	}

	var cli defi.DmailSender
	switch network {
	case v1.Network_ZKSYNCERA:
		cli, err = zksyncera.NewClient(&zksyncera.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint})
	case v1.Network_StarkNet:
		cli, err = starknet.NewClient(&starknet.ClientConfig{
			HttpCli:     proxy.Cli,
			RPCEndpoint: c.RPCEndpoint,
			Proxy:       c.ProxyString,
		})
	default:
		return nil, errors.New("network is not supported for DmailSender")
	}
	return cli, err
}
