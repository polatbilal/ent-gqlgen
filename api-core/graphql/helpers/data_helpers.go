package helpers

import (
	"context"
	"fmt"

	"github.com/polatbilal/gqlgen-ent/api-core/ent"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/companyengineer"
	"github.com/polatbilal/gqlgen-ent/api-core/graphql/model"
	"github.com/polatbilal/gqlgen-ent/api-core/middlewares"
)

func ValidateAndGetEngineers(ctx context.Context, input model.JobEngineerInput) (map[string]*ent.CompanyEngineer, error) {
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

// İlişki yönetimi için yardımcı fonksiyonlar
func HandleJobRelations(ctx context.Context, client *ent.Client, jobDetail *ent.JobDetail, company *ent.CompanyDetail, engineers map[string]*ent.CompanyEngineer) error {
	relations, err := jobDetail.QueryRelations().Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			// İlişki yoksa oluştur
			return createJobRelations(ctx, client, jobDetail, company, engineers)
		}
		return fmt.Errorf("iş ilişkileri bulunamadı: %v", err)
	}

	// İlişki varsa güncelle
	return updateJobRelations(ctx, client, relations, company, engineers)
}

func createJobRelations(ctx context.Context, client *ent.Client, jobDetail *ent.JobDetail, company *ent.CompanyDetail, engineers map[string]*ent.CompanyEngineer) error {
	builder := client.JobRelations.Create().
		SetJob(jobDetail).
		SetCompany(company)

	builder = setEngineersCreate(builder, engineers)

	_, err := builder.Save(ctx)
	if err != nil {
		return fmt.Errorf("iş ilişkileri oluşturulamadı: %v", err)
	}

	return nil
}

func updateJobRelations(ctx context.Context, client *ent.Client, relations *ent.JobRelations, company *ent.CompanyDetail, engineers map[string]*ent.CompanyEngineer) error {
	builder := client.JobRelations.UpdateOne(relations).
		SetCompany(company)

	builder = setEngineersUpdate(builder, engineers)

	_, err := builder.Save(ctx)
	if err != nil {
		return fmt.Errorf("iş ilişkileri güncellenemedi: %v", err)
	}

	return nil
}

func setEngineersCreate(builder *ent.JobRelationsCreate, engineers map[string]*ent.CompanyEngineer) *ent.JobRelationsCreate {
	if inspector, ok := engineers["inspector"]; ok {
		builder = builder.SetInspector(inspector)
	}
	if static, ok := engineers["static"]; ok {
		builder = builder.SetStatic(static)
	}
	if architect, ok := engineers["architect"]; ok {
		builder = builder.SetArchitect(architect)
	}
	if mechanic, ok := engineers["mechanic"]; ok {
		builder = builder.SetMechanic(mechanic)
	}
	if electric, ok := engineers["electric"]; ok {
		builder = builder.SetElectric(electric)
	}
	if controller, ok := engineers["controller"]; ok {
		builder = builder.SetController(controller)
	}
	if mechanicController, ok := engineers["mechanicController"]; ok {
		builder = builder.SetMechaniccontroller(mechanicController)
	}
	if electricController, ok := engineers["electricController"]; ok {
		builder = builder.SetElectriccontroller(electricController)
	}
	return builder
}

func setEngineersUpdate(builder *ent.JobRelationsUpdateOne, engineers map[string]*ent.CompanyEngineer) *ent.JobRelationsUpdateOne {
	if inspector, ok := engineers["inspector"]; ok {
		builder = builder.SetInspector(inspector)
	}
	if static, ok := engineers["static"]; ok {
		builder = builder.SetStatic(static)
	}
	if architect, ok := engineers["architect"]; ok {
		builder = builder.SetArchitect(architect)
	}
	if mechanic, ok := engineers["mechanic"]; ok {
		builder = builder.SetMechanic(mechanic)
	}
	if electric, ok := engineers["electric"]; ok {
		builder = builder.SetElectric(electric)
	}
	if controller, ok := engineers["controller"]; ok {
		builder = builder.SetController(controller)
	}
	if mechanicController, ok := engineers["mechanicController"]; ok {
		builder = builder.SetMechaniccontroller(mechanicController)
	}
	if electricController, ok := engineers["electricController"]; ok {
		builder = builder.SetElectriccontroller(electricController)
	}
	return builder
}
