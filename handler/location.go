package handler

import (
	"2023-it-planeta-web-api/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) getLocation(c *gin.Context) {
	stringID := c.Param("pointId")
	if stringID == stringEmpty || stringID == stringNull {
		h.sendBadRequest(c, "id is not valid")
		return
	}

	id, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		h.sendBadRequest(c, "id is not valid")
		return
	}

	if id <= 0 {
		h.sendBadRequest(c, "id is not valid")
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
