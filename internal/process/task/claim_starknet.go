package task

import (
	"context"
	"net/http"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/starknet"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/pkg/errors"
)

type ClaimStarknet struct {
	cancel func()
}

func (t *ClaimStarknet) Stop() error {
	t.cancel()
	return nil
}

func (t *ClaimStarknet) Type() v1.TaskType {
	return v1.TaskType_StarknetClaim
}

func (t *ClaimStarknet) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_StarknetClaim)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_StarknetClaim) call an ambulance!")
	}

	p := l.StarknetClaim

	tm := taskTimeout
	if p.Network == v1.Network_StarkNet {
		tm = taskStarkNetTimeout
	}

	taskContext, cancel := context.WithTimeout(ctx, tm)
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

	client, err := starknet.NewClient(&starknet.ClientConfig{
		HttpCli:     &http.Client{},
		RPCEndpoint: s.BaseConfig().RPCEndpoint,
		Proxy:       s.BaseConfig().ProxyString,
	})
	if err != nil {
		return nil, errors.Wrap(err, "starknet.NewClient")
	}

	if p.GetTx().GetTxId() == "" {

		estimation, err := EstimateDmailSwapCost(taskContext, p, profile)
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

func (t *ClaimStarknet) Execute(ctx context.Context, p *v1.SimpleTask, client *starknet.Client, profile *halp.Profile, estimation *v1.EstimationTx) (*bozdo.DefaultRes, *bozdo.Gas, error) {
	s, err := profile.GetNetworkSettings(ctx, p.Network)
	if err != nil {
		return nil, nil, err
	}

	estimateOnly := estimation == nil
	var Gas *bozdo.Gas
	if estimateOnly {
		Gas = nil
	} else {
		gas, err := GasManager(estimation, s.Source, p.Network, t.Type())
		if err != nil {
			return nil, nil, err
		}
		Gas = gas
	}

	res, err := client.Claim(ctx, &defi.SimpleReq{
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

func EstimateStarknetClaimCost(ctx context.Context, p *v1.SimpleTask, profile *halp.Profile) (*v1.EstimationTx, error) {

	s, err := profile.GetNetworkSettings(ctx, p.Network)
	if err != nil {
		return nil, err
	}

	client, err := starknet.NewClient(&starknet.ClientConfig{
		HttpCli:     &http.Client{},
		RPCEndpoint: s.BaseConfig().RPCEndpoint,
		Proxy:       s.BaseConfig().ProxyString,
	})
	if err != nil {
		return nil, err
	}

	swap, err := client.Claim(ctx, &defi.SimpleReq{
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
