package task

import (
	"context"

	"github.com/hardstylez72/cry/internal/exchange"
	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WithdrawExchange struct {
	cancel func()
}

func (t *WithdrawExchange) Stop() error {
	t.cancel()
	return nil
}

func (t *WithdrawExchange) Type() v1.TaskType {
	return v1.TaskType_WithdrawExchange
}

func (t *WithdrawExchange) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_WithdrawExchangeTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_WithdrawExchangeTask) call an ambulance!")
	}
	p := l.WithdrawExchangeTask

	withdrawer, err := a.WithdrawerRepository.GetWithdrawer(ctx, p.WithdrawerId, a.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "a.WithdrawerRepository.GetWithdrawer")
	}

	profile, err := a.Halper.Profile(ctx, a.ProfileId)
	if err != nil {
		return nil, errors.Wrap(err, "ProfileRepository.GetProfile")
	}

	exchangeWithdrawer, err := uniclient.NewExchangeWithdrawer(withdrawer, profile.DB)
	if err != nil {
		return nil, errors.Wrap(err, "uniclient.NewExchangeWithdrawer")
	}

	switch a.Task.Status {
	case v1.ProcessStatus_StatusDone, v1.ProcessStatus_StatusError:
		return a.Task, nil
	case v1.ProcessStatus_StatusRunning:
		if p.WithdrawOrderId == nil {
			task.Status = v1.ProcessStatus_StatusReady
			if err := a.UpdateTask(ctx, task); err != nil {
				return nil, errors.Wrap(err, "a.UpdateTask")
			}
		}

		txId, err := exchangeWithdrawer.WaitConfirm(taskContext, *p.WithdrawOrderId)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				return a.Task, nil
			}
			return nil, errors.Wrap(err, "exchangeWithdrawer.WaitConfirm")
		}

		p.TxId = txId
		after := v1.ProcessStatus_StatusDone
		task.Status = after
		task.Task.Description = p.String()
		task.FinishedAt = timestamppb.Now()
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, errors.Wrap(err, "a.UpdateTask")
		}
		return task, nil

	case v1.ProcessStatus_StatusRetry, v1.ProcessStatus_StatusReady:

		if p.WithdrawOrderId != nil {
			after := v1.ProcessStatus_StatusRunning
			a.Task.Status = after
			if err := a.UpdateTask(ctx, task); err != nil {
				return nil, errors.Wrap(err, "a.UpdateTask")
			}
			return task, nil
		}

		after := v1.ProcessStatus_StatusRunning
		a.Task.Status = after
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, errors.Wrap(err, "a.UpdateTask")
		}
	}

	if p.GetSendToRelatedProfile() {

		if profile.Type != v1.ProfileType_StarkNet {
			return nil, errors.New("profile must have starknet type")
		}

		id, err := a.ProfileRepository.GetRelatedProfile(ctx, &repository.GetRelatedProfileReq{
			ProfileId:          profile.Id,
			ProfileType:        profile.Type.String(),
			ProfileSubType:     profile.SubType.String(),
			WantProfileType:    v1.ProfileType_EVM.String(),
			WantProfileSubType: v1.ProfileSubType_Metamask.String(),
			UserId:             profile.UserId,
		})
		if err != nil {
			return nil, errors.New("related EVM profile not found")
		}

		rp, err := a.Halper.Profile(ctx, *id)
		if err != nil {
			return nil, errors.New("related EVM profile not found")
		}

		p.WithdrawAddr = &rp.Addr
	} else {
		p.WithdrawAddr = &profile.Addr
	}

	if *p.SendAllCoins {
		b, err := exchangeWithdrawer.GetFundingBalance(ctx, p.Token)
		if err != nil {
			return nil, errors.Wrap(err, "exchangeWithdrawer.GetFundingBalance")
		}
		amount := lib.FloatToString(b)
		p.Amount = &amount
	} else {
		min, err := lib.StringToFloat(p.AmountMin)
		if err != nil {
			return nil, err
		}
		max, err := lib.StringToFloat(p.AmountMax)
		if err != nil {
			return nil, err
		}
		am := lib.FloatToString(lib.RandFloatRange(min, max))
		p.Amount = &am
	}

	res, err := exchangeWithdrawer.Withdraw(taskContext, &exchange.WithdrawRequest{
		ToAddress: *p.WithdrawAddr,
		Amount:    *p.Amount,
		Network:   p.Network,
		Token:     p.Token,
	})
	if err != nil {
		return nil, errors.Wrap(err, "exchangeWithdrawer.Withdraw")
	}

	p.WithdrawOrderId = &res.WithdrawId

	if err := a.UpdateTask(ctx, task); err != nil {
		return nil, errors.Wrap(err, "UpdateTask")
	}
	return task, nil
}
