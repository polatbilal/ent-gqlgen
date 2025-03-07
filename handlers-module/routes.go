package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polatbilal/gqlgen-ent/handlers-module/external"
	"github.com/polatbilal/gqlgen-ent/handlers-module/sync"
)

// SetupRoutes tüm API endpoint'lerini yapılandırır
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
	companyGroup.Post("/list", external.YDKCompanies)
	companyGroup.Get("/sync", sync.CompanySync)

	// İş ile ilgili endpoint'ler
	jobsGroup := ydkGroup.Group("/jobs")
	jobsGroup.Post("/list", external.YibfList)             // Sadece liste çeker
	jobsGroup.Post("/detail", external.YibfDetailHandler)  // Liste içindeki yibf numarasına göre detayını çeker
	jobsGroup.Post("/payments", external.ProgressPayments) // Tek bir işin detayını çeker

	// Senkronizasyon endpoint'i
	jobsGroup.Post("/sync", sync.SyncYibf) // Liste + detay + DB kayıt
}
