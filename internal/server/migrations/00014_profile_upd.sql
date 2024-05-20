-- +goose Up
begin;
alter table if exists profiles
       add if not exists num integer default 0;


end;

-- +goose Down
alter table if exists profiles
       drop if exists num;