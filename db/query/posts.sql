-- name: CreatePost :one
INSERT INTO posts (title, description, published_at, url, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: GetPostByID :one
SELECT *
FROM posts
WHERE id = $1
LIMIT 1;
-- name: ListPosts :many
SELECT *
FROM posts
ORDER BY id
LIMIT $1 OFFSET $2;
-- name: ListPostsForUser :many
SELECT posts.*
FROM posts
    JOIN feed_follows ON feed_follows.feed_id = posts.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2 OFFSET $3;