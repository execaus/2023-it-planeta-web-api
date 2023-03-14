package handler

import (
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAnimalType(c *gin.Context) {
	id, err := utils.GetNumberParam(c, "typeId")
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
	id, err := utils.GetNumberParam(c, "typeId")
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
	id, err := utils.GetNumberParam(c, "typeId")
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

func (h *Handler) linkAnimalTypeToAnimal(c *gin.Context) {
	animalID, err := utils.GetNumberParam(c, "animalId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	typeID, err := utils.GetNumberParam(c, "typeId")
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

	// 404 Тип животного с typeId не найден
	isExistType, err := h.services.AnimalType.IsExistByID(typeID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExistType {
		h.sendNotFound(c)
		return
	}

	// 409 Тип животного с typeId уже есть у животного с animalId
	isLinked, err := h.services.Animal.IsLinkedAnimalType(animalID, typeID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if isLinked {
		h.sendConflict(c)
		return
	}

	if err = h.services.Animal.LinkAnimalType(animalID, typeID); err != nil {
		h.sendInternalServerError(c)
		return
	}

	var output models.LinkAnimalTypeToAnimalOutput
	if err = output.Load(h.services, animalID); err != nil {
		h.sendInternalServerError(c)
		return
	}

	h.sendOKWithBody(c, &output)
}

func (h *Handler) updateAnimalTypeToAnimal(c *gin.Context) {
	animalID, err := utils.GetNumberParam(c, "animalId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	var input models.UpdateAnimalTypeToAnimalInput
	if err = c.BindJSON(&input); err != nil {
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

	// 404 Тип животного с oldTypeId не найден
	isExistOldType, err := h.services.AnimalType.IsExistByID(input.OldTypeId)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}
	if !isExistOldType {
		h.sendNotFound(c)
		return
	}

	// 404 Тип животного с newTypeId не найден
	isExistNewType, err := h.services.AnimalType.IsExistByID(input.OldTypeId)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}
	if !isExistNewType {
		h.sendNotFound(c)
		return
	}

	// 404 Типа животного с oldTypeId нет у животного с animalId
	isAnimalHaveOldType, err := h.services.Animal.IsLinkedAnimalType(animalID, input.OldTypeId)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}
	if !isAnimalHaveOldType {
		h.sendNotFound(c)
		return
	}

	// 409 Тип животного с newTypeId уже есть у животного с animalId
	isAnimalHaveNewType, err := h.services.Animal.IsLinkedAnimalType(animalID, input.NewTypeId)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}
	if isAnimalHaveNewType {
		h.sendConflict(c)
		return
	}

	// 409 Животное с animalId уже имеет типы с oldTypeId и newTypeId
	if isAnimalHaveOldType && isAnimalHaveNewType {
		h.sendConflict(c)
		return
	}

	if err = h.services.Animal.UpdateAnimalTypeToAnimal(animalID, &input); err != nil {
		h.sendInternalServerError(c)
		return
	}

	var output models.UpdateAnimalTypeToAnimalOutput
	if err = output.Load(h.services, animalID); err != nil {
		h.sendInternalServerError(c)
		return
	}

	h.sendOKWithBody(c, &output)
}
