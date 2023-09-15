package v1

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IssueService struct {
	v1.UnimplementedIssueServiceServer
	r                 repository.IssueRepository
	processRepository repository.ProcessRepository
	userRepository    repository.UserRepository
}

func NewIssueService(r repository.IssueRepository, processRepository repository.ProcessRepository, userRepository repository.UserRepository) *IssueService {
	return &IssueService{
		r:                 r,
		processRepository: processRepository,
		userRepository:    userRepository,
	}
}

func (s *IssueService) Issues(ctx context.Context, req *v1.IssuesReq) (*v1.IssuesRes, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	tmp, err := s.r.Issues(ctx, &repository.Offset{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*v1.Issue, len(tmp))
	for i := range tmp {
		t := tmp[i]
		out[i] = &v1.Issue{
			Status:      v1.IssueStatus(v1.IssueStatus_value[t.Status]),
			Id:          t.Id,
			CreatedAt:   timestamppb.New(t.CreatedAt),
			FinishedAt:  nil,
			CreatedBy:   t.CreatorId,
			Title:       t.Title,
			Description: t.Description,
			My:          t.CreatorId == userId,
		}

		if t.FinishedAt.Valid {
			out[i].FinishedAt = timestamppb.New(t.FinishedAt.Time)
		}
	}
	return &v1.IssuesRes{
		Items: out,
	}, nil

}
func (s *IssueService) IssuesUser(ctx context.Context, req *v1.IssuesReq) (*v1.IssuesRes, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}
	tmp, err := s.r.IssuesByUser(ctx, userId, &repository.Offset{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*v1.Issue, len(tmp))
	for i := range tmp {
		t := tmp[i]
		out[i] = &v1.Issue{
			Status:      v1.IssueStatus(v1.IssueStatus_value[t.Status]),
			Id:          t.Id,
			CreatedAt:   timestamppb.New(t.CreatedAt),
			FinishedAt:  nil,
			CreatedBy:   t.CreatorId,
			Title:       t.Title,
			Description: t.Description,
			My:          t.CreatorId == userId,
		}

		if t.FinishedAt.Valid {
			out[i].FinishedAt = timestamppb.New(t.FinishedAt.Time)
		}
	}
	return &v1.IssuesRes{
		Items: out,
	}, nil

}
func (s *IssueService) CreateIssue(ctx context.Context, req *v1.CreateIssueReq) (*v1.CreateIssueRes, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	i := &repository.Issue{
		Id:          uuid.New().String(),
		CreatorId:   userId,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   time.Now(),
		FinishedAt:  sql.NullTime{},
		Status:      v1.IssueStatus_Created.String(),
	}

	if req.GetProcessId() != "" && req.GetTaskid() != "" {
		task, err := s.processRepository.GetProcessTask(ctx, req.GetTaskid())
		if err != nil {
			return nil, err
		}
		t, err := task.ToPB()
		if err != nil {
			return nil, err
		}

		if t.Error != nil {
			i.Title = "[" + t.Task.TaskType.String() + "] " + *t.Error
		}

		i.Description = req.Description + " \n\r " + t.Task.Description

	}

	if err := s.r.CreateIssue(ctx, i); err != nil {
		return nil, err
	}
	return &v1.CreateIssueRes{}, nil
}
func (s *IssueService) DeleteIssue(ctx context.Context, req *v1.DeleteIssueReq) (*v1.DeleteIssueRes, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	if err := s.r.DeleteIssue(ctx, userId, req.IssueId); err != nil {
		return nil, err
	}

	return &v1.DeleteIssueRes{}, nil
}

func (s *IssueService) Issue(ctx context.Context, req *v1.IssueReq) (*v1.IssueRes, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	t, err := s.r.Issue(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	i := &v1.Issue{
		Status:      v1.IssueStatus(v1.IssueStatus_value[t.Status]),
		Id:          t.Id,
		CreatedAt:   timestamppb.New(t.CreatedAt),
		FinishedAt:  nil,
		CreatedBy:   t.CreatorId,
		Title:       t.Title,
		Description: t.Description,
		My:          t.CreatorId == userId,
	}

	if t.FinishedAt.Valid {
		i.FinishedAt = timestamppb.New(t.FinishedAt.Time)
	}

	comments, err := s.r.IssueComments(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	tmp := make([]*v1.IssueComment, len(comments))
	for i, c := range comments {
		tmp[i] = &v1.IssueComment{
			Id:          c.Id,
			IssueId:     c.IssueId,
			CreatorId:   c.UserId,
			Description: c.Description,
			CreatedAt:   timestamppb.New(c.CreatedAt),
		}
	}

	return &v1.IssueRes{
		Issue:    i,
		Comments: tmp,
	}, nil
}
func (s *IssueService) IssueUpdateStatus(ctx context.Context, req *v1.IssueUpdateStatusReq) (*v1.IssueUpdateStatusRes, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	groups, err := user.Groups(ctx, s.userRepository)
	if err != nil {
		return nil, err
	}
	
	if !groups.Get(user.GroupSupport) {
		return nil, status.New(codes.PermissionDenied, "user is not in support group").Err()
	}

	if err := s.r.UpdateStatus(ctx, req.IssueId, req.Status.String(), userId); err != nil {
		return nil, err
	}
	return &v1.IssueUpdateStatusRes{}, nil
}
func (s *IssueService) AddComment(ctx context.Context, req *v1.AddCommentReq) (*v1.AddCommentRes, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	c := &repository.IssueComment{
		Id:          uuid.New().String(),
		IssueId:     req.IssueId,
		Description: req.Text,
		UserId:      userId,
		CreatedAt:   time.Now(),
	}

	if err := s.r.CreateIssueComment(ctx, c); err != nil {
		return nil, err
	}

	return &v1.AddCommentRes{}, nil
}
func (s *IssueService) DeleteComment(ctx context.Context, req *v1.DeleteCommentReq) (*v1.DeleteCommentRes, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	if err := s.r.DeleteIssueComment(ctx, userId, req.CommentId); err != nil {
		return nil, err
	}
	return &v1.DeleteCommentRes{}, nil
}
