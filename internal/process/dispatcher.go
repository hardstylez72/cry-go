package process

import (
	"context"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/orbiter"
	"github.com/hardstylez72/cry/internal/defi/starknet"
	"github.com/hardstylez72/cry/internal/lib"
	"github.com/hardstylez72/cry/internal/log"
	"github.com/hardstylez72/cry/internal/pay"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/process/task"
	"github.com/hardstylez72/cry/internal/server/config"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/settings"
	"github.com/hardstylez72/cry/internal/snapshot"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	SwapStatistic *v1.SwapStatRes
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

	go func() {
		for {
			ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
			s, err := d.SwapStat(ctx)
			cancel()
			if err != nil {
				log.Log.Error(zap.Error(err), "d.SwapStat")
			}
			d.SwapStatistic = s
			time.Sleep(time.Minute * 10)
		}
	}()

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
		e, err = task.NewSyncSwapTask().EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_ZkSyncOfficialBridgeToEthereum:
		p := t.Task.Task.(*v1.Task_ZkSyncOfficialBridgeToEthereumTask).ZkSyncOfficialBridgeToEthereumTask
		e, err = task.EstimateZkSyncOfficialBridgeToEthSwapCost(ctx, profile, p)
	case v1.TaskType_StargateBridge:
		p := t.Task.Task.(*v1.Task_StargateBridgeTask).StargateBridgeTask
		e, err = task.EstimateStargateBridgeSwapCost(ctx, p, profile)
	case v1.TaskType_OrbiterBridge:
		p := t.Task.Task.(*v1.Task_OrbiterBridgeTask).OrbiterBridgeTask
		e, err = task.EstimateOrbiterBridgeCost(ctx, d.orbiterService, profile, p)
	case v1.TaskType_ZkSyncOfficialBridgeFromEthereum:
		p := t.Task.Task.(*v1.Task_ZkSyncOfficialBridgeFromEthereumTask).ZkSyncOfficialBridgeFromEthereumTask
		e, err = task.EstimateZkSyncOfficialBridgeFromEthSwapCost(ctx, profile, p)
	case v1.TaskType_WETH:
		p := t.Task.Task.(*v1.Task_WETHTask).WETHTask
		e, err = task.EstimateWethTaskCost(ctx, p, profile)
	case v1.TaskType_MuteioSwap:
		p := t.Task.Task.(*v1.Task_MuteioSwapTask).MuteioSwapTask
		e, err = task.NewMuteioSwapTask().EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_OkexDeposit:
		p := t.Task.Task.(*v1.Task_OkexDepositTask).OkexDepositTask
		e, err = task.EstimateOkexDepositCost(ctx, profile, p, d.runner.withdrawerRepository)
	case v1.TaskType_SyncSwapLP:
		p := t.Task.Task.(*v1.Task_SyncSwapLPTask).SyncSwapLPTask
		e, err = task.EstimateSyncSwapLPCost(ctx, profile, p, nil)
	case v1.TaskType_MaverickSwap:
		p := t.Task.Task.(*v1.Task_MaverickSwapTask).MaverickSwapTask
		e, err = task.NewMaverickSwapTask().EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_SpaceFISwap:
		p := t.Task.Task.(*v1.Task_SpaceFiSwapTask).SpaceFiSwapTask
		e, err = task.NewSpaceFiSwapTask().EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_VelocoreSwap:
		p := t.Task.Task.(*v1.Task_VelocoreSwapTask).VelocoreSwapTask
		e, err = task.NewVelocoreSwapTask().EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_IzumiSwap:
		p := t.Task.Task.(*v1.Task_IzumiSwapTask).IzumiSwapTask
		e, err = task.NewIzumiSwapTask().EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_VeSyncSwap:
		p := t.Task.Task.(*v1.Task_VeSyncSwapTask).VeSyncSwapTask
		e, err = task.NewVeSyncSwapTask().EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_EzkaliburSwap:
		p := t.Task.Task.(*v1.Task_EzkaliburSwapTask).EzkaliburSwapTask
		e, err = task.NewEzkaliburSwapTask().EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_ZkSwap:
		p := t.Task.Task.(*v1.Task_ZkSwapTask).ZkSwapTask
		e, err = task.NewZkSwapTask().EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_TraderJoeSwap:
		p := t.Task.Task.(*v1.Task_TraderJoeSwapTask).TraderJoeSwapTask
		e, err = task.NewTraderJoeSwapTask().EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_MerklyMintAndBridgeNFT:
		p := t.Task.Task.(*v1.Task_MerklyMintAndBridgeNFTTask).MerklyMintAndBridgeNFTTask
		e1, err := task.EstimateMerklyMintCost(ctx, profile, p, nil)
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
		e, err = (&task.StarketSwapHalper{v1.TaskType_Swap10k}).EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_PancakeSwap:
		p := t.Task.Task.(*v1.Task_PancakeSwapTask).PancakeSwapTask
		e, err = task.NewPancakeSwapTask().EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_SithSwap:
		p := t.Task.Task.(*v1.Task_SithSwapTask).SithSwapTask
		e, err = (&task.StarketSwapHalper{v1.TaskType_SithSwap}).EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_JediSwap:
		p := t.Task.Task.(*v1.Task_JediSwapTask).JediSwapTask
		e, err = (&task.StarketSwapHalper{v1.TaskType_JediSwap}).EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_MySwap:
		p := t.Task.Task.(*v1.Task_MySwapTask).MySwapTask
		e, err = (&task.StarketSwapHalper{v1.TaskType_MySwap}).EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_ProtossSwap:
		p := t.Task.Task.(*v1.Task_ProtosSwapTask).ProtosSwapTask
		e, err = (&task.StarketSwapHalper{v1.TaskType_ProtossSwap}).EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_StarkNetBridge:
		p := t.Task.Task.(*v1.Task_StarkNetBridgeTask).StarkNetBridgeTask
		from, to, errr := task.LiquidityBridgeProfiles(ctx, d.haalp, d.runner.profileRepository, p, profile)
		if errr != nil {
			return nil, errr
		}
		e, err = (&task.DefaultLiquidityBridgeTaskHalper{v1.TaskType_StarkNetBridge}).EstimateCost(ctx, from, to, p, nil, nil)
	case v1.TaskType_Dmail:
		p := t.Task.Task.(*v1.Task_DmailTask).DmailTask
		e, err = task.EstimateDmailSwapCost(ctx, p, profile)
	case v1.TaskType_StarkNetIdMint:
		p := t.Task.Task.(*v1.Task_StarkNetIdMintTask).StarkNetIdMintTask
		e, err = task.EstimateMintCost(ctx, p, profile)
	case v1.TaskType_OdosSwap:
		p := t.Task.Task.(*v1.Task_OdosSwapTask).OdosSwapTask
		e, err = task.NewOdosSwapTask().EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_AcrossBridge:
		p := t.Task.Task.(*v1.Task_AcrossBridgeTask).AcrossBridgeTask
		e, err = task.NewAcrossBridgeTask().EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_AvnuSwap:
		p := t.Task.Task.(*v1.Task_AvnuSwapTask).AvnuSwapTask
		e, err = task.NewAvnuSwapTask().EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_FibrousSwap:
		p := t.Task.Task.(*v1.Task_FibrousSwapTask).FibrousSwapTask
		e, err = task.NewFibrousSwapTask().EstimateCost(ctx, profile, p, nil)
	case v1.TaskType_ZkLendLP:
		p := t.Task.Task.(*v1.Task_ZkLendLPTask).ZkLendLPTask
		e, err = task.NewZkLendLPTask().EstimateLPCost(ctx, profile, p, nil)
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

		processIds, err := d.r.ProcessIds(ctx, v1.ProcessStatus_StatusRunning)
		if err != nil {
			log.Log.Error(errors.Wrap(err, "ProcessIds"))
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
					log.Log.Error(errors.Wrap(err, "RunDispatcher.ListProcessIdsForAutoRetry"))
					continue
				}

				readyIds, err := d.r.ProcessIdsReady(ctx, time.Now().UTC())
				if err != nil {
					log.Log.Error(errors.Wrap(err, "RunDispatcher.ProcessIdsReady"))
					continue
				}

				for _, el := range readyIds {
					d.StartProcess(el)
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
				pIds, err := d.r.ProcessIds(ctx, v1.ProcessStatus_StatusRunning)
				if err != nil {
					log.Log.Error(errors.Wrap(err, "ProcessIds"))
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

func (d *Dispatcher) SwapStat(ctx context.Context) (*v1.SwapStatRes, error) {

	wg := sync.WaitGroup{}
	wg.Add(4)
	zkETHUSDC := make([]*v1.SwapStatItem, 0)
	zkSyncProfile, err := d.haalp.Profile(ctx, config.CFG.TestEVMWalletId)
	if err != nil {
		return nil, err
	}

	starknetProfile, err := d.haalp.Profile(ctx, config.CFG.TestStarkNetWalletId)
	if err != nil {
		return nil, err
	}

	go func() {
		defer wg.Done()
		zk, err := d.estimateZkSyncSwapRatio(ctx, v1.Token_ETH, v1.Token_USDC, v1.Network_ZKSYNCERA, zkSyncProfile)
		if err != nil {
			return
		}

		zkETHUSDC = zk.Values()
		sort.Slice(zkETHUSDC, func(i, j int) bool {
			return zkETHUSDC[i].RateRatio < zkETHUSDC[j].RateRatio
		})
	}()

	zkUSDCETH := make([]*v1.SwapStatItem, 0)
	go func() {
		defer wg.Done()
		zk, err := d.estimateZkSyncSwapRatio(ctx, v1.Token_USDC, v1.Token_ETH, v1.Network_ZKSYNCERA, zkSyncProfile)
		if err != nil {
			return
		}

		zkUSDCETH = zk.Values()
		sort.Slice(zkUSDCETH, func(i, j int) bool {
			return zkUSDCETH[i].RateRatio < zkUSDCETH[j].RateRatio
		})
	}()

	starknetUSDCETH := make([]*v1.SwapStatItem, 0)
	go func() {
		defer wg.Done()
		zk, err := d.estimateStarknetSwapRatio(ctx, v1.Token_USDC, v1.Token_ETH, v1.Network_StarkNet, starknetProfile)
		if err != nil {
			return
		}
		starknetUSDCETH = zk.Values()

		sort.Slice(starknetUSDCETH, func(i, j int) bool {
			return starknetUSDCETH[i].RateRatio < starknetUSDCETH[j].RateRatio
		})
	}()

	starknetETHUSDC := make([]*v1.SwapStatItem, 0)
	go func() {
		defer wg.Done()
		zk, err := d.estimateStarknetSwapRatio(ctx, v1.Token_ETH, v1.Token_USDC, v1.Network_StarkNet, starknetProfile)
		if err != nil {
			return
		}

		starknetETHUSDC = zk.Values()
		sort.Slice(starknetETHUSDC, func(i, j int) bool {
			return starknetETHUSDC[i].RateRatio < starknetETHUSDC[j].RateRatio
		})
	}()

	wg.Wait()

	return &v1.SwapStatRes{
		ZkSyncETHUSDC:   zkETHUSDC,
		ZkSyncUSDCETH:   zkUSDCETH,
		StarknetETHUSDC: starknetETHUSDC,
		StarknetUSDCETH: starknetUSDCETH,
		Updated:         timestamppb.Now(),
	}, nil
}

func (d *Dispatcher) estimateZkSyncSwapRatio(ctx context.Context, t1, t2 v1.Token, network v1.Network, profile *halp.Profile) (*lib.Map[v1.TaskType, *v1.SwapStatItem], error) {

	m := lib.NewMap[v1.TaskType, *v1.SwapStatItem]()

	era := &v1.DefaultSwap{
		Amount: &v1.Amount{
			Kind: &v1.Amount_SendPercent{
				SendPercent: 10,
			},
		},
		Network:   network,
		FromToken: t1,
		ToToken:   t2,
	}

	s, err := profile.GetNetworkSettings(ctx, network)
	if err != nil {
		return nil, err
	}

	tasks := []*task.ZkSyncSwap{
		task.NewZkSwapTask(),
		task.NewSyncSwapTask(),
		task.NewMaverickSwapTask(),
		task.NewVelocoreSwapTask(),
		task.NewMuteioSwapTask(),
		task.NewEzkaliburSwapTask(),
		task.NewPancakeZkSyncSwapTask(),
		task.NewVeSyncSwapTask(),
		task.NewIzumiSwapTask(),
		task.NewSpaceFiSwapTask(),
		task.NewOdosSwapTaskZksync(),
	}

	for _, tasker := range tasks {
		e, err := tasker.EstimateCost(ctx, profile, era, nil)
		if err == nil {
			rate := RateRatio(e)
			if rate != nil {
				slip := GetSlippage(tasker.TaskType, s.Source)
				m.Set(tasker.TaskType, &v1.SwapStatItem{
					Type:      tasker.Type(),
					Slippage:  slip,
					RateRatio: *rate,
					Sum:       *rate + slip,
				})
			} else {
				log.Log.Error("swap stat failed: " + tasker.TaskType.String() + t1.String() + t2.String())
			}
		} else {
			log.Log.Error(zap.Error(err), "swap stat failed: "+tasker.TaskType.String()+t1.String()+t2.String())
		}

	}

	return m, nil
}

func (d *Dispatcher) estimateStarknetSwapRatio(ctx context.Context, t1, t2 v1.Token, network v1.Network, profile *halp.Profile) (*lib.Map[v1.TaskType, *v1.SwapStatItem], error) {

	m := lib.NewMap[v1.TaskType, *v1.SwapStatItem]()

	era := &v1.DefaultSwap{
		Amount: &v1.Amount{
			Kind: &v1.Amount_SendPercent{
				SendPercent: 10,
			},
		},
		Network:   network,
		FromToken: t1,
		ToToken:   t2,
	}

	s, err := profile.GetNetworkSettings(ctx, network)
	if err != nil {
		return nil, err
	}

	tasks := []*task.StarkNetSwap{
		task.NewSithSwapTask(),
		task.NewSwap10kSwapTask(),
		task.NewProtossSwapTask(),
		task.NewMySwapSwapTask(),
		task.NewJediSwapTask(),
		task.NewAvnuSwapTask(),
		task.NewFibrousSwapTask(),
	}

	for _, tasker := range tasks {
		e, err := tasker.EstimateCost(ctx, profile, era, nil)
		if err == nil {
			rate := RateRatio(e)
			if rate != nil {
				slip := GetSlippage(tasker.TaskType, s.Source)
				m.Set(tasker.TaskType, &v1.SwapStatItem{
					Type:      tasker.Type(),
					Slippage:  slip,
					RateRatio: *rate,
					Sum:       *rate + slip,
				})
			} else {
				log.Log.Error("swap stat failed: " + tasker.TaskType.String() + t1.String() + t2.String())
			}
		} else {
			log.Log.Error(zap.Error(err), "swap stat failed: "+tasker.TaskType.String()+t1.String()+t2.String())
		}

	}

	return m, nil
}

func GetSlippage(t v1.TaskType, s *v1.NetworkSettings) float64 {
	if s == nil {
		return -1
	}

	v, ok := s.TaskSettings[t.String()]
	if !ok {
		return -1
	}

	f, err := lib.StringToFloat(v.GetSlippage())
	if err != nil {
		return -1
	}

	return f
}

func RateRatio(e *v1.EstimationTx) *float64 {
	if e == nil || e.Details == nil {
		return nil
	}

	var limit float64 = 0
	for _, d := range e.Details {
		if d.Key == bozdo.TxDetailEaxchnageRateRatio {
			value := strings.ReplaceAll(d.Value, " %", "")
			f, err := lib.StringToFloat(value)
			if err != nil {
				return nil
			}
			limit = f
			break
		}
	}

	if limit == 0 {
		return nil
	}

	return &limit
}
