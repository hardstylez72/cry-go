package uniclient

import (
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/arbitrum"
	"github.com/hardstylez72/cry/internal/defi/base"
	"github.com/hardstylez72/cry/internal/defi/zksyncera"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/socks5"
	"github.com/pkg/errors"
)

func NewSwapper(network v1.Network, c *BaseClientConfig, taskType v1.TaskType) (defi.Swapper, error) {

	proxy, err := socks5.NewSock5ProxyString(c.ProxyString, c.UserAgentHeader)
	if err != nil {
		return nil, err
	}

	switch network {
	case v1.Network_ARBITRUM:
		return arbitrum.NewClient(&arbitrum.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint})
	case v1.Network_Base:
		return base.NewClient(&base.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint})
	case v1.Network_ZKSYNCERA:
		return zksyncera.NewClient(&zksyncera.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint})

	default:
		return nil, errors.New("unsupported network: " + taskType.String())
	}
}
