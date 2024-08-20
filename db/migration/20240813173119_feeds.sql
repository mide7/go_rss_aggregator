-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS feeds (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    "title" varchar(255) NOT NULL,
    "url" varchar(255) UNIQUE NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    "user_id" uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS feeds;
-- +goose StatementEnd