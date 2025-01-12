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
		return fmt.Sprintf("%v,%v", coords[0], coords[1])
	}
	return ""
}

// Tarihleri karşılaştırmak için yardımcı fonksiyon
func CompareDates(date1, date2 string) bool {
	// İlk tarihi parse et (dd/MM/yyyy formatı)
	t1, err1 := time.Parse("02/01/2006", date1)
	if err1 != nil {
		// yyyy-MM-dd formatını dene
		t1, err1 = time.Parse("2006-01-02", date1)
		if err1 != nil {
			return false
		}
	}

	// İkinci tarihi parse et
	t2, err2 := time.Parse("02/01/2006", date2)
	if err2 != nil {
		// yyyy-MM-dd formatını dene
		t2, err2 = time.Parse("2006-01-02", date2)
		if err2 != nil {
			return false
		}
	}

	// Unix timestamp'lerini karşılaştır
	return t1.Unix() == t2.Unix()
}
