package handler

import (
	"2023-it-planeta-web-api/ctype"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
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

func getParamID(c *gin.Context, key string) (int64, error) {
	stringID := c.Param(key)
	if stringID == stringEmpty || stringID == stringNull {
		return 0, errors.New("id is not valid")
	}

	id, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		return 0, errors.New("id is not valid")
	}

	if id <= 0 {
		return 0, errors.New("id is not valid")
	}

	return id, nil
}
