package utils

import (
	"2023-it-planeta-web-api/ctypes"
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

func ConvertDateToISO8601(date time.Time) string {
	return date.Format(time.RFC3339)
}

func ConvertNullDateToISO8601(date sql.NullTime) ctypes.TimeOrNil {
	if !date.Valid {
		return nil
	}
	return ConvertDateToISO8601(date.Time)
}

func IsISO8601Date(str string) bool {
	_, err := time.Parse(time.RFC3339, str)
	return err == nil
}

func GetNumberParam(c *gin.Context, key string) (int64, error) {
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

func HasDuplicates(collection []*int64) bool {
	seen := make(map[int64]bool)
	for _, element := range collection {
		if seen[*element] == true {
			return true // Дубликат найден
		} else {
			seen[*element] = true
		}
	}
	return false // Дубликатов нет
}
