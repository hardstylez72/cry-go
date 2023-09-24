package task

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

func NewAcrossBridgeTask() *DefaultBridge {
	return NewDefaultBridgeTask(v1.TaskType_AcrossBridge, func(a *Input) (*v1.DefaultBridge, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_AcrossBridgeTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_AcrossBridge) call an ambulance!")
		}
		return l.AcrossBridgeTask, nil
	})
}

func NewDefaultBridgeTask(taskType v1.TaskType, extractor func(a *Input) (*v1.DefaultBridge, error)) *DefaultBridge {
	return &DefaultBridge{
		taskType:  taskType,
		extractor: extractor,
		cancel:    nil,
		DefaultBridgeTaskHalper: &DefaultBridgeTaskHalper{
			TaskType: taskType,
		},
	}
}

type DefaultBridge struct {
	taskType  v1.TaskType
	extractor func(a *Input) (*v1.DefaultBridge, error)
	cancel    func()
	*DefaultBridgeTaskHalper
}

type DefaultBridgeTaskHalper struct {
	TaskType v1.TaskType
}

func (t *DefaultBridge) Stop() error {
	t.cancel()
	return nil
}

func (t *DefaultBridge) Type() v1.TaskType {
	return t.taskType
}

func (t *DefaultBridge) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task

	p, err := t.extractor(a)
	if err != nil {
		return nil, err
	}

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
	if profile.Type != v1.ProfileType_EVM {
		return nil, errors.New("task is available only for emv account")
	}

	s, err := profile.GetNetworkSettings(ctx, p.FromNetwork)
	if err != nil {
		return nil, err
	}

	client, err := uniclient.NewBridger(p.FromNetwork, s.BaseConfig(), t.taskType)
	if err != nil {
		return nil, err
	}

	if p.GetTx().GetTxId() == "" {

		estimation, err := t.EstimateCost(taskContext, profile, p, client)
		if err != nil {
			return nil, errors.Wrap(err, "EstimateSwapCost of "+t.taskType.String())
		}
		res, gas, err := t.Execute(taskContext, profile, p, client, estimation)
		if err != nil {
			return nil, errors.Wrap(err, "Swap of "+t.taskType.String())
		}

		p.ApproveTx = NewTx(res.ApproveTx, gas)
		if err := a.AddTx2(ctx, p.ApproveTx); err != nil {
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

	if !p.GetReceived() {
		if err := client.WaitForConfirm(ctx, p.GetTx().GetTxId()); err != nil {
			return nil, err
		}
		p.Received = true
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	if p.GetTx().GetTxCompleted() && p.GetReceived() {
		task.Status = v1.ProcessStatus_StatusDone
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	return task, nil
}

func (h *DefaultBridgeTaskHalper) Execute(ctx context.Context, profile *halp.Profile, p *v1.DefaultBridge, client defi.Bridger, estimation *v1.EstimationTx) (*bozdo.DefaultRes, *bozdo.Gas, error) {

	s, err := profile.GetNetworkSettings(ctx, p.FromNetwork)
	if err != nil {
		return nil, nil, err
	}
	if client == nil {
		client, err = uniclient.NewBridger(p.FromNetwork, s.BaseConfig(), h.TaskType)
		if err != nil {
			return nil, nil, err
		}
	}

	balance, err := client.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: profile.Addr,
		Token:         p.FromToken,
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "client.GetFundingBalance")
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
		am = bozdo.Percent(am, 90)
		Gas = nil
	} else {
		gas, err := GasManager(estimation, s.Source, p.FromNetwork, h.TaskType)
		if err != nil {
			return nil, nil, err
		}
		Gas = gas

		balanceNative, err := client.GetBalance(ctx, &defi.GetBalanceReq{
			WalletAddress: profile.Addr,
			Token:         client.GetNetworkToken(),
		})
		if err != nil {
			return nil, nil, errors.Wrap(err, "client.GetFundingBalance")
		}
		if balanceNative.WEI.Cmp(&Gas.TotalGas) <= 0 {
			return nil, nil, ErrProfileHasInsufficientBalance(v1.Token_ETH, &Gas.TotalGas, balance.WEI)
		}
	}

	res, err := client.Bridge(ctx, &defi.DefaultBridgeReq{
		FromNetwork:  p.FromNetwork,
		ToNetwork:    p.ToNetwork,
		PK:           profile.WalletPK,
		Amount:       am,
		FromToken:    p.FromToken,
		ToToken:      p.ToToken,
		Gas:          Gas,
		Slippage:     getSlippage(s.Source, h.TaskType),
		EstimateOnly: estimateOnly,
		Debug:        false,
		SubType:      profile.SubType,
	}, h.TaskType)
	if err != nil {
		return nil, nil, err
	}
	return res, Gas, nil
}

func (h *DefaultBridgeTaskHalper) EstimateCost(ctx context.Context, profile *halp.Profile, p *v1.DefaultBridge, client defi.Bridger) (*v1.EstimationTx, error) {
	res, _, err := h.Execute(ctx, profile, p, client, nil)
	if err != nil {
		return nil, err
	}

	return GasStation(res.ECost, p.FromNetwork), nil
}
