package settings

import (
	"context"
	"encoding/json"
	"time"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/arbitrum"
	"github.com/hardstylez72/cry/internal/defi/avalanche"
	"github.com/hardstylez72/cry/internal/defi/base"
	"github.com/hardstylez72/cry/internal/defi/bnb"
	"github.com/hardstylez72/cry/internal/defi/etherium"
	"github.com/hardstylez72/cry/internal/defi/optimism"
	"github.com/hardstylez72/cry/internal/defi/poligon"
	"github.com/hardstylez72/cry/internal/defi/starknet"
	"github.com/hardstylez72/cry/internal/defi/zksyncera"
	"github.com/hardstylez72/cry/internal/defi/zksynclite"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

const LastSettingsUpdateTime = "2023-10-11 7:40:05"

type GetSettingsNetworkRequest struct {
	Network v1.Network
	UserId  string
}

type Service struct {
	rep repository.SettingsRepository
}

type ResolveNet struct {
	Network v1.Network
	GasMax  string
	RPC     string
}

var Networks = []v1.Network{
	v1.Network_ARBITRUM,
	v1.Network_ZKSYNCERA,
	v1.Network_OPTIMISM,
	v1.Network_Etherium,
	v1.Network_BinanaceBNB,
	v1.Network_ZKSYNCLITE,
	v1.Network_AVALANCHE,
	v1.Network_POLIGON,
	v1.Network_StarkNet,
	v1.Network_Base,
}

func NewService(rep repository.SettingsRepository) *Service {
	return &Service{rep: rep}
}

func (s *Service) GetSettings(ctx context.Context, userId string, network v1.Network) (*v1.NetworkSettings, error) {
	return s.rep.GetSettings(ctx, userId, network)
}
func (s *Service) UpdateSettings(ctx context.Context, userId string, in *v1.NetworkSettings) error {
	return s.rep.UpdateSettings(ctx, userId, in)
}
func (s *Service) ResolveAllSettings(ctx context.Context, userId string, lastUpdate time.Time) error {
	for _, n := range Networks {

		updated, err := s.rep.GetSettingsDate(ctx, userId, n)
		if err != nil {
			if !errors.Is(err, repository.ErrNotFound) {
				return err
			}
		}

		forceUpdate := updated == nil || lastUpdate.After(*updated)
		if forceUpdate {
			if _, err := s.ResolveSettings(ctx, userId, n, false); err != nil {
				return err
			}
		}
	}
	return nil
}
func (s *Service) ResolveSettings(ctx context.Context, userId string, network v1.Network, force bool) (*v1.NetworkSettings, error) {

	rep := s.rep

	current, err := rep.GetSettings(ctx, userId, network)
	if err != nil {
		if !errors.Is(err, repository.ErrNotFound) {
			return nil, err
		}
	}

	before, err := json.Marshal(current)
	if err != nil {
		return nil, err
	}

	resolved, err := resolveSettings(current, network, force)
	if err != nil {
		return nil, err
	}

	after, err := json.Marshal(resolved)
	if err != nil {
		return nil, err
	}

	if string(after) == string(before) {
		return current, nil
	}
	if err := s.rep.UpdateSettings(ctx, userId, resolved); err != nil {
		return nil, err
	}

	return resolved, nil

}

func resolveSettings(in *v1.NetworkSettings, network v1.Network, force bool) (*v1.NetworkSettings, error) {

	s := ResolveNet{
		Network: network,
	}

	eth := "10000000000000000"
	poligonMax := "40000000000000000000"
	avalancheMax := "1000000000000000000"
	bnbGas := "100000000000000000"

	switch network {
	case v1.Network_POLIGON:
		s.RPC = poligon.MainNetURL
		s.GasMax = poligonMax
	case v1.Network_BinanaceBNB:
		s.RPC = bnb.MainNetURL
		s.GasMax = bnbGas
	case v1.Network_ARBITRUM:
		s.RPC = arbitrum.MainNetURL
		s.GasMax = eth
	case v1.Network_OPTIMISM:
		s.RPC = optimism.MainNetURL
		s.GasMax = eth
	case v1.Network_Etherium:
		s.RPC = etherium.MainNetURL
		s.GasMax = eth
	case v1.Network_ZKSYNCERA:
		s.RPC = zksyncera.MainNetURL
		s.GasMax = eth
	case v1.Network_ZKSYNCLITE:
		s.RPC = zksynclite.MainNetURL
		s.GasMax = eth
	case v1.Network_AVALANCHE:
		s.RPC = avalanche.MainNetURL
		s.GasMax = avalancheMax
	case v1.Network_StarkNet:
		s.RPC = starknet.MainnetRPC
		s.GasMax = eth
	case v1.Network_Base:
		s.RPC = base.MainNetURL
		s.GasMax = eth
	}

	return resolveNetworkSettings(in, s, force)
}
func resolveNetworkSettings(in *v1.NetworkSettings, s ResolveNet, force bool) (*v1.NetworkSettings, error) {

	out := in
	if force {
		out = nil
	}
	if out == nil {
		out = &v1.NetworkSettings{
			Network:      s.Network,
			TaskSettings: map[string]*v1.TaskSettings{},
		}
	}

	if out.TaskSettings == nil {
		out.TaskSettings = map[string]*v1.TaskSettings{}
	}

	if out.Id == "" {
		cli, err := uniclient.NewBaseClient(s.Network, &uniclient.BaseClientConfig{
			RPCEndpoint: s.RPC,
		})
		if err != nil {
			return nil, err
		}

		out.Id = cli.GetNetworkId().String()
		out.RpcEndpoint = s.RPC
		out.MaxGas = s.GasMax
		out.GasMultiplier = 1
	}

	// настройки сваполок и бриджей
	for taskType, slippage := range defi.SlippageMap {
		v, exist := out.TaskSettings[taskType.String()]
		if !exist {
			tmp := slippage
			out.TaskSettings[taskType.String()] = &v1.TaskSettings{
				Slippage: &tmp,
			}
		} else {
			if v.SwapRateRatio == nil {
				tmp := "1"
				out.TaskSettings[taskType.String()].SwapRateRatio = &tmp
			}

			if v.SwapUseExchangeRate == nil {
				tmp := false
				out.TaskSettings[taskType.String()].SwapUseExchangeRate = &tmp
			}
		}
	}

	return out, nil
}
