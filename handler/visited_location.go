package handler

import (
	"2023-it-planeta-web-api/models"
	"github.com/gin-gonic/gin"
	"time"
)

func (h *Handler) createVisitedLocation(c *gin.Context) {
	animalID, err := getNumberParam(c, "animalId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	pointID, err := getNumberParam(c, "pointId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	// У животного lifeStatus = "DEAD"
	isAnimalDead, err := h.services.Animal.IsDead(animalID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if isAnimalDead {
		h.sendBadRequest(c, "animal is dead")
		return
	}

	// Животное находится в точке чипирования и никуда не перемещалось, попытка добавить точку локации,
	// равную точке чипирования.
	visitedLocations, err := h.services.Location.GetVisitedAnimal(animalID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	chippingLocation, err := h.services.Animal.GetChippingLocation(animalID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	isChippingPoint := chippingLocation.ID == pointID
	isAnimalDontMove := len(visitedLocations) == 0

	if isAnimalDontMove && isChippingPoint {
		h.sendBadRequest(c, "point is chipping location or animal don't move")
		return
	}

	// Попытка добавить точку локации, в которой уже находится животное
	currentLocation, err := h.services.Animal.GetCurrentLocation(animalID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if currentLocation.Location == pointID {
		h.sendBadRequest(c, "there is already an animal at the location point")
		return
	}

	visitedLocation, err := h.services.Animal.CreateVisitedLocation(animalID, pointID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	output := models.CreateVisitedLocationOutput{
		ID:                           visitedLocation.ID,
		DateTimeOfVisitLocationPoint: visitedLocation.Date.Format(time.RFC3339),
		LocationPointID:              visitedLocation.Location,
	}

	h.sendOKWithBody(c, &output)
}
