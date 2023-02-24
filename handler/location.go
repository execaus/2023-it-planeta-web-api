package handler

import (
	"2023-it-planeta-web-api/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getLocation(c *gin.Context) {
	id, err := getParamID(c, "pointId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	isExist, err := h.services.Location.IsExistByID(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExist {
		h.sendNotFound(c)
		return
	}

	location, err := h.services.Location.Get(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	output := &models.GetLocationOutput{
		ID:        location.ID,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
	}

	h.sendOKWithBody(c, &output)
}

func (h *Handler) createLocation(c *gin.Context) {
	var input models.CreateLocationInput
	if err := c.BindJSON(&input); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	isExist, err := h.services.Location.IsExistByCoordinates(input.Latitude, input.Longitude)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if isExist {
		h.sendConflict(c)
		return
	}

	location, err := h.services.Location.Create(input.Latitude, input.Longitude)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	output := &models.CreateLocationOutput{
		ID:        location.ID,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
	}

	h.sendCreatedWithBody(c, &output)
}

func (h *Handler) updateLocation(c *gin.Context) {
	id, err := getParamID(c, "pointId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	isPointExist, err := h.services.Location.IsExistByID(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if isPointExist {
		h.sendNotFound(c)
		return
	}

	var input models.UpdateLocationInput
	if err = c.BindJSON(&input); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	isCoordinatesExist, err := h.services.Location.IsExistByCoordinates(input.Latitude, input.Longitude)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if isCoordinatesExist {
		h.sendConflict(c)
		return
	}

	location, err := h.services.Location.Update(id, input.Latitude, input.Longitude)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	output := &models.UpdateLocationOutput{
		ID:        location.ID,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
	}

	h.sendCreatedWithBody(c, &output)
}

func (h *Handler) removeLocation(c *gin.Context) {
	id, err := getParamID(c, "pointId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	isPointExist, err := h.services.Location.IsExistByID(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if isPointExist {
		h.sendNotFound(c)
		return
	}

	isLinkedAnimal, err := h.services.Location.IsLinkedAnimal(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if isLinkedAnimal {
		h.sendBadRequest(c, "the location point is linked with the animal")
		return
	}

	if err = h.services.Location.Remove(id); err != nil {
		h.sendInternalServerError(c)
		return
	}

	h.sendOk(c)
}
