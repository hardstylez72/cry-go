-- +goose Up
alter table if exists process
    add if not exists parallel boolean not null default false;
-- +goose Down
alter table if exists process
    drop if exists parallel;

