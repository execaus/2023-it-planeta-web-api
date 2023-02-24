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
