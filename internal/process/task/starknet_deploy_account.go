package task

import (
	"context"
	"time"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/starknet"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/socks5"
	"github.com/pkg/errors"
)

type DeployStarkNetAccountTask struct {
	cancel func()
}

func (t *DeployStarkNetAccountTask) Stop() error {
	t.cancel()
	return nil
}

func (t *DeployStarkNetAccountTask) Type() v1.TaskType {
	return v1.TaskType_DeployStarkNetAccount
}

func (t *DeployStarkNetAccountTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_DeployStarkNetAccountTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_DeployStarkNetAccountTask) call an ambulance!")
	}

	p := l.DeployStarkNetAccountTask

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

	client, err := NewStarkNetClient(profile)
	if err != nil {
		return nil, err
	}

	if p.GetTx().GetTxId() == "" {

		estimation, err := EstimateDeployStarkNetAccountCost(taskContext, profile, p, client)
		if err != nil {
			return nil, errors.Wrap(err, "EstimateDeployStarkNetAccountCost")
		}
		res, gas, err := DeployStarkNetAccount(taskContext, profile, p, client, estimation)
		if err != nil {
			return nil, errors.Wrap(err, "DeployStarkNetAccount")
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

func DeployStarkNetAccount(ctx context.Context, profile *halp.Profile, p *v1.DeployStarkNetAccountTask, client *starknet.Client, estimation *v1.EstimationTx) (*bozdo.DefaultRes, *bozdo.Gas, error) {

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

	isDeployed, err := client.IsAccountDeployed(ctx, profile.WalletPK)
	if err != nil {
		return nil, nil, err
	}
	if *isDeployed {
		return nil, nil, errors.New("account already deployed on starknet")
	}

	balance, err := client.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: profile.Addr,
		Token:         v1.Token_ETH,
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "client.GetBalance")
	}

	estimateOnly := estimation == nil
	var Gas *bozdo.Gas
	if estimateOnly {
		Gas = nil
	} else {
		if balance.Float == 0 {
			return nil, nil, ErrAccountIsZero
		}

		gas, err := GasManager(estimation, s.Source, p.Network)
		if err != nil {
			return nil, nil, err
		}
		Gas = gas
	}

	res, err := client.DeployAccount(ctx, &starknet.DeployAccountReq{
		PK: profile.WalletPK,
		BaseReq: bozdo.BaseReq{
			EstimateOnly: estimateOnly,
			Gas:          nil,
			Debug:        false,
		},
	})
	if err != nil {
		return nil, nil, err
	}
	return res, Gas, nil
}

func EstimateDeployStarkNetAccountCost(ctx context.Context, profile *halp.Profile, p *v1.DeployStarkNetAccountTask, client *starknet.Client) (*v1.EstimationTx, error) {
	res, _, err := DeployStarkNetAccount(ctx, profile, p, client, nil)
	if err != nil {
		return nil, err
	}

	return GasStation(res.ECost, p.Network), nil
}

func NewStarkNetClient(profile *halp.Profile) (*starknet.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	s, err := profile.GetNetworkSettings(ctx, v1.Network_StarkNet)
	if err != nil {
		return nil, err
	}

	proxy, err := socks5.NewSock5ProxyString(s.BaseConfig().ProxyString, s.BaseConfig().UserAgentHeader)
	if err != nil {
		return nil, err
	}
	client, err := starknet.NewClient(&starknet.ClientConfig{
		HttpCli:     proxy.Cli,
		RPCEndpoint: s.BaseConfig().RPCEndpoint,
		Proxy:       s.BaseConfig().ProxyString,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
