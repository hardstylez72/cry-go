-- +goose Up
alter table if exists okex_deposit_addr_profile
       add if not exists sub_type text default 'Metamask';

-- +goose Down
alter table if exists okex_deposit_addr_profile
       drop column if exists sub_type;
