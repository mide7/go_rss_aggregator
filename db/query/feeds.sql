-- name: CreateFeed :one
INSERT INTO feeds (title, url, user_id)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetFeedByID :one
SELECT *
FROM feeds
WHERE id = $1
LIMIT 1;
-- name: ListFeeds :many
SELECT *
FROM feeds
ORDER BY id
LIMIT $1 OFFSET $2;
-- name: GetNextFeedsToFetch :many
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;
-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fetched_at = now(),
    updated_at = now()
WHERE id = $1
RETURNING *;