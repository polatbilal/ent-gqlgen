package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polatbilal/gqlgen-ent/handlers-module/handlers/external"
	"github.com/polatbilal/gqlgen-ent/handlers-module/handlers/sync"
)

// SetupRoutes tüm API endpoint'lerini yapılandırır
func SetupRoutes(app *fiber.App) {
	// YDK API ana grubu
	ydkGroup := app.Group("/ydk")

	// Company ile ilgili endpoint'ler
	companyGroup := ydkGroup.Group("/company")
	companyGroup.Post("/inspectors", external.YDKInspectors)
	companyGroup.Post("/list", external.YDKCompanies)
	companyGroup.Get("/sync", sync.CompanySync)

	// İş ile ilgili endpoint'ler
	jobsGroup := ydkGroup.Group("/jobs")
	jobsGroup.Get("/list", external.YibfList) // Sadece liste çeker
	// jobsGroup.Get("/detail/:id", external.YibfDetail) // Tek bir işin detayını çeker
	jobsGroup.Post("/payments", external.ProgressPayments) // Tek bir işin detayını çeker

	// Senkronizasyon endpoint'i
	jobsGroup.Get("/sync", sync.SyncYibf) // Liste + detay + DB kayıt
}
