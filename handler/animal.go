package handler

import (
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

	output := &models.GetAnimalOutput{
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

	h.sendOKWithBody(c, output)
}
