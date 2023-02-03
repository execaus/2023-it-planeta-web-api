package repository

import "2023-it-planeta-web-api/queries"

type Account interface {
}

type Repository struct {
	Account
}

func NewRepository(db *queries.Queries) *Repository {
	return &Repository{
		Account: NewAccountPostgres(db),
	}
}
