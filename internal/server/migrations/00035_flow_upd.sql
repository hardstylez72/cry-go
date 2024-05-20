-- +goose Up
alter table if exists flow
       add if not exists version integer not null default 1;
-- +goose Down
alter table if exists flow
       drop if exists version;