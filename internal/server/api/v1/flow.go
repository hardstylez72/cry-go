package v1

import (
	"context"

	"github.com/google/uuid"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/user"
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
	if err := f.FromPB(req.Flow, userId); err != nil {
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
		Id:        uuid.New().String(),
		Label:     req.Label,
		Tasks:     req.Tasks,
		CreatedAt: timestamppb.Now(),
	}

	a := &repository.Flow{}
	if err := a.FromPB(pb, userId); err != nil {
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
	tmp := make([]*v1.Flow, 0)
	for i := range out {

		p, err := out[i].ToPB()
		if err != nil {
			continue
		}

		tmp = append(tmp, p)
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
