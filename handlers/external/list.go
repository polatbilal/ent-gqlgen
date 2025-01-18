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
	// jwtToken := c.GetHeader("Authorization")
	// if jwtToken == "" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT Token gerekli"})
	// 	return
	// }

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

	// İki farklı durum için sorgu yapıları
	requestBodies := []map[string]interface{}{
		{
			"requireTotalCount": true,
			"searchOperation":   "contains",
			"searchValue":       nil,
			"skip":              5,
			"take":              4,
			"filter":            []interface{}{"state.id", "=", 1},
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
		},
		{
			"requireTotalCount": true,
			"searchOperation":   "contains",
			"searchValue":       nil,
			"skip":              5,
			"take":              4,
			"filter":            []interface{}{"state.id", "=", 2},
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
		},
		{
			"requireTotalCount": true,
			"searchOperation":   "contains",
			"searchValue":       nil,
			"skip":              5,
			"take":              4,
			"filter":            []interface{}{"state.id", "=", 3},
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
		},
		{
			"requireTotalCount": true,
			"searchOperation":   "contains",
			"searchValue":       nil,
			"skip":              5,
			"take":              4,
			"filter":            []interface{}{"state.id", "=", 4},
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
		},
		{
			"requireTotalCount": true,
			"searchOperation":   "contains",
			"searchValue":       nil,
			"skip":              5,
			"take":              4,
			"filter":            []interface{}{"state.id", "=", 5},
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
		},
		{
			"requireTotalCount": true,
			"searchOperation":   "contains",
			"searchValue":       nil,
			"skip":              5,
			"take":              4,
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
		},
	}

	type responseStruct struct {
		Items []struct {
			ID int `json:"id"`
		} `json:"items"`
		TotalCount int `json:"totalCount"`
	}

	var allYibfIDs []int
	var totalCount int

	// Her bir sorgu için istekleri yap
	for _, requestBody := range requestBodies {
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

		var response responseStruct
		if err := json.Unmarshal(body, &response); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("FindAll yanıt parse hatası: %v", err)})
			return
		}

		// ID'leri topla
		for _, item := range response.Items {
			allYibfIDs = append(allYibfIDs, item.ID)
		}
		totalCount += response.TotalCount
	}

	log.Printf("Toplam kayıt sayısı: %d", totalCount)
	log.Printf("Dönen kayıt sayısı: %d", len(allYibfIDs))

	result := gin.H{
		"success": true,
		"data": map[string]interface{}{
			"total": totalCount,
			"items": allYibfIDs,
		},
	}

	c.Set("response", result)
	c.JSON(http.StatusOK, result)
}
