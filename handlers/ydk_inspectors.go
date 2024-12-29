package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type YDKInspectorResponse struct {
	GroupCount int `json:"groupCount"`
	Items      []struct {
		Active      bool `json:"active"`
		Application struct {
			Active           bool `json:"active"`
			ApplicationState struct {
				Active bool   `json:"active"`
				ID     int    `json:"id"`
				Name   string `json:"name"`
			} `json:"applicationState"`
			ApplicationType struct {
				Active bool   `json:"active"`
				ID     int    `json:"id"`
				Name   string `json:"name"`
			} `json:"applicationType"`
			Person struct {
				ID              int    `json:"id"`
				IdentityNumber  string `json:"identityNumber"`
				FullName        string `json:"fullName"`
				LastPhoneNumber string `json:"lastPhoneNumber"`
				LastEPosta      string `json:"lastEPosta"`
				AddressStr      string `json:"addressStr"`
			} `json:"person"`
			Title struct {
				Name string `json:"name"`
			} `json:"title"`
			Tasks struct {
				Name string `json:"name"`
			} `json:"tasks"`
			RegistrationNumber string `json:"occupationalRegistrationNumber"`
			DocumentNumber     int    `json:"documentNumber"`
		} `json:"application"`
		Department struct {
			FileNumber int `json:"fileNumber"`
		} `json:"department"`
		TaskStartDate int64 `json:"taskStartDate"`
	} `json:"items"`
	TotalCount int `json:"totalCount"`
}

func YDKInspectors(c *gin.Context) {
	// GraphQL için JWT token
	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT Token gerekli"})
		return
	}

	// YDK API için token
	ydkToken := c.GetHeader("YDK-Token")
	if ydkToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "YDK Token gerekli"})
		return
	}

	service := &ExternalService{
		baseURL: os.Getenv("YDK_BASE_URL"),
		client:  &http.Client{},
	}

	if service.baseURL == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "YDK_BASE_URL environment variable is not set"})
		return
	}

	requestBody := map[string]interface{}{
		"requireTotalCount": true,
		"searchOperation":   "contains",
		"searchValue":       nil,
		"userData":          map[string]interface{}{},
		"sort": []map[string]interface{}{
			{
				"selector": "application.occupationalRegistrationNumber",
				"desc":     false,
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body oluşturma hatası: " + err.Error()})
		return
	}

	url := service.baseURL + "/api/departmentEmployee/findAllYdkApprovedEmployees"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ydkToken))
	req.Header.Set("Content-Type", "application/json")

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

	// Ham veriyi logla
	log.Printf("YDK API Ham Yanıt: %+v\n", string(body))

	// Struct'a parse et
	var ydkResponse YDKInspectorResponse
	if err := json.Unmarshal(body, &ydkResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parse hatası: " + err.Error()})
		return
	}

	log.Printf("Toplam denetçi sayısı: %d\n", ydkResponse.TotalCount)

	// Struct'tan gelen verileri sadeleştir
	type SimplifiedInspector struct {
		Name        string `json:"Name"`
		TcNo        string `json:"TcNo"`
		Phone       string `json:"Phone"`
		Email       string `json:"Email"`
		Address     string `json:"Address"`
		Career      string `json:"Career"`
		Position    string `json:"Position"`
		RegNo       string `json:"RegNo"`
		CertNo      int    `json:"CertNo"`
		YDSID       int    `json:"YDSID"`
		CompanyCode int    `json:"CompanyCode"`
		Employment  int64  `json:"Employment"`
	}

	var simplifiedData []SimplifiedInspector
	for _, item := range ydkResponse.Items {
		inspector := SimplifiedInspector{
			Name:        item.Application.Person.FullName,
			Address:     item.Application.Person.AddressStr,
			Email:       item.Application.Person.LastEPosta,
			TcNo:        item.Application.Person.IdentityNumber,
			Phone:       item.Application.Person.LastPhoneNumber,
			RegNo:       item.Application.RegistrationNumber,
			CertNo:      item.Application.DocumentNumber,
			Career:      item.Application.Title.Name,
			Position:    item.Application.Tasks.Name,
			Employment:  item.TaskStartDate,
			YDSID:       item.Application.Person.ID,
			CompanyCode: item.Department.FileNumber,
		}
		simplifiedData = append(simplifiedData, inspector)
	}

	// GraphQL client oluştur
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	graphqlClient := GraphQLClient{
		URL: fmt.Sprintf("%s://%s/graphql", scheme, c.Request.Host),
	}

	// Her bir denetçi için GraphQL mutation'ı çalıştır
	successCount := 0
	var processLogs []map[string]interface{}

	for _, inspector := range simplifiedData {
		log.Printf("Denetçi işleniyor: %s\n", inspector.Name)

		processLog := map[string]interface{}{
			"name": inspector.Name,
			"time": time.Now().Format("2006-01-02 15:04:05"),
		}

		// Unix timestamp'i tarihe çevir
		employment := time.Unix(inspector.Employment/1000, 0).Format("2006-01-02")

		mutation := `
		mutation CreateEngineer($input: CompanyEngineerInput!) {
			createEngineer(input: $input) {
				id
				Name
				TcNo
				RegNo
				Email
				Phone
				Career
				Position
				Employment
				Address
				Company {
					id
				}
			}
		}
		`

		variables := map[string]interface{}{
			"input": map[string]interface{}{
				"Name":        inspector.Name,
				"TcNo":        inspector.TcNo,
				"RegNo":       inspector.RegNo,
				"Email":       inspector.Email,
				"Phone":       inspector.Phone,
				"Career":      inspector.Career,
				"Position":    inspector.Position,
				"Address":     inspector.Address,
				"CertNo":      inspector.CertNo,
				"CompanyCode": inspector.CompanyCode,
				"YDSID":       inspector.YDSID,
				"Employment":  employment,
			},
		}

		processLog["variables"] = variables

		if err := graphqlClient.Execute(mutation, variables, jwtToken); err != nil {
			errMsg := fmt.Sprintf("Denetçi eklenirken hata oluştu: %v", err)

			if strings.Contains(err.Error(), "sahip denetçi zaten mevcut") {
				errMsg = fmt.Sprintf("Bu sicil numarasına sahip denetçi zaten mevcut: %s", inspector.RegNo)
				processLog["status"] = "skipped"
			} else {
				processLog["status"] = "error"
			}

			log.Printf("Hata: %s - Denetçi: %s\n", errMsg, inspector.Name)
			processLog["error"] = errMsg
			processLogs = append(processLogs, processLog)
			continue
		}

		successCount++
		processLog["status"] = "success"
		processLogs = append(processLogs, processLog)
		log.Printf("Başarılı: Denetçi eklendi - %s\n", inspector.Name)
	}

	log.Printf("İşlem tamamlandı. Toplam: %d, Başarılı: %d, Atlanan: %d\n",
		len(simplifiedData),
		successCount,
		len(simplifiedData)-successCount)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d adet denetçiden %d tanesi başarıyla eklendi, %d tanesi zaten mevcut",
			len(simplifiedData),
			successCount,
			len(simplifiedData)-successCount),
		"totalCount":   len(simplifiedData),
		"successCount": successCount,
		"skippedCount": len(simplifiedData) - successCount,
		"logs":         processLogs,
	})
}
