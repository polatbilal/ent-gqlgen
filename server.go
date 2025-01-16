// main.go

package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/polatbilal/gqlgen-ent/database"
	"github.com/polatbilal/gqlgen-ent/ent/migrate"
	"github.com/polatbilal/gqlgen-ent/graph/resolvers"
	"github.com/polatbilal/gqlgen-ent/handlers"
	"github.com/polatbilal/gqlgen-ent/middlewares"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func main() {
	r := gin.Default()

	// CORS ayarlarını middleware'den önce yapılandırın
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // Tüm originlere izin ver
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
		"YDK-Token",
		"Accept",
		"X-Requested-With",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Methods",
	}
	config.ExposeHeaders = []string{"Authorization", "Content-Length"}
	config.MaxAge = 12 * time.Hour

	r.Use(cors.New(config))
	r.Use(middlewares.AuthMiddleware())

	// Handlers API endpoint'lerini yapılandır
	handlers.SetupRoutes(r)

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Database connection
	client, err := database.GetClient()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer client.Close()

	// Run the migration here
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); !errors.Is(err, nil) {
		log.Fatalf("Error: failed creating schema resources %v\n", err)
	}

	// Redis bağlantısını başlat
	if err := database.InitRedis(); err != nil {
		log.Fatalf("Redis başlatma hatası: %v", err)
	}
	defer database.RedisClient.Close()

	// Configure the GraphQL server and start
	srv := handler.NewDefaultServer(resolvers.NewSchema(client))
	{
		r.POST("/graphql", func(c *gin.Context) {

			if c.Request.Method == "OPTIONS" {
				c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
				c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
				c.Writer.Header().Set("Access-Control-Max-Age", "86400")
				c.AbortWithStatus(http.StatusNoContent)
				return
			}

			srv.ServeHTTP(c.Writer, c.Request)

			// Handler'ın döndürdüğü hata nesnesini al
			err := c.Errors.Last()
			if err != nil {
				log.Printf("GraphQL Error: %v\n", err)

				gqlErr, ok := err.Err.(gqlerror.List)
				if ok && len(gqlErr) > 0 {
					c.JSON(http.StatusBadRequest, gin.H{
						"errors": gqlErr,
					})
					return
				}

				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
		})

		r.GET("/playground", func(c *gin.Context) {
			playground.Handler("Graphql", "/graphql").ServeHTTP(c.Writer, c.Request)
		})
	}

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.AbortWithStatus(http.StatusNoContent)
	})

	r.Run(":4000")
}
