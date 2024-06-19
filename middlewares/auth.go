package middlewares

import (
	"context"
	"fmt"
	"gqlgen-ent/database"
	"gqlgen-ent/ent"
	"gqlgen-ent/service"

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
			return c.String(echo.ErrForbidden.Code, "Invalid token")
		}

		validate, err := service.JwtValidate(context.Background(), auth)
		if err != nil {
			if err.Error() == "token expired" {
				return c.String(echo.ErrUnauthorized.Code, "Token has expired")
			} else if err.Error() == "token invalidated" {
				return c.String(echo.ErrForbidden.Code, "Token has been invalidated")
			} else {
				return c.String(echo.ErrForbidden.Code, "Invalid token")
			}
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
