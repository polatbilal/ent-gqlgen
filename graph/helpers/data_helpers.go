package helpers

import (
	"context"
	"fmt"

	"github.com/polatbilal/gqlgen-ent/ent"
	"github.com/polatbilal/gqlgen-ent/ent/companyengineer"
	"github.com/polatbilal/gqlgen-ent/graph/model"
	"github.com/polatbilal/gqlgen-ent/middlewares"
)

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
