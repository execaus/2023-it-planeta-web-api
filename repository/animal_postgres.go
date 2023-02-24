package repository

import (
	"2023-it-planeta-web-api/queries"
	"context"
	"github.com/sirupsen/logrus"
)

type AnimalPostgres struct {
	db *queries.Queries
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
