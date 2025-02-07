package services

import (
	"context"
	"log"
	"time"

	"github.com/polatbilal/gqlgen-ent/ent"
	"github.com/polatbilal/gqlgen-ent/ent/companydetail"
	"github.com/polatbilal/gqlgen-ent/ent/companyuser"
	"github.com/polatbilal/gqlgen-ent/ent/user"
)

type LicenseChecker struct {
	client *ent.Client
	ticker *time.Ticker
	done   chan bool
}

func NewLicenseChecker(client *ent.Client) *LicenseChecker {
	return &LicenseChecker{
		client: client,
		ticker: time.NewTicker(1 * time.Hour), // Her saat başı kontrol et
		done:   make(chan bool),
	}
}

func (l *LicenseChecker) Start() {
	// Hemen ilk kontrolü yap
	if err := l.checkExpiredLicenses(); err != nil {
		log.Printf("İlk lisans kontrolü sırasında hata: %v", err)
	}

	go func() {
		for {
			select {
			case <-l.done:
				return
			case <-l.ticker.C:
				if err := l.checkExpiredLicenses(); err != nil {
					log.Printf("Lisans kontrolü sırasında hata: %v", err)
				}
			}
		}
	}()
}

func (l *LicenseChecker) Stop() {
	l.ticker.Stop()
	l.done <- true
}

func (l *LicenseChecker) checkExpiredLicenses() error {
	ctx := context.Background()

	now := time.Now().UTC() // UTC zamanını kullan
	log.Printf("Lisans kontrolü yapılıyor. Şu anki UTC zaman: %v", now)

	// Süresi dolmuş admin kullanıcılarını bul
	expiredAdmins, err := l.client.User.Query().
		Where(
			user.And(
				user.RoleEQ("Admin"),
				user.LicenseExpireDateLT(now),
			),
		).All(ctx)

	if err != nil {
		return err
	}

	// Her bir admin için işlem yap
	for _, admin := range expiredAdmins {
		// Admin'in bağlı olduğu şirketleri bul
		companies, err := admin.QueryCompanies().
			QueryCompany().
			All(ctx)

		if err != nil {
			log.Printf("Admin ID %d için şirket bilgileri alınamadı: %v", admin.ID, err)
			continue
		}

		// Şirket ID'lerini topla
		companyIDs := make([]int, len(companies))
		for i, company := range companies {
			companyIDs[i] = company.ID
		}

		// Bu şirketlere bağlı tüm aktif kullanıcıları bul ve deaktif et
		_, err = l.client.User.Update().
			Where(
				user.And(
					user.HasCompaniesWith(
						companyuser.HasCompanyWith(
							companydetail.IDIn(companyIDs...),
						),
					),
					user.ActiveEQ(true),
				),
			).
			SetActive(false).
			Save(ctx)

		if err != nil {
			log.Printf("Admin ID %d'nin şirketlerindeki kullanıcılar deaktif edilemedi: %v", admin.ID, err)
			continue
		}

		log.Printf("Admin ID %d için lisans süresi dolmuş, ilgili şirketlerdeki kullanıcılar deaktif edildi", admin.ID)
	}

	return nil
}
