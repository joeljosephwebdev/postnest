-- name: CreateFeedFollow :one
INSERT INTO feeds_follows (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedFollowsForUser :many
SELECT 
feeds_follows.*,
users.name AS username,
feeds.name AS feedname,
feeds.url AS url
FROM feeds_follows
INNER JOIN users ON feeds_follows.user_id = users.id
INNER JOIN feeds ON feeds_follows.feed_id = feeds.id
WHERE users.name = $1;