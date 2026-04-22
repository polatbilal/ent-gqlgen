package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/polatbilal/ent-gqlgen/ent"
	"github.com/polatbilal/ent-gqlgen/ent/companydetail"
	"github.com/polatbilal/ent-gqlgen/ent/financeaccount"
	"github.com/polatbilal/ent-gqlgen/ent/financeclass"
	"github.com/polatbilal/ent-gqlgen/ent/financegroup"
)

// CreateFinanceAccountForEntity, yeni bir entity (Mühendis veya Yapı Sahibi) oluşturulduğunda
// ona karşılık gelen ve verileri doldurulmuş yeni bir FinanceAccount kaydı oluşturur.
func CreateFinanceAccountForEntity(ctx context.Context, client *ent.Client, companyId int, name string, tcNo string, taxNo string, taxAdmin string, phone string, email string, address string, groupName string, entityType string) error {
	// 1. İlgili grubu bul (Yapı Sahibi, Mühendis vb.)
	group, err := client.FinanceGroup.Query().
		Where(financegroup.NameEQ(groupName)).
		Only(ctx)
	if err != nil {
		log.Printf("⚠️ Finans grubu bulunamadı (%s), FinanceAccount oluşturulmadı: %v", groupName, err)
		return fmt.Errorf("finans grubu bulunamadı (%s): %v", groupName, err)
	}

	// 2. İlgili sınıfı/tipi (MÜŞTERİ, PERSONEL) şirket bazlı bul
	// Not: Her şirketin kendi "Accounts" kategorisindeki sınıfları olmalı.
	types, err := client.FinanceClass.Query().
		Where(
			financeclass.Category("Accounts"),
			financeclass.NameEqualFold(entityType),
			financeclass.HasCompanyWith(companydetail.IDEQ(companyId)),
		).Only(ctx)
	if err != nil {
		log.Printf("⚠️ Finans tipi bulunamadı (Tip: %s, Şirket: %d), FinanceAccount oluşturulmadı: %v", entityType, companyId, err)
		return fmt.Errorf("finans tipi bulunamadı (%s, Şirket: %d): %v", entityType, companyId, err)
	}

	// 3. Uniqueness kontrolü (Bu şirket altında bu hesap var mı?)
	query := client.FinanceAccount.Query().Where(financeaccount.HasCompanyWith(companydetail.IDEQ(companyId)))
	var exists bool
	if tcNo != "" {
		exists, _ = query.Clone().Where(financeaccount.TcNoEQ(tcNo)).Exist(ctx)
	} else if taxNo != "" {
		exists, _ = query.Clone().Where(financeaccount.TaxNoEQ(taxNo)).Exist(ctx)
	} else if name != "" {
		exists, _ = query.Clone().Where(financeaccount.NameEQ(name)).Exist(ctx)
	}

	if exists {
		log.Printf("ℹ️  Bu kişi için zaten bir FinanceAccount mevcut: %s (Company: %d).", name, companyId)
		return nil
	}

	// 4. FinanceAccount kaydını oluştur
	create := client.FinanceAccount.Create().
		SetCompanyID(companyId).
		SetName(name).
		SetGroupID(group.ID).
		SetTypeID(types.ID).
		SetNote(fmt.Sprintf("%s referansıyla otomatik eklendi.", entityType))

	if tcNo != "" {
		create.SetTcNo(tcNo)
	}
	if taxNo != "" {
		create.SetNillableTaxNo(&taxNo)
	}
	if taxAdmin != "" {
		create.SetNillableTaxAdmin(&taxAdmin)
	}
	if phone != "" {
		create.SetNillablePhone(&phone)
	}
	if email != "" {
		create.SetNillableEmail(&email)
	}
	if address != "" {
		create.SetNillableAddress(&address)
	}

	_, err = create.Save(ctx)
	if err != nil {
		log.Printf("⚠️ FinanceAccount oluşturulamadı (%s, İsim: %s): %v", entityType, name, err)
		return fmt.Errorf("finans hesabı oluşturulamadı: %v", err)
	}

	log.Printf("✅ FinanceAccount oluşturuldu: %s (Company: %d) → %s", name, companyId, types.Name)
	return nil
}

// MigrateExistingAccounts taraması
func MigrateExistingAccounts(ctx context.Context, client *ent.Client) (string, error) {
	var created, failed int

	// 1. JobOwner'lar
	owners, err := client.JobOwner.Query().
		WithOwners(func(q *ent.JobRelationsQuery) { q.WithCompany() }).
		All(ctx)
	if err != nil {
		return "", fmt.Errorf("job owner'lar sorgulanırken hata: %v", err)
	}
	for _, owner := range owners {
		companyIDs := make(map[int]bool)
		for _, rel := range owner.Edges.Owners {
			if rel.Edges.Company != nil {
				companyIDs[rel.Edges.Company.ID] = true
			}
		}

		if len(companyIDs) == 0 {
			continue
		}

		for cID := range companyIDs {
			// "Müşteri" -> EqualFold sayesinde "MÜŞTERİ" ile eşleşecek
			if err := CreateFinanceAccountForEntity(ctx, client, cID, owner.Name, owner.TcNo, owner.TaxNo, owner.TaxAdmin, owner.Phone, owner.Email, owner.Address, "Yapı Sahibi", "Müşteri"); err != nil {
				failed++
			} else {
				created++
			}
		}
	}

	// 2. CompanyEngineer'lar
	engineers, err := client.CompanyEngineer.Query().WithCompany().All(ctx)
	if err != nil {
		return "", fmt.Errorf("mühendisler sorgulanırken hata: %v", err)
	}
	for _, eng := range engineers {
		if eng.Edges.Company == nil {
			continue
		}
		cID := eng.Edges.Company.ID
		if err := CreateFinanceAccountForEntity(ctx, client, cID, eng.Name, eng.TcNo, "", "", eng.Phone, eng.Email, eng.Address, "Mühendis", "Personel"); err != nil {
			failed++
		} else {
			created++
		}
	}

	result := fmt.Sprintf("Migration tamamlandı: %d yeni hesap oluşturuldu/atlandı, %d hata.", created, failed)
	log.Printf("📊 %s", result)
	return result, nil
}
