package v1

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/task"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/user"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlowService struct {
	v1.UnimplementedFlowServiceServer
	repository    repository.FlowRepository
	flowSharedRep repository.FlowSharedRepository
}

func NewFlowService(repository repository.FlowRepository, flowSharedRep repository.FlowSharedRepository) *FlowService {
	return &FlowService{
		repository:    repository,
		flowSharedRep: flowSharedRep,
	}
}

func (s *FlowService) UpdateFlow(ctx context.Context, req *v1.UpdateFlowRequest) (*v1.UpdateFlowResponse, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	parentFlowId := req.Flow.Id

	req.Flow.Id = uuid.New().String()
	var f repository.Flow
	if err := f.FromPB(req.Flow, userId, 1); err != nil {
		return nil, err
	}

	if err := s.repository.UpdateFlow(ctx, parentFlowId, &f); err != nil {
		return nil, err
	}

	flow, err := s.repository.GetFlow(ctx, userId, req.Flow.Id)
	if err != nil {
		return nil, err
	}

	t, err := flow.ToPB()
	if err != nil {
		return nil, err
	}

	return &v1.UpdateFlowResponse{
		Flow: t,
	}, nil
}

func (s *FlowService) CreateFlow(ctx context.Context, req *v1.CreateFlowRequest) (*v1.CreateFlowResponse, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}
	pb := &v1.Flow{
		Id:          uuid.New().String(),
		Label:       req.Label,
		Tasks:       req.Tasks,
		CreatedAt:   timestamppb.Now(),
		RandomTasks: req.GetRandomTasks(),
	}

	a := &repository.Flow{}
	if err := a.FromPB(pb, userId, 1); err != nil {
		return nil, err
	}

	if err := s.repository.CreateFlow(ctx, a); err != nil {
		return nil, err
	}

	out, err := s.repository.GetFlow(ctx, userId, a.Id)
	if err != nil {
		return nil, err
	}

	pb, err = out.ToPB()
	if err != nil {
		return nil, err
	}

	return &v1.CreateFlowResponse{Flow: pb}, nil
}

func (s *FlowService) ListFlow(ctx context.Context, req *v1.ListFlowRequest) (*v1.ListFlowResponse, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	out, err := s.repository.ListFlows(ctx, userId)
	if err != nil {
		return nil, err
	}
	tmp := make([]*v1.FlowListItem, 0)
	for i := range out {

		tmp = append(tmp, &v1.FlowListItem{
			Id:        out[i].Id,
			Label:     out[i].Label,
			NextId:    &out[i].NextId.String,
			CreatedAt: timestamppb.New(out[i].CreatedAt),
			DeletedAt: timestamppb.New(out[i].DeletedAt.Time),
			Version:   int64(out[i].Version),
		})
	}

	return &v1.ListFlowResponse{
		Flows: tmp,
	}, nil
}
func (s *FlowService) DeleteFlow(ctx context.Context, req *v1.DeleteFlowRequest) (*v1.DeleteFlowResponse, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	if err := s.repository.DeleteFlow(ctx, userId, req.Id); err != nil {
		return nil, err
	}
	return &v1.DeleteFlowResponse{}, nil
}

func (s *FlowService) GetFlow(ctx context.Context, req *v1.GetFlowRequest) (*v1.GetFlowResponse, error) {
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	t, err := s.repository.GetFlow(ctx, userId, req.Id)
	if err != nil {
		return nil, err
	}
	pb, err := t.ToPB()
	if err != nil {
		return nil, err
	}
	return &v1.GetFlowResponse{Flow: pb}, nil
}

func (s *FlowService) CopyFlow(ctx context.Context, req *v1.CopyFlowReq) (*v1.CopyFlowRes, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	flow, err := s.repository.GetFlow(ctx, userId, req.GetId())
	if err != nil {
		return nil, err
	}

	flow.Id = uuid.New().String()
	flow.Label = flow.Label + " [копия]"

	err = s.repository.CreateFlow(ctx, flow)
	if err != nil {
		return nil, err
	}

	return &v1.CopyFlowRes{
		Id: flow.Id,
	}, nil
}

func (s *FlowService) CreateFlowV2(ctx context.Context, req *v1.CreateFlowV2Req) (*v1.CreateFlowV2Res, error) {

	payload, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	flow := &repository.Flow{
		Id:        uuid.New().String(),
		Label:     req.Label,
		Payload:   string(payload),
		NextId:    sql.NullString{},
		UserId:    userId,
		CreatedAt: time.Now(),
		DeletedAt: sql.NullTime{},
		Version:   2,
	}

	if err := s.repository.CreateFlow(ctx, flow); err != nil {
		return nil, err
	}

	return &v1.CreateFlowV2Res{Id: flow.Id}, nil
}

func (s *FlowService) GetFlowV2(ctx context.Context, req *v1.GetFlowV2Req) (*v1.GetFlowV2Res, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	flow, err := s.repository.GetFlow(ctx, userId, req.Id)
	if err != nil {
		return nil, err
	}

	var tmp v1.CreateFlowV2Req

	if err := protojson.Unmarshal([]byte(flow.Payload), &tmp); err != nil {
		return nil, err
	}

	return &v1.GetFlowV2Res{
		Id:     flow.Id,
		Label:  flow.Label,
		Blocks: tmp.Blocks,
	}, nil
}

func (s *FlowService) UpdateFlowV2(ctx context.Context, req *v1.UpdateFlowV2Request) (*v1.UpdateFlowV2Response, error) {
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	parentFlowId := req.Id

	var p v1.CreateFlowV2Req
	p.Label = req.Label
	p.Blocks = req.Blocks

	payload, err := protojson.Marshal(&p)
	if err != nil {
		return nil, err
	}
	req.Id = uuid.New().String()
	f := repository.Flow{
		Id:        req.Id,
		Label:     req.Label,
		Payload:   string(payload),
		NextId:    sql.NullString{},
		UserId:    userId,
		CreatedAt: time.Now(),
		DeletedAt: sql.NullTime{},
		Version:   2,
	}

	if err := s.repository.UpdateFlow(ctx, parentFlowId, &f); err != nil {
		return nil, err
	}

	out, err := s.GetFlowV2(ctx, &v1.GetFlowV2Req{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateFlowV2Response{
		Id:     out.Id,
		Label:  out.Label,
		Blocks: out.Blocks,
	}, nil
}

func (s *FlowService) OnlyRandomFlowPreview(ctx context.Context, req *v1.OnlyRandomFlowPreviewReq) (*v1.OnlyRandomFlowPreviewRes, error) {

	block := &v1.FlowBlock_Rand{
		Rand: &v1.FlowBlockRand{
			StartToken:   req.StartToken,
			FinishToken:  req.FinishToken,
			StartNetwork: req.StartNetwork,
			Tasks:        req.Tasks,
			TaskCount:    req.TaskCount,
		},
	}

	dreams, err := task.RandomizeRandomBlock(block, req.IgnoreStartToken, req.IgnoreFinishToken, req.MinDelay, req.MaxDelay)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	uf := make([]*v1.UniqueFlow, 0)
	for _, tasks := range dreams {
		uf = append(uf, &v1.UniqueFlow{
			Tasks: tasks,
		})
	}

	tokens := task.TokenCombo(dreams)

	if len(uf) > 2 {
		uf = uf[:3]
	}

	return &v1.OnlyRandomFlowPreviewRes{
		Flow:          uf,
		UniquePercent: float64(len(dreams)),
		Tokens:        tokens,
	}, nil
}

func (s *FlowService) OnlyRandomFlowFromTokens(ctx context.Context, req *v1.OnlyRandomFlowPreviewReq) (*v1.OnlyRandomFlowFromTokensRes, error) {

	tasks := task.ExtractItems(req.Tasks, req.StartNetwork, nil, nil, false)

	m := map[v1.Token]bool{}
	tokens := make([]v1.Token, 0)
	for _, randomTask := range tasks {

		if randomTask.FromToken() == nil {
			continue
		}

		m[*randomTask.FromToken()] = true
	}

	for token, _ := range m {
		tokens = append(tokens, token)
	}

	return &v1.OnlyRandomFlowFromTokensRes{
		Tokens: tokens,
	}, nil
}

func (s *FlowService) FlowPreview(ctx context.Context, req *v1.FlowPreviewReq) (*v1.FlowPreviewRes, error) {

	dreams, err := LetTheDreamsComeTrue(req.GetBlocks())
	if err != nil {
		return nil, err
	}
	uf := make([]*v1.UniqueFlow, 0)
	for _, tasks := range dreams {
		uf = append(uf, &v1.UniqueFlow{
			Tasks: tasks,
		})
	}

	if len(uf) > 2 {
		uf = uf[:3]
	}

	return &v1.FlowPreviewRes{
		Flow:          uf,
		UniquePercent: 1,
	}, nil
}

func LetTheDreamsComeTrue(blocks []*v1.FlowBlock) ([][]*v1.Task, error) {
	flows := [][][]*v1.Task{}
	randomBlocksCount := 0
	manBlocksCount := 0

	for _, block := range blocks {

		switch (block.Block).(type) {
		case *v1.FlowBlock_Rand:
			b := block.Block.(*v1.FlowBlock_Rand)

			cases, err := task.RandomizeRandomBlock(b, false, false, b.Rand.MinDelay, b.Rand.MaxDelay)
			if err != nil {
				return nil, err
			}

			flows = append(flows, cases)
			randomBlocksCount++

		case *v1.FlowBlock_Man:
			b := block.Block.(*v1.FlowBlock_Man)

			cases := [][]*v1.Task{}
			if len(b.Man.RandomTasks) == 0 {
				cases = [][]*v1.Task{b.Man.Tasks}
			} else {
				for _, randomTask := range b.Man.RandomTasks {
					cases = append(cases, task.InsertRandom[*v1.Task](b.Man.Tasks, randomTask))
				}
			}

			flows = append(flows, cases)

			manBlocksCount++
		}
	}

	maxSize := 0
	for _, flow := range flows {
		if len(flow) > maxSize {
			maxSize = len(flow)
		}
	}

	out := make([][]*v1.Task, maxSize)
	for _, flow := range flows {

		flow = task.RandomN(flow, maxSize)

		for i, tasks := range flow {
			out[i] = append(out[i], tasks...)
		}
	}

	return out, nil
}
