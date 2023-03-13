package handler

import (
	"2023-it-planeta-web-api/constants"
	"2023-it-planeta-web-api/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAnimal(c *gin.Context) {
	id, err := getNumberParam(c, "animalId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	animal, err := h.services.Animal.Get(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	animalTypes, err := h.services.AnimalType.GetByAnimalID(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	visitedLocations, err := h.services.Location.GetVisitedAnimal(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	animalTypesID := make([]int64, len(animalTypes))
	for i, animalType := range animalTypes {
		animalTypesID[i] = animalType.AnimalType
	}

	visitedLocationsID := make([]int64, len(visitedLocations))
	for i, location := range visitedLocations {
		visitedLocationsID[i] = location.ID
	}

	output := models.GetAnimalOutput{
		ID:                 animal.ID,
		AnimalTypes:        animalTypesID,
		Weight:             animal.Weight,
		Length:             animal.Length,
		Height:             animal.Height,
		Gender:             animal.Gender,
		LifeStatus:         animal.LifeStatus,
		ChippingDateTime:   convertDateToISO8601(animal.ChippingDate),
		ChipperID:          animal.Chipper,
		ChippingLocationID: animal.ChippingLocation,
		VisitedLocations:   visitedLocationsID,
		DeathDateTime:      convertNullDateToISO8601(animal.DeathDate),
	}

	h.sendOKWithBody(c, &output)
}

func (h *Handler) getAnimals(c *gin.Context) {
	var input models.GetAnimalsInput

	if err := c.ShouldBindQuery(&input); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	if input.From != nil && *input.From < 0 {
		h.sendBadRequest(c, "invalid parameter from")
		return
	}

	if input.Size != nil && *input.Size <= 0 {
		h.sendBadRequest(c, "invalid parameter size")
		return
	}

	if input.StartDateTime != nil && !IsISO8601Date(*input.StartDateTime) {
		h.sendBadRequest(c, "invalid parameter start date time")
		return
	}

	if input.EndDateTime != nil && !IsISO8601Date(*input.EndDateTime) {
		h.sendBadRequest(c, "invalid parameter end date time")
		return
	}

	if input.ChipperID != nil && *input.ChipperID <= 0 {
		h.sendBadRequest(c, "invalid parameter chipper id")
		return
	}

	if input.ChipperID != nil && *input.ChipperID <= 0 {
		h.sendBadRequest(c, "invalid parameter chipping location")
		return
	}

	if input.LifeStatus != nil {
		if !constants.IsAnimalLifeStatus(*input.LifeStatus) {
			h.sendBadRequest(c, "invalid parameter life status")
			return
		}
	}

	if input.Gender != nil {
		if !constants.IsAnimalGender(*input.Gender) {
			h.sendBadRequest(c, "invalid parameter gender")
			return
		}
	}

	animals, err := h.services.Animal.GetList(&input)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	output := make([]*models.GetAnimalsOutput, len(animals))
	for i, animal := range animals {
		animalTypes, err := h.services.AnimalType.GetByAnimalID(animal.ID)
		if err != nil {
			h.sendInternalServerError(c)
			return
		}

		visitedLocations, err := h.services.Location.GetVisitedAnimal(animal.ID)
		if err != nil {
			h.sendInternalServerError(c)
			return
		}

		animalTypesID := make([]int64, len(animalTypes))
		for i, animalType := range animalTypes {
			animalTypesID[i] = animalType.AnimalType
		}

		visitedLocationsID := make([]int64, len(visitedLocations))
		for i, location := range visitedLocations {
			visitedLocationsID[i] = location.ID
		}

		output[i] = &models.GetAnimalsOutput{
			ID:                 animal.ID,
			AnimalTypes:        animalTypesID,
			Weight:             animal.Weight,
			Length:             animal.Length,
			Height:             animal.Height,
			Gender:             animal.Gender,
			LifeStatus:         animal.LifeStatus,
			ChippingDateTime:   convertDateToISO8601(animal.ChippingDate),
			ChipperID:          animal.Chipper,
			ChippingLocationID: animal.ChippingLocation,
			VisitedLocations:   visitedLocationsID,
			DeathDateTime:      convertNullDateToISO8601(animal.DeathDate),
		}
	}

	h.sendOKWithBody(c, output)
}
