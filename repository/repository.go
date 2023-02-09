package repository

import (
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/queries"
)

type Account interface {
	IsExist(email string) (bool, error)
	Registration(input *models.RegistrationAccountInput) (*queries.Account, error)
	Get(id int32) (*queries.Account, error)
}

type Repository struct {
	Account
}

func NewRepository(db *queries.Queries) *Repository {
	return &Repository{
		Account: NewAccountPostgres(db),
	}
}
