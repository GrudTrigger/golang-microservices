-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders
ALTER COLUMN part_uuid TYPE text[]
USING ARRAY[part_uuid];
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
