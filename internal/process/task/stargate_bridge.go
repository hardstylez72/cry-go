package task

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/log"
	"github.com/hardstylez72/cry/internal/lzscan"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type StargateTask struct {
	cancel func()
}

func (t *StargateTask) Stop() error {
	t.cancel()
	return nil
}

func (t *StargateTask) Type() v1.TaskType {
	return v1.TaskType_StargateBridge
}

func (t *StargateTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_StargateBridgeTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_StargateBridgeTask) call an ambulance!")
	}

	p := l.StargateBridgeTask

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

	client, err := NewSwapper(taskContext, a)
	if err != nil {
		return nil, err
	}

	if p.GetTx().GetTxId() == "" {

		estimation, err := EstimateStargateBridgeSwapCost(taskContext, p, profile)
		if err != nil {
			return nil, err
		}

		res, gas, err := t.Swap(taskContext, p, client, profile, estimation)
		if err != nil {
			return nil, err
		}

		if err := a.AddTx(ctx, res.Allowance); err != nil {
			return nil, err
		}
		p.Tx = NewTx(res.Swap, gas)
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

	if p.LayerZeroStatus == nil || *p.LayerZeroStatus != lzscan.StatusDELIVERED {
		s := lzscan.NewService()
		lzUrl, err := s.GetTxUrl(taskContext, p.GetTx().GetTxId())
		if err != nil {
			if err == lzscan.ErrNotFound {
				return task, nil
			}
			return nil, err
		}
		p.LzscanUrl = &lzUrl
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}

		status, err := s.WaitConfirm(taskContext, p.GetTx().GetTxId())
		if err != nil {
			return nil, err
		}

		p.LzscanUrl = &lzUrl
		p.LayerZeroStatus = &status
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	switch p.GetLayerZeroStatus() {
	case lzscan.StatusDELIVERED:
		task.Status = v1.ProcessStatus_StatusDone
		task.FinishedAt = timestamppb.Now()
	case lzscan.StatusINFLIGHT:
		task.Status = v1.ProcessStatus_StatusRunning
	case lzscan.StatusFAILED:
		task.Status = v1.ProcessStatus_StatusError
	}

	if err := a.UpdateTask(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

type SwapRes struct {
	Swap      *bozdo.Transaction
	Allowance *bozdo.Transaction
	Swapper   defi.StargateSwapper
}

func NewSwapper(ctx context.Context, a *Input) (defi.StargateSwapper, error) {

	l, ok := a.Task.Task.Task.(*v1.Task_StargateBridgeTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_StargateBridgeTask) call an ambulance!")
	}

	profile, err := a.ProfileRepository.GetProfile(ctx, a.ProfileId)
	if err != nil {
		return nil, err
	}

	proxyString := ""
	if profile.Proxy.Valid {
		proxyString = profile.Proxy.String
	}

	stgs, err := a.SettingsService.GetSettings(ctx, a.UserId, l.StargateBridgeTask.FromNetwork)
	if err != nil {
		return nil, err
	}

	swapper, err := uniclient.NewStargateSwapper(l.StargateBridgeTask.FromNetwork, &uniclient.BaseClientConfig{
		ProxyString: proxyString,
		RPCEndpoint: stgs.RpcEndpoint,
	})
	if err != nil {
		return nil, err
	}

	return swapper, nil
}

func (t *StargateTask) Swap(ctx context.Context, p *v1.StargateBridgeTask, swapper defi.StargateSwapper, profile *halp.Profile, estimation *v1.EstimationTx) (*SwapRes, *bozdo.Gas, error) {

	s, err := profile.GetNetworkSettings(ctx, p.FromNetwork)
	if err != nil {
		return nil, nil, err
	}
	b, err := swapper.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: profile.Addr,
		Token:         p.FromToken,
	})
	if err != nil {
		return nil, nil, err
	}

	am, err := defi.ResolveAmount(p.Amount, b.WEI)
	if err != nil {
		return nil, nil, err
	}

	gas, err := GasManager(estimation, s.Source, p.FromNetwork)
	if err != nil {
		return nil, nil, err
	}

	if p.FromToken == swapper.GetNetworkToken() {
		am = ResolveNetworkTokenAmount(b.WEI, &gas.TotalGas, am)
	}

	println("balance: " + defi.WEIToEther(b.WEI).String())
	println("amount: " + defi.WEIToEther(am).String())
	println("diff: " + defi.WEIToEther(new(big.Int).Sub(b.WEI, am)).String())

	swap, err := swapper.StargateBridgeSwap(ctx, &defi.DefaultBridgeReq{
		FromNetwork: p.FromNetwork,
		ToNetwork:   p.ToNetwork,
		WalletPK:    profile.WalletPK,
		Amount:      am,
		FromToken:   p.FromToken,
		ToToken:     p.ToToken,
		Gas:         gas,
		Slippage:    getSlippage(s.Source, v1.TaskType_StargateBridge),
	})
	if err != nil {
		return nil, nil, err
	}

	result := &SwapRes{
		Swapper:   swapper,
		Swap:      swap.Tx,
		Allowance: swap.ApproveTx,
	}

	return result, gas, nil
}

type Tx struct {
	Url  string
	TxId string
}

func GasManager(e *v1.EstimationTx, s *v1.NetworkSettings, n v1.Network) (*bozdo.Gas, error) {

	if e.GasLimit == nil {
		e.GasLimit = &v1.AmUni{}
	}
	limit, ok := big.NewInt(0).SetString(e.GasLimit.Wei, 10)
	if !ok {
		log.Log.Error("GasLimit: " + e.GasLimit.Wei)
		limit = big.NewInt(0)
	}

	if e.GasPrice == nil {
		e.GasPrice = &v1.AmUni{}
	}
	gasPrice, ok := big.NewInt(0).SetString(e.GasPrice.Wei, 10)
	if !ok {
		log.Log.Error("GasPrice: " + e.GasPrice.Wei)
		gasPrice = big.NewInt(0)
	}

	if e.Gas == nil {
		e.Gas = &v1.AmUni{}
	}
	totalGas, ok := big.NewInt(0).SetString(e.Gas.Wei, 10)
	if !ok {
		log.Log.Error("Gas total: " + e.Gas.Wei)
		totalGas = big.NewInt(0)
	}

	beforeMultiplier := big.NewInt(totalGas.Int64())

	maxGas := big.NewInt(0)
	tmp, ok := big.NewInt(0).SetString(s.GetMaxGas(), 10)
	if ok {
		maxGas = tmp
	}

	gas := &bozdo.Gas{
		Network:             n,
		MaxGas:              *maxGas,
		GasBeforeMultiplier: *beforeMultiplier,
		GasLimit:            *limit,
		GasPrice:            *gasPrice,
		TotalGas:            *totalGas,
	}

	gas = GasMultiplier(&s.GasMultiplier, gas)

	if maxGas.Cmp(&gas.TotalGas) <= -1 {
		max := defi.AmountUni(maxGas, n)
		estimated := defi.AmountUni(&gas.TotalGas, n)
		return nil, ErrGasIsOverMax(max.Usd, estimated.Usd)
	}

	return gas, nil
}

func EstimateStargateBridgeSwapCost(ctx context.Context, p *v1.StargateBridgeTask, profile *halp.Profile) (*v1.EstimationTx, error) {

	s, err := profile.GetNetworkSettings(ctx, p.FromNetwork)
	if err != nil {
		return nil, err
	}

	swapper, err := uniclient.NewStargateSwapper(p.FromNetwork, s.BaseConfig())
	if err != nil {
		return nil, err
	}

	b, err := swapper.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: profile.Addr,
		Token:         p.FromToken,
	})
	if err != nil {
		return nil, err
	}

	am, err := defi.ResolveAmount(p.Amount, b.WEI)
	if err != nil {
		return nil, err
	}

	swap, err := swapper.StargateBridgeSwap(ctx, &defi.DefaultBridgeReq{
		FromNetwork:  p.FromNetwork,
		ToNetwork:    p.ToNetwork,
		WalletPK:     profile.WalletPK,
		Amount:       bozdo.Percent(am, 90),
		FromToken:    p.FromToken,
		ToToken:      p.ToToken,
		Gas:          nil,
		EstimateOnly: true,
		Debug:        false,
		Slippage:     getSlippage(s.Source, v1.TaskType_StargateBridge),
	})
	if err != nil {
		return nil, err
	}

	if p.FromToken == swapper.GetNetworkToken() {
		am = ResolveNetworkTokenAmount(b.WEI, swap.ECost.TotalGasWei, am)
	}

	return GasStation(swap.ECost, p.FromNetwork), nil
}
