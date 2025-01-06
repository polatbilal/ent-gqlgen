package service

import (
	"fmt"
	"strconv"
	"time"
)

// String'i int'e dönüştürür, hata durumunda 0 döner
func SafeStringToInt(s string) int {
	if s == "" {
		return 0
	}
	if val, err := strconv.Atoi(s); err == nil {
		return val
	}
	return 0
}

// Unix timestamp'i tarihe dönüştürür, geçersiz timestamp için boş string döner
func SafeUnixToDate(timestamp int64) string {
	if timestamp > 0 {
		// Milisaniyeyi saniyeye çevir
		seconds := timestamp / 1000
		return time.Unix(seconds, 0).Local().Format("2006-01-02")
	}
	return ""
}

// Int'i string'e dönüştürür
func SafeIntToString(i int) string {
	return strconv.Itoa(i)
}

// Koordinat array'ini string'e dönüştürür
func CoordinatesToString(coords []float64) string {
	if len(coords) >= 2 {
		return fmt.Sprintf("%.6f, %.6f", coords[0], coords[1])
	}
	return ""
}
