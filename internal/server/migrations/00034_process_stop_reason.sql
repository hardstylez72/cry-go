-- +goose Up
alter table if exists process
    add if not exists stop_reason text null;
-- +goose Down
alter table if exists process
    drop if exists stop_reason;

