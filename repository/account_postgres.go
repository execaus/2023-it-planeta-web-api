package repository

import (
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/queries"
	"context"
	"database/sql"
	"github.com/execaus/exloggo"
)

type AccountPostgres struct {
	db *queries.Queries
}

func (r *AccountPostgres) IsLinkedAnimal(accountID int64) (bool, error) {
	isLinked, err := r.db.IsAccountLinkedAnimal(context.Background(), accountID)
	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}
	return isLinked, nil
}

func (r *AccountPostgres) IsExistByEmailExcept(email string, animalID int64) (bool, error) {
	isExist, err := r.db.IsExistAccountByEmailExcept(context.Background(), queries.IsExistAccountByEmailExceptParams{
		Email: email,
		ID:    animalID,
	})
	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}
	return isExist, nil
}

func (r *AccountPostgres) GetByEmail(login string) (*queries.Account, error) {
	account, err := r.db.GetAccountByEmail(context.Background(), login)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		exloggo.Error(err.Error())
		return nil, err
	}

	return &account, nil
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
