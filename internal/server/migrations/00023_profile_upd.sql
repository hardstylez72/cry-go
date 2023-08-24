-- +goose Up
alter table if exists profiles
       add if not exists sub_type text not null default 'Metamask';

alter table if exists profiles
       add if not exists seed bytea not null default '\000'::bytea;

-- +goose Down
alter table if exists profiles
       drop column if exists sub_type;

alter table if exists profiles
       drop column if exists seed;