package v1

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/user"
)

func (s *FlowService) ShareFlow(ctx context.Context, req *v1.ShareFlowReq) (*v1.ShareFlowRes, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	if err := Wellcome(ctx, user.GroupPublisher); err != nil {
		return nil, err
	}

	flow, err := s.repository.GetFlow(ctx, userId, req.GetId())
	if err != nil {
		return nil, err
	}

	pb, err := flow.ToPB()
	if err != nil {
		return nil, err
	}
	for _, el := range pb.Tasks {
		switch el.TaskType {
		case v1.TaskType_WithdrawExchange:
			t, ok := el.Task.(*v1.Task_WithdrawExchangeTask)
			if ok {
				t.WithdrawExchangeTask.WithdrawerId = ""
				t.WithdrawExchangeTask.WithdrawAddr = nil
				t.WithdrawExchangeTask.WithdrawOrderId = nil
			}
		case v1.TaskType_ExchangeSwap:
			t, ok := el.Task.(*v1.Task_ExchangeSwapTask)
			if ok {
				t.ExchangeSwapTask.WithdrawerId = ""
			}
		}
	}

	var t repository.Flow
	if err := t.FromPB(pb, userId, int64(flow.Version)); err != nil {
		return nil, err
	}

	f := &repository.FlowShared{
		Id:          uuid.New().String(),
		ParentId:    t.Id,
		Label:       t.Label,
		Description: req.Description,
		Payload:     t.Payload,
		UserId:      userId,
		CreatedAt:   time.Now(),
		DeletedAt:   sql.NullTime{},
	}

	if err := s.flowSharedRep.CreateFlowShared(ctx, f); err != nil {
		return nil, err
	}
	return &v1.ShareFlowRes{Id: f.Id}, nil
}
func (s *FlowService) HideFlow(ctx context.Context, req *v1.HideFlowReq) (*v1.HideFlowRes, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	if err := Wellcome(ctx, user.GroupPublisher); err != nil {
		return nil, err
	}

	if err := s.flowSharedRep.DeleteFlowShared(ctx, userId, req.GetId()); err != nil {
		return nil, err
	}

	return &v1.HideFlowRes{}, nil
}
func (s *FlowService) SharedFlows(ctx context.Context, req *v1.SharedFlowsReq) (*v1.SharedFlowsRes, error) {

	items, err := s.flowSharedRep.ListSharedFlow(ctx)
	if err != nil {
		return nil, err
	}

	tmp := make([]*v1.FlowShared, 0)

	for _, i := range items {

		pb, err := i.ToPB()
		if err != nil {
			return nil, err
		}

		tmp = append(tmp, pb)
	}

	return &v1.SharedFlowsRes{Items: tmp}, err
}
func (s *FlowService) SharedFlow(ctx context.Context, req *v1.SharedFlowReq) (*v1.SharedFlowRes, error) {

	f, err := s.flowSharedRep.SharedFlow(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	pb, err := f.ToPB()
	if err != nil {
		return nil, err
	}

	return &v1.SharedFlowRes{Flow: pb}, nil
}
func (s *FlowService) UseSharedFlow(ctx context.Context, req *v1.UseSharedFlowReq) (*v1.UseSharedFlowRes, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	flow, err := s.flowSharedRep.SharedFlow(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	f := &repository.Flow{
		Id:        uuid.New().String(),
		Label:     flow.Label,
		Payload:   flow.Payload,
		NextId:    sql.NullString{},
		UserId:    userId,
		CreatedAt: time.Now(),
		DeletedAt: sql.NullTime{},
	}

	if err := s.repository.CreateFlow(ctx, f); err != nil {
		return nil, err
	}

	return &v1.UseSharedFlowRes{Id: f.Id}, nil
}
