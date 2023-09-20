package task

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/nft/merkly"
	"github.com/hardstylez72/cry/internal/defi/zksyncera"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

type MerklyMintAndBridgeNFTTask struct {
	cancel func()
}

func (t *MerklyMintAndBridgeNFTTask) Stop() error {
	t.cancel()
	return nil
}

func (t *MerklyMintAndBridgeNFTTask) Type() v1.TaskType {
	return v1.TaskType_MerklyMintAndBridgeNFT
}

func (t *MerklyMintAndBridgeNFTTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_MerklyMintAndBridgeNFTTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_MerklyMintAndBridgeTask) call an ambulance!")
	}

	p := l.MerklyMintAndBridgeNFTTask

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

	s, err := profile.GetNetworkSettings(ctx, p.FromNetwork)
	if err != nil {
		return nil, err
	}

	client, err := uniclient.NewMerklyMintAndBridgeNFT(p.FromNetwork, s.BaseConfig())
	if err != nil {
		return nil, err
	}

	// mint
	if p.GetMintTx().GetTxId() == "" {

		estimation, err := EstimateMerklyMintCost(taskContext, profile, p, client)
		if err != nil {
			return nil, errors.Wrap(err, "EstimateMerklyMintCost")
		}
		res, gas, err := MerklyMintNFT(taskContext, profile, p, client, estimation)
		if err != nil {
			return nil, errors.Wrap(err, "MerklyMintNFT")
		}

		p.MintTx = NewTx(res.Tx, gas)

		if err := a.AddTx(ctx, res.ApproveTx); err != nil {
			return nil, err
		}

		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	// mint wait
	if !p.GetMintTx().GetTxCompleted() {
		if err := WaitTxComplete(taskContext, p.GetMintTx(), task, client, a); err != nil {
			return nil, err
		}
		if err := a.AddTx2(ctx, p.GetMintTx()); err != nil {
			return nil, err
		}
	}

	nftId, err := client.GetMerklyNFTId(ctx, common.HexToHash(p.GetMintTx().GetTxId()))
	if err != nil {
		tx := p.MintTx
		if tx.RetryCount > 20 {
			tx.TxCompleted = false
			tx.RetryCount = 0
			tx.TxId = ""
			if err := a.UpdateTask(ctx, task); err != nil {
				return nil, err
			}
		} else {
			tx.RetryCount++
			if err := a.UpdateTask(ctx, task); err != nil {
				return nil, err
			}
			return a.Task, nil
		}

		return nil, err
	}
	tmp := nftId.String()
	p.NftId = &tmp

	// bridge
	if p.GetBridgeTx().GetTxId() == "" {
		estimation, err := EstimateMerklyBridgeCost(taskContext, profile, p, client)
		if err != nil {
			return nil, errors.Wrap(err, "EstimateMerklyBridgeCost")
		}
		res, gas, err := MerklyBridgeNFT(taskContext, profile, p, client, estimation)
		if err != nil {
			return nil, errors.Wrap(err, "MerklyBridgeNFT")
		}

		p.BridgeTx = NewTx(res.Tx, gas)

		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
		if err := a.AddTx(ctx, res.ApproveTx); err != nil {
			return nil, err
		}
	}

	// bridge wait
	if !p.GetBridgeTx().GetTxCompleted() {
		if err := WaitTxComplete(taskContext, p.GetBridgeTx(), task, client, a); err != nil {
			return nil, err
		}
		if err := a.AddTx2(ctx, p.GetBridgeTx()); err != nil {
			return nil, err
		}
	}

	// wait
	if p.GetBridgeTx().GetTxCompleted() {
		task.Status = v1.ProcessStatus_StatusDone
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	return task, nil
}

func MerklyMintNFT(ctx context.Context, profile *halp.Profile, p *v1.MerklyMintAndBridgeNFTTask, client defi.MintAndBridgeNFT, estimation *v1.EstimationTx) (_ *bozdo.DefaultRes, _ *bozdo.Gas, _ error) {

	s, err := profile.GetNetworkSettings(ctx, p.FromNetwork)
	if err != nil {
		return nil, nil, err
	}
	if client == nil {
		client, err = uniclient.NewMerklyMintAndBridgeNFT(p.FromNetwork, s.BaseConfig())
		if err != nil {
			return nil, nil, err
		}
	}

	wallet, err := zksyncera.NewWalletTransactor(profile.WalletPK, client.GetNetworkId())
	if err != nil {
		return nil, nil, err
	}

	estimateOnly := estimation == nil
	var Gas *bozdo.Gas
	if estimateOnly {
		Gas = nil
	} else {
		gas, err := GasManager(estimation, s.Source, p.FromNetwork, v1.TaskType_MerklyMintAndBridgeNFT)
		if err != nil {
			return nil, nil, err
		}
		Gas = gas
	}

	res, err := client.MerklyMintNft(ctx, &merkly.MintNFTReq{
		WalletPK: wallet.PK,
		BaseReq: bozdo.BaseReq{
			EstimateOnly: estimateOnly,
			Gas:          Gas,
			Debug:        false,
		},
	})
	if err != nil {
		return nil, nil, err
	}
	return res, Gas, nil
}

func MerklyBridgeNFT(ctx context.Context, profile *halp.Profile, p *v1.MerklyMintAndBridgeNFTTask, client defi.MintAndBridgeNFT, estimation *v1.EstimationTx) (*bozdo.DefaultRes, *bozdo.Gas, error) {

	s, err := profile.GetNetworkSettings(ctx, p.FromNetwork)
	if err != nil {
		return nil, nil, err
	}
	if client == nil {
		client, err = uniclient.NewMerklyMintAndBridgeNFT(p.FromNetwork, s.BaseConfig())
		if err != nil {
			return nil, nil, err
		}
	}

	wallet, err := zksyncera.NewWalletTransactor(profile.WalletPK, client.GetNetworkId())
	if err != nil {
		return nil, nil, err
	}

	estimateOnly := estimation == nil
	var Gas *bozdo.Gas
	if estimateOnly {
		Gas = nil
	} else {
		gas, err := GasManager(estimation, s.Source, p.FromNetwork, v1.TaskType_MerklyMintAndBridgeNFT)
		if err != nil {
			return nil, nil, err
		}
		Gas = gas
	}

	nftId, ok := big.NewInt(0).SetString(p.GetNftId(), 10)
	if !ok {
		return nil, nil, errors.New("invalid nft id: " + p.GetNftId())
	}

	res, err := client.MerklyBridgeNft(ctx, &merkly.BridgeNFTReq{
		WalletPK:    wallet.PK,
		FromNetwork: p.FromNetwork,
		ToNetwork:   p.ToNetwork,
		NFTId:       nftId,
		BaseReq: bozdo.BaseReq{
			EstimateOnly: estimateOnly,
			Gas:          Gas,
			Debug:        false,
		},
	})
	if err != nil {
		return nil, nil, err
	}
	return res, Gas, nil
}

func EstimateMerklyMintCost(ctx context.Context, profile *halp.Profile, p *v1.MerklyMintAndBridgeNFTTask, client defi.MintAndBridgeNFT) (_ *v1.EstimationTx, _ error) {
	res, _, err := MerklyMintNFT(ctx, profile, p, client, nil)
	if err != nil {
		return nil, err
	}

	return GasStation(res.ECost, p.FromNetwork), nil
}

func EstimateMerklyBridgeCost(ctx context.Context, profile *halp.Profile, p *v1.MerklyMintAndBridgeNFTTask, client defi.MintAndBridgeNFT) (*v1.EstimationTx, error) {
	res, _, err := MerklyBridgeNFT(ctx, profile, p, client, nil)
	if err != nil {
		return nil, err
	}

	return GasStation(res.ECost, p.FromNetwork), nil
}
