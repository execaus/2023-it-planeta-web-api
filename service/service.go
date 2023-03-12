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
	IsDead(id int64) (bool, error)
	GetChippingLocation(animalID int64) (*queries.LocationPoint, error)
	GetCurrentLocation(animalID int64) (*queries.AnimalVisitedLocation, error)
	CreateVisitedLocation(animalID int64, pointID int64) (*queries.AnimalVisitedLocation, error)
}

type AnimalType interface {
	GetByAnimalID(id int64) ([]queries.AnimalToType, error)
	IsExistByID(id int64) (bool, error)
	IsExistByType(animalType string) (bool, error)
	Create(animalType string) (*queries.AnimalType, error)
	GetByID(id int64) (*queries.AnimalType, error)
	Update(id int64, animalType string) (*queries.AnimalType, error)
	IsLinkedAnimal(id int64) (bool, error)
	Remove(id int64) error
}

type Location interface {
	GetVisitedAnimal(id int64) ([]queries.AnimalVisitedLocation, error)
	Get(id int64) (*queries.LocationPoint, error)
	IsExistByID(id int64) (bool, error)
	IsExistByCoordinates(latitude float64, longitude float64) (bool, error)
	Create(latitude float64, longitude float64) (*queries.LocationPoint, error)
	Update(id int64, latitude float64, longitude float64) (*queries.LocationPoint, error)
	IsLinkedAnimal(id int64) (bool, error)
	Remove(id int64) error
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
