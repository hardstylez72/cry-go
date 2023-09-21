package v1

import (
	"context"

	"github.com/hardstylez72/cry/internal/server/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Wellcome(ctx context.Context, groups ...string) error {
	g, err := user.GetUserGroups(ctx)
	if err != nil {
		return status.New(codes.Unauthenticated, "").Err()
	}

	for _, target := range groups {
		if g.Get(target) {
			return nil
		}
	}
	return status.New(codes.Unauthenticated, "user is not allowed to to the action").Err()
}

func NotWellcome(ctx context.Context, groups ...string) error {
	g, err := user.GetUserGroups(ctx)
	if err != nil {
		return status.New(codes.Unauthenticated, "").Err()
	}

	for _, target := range groups {
		if g.Get(target) {
			return status.New(codes.Unauthenticated, "action forbidden for group: "+target).Err()
		}
	}
	return nil
}
