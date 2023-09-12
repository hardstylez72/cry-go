package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/hardstylez72/cry/internal/server/repository/pg"
	"github.com/jmoiron/sqlx"
)

type IssueStatus = string

type Issue struct {
	Id          string       `db:"id"`
	CreatorId   string       `db:"creator_id"`
	Title       string       `db:"title"`
	Description string       `db:"description"`
	CreatedAt   time.Time    `db:"created_at"`
	FinishedAt  sql.NullTime `db:"finished_at"`
	Status      string       `db:"status"`
}

type IssueComment struct {
	Id          string    `db:"id"`
	IssueId     string    `db:"issue_id"`
	Description string    `db:"description"`
	UserId      string    `db:"creator_id"`
	CreatedAt   time.Time `db:"created_at"`
}

type Offset struct {
	Limit  int64
	Offset int64
}

type IssueRepository interface {
	IssuesByUser(ctx context.Context, userId string, o *Offset) ([]Issue, error)
	Issues(ctx context.Context, o *Offset) ([]Issue, error)
	CreateIssue(ctx context.Context, i *Issue) error
	DeleteIssue(ctx context.Context, userId, issueId string) error
	UpdateStatus(ctx context.Context, issueId, status string) error
	Issue(ctx context.Context, issueId string) (*Issue, error)

	IssueComments(ctx context.Context, issueId string) ([]IssueComment, error)
	CreateIssueComment(ctx context.Context, i *IssueComment) error
	DeleteIssueComment(ctx context.Context, userId, issueId string) error
}

var IssueCommentCols = []string{
	"id",
	"issue_id",
	"creator_id",
	"description",
	"created_at",
}

var IssueCols = []string{
	"id",
	"creator_id",
	"title",
	"description",
	"created_at",
	"finished_at",
	"status",
}

var (
	issueHelper   = NewHelper(IssueCols)
	commentHelper = NewHelper(IssueCommentCols)
)

type issueRepository struct {
	conn    *sqlx.DB
	issue   *Helper
	comment *Helper
}

func NewIssueRepository(conn *sqlx.DB) *issueRepository {
	return &issueRepository{
		conn:    conn,
		issue:   issueHelper,
		comment: commentHelper,
	}
}

func (db *issueRepository) UpdateStatus(ctx context.Context, issueId, status string) error {
	if _, err := db.conn.ExecContext(ctx, "update issues set status = $1 where id = $2 ", status, issueId); err != nil {
		return err
	}
	return nil
}

func (db *issueRepository) IssuesByUser(ctx context.Context, userId string, o *Offset) ([]Issue, error) {
	q := Join(`select `, db.issue.Cols(), ` from issues where creator_id = $1 order by created_at desc`)
	var out []Issue
	if err := db.conn.SelectContext(ctx, &out, q, userId); err != nil {
		return nil, err
	}
	return out, nil
}

func (db *issueRepository) Issues(ctx context.Context, o *Offset) ([]Issue, error) {
	q := Join(`select `, db.issue.Cols(), ` from issues order by created_at desc`)
	var out []Issue
	if err := db.conn.SelectContext(ctx, &out, q); err != nil {
		return nil, err
	}
	return out, nil
}

func (db *issueRepository) Issue(ctx context.Context, issueId string) (*Issue, error) {
	q := Join(`select `, db.issue.Cols(), ` from issues where id = $1`)
	var out Issue
	if err := db.conn.GetContext(ctx, &out, q, issueId); err != nil {
		return nil, err
	}
	return &out, nil
}

func (db *issueRepository) CreateIssue(ctx context.Context, i *Issue) error {
	q := Join("insert into issues (", db.issue.Cols(), `) values (`, db.issue.ColsColon(), ")")

	if _, err := db.conn.NamedExecContext(ctx, q, i); err != nil {
		return pg.PgError(err)
	}
	return nil
}

func (db *issueRepository) DeleteIssue(ctx context.Context, userId, issueId string) error {
	if _, err := db.conn.ExecContext(ctx, "delete from issues where id = $1 and creator_id = $2", issueId, userId); err != nil {
		return err
	}
	return nil
}

func (db *issueRepository) IssueComments(ctx context.Context, issueId string) ([]IssueComment, error) {
	q := Join(`select `, db.comment.Cols(), ` from issue_comments where issue_id = $1 order by created_at asc`)
	var out = make([]IssueComment, 0)
	if err := db.conn.SelectContext(ctx, &out, q, issueId); err != nil {
		return nil, err
	}
	return out, nil
}

func (db *issueRepository) CreateIssueComment(ctx context.Context, i *IssueComment) error {
	q := Join("insert into issue_comments (", db.comment.Cols(), `) values (`, db.comment.ColsColon(), ")")

	if _, err := db.conn.NamedExecContext(ctx, q, i); err != nil {
		return pg.PgError(err)
	}
	return nil
}

func (db *issueRepository) DeleteIssueComment(ctx context.Context, userId, issueId string) error {
	if _, err := db.conn.ExecContext(ctx, "delete from issue_comments where id = $1 and creator_id = $2", issueId, userId); err != nil {
		return err
	}
	return nil
}
