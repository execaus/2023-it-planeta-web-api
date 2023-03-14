package models

import (
	"2023-it-planeta-web-api/ctypes"
	"2023-it-planeta-web-api/service"
	"2023-it-planeta-web-api/utils"
)

type outputAnimal struct {
	ID                 int64            `json:"id"`
	AnimalTypes        []int64          `json:"animalTypes"`
	Weight             float64          `json:"weight"`
	Length             float64          `json:"length"`
	Height             float64          `json:"height"`
	Gender             string           `json:"gender"`
	LifeStatus         string           `json:"lifeStatus"`
	ChippingDateTime   string           `json:"chippingDateTime"`
	ChipperID          int64            `json:"chipperId"`
	ChippingLocationID int64            `json:"chippingLocationId"`
	VisitedLocations   []int64          `json:"visitedLocations"`
	DeathDateTime      ctypes.TimeOrNil `json:"deathDateTime"`
}

func (a *outputAnimal) Load(services *service.Service, animalID int64) error {
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

	a.ID = animal.ID
	a.AnimalTypes = animalTypesID
	a.Weight = animal.Weight
	a.Length = animal.Length
	a.Height = animal.Height
	a.Gender = animal.Gender
	a.LifeStatus = animal.LifeStatus
	a.ChippingDateTime = utils.ConvertDateToISO8601(animal.ChippingDate)
	a.ChipperID = animal.Chipper
	a.ChippingLocationID = animal.ChippingLocation
	a.VisitedLocations = visitedLocationsID
	a.DeathDateTime = utils.ConvertNullDateToISO8601(animal.DeathDate)

	return nil
}
