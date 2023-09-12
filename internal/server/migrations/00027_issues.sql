-- +goose Up
create table if not exists  issues (
    id uuid primary key,
    creator_id uuid references users (id),
    title text not null,
    description text not null,
    created_at timestamp not null default now(),
    finished_at timestamp null,
    status text not null
);

create  table if not exists issue_comments (
       id uuid primary key,
       issue_id uuid  references issues (id),
       creator_id uuid references users (id),
       description text not null,
       created_at timestamp not null default now()
);
-- +goose Down
alter table if exists issue_comments
       drop column if exists sub_type;

alter table if exists issues
       drop column if exists sub_type;
