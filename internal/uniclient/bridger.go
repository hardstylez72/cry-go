package uniclient

import (
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/zksyncera"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/socks5"
	"github.com/pkg/errors"
)

func NewBridger(network v1.Network, c *BaseClientConfig, taskType v1.TaskType) (defi.Bridger, error) {

	proxy, err := socks5.NewSock5ProxyString(c.ProxyString, c.UserAgentHeader)
	if err != nil {
		return nil, err
	}

	switch taskType {

	case v1.TaskType_AcrossBridge:
		switch network {
		case v1.Network_ZKSYNCERA:
			c, err := zksyncera.NewMainNetClient(&zksyncera.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint})
			if err != nil {
				return nil, err
			}
			return &zksyncera.AcrossBridge{Client: c, Debug: true, HttpCli: proxy.Cli}, nil
		default:
			return nil, errors.New("network is not supported for Transfer")
		}

	default:
		return nil, errors.New("unsupported taskType: " + taskType.String())
	}
}
