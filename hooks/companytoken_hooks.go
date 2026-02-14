package hooks

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/polatbilal/ent-gqlgen/ent"
	"github.com/polatbilal/ent-gqlgen/ent/hook"
	"github.com/polatbilal/ent-gqlgen/utils/crypto"
)

// CompanyTokenEncryptionHook şifreleme hook'unu döndürür
func CompanyTokenEncryptionHook() ent.Hook {
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.CompanyTokenFunc(func(ctx context.Context, m *ent.CompanyTokenMutation) (ent.Value, error) {
				// Encryption key'i environment'dan al
				encryptionKey := os.Getenv("ENCRYPTION_KEY")
				if len(encryptionKey) != 32 {
					log.Println("Warning: ENCRYPTION_KEY must be exactly 32 bytes for AES-256")
					return next.Mutate(ctx, m)
				}

				// YDKPassword set ediliyorsa şifrele
				if password, exists := m.YDKPassword(); exists && password != "" {
					encryptedPassword, err := crypto.Encrypt(password, []byte(encryptionKey))
					if err != nil {
						return nil, fmt.Errorf("failed to encrypt YDKPassword: %w", err)
					}
					m.SetYDKPassword(encryptedPassword)
				}

				return next.Mutate(ctx, m)
			})
		},
		ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
	)
}

// RegisterCompanyTokenHooks client'a hook'ları register eder
func RegisterCompanyTokenHooks(client *ent.Client) {
	client.CompanyToken.Use(CompanyTokenEncryptionHook())
}
