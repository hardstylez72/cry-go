package task

import (
	"context"
	"time"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

type MintTask struct {
	cancel func()
}

func (t *MintTask) Stop() error {
	t.cancel()
	return nil
}

func (t *MintTask) Type() v1.TaskType {
	return v1.TaskType_StarkNetIdMint
}

func Timeout(n v1.Network) time.Duration {
	tm := taskTimeout
	if n == v1.Network_StarkNet {
		tm = taskStarkNetTimeout
	}

	return tm
}

func (t *MintTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_StarkNetIdMintTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_StarkNetIdMint) call an ambulance!")
	}

	p := l.StarkNetIdMintTask

	taskContext, cancel := context.WithTimeout(ctx, Timeout(p.Network))
	defer cancel()

	t.cancel = cancel

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

	s, err := profile.GetNetworkSettings(ctx, p.Network)
	if err != nil {
		return nil, err
	}

	client, err := uniclient.NewMintClient(p.Network, s.BaseConfig())
	if err != nil {
		return nil, err
	}

	if p.GetApproveTx().GetTxId() == "" && p.Network == v1.Network_StarkNet {
		client, err := NewStarkNetClient(profile)
		if err != nil {
			return nil, err
		}

		txId, err := StarkNetApprove(taskContext, v1.Token_ETH, client, profile, t.Type())
		if err != nil {
			return nil, err
		}

		if txId != nil {
			p.ApproveTx = NewStarkNetApproveTx(*txId)
			if err := a.AddTx2(ctx, p.ApproveTx); err != nil {
				return nil, err
			}
			if err := a.UpdateTask(ctx, task); err != nil {
				return nil, err
			}
		}
	}

	if err := WaitTxComplete(taskContext, p.ApproveTx, task, client, a); err != nil {
		return nil, err
	}

	if p.GetTx().GetTxId() == "" {

		estimation, err := EstimateMintCost(taskContext, p, profile)
		if err != nil {
			return nil, err
		}

		res, gas, err := t.Execute(taskContext, p, client, profile, estimation)
		if err != nil {
			return nil, err
		}

		if err := a.AddTx(ctx, res.ApproveTx); err != nil {
			return nil, err
		}
		p.Tx = NewTx(res.Tx, gas)

		if err := a.AddTx2(ctx, p.Tx); err != nil {
			return nil, err
		}
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	if err := WaitTxComplete(taskContext, p.Tx, task, client, a); err != nil {
		return nil, err
	}

	task.Status = v1.ProcessStatus_StatusDone
	if err := a.UpdateTask(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (t *MintTask) Execute(ctx context.Context, p *v1.SimpleTask, client defi.Minter, profile *halp.Profile, estimation *v1.EstimationTx) (*bozdo.DefaultRes, *bozdo.Gas, error) {
	s, err := profile.GetNetworkSettings(ctx, p.Network)
	if err != nil {
		return nil, nil, err
	}

	estimateOnly := estimation == nil
	var Gas *bozdo.Gas
	if estimateOnly {
		Gas = nil
	} else {
		gas, err := GasManager(estimation, s.Source, p.Network)
		if err != nil {
			return nil, nil, err
		}
		Gas = gas
	}

	res, err := client.Mint(ctx, &defi.SimpleReq{
		PK:      profile.WalletPK,
		SubType: profile.SubType,
		BaseReq: &bozdo.BaseReq{
			EstimateOnly: false,
			Debug:        true,
			Gas:          Gas,
		},
	})
	if err != nil {
		return nil, nil, err
	}

	return res, Gas, nil
}

func EstimateMintCost(ctx context.Context, p *v1.SimpleTask, profile *halp.Profile) (*v1.EstimationTx, error) {

	s, err := profile.GetNetworkSettings(ctx, p.Network)
	if err != nil {
		return nil, err
	}

	if p.Network == v1.Network_StarkNet {
		client, err := NewStarkNetClient(profile)
		if err != nil {
			return nil, err
		}

		approve, err := StarkNetApprove(ctx, v1.Token_ETH, client, profile, v1.TaskType_StarkNetIdMint)
		if err != nil {
			return nil, err
		}

		if approve != nil {
			if err := client.WaitTxComplete(ctx, *approve); err != nil {
				return nil, err
			}
		}
	}

	client, err := uniclient.NewMintClient(p.Network, s.BaseConfig())
	if err != nil {
		return nil, err
	}

	swap, err := client.Mint(ctx, &defi.SimpleReq{
		PK:      profile.WalletPK,
		SubType: profile.SubType,
		BaseReq: &bozdo.BaseReq{
			EstimateOnly: true,
			Debug:        false,
		},
	})
	if err != nil {
		return nil, err
	}

	return GasStation(swap.ECost, p.Network), nil
}