-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostsForUser :many
WITH feeds_id AS (
  SELECT feed_id FROM feeds_follows WHERE feeds_follows.user_id = $1
)
SELECT * FROM posts 
INNER JOIN feeds_follows ON  posts.feed_id = feeds_follows.feed_id
WHERE posts.feed_id IN (SELECT feed_id FROM feeds_id)
ORDER BY published_at DESC
LIMIT $2;