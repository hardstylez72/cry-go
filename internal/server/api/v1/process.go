package v1

import (
	"context"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process"
	"github.com/hardstylez72/cry/internal/process/task"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/user"
	"github.com/hardstylez72/cry/internal/settings"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProcessService struct {
	v1.UnimplementedProcessServiceServer
	processRepository repository.ProcessRepository
	flowRepository    repository.FlowRepository
	dispatcher        *process.Dispatcher
	settingsService   *settings.Service
}

func NewProcessService(
	processRepository repository.ProcessRepository,
	dispatcher *process.Dispatcher,
	flowRepository repository.FlowRepository,
	settingsService *settings.Service,
) *ProcessService {
	return &ProcessService{
		processRepository: processRepository,
		dispatcher:        dispatcher,
		flowRepository:    flowRepository,
		settingsService:   settingsService,
	}
}

func (s *ProcessService) EnableAutoRetry(ctx context.Context, req *v1.EnableAutoRetryRequest) (*v1.EnableAutoRetryResponse, error) {
	err := s.processRepository.UpdateProcessAutoRetry(ctx, req.ProcessId, true)
	if err != nil {
		return nil, err
	}
	return &v1.EnableAutoRetryResponse{}, nil
}
func (s *ProcessService) DisableAutoRetry(ctx context.Context, req *v1.DisableAutoRetryRequest) (*v1.DisableAutoRetryResponse, error) {
	err := s.processRepository.UpdateProcessAutoRetry(ctx, req.ProcessId, false)
	if err != nil {
		return nil, err
	}
	return &v1.DisableAutoRetryResponse{}, nil
}
func (s *ProcessService) SkipProcessTask(ctx context.Context, req *v1.SkipProcessTaskRequest) (*v1.SkipProcessTaskResponse, error) {
	if err := s.dispatcher.SkipTask(ctx, req.ProcessId, req.ProfileId, req.TaskId); err != nil {
		return nil, err
	}
	return &v1.SkipProcessTaskResponse{}, nil
}
func (s *ProcessService) CreateProcess(ctx context.Context, req *v1.CreateProcessRequest) (*v1.CreateProcessResponse, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	f, err := s.flowRepository.GetFlow(ctx, userId, req.FlowId)
	if err != nil {
		return nil, err
	}

	flow, err := f.ToPB()
	if err != nil {
		return nil, err
	}

	profiles := make([]*v1.ProcessProfile, 0)

	for profileWeight, profileId := range req.ProfileIds {

		tasks := []v1.Task{}
		for i := range flow.Tasks {
			tasks = append(tasks, *flow.Tasks[i])
		}

		randomTasks := []v1.Task{}
		for i := range flow.RandomTasks {
			randomTasks = append(randomTasks, *flow.RandomTasks[i])
		}

		processTasks, err := addRandomTasks(tasks, randomTasks)
		if err != nil {
			return nil, errors.Wrap(err, "lib.Marshal")
		}

		for i := range processTasks {
			println(processTasks[i].Task.TaskType.String() + strconv.Itoa(int(processTasks[i].Task.Weight)))
		}

		profiles = append(profiles, &v1.ProcessProfile{
			ProfileId: profileId,
			Id:        uuid.New().String(),
			Weight:    int64(profileWeight),
			Tasks:     processTasks,
			Status:    v1.ProcessStatus_StatusReady,
		})
	}

	status := v1.ProcessStatus_StatusStop
	if req.RunAfter != nil {
		status = v1.ProcessStatus_StatusReady
	}

	var pb = &v1.Process{
		Id:         uuid.New().String(),
		Status:     status,
		Profiles:   profiles,
		FlowId:     req.FlowId,
		CreatedAt:  timestamppb.Now(),
		UpdatedAt:  timestamppb.Now(),
		FinishedAt: nil,
		StartedAt:  nil,
		RunAfter:   req.RunAfter,
	}

	a := &repository.ProcessArg{}
	if err := a.FromPB(pb, userId); err != nil {
		return nil, err
	}

	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()

	if err := s.processRepository.CreateProcess(ctx, a); err != nil {
		return nil, err
	}

	res, err := s.processRepository.GetProcessArg(ctx, &v1.GetProcessRequest{Id: a.Id}, userId)
	if err != nil {
		return nil, err
	}

	return &v1.CreateProcessResponse{
		Process: res.Process,
	}, nil
}

func addRandomTasks(tasks, randomTasks []v1.Task) ([]*v1.ProcessTask, error) {

	// preparation
	randomProcessTasks := make([]*v1.ProcessTask, 0, len(randomTasks))

	for i := range randomTasks {
		t := randomTasks[i]

		marshal, err := task.GetTaskDesc(&t)
		if err != nil {
			return nil, errors.Wrap(err, "lib.Marshal")
		}
		description := string(marshal)
		t.Description = description

		randomProcessTasks = append(randomProcessTasks, &v1.ProcessTask{
			Id:           uuid.New().String(),
			Task:         &t,
			Status:       v1.ProcessStatus_StatusReady,
			Transactions: []string{},
			FinishedAt:   nil,
			Error:        nil,
			StartedAt:    nil,
		})
	}

	processTasks := make([]*v1.ProcessTask, 0, len(randomTasks)+len(tasks))
	for i := range tasks {
		t := tasks[i]

		marshal, err := task.GetTaskDesc(&t)
		if err != nil {
			return nil, errors.Wrap(err, "lib.Marshal")
		}
		description := string(marshal)
		t.Description = description

		processTasks = append(processTasks, &v1.ProcessTask{
			Id:         uuid.New().String(),
			Task:       &t,
			Status:     v1.ProcessStatus_StatusReady,
			FinishedAt: nil,
			Error:      nil,
			StartedAt:  nil,
		})
	}

	// random

	capacity := len(randomTasks) + len(tasks)
	out := make([]*v1.ProcessTask, 0, capacity)

	m := map[int]*v1.ProcessTask{}
	for _, randomTask := range randomProcessTasks {

		for {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			position := r.Intn(capacity)
			_, exist := m[position]
			if !exist {
				m[position] = randomTask
				break
			}
		}
	}

	j := 0
	for i := 0; i < capacity; i++ {

		if j >= len(processTasks) {
			for i2, randomTask := range m {
				if i2 >= i {
					out = append(out, randomTask)
				}
			}
			break
		}
		t := processTasks[j]

		randomTask, needInsert := m[i]
		if needInsert {
			out = append(out, randomTask)
		} else {
			out = append(out, t)
			j++
		}
	}

	for i := range out {
		out[i].Task.Weight = int64(i)
	}

	return out, nil
}

func (s *ProcessService) GetProcess(ctx context.Context, req *v1.GetProcessRequest) (*v1.GetProcessResponse, error) {
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	return s.processRepository.GetProcessArg(ctx, req, userId)
}
func (s *ProcessService) ListProcess(ctx context.Context, req *v1.ListProcessRequest) (*v1.ListProcessResponse, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	statuses := make([]string, len(req.Statuses))
	for _, s := range req.Statuses {
		statuses = append(statuses, s.String())
	}

	processes, err := s.processRepository.ListProcessByUser(ctx, userId, statuses, int(req.Offset), 15)
	if err != nil {
		return nil, err
	}

	weight := map[v1.ProcessStatus]int{
		v1.ProcessStatus_StatusDone:    1,
		v1.ProcessStatus_StatusError:   99,
		v1.ProcessStatus_StatusRunning: 97,
		v1.ProcessStatus_StatusStop:    98,
		v1.ProcessStatus_StatusRetry:   3,
		v1.ProcessStatus_StatusReady:   4,
	}

	sort.Slice(processes, func(i, j int) bool {
		return weight[processes[i].Status]-weight[processes[j].Status] > 0
	})

	return &v1.ListProcessResponse{
		Processes: processes,
	}, nil
}
func (s *ProcessService) RetryProcess(ctx context.Context, req *v1.RetryProcessRequest) (*v1.RetryProcessResponse, error) {

	err := s.dispatcher.RetryProcessProfile(ctx, req.ProcessId, req.ProfileId, req.TaskId)
	if err != nil {
		return nil, err
	}

	return &v1.RetryProcessResponse{}, nil
}
func (s *ProcessService) GetProcessUpdatedAt(ctx context.Context, req *v1.GetProcessUpdatedAtRequest) (*v1.GetProcessUpdatedAtResponse, error) {

	at, err := s.processRepository.GetProcessUpdatedAt(ctx, req.GetProcessId())
	if err != nil {
		return nil, err
	}

	return &v1.GetProcessUpdatedAtResponse{
		UpdatedAt: timestamppb.New(*at),
	}, nil
}
func (s *ProcessService) GetProcessTaskHistory(ctx context.Context, req *v1.GetProcessTaskHistoryRequest) (*v1.GetProcessTaskHistoryResponse, error) {
	records, err := s.processRepository.GetProcessTaskHistory(ctx, req.GetTaskId())
	if err != nil {
		return nil, err
	}

	return &v1.GetProcessTaskHistoryResponse{
		Records: records,
	}, nil
}
func (s *ProcessService) StopProcess(ctx context.Context, req *v1.StopProcessRequest) (*v1.StopProcessResponse, error) {
	if err := s.dispatcher.StopProcess(ctx, req.ProcessId); err != nil {
		return nil, err
	}

	return &v1.StopProcessResponse{}, nil
}
func (s *ProcessService) ResumeProcess(ctx context.Context, req *v1.ResumeProcessRequest) (*v1.ResumeProcessResponse, error) {
	s.dispatcher.ResumeProcess(ctx, req.ProcessId)
	return &v1.ResumeProcessResponse{}, nil
}
func (s *ProcessService) CancelProcess(ctx context.Context, req *v1.CancelProcessRequest) (*v1.CancelProcessResponse, error) {

	if err := s.dispatcher.KillProcess(ctx, req.ProcessId); err != nil {
		return nil, err
	}

	return &v1.CancelProcessResponse{}, nil
}
func (s *ProcessService) EstimateCost(ctx context.Context, req *v1.EstimateCostRequest) (*v1.EstimateCostResponse, error) {
	etimations, err := s.dispatcher.EstimateTaskCost(ctx, req.ProfileId, req.TaskId)
	if err != nil {
		e := err.Error()
		return &v1.EstimateCostResponse{
			Error: &e,
		}, nil
	}
	return &v1.EstimateCostResponse{
		Error:       nil,
		Estimations: etimations,
	}, nil
}

func (s *ProcessService) GetTaskTransactions(ctx context.Context, req *v1.GetTaskTransactionsReq) (*v1.GetTaskTransactionsRes, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.processRepository.TransactionsByTaskId(ctx, userId, req.TaskId)
	if err != nil {
		return nil, err
	}

	tmp := make([]*v1.Transaction, len(res))

	for i := range res {
		tmp[i] = res[i].ToPB()
	}

	return &v1.GetTaskTransactionsRes{
		Transactions: tmp,
	}, nil
}
func (s *ProcessService) GetProfileTransactions(ctx context.Context, req *v1.GetProfileTransactionsReq) (*v1.GetProfileTransactionsRes, error) {
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}
	res, err := s.processRepository.TransactionsByProfileId(ctx, userId, req.ProfileId)
	if err != nil {
		return nil, err
	}

	tmp := make([]*v1.Transaction, len(res))

	for i := range res {
		tmp[i] = res[i].ToPB()
	}

	return &v1.GetProfileTransactionsRes{
		Transactions: tmp,
	}, nil
}

func (s *ProcessService) StopAllProcess(ctx context.Context, req *v1.StopAllProcessReq) (*v1.StopAllProcessRes, error) {
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}
	ids, err := s.processRepository.ProcessIDsUser(ctx, userId, v1.ProcessStatus_StatusRunning)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	wg.Add(len(ids))

	for _, id := range ids {
		go func(id string) {
			defer wg.Done()
			_ = s.dispatcher.StopProcess(ctx, id)
		}(id)
	}

	wg.Wait()

	return &v1.StopAllProcessRes{}, nil
}

func (s *ProcessService) ResumeAllProcess(ctx context.Context, req *v1.ResumeAllProcessReq) (*v1.ResumeAllProcessRes, error) {
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}
	ids, err := s.processRepository.ProcessIDsUser(ctx, userId, v1.ProcessStatus_StatusStop)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	wg.Add(len(ids))

	for _, id := range ids {
		go func(id string) {
			defer wg.Done()
			_ = s.dispatcher.ResumeProcess(ctx, id)
		}(id)
	}

	wg.Wait()

	return &v1.ResumeAllProcessRes{}, nil
}
