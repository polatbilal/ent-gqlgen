package middlewares

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/polatbilal/ent-gqlgen/database"
	"github.com/polatbilal/ent-gqlgen/ent"
	"github.com/polatbilal/ent-gqlgen/ent/user"
	"github.com/polatbilal/ent-gqlgen/services"
)

type authString string

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")

		if auth == "" {
			fmt.Println("Authorization hatası")
			return c.Next()
		}

		bearer := "Bearer "
		if len(auth) > len(bearer) && auth[:len(bearer)] == bearer {
			auth = auth[len(bearer):]
		} else {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Veritabanından token'ı kontrol et
		client, err := database.GetClient()
		if err != nil {
			fmt.Println("veritabanına bağlanırken hata oluştu")
			return c.Next()
		}

		// Token'ı veritabanında ara
		_, err = client.User.Query().
			Where(user.RefreshToken(auth)).
			Only(c.Context())
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token has expired",
			})
		}

		validate, err := services.JwtValidate(c.Context(), auth)
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		costumClaim, _ := validate.Claims.(*services.JwtCustomClaim)

		// Context'e client'ı ent.Client key'i ile ekle
		ctx := context.WithValue(c.Context(), "dbClient", client)
		ctx = context.WithValue(ctx, authString("auth"), costumClaim)
		c.SetUserContext(ctx)

		return c.Next()
	}
}

func CtxValue(ctx context.Context) *services.JwtCustomClaim {
	raw, _ := ctx.Value(authString("auth")).(*services.JwtCustomClaim)
	return raw
}

func GetClientFromContext(ctx context.Context) *ent.Client {
	client, err := database.GetClient()
	if err != nil {
		return nil
	}
	return client
}
