package sync

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/polatbilal/gqlgen-ent/handlers-module/external"
	"github.com/polatbilal/gqlgen-ent/handlers-module/service"
)

func SyncYibf(c *fiber.Ctx) error {
	// JWT token kontrolü
	jwtToken := c.Get("Authorization")
	if jwtToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT Token gerekli"})
	}

	// Request body'den parametreleri al
	var requestParams service.FrontendRequest
	if err := c.BodyParser(&requestParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz request body: " + err.Error(),
		})
	}

	// Parametreleri kontrol et
	if requestParams.CompanyCode == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "companyCode parametresi gerekli"})
	}

	// YİBF listesini al
	if err := external.YibfList(c); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Liste alınamadı"})
	}

	// Liste yanıtını al
	listMap := c.Locals("response").(fiber.Map)
	if !listMap["success"].(bool) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Liste alınamadı"})
	}

	data := listMap["data"].(fiber.Map)
	yibfIDs := data["items"].([]int)
	total := data["total"].(int)

	log.Printf("Toplam YİBF sayısı: %d", total)
	log.Printf("İşlenecek YİBF sayısı: %d", len(yibfIDs))

	// Detay bilgilerini al
	if err := external.YibfDetail(c, yibfIDs, requestParams.CompanyCode); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Detaylar alınamadı"})
	}

	// Detay yanıtını al
	detailMap := c.Locals("response").(fiber.Map)
	if !detailMap["success"].(bool) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Detaylar işlenemedi"})
	}

	detailData := detailMap["data"].(fiber.Map)

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"total_yibf":      total,
			"processed_count": detailData["processed_count"],
			"processed_ids":   detailData["processed_ids"],
			"failed_count":    detailData["failed_count"],
			"failed_ids":      detailData["failed_ids"],
			"results":         detailData["results"],
		},
	})
}
