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
