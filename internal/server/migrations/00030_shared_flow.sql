-- +goose Up
create table if not exists shared_flows (
    id uuid primary key,
    parent_id uuid references flow (id),
    label text not null,
    description text not null,
    payload text not null,
    user_id uuid not null references users (id),
    created_at timestamp not null default now(),
    deleted_at timestamp null
);

-- +goose Down
drop table if exists shared_flows;
