-- +goose Up
alter table if exists profiles
       add if not exists type text not null default 'EVM';
-- +goose Down
alter table if exists profiles
       drop if exists type;