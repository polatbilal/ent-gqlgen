package sync

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/polatbilal/gqlgen-ent/handlers/external"
)

func SyncYibf(c *gin.Context) {
	// JWT token kontrolü
	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT Token gerekli"})
		return
	}

	// CompanyCode parametresini al
	companyCode := c.Query("companyCode")
	if companyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CompanyCode parametresi gerekli"})
		return
	}

	// YİBF listesini al
	listContext := &gin.Context{
		Request: c.Request.Clone(c.Request.Context()),
		Writer:  c.Writer,
		Keys:    make(map[string]interface{}),
	}
	listContext.Request.Header.Set("Authorization", jwtToken)

	// YibfList fonksiyonunu çağır
	external.YibfList(listContext)

	// Yanıtı kontrol et
	if len(listContext.Errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Liste alınamadı: %v", listContext.Errors)})
		return
	}

	// Liste yanıtını al
	listResult, exists := listContext.Get("response")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Liste yanıtı bulunamadı"})
		return
	}

	listResponse := listResult.(gin.H)
	if !listResponse["success"].(bool) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Liste alınamadı"})
		return
	}

	data := listResponse["data"].(map[string]interface{})
	yibfIDs := data["items"].([]int)
	total := data["total"].(int)

	log.Printf("Toplam YİBF sayısı: %d", total)
	log.Printf("İşlenecek YİBF sayısı: %d", len(yibfIDs))

	// Detay bilgilerini al
	detailContext := &gin.Context{
		Request: c.Request.Clone(c.Request.Context()),
		Writer:  c.Writer,
		Keys:    make(map[string]interface{}),
	}
	detailContext.Request.Header.Set("Authorization", jwtToken)

	// YibfDetail fonksiyonunu çağır
	external.YibfDetail(detailContext, yibfIDs, companyCode)

	// Yanıtı kontrol et
	if len(detailContext.Errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Detaylar alınamadı: %v", detailContext.Errors)})
		return
	}

	// Detay yanıtını al
	detailResult, exists := detailContext.Get("response")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Detay yanıtı bulunamadı"})
		return
	}

	detailResponse := detailResult.(gin.H)
	if !detailResponse["success"].(bool) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Detaylar işlenemedi"})
		return
	}

	detailData := detailResponse["data"].(map[string]interface{})

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": map[string]interface{}{
			"total_yibf":      total,
			"processed_count": detailData["processed_count"],
			"processed_ids":   detailData["processed_ids"],
			"failed_count":    detailData["failed_count"],
			"failed_ids":      detailData["failed_ids"],
			"results":         detailData["results"],
		},
	})
}
