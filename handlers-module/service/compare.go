package service

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/polatbilal/gqlgen-ent/api-core/graphql/model"
)

// Yardımcı fonksiyonlar
func getFieldValue(obj interface{}, fieldName string) interface{} {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		return nil
	}

	return field.Interface()
}

func compareValues(current, new interface{}, fieldName string) bool {
	if current == nil && new == nil {
		return true
	}
	if current == nil || new == nil {
		return false
	}

	// Tarih karşılaştırması
	if strings.Contains(fieldName, "Date") {
		// Unix timestamp kontrolü ve dönüşümü
		var currentDateStr, newDateStr string

		// Current değer için kontrol (GraphQL'den gelen DD/MM/YYYY formatında)
		switch v := current.(type) {
		case string:
			if strings.Contains(v, "/") {
				// DD/MM/YYYY formatını YYYY-MM-DD'ye çevir
				parts := strings.Split(v, "/")
				if len(parts) == 3 {
					currentDateStr = fmt.Sprintf("%s-%s-%s", parts[2], parts[1], parts[0])
				} else {
					currentDateStr = v
				}
			} else {
				currentDateStr = v
			}
		case float64:
			currentDateStr = SafeUnixToDate(int64(v))
		case int64:
			currentDateStr = SafeUnixToDate(v)
		case int:
			currentDateStr = SafeUnixToDate(int64(v))
		default:
			currentDateStr = fmt.Sprintf("%v", current)
		}

		// New değer için kontrol (YDK'dan gelen Unix timestamp)
		switch v := new.(type) {
		case string:
			if strings.Contains(v, "/") {
				// DD/MM/YYYY formatını YYYY-MM-DD'ye çevir
				parts := strings.Split(v, "/")
				if len(parts) == 3 {
					newDateStr = fmt.Sprintf("%s-%s-%s", parts[2], parts[1], parts[0])
				} else {
					newDateStr = v
				}
			} else {
				newDateStr = v
			}
		case float64:
			newDateStr = SafeUnixToDate(int64(v))
		case int64:
			newDateStr = SafeUnixToDate(v)
		case int:
			newDateStr = SafeUnixToDate(int64(v))
		default:
			newDateStr = fmt.Sprintf("%v", new)
		}

		// Her iki tarih de YYYY-MM-DD formatına dönüştürüldü, şimdi karşılaştır
		return CompareDates(currentDateStr, newDateStr)
	}

	// Koordinat karşılaştırması
	if fieldName == "Coordinates" {
		currentStr := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%v", current), " ", ""), ",", ".")
		newStr := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%v", new), " ", ""), ",", ".")
		return currentStr == newStr
	}

	// Genel karşılaştırma
	return fmt.Sprintf("%v", current) == fmt.Sprintf("%v", new)
}

// CompareJobData iş verilerini karşılaştırır
func CompareJobData(current *model.JobInput, new map[string]interface{}) map[string]interface{} {
	changes := make(map[string]interface{})

	for key, newValue := range new {
		if key == "YibfNo" || key == "CompanyCode" {
			continue
		}

		currentValue := getFieldValue(current, key)
		if !compareValues(currentValue, newValue, key) {
			changes[key] = newValue
		}
	}

	return changes
}

// CompareOwnerData mal sahibi verilerini karşılaştırır
func CompareOwnerData(current *model.JobOwnerInput, new map[string]interface{}) map[string]interface{} {
	changes := make(map[string]interface{})

	for key, newValue := range new {
		if key == "YibfNo" {
			continue
		}

		currentValue := getFieldValue(current, key)
		if !compareValues(currentValue, newValue, key) {
			changes[key] = newValue
		}
	}

	return changes
}

// CompareContractorData müteahhit verilerini karşılaştırır
func CompareContractorData(current *model.JobContractorInput, new map[string]interface{}) map[string]interface{} {
	changes := make(map[string]interface{})

	for key, newValue := range new {
		if key == "YibfNo" {
			continue
		}

		currentValue := getFieldValue(current, key)
		if !compareValues(currentValue, newValue, key) {
			changes[key] = newValue
		}
	}

	return changes
}

// CompareSupervisorData şantiye şefi verilerini karşılaştırır
func CompareSupervisorData(current *model.JobSupervisorInput, new map[string]interface{}) map[string]interface{} {
	changes := make(map[string]interface{})

	for key, newValue := range new {
		if key == "YibfNo" {
			continue
		}

		currentValue := getFieldValue(current, key)
		if !compareValues(currentValue, newValue, key) {
			changes[key] = newValue
		}
	}

	return changes
}

// CompareAuthorData proje müellifi verilerini karşılaştırır
func CompareAuthorData(current *model.JobAuthorInput, new map[string]interface{}) map[string]interface{} {
	changes := make(map[string]interface{})

	for key, newValue := range new {
		if key == "YibfNo" {
			continue
		}

		currentValue := getFieldValue(current, key)
		if !compareValues(currentValue, newValue, key) {
			changes[key] = newValue
		}
	}

	return changes
}

// CompareEngineerData denetçi verilerini karşılaştırır
func CompareEngineerData(current *model.JobEngineerInput, new map[string]interface{}) map[string]interface{} {
	changes := make(map[string]interface{})

	for key, newValue := range new {
		if key == "YibfNo" {
			continue
		}

		currentValue := getFieldValue(current, key)
		if !compareValues(currentValue, newValue, key) {
			changes[key] = newValue
		}
	}

	return changes
}
