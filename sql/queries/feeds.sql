-- name: CreateFeed :one
INSERT INTO Feeds (ID, user_id, name, url, created_at, updated_at)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM Feeds;

-- name: GetFeedByURL :one
SELECT * FROM Feeds
WHERE url = $1;

-- name: GetIdByURL :one
SELECT ID FROM Feeds
WHERE url = $1
LIMIT 1;

-- name: GetFeedById :one
SELECT name FROM Feeds
WHERE ID = $1
LIMIT 1;

-- name: MarkFeedFetch :exec
UPDATE Feeds
SET 
  last_fetched_at = CURRENT_TIMESTAMP,
  updated_at = CURRENT_TIMESTAMP
WHERE ID = $1;

-- name: GetNextFeedToFetch :one
SELECT ID, name, url, last_fetched_at
FROM Feeds
ORDER BY last_fetched_at NULLS FIRST, updated_at ASC
LIMIT 1;
