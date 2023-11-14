package zksyncera

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func (c *Client) Swap(ctx context.Context, req *defi.DefaultSwapReq, taskType v1.TaskType) (*bozdo.DefaultRes, error) {
	switch taskType {
	case v1.TaskType_VeSyncSwap:
		return c.VeSyncSwap(ctx, req)
	case v1.TaskType_VelocoreSwap:
		return c.VelocoreSwap(ctx, req)
	case v1.TaskType_IzumiSwap:
		return c.IzumiSwap(ctx, req)
	case v1.TaskType_MaverickSwap:
		return c.MaverickSwap(ctx, req)
	case v1.TaskType_PancakeSwap:
		return c.PancakeSwap(ctx, req)
	case v1.TaskType_SpaceFISwap:

		tokenLimitChecker, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
			Token:       req.FromToken,
			WalletPK:    req.WalletPK,
			Amount:      req.Amount,
			SpenderAddr: c.Cfg.SpaceFI.Router,
		})
		if err != nil {
			return nil, errors.Wrap(err, "TokenLimitChecker")
		}

		res, err := c.SpaceFiSwap(ctx, req)
		if err != nil {
			return nil, err
		}
		if tokenLimitChecker.ApproveTx != nil {
			res.ApproveTx = tokenLimitChecker.ApproveTx
		}
		return res, nil

	case v1.TaskType_MuteioSwap:
		return c.MuteIOSwap(ctx, req)
	case v1.TaskType_SyncSwap:
		return c.SyncSwap(ctx, req)
	case v1.TaskType_ZkSwap:
		return c.ZkSwap(ctx, req)
	case v1.TaskType_EzkaliburSwap:
		return c.EzkaliburSwap(ctx, req)
	case v1.TaskType_OdosSwap:
		return c.OdosSwap(ctx, req)
	case v1.TaskType_KyberSwap:
		return c.KeyberSwap(ctx, req)
	default:
		return nil, errors.New("unsupported task type: " + taskType.String())
	}

}
