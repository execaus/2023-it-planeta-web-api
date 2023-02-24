package service

import (
	"2023-it-planeta-web-api/queries"
	"2023-it-planeta-web-api/repository"
)

type AnimalTypeService struct {
	repo repository.AnimalType
}

func (s *AnimalTypeService) GetFromAnimal(id int64) ([]queries.AnimalToType, error) {
	return s.repo.GetFromAnimal(id)
}

func NewAnimalTypeService(repo repository.AnimalType) *AnimalTypeService {
	return &AnimalTypeService{repo: repo}
}
