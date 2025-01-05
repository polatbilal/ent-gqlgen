package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"
	"fmt"

	"github.com/polatbilal/gqlgen-ent/ent"
	"github.com/polatbilal/gqlgen-ent/ent/companydetail"
	"github.com/polatbilal/gqlgen-ent/graph/model"
	"github.com/polatbilal/gqlgen-ent/middlewares"
)

// CreateCompany is the resolver for the createCompany field.
func (r *mutationResolver) CreateCompany(ctx context.Context, input model.CompanyDetailInput) (*ent.CompanyDetail, error) {
	client := middlewares.GetClientFromContext(ctx)

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
		SetNillableChamberRegisterNo(input.ChamberRegisterNo).
		SetNillableVisaDate(input.VisaDate).
		SetNillableVisaEndDate(input.VisaEndDate).
		SetNillableVisaFinishedFor90Days(input.VisaFinishedFor90days).
		SetNillableCorePersonAbsent90Days(input.CorePersonAbsent90days).
		SetNillableIsClosed(input.IsClosed).
		SetNillableOwnerName(input.OwnerName).
		SetNillableOwnerTcNo(input.OwnerTcNo).
		SetNillableOwnerAddress(input.OwnerAddress).
		SetNillableOwnerPhone(input.OwnerPhone).
		SetNillableOwnerEmail(input.OwnerEmail).
		SetNillableOwnerRegisterNo(input.OwnerRegisterNo).
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
		SetNillableChamberRegisterNo(input.ChamberRegisterNo).
		SetNillableVisaDate(input.VisaDate).
		SetNillableVisaEndDate(input.VisaEndDate).
		SetNillableVisaFinishedFor90Days(input.VisaFinishedFor90days).
		SetNillableCorePersonAbsent90Days(input.CorePersonAbsent90days).
		SetNillableIsClosed(input.IsClosed).
		SetNillableOwnerName(input.OwnerName).
		SetNillableOwnerTcNo(input.OwnerTcNo).
		SetNillableOwnerAddress(input.OwnerAddress).
		SetNillableOwnerPhone(input.OwnerPhone).
		SetNillableOwnerEmail(input.OwnerEmail).
		SetNillableOwnerRegisterNo(input.OwnerRegisterNo).
		SetNillableOwnerCareer(input.OwnerCareer).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("şirket güncellenemedi: %v", err)
	}

	return updatedCompany, nil
}

// CompanyToken is the resolver for the CompanyToken field.
func (r *mutationResolver) CompanyToken(ctx context.Context, departmentID *int, input *model.CompanyTokenInput) (*ent.CompanyToken, error) {
	client := middlewares.GetClientFromContext(ctx)

	company, err := client.CompanyDetail.Query().Where(companydetail.DepartmentIdEQ(*departmentID)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("şirket bulunamadı: %v", err)
	}

	createCompanyToken, err := client.CompanyToken.Create().
		SetToken(*input.Token).
		SetCompany(company).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("şirket token oluşturulamadı: %v", err)
	}

	return createCompanyToken, nil
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

// CompanyToken is the resolver for the companyToken field.
func (r *queryResolver) CompanyToken(ctx context.Context, departmentID *int) (*ent.CompanyToken, error) {
	client := middlewares.GetClientFromContext(ctx)

	company, err := client.CompanyDetail.Query().Where(companydetail.DepartmentIdEQ(*departmentID)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("şirket bulunamadı: %v", err)
	}

	companyToken, err := company.QueryTokens().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("şirket token bulunamadı: %v", err)
	}

	return companyToken, nil
}
