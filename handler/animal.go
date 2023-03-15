package handler

import (
	"2023-it-planeta-web-api/constants"
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/models_output"
	"2023-it-planeta-web-api/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAnimal(c *gin.Context) {
	animalID, err := utils.GetNumberParam(c, "animalId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	var output models.GetAnimalOutput
	if err = h.services.Animal.FillAnimalOutput(&output, h.services, animalID); err != nil {
		h.sendInternalServerError(c)
		return
	}

	h.sendOKWithBody(c, &output)
}

func (h *Handler) getAnimals(c *gin.Context) {
	var input models.GetAnimalsInput

	if err := c.ShouldBindQuery(&input); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	animals, err := h.services.Animal.GetList(&input)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	var output models.GetAnimalsOutput
	for i, animal := range animals {
		var outputAnimal models_output.OutputAnimal
		if err = h.services.Animal.FillAnimalOutput(&outputAnimal, h.services, animal.ID); err != nil {
			h.sendInternalServerError(c)
			return
		}

		output[i] = &outputAnimal
	}

	h.sendOKWithBody(c, output)
}

func (h *Handler) createAnimal(c *gin.Context) {
	var input models.CreateAnimalInput

	if err := c.BindJSON(&input); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	// 409 Массив animalTypes содержит дубликаты
	hasDuplicates := utils.HasDuplicates(input.AnimalTypes)
	if hasDuplicates {
		h.sendConflict(c)
		return
	}

	// 404 Тип животного не найден
	for _, animalType := range input.AnimalTypes {
		isExistAnimal, err := h.services.AnimalType.IsExistByID(*animalType)
		if err != nil {
			h.sendInternalServerError(c)
			return
		}
		if !isExistAnimal {
			h.sendNotFound(c)
			return
		}
	}

	// 404 Аккаунт с chipperId не найден
	isExistChipper, err := h.services.Account.IsExistByID(input.ChipperID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExistChipper {
		h.sendNotFound(c)
		return
	}

	// 404 Точка локации с chippingLocationId не найдена
	isExistChippingLocation, err := h.services.Location.IsExistByID(input.ChippingLocationID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExistChippingLocation {
		h.sendNotFound(c)
		return
	}

	animal, err := h.services.Animal.Create(&input)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	var output models.CreateAnimalOutput
	if err = h.services.Animal.FillAnimalOutput(&output, h.services, animal.ID); err != nil {
		h.sendInternalServerError(c)
		return
	}

	h.sendOKWithBody(c, &output)
}

func (h *Handler) updateAnimal(c *gin.Context) {
	animalID, err := utils.GetNumberParam(c, "animalId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	var input models.UpdateAnimalInput
	if err = c.BindJSON(&input); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	if err = input.Validate(); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	// 400 Установка lifeStatus = “ALIVE”, если у животного lifeStatus = “DEAD”
	currentAnimal, err := h.services.Animal.Get(animalID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if currentAnimal.LifeStatus == constants.AnimalLifeStatusDead &&
		input.LifeStatus == constants.AnimalLifeStatusAlive {
		h.sendBadRequest(c, "invalid value life status")
		return
	}

	// 400 Новая точка чипирования совпадает с первой посещенной точкой локации
	visitedLocations, err := h.services.Animal.GetVisitedLocations(animalID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if len(visitedLocations) > 0 && visitedLocations[0].Location == input.ChippingLocationID {
		h.sendBadRequest(c, "the new chipping point coincides with the first visited location point")
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

	// 404 Аккаунт с chipperId не найден
	isExistChipper, err := h.services.Account.IsExistByID(input.ChipperID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExistChipper {
		h.sendNotFound(c)
		return
	}

	// 404 Точка локации с chippingLocationId не найдена
	isExistChippingLocation, err := h.services.Location.IsExistByID(input.ChippingLocationID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExistChippingLocation {
		h.sendNotFound(c)
		return
	}

	_, err = h.services.Animal.Update(animalID, &input)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	var output models.UpdateAnimalOutput
	if err = h.services.Animal.FillAnimalOutput(&output, h.services, animalID); err != nil {
		h.sendInternalServerError(c)
		return
	}

	h.sendOKWithBody(c, &output)
}

func (h *Handler) removeAnimal(c *gin.Context) {
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

	// 400 Животное покинуло локацию чипирования, при этом есть другие посещенные точки
	visitedLocations, err := h.services.Animal.GetVisitedLocations(animalID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}
	if len(visitedLocations) > 0 {
		h.sendBadRequest(c, "the animal has left the chipping location and there are other visited locations")
		return
	}

	if err = h.services.Animal.Remove(animalID); err != nil {
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
	if err = h.services.Animal.FillAnimalOutput(&output, h.services, animalID); err != nil {
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
	if err = h.services.Animal.FillAnimalOutput(&output, h.services, animalID); err != nil {
		h.sendInternalServerError(c)
		return
	}

	h.sendOKWithBody(c, &output)
}

func (h *Handler) removeAnimalTypeToAnimal(c *gin.Context) {
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
	isExistType, err := h.services.AnimalType.IsExistByID(animalID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExistType {
		h.sendNotFound(c)
		return
	}

	// 404 У животного с animalId нет типа с typeId
	isLinkedType, err := h.services.Animal.IsLinkedAnimalType(animalID, typeID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isLinkedType {
		h.sendNotFound(c)
		return
	}

	// 400 У животного только один тип и это тип с typeId
	animalTypes, err := h.services.AnimalType.GetByAnimalID(animalID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if len(animalTypes) == 1 && animalTypes[0].AnimalType == typeID {
		h.sendBadRequest(c, "the animal has only one type and it is the type with typeId")
		return
	}

	if err = h.services.Animal.RemoveAnimalType(animalID, typeID); err != nil {
		h.sendInternalServerError(c)
		return
	}

	var output models.RemoveAnimalTypeToAnimalOutput
	if err = h.services.Animal.FillAnimalOutput(&output, h.services, animalID); err != nil {
		h.sendInternalServerError(c)
		return
	}

	h.sendOKWithBody(c, &output)
}
