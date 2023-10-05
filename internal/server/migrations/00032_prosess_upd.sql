-- +goose Up
alter table if exists process
       add if not exists run_after timestamp null;

-- +goose Down
alter table if exists process
       drop column if exists run_after;
