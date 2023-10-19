-- name: CreateUser :one
INSERT INTO users (username, FirstName, LastName, PasswordHash)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UserExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE username = $1);

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- seach for users by username
-- name: SearchUsers :many
SELECT username, FirstName, LastName FROM users WHERE username LIKE $1;
