package task

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/starknet"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/pkg/errors"
)

type Swap10kTask struct {
	cancel func()
}

func (t *Swap10kTask) Stop() error {
	t.cancel()
	return nil
}

func (t *Swap10kTask) Reset(ctx context.Context, a *Input) error {
	task := a.Task

	if err := a.UpdateTask(ctx, task); err != nil {
		return err
	}

	return nil
}

func (t *Swap10kTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_Swap10K)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_Swap10kTask) call an ambulance!")
	}

	p := l.Swap10K

	switch a.Task.Status {
	case v1.ProcessStatus_StatusDone, v1.ProcessStatus_StatusError:
		return a.Task, nil
	case v1.ProcessStatus_StatusRetry, v1.ProcessStatus_StatusReady, v1.ProcessStatus_StatusRunning:

		task.Status = v1.ProcessStatus_StatusRunning
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	profile, err := a.Halper.Profile(ctx, a.ProfileId)
	if err != nil {
		return nil, err
	}
	if profile.Type != v1.ProfileType_StarkNet {
		return nil, errors.New("task is available only for starknet account")
	}

	client, err := NewStarkNetClient(profile)
	if err != nil {
		return nil, err
	}

	if p.GetTx().GetTxId() == "" {

		estimation, err := EstimateSwap10kCost(taskContext, profile, p, client)
		if err != nil {
			return nil, errors.Wrap(err, "EstimateSwap10kCost")
		}
		res, gas, err := Swap10k(taskContext, profile, p, client, estimation)
		if err != nil {
			return nil, errors.Wrap(err, "Swap10k")
		}

		p.Tx = NewTx(res.Tx, gas)
		if err := a.AddTx(ctx, res.ApproveTx); err != nil {
			return nil, err
		}

		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	if err := WaitTxComplete(taskContext, p.Tx, task, client, a); err != nil {
		return nil, err
	}
	if err := a.AddTx2(ctx, p.Tx); err != nil {
		return nil, err
	}

	if p.GetTx().GetTxCompleted() {
		task.Status = v1.ProcessStatus_StatusDone
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	return task, nil
}

func Swap10k(ctx context.Context, profile *halp.Profile, p *v1.DefaultSwap, client *starknet.Client, estimation *v1.EstimationTx) (*bozdo.DefaultRes, *bozdo.Gas, error) {

	s, err := profile.GetNetworkSettings(ctx, p.Network)
	if err != nil {
		return nil, nil, err
	}
	if client == nil {
		client, err = NewStarkNetClient(profile)
		if err != nil {
			return nil, nil, err
		}
	}

	balance, err := client.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: profile.Addr,
		Token:         p.FromToken,
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "client.GetBalance")
	}

	am, err := defi.ResolveAmount(p.Amount, balance.WEI)
	if err != nil {
		return nil, nil, err
	}

	if am == nil || am.Cmp(big.NewInt(0)) == 0 {
		return nil, nil, errors.New("not enough balance of " + p.FromToken.String())
	}

	estimateOnly := estimation == nil
	var Gas *bozdo.Gas
	if estimateOnly {
		am = bozdo.Percent(am, 50)
		Gas = nil
	} else {
		gas, err := GasManager(estimation, s.Source, p.Network)
		if err != nil {
			return nil, nil, err
		}
		Gas = gas
	}

	res, err := client.Swap10k(ctx, &defi.DefaultSwapReq{
		Network:      v1.Network_StarkNet,
		Amount:       am,
		FromToken:    p.FromToken,
		ToToken:      p.ToToken,
		WalletPK:     profile.WalletPK,
		EstimateOnly: estimateOnly,
		Gas:          Gas,
		Debug:        false,
		Slippage:     getSlippage(s.Source, v1.TaskType_Swap10k),
	})
	if err != nil {
		return nil, nil, err
	}
	return res, Gas, nil
}

func EstimateSwap10kCost(ctx context.Context, profile *halp.Profile, p *v1.DefaultSwap, client *starknet.Client) (*v1.EstimationTx, error) {
	res, _, err := Swap10k(ctx, profile, p, client, nil)
	if err != nil {
		return nil, err
	}

	return GasStation(res.ECost, p.Network), nil
}
