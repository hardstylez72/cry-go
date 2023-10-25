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

func NewZkLendLPTask() *DefaultLPTask {
	return &DefaultLPTask{
		taskType: v1.TaskType_ZkLendLP,
		extractor: func(a *Input) (*v1.DefaultLP, error) {
			l, ok := a.Task.Task.Task.(*v1.Task_ZkLendLPTask)
			if !ok {
				return nil, errors.New("Task.(*v1.Task_ZkLandLPTask) call an ambulance!")
			}
			return l.ZkLendLPTask, nil
		},
		DefaultLPTaskHalper: &DefaultLPTaskHalper{
			TaskType: v1.TaskType_ZkLendLP,
		},
	}
}

func NewAaveLPTask() *DefaultLPTask {
	return &DefaultLPTask{
		taskType: v1.TaskType_AaveLP,
		extractor: func(a *Input) (*v1.DefaultLP, error) {
			l, ok := a.Task.Task.Task.(*v1.Task_AaveLPTask)
			if !ok {
				return nil, errors.New("Task.(*v1.Task_AaveLPTask) call an ambulance!")
			}
			return l.AaveLPTask, nil
		},
		DefaultLPTaskHalper: &DefaultLPTaskHalper{
			TaskType: v1.TaskType_AaveLP,
		},
	}
}

type DefaultLPTask struct {
	taskType  v1.TaskType
	extractor func(a *Input) (*v1.DefaultLP, error)
	cancel    func()
	*DefaultLPTaskHalper
}

type DefaultLPTaskHalper struct {
	TaskType v1.TaskType
}

func (t *DefaultLPTask) Stop() error {
	t.cancel()
	return nil
}

func (t *DefaultLPTask) Type() v1.TaskType {
	return t.taskType
}

func (t *DefaultLPTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

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

	s, err := profile.GetNetworkSettings(ctx, p.Network)
	if err != nil {
		return nil, err
	}
	client, err := uniclient.NewLPClient(p.Network, s.BaseConfig())
	if err != nil {
		return nil, err
	}

	if profile.Type == v1.ProfileType_StarkNet {

		client, err := NewStarkNetClient(profile)
		if err != nil {
			return nil, err
		}

		for i := range p.Tokens {
			token := p.Tokens[i]
			txId, err := StarkNetApprove(taskContext, token.Token, client, profile, t.taskType)
			if err != nil {
				return nil, err
			}

			if txId != nil {
				token.ApproveTx = NewStarkNetApproveTx(*txId)
				if err := a.AddTx2(ctx, token.ApproveTx); err != nil {
					return nil, err
				}
				if err := a.UpdateTask(ctx, task); err != nil {
					return nil, err
				}
			}
		}
	}

	if p.GetTx().GetTxId() == "" {

		estimation, err := t.EstimateLPCost(taskContext, profile, p, client)
		if err != nil {
			return nil, errors.Wrap(err, "EstimateLPCost")
		}
		res, gas, err := t.Execute(taskContext, profile, p, client, estimation)
		if err != nil {
			return nil, errors.Wrap(err, "Execute")
		}

		p.Tx = NewTx(res.Tx, gas)
		if err := a.AddTx2(ctx, p.Tx); err != nil {
			return nil, err
		}

		for i := range res.Approves {
			approve := &res.Approves[i]
			if err := a.AddTx(ctx, approve); err != nil {
				return nil, err
			}
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

func (h *DefaultLPTaskHalper) Execute(ctx context.Context, profile *halp.Profile, p *v1.DefaultLP, client defi.LP, estimation *v1.EstimationTx) (*defi.LPRes, *bozdo.Gas, error) {

	s, err := profile.GetNetworkSettings(ctx, p.Network)
	if err != nil {
		return nil, nil, err
	}
	if client == nil {
		client, err = uniclient.NewLPClient(p.Network, s.BaseConfig())
		if err != nil {
			return nil, nil, err
		}
	}

	var am = big.NewInt(0)

	if p.Add {

		if len(p.Tokens) == 0 {
			return nil, nil, errors.New("invalid params. tokens are empty")
		}
		tokenA := p.Tokens[0].Token

		balance, err := client.GetBalance(ctx, &defi.GetBalanceReq{
			WalletAddress: profile.Addr,
			Token:         tokenA,
		})
		if err != nil {
			return nil, nil, errors.Wrap(err, "client.GetFundingBalance")
		}
		am, err = defi.ResolveAmount(p.Amount, balance.WEI)
		if err != nil {
			return nil, nil, err
		}

		if am == nil || am.Cmp(big.NewInt(0)) == 0 {
			return nil, nil, errors.New("not enough balance of " + p.A.String())
		}
	} else {
		// todo:/...
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
	}

	res, err := client.LP(ctx, &defi.LPReq{
		Amount:       am,
		Tokens:       CastLPTokens(p.Tokens),
		PK:           profile.WalletPK,
		Add:          p.Add,
		EstimateOnly: estimateOnly,
		Gas:          Gas,
		PSubType:     profile.SubType,
		Debug:        true,
	}, h.TaskType)
	if err != nil {
		return nil, nil, err
	}

	return res, Gas, nil
}

func (h *DefaultLPTaskHalper) EstimateLPCost(ctx context.Context, profile *halp.Profile, p *v1.DefaultLP, client defi.LP) (*v1.EstimationTx, error) {
	res, _, err := h.Execute(ctx, profile, p, client, nil)
	if err != nil {
		return nil, err
	}

	return GasStation(res.ECost, p.Network), nil
}

func CastLPTokens(in []*v1.LPToken) []v1.Token {
	out := make([]v1.Token, 0)

	for _, el := range in {
		out = append(out, el.Token)
	}
	return out
}
