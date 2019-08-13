-- +goose Up
alter table article add column user_id int(11);

-- -goose Down
alter table article drop column user_id RESTRICT;