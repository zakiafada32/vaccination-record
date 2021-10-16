package utils

import (
	"log"
	"time"
)

func ConvertStringToDate(s string) time.Time {
	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		log.Println("cannot convert string to date")
	}

	return date
}

func ConvertDateToString(t time.Time) string {
	data := t.Format("2006-01-02")
	return data
}
