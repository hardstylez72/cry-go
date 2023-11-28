package task

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/exchange/pub"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

func NewKyberSwapTask() *DefaultSwapTask {
	return NewDefaultSwapTaskTask(v1.TaskType_KyberSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_KyberSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_WoofiSwapTask) call an ambulance!")
		}
		return l.KyberSwapTask, nil
	})
}

func NewWoofiSwapTask() *DefaultSwapTask {
	return NewDefaultSwapTaskTask(v1.TaskType_WoofiSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_WoofiSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_WoofiSwapTask) call an ambulance!")
		}
		return l.WoofiSwapTask, nil
	})
}

func NewTraderJoeSwapTask() *DefaultSwapTask {
	return NewDefaultSwapTaskTask(v1.TaskType_TraderJoeSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_TraderJoeSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_TraderJoeSwapTask) call an ambulance!")
		}
		return l.TraderJoeSwapTask, nil
	})
}

func NewPancakeSwapTask() *DefaultSwapTask {
	return NewDefaultSwapTaskTask(v1.TaskType_PancakeSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_PancakeSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_PancakeSwapTask) call an ambulance!")
		}
		return l.PancakeSwapTask, nil
	})
}

func NewOdosSwapTask() *DefaultSwapTask {
	return NewDefaultSwapTaskTask(v1.TaskType_OdosSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_OdosSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_OdosSwapTask) call an ambulance!")
		}
		return l.OdosSwapTask, nil
	})
}

type DefaultSwapTask struct {
	taskType  v1.TaskType
	extractor func(a *Input) (*v1.DefaultSwap, error)
	cancel    func()
	*DefaultSwapTaskHalper
}

func NewDefaultSwapTaskTask(taskType v1.TaskType, extractor func(a *Input) (*v1.DefaultSwap, error)) *DefaultSwapTask {
	return &DefaultSwapTask{
		taskType:  taskType,
		extractor: extractor,
		cancel:    nil,
		DefaultSwapTaskHalper: &DefaultSwapTaskHalper{
			TaskType: taskType,
		},
	}
}

type DefaultSwapTaskHalper struct {
	TaskType v1.TaskType
}

func (t *DefaultSwapTask) Stop() error {
	t.cancel()
	return nil
}

func (t *DefaultSwapTask) Type() v1.TaskType {
	return t.taskType
}

func (t *DefaultSwapTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

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

	s, err := profile.GetNetworkSettings(ctx, p.Network)
	if err != nil {
		return nil, err
	}

	client, err := uniclient.NewSwapper(p.Network, s.BaseConfig(), t.taskType)
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

	if p.GetTx().GetTxCompleted() {
		task.Status = v1.ProcessStatus_StatusDone
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	return task, nil
}

func (h *DefaultSwapTaskHalper) Execute(ctx context.Context, profile *halp.Profile, p *v1.DefaultSwap, client defi.Swapper, estimation *v1.EstimationTx) (*bozdo.DefaultRes, *bozdo.Gas, error) {

	s, err := profile.GetNetworkSettings(ctx, p.Network)
	if err != nil {
		return nil, nil, errors.Wrap(err, "profile.GetNetworkSettings")
	}
	if client == nil {
		client, err = uniclient.NewSwapper(p.Network, s.BaseConfig(), h.TaskType)
		if err != nil {
			return nil, nil, errors.Wrap(err, "uniclient.NewSwapper")
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
		return nil, nil, errors.Wrap(err, "defi.ResolveAmount")
	}

	if am == nil || am.Cmp(big.NewInt(0)) == 0 {
		return nil, nil, &defi.ErrOutOfGas{N: p.Network, Token: p.FromToken}
	}

	estimateOnly := estimation == nil
	var Gas *bozdo.Gas
	if estimateOnly {
		am = bozdo.Percent(am, 90)
		Gas = nil
	} else {
		gas, err := GasManager(estimation, s.Source, p.Network, h.TaskType)
		if err != nil {
			return nil, nil, errors.Wrap(err, "GasManager")
		}
		Gas = gas

		balanceNative, err := client.GetBalance(ctx, &defi.GetBalanceReq{
			WalletAddress: profile.Addr,
			Token:         client.GetNetworkToken(),
		})
		if err != nil {
			return nil, nil, errors.Wrap(err, "client.GetBalance")
		}
		if balanceNative.WEI.Cmp(&Gas.TotalGas) <= 0 {
			return nil, nil, ErrProfileHasInsufficientBalance(p.Network, v1.Token_ETH, &Gas.TotalGas, balance.WEI)
		}
	}

	res, err := client.Swap(ctx, &defi.DefaultSwapReq{
		Network:      p.Network,
		Amount:       am,
		FromToken:    p.FromToken,
		ToToken:      p.ToToken,
		WalletPK:     profile.WalletPK,
		EstimateOnly: estimateOnly,
		Gas:          Gas,
		Debug:        true,
		Slippage:     getSlippage(s.Source, h.TaskType),
		SubType:      profile.SubType,
		ExchangeRate: pub.GetExchangeRate(p.FromToken, p.ToToken),
	}, h.TaskType)
	if err != nil {
		return nil, nil, errors.Wrap(err, "client.Swap")
	}
	return res, Gas, nil
}

func (h *DefaultSwapTaskHalper) EstimateCost(ctx context.Context, profile *halp.Profile, p *v1.DefaultSwap, client defi.Swapper) (*v1.EstimationTx, error) {
	res, _, err := h.Execute(ctx, profile, p, client, nil)
	if err != nil {
		return nil, errors.Wrap(err, "task.Execute")
	}

	return GasStation(res.ECost, p.Network), nil
}
