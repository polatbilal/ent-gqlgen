package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/polatbilal/gqlgen-ent/handlers-module/service"
)

func YibfList(c *fiber.Ctx) error {

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

	// Token bilgisini veritabanından al
	companyToken, err := service.GetCompanyTokenFromDB(c.Context(), requestParams.CompanyCode)
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

	requestBodies := make([]map[string]interface{}, 0)

	// State ID'leri için döngü
	for stateID := 1; stateID <= 6; stateID++ {
		requestBody := map[string]interface{}{
			"requireTotalCount": true,
			"searchOperation":   "contains",
			"searchValue":       nil,
			"skip":              40,
			"take":              4,
			"filter":            []interface{}{"state.id", "=", stateID},
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
		requestBodies = append(requestBodies, requestBody)
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
