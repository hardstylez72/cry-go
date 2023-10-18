package task

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/zksyncera"
	"github.com/hardstylez72/cry/internal/exchange/pub"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/pkg/errors"
)

type SyncSwapTask struct {
	*ZkSyncSwap
}

// deprecated
func NewOdosSwapTaskZksync() *ZkSyncSwap {
	return NewZkSyncSwapTask(v1.TaskType_OdosSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_OdosSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_OdosSwapTask) call an ambulance!")
		}
		return l.OdosSwapTask, nil
	})
}

func NewSyncSwapTask() *ZkSyncSwap {
	return NewZkSyncSwapTask(v1.TaskType_SyncSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_SyncSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_SyncSwapTask) call an ambulance!")
		}
		return l.SyncSwapTask, nil
	})
}

func NewMuteioSwapTask() *ZkSyncSwap {
	return NewZkSyncSwapTask(v1.TaskType_MuteioSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_MuteioSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_MuteioSwapTask) call an ambulance!")
		}
		return l.MuteioSwapTask, nil
	})
}

func NewMaverickSwapTask() *ZkSyncSwap {
	return NewZkSyncSwapTask(v1.TaskType_MaverickSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_MaverickSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_MaverickSwapTask) call an ambulance!")
		}
		return l.MaverickSwapTask, nil
	})
}

func NewSpaceFiSwapTask() *ZkSyncSwap {
	return NewZkSyncSwapTask(v1.TaskType_SpaceFISwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_SpaceFiSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_SpaceFiSwapTask) call an ambulance!")
		}
		return l.SpaceFiSwapTask, nil
	})
}

func NewVelocoreSwapTask() *ZkSyncSwap {
	return NewZkSyncSwapTask(v1.TaskType_VelocoreSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_VelocoreSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_VelocoreSwapTask) call an ambulance!")
		}
		return l.VelocoreSwapTask, nil
	})
}

func NewIzumiSwapTask() *ZkSyncSwap {
	return NewZkSyncSwapTask(v1.TaskType_IzumiSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_IzumiSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_IzumiSwapTask) call an ambulance!")
		}
		return l.IzumiSwapTask, nil
	})
}

func NewVeSyncSwapTask() *ZkSyncSwap {
	return NewZkSyncSwapTask(v1.TaskType_VeSyncSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_VeSyncSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_VeSyncSwapTask) call an ambulance!")
		}
		return l.VeSyncSwapTask, nil
	})
}

func NewEzkaliburSwapTask() *ZkSyncSwap {
	return NewZkSyncSwapTask(v1.TaskType_EzkaliburSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_EzkaliburSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_EzkaliburSwapTask) call an ambulance!")
		}
		return l.EzkaliburSwapTask, nil
	})
}

func NewZkSwapTask() *ZkSyncSwap {
	return NewZkSyncSwapTask(v1.TaskType_ZkSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_ZkSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_ZkSwapTask) call an ambulance!")
		}
		return l.ZkSwapTask, nil
	})
}

// deprecated
func NewPancakeZkSyncSwapTask() *ZkSyncSwap {
	return NewZkSyncSwapTask(v1.TaskType_PancakeSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_PancakeSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_PancakeSwapTask) call an ambulance!")
		}
		return l.PancakeSwapTask, nil
	})
}

type ZkSyncSwap struct {
	taskType  v1.TaskType
	extractor func(a *Input) (*v1.DefaultSwap, error)
	cancel    func()
	*ZkSyncSwapHalper
}

func NewZkSyncSwapTask(taskType v1.TaskType, extractor func(a *Input) (*v1.DefaultSwap, error)) *ZkSyncSwap {
	return &ZkSyncSwap{
		taskType:  taskType,
		extractor: extractor,
		cancel:    nil,
		ZkSyncSwapHalper: &ZkSyncSwapHalper{
			TaskType: taskType,
		},
	}
}

type ZkSyncSwapHalper struct {
	TaskType v1.TaskType
}

func (t *ZkSyncSwap) Stop() error {
	t.cancel()
	return nil
}

func (t *ZkSyncSwap) Type() v1.TaskType {
	return t.taskType
}

func (t *ZkSyncSwap) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskStarkNetTimeout)
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

	client, _, err := NewZkSyncClient(profile, v1.Network_ZKSYNCERA)
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

func (h *ZkSyncSwapHalper) Execute(ctx context.Context, profile *halp.Profile, p *v1.DefaultSwap, client *zksyncera.Client, estimation *v1.EstimationTx) (*bozdo.DefaultRes, *bozdo.Gas, error) {

	s, err := profile.GetNetworkSettings(ctx, p.Network)
	if err != nil {
		return nil, nil, err
	}
	if client == nil {
		client, _, err = NewZkSyncClient(profile, v1.Network_ZKSYNCERA)
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
		gas, err := GasManager(estimation, s.Source, p.Network, h.TaskType)
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
			return nil, nil, ErrProfileHasInsufficientBalance(v1.Token_ETH, &Gas.TotalGas, balanceNative.WEI)
		}
	}

	res, err := client.Swap(ctx, &defi.DefaultSwapReq{
		Network:      v1.Network_ZKSYNCERA,
		Amount:       am,
		FromToken:    p.FromToken,
		ToToken:      p.ToToken,
		WalletPK:     profile.WalletPK,
		EstimateOnly: estimateOnly,
		Gas:          Gas,
		Debug:        false,
		Slippage:     getSlippage(s.Source, h.TaskType),
		SubType:      profile.SubType,
		ExchangeRate: pub.GetExchangeRate(p.FromToken, p.ToToken),
	}, h.TaskType)
	if err != nil {
		return nil, nil, err
	}
	return res, Gas, nil
}

func (h *ZkSyncSwapHalper) EstimateCost(ctx context.Context, profile *halp.Profile, p *v1.DefaultSwap, client *zksyncera.Client) (*v1.EstimationTx, error) {
	res, _, err := h.Execute(ctx, profile, p, client, nil)
	if err != nil {
		return nil, err
	}

	return GasStation(res.ECost, p.Network), nil
}
