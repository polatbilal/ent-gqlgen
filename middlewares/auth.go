package middlewares

import (
	"context"
	"fmt"
	"gqlgen-ent/database"
	"gqlgen-ent/ent"
	"gqlgen-ent/service"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

type authString string

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")

		if auth == "" {
			fmt.Println("Authorization hatası")
			return next(c)
		}

		bearer := "Bearer "
		if len(auth) > len(bearer) && auth[:len(bearer)] == bearer {
			auth = auth[len(bearer):]
		} else {
			return c.JSON(echo.ErrForbidden.Code, map[string]string{
				"error": "Invalid token",
			})
		}

		// Redis'ten token'ı kontrol et
		_, err := database.RedisClient.Get(c.Request().Context(), auth).Result()
		if err == redis.Nil {
			return c.JSON(echo.ErrUnauthorized.Code, map[string]string{
				"error": "Token has expired",
			})
		} else if err != nil {
			return c.JSON(echo.ErrForbidden.Code, map[string]string{
				"error": "Invalid token",
			})
		}

		validate, err := service.JwtValidate(c.Request().Context(), auth)
		if err != nil {
			return c.JSON(echo.ErrForbidden.Code, map[string]string{
				"error": "Invalid token",
			})
		}

		costumClaim, _ := validate.Claims.(*service.JwtCustomClaim)

		// Veritabanına bağlan
		client, err := database.GetClient(costumClaim.CompanyCode)
		if err != nil {
			fmt.Println("veritabanına bağlanırken hata oluştu")
			return next(c)
		}

		ctx := context.WithValue(c.Request().Context(), "companyCode", costumClaim.CompanyCode)
		ctx = context.WithValue(ctx, "dbClient", client)
		ctx = context.WithValue(ctx, authString("auth"), costumClaim) // Auth bilgisini de ekle
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}

func CtxValue(ctx context.Context) *service.JwtCustomClaim {
	raw, _ := ctx.Value(authString("auth")).(*service.JwtCustomClaim)
	return raw
}

func GetClientFromContext(ctx context.Context) *ent.Client {
	customClaim := CtxValue(ctx)
	client, err := database.GetClient(customClaim.CompanyCode)
	if err != nil {
		return nil
	}
	return client
}
