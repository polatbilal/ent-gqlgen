package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.61

import (
	"context"
	"fmt"

	"github.com/polatbilal/gqlgen-ent/ent"
	"github.com/polatbilal/gqlgen-ent/ent/companydetail"
	"github.com/polatbilal/gqlgen-ent/graph/generated"
	"github.com/polatbilal/gqlgen-ent/graph/model"
	"github.com/polatbilal/gqlgen-ent/middlewares"
	"github.com/polatbilal/gqlgen-ent/tools"
)

// VisaDate is the resolver for the VisaDate field.
func (r *companyDetailResolver) VisaDate(ctx context.Context, obj *ent.CompanyDetail) (*string, error) {
	if obj.VisaDate.IsZero() {
		return nil, nil
	}
	visaDate := obj.VisaDate.Format("2006-01-02")
	return &visaDate, nil
}

// VisaEndDate is the resolver for the VisaEndDate field.
func (r *companyDetailResolver) VisaEndDate(ctx context.Context, obj *ent.CompanyDetail) (*string, error) {
	if obj.VisaEndDate.IsZero() {
		return nil, nil
	}
	visaEndDate := obj.VisaEndDate.Format("2006-01-02")
	return &visaEndDate, nil
}

// CreateCompany is the resolver for the createCompany field.
func (r *mutationResolver) CreateCompany(ctx context.Context, input model.CompanyDetailInput) (*ent.CompanyDetail, error) {
	client := middlewares.GetClientFromContext(ctx)

	visaDatePtr, err := tools.ParseDate(input.VisaDate)
	if err != nil {
		return nil, fmt.Errorf("visa date dönüşüm hatası: %v", err)
	}

	visaEndDatePtr, err := tools.ParseDate(input.VisaEndDate)
	if err != nil {
		return nil, fmt.Errorf("visa end date dönüşüm hatası: %v", err)
	}

	createCompany, err := client.CompanyDetail.Create().
		SetCompanyCode(input.CompanyCode).
		SetName(input.Name).
		SetNillableAddress(input.Address).
		SetNillablePhone(input.Phone).
		SetNillableEmail(input.Email).
		SetNillableWebsite(input.Website).
		SetNillableTaxAdmin(input.TaxAdmin).
		SetNillableTaxNo(input.TaxNo).
		SetNillableChamberInfo(input.ChamberInfo).
		SetNillableChamberRegNo(input.ChamberRegNo).
		SetNillableVisaDate(visaDatePtr).
		SetNillableVisaEndDate(visaEndDatePtr).
		SetNillableVisaFinishedFor90Days(input.VisaFinishedFor90days).
		SetNillableCorePersonAbsent90Days(input.CorePersonAbsent90days).
		SetNillableIsClosed(input.IsClosed).
		SetNillableOwnerName(input.OwnerName).
		SetNillableOwnerTcNo(input.OwnerTcNo).
		SetNillableOwnerAddress(input.OwnerAddress).
		SetNillableOwnerPhone(input.OwnerPhone).
		SetNillableOwnerEmail(input.OwnerEmail).
		SetNillableOwnerRegNo(input.OwnerRegNo).
		SetNillableOwnerBirthDate(input.OwnerBirthDate).
		SetNillableOwnerCareer(input.OwnerCareer).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("şirket oluşturulamadı: %v", err)
	}

	return createCompany, nil
}

// UpdateCompany is the resolver for the updateCompany field.
func (r *mutationResolver) UpdateCompany(ctx context.Context, input model.CompanyDetailInput) (*ent.CompanyDetail, error) {
	client := middlewares.GetClientFromContext(ctx)
	visaDatePtr, err := tools.ParseDate(input.VisaDate)
	if err != nil {
		return nil, fmt.Errorf("visa date dönüşüm hatası: %v", err)
	}
	visaEndDatePtr, err := tools.ParseDate(input.VisaEndDate)
	if err != nil {
		return nil, fmt.Errorf("visa end date dönüşüm hatası: %v", err)
	}

	company, err := client.CompanyDetail.Query().Where(companydetail.CompanyCodeEQ(input.CompanyCode)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("şirket bulunamadı: %v", err)
	}

	updatedCompany, err := client.CompanyDetail.UpdateOneID(company.ID).
		SetName(input.Name).
		SetCompanyCode(input.CompanyCode).
		SetNillableAddress(input.Address).
		SetNillablePhone(input.Phone).
		SetNillableEmail(input.Email).
		SetNillableWebsite(input.Website).
		SetNillableTaxAdmin(input.TaxAdmin).
		SetNillableTaxNo(input.TaxNo).
		SetNillableChamberInfo(input.ChamberInfo).
		SetNillableChamberRegNo(input.ChamberRegNo).
		SetNillableVisaDate(visaDatePtr).
		SetNillableVisaEndDate(visaEndDatePtr).
		SetNillableVisaFinishedFor90Days(input.VisaFinishedFor90days).
		SetNillableCorePersonAbsent90Days(input.CorePersonAbsent90days).
		SetNillableIsClosed(input.IsClosed).
		SetNillableOwnerName(input.OwnerName).
		SetNillableOwnerTcNo(input.OwnerTcNo).
		SetNillableOwnerAddress(input.OwnerAddress).
		SetNillableOwnerPhone(input.OwnerPhone).
		SetNillableOwnerEmail(input.OwnerEmail).
		SetNillableOwnerRegNo(input.OwnerRegNo).
		SetNillableOwnerBirthDate(input.OwnerBirthDate).
		SetNillableOwnerCareer(input.OwnerCareer).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("şirket güncellenemedi: %v", err)
	}

	return updatedCompany, nil
}

// CompanyByCode is the resolver for the companyByCode field.
func (r *queryResolver) CompanyByCode(ctx context.Context, companyCode int) (*ent.CompanyDetail, error) {
	client := middlewares.GetClientFromContext(ctx)
	company, err := client.CompanyDetail.Query().Where(companydetail.CompanyCodeEQ(companyCode)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("şirket bulunamadı: %v", err)
	}
	return company, nil
}

// CompanyDetail returns generated.CompanyDetailResolver implementation.
func (r *Resolver) CompanyDetail() generated.CompanyDetailResolver { return &companyDetailResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type companyDetailResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
