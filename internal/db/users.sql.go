// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, FirstName, LastName, PasswordHash)
VALUES ($1, $2, $3, $4)
RETURNING id, username, firstname, lastname, passwordhash
`

type CreateUserParams struct {
	Username     string
	Firstname    string
	Lastname     string
	Passwordhash string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Firstname,
		arg.Lastname,
		arg.Passwordhash,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Firstname,
		&i.Lastname,
		&i.Passwordhash,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, firstname, lastname, passwordhash FROM users WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Firstname,
		&i.Lastname,
		&i.Passwordhash,
	)
	return i, err
}

const searchUsers = `-- name: SearchUsers :many
SELECT username, FirstName, LastName FROM users WHERE username LIKE $1
`

type SearchUsersRow struct {
	Username  string
	Firstname string
	Lastname  string
}

// seach for users by username
func (q *Queries) SearchUsers(ctx context.Context, username string) ([]SearchUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, searchUsers, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchUsersRow
	for rows.Next() {
		var i SearchUsersRow
		if err := rows.Scan(&i.Username, &i.Firstname, &i.Lastname); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const userExists = `-- name: UserExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)
`

func (q *Queries) UserExists(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, userExists, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}