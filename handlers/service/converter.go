package service

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// String'i int'e dönüştürür, hata durumunda 0 döner
func SafeStringToInt(s string) int {
	if s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		// Bilimsel gösterimden normal formata çevir
		if strings.Contains(s, "e+") {
			if f, err := strconv.ParseFloat(s, 64); err == nil {
				return int(f)
			}
		}
		return 0
	}
	return i
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

// Sayısal değerleri karşılaştırmak için yardımcı fonksiyon
func CompareValues(v1, v2 interface{}) bool {
	// Nil kontrolü
	if v1 == nil || v2 == nil {
		return v1 == v2
	}

	// String karşılaştırması
	s1, ok1 := v1.(string)
	s2, ok2 := v2.(string)
	if ok1 && ok2 {
		return strings.TrimSpace(s1) == strings.TrimSpace(s2)
	}

	// Sayısal değer karşılaştırması
	n1, err1 := ConvertToFloat64(v1)
	n2, err2 := ConvertToFloat64(v2)
	if err1 == nil && err2 == nil {
		return math.Abs(n1-n2) < 0.000001 // Küçük farkları göz ardı et
	}

	// Diğer tipleri string'e çevirip karşılaştır
	return fmt.Sprintf("%v", v1) == fmt.Sprintf("%v", v2)
}

// Herhangi bir değeri float64'e çevirmeye çalışır
func ConvertToFloat64(v interface{}) (float64, error) {
	switch val := v.(type) {
	case float64:
		return val, nil
	case float32:
		return float64(val), nil
	case int:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case string:
		// Bilimsel notasyonu ve normal sayıları parse et
		return strconv.ParseFloat(strings.TrimSpace(val), 64)
	}
	// Diğer tipleri string'e çevirip parse etmeyi dene
	str := fmt.Sprintf("%v", v)
	return strconv.ParseFloat(strings.TrimSpace(str), 64)
}

// Tarih formatını dönüştürür (YYYY-MM-DD -> DD/MM/YYYY)
func ConvertDateFormat(date string) string {
	if date == "" {
		return ""
	}
	parts := strings.Split(date, "-")
	if len(parts) == 3 {
		return fmt.Sprintf("%s/%s/%s", parts[2], parts[1], parts[0])
	}
	return date
}

// Alan tipini kontrol eden yardımcı fonksiyonlar
func IsDateField(field string) bool {
	dateFields := map[string]bool{
		"VisaDate":       true,
		"VisaEndDate":    true,
		"ContractDate":   true,
		"LicenseDate":    true,
		"CompletionDate": true,
		"StartDate":      true,
	}
	return dateFields[field]
}

func IsNumericField(field string) bool {
	numericFields := map[string]bool{
		"CompanyCode": true,
	}
	return numericFields[field]
}

func IsStringNumericField(field string) bool {
	stringNumericFields := map[string]bool{
		"TaxNo":     true,
		"OwnerTcNo": true,
		"TcNo":      true,
	}
	return stringNumericFields[field]
}

// Değer karşılaştırma fonksiyonu
func CompareFieldValues(key string, currentValue, newValue interface{}) (bool, string) {
	if newValue == nil && currentValue == nil {
		return false, ""
	}

	if IsDateField(key) {
		newDateStr := fmt.Sprintf("%v", newValue)
		currentDateStr := fmt.Sprintf("%v", currentValue)

		if newValue != nil && newDateStr != "" {
			newDateStr = ConvertDateFormat(newDateStr)
		}

		if !CompareDates(currentDateStr, newDateStr) {
			return true, fmt.Sprintf("Değişiklik tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentValue, newValue)
		}
		return false, ""
	}

	if IsStringNumericField(key) {
		newStr := fmt.Sprintf("%v", newValue)
		currentStr := fmt.Sprintf("%v", currentValue)

		if strings.Contains(currentStr, "e+") {
			if f, err := strconv.ParseFloat(currentStr, 64); err == nil {
				currentStr = fmt.Sprintf("%.0f", f)
			}
		}

		if strings.TrimSpace(newStr) != strings.TrimSpace(currentStr) {
			return true, fmt.Sprintf("Değişiklik tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentValue, newValue)
		}
		return false, ""
	}

	if IsNumericField(key) {
		n1, err1 := ConvertToFloat64(currentValue)
		n2, err2 := ConvertToFloat64(newValue)
		if err1 == nil && err2 == nil && math.Abs(n1-n2) > 0.001 {
			return true, fmt.Sprintf("Değişiklik tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentValue, newValue)
		}
		return false, ""
	}

	// Diğer alanlar için string karşılaştırması
	if !CompareValues(currentValue, newValue) {
		return true, fmt.Sprintf("Değişiklik tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentValue, newValue)
	}
	return false, ""
}
