-- name: CreateFeedFollow :one
INSERT INTO feed_follows (user_id, feed_id)
VALUES ($1, $2)
RETURNING *;
-- name: ListFeedFollows :many
SELECT feed_follows.id,
    feed_follows.user_id,
    feed_follows.feed_id,
    feed_follows.created_at,
    feed_follows.updated_at,
    json_build_object('id', users.id, 'name', users.name) AS user
FROM feed_follows
    INNER JOIN users ON users.id = feed_follows.user_id
ORDER BY feed_follows.id
LIMIT $1 OFFSET $2;
-- name: ListFeedFollowsByUserID :many
SELECT *
FROM feed_follows
WHERE feed_follows.user_id = $1
ORDER BY feed_follows.id
LIMIT $2 OFFSET $3;
-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE id = $1
    AND user_id = $2;