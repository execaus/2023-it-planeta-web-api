package repository

import (
	"2023-it-planeta-web-api/queries"
	"context"
	"github.com/sirupsen/logrus"
)

type AnimalPostgres struct {
	db *queries.Queries
}

func (r *AnimalPostgres) Update(params *queries.UpdateAnimalParams) (*queries.Animal, error) {
	animal, err := r.db.UpdateAnimal(context.Background(), *params)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &animal, nil
}

func (r *AnimalPostgres) BindAnimalType(animalID int64, animalType int64) (*queries.AnimalToType, error) {
	row, err := r.db.BindAnimalTypeToAnimal(context.Background(), queries.BindAnimalTypeToAnimalParams{
		Animal:     animalID,
		AnimalType: animalType,
	})
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &row, err
}

func (r *AnimalPostgres) Create(params *queries.CreateAnimalParams) (*queries.Animal, error) {
	animal, err := r.db.CreateAnimal(context.Background(), *params)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &animal, nil
}

func (r *AnimalPostgres) GetList(params *queries.GetAnimalsParams) ([]queries.Animal, error) {
	animals, err := r.db.GetAnimals(context.Background(), *params)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return animals, nil
}

func (r *AnimalPostgres) GetVisitedLocationList(
	params *queries.GetVisitedLocationListParams) ([]queries.AnimalVisitedLocation, error) {
	visitedLocations, err := r.db.GetVisitedLocationList(context.Background(), *params)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return visitedLocations, nil
}

func (r *AnimalPostgres) RemoveVisitedLocation(visitedLocationPointID int64) error {
	_, err := r.db.RemoveVisitedLocation(context.Background(), visitedLocationPointID)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	return nil
}

func (r *AnimalPostgres) UpdateVisitedLocation(
	visitedLocationPointID int64, locationPointID int64) (*queries.AnimalVisitedLocation, error) {
	point, err := r.db.UpdateVisitedLocation(context.Background(), queries.UpdateVisitedLocationParams{
		Location: locationPointID,
		ID:       visitedLocationPointID,
	})
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &point, nil
}

func (r *AnimalPostgres) IsLinkedVisitedLocation(animalID int64, visitedLocationPointID int64) (bool, error) {
	isLinked, err := r.db.IsLinkedAnimalToVisitedLocation(
		context.Background(),
		queries.IsLinkedAnimalToVisitedLocationParams{
			ID:     visitedLocationPointID,
			Animal: animalID,
		})
	if err != nil {
		logrus.Error(err.Error())
		return false, err
	}

	return isLinked, nil
}

func (r *AnimalPostgres) IsExistVisitedLocationByID(visitedLocationID int64) (bool, error) {
	isExist, err := r.db.IsExistVisitedLocationByID(context.Background(), visitedLocationID)
	if err != nil {
		logrus.Error(err.Error())
		return false, err
	}

	return isExist, nil
}

func (r *AnimalPostgres) IsExistByID(animalID int64) (bool, error) {
	isExist, err := r.db.IsExistAnimalByID(context.Background(), animalID)
	if err != nil {
		logrus.Error(err.Error())
		return false, err
	}

	return isExist, nil
}

func (r *AnimalPostgres) GetVisitedLocations(animalID int64) ([]queries.AnimalVisitedLocation, error) {
	points, err := r.db.GetVisitedLocations(context.Background(), animalID)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return points, nil
}

func (r *AnimalPostgres) GetVisitedLocation(visitedPointID int64) (*queries.AnimalVisitedLocation, error) {
	point, err := r.db.GetVisitedLocation(context.Background(), visitedPointID)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &point, nil
}

func (r *AnimalPostgres) CreateVisitedLocation(animalID int64, pointID int64) (*queries.AnimalVisitedLocation, error) {
	point, err := r.db.CreateVisitedLocation(context.Background(), queries.CreateVisitedLocationParams{
		Location: pointID,
		Animal:   animalID,
	})
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &point, nil
}

func (r *AnimalPostgres) GetCurrentLocation(animalID int64) (*queries.AnimalVisitedLocation, error) {
	point, err := r.db.GetCurrentLocation(context.Background(), animalID)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &point, nil
}

func (r *AnimalPostgres) GetChippingLocation(animalID int64) (*queries.LocationPoint, error) {
	point, err := r.db.GetChippingLocation(context.Background(), animalID)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &point, nil
}

func (r *AnimalPostgres) Get(id int64) (*queries.Animal, error) {
	animal, err := r.db.GetAnimal(context.Background(), id)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &animal, nil
}

func NewAnimalPostgres(db *queries.Queries) *AnimalPostgres {
	return &AnimalPostgres{db: db}
}
