package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"
	"fmt"

	"github.com/polatbilal/gqlgen-ent/ent"
	"github.com/polatbilal/gqlgen-ent/ent/companydetail"
	"github.com/polatbilal/gqlgen-ent/ent/companyuser"
	"github.com/polatbilal/gqlgen-ent/ent/jobdetail"
	"github.com/polatbilal/gqlgen-ent/ent/jobrelations"
	"github.com/polatbilal/gqlgen-ent/ent/user"
	"github.com/polatbilal/gqlgen-ent/graphql/generated"
	"github.com/polatbilal/gqlgen-ent/graphql/helpers"
	"github.com/polatbilal/gqlgen-ent/graphql/model"
	"github.com/polatbilal/gqlgen-ent/middlewares"
)

// CompanyCode is the resolver for the CompanyCode field.
func (r *jobDetailResolver) CompanyCode(ctx context.Context, obj *ent.JobDetail) (int, error) {
	client := middlewares.GetClientFromContext(ctx)
	// Önce JobRelations'ı bul
	relations, err := client.JobRelations.Query().Where(jobrelations.HasJobWith(jobdetail.YibfNoEQ(obj.YibfNo))).Only(ctx)
	if err != nil {
		return 0, fmt.Errorf("ilişkiler bulunamadı: %v", err)
	}

	// Sonra Company'i bul
	company, err := relations.QueryCompany().Only(ctx)
	if err != nil {
		return 0, fmt.Errorf("şirket bulunamadı: %v", err)
	}

	return int(company.CompanyCode), nil
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
		Where(companydetail.CompanyCodeEQ(*input.CompanyCode)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("şirket bulunamadı (kod: %d): %v", input.CompanyCode, err)
	}

	// İş detayını oluştur
	jobDetail, err := client.JobDetail.Create().
		SetYibfNo(*input.YibfNo).
		SetTitle(*input.Title).
		SetNillableAdministration(input.Administration).
		SetNillableState(input.State).
		SetNillableIsland(input.Island).
		SetNillableParcel(input.Parcel).
		SetNillableSheet(input.Sheet).
		SetNillableDistributionDate(input.DistributionDate).
		SetNillableContractDate(input.ContractDate).
		SetNillableStartDate(input.StartDate).
		SetNillableLicenseDate(input.LicenseDate).
		SetNillableLicenseNo(input.LicenseNo).
		SetNillableCompletionDate(input.CompletionDate).
		SetNillableLandArea(input.LandArea).
		SetNillableTotalArea(input.TotalArea).
		SetNillableConstructionArea(input.ConstructionArea).
		SetNillableLeftArea(input.LeftArea).
		SetNillableYDSAddress(input.YDSAddress).
		SetNillableAddress(input.Address).
		SetNillableBuildingClass(input.BuildingClass).
		SetNillableBuildingType(input.BuildingType).
		SetNillableLevel(input.Level).
		SetNillableUnitPrice(input.UnitPrice).
		SetNillableFloorCount(input.FloorCount).
		SetNillableBKSReferenceNo(input.BKSReferenceNo).
		SetNillableCoordinates(input.Coordinates).
		SetNillableFolderNo(input.FolderNo).
		SetNillableUploadedFile(input.UploadedFile).
		SetNillableIndustryArea(input.IndustryArea).
		SetNillableClusterStructure(input.ClusterStructure).
		SetNillableIsLicenseExpired(input.IsLicenseExpired).
		SetNillableIsCompleted(input.IsCompleted).
		SetNillableNote(input.Note).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("iş detayı oluşturulamadı: %v", err)
	}

	// İlişkileri oluştur
	jobRelations := client.JobRelations.Create().
		SetYibfNo(*input.YibfNo).
		SetJob(jobDetail).
		SetCompany(company)

	_, err = jobRelations.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("iş ilişkileri oluşturulamadı: %v", err)
	}

	// Progress bilgilerini iş ayrıntısına ekle
	var one, two, three, four, five, six float64

	if input.Level != nil {
		level := float64(*input.Level)
		remaining := level

		// Sabit değerlere göre sıralı dağıtım:
		// one: 10, two: 10, three: 40, four: 20, five: 15, six: 5

		// İlk dört seviyeyi sırayla doldur
		if remaining >= 10 {
			one = 10
			remaining -= 10
		} else {
			one = remaining
			remaining = 0
		}

		if remaining >= 10 {
			two = 10
			remaining -= 10
		} else {
			two = remaining
			remaining = 0
		}

		if remaining >= 40 {
			three = 40
			remaining -= 40
		} else {
			three = remaining
			remaining = 0
		}

		if remaining >= 20 {
			four = 20
			remaining -= 20
		} else {
			four = remaining
			remaining = 0
		}

		if remaining >= 15 {
			five = 15
			remaining -= 15
		} else {
			five = remaining
			remaining = 0
		}

		if remaining > 0 {
			if remaining <= 5 {
				six = remaining
			} else {
				six = 5
			}
		}
	}

	p, err := client.JobProgress.Create().
		SetYibfNo(*input.YibfNo).
		SetOne(int(one)).
		SetTwo(int(two)).
		SetThree(int(three)).
		SetFour(int(four)).
		SetFive(int(five)).
		SetSix(int(six)).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("progress oluşturulamadı: %v, %s", err, p)
	}

	// Progress'i JobRelations ile ilişkilendir
	relations, err := jobDetail.QueryRelations().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("ilişkiler bulunamadı: %v", err)
	}

	_, err = relations.Update().SetProgress(p).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("progress ilişkilendirilemedi: %v", err)
	}

	return jobDetail, nil
}

// UpdateJob is the resolver for the updateJob field.
func (r *mutationResolver) UpdateJob(ctx context.Context, yibfNo int, input model.JobInput) (*ent.JobDetail, error) {
	client := middlewares.GetClientFromContext(ctx)

	// İş detayını bul
	jobDetail, err := client.JobDetail.Query().Where(jobdetail.YibfNoEQ(yibfNo)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("iş ayrıntısı bulunamadı: %v", err)
	}

	// İş detayını güncelle
	jobDetail, err = client.JobDetail.UpdateOne(jobDetail).
		SetNillableYibfNo(input.YibfNo).
		SetNillableTitle(input.Title).
		SetNillableAdministration(input.Administration).
		SetNillableState(input.State).
		SetNillableIsland(input.Island).
		SetNillableParcel(input.Parcel).
		SetNillableSheet(input.Sheet).
		SetNillableDistributionDate(input.DistributionDate).
		SetNillableContractDate(input.ContractDate).
		SetNillableStartDate(input.StartDate).
		SetNillableLicenseDate(input.LicenseDate).
		SetNillableLicenseNo(input.LicenseNo).
		SetNillableCompletionDate(input.CompletionDate).
		SetNillableLandArea(input.LandArea).
		SetNillableTotalArea(input.TotalArea).
		SetNillableConstructionArea(input.ConstructionArea).
		SetNillableLeftArea(input.LeftArea).
		SetNillableYDSAddress(input.YDSAddress).
		SetNillableAddress(input.Address).
		SetNillableBuildingClass(input.BuildingClass).
		SetNillableBuildingType(input.BuildingType).
		SetNillableLevel(input.Level).
		SetNillableUnitPrice(input.UnitPrice).
		SetNillableFloorCount(input.FloorCount).
		SetNillableBKSReferenceNo(input.BKSReferenceNo).
		SetNillableCoordinates(input.Coordinates).
		SetNillableFolderNo(input.FolderNo).
		SetNillableUploadedFile(input.UploadedFile).
		SetNillableIndustryArea(input.IndustryArea).
		SetNillableClusterStructure(input.ClusterStructure).
		SetNillableIsLicenseExpired(input.IsLicenseExpired).
		SetNillableIsCompleted(input.IsCompleted).
		SetNillableNote(input.Note).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("iş detayı güncellenemedi: %v", err)
	}

	// Level değişmişse progress'i güncelle
	if input.Level != nil {
		// İlişkili progress'i bul
		relations, err := jobDetail.QueryRelations().WithProgress().Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("ilişkiler bulunamadı: %v", err)
		}

		// Mevcut progress değerlerinin toplamını kontrol et
		progress := relations.Edges.Progress
		if progress != nil {
			currentTotal := float64(progress.One + progress.Two + progress.Three + progress.Four + progress.Five + progress.Six)

			// Progress toplamı level'dan büyükse güncelleme yapma
			if currentTotal > float64(*input.Level) {
				fmt.Printf("Progress toplamı (%.0f) level değerinden (%.0f) büyük olduğu için güncelleme yapılmayacak\n",
					currentTotal, float64(*input.Level))
				return jobDetail, nil
			}

			// Progress toplamı level'dan küçükse güncelle
			var one, two, three, four, five, six float64
			level := float64(*input.Level)
			remaining := level

			// Sabit değerlere göre sıralı dağıtım
			if remaining >= 10 {
				one = 10
				remaining -= 10
			} else {
				one = remaining
				remaining = 0
			}

			if remaining >= 10 {
				two = 10
				remaining -= 10
			} else {
				two = remaining
				remaining = 0
			}

			if remaining >= 40 {
				three = 40
				remaining -= 40
			} else {
				three = remaining
				remaining = 0
			}

			if remaining >= 20 {
				four = 20
				remaining -= 20
			} else {
				four = remaining
				remaining = 0
			}

			if remaining >= 15 {
				five = 15
				remaining -= 15
			} else {
				five = remaining
				remaining = 0
			}

			if remaining > 0 {
				if remaining <= 5 {
					six = remaining
				} else {
					six = 5
				}
			}

			// Progress'i güncelle
			_, err = progress.Update().
				SetOne(int(one)).
				SetTwo(int(two)).
				SetThree(int(three)).
				SetFour(int(four)).
				SetFive(int(five)).
				SetSix(int(six)).
				Save(ctx)

			if err != nil {
				return nil, fmt.Errorf("progress güncellenemedi: %v", err)
			}

			fmt.Println("Progress başarıyla güncellendi!")
		} else {
			fmt.Println("Progress bulunamadı, yeni oluşturuluyor...")
			// Eğer progress yoksa yeni oluştur
			var one, two, three, four, five, six float64
			level := float64(*input.Level)
			remaining := level

			// Sabit değerlere göre sıralı dağıtım
			if remaining >= 10 {
				one = 10
				remaining -= 10
			} else {
				one = remaining
				remaining = 0
			}

			if remaining >= 10 {
				two = 10
				remaining -= 10
			} else {
				two = remaining
				remaining = 0
			}

			if remaining >= 40 {
				three = 40
				remaining -= 40
			} else {
				three = remaining
				remaining = 0
			}

			if remaining >= 20 {
				four = 20
				remaining -= 20
			} else {
				four = remaining
				remaining = 0
			}

			if remaining >= 15 {
				five = 15
				remaining -= 15
			} else {
				five = remaining
				remaining = 0
			}

			if remaining > 0 {
				if remaining <= 5 {
					six = remaining
				} else {
					six = 5
				}
			}

			p, err := client.JobProgress.Create().
				SetYibfNo(yibfNo).
				SetOne(int(one)).
				SetTwo(int(two)).
				SetThree(int(three)).
				SetFour(int(four)).
				SetFive(int(five)).
				SetSix(int(six)).
				Save(ctx)

			if err != nil {
				return nil, fmt.Errorf("progress oluşturulamadı: %v", err)
			}

			// Progress'i JobRelations ile ilişkilendir
			_, err = relations.Update().SetProgress(p).Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("progress ilişkilendirilemedi: %v", err)
			}
		}
	}

	return jobDetail, nil
}

// UpdateJobEngineer is the resolver for the updateJobEngineer field.
func (r *mutationResolver) UpdateJobEngineer(ctx context.Context, input model.JobEngineerInput) (*model.JobEngineer, error) {
	client := middlewares.GetClientFromContext(ctx)

	jobDetail, err := client.JobDetail.Query().Where(jobdetail.YibfNoEQ(*input.YibfNo)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("iş detayı bulunamadı: %w", err)
	}

	jobRelations, err := jobDetail.QueryRelations().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("iş ilişkileri bulunamadı: %w", err)
	}

	// ValidateAndGetEngineers fonksiyonunu kullanarak mühendisleri doğrula ve getir
	engineers, err := helpers.ValidateAndGetEngineers(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("mühendisler doğrulanırken hata: %w", err)
	}

	// Mühendis ilişkilerini güncelle
	update := client.JobRelations.UpdateOne(jobRelations)

	if inspector, ok := engineers["inspector"]; ok {
		update.SetInspector(inspector)
	}
	if static, ok := engineers["static"]; ok {
		update.SetStatic(static)
	}
	if architect, ok := engineers["architect"]; ok {
		update.SetArchitect(architect)
	}
	if mechanic, ok := engineers["mechanic"]; ok {
		update.SetMechanic(mechanic)
	}
	if electric, ok := engineers["electric"]; ok {
		update.SetElectric(electric)
	}
	if controller, ok := engineers["controller"]; ok {
		update.SetController(controller)
	}
	if mechanicController, ok := engineers["mechanicController"]; ok {
		update.SetMechaniccontroller(mechanicController)
	}
	if electricController, ok := engineers["electricController"]; ok {
		update.SetElectriccontroller(electricController)
	}

	if err := update.Exec(ctx); err != nil {
		return nil, fmt.Errorf("mühendis ilişkileri güncellenirken hata: %w", err)
	}

	// Güncel ilişkileri getir
	updatedRelations, err := client.JobRelations.Query().
		Where(jobrelations.YibfNo(*input.YibfNo)).
		WithInspector().
		WithStatic().
		WithArchitect().
		WithMechanic().
		WithElectric().
		WithController().
		WithMechaniccontroller().
		WithElectriccontroller().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("güncel ilişkiler alınırken hata: %w", err)
	}

	return &model.JobEngineer{
		YibfNo:             input.YibfNo,
		Inspector:          updatedRelations.Edges.Inspector,
		Static:             updatedRelations.Edges.Static,
		Architect:          updatedRelations.Edges.Architect,
		Mechanic:           updatedRelations.Edges.Mechanic,
		Electric:           updatedRelations.Edges.Electric,
		Controller:         updatedRelations.Edges.Controller,
		MechanicController: updatedRelations.Edges.Mechaniccontroller,
		ElectricController: updatedRelations.Edges.Electriccontroller,
	}, nil
}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, yibfNo int) (*ent.JobDetail, error) {
	client := middlewares.GetClientFromContext(ctx)
	customClaim := middlewares.CtxValue(ctx)

	// Kullanıcının şirketlerini al
	userCompanies, err := client.CompanyUser.Query().
		Where(companyuser.HasUserWith(user.IDEQ(customClaim.ID))).
		QueryCompany().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("kullanıcı şirketleri alınamadı: %v", err)
	}

	// Kullanıcının yetkili olduğu şirket kodlarını topla
	var companyCodes []int
	for _, company := range userCompanies {
		companyCodes = append(companyCodes, company.CompanyCode)
	}

	// İşi sorgula
	job, err := client.JobDetail.Query().
		Where(jobdetail.YibfNoEQ(yibfNo)).
		WithRelations(func(q *ent.JobRelationsQuery) {
			q.Where(
				jobrelations.HasCompanyWith(
					companydetail.CompanyCodeIn(companyCodes...),
				),
			)
		}).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("iş bulunamadı veya bu işe erişim yetkiniz yok")
		}
		return nil, fmt.Errorf("iş sorgulanırken hata oluştu: %v", err)
	}

	return job, nil
}

// Jobs is the resolver for the jobs field.
func (r *queryResolver) Jobs(ctx context.Context) ([]*ent.JobDetail, error) {
	client := middlewares.GetClientFromContext(ctx)
	customClaim := middlewares.CtxValue(ctx)

	// Kullanıcının şirketlerini al
	userCompanies, err := client.CompanyUser.Query().
		Where(companyuser.HasUserWith(user.IDEQ(customClaim.ID))).
		QueryCompany().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("kullanıcı şirketleri alınamadı: %v", err)
	}

	// Kullanıcının yetkili olduğu şirket kodlarını topla
	var companyCodes []int
	for _, company := range userCompanies {
		companyCodes = append(companyCodes, company.CompanyCode)
	}

	// İşleri sorgula
	jobs, err := client.JobDetail.Query().
		Where(jobdetail.StateNEQ("Bitmiş")).
		WithRelations(func(q *ent.JobRelationsQuery) {
			q.Where(
				jobrelations.HasCompanyWith(
					companydetail.CompanyCodeIn(companyCodes...),
				),
			)
		}).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("işler sorgulanırken hata oluştu: %v", err)
	}

	return jobs, nil
}

// JobCounts is the resolver for the jobCounts field.
func (r *queryResolver) JobCounts(ctx context.Context, companyCode *int) (*model.JobCounts, error) {
	client := middlewares.GetClientFromContext(ctx)
	customClaim := middlewares.CtxValue(ctx)

	// Kullanıcının şirketlerini al
	userCompanies, err := client.CompanyUser.Query().
		Where(companyuser.HasUserWith(user.IDEQ(customClaim.ID))).
		QueryCompany().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("kullanıcı şirketleri alınamadı: %v", err)
	}

	// Şirket kodlarını ve erişim kontrolünü yap
	var companyCodes []int
	if companyCode != nil {
		hasAccess := false
		for _, company := range userCompanies {
			if company.CompanyCode == *companyCode {
				hasAccess = true
				companyCodes = []int{*companyCode}
				break
			}
		}
		if !hasAccess {
			return nil, fmt.Errorf("bu şirkete erişim yetkiniz yok")
		}
	} else {
		for _, company := range userCompanies {
			companyCodes = append(companyCodes, company.CompanyCode)
		}
	}

	// İşleri sorgula
	jobs, err := client.JobDetail.Query().
		WithRelations(func(q *ent.JobRelationsQuery) {
			q.Where(
				jobrelations.HasCompanyWith(
					companydetail.CompanyCodeIn(companyCodes...),
				),
			)
		}).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("işler sorgulanırken hata oluştu: %v", err)
	}

	// Durumları say
	var current, pending, completed int
	for _, job := range jobs {
		switch job.State {
		case "Güncel":
			current++
		case "Ruhsat Bekleyen":
			pending++
		case "Bitmiş":
			completed++
		}
	}

	return &model.JobCounts{
		Current:   current,
		Pending:   pending,
		Completed: completed,
		Total:     len(jobs),
	}, nil
}

// JobEngineer is the resolver for the jobEngineer field.
func (r *queryResolver) JobEngineer(ctx context.Context, yibfNo int) (*model.JobEngineer, error) {
	client := middlewares.GetClientFromContext(ctx)

	jobRelations, err := client.JobRelations.Query().
		Where(jobrelations.YibfNoEQ(yibfNo)).
		WithInspector().
		WithStatic().
		WithArchitect().
		WithMechanic().
		WithElectric().
		WithController().
		WithMechaniccontroller().
		WithElectriccontroller().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("iş mühendisi bulunamadı: %v", err)
	}

	return &model.JobEngineer{
		YibfNo:             &yibfNo,
		Inspector:          jobRelations.Edges.Inspector,
		Static:             jobRelations.Edges.Static,
		Architect:          jobRelations.Edges.Architect,
		Mechanic:           jobRelations.Edges.Mechanic,
		Electric:           jobRelations.Edges.Electric,
		Controller:         jobRelations.Edges.Controller,
		MechanicController: jobRelations.Edges.Mechaniccontroller,
		ElectricController: jobRelations.Edges.Electriccontroller,
	}, nil
}

// JobDetail returns generated.JobDetailResolver implementation.
func (r *Resolver) JobDetail() generated.JobDetailResolver { return &jobDetailResolver{r} }

type jobDetailResolver struct{ *Resolver }
