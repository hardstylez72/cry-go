package process

import (
	"context"
	"time"

	"github.com/hardstylez72/cry/internal/defi/starknet"
	"github.com/hardstylez72/cry/internal/lib"
	"github.com/hardstylez72/cry/internal/log"
	"github.com/hardstylez72/cry/internal/orbiter"
	"github.com/hardstylez72/cry/internal/pay"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/process/task"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/settings"
	"github.com/hardstylez72/cry/internal/snapshot"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Signal = string

const (
	SignalStop    Signal = "stop"
	SignalWakeup  Signal = "wakeup"
	SignalTimeout Signal = "timeout"
)

type processId = string
type pTable struct {
	ppTable *lib.Map[string, bool]
	ppOrder []string
	signals chan Signal
}

type Dispatcher struct {
	pTable          *lib.Map[processId, *pTable]
	r               repository.ProcessRepository
	runner          *Runner
	queue           chan processId
	snapshotService snapshot.Voter
	settingsService *settings.Service
	stat            *Stat
	haalp           *halp.Halp
	payService      *pay.Service
	orbiterService  *orbiter.Service
	starknetcClient *starknet.Client
}

func NewDispatcher(
	r repository.ProcessRepository,
	runner *Runner,
	snapshotService snapshot.Voter,
	settingsService *settings.Service,
	payService *pay.Service,
	orbiterService *orbiter.Service,
	starknetcClient *starknet.Client,
) *Dispatcher {
	d := &Dispatcher{
		pTable:          lib.NewMap[processId, *pTable](),
		r:               r,
		runner:          runner,
		queue:           make(chan processId, 1000),
		snapshotService: snapshotService,
		settingsService: settingsService,
		stat: &Stat{
			ActiveProcesses: NewCounter(),
			ActiveProfiles:  NewCounter(),
			ActiveTasks:     NewCounter(),
		},
		haalp:           halp.NewHalp(runner.profileRepository, settingsService, starknetcClient),
		payService:      payService,
		orbiterService:  orbiterService,
		starknetcClient: starknetcClient,
	}

	return d
}

type Stat struct {
	ActiveProcesses *Counter
	ActiveProfiles  *Counter
	ActiveTasks     *Counter
}

func (d *Dispatcher) GetStat() *Stat {
	return d.stat
}

func (d *Dispatcher) EstimateTaskCost(ctx context.Context, profileId, taskId string) ([]*v1.EstimationTx, error) {

	out := make([]*v1.EstimationTx, 0)
	var e *v1.EstimationTx
	taskDB, err := d.r.GetProcessTask(ctx, taskId)
	if err != nil {
		return nil, err
	}

	t, err := taskDB.ToPB()
	if err != nil {
		return nil, err
	}

	profile, err := d.haalp.Profile(ctx, profileId)
	if err != nil {
		return nil, err
	}

	switch t.Task.TaskType {
	case v1.TaskType_SyncSwap:
		p := t.Task.Task.(*v1.Task_SyncSwapTask).SyncSwapTask
		e, err = task.EstimateSyncSwapCost(ctx, profile, p)
	case v1.TaskType_ZkSyncOfficialBridgeToEthereum:
		p := t.Task.Task.(*v1.Task_ZkSyncOfficialBridgeToEthereumTask).ZkSyncOfficialBridgeToEthereumTask
		e, err = task.EstimateZkSyncOfficialBridgeToEthSwapCost(ctx, profile, p)
	case v1.TaskType_StargateBridge:
		p := t.Task.Task.(*v1.Task_StargateBridgeTask).StargateBridgeTask
		e, err = task.EstimateStargateBridgeSwapCost(ctx, p, profile)
	case v1.TaskType_OrbiterBridge:
		p := t.Task.Task.(*v1.Task_OrbiterBridgeTask).OrbiterBridgeTask
		e, err = task.EstimateOrbiterBridgeCost(ctx, d.orbiterService, profile, p)
	//case v1.TaskType_ZkSyncOfficialBridgeFromEthereum:
	//	p := arg.Task.Task.Task.(*v1.Task_ZkSyncOfficialBridgeFromEthereumTask).ZkSyncOfficialBridgeFromEthereumTask
	//	return task.EstimateZkSyncOfficialBridgeFromEthSwapCost(ctx, profile, p)
	case v1.TaskType_WETH:
		p := t.Task.Task.(*v1.Task_WETHTask).WETHTask
		e, err = task.EstimateWethTaskCost(ctx, p, profile)
	case v1.TaskType_MuteioSwap:
		p := t.Task.Task.(*v1.Task_MuteioSwapTask).MuteioSwapTask
		e, err = task.EstimateMuteioSwapCost(ctx, profile, p)
	case v1.TaskType_OkexDeposit:
		p := t.Task.Task.(*v1.Task_OkexDepositTask).OkexDepositTask
		e, err = task.EstimateOkexDepositCost(ctx, profile, p, d.runner.withdrawerRepository)
	case v1.TaskType_SyncSwapLP:
		p := t.Task.Task.(*v1.Task_SyncSwapLPTask).SyncSwapLPTask
		e, err = task.EstimateSyncSwapLPCost(ctx, profile, p, nil)
	case v1.TaskType_MaverickSwap:
		p := t.Task.Task.(*v1.Task_MaverickSwapTask).MaverickSwapTask
		e, err = task.EstimateMaverickSwapCost(ctx, profile, p, nil)
	case v1.TaskType_SpaceFISwap:
		p := t.Task.Task.(*v1.Task_SpaceFiSwapTask).SpaceFiSwapTask
		e, err = task.EstimateSpaceFiSwapCost(ctx, profile, p, nil)
	case v1.TaskType_VelocoreSwap:
		p := t.Task.Task.(*v1.Task_VelocoreSwapTask).VelocoreSwapTask
		e, err = task.EstimateVelocoreSwapCost(ctx, profile, p, nil)
	case v1.TaskType_IzumiSwap:
		p := t.Task.Task.(*v1.Task_IzumiSwapTask).IzumiSwapTask
		e, err = task.EstimateIzumiSwapCost(ctx, profile, p, nil)
	case v1.TaskType_VeSyncSwap:
		p := t.Task.Task.(*v1.Task_VeSyncSwapTask).VeSyncSwapTask
		e, err = task.EstimateVeSyncSwapCost(ctx, profile, p, nil)
	case v1.TaskType_EzkaliburSwap:
		p := t.Task.Task.(*v1.Task_EzkaliburSwapTask).EzkaliburSwapTask
		e, err = task.EstimateEzkaliburSwapCost(ctx, profile, p, nil)
	case v1.TaskType_ZkSwap:
		p := t.Task.Task.(*v1.Task_ZkSwapTask).ZkSwapTask
		e, err = task.EstimateZkSwapCost(ctx, profile, p, nil)
	case v1.TaskType_TraderJoeSwap:
		p := t.Task.Task.(*v1.Task_TraderJoeSwapTask).TraderJoeSwapTask
		e, err = task.EstimateTraderJoeSwapCost(ctx, profile, p, nil)
	case v1.TaskType_MerklyMintAndBridgeNFT:
		p := t.Task.Task.(*v1.Task_MerklyMintAndBridgeNFTTask).MerklyMintAndBridgeNFTTask
		e1, _, err := task.EstimateMerklyMintCost(ctx, profile, p, nil)
		if err != nil {
			return nil, errors.Wrap(err, "EstimateMerklyMintCost")
		}
		out = append(out, e1)
		return out, nil
	case v1.TaskType_DeployStarkNetAccount:
		p := t.Task.Task.(*v1.Task_DeployStarkNetAccountTask).DeployStarkNetAccountTask
		e, err = task.EstimateDeployStarkNetAccountCost(ctx, profile, p, nil)
	case v1.TaskType_Swap10k:
		p := t.Task.Task.(*v1.Task_Swap10K).Swap10K
		e, err = task.EstimateSwap10kCost(ctx, profile, p, nil)

	default:
		return nil, errors.New("task: " + t.Task.TaskType.String() + " can not be estimated")
	}
	if err != nil {
		return nil, err
	}

	out = append(out, e)
	return out, nil
}
func (d *Dispatcher) SkipTask(ctx context.Context, processId, profileId, taskId string) error {

	if err := d.StopProcess(ctx, processId); err != nil {
		return err
	}

	if err := d.r.UpdateProcessTaskStatus(ctx, v1.ProcessStatus_StatusDone.String(), taskId, processId); err != nil {
		return err
	}

	if err := d.r.UpdateProcessProfileStatus(ctx, profileId, processId, v1.ProcessStatus_StatusReady.String()); err != nil {
		return err
	}

	return nil
}
func (d *Dispatcher) RetryProcessProfile(ctx context.Context, processId, profileId, taskId string) error {

	if err := d.StopProcess(ctx, processId); err != nil {
		return err
	}

	status, err := d.r.GetTaskStatus(ctx, taskId)
	if err != nil {
		return err
	}

	switch *status {
	case v1.ProcessStatus_StatusError, v1.ProcessStatus_StatusRetry, v1.ProcessStatus_StatusDone:
	default:
		return errors.New("forbidden action")
	}

	if err := d.r.UpdateProcessTaskStatus(ctx, v1.ProcessStatus_StatusRetry.String(), taskId, processId); err != nil {
		return err
	}
	if err := d.r.UpdateProcessProfileStatus(ctx, profileId, processId, v1.ProcessStatus_StatusReady.String()); err != nil {
		return err
	}

	go d.StartProcess(processId)

	return nil
}
func (d *Dispatcher) RunDispatcher(ctx context.Context) {

	go func() {

		processIds, err := d.r.ListProcessIdsByStatus(ctx, v1.ProcessStatus_StatusReady, v1.ProcessStatus_StatusRunning)
		if err != nil {
			log.Log.Error(errors.Wrap(err, "ListProcessIdsByStatus"))
		}

		for _, pId := range processIds {
			d.queue <- pId
		}

		ticker := time.NewTicker(time.Second * 10)
		resolver := time.NewTicker(time.Minute * 15)

		for {
			select {
			case <-resolver.C:
				log.Log.Info("resolver.start")
				pIds, err := d.r.ListProcessIdsForAutoRetry(ctx)
				if err != nil {
					log.Log.Error(errors.Wrap(err, "ListProcessIdsByStatus"))
					continue
				}

				for _, pId := range pIds {
					if err := d.RetryProcess(ctx, pId); err != nil {
						log.Log.With("processId", pId).Error("RetryProcess", zap.Error(err))
						continue
					}
					log.Log.With("processId", pId).Info("resolver.retry")
				}
				log.Log.Info("resolver.end")

			case <-ticker.C:
				pIds, err := d.r.ListProcessIdsByStatus(ctx, v1.ProcessStatus_StatusReady, v1.ProcessStatus_StatusRunning)
				if err != nil {
					log.Log.Error(errors.Wrap(err, "ListProcessIdsByStatus"))
					continue
				}

				for _, pId := range pIds {
					d.queue <- pId
				}
			case _ = <-ctx.Done():
				return
			}
		}
	}()

	for pId := range d.queue {

		l := log.Log.With("processId", pId)
		//l.Debug("got from queue")

		_, running := d.pTable.Get(pId)
		if running {
			//l.Debug("process already running")
			continue
		}

		if err := d.initProcess(ctx, pId); err != nil {
			l.Error("initProcess", zap.Error(err))
			continue
		}

		go func(pId string) {
			if err := d.RunP(ctx, pId); err != nil {
				l.Error("RunDispatcher.RunP", zap.Error(err))
			}
			d.pTable.Remove(pId)
			l.Debug("process finished")
		}(pId)
	}
}
func (d *Dispatcher) initProcess(ctx context.Context, id processId) error {
	profileIds, err := d.r.GetProcessProfileIds(ctx, id)
	if err != nil {
		return errors.Wrap(err, "GetProcessProfileIds")
	}

	profileMap := &pTable{
		ppTable: lib.NewMap[string, bool](),
		ppOrder: profileIds,
		signals: make(chan Signal),
	}
	for _, profileId := range profileIds {
		profileMap.ppTable.Set(profileId, false)
	}

	d.pTable.Set(id, profileMap)

	return nil
}
