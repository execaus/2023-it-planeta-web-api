package service

import (
	"2023-it-planeta-web-api/queries"
	"2023-it-planeta-web-api/repository"
)

type AnimalTypeService struct {
	repo repository.AnimalType
}

func (s *AnimalTypeService) Update(id int64, animalType string) (*queries.AnimalType, error) {
	params := queries.UpdateAnimalTypeParams{
		Value: animalType,
		ID:    id,
	}
	return s.repo.Update(&params)
}

func (s *AnimalTypeService) IsExistByType(animalType string) (bool, error) {
	return s.repo.IsExistByType(animalType)
}

func (s *AnimalTypeService) Create(animalType string) (*queries.AnimalType, error) {
	return s.repo.Create(animalType)
}

func (s *AnimalTypeService) IsExistByID(id int64) (bool, error) {
	return s.repo.IsExistByID(id)
}

func (s *AnimalTypeService) GetByID(id int64) (*queries.AnimalType, error) {
	return s.repo.GetByID(id)
}

func (s *AnimalTypeService) GetByAnimalID(id int64) ([]queries.AnimalToType, error) {
	return s.repo.GetByAnimalID(id)
}

func NewAnimalTypeService(repo repository.AnimalType) *AnimalTypeService {
	return &AnimalTypeService{repo: repo}
}
