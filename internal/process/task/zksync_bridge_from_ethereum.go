package task

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/zksyncera"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

type ZksyncOfficialBridgeFromEthereumTask struct {
	cancel func()
}

func (t *ZksyncOfficialBridgeFromEthereumTask) Stop() error {
	t.cancel()
	return nil
}

func (t *ZksyncOfficialBridgeFromEthereumTask) Type() v1.TaskType {
	return v1.TaskType_ZkSyncOfficialBridgeFromEthereum
}

func (t *ZksyncOfficialBridgeFromEthereumTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_ZkSyncOfficialBridgeFromEthereumTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_ZkSyncOfficialBridgeFromEthereumTask) call an ambulance!")
	}

	p := l.ZkSyncOfficialBridgeFromEthereumTask

	profile, err := a.Halper.Profile(ctx, a.ProfileId)
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

	zksyncClient, _, err := NewZkSyncClient(profile, v1.Network_ZKSYNCERA)
	if err != nil {
		return nil, err
	}

	ethClient, err := NewEthClient(taskContext, a.Halper, a.ProfileId, v1.Network_Etherium)
	if err != nil {
		return nil, err
	}

	if p.GetTx().GetTxId() == "" {

		tx, err := t.Withdrawal(taskContext, a, zksyncClient, profile, ethClient)
		if err != nil {
			return nil, err
		}

		p.Tx = NewTx(tx, nil)
		if err := a.AddTx2(ctx, p.Tx); err != nil {
			return nil, err
		}
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	if err := WaitTxComplete(taskContext, p.Tx, task, ethClient, a); err != nil {
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

func (t *ZksyncOfficialBridgeFromEthereumTask) Withdrawal(ctx context.Context, a *Input, client *zksyncera.Client, profile *halp.Profile, ethClient defi.Networker) (*bozdo.Transaction, error) {

	s, err := profile.GetNetworkSettings(ctx, v1.Network_Etherium)
	if err != nil {
		return nil, err
	}

	l, ok := a.Task.Task.Task.(*v1.Task_ZkSyncOfficialBridgeFromEthereumTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_ZkSyncOfficialBridgeFromEthereumTask) call an ambulance!")
	}

	p := l.ZkSyncOfficialBridgeFromEthereumTask

	b, err := ethClient.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: profile.Addr,
		Token:         v1.Token_ETH,
	})
	if err != nil {
		return nil, err
	}

	if b.Float == 0 {
		return nil, errors.New("not enough balance in ETH")
	}

	am, err := defi.ResolveAmount(p.Amount, b.WEI)
	if err != nil {
		return nil, err
	}

	estimate, err := EstimateZkSyncOfficialBridgeFromEthSwapCost(ctx, profile, p)
	if err != nil {
		return nil, err
	}

	gas, err := GasManager(estimate, s.Source, v1.Network_Etherium, t.Type())
	if err != nil {
		return nil, err
	}

	am = ResolveNetworkTokenAmount(b.WEI, &gas.TotalGas, am)

	res, err := client.BridgeFromEthereumNetwork(ctx, &zksyncera.L1L2BridgeReq{
		Amount:   am,
		WalletPK: profile.WalletPK,
	})
	if err != nil {
		return nil, err
	}

	result := res.TxHash

	return result, nil
}

// 2503241972553024
// 3890138828593930
// 6643704998402254
func EstimateZkSyncOfficialBridgeFromEthSwapCost(ctx context.Context, profile *halp.Profile, p *v1.ZkSyncOfficialBridgeFromEthereumTask) (*v1.EstimationTx, error) {

	network := v1.Network_ZKSYNCERA

	s, err := profile.GetNetworkSettings(ctx, network)
	if err != nil {
		return nil, err
	}

	client, err := uniclient.NewZkSyncOfficialBridge(network, s.BaseConfig())
	if err != nil {
		return nil, err
	}

	b, err := client.GetBalance(ctx, &defi.GetBalanceReq{WalletAddress: profile.Addr, Token: v1.Token_ETH})
	if err != nil {
		return nil, err
	}

	am, err := defi.ResolveAmount(p.Amount, b.WEI)
	if err != nil {
		return nil, err
	}

	res, err := client.BridgeFromEthereumNetwork(ctx, &zksyncera.L1L2BridgeReq{
		Amount:       am,
		WalletPK:     profile.WalletPK,
		EstimateOnly: true,
		Gas:          nil,
	})
	if err != nil {
		return nil, err
	}

	return &v1.EstimationTx{
		Balance:  defi.AmountUni(b.WEI, v1.Network_Etherium),
		Value:    defi.AmountUni(am, v1.Network_Etherium),
		Gas:      defi.AmountUni(res.EstimatedGasCost.TotalGasWei, v1.Network_Etherium),
		GasLimit: defi.AmountUni(res.EstimatedGasCost.GasLimit, v1.Network_Etherium),
		GasPrice: defi.AmountUni(res.EstimatedGasCost.GasPrice, v1.Network_Etherium),
	}, nil
}

func NewEthClient(ctx context.Context, halper *halp.Halp, profileId string, n v1.Network) (defi.Networker, error) {

	profile, err := halper.Profile(ctx, profileId)
	if err != nil {
		return nil, err
	}

	s, err := profile.GetNetworkSettings(ctx, n)
	if err != nil {
		return nil, err
	}

	swapper, err := uniclient.NewBaseClient(n, s.BaseConfig())
	if err != nil {
		return nil, err
	}

	return swapper, nil
}
