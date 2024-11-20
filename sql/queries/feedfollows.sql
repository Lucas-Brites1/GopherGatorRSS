-- name: CreateFeedFollow :one
INSERT INTO Feed_Follows (ID, user_id, feed_id, created_at, updated_at)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;

-- name: GetFeedFollowsForUser :many
SELECT 
  U.name AS user_name,
  Fe.name AS feed_name,
  Fe.url AS feed_url
FROM 
  Feed_Follows as F
INNER JOIN
  Users AS U ON F.user_id = U.ID
INNER JOIN 
  Feeds AS Fe ON F.feed_id = Fe.ID
WHERE
  U.ID = $1;

-- name: Unfollow :exec 
DELETE FROM Feed_Follows 
WHERE feed_id = (SELECT ID FROM Feeds WHERE url = $1 LIMIT 1)
AND Feed_Follows.user_id = $2;