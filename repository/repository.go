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
	GetChippingLocation(animalID int64) (*queries.LocationPoint, error)
	GetCurrentLocation(animalID int64) (*queries.AnimalVisitedLocation, error)
	CreateVisitedLocation(animalID int64, pointID int64) (*queries.AnimalVisitedLocation, error)
	GetVisitedLocation(visitedPointID int64) (*queries.AnimalVisitedLocation, error)
	GetVisitedLocations(animalID int64) ([]queries.AnimalVisitedLocation, error)
	IsExistByID(animalID int64) (bool, error)
	IsExistVisitedLocationByID(visitedLocationID int64) (bool, error)
	IsLinkedVisitedLocation(animalID int64, visitedLocationPointID int64) (bool, error)
	UpdateVisitedLocation(visitedLocationPointID int64, locationPointID int64) (*queries.AnimalVisitedLocation, error)
	RemoveVisitedLocation(visitedLocationPointID int64) error
	GetVisitedLocationList(params *queries.GetVisitedLocationListParams) ([]queries.AnimalVisitedLocation, error)
	GetList(params *queries.GetAnimalsParams) ([]queries.Animal, error)
	Create(params *queries.CreateAnimalParams) (*queries.Animal, error)
	BindAnimalType(animalID int64, animalType int64) (*queries.AnimalToType, error)
	Update(params *queries.UpdateAnimalParams) (*queries.Animal, error)
	Remove(animalID int64) error
}

type AnimalType interface {
	GetByAnimalID(id int64) ([]queries.AnimalToType, error)
	IsExistByID(id int64) (bool, error)
	IsExistByType(animalType string) (bool, error)
	Create(animalType string) (*queries.AnimalType, error)
	GetByID(id int64) (*queries.AnimalType, error)
	Update(params *queries.UpdateAnimalTypeParams) (*queries.AnimalType, error)
	IsLinkedAnimal(id int64) (bool, error)
	Remove(id int64) (*queries.AnimalType, error)
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
