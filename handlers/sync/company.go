package sync

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/polatbilal/gqlgen-ent/handlers/external"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Geliştirme ortamı için tüm originlere izin veriyoruz
	},
}

func CompanySync(c *gin.Context) {
	// Token'ı query parametresinden al
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token bulunamadı"})
		return
	}

	// Token'ı request header'a ekle
	c.Request.Header.Set("Authorization", "Bearer "+token)

	// WebSocket bağlantısını yükselt
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "WebSocket bağlantısı kurulamadı"})
		return
	}
	defer ws.Close()

	// Yanıtları yakalamak için ResponseWriter'ı wrap et
	companyWriter := &ResponseCapturer{ResponseWriter: c.Writer}

	// Token kontrolü için önce companies'e istek at
	companyCtx := &gin.Context{
		Request: c.Request.Clone(c.Request.Context()),
		Writer:  companyWriter,
		Params:  c.Params,
		Keys:    c.Keys,
	}
	companyCtx.Request.Header = c.Request.Header.Clone()

	// Token kontrolü için companies'e istek at
	external.YDKCompanies(companyCtx)

	// Token kontrolü için yanıtı parse et
	var tokenCheckResponse map[string]interface{}
	if err := json.Unmarshal(companyWriter.body, &tokenCheckResponse); err != nil {
		sendWSNotification(ws, "error", "Token kontrolü başarısız")
		return
	}

	// Token hatası varsa işlemi başlatma
	if _, hasError := tokenCheckResponse["error"]; hasError {
		// YDK API'den gelen JSON yanıtını kontrol et
		if errType, ok := tokenCheckResponse["error"].(string); ok {
			if errType == "invalid_token" {
				sendWSNotification(ws, "error", "YDS Oturum süresi dolmuş. Lütfen YDS Oturumunuzu yenileyiniz.")
				return
			}
		}
		// Diğer hatalar için genel mesaj
		sendWSNotification(ws, "error", "Token hatası oluştu")
		return
	}

	// Token geçerliyse senkronizasyonu başlat
	sendWSNotification(ws, "info", "Senkronizasyon başladı")

	// Şirket yanıtını kullan
	companyResponse := tokenCheckResponse

	// Şirket yanıtında hata varsa işlemi sonlandır
	if errMsg, hasError := companyResponse["error"]; hasError {
		errorMessage := errMsg.(string)
		sendWSNotification(ws, "error", "Şirket bilgileri senkronizasyonu başarısız: "+errorMessage)
		return
	}

	// Diğer writer'ları hazırla
	inspectorWriter := &ResponseCapturer{ResponseWriter: c.Writer}
	// yibfWriter := &ResponseCapturer{ResponseWriter: c.Writer}

	// Diğer context'leri hazırla
	inspectorCtx := &gin.Context{
		Request: c.Request.Clone(c.Request.Context()),
		Writer:  inspectorWriter,
		Params:  c.Params,
		Keys:    c.Keys,
	}
	inspectorCtx.Request.Header = c.Request.Header.Clone()

	// yibfCtx := &gin.Context{
	// 	Request: c.Request.Clone(c.Request.Context()),
	// 	Writer:  yibfWriter,
	// 	Params:  c.Params,
	// 	Keys:    c.Keys,
	// }
	// yibfCtx.Request.Header = c.Request.Header.Clone()

	// Denetçi senkronizasyonu başladı bildirimi
	sendWSNotification(ws, "info", "Denetçi bilgileri senkronizasyonu başladı")
	external.YDKInspectors(inspectorCtx)

	// Denetçi yanıtını parse et
	var inspectorResponse map[string]interface{}
	if err := json.Unmarshal(inspectorWriter.body, &inspectorResponse); err != nil {
		sendWSNotification(ws, "error", "Denetçi bilgileri senkronizasyonu başarısız")
		return
	}

	// Denetçi yanıtında hata varsa işlemi sonlandır
	if errMsg, hasError := inspectorResponse["error"]; hasError {
		errorMessage := errMsg.(string)
		if inspectorResponse["status"] == float64(401) || inspectorResponse["status"] == float64(http.StatusUnauthorized) {
			sendWSNotification(ws, "error", "Token süresi dolmuş veya geçersiz. Lütfen yeniden giriş yapın.")
			return
		}
		sendWSNotification(ws, "error", "Denetçi bilgileri senkronizasyonu başarısız: "+errorMessage)
		return
	}

	// Denetçi senkronizasyonu tamamlandı bildirimi
	sendWSNotification(ws, "success", "Denetçi bilgileri senkronizasyonu tamamlandı")

	// YİBF senkronizasyonu başladı bildirimi
	sendWSNotification(ws, "info", "YİBF bilgileri senkronizasyonu başladı")
	// external.YibfList(yibfCtx)

	// YİBF yanıtını parse et
	// var yibfResponse map[string]interface{}
	// if err := json.Unmarshal(yibfWriter.body, &yibfResponse); err != nil {
	// 	sendWSNotification(ws, "error", "YİBF bilgileri senkronizasyonu başarısız")
	// 	return
	// }

	// YİBF yanıtında hata varsa işlemi sonlandır
	// if errMsg, hasError := yibfResponse["error"]; hasError {
	// 	errorMessage := errMsg.(string)
	// 	if yibfResponse["status"] == float64(401) || yibfResponse["status"] == float64(http.StatusUnauthorized) {
	// 		sendWSNotification(ws, "error", "Token süresi dolmuş veya geçersiz. Lütfen yeniden giriş yapın.")
	// 		return
	// 	}
	// 	sendWSNotification(ws, "error", "YİBF bilgileri senkronizasyonu başarısız: "+errorMessage)
	// 	return
	// }

	// YİBF senkronizasyonu tamamlandı bildirimi
	sendWSNotification(ws, "success", "YİBF bilgileri senkronizasyonu tamamlandı")

	// Birleştirilmiş yanıtı hazırla
	response := gin.H{
		"type":    "result",
		"status":  "success",
		"message": "Senkronizasyon tamamlandı",
		"details": gin.H{
			"company": companyResponse,
			"inspector": gin.H{
				"totalCount":   inspectorResponse["totalCount"],
				"successCount": inspectorResponse["successCount"],
				"skippedCount": inspectorResponse["skippedCount"],
				"message":      inspectorResponse["message"],
			},
			// "yibf": gin.H{
			// 	"total":           yibfResponse["total"],
			// 	"processed_count": yibfResponse["processed_count"],
			// 	"failed_count":    yibfResponse["failed_count"],
			// 	"success":         yibfResponse["success"],
			// },
		},
	}

	// Senkronizasyon tamamlandı bildirimi
	sendWSNotification(ws, "success", "Senkronizasyon başarıyla tamamlandı")

	// Son yanıtı gönder
	ws.WriteJSON(response)
}

// ResponseCapturer yanıtları yakalamak için kullanılır
type ResponseCapturer struct {
	gin.ResponseWriter
	body []byte
}

func (w *ResponseCapturer) Write(b []byte) (int, error) {
	w.body = b
	return len(b), nil
}

// WebSocket üzerinden bildirim gönderme fonksiyonu
func sendWSNotification(ws *websocket.Conn, status string, message string) {
	notification := gin.H{
		"type":    "notification",
		"status":  status,
		"message": message,
		"time":    time.Now().Format(time.RFC3339),
	}

	ws.WriteJSON(notification)
}
