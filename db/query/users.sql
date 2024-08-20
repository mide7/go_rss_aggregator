-- name: CreateUser :one
INSERT INTO users (name)
VALUES ($1)
RETURNING *;
-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;
-- name: GetUserByAPIKey :one
SELECT *
FROM users
WHERE api_key = $1
LIMIT 1;
-- name: ListUsers :many
SELECT *
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;