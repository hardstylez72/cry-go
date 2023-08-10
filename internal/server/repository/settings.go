package repository

import (
	"context"
	"database/sql"
	"time"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
)

var ErrNotFound = errors.New("not found")

func (r *pgRepository) GetSettings(ctx context.Context, userId string, network v1.Network) (*v1.NetworkSettings, error) {
	var payload string
	if err := r.conn.GetContext(ctx, &payload, "select payload from settings_v2 where user_id = $1 and key = $2", userId, network.String()); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(ErrNotFound, "settings not found")
		}
		return nil, err
	}

	var s v1.NetworkSettings
	if err := protojson.Unmarshal([]byte(payload), &s); err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *pgRepository) UpdateSettings(ctx context.Context, userId string, request *v1.NetworkSettings) error {
	p, err := protojson.Marshal(request)
	if err != nil {
		return err
	}

	s := string(p)

	if _, err := r.conn.ExecContext(ctx, `insert into settings_v2 (user_id, payload, key) values ($1, $2, $3) on conflict (user_id, key) do update set payload = $2`, userId, &s, request.GetNetwork().String()); err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) GetSettingsDate(ctx context.Context, userId string, network v1.Network) (*time.Time, error) {
	var updatedAt time.Time
	if err := r.conn.GetContext(ctx, &updatedAt, "select updated_at from settings_v2 where user_id = $1 and key = $2", userId, network.String()); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(ErrNotFound, "settings date not found")
		}
		return nil, err
	}

	return &updatedAt, nil
}
