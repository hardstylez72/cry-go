-- +goose Up
alter table if exists users
       add if not exists controlled_by uuid null references users (id);

-- +goose Down
alter table if exists users
       drop column if exists controlled_by;
