package service

import (
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/queries"
	"2023-it-planeta-web-api/repository"
)

type AccountService struct {
	repo repository.Account
}

func (s *AccountService) Get(id int) (*queries.Account, error) {
	idNumber64 := int64(id)
	return s.repo.Get(idNumber64)
}

func (s *AccountService) IsExistByEmail(email string) (bool, error) {
	return s.repo.IsExistByEmail(email)
}

func (s *AccountService) IsExistById(id int) (bool, error) {
	idNumber64 := int64(id)
	return s.repo.IsExistById(idNumber64)
}

func (s *AccountService) Registration(input *models.RegistrationAccountInput) (*models.RegistrationAccountOutput, error) {
	account, err := s.repo.Registration(input)
	if err != nil {
		return nil, err
	}

	return &models.RegistrationAccountOutput{
		Id:        account.ID,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Email:     account.Email,
	}, nil
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}
