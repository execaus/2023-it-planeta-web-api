package service

import "2023-it-planeta-web-api/repository"

type Account interface {
}

type Service struct {
	Account
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Account: NewAccountService(repos.Account),
	}
}
