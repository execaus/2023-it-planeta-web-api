package service

import (
	"2023-it-planeta-web-api/queries"
	"2023-it-planeta-web-api/repository"
)

type LocationService struct {
	repo repository.Location
}

func (r *LocationService) GetVisitedAnimal(id int64) ([]queries.AnimalVisitedLocation, error) {
	return r.repo.GetVisitedAnimal(id)
}

func NewLocationService(repo repository.Location) *LocationService {
	return &LocationService{repo: repo}
}
