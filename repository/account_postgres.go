package repository

import (
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/queries"
	"context"
	"github.com/sirupsen/logrus"
)

type AccountPostgres struct {
	db *queries.Queries
}

func (r *AccountPostgres) IsExist(email string) (bool, error) {
	isExist, err := r.db.IsExistAccount(context.Background(), email)
	if err != nil {
		logrus.Error(err.Error())
		return false, err
	}
	return isExist, nil
}

func (r *AccountPostgres) Registration(input *models.RegistrationAccountInput) (*queries.Account, error) {
	account, err := r.db.CreateAccount(context.Background(), queries.CreateAccountParams{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	})
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &account, nil
}

func NewAccountPostgres(db *queries.Queries) *AccountPostgres {
	return &AccountPostgres{db: db}
}
