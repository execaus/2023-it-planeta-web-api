package service

import (
	"2023-it-planeta-web-api/queries"
	"2023-it-planeta-web-api/repository"
)

type LocationService struct {
	repo repository.Location
}

func (s *LocationService) IsExist(id int64) (bool, error) {
	return s.repo.IsExist(id)
}

func (s *LocationService) Get(id int64) (*queries.LocationPoint, error) {
	return s.repo.Get(id)
}

func (s *LocationService) GetVisitedAnimal(id int64) ([]queries.AnimalVisitedLocation, error) {
	return s.repo.GetVisitedAnimal(id)
}

func NewLocationService(repo repository.Location) *LocationService {
	return &LocationService{repo: repo}
}
