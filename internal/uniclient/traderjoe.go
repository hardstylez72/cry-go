package uniclient

import (
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/arbitrum"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/socks5"
	"github.com/pkg/errors"
)

func NewTraderJoe(network v1.Network, c *BaseClientConfig) (defi.TraderJoeSwap, error) {

	proxy, err := socks5.NewSock5ProxyString(c.ProxyString, c.UserAgentHeader)
	if err != nil {
		return nil, err
	}

	var cli defi.TraderJoeSwap
	switch network {
	case v1.Network_ARBITRUM:
		cli, err = arbitrum.NewClient(&arbitrum.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint})
	default:
		return nil, errors.New("network is not supported for Transfer")
	}

	if err != nil {
		return nil, err
	}

	return cli, nil
}
