-- +goose Up
alter table if exists issues
       add if not exists deleted_at timestamp null;

-- +goose Down
alter table if exists issues
       drop column if exists deleted_at;
