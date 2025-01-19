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

	"github.com/gofiber/fiber/v2"
	"github.com/polatbilal/gqlgen-ent/handlers/service"
)

func YibfList(c *fiber.Ctx) error {
	// CompanyCode parametresini al
	companyCode := c.Query("companyCode")
	if companyCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "CompanyCode parametresi gerekli"})
	}

	// CompanyCode'u integer'a çevir
	companyCodeInt, err := strconv.Atoi(companyCode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz CompanyCode formatı"})
	}

	// Token bilgisini veritabanından al
	companyToken, err := service.GetCompanyTokenFromDB(c.Context(), companyCodeInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if companyToken.Token == "" || companyToken.DepartmentId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçerli token veya department ID bulunamadı"})
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
			// "skip":              5,
			// "take":              4,
			"filter":   []interface{}{"state.id", "=", 1},
			"userData": struct{}{},
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
			// "skip":              5,
			// "take":              4,
			"filter":   []interface{}{"state.id", "=", 2},
			"userData": struct{}{},
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
			// "skip":              5,
			// "take":              4,
			"filter":   []interface{}{"state.id", "=", 3},
			"userData": struct{}{},
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
			// "skip":              5,
			// "take":              4,
			"filter":   []interface{}{"state.id", "=", 4},
			"userData": struct{}{},
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
			// "skip":              5,
			// "take":              4,
			"filter":   []interface{}{"state.id", "=", 5},
			"userData": struct{}{},
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
			// "skip":              5,
			// "take":              4,
			"filter":   []interface{}{"state.id", "=", 6},
			"userData": struct{}{},
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
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Request body oluşturma hatası: %v", err)})
		}

		req, err := http.NewRequest("POST", svc.BaseURL+service.ENDPOINT_YIBF_ALL, bytes.NewBuffer(jsonBody))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Request oluşturma hatası: %v", err)})
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", companyToken.Token))
		req.Header.Set("Content-Type", "application/json")

		resp, err := svc.Client.Do(req)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("FindAll isteği başarısız: %v", err)})
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Response body okuma hatası: %v", err)})
		}

		var response responseStruct
		if err := json.Unmarshal(body, &response); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("FindAll yanıt parse hatası: %v", err)})
		}

		// ID'leri topla
		for _, item := range response.Items {
			allYibfIDs = append(allYibfIDs, item.ID)
		}
		totalCount += response.TotalCount
	}

	log.Printf("Toplam kayıt sayısı: %d", totalCount)
	log.Printf("Dönen kayıt sayısı: %d", len(allYibfIDs))

	result := fiber.Map{
		"success": true,
		"data": fiber.Map{
			"total": totalCount,
			"items": allYibfIDs,
		},
	}

	c.Locals("response", result)
	return c.JSON(result)
}
