-- +goose Up
create table if not exists settings_v2 (
       key text not null  default 'unknown',
       user_id uuid references users (id),
       payload text not null,
       updated_at timestamp not null default now(),
       unique (user_id, key)
);

-- +goose Down

alter table if exists settings_v2
       drop if exists num;