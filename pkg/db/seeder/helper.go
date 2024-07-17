package seeder

import (
	"log"
	"time"
)

func stringToPointer(s string) *string {
	return &s
}

func stringToDate(s string) *time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		log.Fatalf("Error parsing date: %v", err)
	}
	return &t
}
