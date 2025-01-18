// main.go

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/polatbilal/gqlgen-ent/database"
	"github.com/polatbilal/gqlgen-ent/ent/migrate"
	"github.com/polatbilal/gqlgen-ent/graph/resolvers"
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

	// Logger middleware'i ekle
	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("\n[REQUEST] Method: %s, Path: %s\n", c.Method(), c.Path())
		fmt.Printf("[HEADERS] Origin: %s\n", c.Get("Origin"))
		fmt.Printf("[HEADERS] Referer: %s\n", c.Get("Referer"))
		fmt.Printf("[HEADERS] Host: %s\n", c.Get("Host"))
		fmt.Printf("[HEADERS] User-Agent: %s\n", c.Get("User-Agent"))

		start := time.Now()
		err := c.Next()
		log.Printf("[%s] %s - %v", c.Method(), c.Path(), time.Since(start))
		return err
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
		// Debug için tüm header'ları logla
		c.Request().Header.VisitAll(func(key, value []byte) {
			fmt.Printf("[HEADER] %s: %s\n", string(key), string(value))
		})

		// Sec-Fetch-Site header'ını kontrol et
		secFetchSite := string(c.Request().Header.Peek("Sec-Fetch-Site"))
		fmt.Printf("[DEBUG] Sec-Fetch-Site: %s\n", secFetchSite)

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
	srv := handler.NewDefaultServer(resolvers.NewSchema(client))

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
