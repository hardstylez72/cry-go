package task

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/exchange"
	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

type ExchangeSwapTask struct {
	cancel func()
}

func (t *ExchangeSwapTask) Stop() error {
	t.cancel()
	return nil
}

func (t *ExchangeSwapTask) Type() v1.TaskType {
	return v1.TaskType_OkexDeposit
}

func (t *ExchangeSwapTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_ExchangeSwapTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_OkexDepositTask) call an ambulance!")
	}

	p := l.ExchangeSwapTask

	switch a.Task.Status {
	case v1.ProcessStatus_StatusDone, v1.ProcessStatus_StatusError:
		return a.Task, nil
	case v1.ProcessStatus_StatusRetry, v1.ProcessStatus_StatusReady:
		task.Status = v1.ProcessStatus_StatusRunning
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}
	withdrawer, err := a.WithdrawerRepository.GetWithdrawer(ctx, p.WithdrawerId, a.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "a.WithdrawerRepository.GetWithdrawer")
	}

	profile, err := a.ProfileRepository.GetProfile(ctx, a.ProfileId)
	if err != nil {
		return nil, errors.Wrap(err, "a.ProfileRepository.GetProfile")
	}

	client, err := uniclient.NewExchangeSwapper(withdrawer, profile)
	if err != nil {
		return nil, err
	}

	percentString, err := defi.GetPercent(p.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "defi.GetPercent")
	}

	float, err := lib.StringToFloat(percentString)
	if err != nil {
		return nil, err
	}
	req := &exchange.SwapReq{
		From:      p.FromToken,
		To:        p.ToToken,
		AmPercent: float,
	}

	if p.GetBefore() == false {

		err = client.Before(taskContext, req)
		if err != nil {
			return nil, errors.Wrap(err, "client.Before")
		}

		tmp := true
		p.Before = &tmp
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	if p.GetTradeId() == "" {
		swap, err := client.Swap(taskContext, req)
		if err != nil {
			return nil, err
		}

		p.TradeId = &swap.TradeId
		p.Pair = &swap.Pair
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	if p.GetTradeId() != "" {
		if err := client.WaitSwapComplete(taskContext, &exchange.SwapRes{
			Pair:    p.GetPair(),
			TradeId: p.GetTradeId(),
		}); err != nil {
			return nil, err
		}

		tmp := true
		p.SwapCompleted = &tmp
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	if p.GetSwapCompleted() && !p.GetAfter() {
		if err := client.After(taskContext, req); err != nil {
			return nil, err
		}
		tmp := true
		p.After = &tmp
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	if p.GetAfter() {
		task.Status = v1.ProcessStatus_StatusDone
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	return task, nil
}
