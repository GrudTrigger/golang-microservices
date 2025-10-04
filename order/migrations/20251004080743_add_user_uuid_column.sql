-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders ADD COLUMN user_uuid varchar;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
