package tools

import (
	"fmt"
	"time"
)

// ParseDate parses a date string in "2006-01-02" format and returns a *time.Time
func ParseDate(dateStr *string) (*time.Time, error) {
	if dateStr == nil {
		return nil, nil
	}

	parsedDate, err := time.Parse("2006-01-02", *dateStr)
	if err != nil {
		return nil, fmt.Errorf("tarih formatı hatası: %v", err)
	}
	return &parsedDate, nil
}
