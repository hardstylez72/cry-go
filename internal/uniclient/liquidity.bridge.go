package uniclient

import (
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/etherium"
	"github.com/hardstylez72/cry/internal/defi/starknet"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/socks5"
	"github.com/pkg/errors"
)

func NewLiquidityBridge(from, to v1.Network, c *BaseClientConfig, taskType v1.TaskType) (defi.LiquidityBridger, defi.Networker, error) {

	proxy, err := socks5.NewSock5ProxyString(c.ProxyString, c.UserAgentHeader)
	if err != nil {
		return nil, nil, err
	}

	switch taskType {

	case v1.TaskType_StarkNetBridge:
		switch from {
		case v1.Network_Etherium:
			client, err := starknet.NewClient(&starknet.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint, Proxy: c.ProxyString})
			if err != nil {
				return nil, nil, err
			}
			networker, err := etherium.NewClient(&etherium.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint})
			if err != nil {
				return nil, nil, err
			}
			return client, networker, nil
		default:
			return nil, nil, errors.New("network is not supported for Transfer")
		}

	default:
		return nil, nil, errors.New("unsupported taskType: " + taskType.String())
	}
}
