package handlers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func StartServer(ctx context.Context) error {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		AppName:           "Handlers Service",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Printf("[HANDLERS] Error: %v", err)
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// CORS ayarları
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, https://dev.bilalpolat.tr, https://main.bilalpolat.tr, https://ydts.bilalpolat.tr",
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
			origin = "http://192.168.1.105:4001"
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

	// Logger middleware'i ekle
	app.Use(logger.New(logger.Config{
		Format:     "[HANDLERS] ${time} | ${status} | ${latency} | ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))

	// Mevcut route'ları ekle
	SetupRoutes(app)

	// Server'ı başlat
	go func() {
		if err := app.Listen(":4001"); err != nil {
			log.Printf("Handlers server error: %v", err)
		}
	}()

	// Context iptal edilene kadar bekle
	<-ctx.Done()
	return app.Shutdown()
}
