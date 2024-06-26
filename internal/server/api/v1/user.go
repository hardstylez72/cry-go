package v1

import (
	"context"

	paycli "github.com/hardstylez72/cry-pay/proto/gen/go/v1"
	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository/pg"
	"github.com/hardstylez72/cry/internal/server/user"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *HelperService) GetUser(ctx context.Context, _ *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, status.New(codes.Unauthenticated, "").Err()
	}

	u, err := s.userRepository.GetUserById(ctx, userId)
	if err != nil {
		if errors.Is(err, pg.EntityNotFound) {
			return nil, status.New(codes.Unauthenticated, "").Err()
		}
		return nil, err
	}

	res, err := s.payService.GetAccount(ctx, &paycli.GetAccountReq{
		Id: u.Id,
	})
	if err != nil {
		return nil, err
	}

	groups, err := user.Groups(&user.User{
		Id:           u.Id,
		Email:        u.Email,
		ControlledBy: u.ControlledBy,
	})
	if err != nil {
		return nil, err
	}

	return &v1.GetUserResponse{
		Id:        u.Id,
		Email:     u.Email,
		Funds:     lib.FloatToString(res.GetAccount().GetFunds()),
		TaskPrice: lib.FloatToString(res.GetAccount().GetTaskPrice()),
		//PayableTasks:    task.PayableTasks,
		//NonpayableTasks: task.NonPayableTasks,
		Groups: groups.Keys(),
		Promo:  res.Account.Promo,
	}, nil
}
