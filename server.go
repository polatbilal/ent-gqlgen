// main.go

package main

import (
	"context"
	"errors"
	"gqlgen-ent/database"
	"gqlgen-ent/graph/resolvers"
	"gqlgen-ent/handlers"
	"gqlgen-ent/middlewares"
	"log"
	"net/http"
	"time"

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

	// CORS ayarlarını middleware'den önce yapılandırın
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // Tüm originlere izin ver
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

	// YDK Inspectors endpoint'ini ekle
	r.GET("/ydk/inspectors", handlers.YDKInspectors)

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Database connection
	companyCode := "3"
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
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

			if c.Request.Method == "OPTIONS" {
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

	r.Run("127.0.0.1:4000")
}
