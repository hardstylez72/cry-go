package v1

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi/starknet"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/scanner"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/socks5"
)

type StatService struct {
	v1.UnimplementedStatServiceServer
	profileRepository repository.ProfileRepository
	starknet          *starknet.Client
}

func NewStatService(profileRepository repository.ProfileRepository, starknet *starknet.Client) *StatService {
	return &StatService{
		profileRepository: profileRepository,
		starknet:          starknet,
	}
}

func (s *StatService) ZkSyncStat(ctx context.Context, req *v1.ZkSyncStatReq) (*v1.ZkSyncStatRes, error) {

	p, err := s.profileRepository.GetProfile(ctx, req.GetProfileId())
	if err != nil {
		return nil, err
	}

	pp, err := p.ToPB(s.starknet)
	if err != nil {
		return nil, err
	}

	proxy, err := socks5.NewSock5ProxyString((p.Proxy.String), "")
	if err != nil {
		return nil, err
	}

	zkscan := scanner.NewScannerZkSync(proxy.Cli)

	agg, err := scanner.Stat(ctx, zkscan, pp.MmskId)
	if err != nil {
		return nil, err
	}

	return &v1.ZkSyncStatRes{
		Stat: agg,
	}, nil
}
