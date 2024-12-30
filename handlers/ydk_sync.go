package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func YDKSync(c *gin.Context) {
	// GraphQL için JWT token kontrolü
	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT Token gerekli"})
		return
	}

	// YDK API için token kontrolü
	ydkTokenStr := c.GetHeader("YDK-Token")
	if ydkTokenStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "YDK Token gerekli"})
		return
	}

	// YDK token'ı JSON'dan parse et
	var ydkToken YDKTokenResponse
	if err := json.Unmarshal([]byte(ydkTokenStr), &ydkToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "YDK Token parse hatası: " + err.Error()})
		return
	}

	// Önce şirket bilgilerini çek ve kaydet
	log.Println("Şirket bilgileri çekiliyor...")
	YDKCompanies(c)

	// Şirket bilgileri başarıyla çekildiyse, denetçi bilgilerini çek
	log.Println("Şirket bilgileri başarıyla kaydedildi, denetçi bilgileri çekiliyor...")
	YDKInspectors(c)

	// Her iki işlemin sonucunu birleştir
	c.JSON(http.StatusOK, gin.H{
		"message": "Senkronizasyon tamamlandı",
		"status":  "success",
	})
}
