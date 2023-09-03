package task

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/orbiter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

type OrbiterBridgeTask struct {
	cancel func()
}

func (t *OrbiterBridgeTask) Stop() error {
	t.cancel()
	return nil
}

func (t *OrbiterBridgeTask) Type() v1.TaskType {
	return v1.TaskType_OrbiterBridge
}

func (t *OrbiterBridgeTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_OrbiterBridgeTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_OrbiterBridgeTask) call an ambulance!")
	}
	p := l.OrbiterBridgeTask

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

	s, err := profile.GetNetworkSettings(ctx, p.FromNetwork)
	if err != nil {
		return nil, err
	}
	client, err := uniclient.NewOrbiterSwapper(p.FromNetwork, s.BaseConfig())
	if err != nil {
		return nil, err
	}

	if p.GetTx().GetTxId() == "" {

		estimation, err := EstimateOrbiterBridgeCost(taskContext, a.Orbiter, profile, p)
		if err != nil {
			return nil, err
		}

		if err != nil {
			return nil, err
		}
		b, err := client.GetBalance(taskContext, &defi.GetBalanceReq{
			WalletAddress: profile.Addr,
			Token:         p.FromToken,
		})
		if err != nil {
			return nil, err
		}

		am, err := defi.ResolveAmount(p.Amount, b.WEI)
		if err != nil {
			return nil, err
		}

		gas, err := GasManager(estimation, s.Source, p.FromNetwork)
		if err != nil {
			return nil, err
		}

		if p.FromToken == client.GetNetworkToken() {
			am = ResolveNetworkTokenAmount(b.WEI, &gas.TotalGas, am)
		}

		res, err := client.OrbiterBridge(taskContext, &defi.OrbiterBridgeReq{
			OrbiterService: a.Orbiter,
			FromNetwork:    p.FromNetwork,
			ToNetwork:      p.ToNetwork,
			FromToken:      p.FromToken,
			ToToken:        p.ToToken,
			Amount:         am,
			WalletPk:       profile.WalletPK,
			Gas:            gas,
		})
		if err != nil {
			return nil, err
		}

		p.Tx = NewTx(res.Tx, gas)
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

func EstimateOrbiterBridgeCost(ctx context.Context, orbiterService *orbiter.Service, profile *halp.Profile, p *v1.OrbiterBridgeTask) (*v1.EstimationTx, error) {

	s, err := profile.GetNetworkSettings(ctx, p.FromNetwork)
	if err != nil {
		return nil, err
	}
	swapper, err := uniclient.NewOrbiterSwapper(p.FromNetwork, s.BaseConfig())
	if err != nil {
		return nil, err
	}
	
	b, err := swapper.GetBalance(ctx, &defi.GetBalanceReq{WalletAddress: profile.Addr, Token: p.FromToken})
	if err != nil {
		return nil, err
	}

	if b.Float == 0 {
		return nil, errors.New("not enough balance in " + p.FromToken.String())
	}

	am, err := defi.ResolveAmount(p.Amount, b.WEI)
	if err != nil {
		return nil, err
	}

	swap, err := swapper.OrbiterBridge(ctx, &defi.OrbiterBridgeReq{
		OrbiterService: orbiterService,
		FromNetwork:    p.FromNetwork,
		ToNetwork:      p.ToNetwork,
		FromToken:      p.FromToken,
		ToToken:        p.ToToken,
		Amount:         bozdo.Percent(am, 90),
		WalletPk:       profile.WalletPK,
		Gas:            nil,
		EstimateOnly:   true,
	})
	if err != nil {
		return nil, err
	}

	if p.FromToken == swapper.GetNetworkToken() {
		am = ResolveNetworkTokenAmount(b.WEI, swap.ECost.TotalGasWei, am)
	}

	return GasStation(swap.ECost, p.FromNetwork), nil
}
