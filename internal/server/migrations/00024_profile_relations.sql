-- +goose Up
create table if not exists  profile_relations (
    p1_id uuid references profiles (id),
    p2_id uuid references profiles (id),
    p1_sub_type text not null,
    p2_sub_type text not null,
    p1_type text not null,
    p2_type text not null,
    unique (p1_id, p2_id, p1_sub_type, p2_sub_type, p1_type, p2_type)
);

-- +goose Down
drop table if exists profile_relations;