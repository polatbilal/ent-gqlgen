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

	"github.com/polatbilal/gqlgen-ent/handlers/service"

	"github.com/gin-gonic/gin"
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

	requestBody := map[string]interface{}{
		"requireTotalCount": true,
		"searchOperation":   "contains",
		"searchValue":       nil,
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
		log.Printf("Request body oluşturma hatası: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req, err := http.NewRequest("POST", svc.BaseURL+service.ENDPOINT_YIBF_ALL, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Printf("Request oluşturma hatası: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", companyToken.Token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := svc.Client.Do(req)
	if err != nil {
		log.Printf("FindAll isteği başarısız: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Response body okuma hatası: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response struct {
		Items []struct {
			ID int `json:"id"`
		} `json:"items"`
		TotalCount int `json:"totalCount"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("FindAll yanıt parse hatası: %v", err)
		log.Printf("Parse edilemeyen yanıt: %s", string(body))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parse hatası: " + err.Error()})
		return
	}

	// ID listesini oluştur
	var idList []int
	for _, item := range response.Items {
		idList = append(idList, item.ID)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"total":   response.TotalCount,
		"items":   idList,
	})
}

// YibfListBatch YDK'dan gelen iş listesini alır ve toplu olarak işler
func YibfListBatch(c *gin.Context) {
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

	requestBody := map[string]interface{}{
		"requireTotalCount": true,
		"searchOperation":   "contains",
		"searchValue":       nil,
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
		log.Printf("Request body oluşturma hatası: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req, err := http.NewRequest("POST", svc.BaseURL+service.ENDPOINT_YIBF_ALL, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Printf("Request oluşturma hatası: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", companyToken.Token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := svc.Client.Do(req)
	if err != nil {
		log.Printf("FindAll isteği başarısız: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Response body okuma hatası: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response struct {
		Items []struct {
			ID int `json:"id"`
		} `json:"items"`
		TotalCount int `json:"totalCount"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("FindAll yanıt parse hatası: %v", err)
		log.Printf("Parse edilemeyen yanıt: %s", string(body))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parse hatası: " + err.Error()})
		return
	}

	// ID listesini oluştur
	var idList []int
	for _, item := range response.Items {
		idList = append(idList, item.ID)
	}

	// ID'leri toplu olarak işle
	processedIDs, failedIDs := ProcessYibfDetailBatch(c, idList)

	c.JSON(http.StatusOK, gin.H{
		"success":         true,
		"total":           response.TotalCount,
		"processed_count": len(processedIDs),
		"processed_ids":   processedIDs,
		"failed_count":    len(failedIDs),
		"failed_ids":      failedIDs,
	})
}
