package repository

import (
	"2023-it-planeta-web-api/queries"
	"context"
	"github.com/sirupsen/logrus"
)

type LocationPostgres struct {
	db *queries.Queries
}

func (r *LocationPostgres) GetVisitedAnimal(id int64) ([]queries.AnimalVisitedLocation, error) {
	points, err := r.db.GetVisitedLocationFromAnimal(context.Background(), id)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return points, nil
}

func NewLocationPostgres(db *queries.Queries) *LocationPostgres {
	return &LocationPostgres{db: db}
}