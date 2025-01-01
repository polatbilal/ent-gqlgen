package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func YDKSync(c *gin.Context) {
	// CORS ve SSE header'larını ayarla
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, YDK-Token")

	// OPTIONS isteğine yanıt ver
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	// Senkronizasyon başladı bildirimi
	sendNotification(c, "info", "Senkronizasyon başladı")

	// Token bilgilerini header veya query'den al
	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		jwtToken = c.Query("jwt")
	}
	if jwtToken == "" {
		sendNotification(c, "error", "JWT Token gerekli")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT Token gerekli"})
		return
	}

	// YDK token bilgisini header veya query'den al
	ydkTokenStr := c.GetHeader("YDK-Token")
	if ydkTokenStr == "" {
		ydkTokenStr = c.Query("ydk")
	}
	if ydkTokenStr == "" {
		sendNotification(c, "error", "YDK Token gerekli")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "YDK Token gerekli"})
		return
	}

	// YDK token'ı JSON'dan parse et
	var ydkToken YDKTokenResponse
	if err := json.Unmarshal([]byte(ydkTokenStr), &ydkToken); err != nil {
		sendNotification(c, "error", "YDK Token parse hatası: "+err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "YDK Token parse hatası: " + err.Error()})
		return
	}

	// Yanıtları yakalamak için ResponseWriter'ı wrap et
	companyWriter := &ResponseCapturer{ResponseWriter: c.Writer}
	inspectorWriter := &ResponseCapturer{ResponseWriter: c.Writer}

	// Context'leri kopyala
	companyCtx := &gin.Context{
		Request: c.Request,
		Writer:  companyWriter,
	}
	inspectorCtx := &gin.Context{
		Request: c.Request,
		Writer:  inspectorWriter,
	}

	// Şirket senkronizasyonu başladı bildirimi
	sendNotification(c, "info", "Şirket bilgileri senkronizasyonu başladı")
	YDKCompanies(companyCtx)

	// Şirket yanıtını parse et
	var companyResponse map[string]interface{}
	if err := json.Unmarshal(companyWriter.body, &companyResponse); err != nil {
		sendNotification(c, "error", "Şirket bilgileri senkronizasyonu başarısız")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Şirket yanıtı parse edilemedi"})
		return
	}

	// Şirket senkronizasyonu tamamlandı bildirimi
	sendNotification(c, "success", "Şirket bilgileri senkronizasyonu tamamlandı")

	// Denetçi senkronizasyonu başladı bildirimi
	sendNotification(c, "info", "Denetçi bilgileri senkronizasyonu başladı")
	YDKInspectors(inspectorCtx)

	// Denetçi yanıtını parse et
	var inspectorResponse map[string]interface{}
	if err := json.Unmarshal(inspectorWriter.body, &inspectorResponse); err != nil {
		sendNotification(c, "error", "Denetçi bilgileri senkronizasyonu başarısız")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Denetçi yanıtı parse edilemedi"})
		return
	}

	// Denetçi senkronizasyonu tamamlandı bildirimi
	sendNotification(c, "success", "Denetçi bilgileri senkronizasyonu tamamlandı")

	// Birleştirilmiş yanıtı hazırla
	response := gin.H{
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
		},
	}

	// Senkronizasyon tamamlandı bildirimi
	sendNotification(c, "success", "Senkronizasyon başarıyla tamamlandı")

	// Yanıtı gönder
	c.JSON(http.StatusOK, response)
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

// Bildirim gönderme fonksiyonu
func sendNotification(c *gin.Context, status string, message string) {
	notification := gin.H{
		"type":    "notification",
		"status":  status,
		"message": message,
		"time":    time.Now().Format(time.RFC3339),
	}

	// SSE formatında bildirim gönder
	c.SSEvent("message", notification)
	c.Writer.Flush()
}
