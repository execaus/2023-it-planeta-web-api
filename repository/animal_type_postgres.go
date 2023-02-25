package repository

import (
	"2023-it-planeta-web-api/queries"
	"context"
	"github.com/sirupsen/logrus"
)

type AnimalTypePostgres struct {
	db *queries.Queries
}

func (r *AnimalTypePostgres) IsExistByType(animalType string) (bool, error) {
	isExist, err := r.db.IsExistAnimalTypeByType(context.Background(), animalType)
	if err != nil {
		logrus.Error(err.Error())
		return false, err
	}

	return isExist, nil
}

func (r *AnimalTypePostgres) Create(animalType string) (*queries.AnimalType, error) {
	animalTypeRow, err := r.db.CreateAnimalType(context.Background(), animalType)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return &animalTypeRow, nil
}

func (r *AnimalTypePostgres) IsExistByID(id int64) (bool, error) {
	isExist, err := r.db.IsExistAnimalTypeByID(context.Background(), id)
	if err != nil {
		logrus.Error(err.Error())
		return false, err
	}

	return isExist, nil
}

func (r *AnimalTypePostgres) GetByID(id int64) (*queries.AnimalType, error) {
	animalType, err := r.db.GetAnimalTypeByID(context.Background(), id)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return &animalType, nil
}

func (r *AnimalTypePostgres) GetByAnimalID(id int64) ([]queries.AnimalToType, error) {
	types, err := r.db.GetAnimalTypesByAnimalID(context.Background(), id)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return types, nil
}

func NewAnimalTypePostgres(db *queries.Queries) *AnimalTypePostgres {
	return &AnimalTypePostgres{db: db}
}
