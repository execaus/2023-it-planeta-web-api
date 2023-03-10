package service

import (
	"2023-it-planeta-web-api/queries"
	"2023-it-planeta-web-api/repository"
)

type LocationService struct {
	repo repository.Location
}

func (s *LocationService) IsSurroundedDuplicatesPoints(
	visitedLocations []queries.AnimalVisitedLocation,
	targetLocationID int64,
) (bool, error) {
	for i, visitedLocation := range visitedLocations {
		if visitedLocation.Location != targetLocationID {
			continue
		}

		targetLocationPoint, err := s.repo.Get(visitedLocation.Location)
		if err != nil {
			return false, err
		}

		if i-1 < len(visitedLocations) && i-1 != -1 {
			prevVisitedLocation := visitedLocations[i-1]
			locationPoint, err := s.repo.Get(prevVisitedLocation.Location)
			if err != nil {
				return false, err
			}

			if locationPoint.Longitude == targetLocationPoint.Longitude &&
				locationPoint.Latitude == targetLocationPoint.Latitude {
				return true, nil
			}
		}

		if i+1 < len(visitedLocations) {
			prevVisitedLocation := visitedLocations[i+1]
			locationPoint, err := s.repo.Get(prevVisitedLocation.Location)
			if err != nil {
				return false, err
			}

			if locationPoint.Longitude == targetLocationPoint.Longitude &&
				locationPoint.Latitude == targetLocationPoint.Latitude {
				return true, nil
			}
		}

		break
	}

	return false, nil
}

func (s *LocationService) Remove(id int64) error {
	return s.repo.Remove(id)
}

func (s *LocationService) IsLinkedAnimal(id int64) (bool, error) {
	isLinked, err := s.repo.IsVisitedAnimal(id)
	if err != nil {
		return false, err
	}
	if isLinked {
		return true, nil
	}
	isLinked, err = s.repo.IsAnimalChipping(id)
	if err != nil {
		return false, err
	}
	return isLinked, nil
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
