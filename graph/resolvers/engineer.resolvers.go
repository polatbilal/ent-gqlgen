package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/polatbilal/gqlgen-ent/ent"
	"github.com/polatbilal/gqlgen-ent/ent/companydetail"
	"github.com/polatbilal/gqlgen-ent/ent/companyengineer"
	"github.com/polatbilal/gqlgen-ent/ent/companyuser"
	"github.com/polatbilal/gqlgen-ent/ent/user"
	"github.com/polatbilal/gqlgen-ent/graph/generated"
	"github.com/polatbilal/gqlgen-ent/graph/model"
	"github.com/polatbilal/gqlgen-ent/middlewares"
)

// CompanyCode is the resolver for the CompanyCode field.
func (r *companyEngineerResolver) CompanyCode(ctx context.Context, obj *ent.CompanyEngineer) (*int, error) {
	client := middlewares.GetClientFromContext(ctx)
	company, err := client.CompanyDetail.Query().
		Where(companydetail.HasEngineersWith(companyengineer.IDEQ(obj.ID))).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("şirket bulunamadı: %v", err)
	}
	return &company.CompanyCode, nil
}

// CreateEngineer is the resolver for the createEngineer field.
func (r *mutationResolver) CreateEngineer(ctx context.Context, input model.CompanyEngineerInput) (*ent.CompanyEngineer, error) {
	client := middlewares.GetClientFromContext(ctx)

	// YDSID ile denetçi kontrolü
	if input.Ydsid != nil && input.CompanyCode != nil {
		// Önce denetçiyi bul
		engineer, err := client.CompanyEngineer.Query().
			Where(companyengineer.YDSIDEQ(*input.Ydsid)).
			WithCompany().
			Only(ctx)

		// Eğer denetçi bulunduysa ve aynı şirkette ise hata döndür
		if err == nil && engineer.Edges.Company != nil && engineer.Edges.Company.CompanyCode == *input.CompanyCode {
			return nil, fmt.Errorf("bu ydsid numarasına (%d) sahip denetçi bu şirkette zaten mevcut", *input.Ydsid)
		}
	}

	// CompanyCode ile şirketi bul
	company, err := client.CompanyDetail.Query().
		Where(companydetail.CompanyCodeEQ(*input.CompanyCode)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("şirket bulunamadı (kod: %d): %v", *input.CompanyCode, err)
	}

	createEngineer, err := client.CompanyEngineer.Create().
		SetNillableName(input.Name).
		SetNillableAddress(input.Address).
		SetNillableEmail(input.Email).
		SetNillableTcNo(input.TcNo).
		SetNillablePhone(input.Phone).
		SetNillableRegisterNo(input.RegisterNo).
		SetNillableCertNo(input.CertNo).
		SetNillableCareer(input.Career).
		SetNillablePosition(input.Position).
		SetNillableYDSID(input.Ydsid).
		SetNillableNote(input.Note).
		SetNillableEmployment(input.Employment).
		SetCompany(company).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("mühendis oluşturulamadı: %v, %s", err, createEngineer)
	}

	return createEngineer, nil
}

// UpdateEngineer is the resolver for the updateEngineer field.
func (r *mutationResolver) UpdateEngineer(ctx context.Context, ydsid int, input model.CompanyEngineerInput) (*ent.CompanyEngineer, error) {
	client := middlewares.GetClientFromContext(ctx)

	// Önce YDSID'ye göre denetçiyi bul
	engineer, err := client.CompanyEngineer.Query().
		Where(companyengineer.YDSIDEQ(ydsid)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("denetçi bulunamadı (YDSID: %d): %v", ydsid, err)
	}

	// Bulunan denetçiyi ID'sine göre güncelle
	updatedEngineer, err := client.CompanyEngineer.UpdateOneID(engineer.ID).
		SetNillableName(input.Name).
		SetNillableAddress(input.Address).
		SetNillableEmail(input.Email).
		SetNillableTcNo(input.TcNo).
		SetNillablePhone(input.Phone).
		SetNillableRegisterNo(input.RegisterNo).
		SetNillableCertNo(input.CertNo).
		SetNillableCareer(input.Career).
		SetNillablePosition(input.Position).
		SetNillableNote(input.Note).
		SetNillableEmployment(input.Employment).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("denetçi güncellenemedi: %v", err)
	}

	return updatedEngineer, nil
}

// UpdateEngineerByYdsid is the resolver for the updateEngineerByYDSID field.
func (r *mutationResolver) UpdateEngineerByYdsid(ctx context.Context, ydsid int, input model.CompanyEngineerInput) (*ent.CompanyEngineer, error) {
	client := middlewares.GetClientFromContext(ctx)

	// Önce YDSID'ye göre denetçiyi bul
	engineer, err := client.CompanyEngineer.Query().
		Where(companyengineer.YDSIDEQ(ydsid)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("denetçi bulunamadı (YDSID: %d): %v", ydsid, err)
	}

	// Denetçi güncelleme builder'ını oluştur
	update := client.CompanyEngineer.UpdateOneID(engineer.ID).
		SetNillableName(input.Name).
		SetNillableAddress(input.Address).
		SetNillableEmail(input.Email).
		SetNillableTcNo(input.TcNo).
		SetNillablePhone(input.Phone).
		SetNillableRegisterNo(input.RegisterNo).
		SetNillableCertNo(input.CertNo).
		SetNillableCareer(input.Career).
		SetNillablePosition(input.Position).
		SetNillableNote(input.Note).
		SetNillableEmployment(input.Employment)

	// CompanyCode varsa şirket bağlantısını güncelle
	if input.CompanyCode != nil {
		company, err := client.CompanyDetail.Query().
			Where(companydetail.CompanyCodeEQ(*input.CompanyCode)).
			Only(ctx)
		if err == nil {
			// Şirket bulunduysa bağlantıyı güncelle
			update.SetCompany(company)
		} else {
			// Şirket bulunamadıysa sadece log at ve devam et
			log.Printf("Şirket bulunamadı (kod: %d): %v - Denetçi güncellemeye devam ediliyor", *input.CompanyCode, err)
		}
	}

	// Güncellemeyi gerçekleştir
	updatedEngineer, err := update.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("denetçi güncellenemedi: %v", err)
	}

	return updatedEngineer, nil
}

// Engineer is the resolver for the Engineer field.
func (r *queryResolver) Engineer(ctx context.Context, filter *model.EngineerFilterInput) ([]*ent.CompanyEngineer, error) {
	client := middlewares.GetClientFromContext(ctx)
	query := client.CompanyEngineer.Query()

	// Kullanıcı bilgilerini context'ten al
	customClaim := middlewares.CtxValue(ctx)
	if customClaim == nil {
		return nil, fmt.Errorf("kullanıcı kimliği bulunamadı")
	}

	// Kullanıcının bağlı olduğu şirketleri bul
	userCompanies, err := client.CompanyUser.Query().
		Where(companyuser.HasUserWith(user.IDEQ(customClaim.ID))).
		QueryCompany().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("kullanıcının şirketleri alınamadı: %v", err)
	}

	if len(userCompanies) == 0 {
		return nil, fmt.Errorf("kullanıcının bağlı olduğu şirket bulunamadı")
	}

	// Şirket ID'lerini topla
	companyIDs := make([]int, len(userCompanies))
	for i, company := range userCompanies {
		companyIDs[i] = company.ID
	}

	// Sorguyu kullanıcının şirketlerine göre filtrele
	query = query.Where(companyengineer.HasCompanyWith(companydetail.IDIn(companyIDs...)))

	// Filtreleri uygula
	if filter != nil {
		if filter.ID != nil && *filter.ID != "" {
			engineerID, err := strconv.Atoi(*filter.ID)
			if err != nil {
				return nil, fmt.Errorf("geçersiz mühendis ID'si: %v", err)
			}
			query = query.Where(companyengineer.IDEQ(engineerID))
		}
		if filter.Ydsid != nil {
			query = query.Where(companyengineer.YDSIDEQ(int(*filter.Ydsid)))
		}
	}

	// Sorguyu çalıştır ve sonuçları al
	engineers, err := query.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("mühendisler sorgulanırken hata oluştu: %v", err)
	}

	return engineers, nil
}

// EngineerByYdsid is the resolver for the engineerByYDSID field.
func (r *queryResolver) EngineerByYdsid(ctx context.Context, ydsid int) (*ent.CompanyEngineer, error) {
	client := middlewares.GetClientFromContext(ctx)

	engineer, err := client.CompanyEngineer.Query().
		Where(companyengineer.YDSIDEQ(ydsid)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("denetçi bulunamadı (YDSID: %d): %v", ydsid, err)
	}
	return engineer, nil
}

// CompanyEngineer returns generated.CompanyEngineerResolver implementation.
func (r *Resolver) CompanyEngineer() generated.CompanyEngineerResolver {
	return &companyEngineerResolver{r}
}

type companyEngineerResolver struct{ *Resolver }
