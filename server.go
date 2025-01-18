// main.go

package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/polatbilal/gqlgen-ent/database"
	"github.com/polatbilal/gqlgen-ent/ent/migrate"
	"github.com/polatbilal/gqlgen-ent/graph/resolvers"
	"github.com/polatbilal/gqlgen-ent/handlers"
	"github.com/polatbilal/gqlgen-ent/middlewares"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// HTTP handler'ı Fiber'a adapte eden yardımcı fonksiyon
func adaptHandler(h http.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Fiber context'inden auth bilgisini al
		ctx := c.UserContext()
		// HTTP request'i auth context'i ile oluştur
		req := c.Request()
		req.Header.SetCanonical([]byte("Authorization"), []byte(c.Get("Authorization")))

		handler := fasthttpadaptor.NewFastHTTPHandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				// Context'i aktar
				r = r.WithContext(ctx)
				h.ServeHTTP(w, r)
			},
		)
		handler(c.Context())
		return nil
	}
}

func main() {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	// CORS ayarları
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS, HEAD, PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With, Referer, Sec-Fetch-Site",
		ExposeHeaders:    "Content-Length, Authorization",
		AllowCredentials: false,
	}))

	// Her istek için header'ları ekle
	app.Use(func(c *fiber.Ctx) error {
		// Sec-Fetch-Site header'ını kontrol et
		secFetchSite := string(c.Request().Header.Peek("Sec-Fetch-Site"))

		var origin string
		if c.Get("Origin") != "" {
			origin = c.Get("Origin")
		} else if secFetchSite == "cross-site" {
			// Cross-site request ise ve origin yoksa, production URL'yi kullan
			origin = "https://dev.bilalpolat.tr"
		} else {
			// Local development için IP adresini kullan
			origin = "http://192.168.1.105:4000"
		}

		// CORS header'larını ayarla
		c.Set("Access-Control-Allow-Origin", origin)
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD, PATCH")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With, Referer, Sec-Fetch-Site")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}

		return c.Next()
	})

	app.Use(middlewares.AuthMiddleware())

	// Rotaları ayarla
	handlers.SetupRoutes(app)

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

	// Configure the GraphQL server
	srv := handler.New(resolvers.NewSchema(client))

	// GraphQL endpoint
	app.Post("/graphql", adaptHandler(srv))

	// GraphQL Playground
	app.Get("/playground", adaptHandler(playground.Handler("GraphQL", "/graphql")))

	// OPTIONS handler
	app.Options("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})

	log.Fatal(app.Listen(":4000"))
}
