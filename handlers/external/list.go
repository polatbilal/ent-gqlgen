package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/polatbilal/gqlgen-ent/handlers/service"
)

func YibfList(c *gin.Context) {
	// GraphQL için JWT token
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

	// CompanyCode'u integer'a çevir
	companyCodeInt, err := strconv.Atoi(companyCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz CompanyCode formatı"})
		return
	}

	// Token bilgisini veritabanından al
	companyToken, err := service.GetCompanyTokenFromDB(c.Request.Context(), companyCodeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if companyToken.Token == "" || companyToken.DepartmentId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçerli token veya department ID bulunamadı"})
		return
	}

	svc := &service.ExternalService{
		BaseURL: os.Getenv("YDK_BASE_URL"),
		Client:  &http.Client{},
	}

	log.Printf("YDK Base URL: %s", svc.BaseURL)

	requestBody := map[string]interface{}{
		"requireTotalCount": true,
		"searchOperation":   "contains",
		"searchValue":       nil,
		"skip":              0,
		"take":              10,
		"filter":            []interface{}{"state.id", "=", 6},
		"userData":          struct{}{},
		"sort": []map[string]interface{}{
			{
				"selector": "distributiondate",
				"desc":     true,
			},
			{
				"selector": "id",
				"desc":     true,
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Request body oluşturma hatası: %v", err)})
		return
	}

	req, err := http.NewRequest("POST", svc.BaseURL+service.ENDPOINT_YIBF_ALL, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Request oluşturma hatası: %v", err)})
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", companyToken.Token))
	req.Header.Set("Content-Type", "application/json")

	log.Printf("İstek gönderiliyor: %s", svc.BaseURL+service.ENDPOINT_YIBF_ALL)

	resp, err := svc.Client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("FindAll isteği başarısız: %v", err)})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Response body okuma hatası: %v", err)})
		return
	}

	var response struct {
		Items []struct {
			ID int `json:"id"`
		} `json:"items"`
		TotalCount int `json:"totalCount"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("FindAll yanıt parse hatası: %v", err)})
		return
	}

	log.Printf("Toplam kayıt sayısı: %d", response.TotalCount)
	log.Printf("Dönen kayıt sayısı: %d", len(response.Items))

	var yibfIDs []int
	for _, item := range response.Items {
		yibfIDs = append(yibfIDs, item.ID)
	}

	result := gin.H{
		"success": true,
		"data": map[string]interface{}{
			"total": response.TotalCount,
			"items": yibfIDs,
		},
	}

	c.Set("response", result)
	c.JSON(http.StatusOK, result)
}
