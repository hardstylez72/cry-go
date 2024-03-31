package v1

import (
	"context"

	"github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/user"
	settings2 "github.com/hardstylez72/cry/internal/settings"
	"github.com/pkg/errors"
)

type SettingsService struct {
	v1.UnimplementedSettingsServiceServer
	settingsService *settings2.Service
}

func NewSettingsService(settingsService *settings2.Service) *SettingsService {
	return &SettingsService{settingsService: settingsService}
}

func (s *SettingsService) ResetSettings(ctx context.Context, req *v1.ResetRequest) (*v1.ResetResponse, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	if err := s.settingsService.UpdateSettings(ctx, userId, nil); err != nil {
		return nil, err
	}
	settings, err := s.settingsService.ResolveSettings(ctx, userId, req.GetNetwork(), true)
	if err != nil {
		return nil, err
	}

	return &v1.ResetResponse{
		Settings: settings,
	}, nil
}

func (s *SettingsService) GetSettings(ctx context.Context, req *v1.GetSettingsRequest) (*v1.GetSettingsResponse, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	settings, err := s.settingsService.GetSettings(ctx, userId, req.Network)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			resolveSettings, err := s.settingsService.ResolveSettings(ctx, userId, req.Network, true)
			if err != nil {
				return nil, err
			}

			return &v1.GetSettingsResponse{
				Settings: resolveSettings,
			}, nil
		}

		return nil, err
	}

	return &v1.GetSettingsResponse{
		Settings: settings,
	}, nil

}
func (s *SettingsService) UpdateSettings(ctx context.Context, req *v1.UpdateSettingsRequest) (*v1.UpdateSettingsResponse, error) {
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	if len(req.GetSettings().TaskSettings) == 0 ||
		req.GetSettings().Id == "" ||
		req.GetSettings().GetRpcEndpoint() == "" {
		return nil, errors.New("invalid input: 400")
	}

	if err := s.settingsService.UpdateSettings(ctx, userId, req.GetSettings()); err != nil {
		return nil, err
	}

	settings, err := s.settingsService.GetSettings(ctx, userId, req.GetSettings().GetNetwork())
	if err != nil {
		return nil, err
	}

	return &v1.UpdateSettingsResponse{
		Settings: settings,
	}, nil
}
