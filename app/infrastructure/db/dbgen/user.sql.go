// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: user.sql

package dbgen

import (
	"context"
)

const findByID = `-- name: FindByUserID :one
SELECT id, email, password, name, profile, created_at, updated_at
FROM users
WHERE id = ?
`

func (q *Queries) FindByID(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, findByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Name,
		&i.Profile,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findByName = `-- name: FindByName :one
SELECT id, email, password, name, profile, created_at, updated_at
FROM users
WHERE name = ?
`

func (q *Queries) FindByName(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, findByName, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Name,
		&i.Profile,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertUser = `-- name: UpsertUser :exec
INSERT INTO users (id,
                   email,
                   password,
                   name,
                   profile,
                   created_at,
                   updated_at)
VALUES (?,
        ?,
        ?,
        ?,
        ?,
        NOW(),
        NOW()) ON DUPLICATE KEY
UPDATE
    email = ?,
    password = ?,
    name = ?,
    profile = ?,
    updated_at = NOW()
`

type UpsertUserParams struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Profile  string `json:"profile"`
}

func (q *Queries) UpsertUser(ctx context.Context, arg UpsertUserParams) error {
	_, err := q.db.ExecContext(ctx, upsertUser,
		arg.ID,
		arg.Email,
		arg.Password,
		arg.Name,
		arg.Profile,
		arg.Email,
		arg.Password,
		arg.Name,
		arg.Profile,
	)
	return err
}