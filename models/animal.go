package models

import "2023-it-planeta-web-api/ctype"

type GetAnimalOutput struct {
	ID                 int64
	AnimalTypes        []int64
	Weight             float64
	Length             float64
	Height             float64
	Gender             string
	LifeStatus         string
	ChippingDateTime   string
	ChipperID          int64
	ChippingLocationID int64
	VisitedLocations   []int64
	DeathDateTime      ctype.TimeOrNil
}

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

type GetAnimalsOutput struct {
	ID                 int64
	AnimalTypes        []int64
	Weight             float64
	Length             float64
	Height             float64
	Gender             string
	LifeStatus         string
	ChippingDateTime   string
	ChipperID          int64
	ChippingLocationID int64
	VisitedLocations   []int64
	DeathDateTime      ctype.TimeOrNil
}
