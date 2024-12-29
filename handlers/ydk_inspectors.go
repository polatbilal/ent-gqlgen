package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

func writeToFile(filename string, data interface{}) error {
	// Logs klasörünü oluştur
	logsDir := "logs"
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		return fmt.Errorf("logs klasörü oluşturma hatası: %v", err)
	}

	// Dosya adına timestamp ekle
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename = filepath.Join(logsDir, fmt.Sprintf("%s_%s.json", filename, timestamp))

	// Veriyi JSON formatına çevir
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("JSON dönüşüm hatası: %v", err)
	}

	// Dosyaya yaz
	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		return fmt.Errorf("dosya yazma hatası: %v", err)
	}

	return nil
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
		baseURL: "https://businessyds.csb.gov.tr",
		client:  &http.Client{},
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

	// Önce ham veriyi kaydet
	var rawResponse interface{}
	if err := json.Unmarshal(body, &rawResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Raw response parse hatası: " + err.Error()})
		return
	}
	if err := writeToFile("ydk_raw_response", rawResponse); err != nil {
		log.Printf("Ham yanıt kaydetme hatası: %v", err)
	}

	// Struct'a parse et
	var ydkResponse YDKInspectorResponse
	if err := json.Unmarshal(body, &ydkResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parse hatası: " + err.Error()})
		return
	}

	// Parse edilmiş struct verisini kaydet
	if err := writeToFile("ydk_parsed_response", ydkResponse); err != nil {
		log.Printf("Parse edilmiş yanıt kaydetme hatası: %v", err)
	}

	// Struct'tan gelen verileri sadeleştir
	type SimplifiedInspector struct {
		Name        string `json:"name"`
		TcNo        string `json:"tcNo"`
		Phone       string `json:"phone"`
		Email       string `json:"email"`
		Address     string `json:"address"`
		Career      string `json:"career"`
		Position    string `json:"position"`
		RegNo       string `json:"regNo"`
		CertNo      int    `json:"certNo"`
		YDSID       int    `json:"id"`
		CompanyCode int    `json:"CompanyCode"`
		Employment  int64  `json:"employment"`
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

	// Sadeleştirilmiş veriyi kaydet
	if err := writeToFile("ydk_simplified_response", map[string]interface{}{
		"totalCount": ydkResponse.TotalCount,
		"groupCount": ydkResponse.GroupCount,
		"items":      simplifiedData,
	}); err != nil {
		log.Printf("Sadeleştirilmiş yanıt kaydetme hatası: %v", err)
	}

	// GraphQL client oluştur
	graphqlClient := GraphQLClient{
		URL: "http://localhost:4000/graphql",
	}

	// Her bir denetçi için GraphQL mutation'ı çalıştır
	successCount := 0
	var processLogs []map[string]interface{}

	for _, inspector := range simplifiedData {
		processLog := map[string]interface{}{
			"name": inspector.Name,
			"time": time.Now().Format("2006-01-02 15:04:05"),
		}

		// Unix timestamp'i tarihe çevir
		employment := time.Unix(inspector.Employment/1000, 0).Format("2006-01-02")

		mutation := `
		mutation CreateEngineer($input: CompanyEngineerInput!) {
			createEngineer(input: $input) {
				YDSID
				Name
				TcNo
				RegNo
				Email
				Phone
				Career
				Position
				Employment
				Address
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
			processLog["error"] = errMsg
			processLogs = append(processLogs, processLog)
			log.Printf("Hata: %s - Denetçi: %s", errMsg, inspector.Name)
			continue
		}

		successCount++
		processLog["status"] = "success"
		processLogs = append(processLogs, processLog)
		log.Printf("Başarılı: Denetçi eklendi - %s", inspector.Name)
	}

	// İşlem loglarını kaydet
	if err := writeToFile("ydk_process_logs", map[string]interface{}{
		"totalCount":   len(simplifiedData),
		"successCount": successCount,
		"logs":         processLogs,
	}); err != nil {
		log.Printf("İşlem logları kaydetme hatası: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      fmt.Sprintf("%d adet denetçiden %d tanesi başarıyla eklendi", len(simplifiedData), successCount),
		"totalCount":   len(simplifiedData),
		"successCount": successCount,
		"logs":         processLogs,
	})
}
