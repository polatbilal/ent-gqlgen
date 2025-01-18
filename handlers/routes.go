package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/polatbilal/gqlgen-ent/handlers/external"
	"github.com/polatbilal/gqlgen-ent/handlers/sync"
)

// SetupRoutes tüm API endpoint'lerini yapılandırır
func SetupRoutes(r *gin.Engine) {
	// YDK API ana grubu
	ydkGroup := r.Group("/ydk")
	{
		// Company ile ilgili endpoint'ler
		companyGroup := ydkGroup.Group("/company")
		{
			companyGroup.POST("/inspectors", external.YDKInspectors)
			companyGroup.GET("/list", external.YDKCompanies)
			companyGroup.GET("/sync", sync.CompanySync)
		}

		// İş ile ilgili endpoint'ler
		jobsGroup := ydkGroup.Group("/jobs")
		{
			// Liste ve detay endpoint'leri
			jobsGroup.GET("/list", external.YibfList) // Sadece liste çeker
			// jobsGroup.GET("/detail/:id", external.YibfDetail) // Tek bir işin detayını çeker

			// Senkronizasyon endpoint'i
			jobsGroup.GET("/sync", sync.SyncYibf) // Liste + detay + DB kayıt
		}
	}
}
