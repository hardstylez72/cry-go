-- +goose Up

create index if not exists process_tasks_profile_id ON process_tasks (profile_id);
create index if not exists process_profiles_profile_id ON process_profiles  (profile_id) ;
create index  if not exists process_profiles_process_id ON process_profiles  (process_id) ;

-- +goose Down

