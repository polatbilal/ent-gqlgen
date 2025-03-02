package service

import (
	"context"
	"fmt"

	"github.com/polatbilal/gqlgen-ent/api-core/database"
	"github.com/polatbilal/gqlgen-ent/api-core/ent"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/companydetail"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/jobdetail"
)

// GetCompanyTokenFromDB şirket token bilgisini veritabanından alır
func GetCompanyTokenFromDB(ctx context.Context, companyCode int) (*ent.CompanyToken, error) {
	// Veritabanı bağlantısını al
	dbClient, err := database.GetClient()
	if err != nil {
		return nil, fmt.Errorf("veritabanı bağlantısı kurulamadı: %w", err)
	}

	// Önce CompanyDetail'den şirketi bul
	company, err := dbClient.CompanyDetail.Query().
		Where(companydetail.CompanyCode(companyCode)).
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("şirket bulunamadı: %w", err)
	}

	// Şirketin token bilgisini al
	companyToken, err := company.QueryTokens().First(ctx)
	if err != nil {
		return nil, fmt.Errorf("token bilgisi bulunamadı: %w", err)
	}

	if companyToken.Token == "" || companyToken.DepartmentId == 0 {
		return nil, fmt.Errorf("geçerli token veya department ID bulunamadı")
	}

	return companyToken, nil
}

// GetCompanyCodeFromYibf iş numarasından şirket kodunu alır
func GetCompanyCodeFromYibf(ctx context.Context, yibfNo int) (int, error) {
	// Veritabanı bağlantısını al
	dbClient, err := database.GetClient()
	if err != nil {
		return 0, fmt.Errorf("veritabanı bağlantısı kurulamadı: %w", err)
	}

	// İş kaydını bul ve şirket kodunu al
	job, err := dbClient.JobDetail.Query().
		Where(jobdetail.YibfNoEQ(yibfNo)).
		QueryCompany().
		Select(companydetail.FieldCompanyCode).
		First(ctx)
	if err != nil {
		return 0, fmt.Errorf("iş kaydı bulunamadı: %w", err)
	}

	return job.CompanyCode, nil
}
