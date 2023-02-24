package repository

import (
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/queries"
)

type Account interface {
	IsExistByEmail(email string) (bool, error)
	IsExistById(id int64) (bool, error)
	Registration(input *models.RegistrationAccountInput) (*queries.Account, error)
	Get(id int64) (*queries.Account, error)
	GetList(params *queries.GetAccountsParams) ([]queries.Account, error)
	Update(params *queries.UpdateAccountParams) (*queries.Account, error)
}

type Repository struct {
	Account
}

func NewRepository(db *queries.Queries) *Repository {
	return &Repository{
		Account: NewAccountPostgres(db),
	}
}
