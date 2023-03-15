package models_output

import (
	"2023-it-planeta-web-api/ctypes"
)

type OutputAnimal struct {
	ID                 int64            `json:"id"`
	AnimalTypes        []int64          `json:"animalTypes"`
	Weight             float64          `json:"weight"`
	Length             float64          `json:"length"`
	Height             float64          `json:"height"`
	Gender             string           `json:"gender"`
	LifeStatus         string           `json:"lifeStatus"`
	ChippingDateTime   string           `json:"chippingDateTime"`
	ChipperID          int64            `json:"chipperId"`
	ChippingLocationID int64            `json:"chippingLocationId"`
	VisitedLocations   []int64          `json:"visitedLocations"`
	DeathDateTime      ctypes.TimeOrNil `json:"deathDateTime"`
}
