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

	isExist, err := h.services.AnimalType.IsExistByID(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExist {
		h.sendNotFound(c)
		return
	}

	animalType, err := h.services.AnimalType.GetByID(id)
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

func (h *Handler) createAnimalType(c *gin.Context) {
	var input models.CreateAnimalTypeInput

	if err := c.BindJSON(&input); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	isExist, err := h.services.AnimalType.IsExistByType(input.Type)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExist {
		h.sendConflict(c)
		return
	}

	animalType, err := h.services.AnimalType.Create(input.Type)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	output := &models.CreateAnimalTypeOutput{
		ID:   animalType.ID,
		Type: animalType.Value,
	}

	h.sendCreatedWithBody(c, output)
}

func (h *Handler) updateAnimalType(c *gin.Context) {
	id, err := getParamID(c, "typeId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	isExist, err := h.services.AnimalType.IsExistByID(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExist {
		h.sendNotFound(c)
		return
	}

	var input models.UpdateAnimalTypeInput
	if err = c.BindJSON(&input); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	isExist, err = h.services.AnimalType.IsExistByType(input.Type)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExist {
		h.sendConflict(c)
		return
	}

	animalType, err := h.services.AnimalType.Update(id, input.Type)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	output := &models.UpdateAnimalTypeOutput{
		ID:   animalType.ID,
		Type: animalType.Value,
	}

	h.sendCreatedWithBody(c, output)
}

func (h *Handler) removeAnimalType(c *gin.Context) {
	id, err := getParamID(c, "typeId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	isExist, err := h.services.AnimalType.IsExistByID(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExist {
		h.sendNotFound(c)
		return
	}

	isLinked, err := h.services.AnimalType.IsLinkedAnimal(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if isLinked {
		h.sendBadRequest(c, "the animal type is linked with the animal")
		return
	}

	if err = h.services.AnimalType.Remove(id); err != nil {
		h.sendInternalServerError(c)
		return
	}

	h.sendOk(c)
}
