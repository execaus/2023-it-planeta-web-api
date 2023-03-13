package models

type CreateVisitedLocationOutput struct {
	ID                           int64  `json:"id"`
	DateTimeOfVisitLocationPoint string `json:"dateTimeOfVisitLocationPoint"`
	LocationPointID              int64  `json:"locationPointId"`
}

type UpdateVisitedLocationInput struct {
	VisitedLocationPointID int64 `json:"visitedLocationPointId"`
	LocationPointID        int64 `json:"locationPointId"`
}

type UpdateVisitedLocationOutput struct {
	ID                           int64  `json:"id"`
	DateTimeOfVisitLocationPoint string `json:"dateTimeOfVisitLocationPoint"`
	LocationPointID              int64  `json:"locationPointId"`
}
