-- +goose Up
alter table article add constraint article_fk_user foreign key (user_id) references user(id);

-- +goose Down
alter table drop constraint article_fk_user;