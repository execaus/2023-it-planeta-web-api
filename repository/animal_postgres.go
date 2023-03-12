package repository

import (
	"2023-it-planeta-web-api/queries"
	"context"
	"github.com/sirupsen/logrus"
)

type AnimalPostgres struct {
	db *queries.Queries
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
