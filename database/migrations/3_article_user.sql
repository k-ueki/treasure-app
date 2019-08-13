-- +goose Up
alter table article add user_id int(11);

-- +goose Down
alter table article drop user_id RESTRICT;