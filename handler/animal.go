package handler

import (
	"2023-it-planeta-web-api/constants"
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAnimal(c *gin.Context) {
	id, err := utils.GetNumberParam(c, "animalId")
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
		ChippingDateTime:   utils.ConvertDateToISO8601(animal.ChippingDate),
		ChipperID:          animal.Chipper,
		ChippingLocationID: animal.ChippingLocation,
		VisitedLocations:   visitedLocationsID,
		DeathDateTime:      utils.ConvertNullDateToISO8601(animal.DeathDate),
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
			ChippingDateTime:   utils.ConvertDateToISO8601(animal.ChippingDate),
			ChipperID:          animal.Chipper,
			ChippingLocationID: animal.ChippingLocation,
			VisitedLocations:   visitedLocationsID,
			DeathDateTime:      utils.ConvertNullDateToISO8601(animal.DeathDate),
		}
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

	visitedLocations, err := h.services.Animal.GetVisitedLocations(animal.ID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	visitedLocationID := make([]int64, len(visitedLocations))
	for i, location := range visitedLocations {
		visitedLocationID[i] = location.ID
	}

	animalsTypeID := make([]int64, len(input.AnimalTypes))
	for i, animalType := range input.AnimalTypes {
		animalsTypeID[i] = *animalType
	}

	output := models.CreateAnimalOutput{
		ID:                 animal.ID,
		AnimalTypes:        animalsTypeID,
		Weight:             animal.Weight,
		Length:             animal.Length,
		Height:             animal.Height,
		Gender:             animal.Gender,
		LifeStatus:         animal.LifeStatus,
		ChippingDateTime:   utils.ConvertDateToISO8601(animal.ChippingDate),
		ChipperID:          animal.Chipper,
		ChippingLocationID: animal.ChippingLocation,
		VisitedLocations:   visitedLocationID,
		DeathDateTime:      nil,
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
	if err = output.Load(h.services, animalID); err != nil {
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
