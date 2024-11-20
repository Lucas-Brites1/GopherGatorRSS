-- name: CreatePost :one
INSERT INTO Posts(ID, title, description, url, feed_id, published_at, created_at, updated_at)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8
)
RETURNING *;

-- name: GetPostsForUser :many
SELECT *
FROM Posts as P
INNER JOIN Feeds as F ON P.feed_id = F.ID
INNER JOIN Feed_Follows FF ON FF.feed_id = F.ID
WHERE FF.user_id = $1
ORDER BY P.published_at DESC
LIMIT $2;