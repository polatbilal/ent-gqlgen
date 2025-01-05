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
	"strings"
	"time"

	"github.com/polatbilal/gqlgen-ent/graph/helpers"
	"github.com/polatbilal/gqlgen-ent/handlers/client"
	"github.com/polatbilal/gqlgen-ent/handlers/service"

	"github.com/gin-gonic/gin"
)

func YDKInspectors(c *gin.Context) {
	// GraphQL için JWT token
	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT Token gerekli"})
		return
	}

	// YDK API için token
	ydkTokenJSON := c.GetHeader("YDK-Token")
	if ydkTokenJSON == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "YDK Token gerekli"})
		return
	}

	// [object Object] kontrolü
	if ydkTokenJSON == "[object Object]" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "YDK Token JSON formatında gönderilmelidir, JavaScript objesi olarak değil"})
		return
	}

	// YDK Token'ı JSON'dan parse et
	var ydkTokenResp service.YDKTokenResponse
	if err := json.Unmarshal([]byte(ydkTokenJSON), &ydkTokenResp); err != nil {
		log.Printf("Parse hatası detayı: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("YDK Token JSON parse hatası: %v - Gelen veri: %s", err, ydkTokenJSON)})
		return
	}

	if ydkTokenResp.AccessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "YDK Token içinde access_token bulunamadı"})
		return
	}

	svc := &service.ExternalService{
		BaseURL: os.Getenv("YDK_BASE_URL"),
		Client:  &http.Client{},
	}

	if svc.BaseURL == "" {
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

	url := svc.BaseURL + service.ENDPOINT_COMPANY_ENGİNNER
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ydkTokenResp.AccessToken))
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

	// Ham veriyi logla
	// log.Printf("YDK API Ham Yanıt: %+v\n", string(body))

	// Struct'a parse et
	var inspectorResponse service.YDKEngineerResponse
	if err := json.Unmarshal(body, &inspectorResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parse hatası: " + err.Error()})
		return
	}

	// log.Printf("Toplam denetçi sayısı: %d\n", inspectorResponse.TotalCount)

	// Struct'tan gelen verileri sadeleştir
	type SimplifiedInspector struct {
		Name        string `json:"Name"`
		TcNo        string `json:"TcNo"`
		Phone       string `json:"Phone"`
		Email       string `json:"Email"`
		Address     string `json:"Address"`
		Career      string `json:"Career"`
		Position    string `json:"Position"`
		RegisterNo  string `json:"RegisterNo"`
		CertNo      int    `json:"CertNo"`
		YDSID       int    `json:"YDSID"`
		CompanyCode int    `json:"CompanyCode"`
		Employment  int64  `json:"Employment"`
	}

	var simplifiedData []SimplifiedInspector
	for _, item := range inspectorResponse.Items {
		inspector := SimplifiedInspector{
			YDSID:       item.Application.User.ID,
			Name:        item.Application.User.Person.FullName,
			Address:     item.Application.User.Person.AddressStr,
			Email:       item.Application.User.Person.LastEPosta,
			TcNo:        item.Application.User.Person.IdentityNumber,
			Phone:       item.Application.User.Person.LastPhoneNumber,
			RegisterNo:  item.Application.RegistrationNumber,
			CertNo:      item.Application.DocumentNumber,
			Career:      item.Application.Title.Name,
			Position:    item.Application.Tasks.Name,
			Employment:  item.TaskStartDate,
			CompanyCode: item.Department.FileNumber,
		}
		simplifiedData = append(simplifiedData, inspector)
	}

	// GraphQL client oluştur
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	graphqlClient := client.GraphQLClient{
		URL: fmt.Sprintf("%s://%s/graphql", scheme, c.Request.Host),
	}

	// Her bir denetçi için GraphQL mutation'ı çalıştır
	successCount := 0
	skippedCount := 0
	var processLogs []map[string]interface{}

	for _, inspector := range simplifiedData {
		log.Printf("Denetçi işleniyor: %s (YDSID: %d)\n", inspector.Name, inspector.YDSID)

		processLog := map[string]interface{}{
			"name": inspector.Name,
			"time": time.Now().Format("2006-01-02 15:04:05"),
		}

		// Unix timestamp'i tarihe çevir
		employment := time.Unix(inspector.Employment/1000, 0).Format("2006-01-02")

		// RegisterNo'yu int'e çevir
		RegisterNoInt, err := strconv.Atoi(inspector.RegisterNo)
		if err != nil {
			errMsg := fmt.Sprintf("RegisterNo int'e çevrilemedi: %v", err)
			log.Printf("Hata: %s - Denetçi: %s\n", errMsg, inspector.Name)
			processLog["status"] = "error"
			processLog["error"] = errMsg
			processLogs = append(processLogs, processLog)
			continue
		}

		// Denetçi verilerini hazırla
		engineerData := map[string]interface{}{
			"Name":        inspector.Name,
			"TcNo":        helpers.SafeStringToInt(inspector.TcNo),
			"RegisterNo":  RegisterNoInt,
			"Email":       inspector.Email,
			"Phone":       inspector.Phone,
			"Career":      inspector.Career,
			"Position":    inspector.Position,
			"Address":     inspector.Address,
			"CertNo":      inspector.CertNo,
			"YDSID":       inspector.YDSID,
			"Employment":  employment,
			"CompanyCode": inspector.CompanyCode,
		}

		// Denetçiyi sorgula
		detailQuery := `
		query GetEngineerDetail($ydsid: Int!) {
			engineerByYDSID(ydsid: $ydsid) {
				id
				Name
				TcNo
				RegisterNo
				Email
				Phone
				Career
				Position
				Address
				CertNo
				YDSID
				Employment
				CompanyCode
			}
		}
		`

		var detailResult struct {
			EngineerByYDSID struct {
				ID          int    `json:"id"`
				Name        string `json:"Name"`
				TcNo        int    `json:"TcNo"`
				RegisterNo  int    `json:"RegisterNo"`
				Email       string `json:"Email"`
				Phone       string `json:"Phone"`
				Career      string `json:"Career"`
				Position    string `json:"Position"`
				Address     string `json:"Address"`
				CertNo      int    `json:"CertNo"`
				YDSID       int    `json:"YDSID"`
				Employment  string `json:"Employment"`
				CompanyCode int    `json:"CompanyCode"`
			} `json:"engineerByYDSID"`
		}

		err = graphqlClient.Execute(detailQuery, map[string]interface{}{"ydsid": inspector.YDSID}, jwtToken, &detailResult)

		if err != nil {
			// Hata mesajını kontrol et
			errMsg := err.Error()

			// Denetçi bulunamadı hatası (company_engineer not found)
			if strings.Contains(errMsg, "company_engineer not found") {
				// Denetçi yoksa ekle
				createMutation := `
				mutation CreateEngineer($input: CompanyEngineerInput!) {
					createEngineer(input: $input) {
						id
						Name
						YDSID
						CompanyCode
					}
				}
				`

				var createResult struct {
					CreateEngineer struct {
						ID          int    `json:"id"`
						Name        string `json:"Name"`
						YDSID       int    `json:"YDSID"`
						CompanyCode int    `json:"CompanyCode"`
					} `json:"createEngineer"`
				}

				if err := graphqlClient.Execute(createMutation, map[string]interface{}{"input": engineerData}, jwtToken, &createResult); err != nil {
					errMsg := fmt.Sprintf("Denetçi eklenirken hata oluştu: %v", err)
					log.Printf("Hata: %s - Denetçi: %s\n", errMsg, inspector.Name)
					processLog["status"] = "error"
					processLog["error"] = errMsg
					processLogs = append(processLogs, processLog)
					continue
				}

				successCount++
				processLog["status"] = "success"
				processLog["message"] = "Denetçi başarıyla eklendi"
				processLogs = append(processLogs, processLog)
				continue

				// Şirket bulunamadı hatası (company_detail not found)
			} else if strings.Contains(errMsg, "company_detail not found") {
				log.Printf("CompanyCode bulunamadı, güncelleme yapılacak - Denetçi: %s", inspector.Name)
				updateMutation := `
				mutation UpdateEngineer($ydsid: Int!, $input: CompanyEngineerInput!) {
					updateEngineerByYDSID(ydsid: $ydsid, input: $input) {
						id
						Name
						YDSID
						CompanyCode
					}
				}
				`

				var updateResult struct {
					UpdateEngineerByYDSID struct {
						ID          int    `json:"id"`
						Name        string `json:"Name"`
						YDSID       int    `json:"YDSID"`
						CompanyCode int    `json:"CompanyCode"`
					} `json:"updateEngineerByYDSID"`
				}

				updateVariables := map[string]interface{}{
					"ydsid": inspector.YDSID,
					"input": engineerData,
				}

				if err := graphqlClient.Execute(updateMutation, updateVariables, jwtToken, &updateResult); err != nil {
					errMsg := fmt.Sprintf("Denetçi güncellenirken hata oluştu: %v", err)
					log.Printf("Hata: %s - Denetçi: %s\n", errMsg, inspector.Name)
					processLog["status"] = "error"
					processLog["error"] = errMsg
					processLogs = append(processLogs, processLog)
					continue
				}

				successCount++
				processLog["status"] = "success"
				processLog["message"] = "Denetçi başarıyla güncellendi (CompanyCode değişikliği)"
				processLogs = append(processLogs, processLog)
				continue
			} else {
				// Diğer hatalar
				errMsg := fmt.Sprintf("Denetçi sorgulanırken hata oluştu: %v", err)
				log.Printf("Hata: %s - Denetçi: %s\n", errMsg, inspector.Name)
				processLog["status"] = "error"
				processLog["error"] = errMsg
				processLogs = append(processLogs, processLog)
				continue
			}
		}

		// Denetçi bulundu, değişiklikleri kontrol et
		currentEngineer := detailResult.EngineerByYDSID
		var changedFields []string
		needsUpdate := false

		// Alanları karşılaştır
		if currentEngineer.CompanyCode != inspector.CompanyCode {
			needsUpdate = true
			changedFields = append(changedFields, fmt.Sprintf("CompanyCode: %v -> %v",
				currentEngineer.CompanyCode, inspector.CompanyCode))
		}
		if currentEngineer.Name != inspector.Name {
			needsUpdate = true
			changedFields = append(changedFields, fmt.Sprintf("Name: %v -> %v",
				currentEngineer.Name, inspector.Name))
		}
		if currentEngineer.TcNo != helpers.SafeStringToInt(inspector.TcNo) {
			needsUpdate = true
			changedFields = append(changedFields, fmt.Sprintf("TcNo: %v -> %v",
				currentEngineer.TcNo, helpers.SafeStringToInt(inspector.TcNo)))
		}
		if currentEngineer.RegisterNo != RegisterNoInt {
			needsUpdate = true
			changedFields = append(changedFields, fmt.Sprintf("RegisterNo: %v -> %v",
				currentEngineer.RegisterNo, RegisterNoInt))
		}
		if currentEngineer.Email != inspector.Email {
			needsUpdate = true
			changedFields = append(changedFields, fmt.Sprintf("Email: %v -> %v",
				currentEngineer.Email, inspector.Email))
		}
		if currentEngineer.Phone != inspector.Phone {
			needsUpdate = true
			changedFields = append(changedFields, fmt.Sprintf("Phone: %v -> %v",
				currentEngineer.Phone, inspector.Phone))
		}
		if currentEngineer.Career != inspector.Career {
			needsUpdate = true
			changedFields = append(changedFields, fmt.Sprintf("Career: %v -> %v",
				currentEngineer.Career, inspector.Career))
		}
		if currentEngineer.Position != inspector.Position {
			needsUpdate = true
			changedFields = append(changedFields, fmt.Sprintf("Position: %v -> %v",
				currentEngineer.Position, inspector.Position))
		}
		if currentEngineer.Address != inspector.Address {
			needsUpdate = true
			changedFields = append(changedFields, fmt.Sprintf("Address: %v -> %v",
				currentEngineer.Address, inspector.Address))
		}
		if currentEngineer.CertNo != inspector.CertNo {
			needsUpdate = true
			changedFields = append(changedFields, fmt.Sprintf("CertNo: %v -> %v",
				currentEngineer.CertNo, inspector.CertNo))
		}

		if needsUpdate {
			// Değişiklik varsa güncelle
			log.Printf("Denetçi için değişiklikler tespit edildi - YDSID: %d, İsim: %s",
				inspector.YDSID, inspector.Name)
			for _, change := range changedFields {
				log.Printf("  - %s", change)
			}

			updateMutation := `
			mutation UpdateEngineer($ydsid: Int!, $input: CompanyEngineerInput!) {
				updateEngineerByYDSID(ydsid: $ydsid, input: $input) {
					id
					Name
					YDSID
					CompanyCode
				}
			}
			`

			var updateResult struct {
				UpdateEngineerByYDSID struct {
					ID          int    `json:"id"`
					Name        string `json:"Name"`
					YDSID       int    `json:"YDSID"`
					CompanyCode int    `json:"CompanyCode"`
				} `json:"updateEngineerByYDSID"`
			}

			updateVariables := map[string]interface{}{
				"ydsid": inspector.YDSID,
				"input": engineerData,
			}

			if err := graphqlClient.Execute(updateMutation, updateVariables, jwtToken, &updateResult); err != nil {
				errMsg := fmt.Sprintf("Denetçi güncellenirken hata oluştu: %v", err)
				log.Printf("Hata: %s - Denetçi: %s\n", errMsg, inspector.Name)
				processLog["status"] = "error"
				processLog["error"] = errMsg
				processLogs = append(processLogs, processLog)
				continue
			}

			successCount++
			processLog["status"] = "success"
			processLog["message"] = "Denetçi başarıyla güncellendi"
			processLogs = append(processLogs, processLog)
		} else {
			// Değişiklik yoksa atla
			skippedCount++
			processLog["status"] = "skipped"
			processLog["message"] = "Denetçi verileri güncel, güncelleme yapılmadı"
			processLogs = append(processLogs, processLog)
			log.Printf("Denetçi atlandı (değişiklik yok) - YDSID: %d, İsim: %s", inspector.YDSID, inspector.Name)
		}
	}

	log.Printf("İşlem tamamlandı. Toplam: %d, Başarılı: %d, Atlanan: %d\n",
		len(simplifiedData),
		successCount,
		skippedCount)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d adet denetçiden %d tanesi başarıyla eklendi/güncellendi, %d tanesi zaten güncel",
			len(simplifiedData),
			successCount,
			skippedCount),
		"totalCount":   len(simplifiedData),
		"successCount": successCount,
		"skippedCount": skippedCount,
		"logs":         processLogs,
	})
}
