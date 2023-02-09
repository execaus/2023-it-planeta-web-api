package service

import (
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/queries"
	"2023-it-planeta-web-api/repository"
)

type Account interface {
	IsExist(email string) (bool, error)
	Registration(input *models.RegistrationAccountInput) (*models.RegistrationAccountOutput, error)
	Get(id int) (*queries.Account, error)
}

type Service struct {
	Account
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Account: NewAccountService(repos.Account),
	}
}
