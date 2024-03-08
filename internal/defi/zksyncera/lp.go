package zksyncera

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func (c *Client) LP(ctx context.Context, req *defi.LPReq, taskType v1.TaskType) (*defi.LPRes, error) {

	switch taskType {
	case v1.TaskType_EraLend:
		return NewEraLendEraLend(c).LP(ctx, req)
	default:
		return nil, errors.New("unsupported task type: " + taskType.String())
	}
}
