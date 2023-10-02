-- +goose Up
ALTER TABLE profiles DROP CONSTRAINT IF EXISTS check_num_uniq;

create unique index if not exists check_num_deleted_at_uniq
       on profiles (user_id, num, "type", (deleted_at is null))
       where deleted_at is null;
-- +goose Down

