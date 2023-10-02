package uniclient

import (
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/starknet"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/socks5"
	"github.com/pkg/errors"
)

func NewLPClient(network v1.Network, c *BaseClientConfig) (defi.LP, error) {

	proxy, err := socks5.NewSock5ProxyString(c.ProxyString, c.UserAgentHeader)
	if err != nil {
		return nil, err
	}

	var cli defi.LP
	switch network {
	case v1.Network_StarkNet:
		cli, err = starknet.NewClient(&starknet.ClientConfig{
			HttpCli:     proxy.Cli,
			RPCEndpoint: c.RPCEndpoint,
			Proxy:       c.ProxyString,
		})
	default:
		return nil, errors.New("network is not supported for LP")
	}
	return cli, err
}
