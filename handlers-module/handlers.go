package handlers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/polatbilal/gqlgen-ent/handlers-module/handlers/external"
	"github.com/polatbilal/gqlgen-ent/handlers-module/handlers/sync"
)

func SetupRoutes(app *fiber.App) {
	// Route'ları burada tanımlayın
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	// YDK API ana grubu
	ydkGroup := app.Group("/ydk")

	// Company ile ilgili endpoint'ler
	companyGroup := ydkGroup.Group("/company")
	companyGroup.Post("/inspectors", external.YDKInspectors)
	companyGroup.Get("/list", external.YDKCompanies)
	companyGroup.Get("/sync", sync.CompanySync)

	// İş ile ilgili endpoint'ler
	jobsGroup := ydkGroup.Group("/jobs")
	jobsGroup.Get("/list", external.YibfList) // Sadece liste çeker
	// jobsGroup.Get("/detail/:id", external.YibfDetail) // Tek bir işin detayını çeker
	jobsGroup.Post("/payments", external.ProgressPayments) // Tek bir işin detayını çeker

	// Senkronizasyon endpoint'i
	jobsGroup.Get("/sync", sync.SyncYibf) // Liste + detay + DB kayıt
}

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
