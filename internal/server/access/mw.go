package access

import (
	"context"

	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Access(r repository.UserRepository) func(ctx context.Context) (context.Context, error) {

	return func(ctx context.Context) (context.Context, error) {

		id, err := user.UserId(ctx)
		if err != nil {
			return nil, status.New(codes.Unauthenticated, "").Err()
		}
		u, err := r.GetUserById(ctx, id)
		if err != nil {
			return nil, status.New(codes.Unauthenticated, "").Err()
		}

		groups, err := user.Groups(&user.User{
			Id:           u.Id,
			Email:        u.Email,
			ControlledBy: u.ControlledBy,
		})
		if err != nil {
			return nil, status.New(codes.Unauthenticated, "").Err()
		}

		if u.ControlledBy.Valid {
			ctx = user.SetControlledByUserIdContext(ctx, u.ControlledBy.String)
		}

		ctx = user.SetUserGroups(ctx, groups)

		return ctx, nil
	}

}
