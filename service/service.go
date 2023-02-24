package service

import (
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/queries"
	"2023-it-planeta-web-api/repository"
)

type Account interface {
	IsExistByEmail(email string) (bool, error)
	IsExistByID(id int64) (bool, error)
	Registration(input *models.RegistrationAccountInput) (*models.RegistrationAccountOutput, error)
	Get(id int64) (*queries.Account, error)
	GetList(input *models.GetAccountsInput) ([]*models.GetAccountsOutput, error)
	Update(id int64, input *models.UpdateAccountInput) (*queries.Account, error)
	Remove(id int64) error
}

type Animal interface {
	Get(id int64) (*queries.Animal, error)
}

type AnimalType interface {
	GetFromAnimal(id int64) ([]queries.AnimalToType, error)
}

type Location interface {
	GetVisitedAnimal(id int64) ([]queries.AnimalVisitedLocation, error)
}

type Service struct {
	Account
	Animal
	AnimalType
	Location
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Account:    NewAccountService(repos.Account),
		Animal:     NewAnimalService(repos.Animal),
		AnimalType: NewAnimalTypeService(repos.AnimalType),
		Location:   NewLocationService(repos.Location),
	}
}
