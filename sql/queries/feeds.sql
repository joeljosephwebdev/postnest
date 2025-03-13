-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT 
feeds.id AS id,
feeds.created_at AS createdAt,
feeds.updated_at AS updatedAt,
feeds.name AS name,
feeds.url AS URL,
users.name AS username
FROM feeds
INNER JOIN users
ON feeds.user_id = users.id;