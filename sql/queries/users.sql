-- name: CreateUser :one
INSERT INTO Users (ID, created_at, updated_at, name)
VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM Users 
WHERE name = $1 LIMIT 1;

-- name: GetNameByID :one
SELECT name FROM Users
WHERE ID = $1 LIMIT 1;

-- name: GetIdByName :one
SELECT ID FROM Users
WHERE name = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM Users;

-- name: Reset :exec
DELETE FROM Users;

