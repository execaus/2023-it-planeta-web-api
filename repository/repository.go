package repository

import (
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/queries"
)

type Account interface {
	IsExistByEmail(email string) (bool, error)
	IsExistByID(id int64) (bool, error)
	Registration(input *models.RegistrationAccountInput) (*queries.Account, error)
	Get(id int64) (*queries.Account, error)
	GetList(params *queries.GetAccountsParams) ([]queries.Account, error)
	Update(params *queries.UpdateAccountParams) (*queries.Account, error)
	Remove(id int64) error
}

type Animal interface {
	Get(id int64) (*queries.Animal, error)
}

type AnimalType interface {
	GetByAnimalID(id int64) ([]queries.AnimalToType, error)
	IsExist(id int64) (bool, error)
	GetById(id int64) (*queries.AnimalType, error)
}

type Location interface {
	GetVisitedAnimal(id int64) ([]queries.AnimalVisitedLocation, error)
	Get(id int64) (*queries.LocationPoint, error)
	IsExistByID(id int64) (bool, error)
	IsExistByCoordinates(params *queries.IsExistLocationByCoordinatesParams) (bool, error)
	Create(params *queries.CreateLocationParams) (*queries.LocationPoint, error)
	Update(params *queries.UpdateLocationParams) (*queries.LocationPoint, error)
	IsVisitedAnimal(id int64) (bool, error)
	IsAnimalChipping(id int64) (bool, error)
	Remove(id int64) error
}

type Repository struct {
	Account
	Animal
	AnimalType
	Location
}

func NewRepository(db *queries.Queries) *Repository {
	return &Repository{
		Account:    NewAccountPostgres(db),
		Animal:     NewAnimalPostgres(db),
		AnimalType: NewAnimalTypePostgres(db),
		Location:   NewLocationPostgres(db),
	}
}
