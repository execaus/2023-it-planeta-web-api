package handler

import (
	"2023-it-planeta-web-api/ctype"
	"database/sql"
	"time"
)

const (
	stringNull  = "null"
	stringEmpty = ""
)

func convertDateToISO8601(date time.Time) string {
	return date.Format(time.RFC3339)
}

func convertNullDateToISO8601(date sql.NullTime) ctype.TimeOrNil {
	if !date.Valid {
		return nil
	}
	return convertDateToISO8601(date.Time)
}
