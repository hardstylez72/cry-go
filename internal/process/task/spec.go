package task

import (
	"context"
	"fmt"
	"time"

	paycli "github.com/hardstylez72/cry-pay/proto/gen/go/v1"
	"github.com/hardstylez72/cry/internal/defi/orbiter"
	"github.com/hardstylez72/cry/internal/pay"
	"github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/server/config"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/settings"
	"github.com/hardstylez72/cry/internal/snapshot"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Input struct {
	ProfileId string
	ProcessId string
	UserId    string
	Task      *v1.ProcessTask

	L                    *zap.SugaredLogger
	ProcessRepository    repository.ProcessRepository
	ProfileRepository    repository.ProfileRepository
	WithdrawerRepository repository.WithdrawerRepository
	SettingsService      *settings.Service

	Snapshot snapshot.Voter

	Halper     *halp.Halp
	PayService *pay.Service
	User       *repository.User
	Orbiter    *orbiter.Service
}

type TaskRes struct {
	Task *v1.ProcessTask
}

type EstimateSyncSwapCostReq struct {
	Task      *v1.ProcessTask
	UserId    string
	ProcessId string
	ProfileId string
}

type Tasker interface {
	Run(ctx context.Context, arg *Input) (*v1.ProcessTask, error)
	Stop() error
	Type() v1.TaskType
}

const (
	taskTimeout         = time.Minute * 2
	taskStarkNetTimeout = time.Minute * 10
)

type Spec struct {
	Payable  bool
	Tasker   Tasker
	Estimate func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error)
	Desc     func(m *v1.Task) ([]byte, error)
	CastR    func(in RandomTask) (*v1.Task, error)
	Input    func(m *v1.Task) (*TaskInput, error)
	Output   func(m *v1.Task) (*TaskOutput, error)
}

type TaskInput struct {
	NoInput bool
	Native  bool
	Token   v1.Token
	Network v1.Network
}

type TaskOutput struct {
	NoOutput bool
	Token    v1.Token
	Network  v1.Network
}

type EstimateArg struct {
	Task                 *v1.Task
	Profile              *halp.Profile
	ProfileService       *halp.Halp
	WithdrawerRepository repository.WithdrawerRepository
	ProfileRepository    repository.ProfileRepository
	OrbiterService       *orbiter.Service
}

var SpecMap = map[v1.TaskType]Spec{
	//OTHER
	v1.TaskType_DeployStarkNetAccount: {
		Payable: true,
		Tasker:  &Wrap{Tasker: &DeployStarkNetAccountTask{}},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_DeployStarkNetAccountTask).DeployStarkNetAccountTask
			return EstimateDeployStarkNetAccountCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_DeployStarkNetAccountTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_DeployStarkNetAccountTask)")
			}
			return Marshal(t.DeployStarkNetAccountTask)
		},
	},
	v1.TaskType_Dmail: {
		Payable: true,
		Tasker:  &Wrap{Tasker: &DmailTask{}},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_DmailTask).DmailTask
			return EstimateDmailSwapCost(ctx, p, a.Profile)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_DmailTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_DmailTask)")
			}
			return Marshal(t.DmailTask)
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			t, ok := m.Task.(*v1.Task_DmailTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_DmailTask)")
			}
			return &TaskInput{
				Native:  true,
				Network: t.DmailTask.Network,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			t, ok := m.Task.(*v1.Task_DmailTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_DmailTask)")
			}
			return &TaskOutput{
				NoOutput: true,
				Network:  t.DmailTask.Network,
			}, nil
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_DmailTask{DmailTask: CastDefaultSimple(in.P)},
			}, nil
		},
	},
	v1.TaskType_Delay: {
		Payable:  false,
		Tasker:   &Wrap{Tasker: &taskDelay{}},
		Estimate: nil,
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_DelayTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_DelayTask)")
			}
			return Marshal(t.DelayTask)
		},
		CastR: nil,
		Input: func(m *v1.Task) (*TaskInput, error) {
			_, ok := m.Task.(*v1.Task_DelayTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_DelayTask)")
			}
			return &TaskInput{
				NoInput: true,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			_, ok := m.Task.(*v1.Task_DelayTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_DelayTask)")
			}
			return &TaskOutput{
				NoOutput: true,
			}, nil
		},
	},
	v1.TaskType_OkexDeposit: {
		Payable: false,
		Tasker:  &Wrap{Tasker: &OkexDepositTask{}},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_OkexDepositTask).OkexDepositTask
			return EstimateOkexDepositCost(ctx, a.Profile, p, a.WithdrawerRepository)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_OkexDepositTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_OkexDepositTask)")
			}
			return Marshal(t.OkexDepositTask)
		},
		CastR: nil,
	},
	v1.TaskType_WithdrawExchange: {
		Payable:  false,
		Tasker:   &Wrap{Tasker: &WithdrawExchange{}},
		Estimate: nil,
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_WithdrawExchangeTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_WithdrawExchangeTask)")
			}
			return Marshal(t.WithdrawExchangeTask)
		},
		CastR: nil,
	},
	v1.TaskType_ExchangeSwap: {
		Payable:  false,
		Tasker:   &Wrap{Tasker: &ExchangeSwapTask{}},
		Estimate: nil,
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_ExchangeSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_ExchangeSwapTask)")
			}
			return Marshal(t.ExchangeSwapTask)
		},
		CastR: nil,
	},
	v1.TaskType_Mock: {
		Payable:  false,
		Tasker:   &Wrap{Tasker: &mockTask{}},
		Estimate: nil,
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_MockTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MockTask)")
			}
			return Marshal(t.MockTask)
		},
	},
	v1.TaskType_SnapshotVote: {
		Payable:  true,
		Tasker:   &Wrap{Tasker: &SnapshotVoteTask{}},
		Estimate: nil,
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_SnapshotVoteTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_SnapshotVoteTask)")
			}
			return Marshal(t.SnapshotVoteTask)
		},
	},
	v1.TaskType_OkexBinance: {
		Payable:  true,
		Tasker:   &Wrap{Tasker: &mockTask{}},
		Estimate: nil,
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_OkexBinanaceTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_OkexDepositTask)")
			}
			return Marshal(t.OkexBinanaceTask)
		},
	},

	// LP
	v1.TaskType_SyncSwapLP: {
		Payable: true,
		Tasker:  &Wrap{Tasker: &SyncSwapLPTask{}},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_SyncSwapLPTask).SyncSwapLPTask
			return EstimateSyncSwapLPCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_SyncSwapLPTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_SyncSwapLPTask)")
			}
			return Marshal(t.SyncSwapLPTask)
		},
	},
	v1.TaskType_ZkLendLP: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewZkLendLPTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_ZkLendLPTask).ZkLendLPTask
			return NewZkLendLPTask().EstimateLPCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_ZkLendLPTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_ZkLandLPTask)")
			}
			return Marshal(t.ZkLendLPTask)
		},
	},
	v1.TaskType_AaveLP: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewAaveLPTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_AaveLPTask).AaveLPTask
			return NewAaveLPTask().EstimateLPCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_AaveLPTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_AaveLPTask)")
			}
			return Marshal(t.AaveLPTask)
		},
	},
	v1.TaskType_NostraLP: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewNostraLPTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_NostraLPTask).NostraLPTask
			return NewNostraLPTask().EstimateLPCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_NostraLPTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_NostraLPTask)")
			}
			return Marshal(t.NostraLPTask)
		},
	},
	// NFT
	v1.TaskType_MintFun: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewMintFunMintTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_MintFunTask).MintFunTask
			return NewMintFunMintTask().EstimateCost(ctx, p, a.Profile)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_MintFunTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MintFunTask)")
			}
			return Marshal(t.MintFunTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_MintFunTask{MintFunTask: CastDefaultSimple(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			_, ok := m.Task.(*v1.Task_MintFunTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MintFunTask)")
			}
			return &TaskInput{
				NoInput: true,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			_, ok := m.Task.(*v1.Task_MintFunTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MintFunTask)")
			}
			return &TaskOutput{
				NoOutput: true,
			}, nil
		},
	},
	v1.TaskType_MintMerkly: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewMerklyMintTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_MintMerklyTask).MintMerklyTask
			return NewMerklyMintTask().EstimateCost(ctx, p, a.Profile)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_MintMerklyTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MintMerklyTask)")
			}
			return Marshal(t.MintMerklyTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_MintMerklyTask{MintMerklyTask: CastDefaultSimple(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			_, ok := m.Task.(*v1.Task_MintMerklyTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MintMerklyTask)")
			}
			return &TaskInput{
				NoInput: true,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			_, ok := m.Task.(*v1.Task_MintMerklyTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MintMerklyTask)")
			}
			return &TaskOutput{
				NoOutput: true,
			}, nil
		},
	},
	v1.TaskType_MintZerius: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewZeriusMintTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_MintZeriusTask).MintZeriusTask
			return NewZeriusMintTask().EstimateCost(ctx, p, a.Profile)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_MintZeriusTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MintZeriusTask)")
			}
			return Marshal(t.MintZeriusTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_MintZeriusTask{MintZeriusTask: CastDefaultSimple(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			_, ok := m.Task.(*v1.Task_MintZeriusTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MintZeriusTask)")
			}
			return &TaskInput{
				NoInput: true,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			_, ok := m.Task.(*v1.Task_MintZeriusTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MintZeriusTask)")
			}
			return &TaskOutput{
				NoOutput: true,
			}, nil
		},
	},
	v1.TaskType_StarkNetIdMint: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewStarkNetIdMintTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_StarkNetIdMintTask).StarkNetIdMintTask
			return NewStarkNetIdMintTask().EstimateCost(ctx, p, a.Profile)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_StarkNetIdMintTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.StarkNetIdMintTask)")
			}
			return Marshal(t.StarkNetIdMintTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_StarkNetIdMintTask{StarkNetIdMintTask: CastDefaultSimple(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			_, ok := m.Task.(*v1.Task_StarkNetIdMintTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_StarkNetIdMintTask)")
			}
			return &TaskInput{
				NoInput: true,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			_, ok := m.Task.(*v1.Task_StarkNetIdMintTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_StarkNetIdMintTask)")
			}
			return &TaskOutput{
				NoOutput: true,
			}, nil
		},
	},
	v1.TaskType_MerklyMintAndBridgeNFT: {
		Payable: true,
		Tasker:  &Wrap{Tasker: &MerklyMintAndBridgeNFTTask{}},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_MerklyMintAndBridgeNFTTask).MerklyMintAndBridgeNFTTask
			return EstimateMerklyMintCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_MerklyMintAndBridgeNFTTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MerklyMintAndBridgeNFTTask)")
			}
			return Marshal(t.MerklyMintAndBridgeNFTTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			from := CastDefaultSimple(in.P).Network
			return &v1.Task{
				TaskType: in.Type,
				Task: &v1.Task_MerklyMintAndBridgeNFTTask{MerklyMintAndBridgeNFTTask: &v1.MerklyMintAndBridgeNFTTask{
					FromNetwork: from,
					ToNetwork:   v1.Network_POLIGON, //todo: fix
				}},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			_, ok := m.Task.(*v1.Task_MerklyMintAndBridgeNFTTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MerklyMintAndBridgeNFTTask)")
			}
			return &TaskInput{
				NoInput: true,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			_, ok := m.Task.(*v1.Task_MerklyMintAndBridgeNFTTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MerklyMintAndBridgeNFTTask)")
			}
			return &TaskOutput{
				NoOutput: true,
			}, nil
		},
	},
	//BRIDGES
	v1.TaskType_ZkSyncOfficialBridgeToEthereum: {
		Payable: true,
		Tasker:  &Wrap{Tasker: &ZksyncOfficialBridgeToEthereumTask{}},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_ZkSyncOfficialBridgeToEthereumTask).ZkSyncOfficialBridgeToEthereumTask
			return EstimateZkSyncOfficialBridgeToEthSwapCost(ctx, a.Profile, p)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_ZkSyncOfficialBridgeToEthereumTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_ZkSyncOfficialBridgeToEthereumTask)")
			}
			return Marshal(t.ZkSyncOfficialBridgeToEthereumTask)
		},
	},
	v1.TaskType_ZkSyncOfficialBridgeFromEthereum: {
		Payable: true,
		Tasker:  &Wrap{Tasker: &ZksyncOfficialBridgeFromEthereumTask{}},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_ZkSyncOfficialBridgeFromEthereumTask).ZkSyncOfficialBridgeFromEthereumTask
			return EstimateZkSyncOfficialBridgeFromEthSwapCost(ctx, a.Profile, p)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_ZkSyncOfficialBridgeFromEthereumTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_ZkSyncOfficialBridgeFromEthereumTask)")
			}
			return Marshal(t.ZkSyncOfficialBridgeFromEthereumTask)
		},
	},
	v1.TaskType_OrbiterBridge: {
		Payable: true,
		Tasker:  &Wrap{Tasker: &OrbiterBridgeTask{}},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_OrbiterBridgeTask).OrbiterBridgeTask
			return EstimateOrbiterBridgeCost(ctx, a.OrbiterService, a.Profile, p)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_OrbiterBridgeTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_OrbiterBridgeTask)")
			}
			return Marshal(t.OrbiterBridgeTask)
		},
	},
	v1.TaskType_StargateBridge: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewStargateBridgeTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_StargateBridge).StargateBridge
			return NewStargateBridgeTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_StargateBridge)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_StargateBridge)")
			}
			return Marshal(t.StargateBridge)
		},
	},
	v1.TaskType_AcrossBridge: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewAcrossBridgeTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_AcrossBridgeTask).AcrossBridgeTask
			return NewAcrossBridgeTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_AcrossBridgeTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_AcrossBridgeTask)")
			}
			return Marshal(t.AcrossBridgeTask)
		},
	},
	v1.TaskType_StarkNetBridge: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewStarkNetBridgeTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_StarkNetBridgeTask).StarkNetBridgeTask
			from, to, errr := LiquidityBridgeProfiles(ctx, a.ProfileService, a.ProfileRepository, p, a.Profile)
			if errr != nil {
				return nil, errr
			}
			return (&DefaultLiquidityBridgeTaskHalper{v1.TaskType_StarkNetBridge}).EstimateCost(ctx, from, to, p, nil, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_StarkNetBridgeTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_StarkNetBridgeTask)")
			}
			return Marshal(t.StarkNetBridgeTask)
		},
	},
	v1.TaskType_CoreDaoBridge: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewCoreDaoBridgeTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_CoreDaoBridge).CoreDaoBridge
			return NewCoreDaoBridgeTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_CoreDaoBridge)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_CoreDaoBridge)")
			}
			return Marshal(t.CoreDaoBridge)
		},
	},

	// SWAPS
	v1.TaskType_TestNetBridgeSwap: {
		Payable:  true,
		Tasker:   &Wrap{Tasker: &TestNetBridgeSwapTask{}},
		Estimate: nil,
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_TestNetBridgeSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_TestNetBridgeSwapTask)")
			}
			return Marshal(t.TestNetBridgeSwapTask)
		},
	},
	v1.TaskType_SyncSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewSyncSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_SyncSwapTask).SyncSwapTask
			return NewSyncSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_SyncSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_SyncSwapTask)")
			}
			return Marshal(t.SyncSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_SyncSwapTask{SyncSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_SyncSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_SyncSwapTask)")
			}
			return &TaskInput{
				Network: p.SyncSwapTask.Network,
				Token:   p.SyncSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_SyncSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_SyncSwapTask)")
			}
			return &TaskOutput{
				Network: p.SyncSwapTask.Network,
				Token:   p.SyncSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_VelocoreSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewVelocoreSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_VelocoreSwapTask).VelocoreSwapTask
			return NewVelocoreSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_VelocoreSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_VelocoreSwapTask)")
			}
			return Marshal(t.VelocoreSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_VelocoreSwapTask{VelocoreSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_VelocoreSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_VelocoreSwapTask)")
			}
			return &TaskInput{
				Network: p.VelocoreSwapTask.Network,
				Token:   p.VelocoreSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_VelocoreSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_VelocoreSwapTask)")
			}
			return &TaskOutput{
				Network: p.VelocoreSwapTask.Network,
				Token:   p.VelocoreSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_SpaceFISwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewSpaceFiSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_SpaceFiSwapTask).SpaceFiSwapTask
			return NewSpaceFiSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_SpaceFiSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_SpaceFiSwapTask)")
			}
			return Marshal(t.SpaceFiSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_SpaceFiSwapTask{SpaceFiSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_SpaceFiSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_SpaceFiSwapTask)")
			}
			return &TaskInput{
				Network: p.SpaceFiSwapTask.Network,
				Token:   p.SpaceFiSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_SpaceFiSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_SpaceFiSwapTask)")
			}
			return &TaskOutput{
				Network: p.SpaceFiSwapTask.Network,
				Token:   p.SpaceFiSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_MaverickSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewMaverickSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_MaverickSwapTask).MaverickSwapTask
			return NewMaverickSwapTask().EstimateCost(ctx, a.Profile, p, nil)

		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_MaverickSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MaverickSwapTask)")
			}
			return Marshal(t.MaverickSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_MaverickSwapTask{MaverickSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_MaverickSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MaverickSwapTask)")
			}
			return &TaskInput{
				Network: p.MaverickSwapTask.Network,
				Token:   p.MaverickSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_MaverickSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MaverickSwapTask)")
			}
			return &TaskOutput{
				Network: p.MaverickSwapTask.Network,
				Token:   p.MaverickSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_MuteioSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewMuteioSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_MuteioSwapTask).MuteioSwapTask
			return NewMuteioSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_MuteioSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MuteioSwapTask)")
			}
			return Marshal(t.MuteioSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_MuteioSwapTask{MuteioSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_MuteioSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MuteioSwapTask)")
			}
			return &TaskInput{
				Network: p.MuteioSwapTask.Network,
				Token:   p.MuteioSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_MuteioSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MuteioSwapTask)")
			}
			return &TaskOutput{
				Network: p.MuteioSwapTask.Network,
				Token:   p.MuteioSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_IzumiSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewIzumiSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_IzumiSwapTask).IzumiSwapTask
			return NewIzumiSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_IzumiSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_IzumiSwapTask)")
			}
			return Marshal(t.IzumiSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_IzumiSwapTask{IzumiSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_IzumiSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_IzumiSwapTask)")
			}
			return &TaskInput{
				Network: p.IzumiSwapTask.Network,
				Token:   p.IzumiSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_IzumiSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_IzumiSwapTask)")
			}
			return &TaskOutput{
				Network: p.IzumiSwapTask.Network,
				Token:   p.IzumiSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_VeSyncSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewVeSyncSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_VeSyncSwapTask).VeSyncSwapTask
			return NewVeSyncSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_VeSyncSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_VeSyncSwapTask)")
			}
			return Marshal(t.VeSyncSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_VeSyncSwapTask{VeSyncSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_VeSyncSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_VeSyncSwapTask)")
			}
			return &TaskInput{
				Network: p.VeSyncSwapTask.Network,
				Token:   p.VeSyncSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_VeSyncSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_VeSyncSwapTask)")
			}
			return &TaskOutput{
				Network: p.VeSyncSwapTask.Network,
				Token:   p.VeSyncSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_EzkaliburSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewEzkaliburSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_EzkaliburSwapTask).EzkaliburSwapTask
			return NewEzkaliburSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_EzkaliburSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_EzkaliburSwapTask)")
			}
			return Marshal(t.EzkaliburSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_EzkaliburSwapTask{EzkaliburSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_EzkaliburSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_EzkaliburSwapTask)")
			}
			return &TaskInput{
				Network: p.EzkaliburSwapTask.Network,
				Token:   p.EzkaliburSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_EzkaliburSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_EzkaliburSwapTask)")
			}
			return &TaskOutput{
				Network: p.EzkaliburSwapTask.Network,
				Token:   p.EzkaliburSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_ZkSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewZkSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_ZkSwapTask).ZkSwapTask
			return NewZkSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_ZkSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_ZkSwapTask)")
			}
			return Marshal(t.ZkSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_ZkSwapTask{ZkSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_ZkSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_ZkSwapTask)")
			}
			return &TaskInput{
				Network: p.ZkSwapTask.Network,
				Token:   p.ZkSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_ZkSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_ZkSwapTask)")
			}
			return &TaskOutput{
				Network: p.ZkSwapTask.Network,
				Token:   p.ZkSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_TraderJoeSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewTraderJoeSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_TraderJoeSwapTask).TraderJoeSwapTask
			return NewTraderJoeSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_TraderJoeSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_TraderJoeSwapTask)")
			}
			return Marshal(t.TraderJoeSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_TraderJoeSwapTask{TraderJoeSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_TraderJoeSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_TraderJoeSwapTask)")
			}
			return &TaskInput{
				Network: p.TraderJoeSwapTask.Network,
				Token:   p.TraderJoeSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_TraderJoeSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_TraderJoeSwapTask)")
			}
			return &TaskOutput{
				Network: p.TraderJoeSwapTask.Network,
				Token:   p.TraderJoeSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_Swap10k: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewSwap10kSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_Swap10K).Swap10K
			return (&StarketSwapHalper{v1.TaskType_Swap10k}).EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_Swap10K)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_Swap10K)")
			}
			return Marshal(t.Swap10K)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_Swap10K{Swap10K: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_Swap10K)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_Swap10K)")
			}
			return &TaskInput{
				Network: p.Swap10K.Network,
				Token:   p.Swap10K.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_Swap10K)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_Swap10K)")
			}
			return &TaskOutput{
				Network: p.Swap10K.Network,
				Token:   p.Swap10K.ToToken,
			}, nil
		},
	},
	v1.TaskType_PancakeSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewPancakeSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_PancakeSwapTask).PancakeSwapTask
			return NewPancakeSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_PancakeSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_PancakeSwapTask)")
			}
			return Marshal(t.PancakeSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_PancakeSwapTask{PancakeSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_PancakeSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_PancakeSwapTask)")
			}
			return &TaskInput{
				Network: p.PancakeSwapTask.Network,
				Token:   p.PancakeSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_PancakeSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_PancakeSwapTask)")
			}
			return &TaskOutput{
				Network: p.PancakeSwapTask.Network,
				Token:   p.PancakeSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_SithSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewSithSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_SithSwapTask).SithSwapTask
			return (&StarketSwapHalper{v1.TaskType_SithSwap}).EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_SithSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_SithSwapTask)")
			}
			return Marshal(t.SithSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_SithSwapTask{SithSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_SithSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_SithSwapTask)")
			}
			return &TaskInput{
				Network: p.SithSwapTask.Network,
				Token:   p.SithSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_SithSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_SithSwapTask)")
			}
			return &TaskOutput{
				Network: p.SithSwapTask.Network,
				Token:   p.SithSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_JediSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewJediSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_JediSwapTask).JediSwapTask
			return (&StarketSwapHalper{v1.TaskType_JediSwap}).EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_JediSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_JediSwapTask)")
			}
			return Marshal(t.JediSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_JediSwapTask{JediSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_JediSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_JediSwapTask)")
			}
			return &TaskInput{
				Network: p.JediSwapTask.Network,
				Token:   p.JediSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_JediSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_JediSwapTask)")
			}
			return &TaskOutput{
				Network: p.JediSwapTask.Network,
				Token:   p.JediSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_MySwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewMySwapSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_MySwapTask).MySwapTask
			return (&StarketSwapHalper{v1.TaskType_MySwap}).EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_MySwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MySwapTask)")
			}
			return Marshal(t.MySwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_MySwapTask{MySwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_MySwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MySwapTask)")
			}
			return &TaskInput{
				Network: p.MySwapTask.Network,
				Token:   p.MySwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_MySwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MySwapTask)")
			}
			return &TaskOutput{
				Network: p.MySwapTask.Network,
				Token:   p.MySwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_ProtossSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewProtossSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_ProtosSwapTask).ProtosSwapTask
			return (&StarketSwapHalper{v1.TaskType_ProtossSwap}).EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_ProtosSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_ProtosSwapTask)")
			}
			return Marshal(t.ProtosSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_ProtosSwapTask{ProtosSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_ProtosSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_ProtosSwapTask)")
			}
			return &TaskInput{
				Network: p.ProtosSwapTask.Network,
				Token:   p.ProtosSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_ProtosSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_ProtosSwapTask)")
			}
			return &TaskOutput{
				Network: p.ProtosSwapTask.Network,
				Token:   p.ProtosSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_OdosSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewOdosSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_OdosSwapTask).OdosSwapTask
			return NewOdosSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_OdosSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_OdosSwapTask)")
			}
			return Marshal(t.OdosSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_OdosSwapTask{OdosSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_OdosSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_OdosSwapTask)")
			}
			return &TaskInput{
				Network: p.OdosSwapTask.Network,
				Token:   p.OdosSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_OdosSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_OdosSwapTask)")
			}
			return &TaskOutput{
				Network: p.OdosSwapTask.Network,
				Token:   p.OdosSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_AvnuSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewAvnuSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_AvnuSwapTask).AvnuSwapTask
			return NewAvnuSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_AvnuSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_AvnuSwapTask)")
			}
			return Marshal(t.AvnuSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_AvnuSwapTask{AvnuSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_AvnuSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_AvnuSwapTask)")
			}
			return &TaskInput{
				Network: p.AvnuSwapTask.Network,
				Token:   p.AvnuSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_AvnuSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_AvnuSwapTask)")
			}
			return &TaskOutput{
				Network: p.AvnuSwapTask.Network,
				Token:   p.AvnuSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_FibrousSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewFibrousSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_FibrousSwapTask).FibrousSwapTask
			return NewFibrousSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_FibrousSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_FibrousSwapTask)")
			}
			return Marshal(t.FibrousSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_FibrousSwapTask{FibrousSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_FibrousSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_FibrousSwapTask)")
			}
			return &TaskInput{
				Network: p.FibrousSwapTask.Network,
				Token:   p.FibrousSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_FibrousSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_FibrousSwapTask)")
			}
			return &TaskOutput{
				Network: p.FibrousSwapTask.Network,
				Token:   p.FibrousSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_WoofiSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewWoofiSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_WoofiSwapTask).WoofiSwapTask
			return NewWoofiSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_WoofiSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_WoofiSwapTask)")
			}
			return Marshal(t.WoofiSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_WoofiSwapTask{WoofiSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_WoofiSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_WoofiSwapTask)")
			}
			return &TaskInput{
				Network: p.WoofiSwapTask.Network,
				Token:   p.WoofiSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_WoofiSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_WoofiSwapTask)")
			}
			return &TaskOutput{
				Network: p.WoofiSwapTask.Network,
				Token:   p.WoofiSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_KyberSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewKyberSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_KyberSwapTask).KyberSwapTask
			return NewKyberSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_KyberSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_MintZeriusTask)")
			}
			return Marshal(t.KyberSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_KyberSwapTask{KyberSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_KyberSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_KyberSwapTask)")
			}
			return &TaskInput{
				Network: p.KyberSwapTask.Network,
				Token:   p.KyberSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_KyberSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_KyberSwapTask)")
			}
			return &TaskOutput{
				Network: p.KyberSwapTask.Network,
				Token:   p.KyberSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_EkuboSwap: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewEkuboSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_EkuboSwapTask).EkuboSwapTask
			return NewEkuboSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_EkuboSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_EkuboSwapTask)")
			}
			return Marshal(t.EkuboSwapTask)
		},
		CastR: func(in RandomTask) (*v1.Task, error) {
			return &v1.Task{
				TaskType: in.Type,
				Task:     &v1.Task_EkuboSwapTask{EkuboSwapTask: CastDefaultSwap(in.P)},
			}, nil
		},
		Input: func(m *v1.Task) (*TaskInput, error) {
			p, ok := m.Task.(*v1.Task_EkuboSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_EkuboSwapTask)")
			}
			return &TaskInput{
				Network: p.EkuboSwapTask.Network,
				Token:   p.EkuboSwapTask.FromToken,
			}, nil
		},
		Output: func(m *v1.Task) (*TaskOutput, error) {
			p, ok := m.Task.(*v1.Task_EkuboSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_EkuboSwapTask)")
			}
			return &TaskOutput{
				Network: p.EkuboSwapTask.Network,
				Token:   p.EkuboSwapTask.ToToken,
			}, nil
		},
	},
	v1.TaskType_WETH: {
		Payable: true,
		Tasker:  &Wrap{Tasker: NewWETHSwapTask()},
		Estimate: func(ctx context.Context, a EstimateArg) (*v1.EstimationTx, error) {
			p := a.Task.Task.(*v1.Task_WethSwapTask).WethSwapTask
			return NewWETHSwapTask().EstimateCost(ctx, a.Profile, p, nil)
		},
		Desc: func(m *v1.Task) ([]byte, error) {
			t, ok := m.Task.(*v1.Task_WethSwapTask)
			if !ok {
				return nil, errors.New("m.Task.(*v1.Task_WethSwapTask)")
			}
			return Marshal(t.WethSwapTask)
		},
	},
}

func CastDefaultSwap(p any) *v1.DefaultSwap {

	pa := p.(*v1.RPswapItem)

	return &v1.DefaultSwap{
		Amount:    pa.Am,
		Network:   pa.Network,
		FromToken: pa.From,
		ToToken:   pa.To,
	}
}

func CastDefaultSimple(p any) *v1.SimpleTask {

	pa := p.(*v1.RPsimple)

	return &v1.SimpleTask{
		Network: pa.Network,
	}
}

func GetTaskDesc(m *v1.Task) ([]byte, error) {
	v, ok := SpecMap[m.TaskType]
	if !ok {
		return nil, errors.New("unknown task type")
	}
	return v.Desc(m)
}

func GetTask(t v1.TaskType) (Tasker, error) {
	tasker, exist := SpecMap[t]
	if !exist {
		return nil, errors.New("unknown task: " + t.String())
	}
	return tasker.Tasker, nil
}

type Wrap struct {
	Tasker Tasker
}

func (w *Wrap) Stop() error {
	return w.Tasker.Stop()
}
func (w *Wrap) Type() v1.TaskType {
	return w.Tasker.Type()
}

type TaskUpdater interface {
	UpdateTask(ctx context.Context, req *v1.ProcessTask) error
}

func (d *Input) UpdateTask(ctx context.Context, req *v1.ProcessTask) error {
	return UpdateTask(ctx, req, d.ProcessRepository, d.PayService, d.User, d.UserId)
}

func UpdateTask(ctx context.Context, after *v1.ProcessTask, d repository.ProcessRepository, payService *pay.Service, user *repository.User, userId string) error {

	t, err := d.GetProcessTask(ctx, after.Id)
	if err != nil {
		return err
	}

	before, err := t.ToPB()
	if err != nil {
		return err
	}

	if after.Status != before.Status {
		if after.Error != nil &&
			after.Status != v1.ProcessStatus_StatusRunning &&
			after.Status != v1.ProcessStatus_StatusDone {
			_ = d.RecordStatusChanged(ctx, before.Id, before.Status, after.Status, *after.Error)
		} else {
			_ = d.RecordStatusChanged(ctx, before.Id, before.Status, after.Status)
		}
	}

	marshal, err := GetTaskDesc(after.Task)
	if err != nil {
		return errors.Wrap(err, "GetTaskDesc")
	}
	after.Task.Description = string(marshal)

	if err := d.UpdateProcessTask(ctx, after, before.Id, t.ProcessId, t.ProfileId); err != nil {
		return err
	}

	if NeedPay(before, after) {
		_, err = payService.FundsServiceClient.TaskCompleted(ctx, &paycli.TaskCompletedReq{
			ProcessId: t.ProcessId,
			ProfileId: t.ProfileId,
			TaskId:    before.Id,
			TaskType:  before.Task.TaskType.String(),
			UserId:    userId,
		})
		if err != nil {
			return err
		}

	}

	return nil
}

func (w *Wrap) Run(ctx context.Context, a *Input) (task *v1.ProcessTask, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("task panic: %+v", r))
			task = nil
		}
	}()

	if IsPayableTask(a.Task.Task.TaskType) {
		funds, err := a.PayService.FundsServiceClient.GetFunds(ctx, &paycli.GetFundsReq{
			Login: a.User.Email,
		})
		if err != nil {
			return nil, err
		}

		if funds.FundsLeft <= 0 {
			e := ErrUserHasNoBalance.Error()
			a.Task.Error = &e
			a.Task.Status = v1.ProcessStatus_StatusError

			if err := a.UpdateTask(ctx, a.Task); err != nil {
				return nil, err
			}

			return nil, ErrUserHasNoBalance
		}
	}

	taskId := a.Task.Id

	l := a.L.With("taskId", taskId).
		With("task", a.Task.Task.TaskType.String()).
		With("user_id", a.UserId)

	l.Debug("task running")
	task, err = w.Tasker.Run(ctx, a)
	if err != nil {
		l.Error(fmt.Sprintf("task [%s] finished with error ", a.Task.Task.TaskType.String()), zap.Error(err))
	} else {
		if err := a.ProcessRepository.UpdateProcessTask(ctx, a.Task, taskId, a.ProcessId, a.ProfileId); err != nil {
			l.Error("arg.ProcessRepository.UpdateProcessTask", err)
		}
		l.Debug("task finished")
	}

	if err != nil {
		// если транзакция еще не обработалась
		if errors.Is(err, ErrTransactionIsNotReady) {
			return a.Task, nil
		}
		e := err.Error()
		a.Task.Error = &e
		a.Task.Status = v1.ProcessStatus_StatusError

		if err := a.UpdateTask(ctx, a.Task); err != nil {
			return nil, err
		}
	} else {
		a.Task.Error = nil
		if err := a.UpdateTask(ctx, a.Task); err != nil {
			return nil, err
		}
	}

	return task, err
}

func IsPayableTask(t v1.TaskType) bool {

	if config.CFG.Env == config.Local {
		return true
	}

	return SpecMap[t].Payable
}

func NeedPay(before, after *v1.ProcessTask) bool {
	if before.Status == after.Status {
		return false
	}

	if !IsPayableTask(after.Task.TaskType) {
		return false
	}

	return after.Status == v1.ProcessStatus_StatusDone
}

func Marshal(m proto.Message) ([]byte, error) {
	return protojson.MarshalOptions{Multiline: true, UseEnumNumbers: false, UseProtoNames: false, EmitUnpopulated: true}.Marshal(m)
}
