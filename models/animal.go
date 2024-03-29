package models

import (
	"2023-it-planeta-web-api/constants"
	"2023-it-planeta-web-api/models_output"
	"2023-it-planeta-web-api/utils"
	"errors"
)

type GetAnimalOutput = models_output.OutputAnimal

type GetAnimalsInput struct {
	StartDateTime      *string `form:"startDateTime"`
	EndDateTime        *string `form:"endDateTime"`
	ChipperID          *int64  `form:"chipperId"`
	ChippingLocationID *int64  `form:"chippingLocationId"`
	LifeStatus         *string `form:"lifeStatus"`
	Gender             *string `form:"gender"`
	From               *int32  `form:"from"`
	Size               *int32  `form:"size"`
}

func (i *GetAnimalsInput) Validate() error {
	if i.From != nil && *i.From < 0 {
		return errors.New("invalid parameter from")
	}

	if i.Size != nil && *i.Size <= 0 {
		return errors.New("invalid parameter size")
	}

	if i.StartDateTime != nil && !utils.IsISO8601Date(*i.StartDateTime) {
		return errors.New("invalid parameter start date time")
	}

	if i.EndDateTime != nil && !utils.IsISO8601Date(*i.EndDateTime) {
		return errors.New("invalid parameter end date time")
	}

	if i.ChipperID != nil && *i.ChipperID <= 0 {
		return errors.New("invalid parameter chipper id")
	}

	if i.ChipperID != nil && *i.ChipperID <= 0 {
		return errors.New("invalid parameter chipping location")
	}

	if i.LifeStatus != nil {
		if !constants.IsAnimalLifeStatus(*i.LifeStatus) {
			return errors.New("invalid parameter life status")
		}
	}

	if i.Gender != nil {
		if !constants.IsAnimalGender(*i.Gender) {
			return errors.New("invalid parameter gender")
		}
	}

	return nil
}

type GetAnimalsOutput = []*models_output.OutputAnimal

type CreateAnimalInput struct {
	AnimalTypes        []*int64 `json:"animalTypes" binding:"required"`
	Weight             float64  `json:"weight" binding:"required,min=0"`
	Length             float64  `json:"length" binding:"required,min=0"`
	Height             float64  `json:"height" binding:"required,min=0"`
	Gender             string   `json:"gender" binding:"required"`
	ChipperID          int64    `json:"chipperId" binding:"required,min=1"`
	ChippingLocationID int64    `json:"chippingLocationId" binding:"required,min=1"`
}

func (i *CreateAnimalInput) Validate() error {
	if len(i.AnimalTypes) <= 0 {
		return errors.New("invalid field animal types")
	}

	for _, animalType := range i.AnimalTypes {
		if animalType == nil {
			return errors.New("invalid element in animal types")
		}
		if *animalType <= 0 {
			return errors.New("invalid value element in animal types")
		}
	}

	if !constants.IsAnimalGender(i.Gender) {
		return errors.New("invalid gender value")
	}

	return nil
}

type CreateAnimalOutput = models_output.OutputAnimal

type UpdateAnimalInput struct {
	Weight             float64 `json:"weight" binding:"required,min=1"`
	Length             float64 `json:"length" binding:"required,min=1"`
	Height             float64 `json:"height" binding:"required,min=1"`
	Gender             string  `json:"gender" binding:"required"`
	LifeStatus         string  `json:"lifeStatus" binding:"required"`
	ChipperID          int64   `json:"chipperId" binding:"required,min=1"`
	ChippingLocationID int64   `json:"chippingLocationId" binding:"required,min=1"`
}

func (i *UpdateAnimalInput) Validate() error {
	if !constants.IsAnimalLifeStatus(i.LifeStatus) {
		return errors.New("invalid field life status")
	}

	if !constants.IsAnimalGender(i.Gender) {
		return errors.New("invalid field gender")
	}

	return nil
}

type UpdateAnimalOutput = models_output.OutputAnimal
type LinkAnimalTypeToAnimalOutput = models_output.OutputAnimal

type UpdateAnimalTypeToAnimalInput struct {
	OldTypeId int64 `json:"oldTypeId" binding:"required,min=1"`
	NewTypeId int64 `json:"newTypeId" binding:"required,min=1"`
}

type UpdateAnimalTypeToAnimalOutput = models_output.OutputAnimal
type RemoveAnimalTypeToAnimalOutput = models_output.OutputAnimal
