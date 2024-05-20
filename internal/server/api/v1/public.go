package v1

import (
	"context"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process"
	"github.com/pkg/errors"
)

type PublicService struct {
	v1.UnimplementedPublicServiceServer
	d *process.Dispatcher
}

func NewPublicService(d *process.Dispatcher) *PublicService {
	return &PublicService{
		d: d,
	}
}

func (s *PublicService) SwapStat(ctx context.Context, req *v1.SwapStatReq) (*v1.SwapStatRes, error) {
	if s.d.SwapStatistic == nil {
		return nil, errors.New("no statistic")
	}
	return s.d.SwapStatistic, nil
}
