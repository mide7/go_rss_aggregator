-- +goose Up
-- +goose StatementBegin
ALTER TABLE feeds
ADD COLUMN IF NOT EXISTS last_fetched_at timestamptz;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE feeds DROP COLUMN IF EXISTS last_fetched_at;
-- +goose StatementEnd