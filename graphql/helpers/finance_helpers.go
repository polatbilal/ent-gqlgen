package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/polatbilal/ent-gqlgen/ent"
	"github.com/polatbilal/ent-gqlgen/ent/companyengineer"
	"github.com/polatbilal/ent-gqlgen/ent/companypersonnel"
	"github.com/polatbilal/ent-gqlgen/ent/financeaccount"
	"github.com/polatbilal/ent-gqlgen/ent/financegroup"
	"github.com/polatbilal/ent-gqlgen/ent/jobowner"
)

// CreateFinanceRelation, bir account oluşturulduğunda otomatik FinanceRelations kaydı oluşturur.
// accountType: "job_owner", "company_engineer", "company_personnel"
// groupName: FinanceGroup tablosundaki Name değeri (ör: "Yapı Sahibi", "Mühendis", "Personel")
func CreateFinanceRelation(ctx context.Context, client *ent.Client, accountType string, accountID int, groupName string) error {
	// İlgili grubu bul
	group, err := client.FinanceGroup.Query().
		Where(financegroup.NameEQ(groupName)).
		Only(ctx)
	if err != nil {
		log.Printf("⚠️ Finans grubu bulunamadı (%s), FinanceRelations oluşturulmadı: %v", groupName, err)
		return fmt.Errorf("finans grubu bulunamadı (%s): %v", groupName, err)
	}

	// FinanceRelations kaydını oluştur
	create := client.FinanceRelations.Create().
		SetGroupID(group.ID)

	switch accountType {
	case "job_owner":
		create.SetJobOwnerID(accountID)
	case "company_engineer":
		create.SetCompanyEngineerID(accountID)
	case "company_personnel":
		create.SetCompanyPersonnelID(accountID)
	case "finance_account":
		create.SetFinanceAccountID(accountID)
	default:
		return fmt.Errorf("bilinmeyen account tipi: %s", accountType)
	}

	_, err = create.Save(ctx)
	if err != nil {
		log.Printf("⚠️ FinanceRelations oluşturulamadı (%s, ID: %d): %v", accountType, accountID, err)
		return fmt.Errorf("finans ilişkisi oluşturulamadı: %v", err)
	}

	log.Printf("✅ FinanceRelations oluşturuldu: %s (ID: %d) → %s grubu", accountType, accountID, groupName)
	return nil
}

// MigrateExistingAccounts, mevcut JobOwner, CompanyEngineer ve CompanyPersonnel
// kayıtları için henüz FinanceRelations kaydı yoksa otomatik oluşturur.
func MigrateExistingAccounts(ctx context.Context, client *ent.Client) (string, error) {
	var created, skipped, failed int

	// 1. FinanceRelations'ı olmayan JobOwner'ları bul
	owners, err := client.JobOwner.Query().
		Where(jobowner.Not(jobowner.HasFinanceRelations())).
		All(ctx)
	if err != nil {
		return "", fmt.Errorf("job owner'lar sorgulanırken hata: %v", err)
	}
	for _, owner := range owners {
		if err := CreateFinanceRelation(ctx, client, "job_owner", owner.ID, "Yapı Sahibi"); err != nil {
			log.Printf("⚠️ JobOwner (ID: %d) için FinanceRelations oluşturulamadı: %v", owner.ID, err)
			failed++
		} else {
			created++
		}
	}

	// 2. FinanceRelations'ı olmayan CompanyEngineer'ları bul
	engineers, err := client.CompanyEngineer.Query().
		Where(companyengineer.Not(companyengineer.HasFinanceRelations())).
		All(ctx)
	if err != nil {
		return "", fmt.Errorf("mühendisler sorgulanırken hata: %v", err)
	}
	for _, eng := range engineers {
		if err := CreateFinanceRelation(ctx, client, "company_engineer", eng.ID, "Mühendis"); err != nil {
			log.Printf("⚠️ CompanyEngineer (ID: %d) için FinanceRelations oluşturulamadı: %v", eng.ID, err)
			failed++
		} else {
			created++
		}
	}

	// 3. FinanceRelations'ı olmayan CompanyPersonnel'leri bul
	personnels, err := client.CompanyPersonnel.Query().
		Where(companypersonnel.Not(companypersonnel.HasFinanceRelations())).
		All(ctx)
	if err != nil {
		return "", fmt.Errorf("personeller sorgulanırken hata: %v", err)
	}
	for _, pers := range personnels {
		if err := CreateFinanceRelation(ctx, client, "company_personnel", pers.ID, "Personel"); err != nil {
			log.Printf("⚠️ CompanyPersonnel (ID: %d) için FinanceRelations oluşturulamadı: %v", pers.ID, err)
			failed++
		} else {
			created++
		}
	}

	// 4. FinanceRelations'ı olmayan FinanceAccount'ları bul
	accounts, err := client.FinanceAccount.Query().
		Where(financeaccount.Not(financeaccount.HasFinanceRelations())).
		All(ctx)
	if err != nil {
		return "", fmt.Errorf("finans hesapları sorgulanırken hata: %v", err)
	}
	for _, acc := range accounts {
		if err := CreateFinanceRelation(ctx, client, "finance_account", acc.ID, "Cari Hesap"); err != nil {
			log.Printf("⚠️ FinanceAccount (ID: %d) için FinanceRelations oluşturulamadı: %v", acc.ID, err)
			failed++
		} else {
			created++
		}
	}

	skipped = 0 // Zaten ilişkisi olanlar sorguya dahil edilmedi

	result := fmt.Sprintf("Migration tamamlandı: %d oluşturuldu, %d atlandı, %d başarısız\n"+
		"Detay: %d JobOwner, %d CompanyEngineer, %d CompanyPersonnel, %d FinanceAccount işlendi",
		created, skipped, failed,
		len(owners), len(engineers), len(personnels), len(accounts))

	log.Printf("📊 %s", result)
	return result, nil
}
