package repository

import (
	"2023-it-planeta-web-api/queries"
	"context"
	"database/sql"
	"github.com/execaus/exloggo"
)

type LocationPostgres struct {
	db *queries.Queries
}

func (r *LocationPostgres) Remove(id int64) error {
	if _, err := r.db.RemoveLocation(context.Background(), id); err != nil {
		exloggo.Error(err.Error())
		return err
	}

	return nil
}

func (r *LocationPostgres) IsVisitedAnimal(id int64) (bool, error) {
	isExist, err := r.db.IsLocationVisitedAnimal(context.Background(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		exloggo.Error(err.Error())
		return false, err
	}

	return isExist, nil
}

func (r *LocationPostgres) IsAnimalChipping(id int64) (bool, error) {
	isExist, err := r.db.IsLocationChippingAnimal(context.Background(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		exloggo.Error(err.Error())
		return false, err
	}

	return isExist, nil
}

func (r *LocationPostgres) Update(params *queries.UpdateLocationParams) (*queries.LocationPoint, error) {
	location, err := r.db.UpdateLocation(context.Background(), *params)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &location, nil
}

func (r *LocationPostgres) Create(params *queries.CreateLocationParams) (*queries.LocationPoint, error) {
	location, err := r.db.CreateLocation(context.Background(), *params)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &location, nil
}

func (r *LocationPostgres) IsExistByCoordinates(params *queries.IsExistLocationByCoordinatesParams) (bool, error) {
	isExist, err := r.db.IsExistLocationByCoordinates(context.Background(), *params)
	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}

	return isExist, nil
}

func (r *LocationPostgres) IsExistByID(id int64) (bool, error) {
	isExist, err := r.db.IsExistLocationByID(context.Background(), id)
	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}

	return isExist, nil
}

func (r *LocationPostgres) Get(id int64) (*queries.LocationPoint, error) {
	location, err := r.db.GetLocation(context.Background(), id)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &location, err
}

func (r *LocationPostgres) GetVisitedAnimal(id int64) ([]queries.AnimalVisitedLocation, error) {
	points, err := r.db.GetVisitedLocationByAnimalID(context.Background(), id)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}
	return points, nil
}

func NewLocationPostgres(db *queries.Queries) *LocationPostgres {
	return &LocationPostgres{db: db}
}
