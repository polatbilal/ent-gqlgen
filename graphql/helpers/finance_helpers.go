package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/polatbilal/ent-gqlgen/ent"
	"github.com/polatbilal/ent-gqlgen/ent/financeaccount"
	"github.com/polatbilal/ent-gqlgen/ent/financegroup"
)

// CreateFinanceAccountForEntity, yeni bir entity (Mühendis veya Yapı Sahibi) oluşturulduğunda
// ona karşılık gelen ve verileri doldurulmuş yeni bir FinanceAccount kaydı oluşturur.
// "ilişki olarak değil ama veriler oraya da kopyalansın şeklinde"
// entityType: "job_owner" veya "company_engineer"
// groupName: FinanceGroup tablosundaki Name değeri (ör: "Yapı Sahibi", "Mühendis")
func CreateFinanceAccountForEntity(ctx context.Context, client *ent.Client, entityType string, name string, tcNo string, taxNo string, taxAdmin string, phone string, email string, address string, groupName string) error {
	// İlgili grubu bul
	group, err := client.FinanceGroup.Query().
		Where(financegroup.NameEQ(groupName)).
		Only(ctx)
	if err != nil {
		log.Printf("⚠️ Finans grubu bulunamadı (%s), FinanceAccount oluşturulmadı: %v", groupName, err)
		return fmt.Errorf("finans grubu bulunamadı (%s): %v", groupName, err)
	}

	// Aynı tc_no (veya isim) ile daha önce bir FinanceAccount var mı kontrol et.
	var exists bool
	if tcNo != "" {
		exists, _ = client.FinanceAccount.Query().Where(financeaccount.TcNoEQ(tcNo)).Exist(ctx)
	} else if taxNo != "" {
		exists, _ = client.FinanceAccount.Query().Where(financeaccount.TaxNoEQ(taxNo)).Exist(ctx)
	} else if name != "" {
		exists, _ = client.FinanceAccount.Query().Where(financeaccount.NameEQ(name)).Exist(ctx)
	}

	if exists {
		log.Printf("ℹ️  Bu kişi için zaten bir FinanceAccount mevcut (veya atlandı): %s (TC: %s).", name, tcNo)
		// İlişki kurulmayacağı için, varolanla eşleştiğini düşünüp hata vermeden dönüyoruz.
		return nil
	}

	// FinanceAccount kaydını oluştur ve verileri kopyala
	create := client.FinanceAccount.Create().
		SetName(name).
		SetGroupID(group.ID).
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

	log.Printf("✅ FinanceAccount oluşturuldu: %s (İsim: %s) → %s grubu", entityType, name, groupName)
	return nil
}

// MigrateExistingAccounts, mevcut JobOwner ve CompanyEngineer kayıtları
// için henüz bir FinanceAccount hesap kaydı yoksa verileri kopyalayarak otomatik oluşturur.
func MigrateExistingAccounts(ctx context.Context, client *ent.Client) (string, error) {
	var created, failed int

	// 1. JobOwner'ları bul
	// owners, err := client.JobOwner.Query().All(ctx)
	// if err != nil {
	// 	return "", fmt.Errorf("job owner'lar sorgulanırken hata: %v", err)
	// }
	// for _, owner := range owners {
	// 	if err := CreateFinanceAccountForEntity(ctx, client, "job_owner", owner.Name, owner.TcNo, owner.TaxNo, owner.TaxAdmin, owner.Phone, owner.Email, owner.Address, "Yapı Sahibi"); err != nil {
	// 		log.Printf("⚠️ JobOwner (%s) için FinanceAccount işlemi başarısız: %v", owner.Name, err)
	// 		failed++
	// 	} else {
	// 		created++
	// 	}
	// }

	// 2. CompanyEngineer'ları bul
	engineers, err := client.CompanyEngineer.Query().All(ctx)
	if err != nil {
		return "", fmt.Errorf("mühendisler sorgulanırken hata: %v", err)
	}
	for _, eng := range engineers {
		if err := CreateFinanceAccountForEntity(ctx, client, "company_engineer", eng.Name, eng.TcNo, "", "", eng.Phone, eng.Email, eng.Address, "Mühendis"); err != nil {
			log.Printf("⚠️ CompanyEngineer (%s) için FinanceAccount işlemi başarısız: %v", eng.Name, err)
			failed++
		} else {
			created++
		}
	}

	// result := fmt.Sprintf("Migration tamamlandı: Tarama %d JobOwner, %d CompanyEngineer üzerinde yapıldı.\nBaşarılı/Mevcut Atlandı İşlem: %d, Başarısız: %d"),
	result := fmt.Sprintf("Migration tamamlandı: %d CompanyEngineer üzerinde yapıldı.\nBaşarılı/Mevcut Atlandı İşlem: %d, Başarısız: %d",
		// len(owners),
		len(engineers), created, failed)

	log.Printf("📊 %s", result)
	return result, nil
}
