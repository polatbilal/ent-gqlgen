package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"
	"fmt"
	"strconv"

	"github.com/polatbilal/gqlgen-ent/ent"
	"github.com/polatbilal/gqlgen-ent/ent/companydetail"
	"github.com/polatbilal/gqlgen-ent/ent/companyengineer"
	"github.com/polatbilal/gqlgen-ent/ent/jobdetail"
	"github.com/polatbilal/gqlgen-ent/ent/joblayer"
	"github.com/polatbilal/gqlgen-ent/ent/jobsupervisor"
	"github.com/polatbilal/gqlgen-ent/graph/generated"
	"github.com/polatbilal/gqlgen-ent/graph/model"
	"github.com/polatbilal/gqlgen-ent/middlewares"
)

// Supervisor is the resolver for the Supervisor field.
func (r *jobDetailResolver) Supervisor(ctx context.Context, obj *ent.JobDetail) (*model.Supervisor, error) {
	client := middlewares.GetClientFromContext(ctx)
	supervisor, err := client.JobSuperVisor.Query().
		Where(jobsupervisor.HasSupervisorsWith(jobdetail.IDEQ(obj.ID))).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("supervisor bulunamadı: %v", err)
	}
	return &model.Supervisor{ID: strconv.Itoa(supervisor.ID)}, nil
}

// Layer is the resolver for the Layer field.
func (r *jobDetailResolver) Layer(ctx context.Context, obj *ent.JobDetail) ([]*ent.JobLayer, error) {
	client := middlewares.GetClientFromContext(ctx)
	return client.JobLayer.Query().Where(joblayer.HasLayerWith(jobdetail.IDEQ(obj.ID))).All(ctx)
}

// CreateJob is the resolver for the createJob field.
func (r *mutationResolver) CreateJob(ctx context.Context, input model.JobInput) (*ent.JobDetail, error) {
	client := middlewares.GetClientFromContext(ctx)

	// YibfNo kontrolü
	_, err := client.JobDetail.Query().Where(jobdetail.YibfNoEQ(*input.YibfNo)).Only(ctx)
	if err == nil {
		return nil, fmt.Errorf("iş zaten mevcut")
	}

	if !ent.IsNotFound(err) {
		return nil, err
	}

	// CompanyCode ile şirketi bul
	company, err := client.CompanyDetail.Query().
		Where(companydetail.CompanyCodeEQ(input.CompanyCode)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("şirket bulunamadı (kod: %d): %v", input.CompanyCode, err)
	}

	var (
		inspector, static, architect, mechanic, electric, controller, mechanicController, electricController *ent.CompanyEngineer
		engineerErr                                                                                          error
	)

	if input.Inspector != nil {
		inspector, engineerErr = client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.Inspector)).
			Only(ctx)
		if engineerErr != nil {
			return nil, fmt.Errorf("denetçi bulunamadı (kod: %d): %v", *input.Inspector, engineerErr)
		}
	}

	if input.Static != nil {
		static, engineerErr = client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.Static)).
			Only(ctx)
		if engineerErr != nil {
			return nil, fmt.Errorf("statik bulunamadı (kod: %d): %v", *input.Static, engineerErr)
		}
	}

	if input.Architect != nil {
		architect, engineerErr = client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.Architect)).
			Only(ctx)
		if engineerErr != nil {
			return nil, fmt.Errorf("mimar bulunamadı (kod: %d): %v", *input.Architect, engineerErr)
		}
	}

	if input.Mechanic != nil {
		mechanic, engineerErr = client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.Mechanic)).
			Only(ctx)
		if engineerErr != nil {
			return nil, fmt.Errorf("mak. müh. bulunamadı (kod: %d): %v", *input.Mechanic, engineerErr)
		}
	}

	if input.Electric != nil {
		electric, engineerErr = client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.Electric)).
			Only(ctx)
		if engineerErr != nil {
			return nil, fmt.Errorf("elektrik bulunamadı (kod: %d): %v", *input.Electric, engineerErr)
		}
	}

	if input.Controller != nil {
		controller, engineerErr = client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.Controller)).
			Only(ctx)
		if engineerErr != nil {
			return nil, fmt.Errorf("kont. elm. bulunamadı (kod: %d): %v", *input.Controller, engineerErr)
		}
	}

	if input.MechanicController != nil {
		mechanicController, engineerErr = client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.MechanicController)).
			Only(ctx)
		if engineerErr != nil {
			return nil, fmt.Errorf("mekanik kont. elm. bulunamadı (kod: %d): %v", *input.MechanicController, engineerErr)
		}
	}

	if input.ElectricController != nil {
		electricController, engineerErr = client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.ElectricController)).
			Only(ctx)
		if engineerErr != nil {
			return nil, fmt.Errorf("elektrik kont. elm. bulunamadı (kod: %d): %v", *input.ElectricController, engineerErr)
		}
	}

	// İş detayını oluştur
	jobDetailCreate := client.JobDetail.Create().
		SetYibfNo(*input.YibfNo).
		SetCompany(company).
		SetNillableContractDate(input.ContractDate).
		SetNillableStartDate(input.StartDate).
		SetNillableCompletionDate(input.CompletionDate).
		SetNillableLicenseDate(input.LicenseDate).
		SetNillableLicenseNo(input.LicenseNo).
		SetNillableAddress(input.Address).
		SetNillableBuildingClass(input.BuildingClass).
		SetNillableBuildingType(input.BuildingType).
		SetNillableNote(input.Note)

	if inspector != nil {
		jobDetailCreate = jobDetailCreate.SetInspector(inspector)
	}

	if static != nil {
		jobDetailCreate = jobDetailCreate.SetStatic(static)
	}

	if architect != nil {
		jobDetailCreate = jobDetailCreate.SetArchitect(architect)
	}

	if mechanic != nil {
		jobDetailCreate = jobDetailCreate.SetMechanic(mechanic)
	}

	if electric != nil {
		jobDetailCreate = jobDetailCreate.SetElectric(electric)
	}

	if controller != nil {
		jobDetailCreate = jobDetailCreate.SetController(controller)
	}

	if mechanicController != nil {
		jobDetailCreate = jobDetailCreate.SetMechaniccontroller(mechanicController)
	}

	if electricController != nil {
		jobDetailCreate = jobDetailCreate.SetElectriccontroller(electricController)
	}

	newJobDetail, err := jobDetailCreate.Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("iş ayrıntısı oluşturulamadı: %v", err)
	}

	// Progress bilgilerini iş ayrıntısına ekle
	p, err := client.JobProgress.Create().
		SetOne(0).
		SetTwo(0).
		SetThree(0).
		SetFour(0).
		SetFive(0).
		SetSix(0).
		AddProgress(newJobDetail).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("progress oluşturulamadı: %v, %s", err, p)
	}
	return newJobDetail, nil
}

// UpdateJob is the resolver for the updateJob field.
func (r *mutationResolver) UpdateJob(ctx context.Context, yibfNo int, input model.JobInput) (*ent.JobDetail, error) {
	client := middlewares.GetClientFromContext(ctx)
	jobDetail, err := client.JobDetail.Query().Where(jobdetail.YibfNoEQ(yibfNo)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("iş ayrıntısı bulunamadı: %v", err)
	}

	var (
		inspector, static, architect, mechanic, electric, controller, mechanicController, electricController *ent.CompanyEngineer
		engineerErr                                                                                          error
	)

	if input.Inspector != nil {
		inspector, engineerErr = client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.Inspector)).
			Only(ctx)
		if engineerErr != nil {
			return nil, fmt.Errorf("denetçi bulunamadı (kod: %d): %v", *input.Inspector, engineerErr)
		}
	}

	if input.Static != nil {
		static, engineerErr = client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.Static)).
			Only(ctx)
		if engineerErr != nil {
			return nil, fmt.Errorf("statik bulunamadı (kod: %d): %v", *input.Static, engineerErr)
		}
	}

	if input.Architect != nil {
		architect, engineerErr = client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.Architect)).
			Only(ctx)
		if engineerErr != nil {
			return nil, fmt.Errorf("mimar bulunamadı (kod: %d): %v", *input.Architect, engineerErr)
		}
	}

	if input.Mechanic != nil {
		mechanic, engineerErr = client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.Mechanic)).
			Only(ctx)
		if engineerErr != nil {
			return nil, fmt.Errorf("mak. müh. bulunamadı (kod: %d): %v", *input.Mechanic, engineerErr)
		}
	}

	if input.Electric != nil {
		electric, engineerErr = client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.Electric)).
			Only(ctx)
		if engineerErr != nil {
			return nil, fmt.Errorf("elektrik bulunamadı (kod: %d): %v", *input.Electric, engineerErr)
		}
	}

	if input.Controller != nil {
		controller, engineerErr = client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.Controller)).
			Only(ctx)
		if engineerErr != nil {
			return nil, fmt.Errorf("kont. elm. bulunamadı (kod: %d): %v", *input.Controller, engineerErr)
		}
	}

	if input.MechanicController != nil {
		mechanicController, engineerErr = client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.MechanicController)).
			Only(ctx)
		if engineerErr != nil {
			return nil, fmt.Errorf("mekanik kont. elm. bulunamadı (kod: %d): %v", *input.MechanicController, engineerErr)
		}
	}

	if input.ElectricController != nil {
		electricController, engineerErr = client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.ElectricController)).
			Only(ctx)
		if engineerErr != nil {
			return nil, fmt.Errorf("elektrik kont. elm. bulunamadı (kod: %d): %v", *input.ElectricController, engineerErr)
		}
	}

	// Mevcut iş detayını güncelle
	jobDetailUpdate := client.JobDetail.UpdateOne(jobDetail).
		SetNillableYibfNo(input.YibfNo).
		SetNillableAdministration(input.Administration).
		SetNillableState(input.State).
		SetNillableSheet(input.Sheet).
		SetNillableFolderNo(input.FolderNo).
		SetNillableContractDate(input.ContractDate).
		SetNillableStartDate(input.StartDate).
		SetNillableCompletionDate(input.CompletionDate).
		SetNillableLicenseDate(input.LicenseDate).
		SetNillableLicenseNo(input.LicenseNo).
		SetNillableConstructionArea(input.ConstructionArea).
		SetNillableLandArea(input.LandArea).
		SetNillableAddress(input.Address).
		SetNillableBuildingClass(input.BuildingClass).
		SetNillableBuildingType(input.BuildingType).
		SetNillableFloorCount(input.FloorCount).
		SetNillableNote(input.Note)

	if inspector != nil {
		jobDetailUpdate = jobDetailUpdate.SetArchitect(inspector)
	}

	if static != nil {
		jobDetailUpdate = jobDetailUpdate.SetStatic(static)
	}

	if architect != nil {
		jobDetailUpdate = jobDetailUpdate.SetArchitect(architect)
	}

	if mechanic != nil {
		jobDetailUpdate = jobDetailUpdate.SetMechanic(mechanic)
	}

	if electric != nil {
		jobDetailUpdate = jobDetailUpdate.SetElectric(electric)
	}

	if controller != nil {
		jobDetailUpdate = jobDetailUpdate.SetController(controller)
	}

	if mechanicController != nil {
		jobDetailUpdate = jobDetailUpdate.SetMechaniccontroller(mechanicController)
	}

	if electricController != nil {
		jobDetailUpdate = jobDetailUpdate.SetElectriccontroller(electricController)
	}

	newJobDetail, err := jobDetailUpdate.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("iş ayrıntısı güncellenemedi: %v", err)
	}

	return newJobDetail, nil
}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, yibfNo int) (*ent.JobDetail, error) {
	client := middlewares.GetClientFromContext(ctx)
	job, err := client.JobDetail.Query().
		Where(jobdetail.YibfNo(yibfNo)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil // Kayıt bulunamadığında nil dönüyoruz, hata değil
		}
		return nil, fmt.Errorf("failed to get job details: %w", err)
	}

	return job, nil
}

// Jobs is the resolver for the jobs field.
func (r *queryResolver) Jobs(ctx context.Context) (*ent.JobDetail, error) {
	client := middlewares.GetClientFromContext(ctx)
	jobs, err := client.JobDetail.Query().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch jobs: %w", err)
	}

	return jobs, nil
}

// JobDetail returns generated.JobDetailResolver implementation.
func (r *Resolver) JobDetail() generated.JobDetailResolver { return &jobDetailResolver{r} }

type jobDetailResolver struct{ *Resolver }
