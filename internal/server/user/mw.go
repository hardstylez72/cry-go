package user

import (
	"context"

	"github.com/hardstylez72/cry/internal/lib"
	"github.com/pkg/errors"
)

type userIdKey struct {
}

type userGroupsKey struct {
}

type ControlledByUser struct {
}

var ErrUserNotFound = errors.New("user not found in context")

const GuestUserId = "2d4980a5-19cd-4e28-a703-b67e7767d0a1"

func SetUserIdContext(ctx context.Context, id string) context.Context {
	ctx = context.WithValue(ctx, &userIdKey{}, id)
	return ctx
}

func SetControlledByUserIdContext(ctx context.Context, id string) context.Context {
	ctx = context.WithValue(ctx, &ControlledByUser{}, id)
	return ctx
}

func getControlledBy(ctx context.Context) (string, error) {
	val := ctx.Value(&ControlledByUser{})
	if val == nil {
		return "", ErrUserNotFound
	}

	u, ok := val.(string)
	if !ok {
		return "", ErrUserNotFound
	}

	return u, nil
}
func UserId(ctx context.Context) (string, error) {
	return getUserId(ctx)
}

func ResolveUserId(ctx context.Context) (string, error) {
	owner, err := getControlledBy(ctx)
	if err == nil {
		return owner, nil
	}
	return getUserId(ctx)
}
func getUserId(ctx context.Context) (string, error) {

	val := ctx.Value(&userIdKey{})
	if val == nil {
		return "", ErrUserNotFound
	}

	u, ok := val.(string)
	if !ok {
		return "", ErrUserNotFound
	}

	return u, nil
}

func SetUserGroups(ctx context.Context, groups *lib.Set[string]) context.Context {
	ctx = context.WithValue(ctx, &userGroupsKey{}, groups)
	return ctx
}

func GetUserGroups(ctx context.Context) (*lib.Set[string], error) {
	val := ctx.Value(&userGroupsKey{})
	if val == nil {
		return nil, ErrUserNotFound
	}

	u, ok := val.(*lib.Set[string])
	if !ok {
		return nil, ErrUserNotFound
	}

	return u, nil
}
