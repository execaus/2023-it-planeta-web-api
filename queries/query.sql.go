// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package queries

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO "Account" (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id, first_name, last_name, email, password
`

type CreateAccountParams struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const getAccount = `-- name: GetAccount :one
SELECT id, first_name, last_name, email, password FROM "Account" WHERE id=$1
`

func (q *Queries) GetAccount(ctx context.Context, id int32) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const isExistAccount = `-- name: IsExistAccount :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE email = $1
)
`

func (q *Queries) IsExistAccount(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, isExistAccount, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
