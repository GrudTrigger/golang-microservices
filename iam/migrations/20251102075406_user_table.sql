-- +goose Up
-- +goose StatementBegin
create table users (
    id serial primary key,
    login varchar not null,
    password varchar not null,
    email varchar not null unique,
    notification_methods JSONB
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users
-- +goose StatementEnd
