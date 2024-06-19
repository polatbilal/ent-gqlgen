package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"
	"gqlgen-ent/ent"
	"gqlgen-ent/ent/jobauthor"
	"gqlgen-ent/ent/jobcontractor"
	"gqlgen-ent/ent/jobdetail"
	"gqlgen-ent/ent/joblayer"
	"gqlgen-ent/ent/jobowner"
	"gqlgen-ent/ent/jobprogress"
	"gqlgen-ent/graph/generated"
	"gqlgen-ent/graph/model"
	"gqlgen-ent/middlewares"
	"time"
)

// ContractDate is the resolver for the ContractDate field.
func (r *jobDetailResolver) ContractDate(ctx context.Context, obj *ent.JobDetail) (*string, error) {
	contractDate := obj.ContractDate.Format("2006-01-02")
	return &contractDate, nil
}

// StartDate is the resolver for the StartDate field.
func (r *jobDetailResolver) StartDate(ctx context.Context, obj *ent.JobDetail) (*string, error) {
	startDate := obj.StartDate.Format("2006-01-02")
	return &startDate, nil
}

// LicenseDate is the resolver for the LicenseDate field.
func (r *jobDetailResolver) LicenseDate(ctx context.Context, obj *ent.JobDetail) (*string, error) {
	licenseDate := obj.LicenseDate.Format("2006-01-02")
	return &licenseDate, nil
}

// Layer is the resolver for the Layer field.
func (r *jobDetailResolver) Layer(ctx context.Context, obj *ent.JobDetail) ([]*ent.JobLayer, error) {
	client := middlewares.GetClientFromContext(ctx)
	return client.JobLayer.Query().Where(joblayer.HasLayerWith(jobdetail.IDEQ(obj.ID))).All(ctx)
}

// CreateJob is the resolver for the createJob field.
func (r *mutationResolver) CreateJob(ctx context.Context, input model.JobInput) (*ent.JobDetail, error) {
	client := middlewares.GetClientFromContext(ctx)

	_, err := client.JobDetail.Query().Where(jobdetail.YibfNoEQ(*input.YibfNo)).Only(ctx)
	if err == nil {
		return nil, fmt.Errorf("kullanıcı adı zaten mevcut")
	}

	if !ent.IsNotFound(err) {
		return nil, err
	}

	var contractDatePtr *time.Time
	if input.ContractDate != nil {
		parsedDate, err := time.Parse("2006-01-02", *input.ContractDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse contract date: %v", err)
		}
		contractDatePtr = &parsedDate
	}

	var startDatePtr *time.Time
	if input.StartDate != nil {
		parsedStartDate, err := time.Parse("2006-01-02", *input.StartDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse start date: %v", err)
		}
		startDatePtr = &parsedStartDate
	}

	var licenseDatePtr *time.Time
	if input.LicenseDate != nil {
		parsedLicenseDate, err := time.Parse("2006-01-02", *input.LicenseDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse license date: %v", err)
		}
		licenseDatePtr = &parsedLicenseDate
	}

	newJobDetail, err := client.JobDetail.Create().
		SetYibfNo(*input.YibfNo).
		SetNillableProvince(input.Province).
		SetNillableIdare(input.Idare).
		SetNillablePafta(input.Pafta).
		SetNillableAda(input.Ada).
		SetNillableParsel(input.Parsel).
		SetNillableFolderNo(input.FolderNo).
		SetNillableStatus(input.Status).
		SetNillableContractDate(contractDatePtr).
		SetNillableStartDate(startDatePtr).
		SetNillableLicenseDate(licenseDatePtr).
		SetNillableLicenseNo(input.LicenseNo).
		SetNillableConstructionArea(input.ConstructionArea).
		SetNillableLandArea(input.LandArea).
		SetNillableDistrict(input.District).
		SetNillableVillage(input.Village).
		SetNillableStreet(input.Street).
		SetNillableBuildingClass(input.BuildingClass).
		SetNillableBuildingType(input.BuildingType).
		SetNillableBuildingBlock(input.BuildingBlock).
		SetNillableFloors(input.Floors).
		SetNillableNote(input.Note).
		SetNillableStarted(input.Started).
		SetNillableUsagePurpose(input.UsagePurpose).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("iş ayrıntısı oluşturulamadı: %v", err)
	}

	// Owner bilgilerini iş ayrıntısına ekle
	for _, ownerInput := range input.Owner {
		_, err := client.JobOwner.Create().
			SetNillableName(ownerInput.Name).
			SetNillableTcNo(ownerInput.TcNo).
			SetNillableTaxAdmin(ownerInput.TaxAdmin).
			SetNillableTaxNo(ownerInput.TaxNo).
			SetNillablePhone(ownerInput.Phone).
			SetNillableEmail(ownerInput.Email).
			SetNillableNote(ownerInput.Note).
			AddOwners(newJobDetail).
			Save(ctx)

		if err != nil {
			return nil, fmt.Errorf("yapı sahibi oluşturulamadı: %v", err)
		}
	}

	// Contractor bilgilerini iş ayrıntısına ekle
	for _, contractorInput := range input.Contractor {
		_, err := client.JobContractor.Create().
			SetNillableName(contractorInput.Name).
			SetNillableTcNo(contractorInput.TcNo).
			SetNillableTaxAdmin(contractorInput.TaxAdmin).
			SetNillableTaxNo(contractorInput.TaxNo).
			SetNillablePhone(contractorInput.Phone).
			SetNillableEmail(contractorInput.Email).
			SetNillableNote(contractorInput.Note).
			AddContractors(newJobDetail).
			Save(ctx)

		if err != nil {
			return nil, fmt.Errorf("müteahhit oluşturulamadı: %v", err)
		}
	}

	// Author bilgilerini iş ayrıntısına ekle
	for _, authorInput := range input.Author {
		_, err := client.JobAuthor.Create().
			SetNillableArchitect(authorInput.Architect).
			SetNillableStatic(authorInput.Static).
			SetNillableMechanic(authorInput.Mechanic).
			SetNillableElectric(authorInput.Electric).
			SetNillableFloor(authorInput.Floor).
			AddAuthors(newJobDetail).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("yazar oluşturulamadı: %v", err)
		}
		// }
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

	var contractDatePtr *time.Time
	if input.ContractDate != nil {
		parsedDate, err := time.Parse("2006-01-02", *input.ContractDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse contract date: %v", err)
		}
		contractDatePtr = &parsedDate
	}

	var startDatePtr *time.Time
	if input.StartDate != nil {
		parsedStartDate, err := time.Parse("2006-01-02", *input.StartDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse start date: %v", err)
		}
		startDatePtr = &parsedStartDate
	}

	var licenseDatePtr *time.Time
	if input.LicenseDate != nil {
		parsedLicenseDate, err := time.Parse("2006-01-02", *input.LicenseDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse license date: %v", err)
		}
		licenseDatePtr = &parsedLicenseDate
	}

	// Mevcut iş detayını güncelle
	jobDetail, err = client.JobDetail.UpdateOne(jobDetail).
		SetNillableYibfNo(input.YibfNo).
		SetNillableProvince(input.Province).
		SetNillableIdare(input.Idare).
		SetNillablePafta(input.Pafta).
		SetNillableAda(input.Ada).
		SetNillableParsel(input.Parsel).
		SetNillableFolderNo(input.FolderNo).
		SetNillableStatus(input.Status).
		SetNillableContractDate(contractDatePtr).
		SetNillableStartDate(startDatePtr).
		SetNillableLicenseDate(licenseDatePtr).
		SetNillableLicenseNo(input.LicenseNo).
		SetNillableConstructionArea(input.ConstructionArea).
		SetNillableLandArea(input.LandArea).
		SetNillableDistrict(input.District).
		SetNillableVillage(input.Village).
		SetNillableStreet(input.Street).
		SetNillableBuildingClass(input.BuildingClass).
		SetNillableBuildingType(input.BuildingType).
		SetNillableBuildingBlock(input.BuildingBlock).
		SetNillableFloors(input.Floors).
		SetNillableNote(input.Note).
		SetNillableStarted(input.Started).
		SetNillableUsagePurpose(input.UsagePurpose).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("iş ayrıntısı güncellenemedi: %v", err)
	}

	// Owner bilgilerini güncelle
	for _, ownerInput := range input.Owner {
		existingOwner, err := client.JobOwner.Query().
			Where(jobowner.HasOwnersWith(jobdetail.IDEQ(jobDetail.ID))). // İlk sahibi kullanarak sorgulama
			Only(ctx)

		if err != nil {
			if ent.IsNotFound(err) {
				_, err := client.JobOwner.Create().
					SetNillableName(ownerInput.Name).
					SetNillableTcNo(ownerInput.TcNo).
					SetNillableTaxAdmin(ownerInput.TaxAdmin).
					SetNillableTaxNo(ownerInput.TaxNo).
					SetNillablePhone(ownerInput.Phone).
					SetNillableEmail(ownerInput.Email).
					SetNillableNote(ownerInput.Note).
					AddOwners(jobDetail). // İş detayına ekle
					Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("yapı sahibi oluşturulamadı: %v", err)
				}
				continue
			}
			return nil, fmt.Errorf("mevcut yapı sahibi bulunamadı: %v", err)
		}

		_, err = client.JobOwner.
			UpdateOne(existingOwner).
			SetNillableName(ownerInput.Name).
			SetNillableTcNo(ownerInput.TcNo).
			SetNillableTaxAdmin(ownerInput.TaxAdmin).
			SetNillableTaxNo(ownerInput.TaxNo).
			SetNillablePhone(ownerInput.Phone).
			SetNillableEmail(ownerInput.Email).
			SetNillableNote(ownerInput.Note).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("yapı sahibi güncellenemedi: %v", err)
		}
	}

	// Contractor bilgilerini güncelle
	for _, contractorInput := range input.Contractor {
		existingContractor, err := client.JobContractor.Query().
			Where(jobcontractor.HasContractorsWith(jobdetail.IDEQ(jobDetail.ID))).
			Only(ctx)

		if err != nil {
			if ent.IsNotFound(err) {
				_, err := client.JobContractor.Create().
					SetNillableName(contractorInput.Name).
					SetNillableTcNo(contractorInput.TcNo).
					SetNillableTaxAdmin(contractorInput.TaxAdmin).
					SetNillableTaxNo(contractorInput.TaxNo).
					SetNillablePhone(contractorInput.Phone).
					SetNillableEmail(contractorInput.Email).
					SetNillableNote(contractorInput.Note).
					AddContractors(jobDetail). // İş detayına ekle
					Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("müteahhit oluşturulamadı: %v", err)
				}
				continue
			}
			return nil, fmt.Errorf("mevcut müteahhit bulunamadı: %v", err)
		}

		_, err = client.JobContractor.
			UpdateOne(existingContractor).
			SetNillableName(contractorInput.Name).
			SetNillableTcNo(contractorInput.TcNo).
			SetNillableTaxAdmin(contractorInput.TaxAdmin).
			SetNillableTaxNo(contractorInput.TaxNo).
			SetNillablePhone(contractorInput.Phone).
			SetNillableEmail(contractorInput.Email).
			SetNillableNote(contractorInput.Note).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("müteahhit güncellenemedi: %v", err)
		}
	}

	// Author bilgilerini güncelle
	for _, authorInput := range input.Author {
		existingAuthor, err := client.JobAuthor.Query().
			Where(jobauthor.HasAuthorsWith(jobdetail.IDEQ(jobDetail.ID))).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				// Mevcut yazar bulunamadı, yeni yazar oluştur
				_, err := client.JobAuthor.Create().
					SetNillableArchitect(authorInput.Architect).
					SetNillableStatic(authorInput.Static).
					SetNillableMechanic(authorInput.Mechanic).
					SetNillableElectric(authorInput.Electric).
					SetNillableFloor(authorInput.Floor).
					AddAuthors(jobDetail). // İş detayına ekle
					Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("yeni yazar oluşturulamadı: %v", err)
				}
				continue
			}
			return nil, fmt.Errorf("mevcut yazar bulunamadı: %v", err)
		}

		// Mevcut yazarı güncelle
		_, err = client.JobAuthor.
			UpdateOne(existingAuthor).
			SetNillableArchitect(authorInput.Architect).
			SetNillableStatic(authorInput.Static).
			SetNillableMechanic(authorInput.Mechanic).
			SetNillableElectric(authorInput.Electric).
			SetNillableFloor(authorInput.Floor).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("yazar güncellenemedi: %v", err)
		}
	}

	// Progress bilgilerini güncelle
	for _, progressInput := range input.Progress {
		existingProgress, err := client.JobProgress.Query().
			Where(jobprogress.HasProgressWith(jobdetail.IDEQ(jobDetail.ID))).
			Only(ctx)

		if err != nil {
			if ent.IsNotFound(err) {
				_, err := client.JobProgress.Create().
					SetNillableOne(progressInput.One).
					SetNillableTwo(progressInput.Two).
					SetNillableThree(progressInput.Three).
					SetNillableFour(progressInput.Four).
					SetNillableFive(progressInput.Five).
					SetNillableSix(progressInput.Six).
					AddProgress(jobDetail). // İş detayına ekle
					Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("ilerleme bilgisi oluşturulamadı: %v", err)
				}
				continue
			}
			return nil, fmt.Errorf("mevcut ilerleme bilgisi bulunamadı: %v", err)
		}

		// Mevcut ilerleme bilgisini güncelle
		_, err = client.JobProgress.
			UpdateOne(existingProgress).
			SetNillableOne(progressInput.One).
			SetNillableTwo(progressInput.Two).
			SetNillableThree(progressInput.Three).
			SetNillableFour(progressInput.Four).
			SetNillableFive(progressInput.Five).
			SetNillableSix(progressInput.Six).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("ilerleme bilgisi güncellenemedi: %v", err)
		}
	}

	return jobDetail, nil
}

// DeleteJob is the resolver for the deleteJob field.
func (r *mutationResolver) DeleteJob(ctx context.Context, yibfNo int) (bool, error) {
	client := middlewares.GetClientFromContext(ctx)

	// İş ayrıntısını yibfNo ile sorgula
	jobDetail, err := client.JobDetail.Query().Where(jobdetail.YibfNoEQ(yibfNo)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, fmt.Errorf("kayıt bulunamadı")
		}
		return false, err
	}

	// İş ayrıntısının deleted sütununu güncelle
	err = client.JobDetail.UpdateOne(jobDetail).SetDeleted(1).Exec(ctx)
	if err != nil {
		return false, fmt.Errorf("iş ayrıntısı silinemedi: %v", err)
	}

	// İlişkili job owner, contractor, author ve progress kayıtlarını güncelle
	_, err = client.JobOwner.Update().Where(jobowner.HasOwnersWith(jobdetail.IDEQ(jobDetail.ID))).SetDeleted(1).Save(ctx)
	if err != nil {
		return false, fmt.Errorf("yapı sahipleri silinemedi: %v", err)
	}

	_, err = client.JobContractor.Update().Where(jobcontractor.HasContractorsWith(jobdetail.IDEQ(jobDetail.ID))).SetDeleted(1).Save(ctx)
	if err != nil {
		return false, fmt.Errorf("müteahhitler silinemedi: %v", err)
	}

	_, err = client.JobAuthor.Update().Where(jobauthor.HasAuthorsWith(jobdetail.IDEQ(jobDetail.ID))).SetDeleted(1).Save(ctx)
	if err != nil {
		return false, fmt.Errorf("yazarlar silinemedi: %v", err)
	}

	return true, nil
}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, yibfNo *int) ([]*ent.JobDetail, error) {
	client := middlewares.GetClientFromContext(ctx)
	query := client.JobDetail.Query()

	// deleted sütunu 0 olanları filtrele
	query = query.Where(jobdetail.DeletedEQ(0))

	if yibfNo != nil {
		query = query.Where(jobdetail.YibfNoEQ(*yibfNo))
	}

	jobDetails, err := query.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get job details: %v", err)
	}

	return jobDetails, nil
}

// JobDetail returns generated.JobDetailResolver implementation.
func (r *Resolver) JobDetail() generated.JobDetailResolver { return &jobDetailResolver{r} }

type jobDetailResolver struct{ *Resolver }
