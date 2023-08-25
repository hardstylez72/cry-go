package task

import (
	"context"
	"strconv"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/repository/pg"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

func NewStarkNetBridgeTask() *DefaultLiquidityBridgeTask {
	return NewDefaultLiquidityBridgeTaskTask(v1.TaskType_StarkNetBridge, func(a *Input) (*v1.LiquidityBridgeTask, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_StarkNetBridgeTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_StarkNetBridgeTask) call an ambulance!")
		}
		return l.StarkNetBridgeTask, nil
	})
}

type DefaultLiquidityBridgeTask struct {
	taskType  v1.TaskType
	extractor func(a *Input) (*v1.LiquidityBridgeTask, error)
	cancel    func()
	*DefaultLiquidityBridgeTaskHalper
}

func NewDefaultLiquidityBridgeTaskTask(taskType v1.TaskType, extractor func(a *Input) (*v1.LiquidityBridgeTask, error)) *DefaultLiquidityBridgeTask {
	return &DefaultLiquidityBridgeTask{
		taskType:  taskType,
		extractor: extractor,
		cancel:    nil,
		DefaultLiquidityBridgeTaskHalper: &DefaultLiquidityBridgeTaskHalper{
			TaskType: taskType,
		},
	}
}

type DefaultLiquidityBridgeTaskHalper struct {
	TaskType v1.TaskType
}

func (t *DefaultLiquidityBridgeTask) Stop() error {
	t.cancel()
	return nil
}

func (t *DefaultLiquidityBridgeTask) Type() v1.TaskType {
	return t.taskType
}

func LiquidityBridgeProfiles(ctx context.Context, halper *halp.Halp, rep repository.ProfileRepository, p *v1.LiquidityBridgeTask, num int) (from *halp.Profile, to *halp.Profile, err error) {

	switch true {
	case p.FromNetwork == v1.Network_Etherium && p.ToNetwork == v1.Network_StarkNet:
		fromP, err := rep.GetProfileByNum(ctx, num, v1.ProfileType_EVM.String())
		if err != nil {
			if errors.Is(err, pg.EntityNotFound) {
				return nil, nil, errors.New("EMV profile with num :" + strconv.Itoa(num) + " not found")
			}
			return nil, nil, err
		}
		from, err = halper.Profile(ctx, fromP.Id)
		if err != nil {
			return nil, nil, err
		}
		toP, err := rep.GetProfileByNum(ctx, num, v1.ProfileType_StarkNet.String())
		if err != nil {
			if errors.Is(err, pg.EntityNotFound) {
				return nil, nil, errors.New("StarkNet profile with num :" + strconv.Itoa(num) + " not found")
			}
			return nil, nil, err
		}
		to, err = halper.Profile(ctx, toP.Id)
		if err != nil {
			return nil, nil, err
		}

		return from, to, nil
	default:
		return nil, nil, errors.New("unsupported pair: " + p.String())
	}
}

func (t *DefaultLiquidityBridgeTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

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

	from, to, err := LiquidityBridgeProfiles(ctx, a.Halper, a.ProfileRepository, p, profile.Num)
	if err != nil {
		return nil, err
	}

	s, err := from.GetNetworkSettings(ctx, p.FromNetwork)
	if err != nil {
		return nil, err
	}

	client, networker, err := uniclient.NewLiquidityBridge(p.FromNetwork, p.ToNetwork, s.BaseConfig(), t.taskType)
	if err != nil {
		return nil, err
	}

	if p.GetTx().GetTxId() == "" {

		estimation, err := t.EstimateCost(taskContext, from, to, p, client, networker)
		if err != nil {
			return nil, errors.Wrap(err, "EstimateSwapCost of "+t.taskType.String())
		}
		res, gas, err := t.Execute(taskContext, from, to, p, client, networker, estimation)
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

	if err := WaitTxComplete(taskContext, p.Tx, task, networker, a); err != nil {
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

func (h *DefaultLiquidityBridgeTaskHalper) Execute(ctx context.Context, from, to *halp.Profile, p *v1.LiquidityBridgeTask, client defi.LiquidityBridger, networker defi.Networker, estimation *v1.EstimationTx) (*bozdo.DefaultRes, *bozdo.Gas, error) {

	s, err := from.GetNetworkSettings(ctx, p.FromNetwork)
	if err != nil {
		return nil, nil, err
	}
	if client == nil {
		client, networker, err = uniclient.NewLiquidityBridge(p.FromNetwork, p.ToNetwork, s.BaseConfig(), h.TaskType)
		if err != nil {
			return nil, nil, err
		}
	}

	balance, err := networker.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: from.Addr,
		Token:         p.Token,
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "client.GetBalance")
	}

	percent, err := defi.GetPercent(p.Amount)
	if err != nil {
		return nil, nil, err
	}

	estimateOnly := estimation == nil
	var Gas *bozdo.Gas
	if estimateOnly {
		Gas = nil
	} else {
		gas, err := GasManager(estimation, s.Source, p.FromNetwork)
		if err != nil {
			return nil, nil, err
		}
		Gas = gas

		balanceNative, err := networker.GetBalance(ctx, &defi.GetBalanceReq{
			WalletAddress: from.Addr,
			Token:         networker.GetNetworkToken(),
		})
		if err != nil {
			return nil, nil, errors.Wrap(err, "client.GetBalance")
		}
		if balanceNative.WEI.Cmp(&Gas.TotalGas) <= 0 {
			return nil, nil, ErrProfileHasInsufficientBalance(networker.GetNetworkToken(), &Gas.TotalGas, balance.WEI)
		}
	}

	res, err := client.LiquidityBridge(ctx, &defi.LiquidityBridgeReq{
		From:         p.FromNetwork,
		To:           p.ToNetwork,
		Percent:      percent,
		Token:        p.Token,
		PkFrom:       from.WalletPK,
		PkTo:         to.WalletPK,
		Proxy:        s.BaseConfig().ProxyString,
		EstimateOnly: estimateOnly,
		FromSubType:  from.SubType,
		ToSubType:    to.SubType,
	}, h.TaskType)
	if err != nil {
		return nil, nil, err
	}
	return res, Gas, nil
}

func (h *DefaultLiquidityBridgeTaskHalper) EstimateCost(ctx context.Context, from, to *halp.Profile, p *v1.LiquidityBridgeTask, client defi.LiquidityBridger, networker defi.Networker) (*v1.EstimationTx, error) {
	res, _, err := h.Execute(ctx, from, to, p, client, networker, nil)
	if err != nil {
		return nil, err
	}

	return GasStation(res.ECost, p.FromNetwork), nil
}
