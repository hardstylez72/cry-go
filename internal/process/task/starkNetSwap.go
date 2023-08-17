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

type SithSwapTask struct {
	*StarkNetSwap
}

type Swap10kTask struct {
	*StarkNetSwap
}

type JediTask struct {
	*StarkNetSwap
}

type StarkNetSwap struct {
	taskType  v1.TaskType
	extractor func(a *Input) (*v1.DefaultSwap, error)
	cancel    func()
	*StarketSwapHalper
}

func NewStarkNetSwapTask(taskType v1.TaskType, extractor func(a *Input) (*v1.DefaultSwap, error)) *StarkNetSwap {
	return &StarkNetSwap{
		taskType:  taskType,
		extractor: extractor,
		cancel:    nil,
		StarketSwapHalper: &StarketSwapHalper{
			TaskType: taskType,
		},
	}
}

type StarketSwapHalper struct {
	TaskType v1.TaskType
}

func (t *StarkNetSwap) Stop() error {
	t.cancel()
	return nil
}

func (t *StarkNetSwap) Type() v1.TaskType {
	return t.taskType
}

func (t *StarkNetSwap) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

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
	if profile.Type != v1.ProfileType_StarkNet {
		return nil, errors.New("task is available only for starknet account")
	}

	client, err := NewStarkNetClient(profile)
	if err != nil {
		return nil, err
	}

	if p.GetApproveTx().GetTxId() == "" {
		txId, err := StarkNetApprove(taskContext, p, client, profile, t.taskType)
		if err != nil {
			return nil, err
		}

		if txId != nil {
			p.ApproveTx = NewStarkNetApproveTx(*txId)
			if err := a.AddTx2(ctx, p.ApproveTx); err != nil {
				return nil, err
			}
		} else {
			p.ApproveTx = &v1.TaskTx{TxId: AlreadyApproved, TxCompleted: true}
		}

		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	if err := WaitTxComplete(taskContext, p.ApproveTx, task, client, a); err != nil {
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

func (h *StarketSwapHalper) Execute(ctx context.Context, profile *halp.Profile, p *v1.DefaultSwap, client *starknet.Client, estimation *v1.EstimationTx) (*bozdo.DefaultRes, *bozdo.Gas, error) {

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

	approve, err := StarkNetApprove(ctx, p, client, profile, h.TaskType)
	if err != nil {
		return nil, nil, err
	}

	if approve != nil {
		if err := client.WaitTxComplete(ctx, *approve); err != nil {
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
		am = bozdo.Percent(am, 90)
		Gas = nil
	} else {
		gas, err := GasManager(estimation, s.Source, p.Network)
		if err != nil {
			return nil, nil, err
		}
		Gas = gas

		balanceNative, err := client.GetBalance(ctx, &defi.GetBalanceReq{
			WalletAddress: profile.Addr,
			Token:         client.NativeToken,
		})
		if err != nil {
			return nil, nil, errors.Wrap(err, "client.GetBalance")
		}
		if balanceNative.WEI.Cmp(&Gas.TotalGas) <= 0 {
			return nil, nil, ErrProfileHasInsufficientBalance(v1.Token_ETH, &Gas.TotalGas, balance.WEI)
		}
	}

	res, err := client.Swap(ctx, &defi.DefaultSwapReq{
		Network:      v1.Network_StarkNet,
		Amount:       am,
		FromToken:    p.FromToken,
		ToToken:      p.ToToken,
		WalletPK:     profile.WalletPK,
		EstimateOnly: estimateOnly,
		Gas:          Gas,
		Debug:        false,
		Slippage:     getSlippage(s.Source, h.TaskType),
	}, h.TaskType)
	if err != nil {
		return nil, nil, err
	}
	return res, Gas, nil
}

func (h *StarketSwapHalper) EstimateCost(ctx context.Context, profile *halp.Profile, p *v1.DefaultSwap, client *starknet.Client) (*v1.EstimationTx, error) {
	res, _, err := h.Execute(ctx, profile, p, client, nil)
	if err != nil {
		return nil, err
	}

	return GasStation(res.ECost, p.Network), nil
}

const AlreadyApproved = "token already approved"

func StarkNetApprove(ctx context.Context, p *v1.DefaultSwap, client *starknet.Client, profile *halp.Profile, taskType v1.TaskType) (*string, error) {

	balance, err := client.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: profile.Addr,
		Token:         p.FromToken,
	})
	if err != nil {
		return nil, errors.Wrap(err, "client.GetBalance")
	}

	spender := ""
	switch taskType {
	case v1.TaskType_SithSwap:
		spender = client.SithSwapRouterAddress()
	case v1.TaskType_Swap10k:
		spender = client.Swap10kRouterAddress()
	case v1.TaskType_JediSwap:
		spender = client.JediSwapRouterAddress()
	default:
		return nil, errors.New("unknown task type: " + taskType.String())
	}

	res, err := client.Approve(ctx, &starknet.ApproveReq{
		Token:       p.FromToken,
		PK:          profile.WalletPK,
		Amount:      balance.WEI,
		SpenderAddr: spender,
	})
	if err != nil {
		return nil, errors.Wrap(err, "client.Approve")
	}

	return res.TxId, nil
}
