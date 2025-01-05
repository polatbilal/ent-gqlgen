package helpers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/polatbilal/gqlgen-ent/ent"
	"github.com/polatbilal/gqlgen-ent/ent/companyengineer"
	"github.com/polatbilal/gqlgen-ent/graph/model"
	"github.com/polatbilal/gqlgen-ent/middlewares"
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

func ValidateAndGetEngineers(ctx context.Context, input model.JobInput) (map[string]*ent.CompanyEngineer, error) {
	client := middlewares.GetClientFromContext(ctx)
	engineers := make(map[string]*ent.CompanyEngineer)
	var engineerErr error

	engineerChecks := map[string]*int{
		"inspector":          input.Inspector,
		"static":             input.Static,
		"architect":          input.Architect,
		"mechanic":           input.Mechanic,
		"electric":           input.Electric,
		"controller":         input.Controller,
		"mechanicController": input.MechanicController,
		"electricController": input.ElectricController,
	}

	errorMessages := map[string]string{
		"inspector":          "denetçi",
		"static":             "statik",
		"architect":          "mimar",
		"mechanic":           "mak. müh.",
		"electric":           "elektrik",
		"controller":         "kont. elm.",
		"mechanicController": "mekanik kont. elm.",
		"electricController": "elektrik kont. elm.",
	}

	for key, id := range engineerChecks {
		if id != nil {
			engineers[key], engineerErr = client.CompanyEngineer.Query().
				Where(companyengineer.YDSIDEQ(*id)).
				Only(ctx)
			if engineerErr != nil {
				return nil, fmt.Errorf("%s bulunamadı (kod: %d): %v", errorMessages[key], *id, engineerErr)
			}
		}
	}

	return engineers, nil
}
