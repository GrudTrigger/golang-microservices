-- +goose Up
-- +goose StatementBegin
create table orders (
    id serial primary key,
    part_uuid varchar not null,
    total_price float not null,
    transaction_uuid varchar,
    payment_method varchar,
    status varchar not null,
    created_at timestamp not null default now(),
    updated_at timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table order
-- +goose StatementEnd
