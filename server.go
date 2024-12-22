// main.go

package main

import (
	"context"
	"errors"
	"gqlgen-ent/database"
	"gqlgen-ent/graph/resolvers"
	"gqlgen-ent/middlewares"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var redisClient *redis.Client

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.AuthMiddleware)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Authorization"},
	}))

	// Database connection
	companyCode := "2"
	client, err := database.GetClient(companyCode)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer client.Close()

	// Run the migration here
	if err := client.Schema.Create(context.Background()); !errors.Is(err, nil) {
		log.Fatalf("Error: failed creating schema resources %v\n", err)
	}

	// Redis bağlantısını başlat
	database.RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis sunucu adresi
		Password: "",               // Redis şifresi (varsa)
		DB:       0,                // Veritabanı numarası
	})

	// Configure the GraphQL server and start
	srv := handler.NewDefaultServer(resolvers.NewSchema(client))
	{
		e.POST("/graphql", func(c echo.Context) error {
			srv.ServeHTTP(c.Response(), c.Request())

			// Handler'ın döndürdüğü hata nesnesini al
			err := c.Get("error")
			if err != nil {
				// GQLGen'den gelen hata varsa işle
				gqlErr, ok := err.(gqlerror.List)
				if ok && len(gqlErr) > 0 {
					// İlk hatayı al ve isteğe özel yanıt döndür
					return c.String(http.StatusInternalServerError, gqlErr[0].Message)
				}
			}

			return nil
		})

		e.GET("/playground", func(c echo.Context) error {
			playground.Handler("Graphql", "/graphql").ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}

	e.Logger.Fatal(e.Start(":4000"))
}
