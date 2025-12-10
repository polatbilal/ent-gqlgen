// main.go

package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/polatbilal/gqlgen-ent/database"
	"github.com/polatbilal/gqlgen-ent/ent/migrate"
	"github.com/polatbilal/gqlgen-ent/graphql/resolvers"
	"github.com/polatbilal/gqlgen-ent/middlewares"
	"github.com/polatbilal/gqlgen-ent/services"
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
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Hata detaylarını loglama
			log.Printf("Error: %v", err)
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Logger middleware'ini ekliyoruz
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))

	// CORS ayarları
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, http://localhost:4000, http://192.168.1.8:4000, https://dev.bilalpolat.tr, https://main.bilalpolat.tr, https://yds.polat.dev",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           300,
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
			origin = "http://192.168.1.8:4000"
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

	// Lisans kontrolü job'ını başlat
	licenseChecker := services.NewLicenseChecker(client)
	licenseChecker.Start()
	defer licenseChecker.Stop()

	// Run the migration here
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); !errors.Is(err, nil) {
		log.Fatalf("Error: failed creating schema resources %v\n", err)
	}

	// Configure the GraphQL server
	srv := handler.NewDefaultServer(resolvers.NewSchema(client))

	// GraphQL endpoint
	app.Post("/graphql", func(c *fiber.Ctx) error {
		// Content-Type header'ını kontrol et
		if !c.Is("json") {
			return c.Status(400).JSON(fiber.Map{
				"error": "Content-Type must be application/json",
			})
		}

		return adaptHandler(srv)(c)
	})

	// GraphQL Playground
	app.Get("/playground", adaptHandler(playground.Handler("GraphQL", "/graphql")))

	// OPTIONS handler
	app.Options("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})

	log.Fatal(app.Listen(":4000"))
}
