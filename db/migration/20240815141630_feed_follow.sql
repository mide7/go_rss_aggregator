-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS feed_follows (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    "user_id" uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    "feed_id" uuid NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    UNIQUE (user_id, feed_id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS feed_follows;
-- +goose StatementEnd