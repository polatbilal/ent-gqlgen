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
	if currentStr, ok := current.(string); ok {
		if strings.Contains(fieldName, "Date") {
			newStr := fmt.Sprintf("%v", new)
			return CompareDates(currentStr, newStr)
		}
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

// Veri karşılaştırma fonksiyonları
func compareJobData(current *model.JobInput, new map[string]interface{}) map[string]interface{} {
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

func compareOwnerData(current *model.JobOwnerInput, new map[string]interface{}) map[string]interface{} {
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

func compareContractorData(current *model.JobContractorInput, new map[string]interface{}) map[string]interface{} {
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

func compareSupervisorData(current *model.JobSupervisorInput, new map[string]interface{}) map[string]interface{} {
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

func compareAuthorData(current *model.JobAuthorInput, new map[string]interface{}) map[string]interface{} {
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
