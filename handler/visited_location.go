package handler

import (
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func (h *Handler) getVisitedLocation(c *gin.Context) {
	var input models.GetVisitedLocationQueryParams

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

	if input.StartDateTime != nil && !utils.IsISO8601Date(*input.StartDateTime) {
		h.sendBadRequest(c, "invalid parameter start date time")
		return
	}

	if input.EndDateTime != nil && !utils.IsISO8601Date(*input.EndDateTime) {
		h.sendBadRequest(c, "invalid parameter end date time")
		return
	}

	animalID, err := utils.GetNumberParam(c, "animalId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	// 404 Животное с animalId не найдено
	isExistAnimal, err := h.services.Animal.IsExistByID(animalID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExistAnimal {
		h.sendNotFound(c)
		return
	}

	visitedLocations, err := h.services.Animal.GetVisitedLocationList(animalID, &input)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	output := make([]*models.GetVisitedLocationOutput, len(visitedLocations))

	for i, visitedLocation := range visitedLocations {
		output[i] = &models.GetVisitedLocationOutput{
			ID:                           visitedLocation.ID,
			DateTimeOfVisitLocationPoint: utils.ConvertDateToISO8601(visitedLocation.Date),
			LocationPointID:              visitedLocation.Location,
		}
	}

	h.sendOKWithBody(c, output)
}

func (h *Handler) createVisitedLocation(c *gin.Context) {
	animalID, err := utils.GetNumberParam(c, "animalId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	pointID, err := utils.GetNumberParam(c, "pointId")
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

	// Животное находится в точке чипирования и никуда не перемещалось
	visitedLocations, err := h.services.Location.GetVisitedAnimal(animalID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	// Попытка добавить точку локации, равную точке чипирования.
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

func (h *Handler) updateVisitedLocation(c *gin.Context) {
	var input models.UpdateVisitedLocationInput

	if err := c.BindJSON(&input); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	animalID, err := utils.GetNumberParam(c, "animalId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	if input.LocationPointID <= 0 || input.VisitedLocationPointID <= 0 {
		h.sendBadRequest(c, "invalid location point id or visited location point id")
		return
	}

	// 400 Обновление первой посещенной точки на точку чипирования
	chippingLocation, err := h.services.Animal.GetChippingLocation(animalID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if chippingLocation.ID == input.LocationPointID {
		h.sendBadRequest(c, "updating the first visited point to a chipping point")
		return
	}

	// 400 Обновление точки на такую же точку
	currentVisitedPoint, err := h.services.Animal.GetVisitedLocation(input.VisitedLocationPointID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if currentVisitedPoint.Location == input.LocationPointID {
		h.sendBadRequest(c, "updating a point to the same point")
		return
	}

	// 400 Обновление точки локации на точку, совпадающую со следующей и/или с предыдущей точками
	visitedLocations, err := h.services.Animal.GetVisitedLocations(animalID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	isSurroundedDuplicatesPoints, err := h.services.Location.IsSurroundedDuplicatesPoints(
		visitedLocations,
		input.VisitedLocationPointID,
	)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if isSurroundedDuplicatesPoints {
		h.sendBadRequest(c, "update the location point to the point "+
			"that matches the next and/or previous point")
		return
	}

	// 404 Животное с animalId не найдено
	isAnimalExist, err := h.services.Animal.IsExistByID(animalID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isAnimalExist {
		h.sendNotFound(c)
		return
	}

	// 404 Объект с информацией о посещенной точке локации с visitedLocationPointId не найден.
	isExistVisitedLocation, err := h.services.Animal.IsExistVisitedLocationByID(input.VisitedLocationPointID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExistVisitedLocation {
		h.sendNotFound(c)
		return
	}

	// 404 У животного нет объекта с информацией о посещенной точке локации с visitedLocationPointId.
	isLinkedVisitedLocationPoint, err := h.services.Animal.IsLinkedVisitedLocation(
		animalID,
		input.VisitedLocationPointID,
	)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isLinkedVisitedLocationPoint {
		h.sendNotFound(c)
		return
	}

	// 404 Точка локации с locationPointId не найден
	isExistLocationPoint, err := h.services.Location.IsExistByID(input.LocationPointID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExistLocationPoint {
		h.sendNotFound(c)
		return
	}

	visitedLocation, err := h.services.Animal.UpdateVisitedLocation(input.VisitedLocationPointID, input.LocationPointID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	output := models.UpdateVisitedLocationOutput{
		ID:                           visitedLocation.ID,
		DateTimeOfVisitLocationPoint: visitedLocation.Date.Format(time.RFC3339),
		LocationPointID:              visitedLocation.Location,
	}

	h.sendOKWithBody(c, &output)
}

func (h *Handler) removeVisitedLocation(c *gin.Context) {
	animalID, err := utils.GetNumberParam(c, "animalId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	visitedLocationID, err := utils.GetNumberParam(c, "visitedPointId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	// 404 Животное с animalId не найдено
	isExistAnimal, err := h.services.Animal.IsExistByID(animalID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExistAnimal {
		h.sendNotFound(c)
		return
	}

	// 404 Объект с информацией о посещенной точке локации с visitedPointId не найден.
	isExistLocation, err := h.services.Animal.IsExistVisitedLocationByID(visitedLocationID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExistLocation {
		h.sendNotFound(c)
		return
	}

	// 404 У животного нет объекта с информацией о посещенной точке локации с visitedPointId
	isLinkedVisitedLocationPoint, err := h.services.Animal.IsLinkedVisitedLocation(
		animalID,
		visitedLocationID,
	)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isLinkedVisitedLocationPoint {
		h.sendNotFound(c)
		return
	}

	if err = h.services.Animal.RemoveVisitedLocationID(animalID, visitedLocationID); err != nil {
		h.sendInternalServerError(c)
		return
	}

	h.sendOk(c)
}
