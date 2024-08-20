-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN IF NOT EXISTS api_key varchar(100) UNIQUE NOT NULL DEFAULT (encode(sha256(random()::text::bytea), 'hex'));
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN IF EXISTS api_key;
-- +goose StatementEnd