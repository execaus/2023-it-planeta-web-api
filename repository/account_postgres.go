package repository

import (
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/queries"
	"context"
	"github.com/execaus/exloggo"
)

type AccountPostgres struct {
	db *queries.Queries
}

func (r *AccountPostgres) Remove(id int64) error {
	_, err := r.db.RemoveAccount(context.Background(), id)
	if err != nil {
		exloggo.Error(err.Error())
		return err
	}

	return nil
}

func (r *AccountPostgres) Update(params *queries.UpdateAccountParams) (*queries.Account, error) {
	account, err := r.db.UpdateAccount(context.Background(), *params)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &account, nil
}

func (r *AccountPostgres) GetList(params *queries.GetAccountsParams) ([]queries.Account, error) {
	accounts, err := r.db.GetAccounts(context.Background(), *params)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return accounts, nil
}

func (r *AccountPostgres) Get(id int64) (*queries.Account, error) {
	account, err := r.db.GetAccount(context.Background(), id)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &account, nil
}

func (r *AccountPostgres) IsExistByEmail(email string) (bool, error) {
	isExist, err := r.db.IsExistAccountByEmail(context.Background(), email)
	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}
	return isExist, nil
}

func (r *AccountPostgres) IsExistByID(id int64) (bool, error) {
	isExist, err := r.db.IsExistAccountByID(context.Background(), id)
	if err != nil {
		exloggo.Error(err.Error())
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
		exloggo.Error(err.Error())
		return nil, err
	}

	return &account, nil
}

func NewAccountPostgres(db *queries.Queries) *AccountPostgres {
	return &AccountPostgres{db: db}
}
