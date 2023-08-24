package uniclient

import (
	"github.com/hardstylez72/cry/internal/defi/zksyncera"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/socks5"
	"github.com/pkg/errors"
)

func NewZkSyncOfficialBridge(network v1.Network, c *BaseClientConfig) (*zksyncera.Client, error) {
	proxy, err := socks5.NewSock5ProxyString(c.ProxyString, c.UserAgentHeader)
	if err != nil {
		return nil, err
	}

	switch network {
	case v1.Network_ZKSYNCERATESTNET:
		return zksyncera.NewTestNetClient(&zksyncera.ClientConfig{
			HttpCli:     proxy.Cli,
			RPCEndpoint: c.RPCEndpoint,
		})
	case v1.Network_ZKSYNCERA:
		return zksyncera.NewMainNetClient(&zksyncera.ClientConfig{
			HttpCli:     proxy.Cli,
			RPCEndpoint: c.RPCEndpoint,
		})
	default:
		return nil, errors.New("unsupported network: " + network.String())
	}
}
