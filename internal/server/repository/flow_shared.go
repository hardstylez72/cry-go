package repository

import (
	"context"
	"database/sql"
	"time"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlowShared struct {
	Id          string       `db:"id"`
	ParentId    string       `db:"parent_id"`
	Label       string       `db:"label"`
	Description string       `db:"description"`
	Payload     string       `db:"payload"`
	UserId      string       `db:"user_id"`
	CreatedAt   time.Time    `db:"created_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

var flowSharedCols = []string{
	"id",
	"parent_id",
	"label",
	"description",
	"payload",
	"user_id",
	"created_at",
	"deleted_at",
}

type FlowSharedRepository interface {
	CreateFlowShared(ctx context.Context, req *FlowShared) error
	ListSharedFlow(ctx context.Context) ([]FlowShared, error)
	SharedFlow(ctx context.Context, id string) (*FlowShared, error)
	DeleteFlowShared(ctx context.Context, userId, flowId string) error
}

var (
	fsc = NewHelper(flowSharedCols)
)

func (a *FlowShared) ToPB() (*v1.FlowShared, error) {

	var f v1.FlowShared
	parser := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err := parser.Unmarshal([]byte(a.Payload), &f); err != nil {
		return nil, err
	}
	p := &v1.FlowShared{
		Id:          a.Id,
		Label:       a.Label,
		Tasks:       f.Tasks,
		CreatedAt:   timestamppb.New(a.CreatedAt),
		ParentId:    a.ParentId,
		Description: a.Description,
		CreatorId:   a.UserId,
	}

	if a.DeletedAt.Valid {
		p.DeletedAt = timestamppb.New(a.DeletedAt.Time)
	}

	return p, nil
}

func (r *pgRepository) CreateFlowShared(ctx context.Context, req *FlowShared) error {

	q := `insert into shared_flows (` + fsc.Cols() + `) values (` + fsc.ColsColon() + `)                                                                 `
	if _, err := r.conn.NamedExecContext(ctx, q, req); err != nil {
		return err
	}
	return nil

	return nil
}

func (r *pgRepository) DeleteFlowShared(ctx context.Context, userId, flowId string) error {
	q := `update shared_flows set deleted_at = now() where user_id = $1 and id = $2`

	if _, err := r.conn.ExecContext(ctx, q, userId, flowId); err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) ListSharedFlow(ctx context.Context) ([]FlowShared, error) {

	q := "select " + fsc.Cols() + " from shared_flows order by created_at desc"
	out := make([]FlowShared, 0)
	if err := r.conn.SelectContext(ctx, &out, q); err != nil {
		return nil, err
	}

	return out, nil
}

func (r *pgRepository) SharedFlow(ctx context.Context, id string) (*FlowShared, error) {

	q := "select " + fsc.Cols() + " from shared_flows where id = $1"
	var out FlowShared
	if err := r.conn.GetContext(ctx, &out, q, id); err != nil {
		return nil, err
	}

	return &out, nil
}
