package sync

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/polatbilal/gqlgen-ent/handlers/external"
)

func CompanySync(c *fiber.Ctx) error {
	// Token'ı query parametresinden al
	token := c.Query("token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token bulunamadı"})
	}

	// WebSocket yükseltmesi için yeni bir handler
	return websocket.New(func(ws *websocket.Conn) {
		defer ws.Close()

		// Token'ı request header'a ekle
		c.Request().Header.Set("Authorization", "Bearer "+token)

		// Token kontrolü için companies'e istek at
		if err := external.YDKCompanies(c); err != nil {
			sendWSNotification(ws, "error", "Token kontrolü başarısız")
			return
		}

		// Token kontrolü için yanıtı al
		companyResponse := c.Locals("response").(fiber.Map)

		// Token hatası varsa işlemi başlatma
		if errMsg, hasError := companyResponse["error"]; hasError {
			if errType, ok := errMsg.(string); ok {
				if errType == "invalid_token" {
					sendWSNotification(ws, "error", "YDS Oturum süresi dolmuş. Lütfen YDS Oturumunuzu yenileyiniz.")
					return
				}
			}
			sendWSNotification(ws, "error", "Token hatası oluştu")
			return
		}

		// Token geçerliyse senkronizasyonu başlat
		sendWSNotification(ws, "info", "Senkronizasyon başladı")

		// Şirket yanıtında hata varsa işlemi sonlandır
		if errMsg, hasError := companyResponse["error"]; hasError {
			errorMessage := errMsg.(string)
			sendWSNotification(ws, "error", "Şirket bilgileri senkronizasyonu başarısız: "+errorMessage)
			return
		}

		// Denetçi senkronizasyonu başladı bildirimi
		sendWSNotification(ws, "info", "Denetçi bilgileri senkronizasyonu başladı")
		if err := external.YDKInspectors(c); err != nil {
			sendWSNotification(ws, "error", "Denetçi bilgileri senkronizasyonu başarısız")
			return
		}

		// Denetçi yanıtını al
		inspectorResponse := c.Locals("response").(fiber.Map)

		// Denetçi yanıtında hata varsa işlemi sonlandır
		if errMsg, hasError := inspectorResponse["error"]; hasError {
			errorMessage := errMsg.(string)
			if status, ok := inspectorResponse["status"]; ok {
				if status == float64(401) || status == float64(http.StatusUnauthorized) {
					sendWSNotification(ws, "error", "Token süresi dolmuş veya geçersiz. Lütfen yeniden giriş yapın.")
					return
				}
			}
			sendWSNotification(ws, "error", "Denetçi bilgileri senkronizasyonu başarısız: "+errorMessage)
			return
		}

		// Denetçi senkronizasyonu tamamlandı bildirimi
		sendWSNotification(ws, "success", "Denetçi bilgileri senkronizasyonu tamamlandı")

		// YİBF senkronizasyonu başladı bildirimi
		sendWSNotification(ws, "info", "YİBF bilgileri senkronizasyonu başladı")

		// YİBF senkronizasyonu tamamlandı bildirimi
		sendWSNotification(ws, "success", "YİBF bilgileri senkronizasyonu tamamlandı")

		// Birleştirilmiş yanıtı hazırla
		response := fiber.Map{
			"type":    "result",
			"status":  "success",
			"message": "Senkronizasyon tamamlandı",
			"details": fiber.Map{
				"company": companyResponse,
				"inspector": fiber.Map{
					"totalCount":   inspectorResponse["totalCount"],
					"successCount": inspectorResponse["successCount"],
					"skippedCount": inspectorResponse["skippedCount"],
					"message":      inspectorResponse["message"],
				},
			},
		}

		// Senkronizasyon tamamlandı bildirimi
		sendWSNotification(ws, "success", "Senkronizasyon başarıyla tamamlandı")

		// Son yanıtı gönder
		if err := ws.WriteJSON(response); err != nil {
			// WebSocket yazma hatası
			return
		}
	})(c)
}

// WebSocket üzerinden bildirim gönderme fonksiyonu
func sendWSNotification(ws *websocket.Conn, status string, message string) {
	notification := fiber.Map{
		"type":    "notification",
		"status":  status,
		"message": message,
		"time":    time.Now().Format(time.RFC3339),
	}

	ws.WriteJSON(notification)
}
