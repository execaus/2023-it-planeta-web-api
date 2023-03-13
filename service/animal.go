package service

import (
	"2023-it-planeta-web-api/constants"
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/queries"
	"2023-it-planeta-web-api/repository"
	"github.com/sirupsen/logrus"
	"time"
)

type AnimalService struct {
	repo repository.Animal
}

func (s *AnimalService) GetVisitedLocationList(
	animalID int64,
	input *models.GetVisitedLocationQueryParams) ([]queries.AnimalVisitedLocation, error) {
	var limit int32
	var offset int32
	var startDateTime *time.Time
	var endDateTime *time.Time

	if input.Size == nil {
		limit = constants.VisitedLocationGetListDefaultLimit
	} else {
		limit = *input.Size
	}

	if input.From == nil {
		offset = constants.VisitedLocationGetListDefaultOffset
	} else {
		offset = *input.From
	}

	if input.StartDateTime == nil {
		startDateTime = nil
	} else {
		parseDate, err := time.Parse(time.RFC3339, *input.StartDateTime)
		if err != nil {
			return nil, err
		}
		startDateTime = &parseDate
	}

	if input.EndDateTime == nil {
		endDateTime = nil
	} else {
		parseDate, err := time.Parse(time.RFC3339, *input.EndDateTime)
		if err != nil {
			return nil, err
		}
		endDateTime = &parseDate
	}

	params := queries.GetVisitedLocationListParams{
		Animal:  animalID,
		Column2: *startDateTime,
		Column3: *endDateTime,
		Offset:  offset,
		Limit:   limit,
	}

	return s.repo.GetVisitedLocationList(&params)
}

func (s *AnimalService) RemoveVisitedLocationID(animalID int64, visitedLocationID int64) error {
	// Если удаляется первая посещенная точка локации,
	// а вторая точка совпадает с точкой чипирования, то она удаляется автоматически
	visitedLocations, err := s.repo.GetVisitedLocations(animalID)
	if err != nil {
		return err
	}

	if visitedLocations[0].ID != visitedLocationID {
		chippingPoint, err := s.repo.GetChippingLocation(animalID)
		if err != nil {
			return err
		}

		secondLocation := visitedLocations[1].Location
		if chippingPoint.ID == secondLocation {
			if err = s.repo.RemoveVisitedLocation(secondLocation); err != nil {
				return err
			}
		}
	}

	if err = s.repo.RemoveVisitedLocation(visitedLocationID); err != nil {
		return err
	}
	return nil
}

func (s *AnimalService) UpdateVisitedLocation(
	visitedLocationPointID int64,
	locationPointID int64) (*queries.AnimalVisitedLocation, error) {
	return s.repo.UpdateVisitedLocation(visitedLocationPointID, locationPointID)
}

func (s *AnimalService) IsLinkedVisitedLocation(animalID int64, visitedLocationPointID int64) (bool, error) {
	return s.repo.IsLinkedVisitedLocation(animalID, visitedLocationPointID)
}

func (s *AnimalService) IsExistVisitedLocationByID(visitedLocationID int64) (bool, error) {
	return s.repo.IsExistVisitedLocationByID(visitedLocationID)
}

func (s *AnimalService) IsExistByID(animalID int64) (bool, error) {
	return s.repo.IsExistByID(animalID)
}

func (s *AnimalService) GetVisitedLocations(animalID int64) ([]queries.AnimalVisitedLocation, error) {
	return s.repo.GetVisitedLocations(animalID)
}

func (s *AnimalService) GetVisitedLocation(visitedPointID int64) (*queries.AnimalVisitedLocation, error) {
	return s.repo.GetVisitedLocation(visitedPointID)
}

func (s *AnimalService) CreateVisitedLocation(animalID int64, pointID int64) (*queries.AnimalVisitedLocation, error) {
	return s.repo.CreateVisitedLocation(animalID, pointID)
}

func (s *AnimalService) GetCurrentLocation(animalID int64) (*queries.AnimalVisitedLocation, error) {
	return s.repo.GetCurrentLocation(animalID)
}

func (s *AnimalService) GetChippingLocation(animalID int64) (*queries.LocationPoint, error) {
	return s.repo.GetChippingLocation(animalID)
}

func (s *AnimalService) IsDead(id int64) (bool, error) {
	animal, err := s.repo.Get(id)
	if err != nil {
		logrus.Error(err.Error())
		return false, err
	}

	isDead := animal.LifeStatus == constants.AnimalLifeStatusDead
	return isDead, nil
}

func (s *AnimalService) Get(id int64) (*queries.Animal, error) {
	return s.repo.Get(id)
}

func NewAnimalService(repo repository.Animal) *AnimalService {
	return &AnimalService{repo: repo}
}
