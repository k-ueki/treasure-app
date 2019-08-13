-- +goose Up
alter table article add user_id int(10) unsigned default null;

-- +goose Down
alter table article drop user_id RESTRICT;