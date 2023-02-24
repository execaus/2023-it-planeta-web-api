package service

import (
	"2023-it-planeta-web-api/queries"
	"2023-it-planeta-web-api/repository"
)

type LocationService struct {
	repo repository.Location
}

func (s *LocationService) Update(id int64, latitude float64, longitude float64) (*queries.LocationPoint, error) {
	params := &queries.UpdateLocationParams{
		ID:        id,
		Latitude:  latitude,
		Longitude: longitude,
	}
	return s.repo.Update(params)
}

func (s *LocationService) Create(latitude float64, longitude float64) (*queries.LocationPoint, error) {
	params := &queries.CreateLocationParams{
		Latitude:  latitude,
		Longitude: longitude,
	}
	return s.repo.Create(params)
}

func (s *LocationService) IsExistByCoordinates(latitude float64, longitude float64) (bool, error) {
	params := &queries.IsExistLocationByCoordinatesParams{
		Latitude:  latitude,
		Longitude: longitude,
	}
	return s.repo.IsExistByCoordinates(params)
}

func (s *LocationService) IsExistByID(id int64) (bool, error) {
	return s.repo.IsExistByID(id)
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
