package handler

import (
	"2023-it-planeta-web-api/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAnimalType(c *gin.Context) {
	id, err := getParamID(c, "typeId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	isExist, err := h.services.AnimalType.IsExist(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExist {
		h.sendNotFound(c)
		return
	}

	animalType, err := h.services.AnimalType.GetById(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	output := &models.GetAnimalTypeOutput{
		ID:   animalType.ID,
		Type: animalType.Value,
	}

	h.sendOKWithBody(c, output)
}
