package middlewares

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/polatbilal/gqlgen-ent/api-core/database"
	"github.com/polatbilal/gqlgen-ent/api-core/ent"
	"github.com/polatbilal/gqlgen-ent/api-core/services"
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

		// Redis'ten token'ı kontrol et
		_, err := database.RedisClient.Get(c.Context(), auth).Result()
		if err == redis.Nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token has expired",
			})
		} else if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		validate, err := services.JwtValidate(c.Context(), auth)
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		costumClaim, _ := validate.Claims.(*services.JwtCustomClaim)

		// Veritabanına bağlan
		client, err := database.GetClient()
		if err != nil {
			fmt.Println("veritabanına bağlanırken hata oluştu")
			return c.Next()
		}

		// Context'e client'ı ent.Client key'i ile ekle
		// ctx := ent.NewContext(c.Context(), client)
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
	// return ent.FromContext(ctx)
	client, err := database.GetClient()
	if err != nil {
		return nil
	}
	return client
}
