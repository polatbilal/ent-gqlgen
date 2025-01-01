package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/polatbilal/gqlgen-ent/handlers/service"

	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token gerekli"})
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
		"skip":              0,
		"take":              3,
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req, err := http.NewRequest("POST", svc.BaseURL+service.ENDPOINT_ALL, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := svc.Client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response struct {
		Items []struct {
			ID int `json:"id"`
		} `json:"items"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parse hatası: " + err.Error()})
		return
	}

	// Her bir ID için detaylı bilgileri almak için goroutine kullanacağız
	var wg sync.WaitGroup
	detailsChan := make(chan map[string]interface{}, len(response.Items))
	errorsChan := make(chan error, len(response.Items))

	for _, item := range response.Items {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			url := fmt.Sprintf("%s/api/yibf/findById/%d", svc.BaseURL, id)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				errorsChan <- err
				return
			}

			req.Header.Set("Authorization", token)

			resp, err := svc.Client.Do(req)
			if err != nil {
				errorsChan <- err
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				errorsChan <- err
				return
			}

			var findByIdResponse struct {
				ID      int `json:"id"`
				Village struct {
					Parent struct {
						Name   string `json:"name"`
						Parent struct {
							Name string `json:"name"`
						} `json:"parent"`
					} `json:"parent"`
				} `json:"village"`
			}

			if err := json.Unmarshal(body, &findByIdResponse); err != nil {
				errorsChan <- fmt.Errorf("ID %d için parse hatası: %v", id, err)
				return
			}

			detailsChan <- map[string]interface{}{
				"id":      findByIdResponse.ID,
				"village": findByIdResponse.Village,
			}
		}(item.ID)
	}

	// Tüm goroutine'lerin tamamlanmasını bekle
	go func() {
		wg.Wait()
		close(detailsChan)
		close(errorsChan)
	}()

	// Sonuçları topla
	var details []map[string]interface{}
	var errors []string

	for i := 0; i < len(response.Items); i++ {
		select {
		case detail := <-detailsChan:
			if detail != nil {
				details = append(details, detail)
			}
		case err := <-errorsChan:
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"details": details,
		"errors":  errors,
	})
}
