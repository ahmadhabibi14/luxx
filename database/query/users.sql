-- name: CreateNewUser :exec
INSERT INTO Users (
  user_id, username, full_name, email, password
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: GetUserByUsername :one
SELECT username FROM Users
WHERE username = ?;

-- name: ListUsers :many
SELECT username, full_name, email, avatar, join_at FROM Users
ORDER BY join_at DESC;

-- name: GetUserDataByUserId :one
SELECT username, full_name, email, avatar, join_at FROM Users
WHERE user_id = ?;