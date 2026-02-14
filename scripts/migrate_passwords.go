package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/polatbilal/ent-gqlgen/database"
	"github.com/polatbilal/ent-gqlgen/utils/crypto"
)

func main() {
	// .env dosyasını yükle
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Encryption key'i al
	encryptionKey := os.Getenv("ENCRYPTION_KEY")
	if len(encryptionKey) != 32 {
		log.Fatal("ENCRYPTION_KEY must be exactly 32 bytes for AES-256")
	}

	// Database connection
	client, err := database.GetClient()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// Tüm CompanyToken kayıtlarını çek
	tokens, err := client.CompanyToken.Query().All(ctx)
	if err != nil {
		log.Fatalf("Error fetching tokens: %v", err)
	}

	log.Printf("Found %d company tokens to process", len(tokens))

	successCount := 0
	skipCount := 0
	errorCount := 0

	for _, token := range tokens {
		// YDKPassword boşsa atla
		if token.YDKPassword == "" {
			log.Printf("[SKIP] Token ID %d: YDKPassword is empty", token.ID)
			skipCount++
			continue
		}

		// Zaten şifreli mi kontrol et (decrypt edebiliyor muyuz?)
		_, err := crypto.Decrypt(token.YDKPassword, []byte(encryptionKey))
		if err == nil {
			// Decrypt başarılı = zaten şifreli
			log.Printf("[SKIP] Token ID %d: Already encrypted", token.ID)
			skipCount++
			continue
		}

		// Decrypt başarısız = düz metin olabilir, şifrele
		log.Printf("[ENCRYPTING] Token ID %d: Encrypting plain text password", token.ID)

		encryptedPassword, err := crypto.Encrypt(token.YDKPassword, []byte(encryptionKey))
		if err != nil {
			log.Printf("[ERROR] Token ID %d: Encryption failed: %v", token.ID, err)
			errorCount++
			continue
		}

		// Güncelle
		err = client.CompanyToken.UpdateOneID(token.ID).
			SetYDKPassword(encryptedPassword).
			Exec(ctx)
		if err != nil {
			log.Printf("[ERROR] Token ID %d: Update failed: %v", token.ID, err)
			errorCount++
			continue
		}

		log.Printf("[SUCCESS] Token ID %d: Password encrypted successfully", token.ID)
		successCount++
	}

	// Özet
	separator := strings.Repeat("=", 50)
	fmt.Println("\n" + separator)
	fmt.Println("MIGRATION SUMMARY")
	fmt.Println(separator)
	fmt.Printf("Total Tokens: %d\n", len(tokens))
	fmt.Printf("✅ Encrypted: %d\n", successCount)
	fmt.Printf("⏭️  Skipped:   %d\n", skipCount)
	fmt.Printf("❌ Errors:    %d\n", errorCount)
	fmt.Println(separator)

	if errorCount > 0 {
		log.Fatal("Migration completed with errors")
	}

	log.Println("Migration completed successfully!")
}
