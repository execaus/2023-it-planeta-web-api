package repository

import "2023-it-planeta-web-api/queries"

type AccountPostgres struct {
	db *queries.Queries
}

func NewAccountPostgres(db *queries.Queries) *AccountPostgres {
	return &AccountPostgres{db: db}
}
