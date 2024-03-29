package service

import (
	"2023-it-planeta-web-api/constants"
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/models_output"
	"2023-it-planeta-web-api/queries"
	"2023-it-planeta-web-api/repository"
	"2023-it-planeta-web-api/utils"
	"database/sql"
	"github.com/execaus/exloggo"

	"time"
)

type AnimalService struct {
	repo repository.Animal
}

func (s *AnimalService) FillAnimalOutput(output *models_output.OutputAnimal, services *Service, animalID int64) error {
	animal, err := services.Animal.Get(animalID)
	if err != nil {
		return err
	}

	animalTypes, err := services.AnimalType.GetByAnimalID(animalID)
	if err != nil {
		return err
	}

	animalTypesID := make([]int64, len(animalTypes))
	for i, animalType := range animalTypes {
		animalTypesID[i] = animalType.AnimalType
	}

	visitedLocations, err := services.Animal.GetVisitedLocations(animalID)
	if err != nil {
		return err
	}

	visitedLocationsID := make([]int64, len(visitedLocations))
	for i, location := range visitedLocations {
		visitedLocationsID[i] = location.ID
	}

	output.ID = animal.ID
	output.AnimalTypes = animalTypesID
	output.Weight = animal.Weight
	output.Length = animal.Length
	output.Height = animal.Height
	output.Gender = animal.Gender
	output.LifeStatus = animal.LifeStatus
	output.ChippingDateTime = utils.ConvertDateToISO8601(animal.ChippingDate)
	output.ChipperID = animal.Chipper
	output.ChippingLocationID = animal.ChippingLocation
	output.VisitedLocations = visitedLocationsID
	output.DeathDateTime = utils.ConvertNullDateToISO8601(animal.DeathDate)

	return nil
}

func (s *AnimalService) RemoveAnimalType(animalID int64, typeID int64) error {
	params := queries.RemoveAnimalTypeToAnimalParams{
		Animal:     animalID,
		AnimalType: typeID,
	}
	return s.repo.RemoveAnimalType(&params)
}

func (s *AnimalService) UpdateAnimalTypeToAnimal(animalID int64, input *models.UpdateAnimalTypeToAnimalInput) error {
	params := queries.UpdateAnimalTypeToAnimalParams{
		AnimalType:   input.NewTypeId,
		Animal:       animalID,
		AnimalType_2: input.OldTypeId,
	}

	return s.repo.UpdateAnimalTypeToAnimal(&params)
}

func (s *AnimalService) LinkAnimalType(animalID int64, typeID int64) error {
	return s.repo.LinkAnimalType(animalID, typeID)
}

func (s *AnimalService) IsLinkedAnimalType(animalID int64, typeID int64) (bool, error) {
	return s.repo.IsLinkedAnimalType(animalID, typeID)
}

func (s *AnimalService) Remove(animalID int64) error {
	return s.repo.Remove(animalID)
}

func (s *AnimalService) Update(animalID int64, input *models.UpdateAnimalInput) (*queries.Animal, error) {
	deathDate := sql.NullTime{
		Time:  time.Time{},
		Valid: false,
	}

	if input.LifeStatus == constants.AnimalLifeStatusDead {
		deathDate.Valid = true
		deathDate.Time = time.Now()
	}

	params := queries.UpdateAnimalParams{
		Weight:           input.Weight,
		Length:           input.Length,
		Height:           input.Height,
		Gender:           input.Gender,
		LifeStatus:       input.LifeStatus,
		Chipper:          input.ChipperID,
		ChippingLocation: input.ChippingLocationID,
		DeathDate:        deathDate,
		ID:               animalID,
	}

	return s.repo.Update(&params)
}

func (s *AnimalService) Create(input *models.CreateAnimalInput) (*queries.Animal, error) {
	account, err := s.repo.Create(&queries.CreateAnimalParams{
		ChippingLocation: input.ChippingLocationID,
		Weight:           input.Weight,
		Length:           input.Length,
		Height:           input.Height,
		Gender:           input.Gender,
		LifeStatus:       constants.AnimalLifeStatusAlive,
		ChippingDate:     time.Now(),
		Chipper:          input.ChipperID,
		DeathDate: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
		Deleted: false,
	})
	if err != nil {
		return nil, err
	}

	for _, animalType := range input.AnimalTypes {
		_, err = s.repo.BindAnimalType(account.ID, *animalType)
		if err != nil {
			return nil, err
		}
	}

	return account, err
}

func (s *AnimalService) GetList(input *models.GetAnimalsInput) ([]queries.Animal, error) {
	var limit int32
	var offset int32
	var endDateTime time.Time
	var startDateTime time.Time

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

	if input.StartDateTime != nil {
		date, err := time.Parse(time.RFC3339, *input.StartDateTime)
		if err != nil {
			exloggo.Error(err.Error())
			return nil, err
		}
		startDateTime = date
	}

	if input.EndDateTime != nil {
		date, err := time.Parse(time.RFC3339, *input.EndDateTime)
		if err != nil {
			exloggo.Error(err.Error())
			return nil, err
		}
		endDateTime = date
	}

	params := queries.GetAnimalsParams{
		ChippingDate:   startDateTime,
		ChippingDate_2: endDateTime,
		Column3:        input.ChipperID,
		Column4:        input.ChippingLocationID,
		Column5:        input.LifeStatus,
		Column6:        input.Gender,
		Limit:          limit,
		Offset:         offset,
	}

	return s.repo.GetList(&params)
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
		t := time.Unix(0, 0)
		startDateTime = &t
	} else {
		parseDate, err := time.Parse(time.RFC3339, *input.StartDateTime)
		if err != nil {
			return nil, err
		}
		startDateTime = &parseDate
	}

	if input.EndDateTime == nil {
		t := time.Now()
		endDateTime = &t
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
		exloggo.Error(err.Error())
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
