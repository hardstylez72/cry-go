package task

import (
	"context"
	"math/big"
	"time"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/pkg/errors"
)

func GasStation(ecost *bozdo.EstimatedGasCost, network v1.Network) *v1.EstimationTx {
	result := &v1.EstimationTx{}
	switch ecost.Type {
	case bozdo.TxTypeStarkNet:
		result.GasLimit = defi.AmountUni(ecost.TotalGasWei, network)
		result.GasPrice = defi.AmountUni(big.NewInt(1), network)
		result.Gas = defi.AmountUni(ecost.TotalGasWei, network)

		result.GasValuePercent = ""
		result.Name = ecost.Name
		result.Details = CastDetails(ecost.Details)
		return result
	default:
		result.GasLimit = defi.AmountUni(ecost.GasLimit, network)
		result.GasPrice = defi.AmountUni(ecost.GasPrice, network)
		totalFee := ecost.TotalGasWei
		if ecost.ExtraFee != nil {
			totalFee = new(big.Int).Add(totalFee, ecost.ExtraFee)
		}
		result.Gas = defi.AmountUni(totalFee, network)

		result.GasValuePercent = ""
		result.Name = ecost.Name
		result.Details = CastDetails(ecost.Details)
		return result
	}
}

func ResolveNetworkTokenAmount(balance, gas, value *big.Int) *big.Int {
	gas = bozdo.BigIntSum(bozdo.Percent(gas, 10), gas)
	need := new(big.Int).Add(gas, value)
	if need.Cmp(balance) <= 0 {
		return value
	}

	return new(big.Int).Sub(balance, gas)
}

func WaitTxComplete(ctx context.Context, ptx *v1.TaskTx, task *v1.ProcessTask, networker defi.Networker, updater TaskUpdater) error {
	if ptx != nil && ptx.TxId != "" {
		tx := ptx
		if err := networker.WaitTxComplete(ctx, tx.TxId); err != nil {
			if errors.Is(err, defi.ErrTxStatusFailed) {
				tx.TxCompleted = false
				tx.RetryCount = 0
				tx.TxId = ""
				if err := updater.UpdateTask(ctx, task); err != nil {
					return err
				}
			}
			if errors.Is(err, defi.ErrTxNotFound) {
				if tx.RetryCount > 5 {
					tx.TxCompleted = false
					tx.RetryCount = 0
					tx.TxId = ""
					if err := updater.UpdateTask(ctx, task); err != nil {
						return err
					}
					return errors.New("transaction not found")
				}
				tx.RetryCount++
				if err := updater.UpdateTask(ctx, task); err != nil {
					return err
				}
				// бывает сеть перегружена, чтобы не спамить
				return ErrTransactionIsNotReady
			}
			return err
		}

		if tx.CompleteTs == 0 {
			tx.CompleteTs = time.Now().Unix()
		}

		if err := updater.UpdateTask(ctx, task); err != nil {
			return err
		}

		if tx.GetNetwork() == v1.Network_StarkNet {
			if !starkNetTxReady(tx, 60*3) {
				return ErrTransactionIsNotReady
			}
		}

		tx.TxCompleted = true
		if err := updater.UpdateTask(ctx, task); err != nil {
			return err
		}
	}
	return nil
}

func GasMultiplier(mul *float64, gas *bozdo.Gas) *bozdo.Gas {

	if mul == nil || *mul == 0 || *mul == 1 {
		gas.GasMultiplier = 1
		return gas
	} else {
		gas.GasMultiplier = *mul
	}

	f := new(big.Float).SetFloat64(*mul)

	newGasLimit := new(big.Float).Mul(big.NewFloat(0).SetInt64(gas.GasLimit.Int64()), f)
	gl, _ := newGasLimit.Int(nil)
	gas.GasLimit = *gl

	total := big.NewInt(0).Mul(&gas.GasLimit, &gas.GasPrice)
	gas.TotalGas = *total

	return gas
}

func (a *Input) AddTx2(ctx context.Context, transactions ...*v1.TaskTx) error {

	for i := range transactions {
		tx := transactions[i]
		if tx == nil {
			continue
		}
		txDb := &repository.Transaction{
			Id:        tx.GetTxId(),
			Code:      tx.GetCode(),
			Network:   tx.GetNetwork().String(),
			Url:       tx.GetUrl(),
			TaskId:    a.Task.Id,
			ProfileId: a.ProfileId,
			ProcessId: a.ProcessId,
			UserId:    a.UserId,
			CreatedAt: time.Now(),
		}
		if err := TransactionAdd(ctx, a.ProcessRepository, txDb); err != nil {
			return err
		}
	}

	return nil
}

func (a *Input) AddTx(ctx context.Context, transactions ...*bozdo.Transaction) error {

	for i := range transactions {
		tx := transactions[i]
		if tx == nil {
			continue
		}
		txDb := &repository.Transaction{
			Id:        tx.Id,
			Code:      tx.Code,
			Network:   tx.Network.String(),
			Url:       tx.Url,
			TaskId:    a.Task.Id,
			ProfileId: a.ProfileId,
			ProcessId: a.ProcessId,
			UserId:    a.UserId,
			CreatedAt: time.Now(),
		}
		if err := TransactionAdd(ctx, a.ProcessRepository, txDb); err != nil {
			return err
		}
	}

	return nil
}

func TransactionAdd(ctx context.Context, rep repository.TransactionRepository, tx *repository.Transaction) error {
	return rep.TransactionAdd(ctx, tx)
}

func NewStarkNetApproveTx(id string) *v1.TaskTx {

	n := v1.Network_StarkNet
	code := bozdo.CodeApprove
	url := "https://starkscan.co/tx/" + id

	txx := &v1.TaskTx{
		TxCompleted: false,
		TxId:        id,
		RetryCount:  0,
		Url:         &url,
		Network:     &n,
		Code:        &code,
		Details:     CastDetails(nil),
		Ts:          time.Now().Unix(),
	}

	return txx
}

func BalanceSnapshotBefore(ctx context.Context, client defi.Networker, token v1.Token, profile *halp.Profile) ([]bozdo.TxDetail, error) {

	result := make([]bozdo.TxDetail, 0)

	if client.GetNetworkToken().String() != token.String() {
		bToken, err := client.GetBalance(ctx, &defi.GetBalanceReq{
			WalletAddress: profile.Addr,
			Token:         token,
		})
		if err != nil {
			return nil, err
		}

		result = append(result, bozdo.NewTokenBalanceBefore(bToken.WEI, client.Network(), token))
	}

	bToken, err := client.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: profile.Addr,
		Token:         client.GetNetworkToken(),
	})
	if err != nil {
		return nil, err
	}

	result = append(result, bozdo.NewNativeBalanceBefore(bToken.WEI, client.Network(), client.GetNetworkToken()))

	return result, nil
}

func BalanceSnapshotAfter(ctx context.Context, client defi.Networker, token v1.Token, profile *halp.Profile) ([]bozdo.TxDetail, error) {

	result := make([]bozdo.TxDetail, 0)

	if client.GetNetworkToken().String() != token.String() {
		bToken, err := client.GetBalance(ctx, &defi.GetBalanceReq{
			WalletAddress: profile.Addr,
			Token:         token,
		})
		if err != nil {
			return nil, err
		}

		result = append(result, bozdo.NewTokenBalanceAfter(bToken.WEI, client.Network(), token))
	}

	bToken, err := client.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: profile.Addr,
		Token:         client.GetNetworkToken(),
	})
	if err != nil {
		return nil, err
	}

	result = append(result, bozdo.NewNativeBalanceAfter(bToken.WEI, client.Network(), client.GetNetworkToken()))

	return result, nil
}

func NewTx(tx *bozdo.Transaction, gas *bozdo.Gas) *v1.TaskTx {

	if tx == nil {
		return nil
	}
	txx := &v1.TaskTx{
		TxCompleted: false,
		TxId:        tx.Id,
		RetryCount:  0,
		Url:         &tx.Url,
		Network:     &tx.Network,
		Code:        &tx.Code,
		Details:     CastDetails(tx.Details),
		Ts:          time.Now().Unix(),
		CompleteTs:  0,
	}

	if gas != nil {
		//txx.GasEstimated = defi.AmountUni(&gas.GasBeforeMultiplier, gas.Network)
		txx.GasResult = defi.AmountUni(&gas.TotalGas, gas.Network)
		//txx.GasLimit = defi.AmountUni(&gas.MaxGas, gas.Network)
		m := float32(gas.GasMultiplier)
		txx.Multiplier = &m
	}
	return txx
}

func CastDetails(in []bozdo.TxDetail) []*v1.TxDetail {
	out := make([]*v1.TxDetail, len(in))
	if in == nil || len(in) == 0 {
		return out
	}
	for i := range in {
		out[i] = &v1.TxDetail{
			Key:   in[i].Key,
			Value: in[i].Value,
		}
	}

	return out
}

func getSlippage(s *v1.NetworkSettings, taskType v1.TaskType) defi.SlippagePercent {
	if s == nil || s.TaskSettings == nil {
		return getSlippageFromConst(taskType)
	}

	v, ok := s.TaskSettings[taskType.String()]
	if !ok {
		return getSlippageFromConst(taskType)
	}
	return v.GetSlippage()
}

func getSlippageFromConst(taskType v1.TaskType) defi.SlippagePercent {
	v, ok := defi.SlippageMap[taskType]
	if !ok {
		return defi.SlippagePercent05
	}
	return v
}
