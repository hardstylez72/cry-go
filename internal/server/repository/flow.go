package repository

import (
	"context"
	"database/sql"
	"time"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository/pg"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Flow struct {
	Id        string         `db:"id"`
	Label     string         `db:"label"`
	Payload   string         `db:"payload"`
	NextId    sql.NullString `db:"next_id"`
	UserId    string         `db:"user_id"`
	CreatedAt time.Time      `db:"created_at"`
	DeletedAt sql.NullTime   `db:"deleted_at"`
	Version   int            `db:"version"`
}

func (a *Flow) FromPB(pb *v1.Flow, userId string, version int64) error {
	a.Id = pb.Id
	a.Label = pb.Label
	b, err := protojson.Marshal(pb)
	if err != nil {
		return err
	}
	a.Payload = string(b)
	a.CreatedAt = pb.CreatedAt.AsTime()
	a.UserId = userId
	if pb.DeletedAt != nil {
		a.DeletedAt.Valid = true
		a.DeletedAt.Time = pb.DeletedAt.AsTime()
	}
	a.Version = int(version)

	return nil
}

func (a *Flow) ToPB() (*v1.Flow, error) {

	var f v1.Flow
	if err := protojson.Unmarshal([]byte(a.Payload), &f); err != nil {
		return nil, err
	}
	p := &v1.Flow{
		Id:          a.Id,
		Label:       a.Label,
		Tasks:       f.Tasks,
		RandomTasks: f.RandomTasks,
		CreatedAt:   timestamppb.New(a.CreatedAt),
	}

	if a.NextId.Valid {
		p.NextId = &a.NextId.String
	}

	if a.DeletedAt.Valid {
		p.DeletedAt = timestamppb.New(a.DeletedAt.Time)
	}

	return p, nil
}

func (r *pgRepository) UpdateFlow(ctx context.Context, parentFlowId string, req *Flow) (err error) {

	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}

	}()

	txx := pg.WrapSqlxTx(tx)

	if err := createFlow(ctx, r.conn, req); err != nil {
		return err
	}
	if err := updateFlowParent(ctx, txx, parentFlowId, req.Id); err != nil {
		return err
	}

	return nil
}

func (r *pgRepository) CreateFlow(ctx context.Context, req *Flow) error {

	if err := createFlow(ctx, r.conn, req); err != nil {
		return err
	}

	return nil
}

func updateFlowParent(ctx context.Context, driver pg.SqlDriver, flowId, nextFlowId string) error {
	q := `update flow set next_id = $1 where id = $2                                                                 `
	if _, err := driver.ExecContext(ctx, q, nextFlowId, flowId); err != nil {
		return err
	}
	return nil
}

func createFlow(ctx context.Context, driver pg.SqlDriver, req *Flow) error {
	q := `insert into flow (id, label, payload, created_at, user_id, next_id, version) values 
      (:id, :label, :payload, :created_at, :user_id, null, :version)                                                                 `
	if _, err := driver.NamedExecContext(ctx, q, req); err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) GetFlow(ctx context.Context, userId, flowId string) (*Flow, error) {
	res, err := getFlow(ctx, r.conn, flowId, userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *pgRepository) GetFlowLabel(ctx context.Context, userId, flowId string) (*string, error) {
	q := `select label from flow where id = $1 and user_id = $2`
	var a string
	if err := r.conn.GetContext(ctx, &a, q, flowId, userId); err != nil {
		return nil, err
	}
	return &a, nil
}

func getFlow(ctx context.Context, conn *sqlx.DB, id, userId string) (*Flow, error) {
	q := `select * from flow where id = $1 and user_id = $2`
	var a Flow
	if err := conn.GetContext(ctx, &a, q, id, userId); err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *pgRepository) ListFlows(ctx context.Context, userId string) ([]Flow, error) {

	res, err := listFlows(ctx, r.conn, userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func listFlows(ctx context.Context, conn *sqlx.DB, userId string) ([]Flow, error) {
	q := "select * from flow where user_id = $1 and next_id is null and deleted_at is null order by created_at desc"
	out := make([]Flow, 0)
	if err := conn.SelectContext(ctx, &out, q, userId); err != nil {
		return nil, err
	}

	return out, nil
}
func (r *pgRepository) DeleteFlow(ctx context.Context, userId, flowId string) error {
	if _, err := r.conn.ExecContext(ctx, "update flow set deleted_at = now() where id = $1 and user_id = $2", flowId, userId); err != nil {
		return err
	}
	return nil
}

type FlowMeta struct {
	Id      string `db:"id"`
	Label   string `db:"label"`
	Version int    `db:"version"`
}

func (r *pgRepository) FlowMeta(ctx context.Context, userId, flowId string) (*FlowMeta, error) {

	var meta FlowMeta

	if err := r.conn.GetContext(ctx, &meta, "select id, label, version from flow where id = $1 and user_id = $2", flowId, userId); err != nil {
		return nil, err
	}
	return &meta, nil
}
