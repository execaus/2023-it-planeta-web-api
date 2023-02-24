package service

import (
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/queries"
	"2023-it-planeta-web-api/repository"
	"database/sql"
	"github.com/sirupsen/logrus"
)

const (
	accountGetListDefaultLimit  = 10
	accountGetListDefaultOffset = 0
)

type AccountService struct {
	repo repository.Account
}

func (s *AccountService) Remove(id int64) error {
	return s.repo.Remove(id)
}

func (s *AccountService) Update(id int64, input *models.UpdateAccountInput) (*queries.Account, error) {
	params := queries.UpdateAccountParams{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
		ID:        id,
	}
	return s.repo.Update(&params)
}

func (s *AccountService) GetList(input *models.GetAccountsInput) ([]*models.GetAccountsOutput, error) {
	var limit int32
	var offset int32

	if input.Size == nil {
		limit = accountGetListDefaultLimit
	}

	if input.From == nil {
		offset = accountGetListDefaultOffset
	}

	params := &queries.GetAccountsParams{
		Column1: sql.NullString{
			String: input.FirstName,
			Valid:  true,
		},
		Column2: sql.NullString{
			String: input.LastName,
			Valid:  true,
		},
		Column3: sql.NullString{
			String: input.Email,
			Valid:  true,
		},
		Limit:  limit,
		Offset: offset,
	}

	repositoryAccounts, err := s.repo.GetList(params)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	accounts := make([]*models.GetAccountsOutput, 0)
	for _, account := range repositoryAccounts {
		accounts = append(accounts, &models.GetAccountsOutput{
			ID:        account.ID,
			FirstName: account.FirstName,
			LastName:  account.LastName,
			Email:     account.Email,
		})
	}

	return accounts, nil
}

func (s *AccountService) Get(id int64) (*queries.Account, error) {
	return s.repo.Get(id)
}

func (s *AccountService) IsExistByEmail(email string) (bool, error) {
	return s.repo.IsExistByEmail(email)
}

func (s *AccountService) IsExistByID(id int64) (bool, error) {
	return s.repo.IsExistByID(id)
}

func (s *AccountService) Registration(
	input *models.RegistrationAccountInput,
) (*models.RegistrationAccountOutput, error) {
	account, err := s.repo.Registration(input)
	if err != nil {
		return nil, err
	}

	return &models.RegistrationAccountOutput{
		ID:        account.ID,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Email:     account.Email,
	}, nil
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}
