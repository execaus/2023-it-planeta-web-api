package models

type CreateVisitedLocationOutput struct {
	ID                           int64  `json:"ID"`
	DateTimeOfVisitLocationPoint string `json:"dateTimeOfVisitLocationPoint"`
	LocationPointID              int64  `json:"locationPointId"`
}
