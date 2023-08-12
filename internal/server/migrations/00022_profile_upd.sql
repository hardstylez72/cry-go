-- +goose Up
ALTER TABLE profiles DROP CONSTRAINT IF EXISTS check_num_uniq;
alter table if exists profiles
       add constraint check_num_uniq
              unique (user_id, num, "type");


-- +goose Down
alter table if exists profiles
       drop constraint if exists check_num_uniq;
