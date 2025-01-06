package middlewares

import (
	"context"
	"fmt"

	"github.com/polatbilal/gqlgen-ent/database"
	"github.com/polatbilal/gqlgen-ent/ent"
	"github.com/polatbilal/gqlgen-ent/services"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type authString string

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")

		if auth == "" {
			fmt.Println("Authorization hatası")
			c.Next() // Devam et
			return
		}

		bearer := "Bearer "
		if len(auth) > len(bearer) && auth[:len(bearer)] == bearer {
			auth = auth[len(bearer):]
		} else {
			c.JSON(403, map[string]string{
				"error": "Invalid token",
			})
			c.Abort() // İşlemi durdur
			return
		}

		// Redis'ten token'ı kontrol et
		_, err := database.RedisClient.Get(c.Request.Context(), auth).Result()
		if err == redis.Nil {
			c.JSON(401, map[string]string{
				"error": "Token has expired",
			})
			c.Abort() // İşlemi durdur
			return
		} else if err != nil {
			c.JSON(403, map[string]string{
				"error": "Invalid token",
			})
			c.Abort() // İşlemi durdur
			return
		}

		validate, err := services.JwtValidate(c.Request.Context(), auth)
		if err != nil {
			c.JSON(403, map[string]string{
				"error": "Invalid token",
			})
			c.Abort() // İşlemi durdur
			return
		}

		costumClaim, _ := validate.Claims.(*services.JwtCustomClaim)

		// Veritabanına bağlan
		client, err := database.GetClient()
		if err != nil {
			fmt.Println("veritabanına bağlanırken hata oluştu")
			c.Next() // Devam et
			return
		}

		ctx := context.WithValue(c.Request.Context(), "dbClient", client)
		ctx = context.WithValue(ctx, authString("auth"), costumClaim) // Auth bilgisini de ekle
		c.Request = c.Request.WithContext(ctx)

		c.Next() // Devam et
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
