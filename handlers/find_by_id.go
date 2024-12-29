package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindById(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token gerekli"})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parametresi gerekli"})
		return
	}

	service := &ExternalService{
		baseURL: "https://businessyds.csb.gov.tr",
		client:  &http.Client{},
	}

	url := fmt.Sprintf("%s/api/yibf/findById/%s", service.baseURL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Authorization", token)

	resp, err := service.client.Do(req)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parse hatası: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"id":      findByIdResponse.ID,
		"village": findByIdResponse.Village,
	})
}
