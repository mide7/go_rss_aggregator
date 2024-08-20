-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS posts (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    "title" varchar(500) NOT NULL,
    "description" text,
    "published_at" timestamptz NOT NULL,
    "url" varchar(500) NOT NULL UNIQUE,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    "feed_id" uuid NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS posts;
-- +goose StatementEnd