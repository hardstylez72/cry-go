package user

import (
	"context"

	"github.com/hardstylez72/cry/internal/lib"
)

const GroupSupport = "support"

var support = map[string]string{
	"korotkovcv77@gmail.com": GroupSupport,
}

type Getter interface {
	GetUserEmail(ctx context.Context, id string) (*string, error)
}

func Groups(ctx context.Context, r Getter) (*lib.Set[string], error) {

	userId, err := GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	email, err := r.GetUserEmail(ctx, userId)
	if err != nil {
		return nil, err
	}

	out := lib.NewSet[string]()
	_, ok := support[*email]
	if ok {
		out.Set(GroupSupport)
	}

	return out, nil
}
