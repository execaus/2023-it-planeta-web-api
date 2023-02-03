package service

import "2023-it-planeta-web-api/repository"

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}
