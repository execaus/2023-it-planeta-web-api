package service

import (
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/queries"
	"2023-it-planeta-web-api/repository"
)

type Account interface {
	IsExistByEmail(email string) (bool, error)
	IsExistById(id int64) (bool, error)
	Registration(input *models.RegistrationAccountInput) (*models.RegistrationAccountOutput, error)
	Get(id int64) (*queries.Account, error)
	GetList(input *models.GetAccountsInput) ([]*models.GetAccountsOutput, error)
	Update(id int64, input *models.UpdateAccountInput) (*queries.Account, error)
}

type Service struct {
	Account
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Account: NewAccountService(repos.Account),
	}
}
