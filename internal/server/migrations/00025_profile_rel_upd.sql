-- +goose Up
alter table if exists profile_relations
       add if not exists user_id uuid references users (id);

-- +goose Down
alter table if exists profile_relations
       drop column if exists user_id;
