package task

import (
	"context"
	"math/big"
	"time"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/starknet"
	"github.com/hardstylez72/cry/internal/exchange/pub"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/pkg/errors"
)

func NewFibrousSwapTask() *StarkNetSwap {
	return NewStarkNetSwapTask(v1.TaskType_FibrousSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_FibrousSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_FibrousSwapTask) call an ambulance!")
		}
		return l.FibrousSwapTask, nil
	})
}

func NewAvnuSwapTask() *StarkNetSwap {
	return NewStarkNetSwapTask(v1.TaskType_AvnuSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_AvnuSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_AvnuSwapTask) call an ambulance!")
		}
		return l.AvnuSwapTask, nil
	})
}

func NewSithSwapTask() *StarkNetSwap {
	return NewStarkNetSwapTask(v1.TaskType_SithSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_SithSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_SithSwapTask) call an ambulance!")
		}
		return l.SithSwapTask, nil
	})
}

func NewJediSwapTask() *StarkNetSwap {
	return NewStarkNetSwapTask(v1.TaskType_JediSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_JediSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_JediSwapTask) call an ambulance!")
		}
		return l.JediSwapTask, nil
	})
}

func NewMySwapSwapTask() *StarkNetSwap {
	return NewStarkNetSwapTask(v1.TaskType_MySwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_MySwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_MySwapTask) call an ambulance!")
		}
		return l.MySwapTask, nil
	})
}

func NewProtossSwapTask() *StarkNetSwap {
	return NewStarkNetSwapTask(v1.TaskType_ProtossSwap, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_ProtosSwapTask)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_ProtosSwapTask) call an ambulance!")
		}
		return l.ProtosSwapTask, nil
	})
}

func NewSwap10kSwapTask() *StarkNetSwap {
	return NewStarkNetSwapTask(v1.TaskType_Swap10k, func(a *Input) (*v1.DefaultSwap, error) {
		l, ok := a.Task.Task.Task.(*v1.Task_Swap10K)
		if !ok {
			return nil, errors.New("Task.(*v1.Task_Swap10K) call an ambulance!")
		}
		return l.Swap10K, nil
	})
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
	if profile.Type != v1.ProfileType_StarkNet {
		return nil, errors.New("task is available only for starknet account")
	}

	client, err := NewStarkNetClient(profile)
	if err != nil {
		return nil, err
	}

	if p.GetApproveTx().GetTxId() == "" {
		txId, err := StarkNetApprove(taskContext, p.FromToken, client, profile, t.taskType)
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

func starkNetTxReady(tx *v1.TaskTx, seconds int64) bool {
	if tx == nil {
		return true
	}

	if tx.GetTs() == 0 {
		return true
	}

	txTime := tx.GetCompleteTs()

	now := time.Now().Unix()

	if now-txTime >= seconds {
		return true
	}

	return false
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

	approve, err := StarkNetApprove(ctx, p.FromToken, client, profile, h.TaskType)
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
			Token:         client.NativeToken,
		})
		if err != nil {
			return nil, nil, errors.Wrap(err, "client.GetFundingBalance")
		}
		if balanceNative.WEI.Cmp(&Gas.TotalGas) <= 0 {
			return nil, nil, ErrProfileHasInsufficientBalance(v1.Token_ETH, &Gas.TotalGas, balanceNative.WEI)
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
		Slippage:     getSlippage(s.Source, h.TaskType),
		Debug:        false,
		SubType:      profile.SubType,
		ExchangeRate: pub.GetExchangeRate(p.FromToken, p.ToToken),
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

func StarkNetApprove(ctx context.Context, token v1.Token, client *starknet.Client, profile *halp.Profile, taskType v1.TaskType) (*string, error) {

	balance, err := client.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: profile.Addr,
		Token:         token,
	})
	if err != nil {
		return nil, errors.Wrap(err, "client.GetFundingBalance")
	}

	spender := ""
	switch taskType {
	case v1.TaskType_SithSwap:
		spender = "0x028c858a586fa12123a1ccb337a0a3b369281f91ea00544d0c086524b759f627"
	case v1.TaskType_Swap10k:
		spender = "0x07a6f98c03379b9513ca84cca1373ff452a7462a3b61598f0af5bb27ad7f76d1"
	case v1.TaskType_JediSwap:
		spender = "0x041fd22b238fa21cfcf5dd45a8548974d8263b3a531a60388411c5e230f97023"
	case v1.TaskType_MySwap:
		spender = "0x010884171baf1914edc28d7afb619b40a4051cfae78a094a55d230f19e944a28"
	case v1.TaskType_ProtossSwap:
		spender = "0x07a0922657e550ba1ef76531454cb6d203d4d168153a0f05671492982c2f7741"
	case v1.TaskType_Dmail:
		spender = "0x0454f0bd015e730e5adbb4f080b075fdbf55654ff41ee336203aa2e1ac4d4309"
	case v1.TaskType_StarkNetIdMint:
		spender = "0x05dbdedc203e92749e2e746e2d40a768d966bd243df04a6b712e222bc040a9af"
	case v1.TaskType_AvnuSwap:
		spender = "0x04270219d365d6b017231b52e92b3fb5d7c8378b05e9abc97724537a80e93b0f"
	case v1.TaskType_FibrousSwap:
		spender = "0x01b23ed400b210766111ba5b1e63e33922c6ba0c45e6ad56ce112e5f4c578e62"
	default:
		return nil, errors.New("StarkNetApprove. unknown task type: " + taskType.String())
	}

	res, err := client.Approve(ctx, &starknet.ApproveReq{
		Token:       token,
		PK:          profile.WalletPK,
		Amount:      balance.WEI,
		SpenderAddr: spender,
		SubType:     profile.SubType,
	})
	if err != nil {
		return nil, errors.Wrap(err, "client.Approve")
	}

	return res.TxId, nil
}
