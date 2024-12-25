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
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func main() {
	r := gin.Default()

	r.Use(middlewares.AuthMiddleware())
	// CORS ayarları
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: false,
	}))

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

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
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Redis bağlantısını kontrol et
	if err := database.RedisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Redis bağlantı hatası: %v", err)
	}

	// Configure the GraphQL server and start
	srv := handler.NewDefaultServer(resolvers.NewSchema(client))
	{
		r.POST("/graphql", func(c *gin.Context) {
			srv.ServeHTTP(c.Writer, c.Request)

			// Handler'ın döndürdüğü hata nesnesini al
			err := c.Errors.Last()
			if err != nil {
				// GQLGen'den gelen hata varsa işle
				gqlErr, ok := err.Err.(gqlerror.List)
				if ok && len(gqlErr) > 0 {
					// İlk hatayı al ve isteğe özel yanıt döndür
					c.String(http.StatusInternalServerError, gqlErr[0].Message)
					return
				}
			}
		})

		r.GET("/playground", func(c *gin.Context) {
			playground.Handler("Graphql", "/graphql").ServeHTTP(c.Writer, c.Request)
		})
	}

	r.Run("127.0.0.1:4000")
}
