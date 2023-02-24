package repository

import (
	"2023-it-planeta-web-api/queries"
	"context"
	"github.com/sirupsen/logrus"
)

type AnimalTypePostgres struct {
	db *queries.Queries
}

func (r *AnimalTypePostgres) GetFromAnimal(id int64) ([]queries.AnimalToType, error) {
	types, err := r.db.GetAnimalTypesFromAnimal(context.Background(), id)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return types, nil
}

func NewAnimalTypePostgres(db *queries.Queries) *AnimalTypePostgres {
	return &AnimalTypePostgres{db: db}
}
