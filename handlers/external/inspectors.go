package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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
	var processLogs []map[string]interface{}

	for _, inspector := range simplifiedData {
		log.Printf("Denetçi işleniyor: %s\n", inspector.Name)

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

		// Önce denetçinin var olup olmadığını kontrol et
		checkQuery := `
		query CheckEngineer($ydsid: Int!) {
			engineer(filter: {YDSID: $ydsid}) {
				id
				YDSID
				Name
			}
		}
		`

		checkVariables := map[string]interface{}{
			"ydsid": inspector.YDSID,
		}

		var checkResult struct {
			Engineer []map[string]interface{} `json:"engineer"`
		}

		err = graphqlClient.Execute(checkQuery, checkVariables, jwtToken, &checkResult)

		// Denetçi verilerini hazırla
		engineerData := map[string]interface{}{
			"Name":        inspector.Name,
			"TcNo":        inspector.TcNo,
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

		// Eğer hata varsa ve bu "not found" hatası değilse
		if err != nil && !strings.Contains(err.Error(), "not found") {
			errMsg := fmt.Sprintf("Denetçi kontrolü sırasında hata oluştu: %v", err)
			log.Printf("Hata: %s - Denetçi: %s\n", errMsg, inspector.Name)
			processLog["status"] = "error"
			processLog["error"] = errMsg
			processLogs = append(processLogs, processLog)
			continue
		}

		// Denetçi var mı kontrolü
		if len(checkResult.Engineer) > 0 {
			// Mevcut denetçi detaylarını al
			detailQuery := `
			query GetEngineerDetail($ydsid: Int!) {
				engineer(filter: {YDSID: $ydsid}) {
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
				Engineer []map[string]interface{} `json:"engineer"`
			}

			err = graphqlClient.Execute(detailQuery, checkVariables, jwtToken, &detailResult)
			if err != nil {
				errMsg := fmt.Sprintf("Denetçi detayları alınırken hata oluştu: %v", err)
				log.Printf("Hata: %s - Denetçi: %s\n", errMsg, inspector.Name)
				processLog["status"] = "error"
				processLog["error"] = errMsg
				processLogs = append(processLogs, processLog)
				continue
			}

			if len(detailResult.Engineer) == 0 {
				errMsg := "Denetçi detayları bulunamadı"
				log.Printf("Hata: %s - Denetçi: %s\n", errMsg, inspector.Name)
				processLog["status"] = "error"
				processLog["error"] = errMsg
				processLogs = append(processLogs, processLog)
				continue
			}

			// Değişiklik var mı kontrol et
			needsUpdate := false

			// Sayısal alanların listesi
			numericFields := map[string]bool{
				"RegisterNo":  true,
				"TcNo":        true,
				"YDSID":       true,
				"CertNo":      true,
				"CompanyCode": true,
			}

			// Değerleri karşılaştır ve farklılıkları logla
			for key, newValue := range engineerData {
				if currentValue, exists := detailResult.Engineer[0][key]; exists {
					// Nil değerleri kontrol et
					if newValue == nil && currentValue == nil {
						continue
					}

					// Sayısal alan kontrolü
					if numericFields[key] {
						// Sayısal değerlere çevir
						var newFloat, currentFloat float64
						switch v := newValue.(type) {
						case float64:
							newFloat = v
						case int:
							newFloat = float64(v)
						case string:
							if f, err := strconv.ParseFloat(v, 64); err == nil {
								newFloat = f
							}
						}

						switch v := currentValue.(type) {
						case float64:
							currentFloat = v
						case int:
							currentFloat = float64(v)
						case string:
							if f, err := strconv.ParseFloat(v, 64); err == nil {
								currentFloat = f
							}
						}

						if math.Abs(newFloat-currentFloat) > 0.001 {
							needsUpdate = true
							log.Printf("Değişiklik tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentValue, newValue)
						}
					} else {
						// String karşılaştırması
						newStr := fmt.Sprintf("%v", newValue)
						currentStr := fmt.Sprintf("%v", currentValue)
						if strings.TrimSpace(newStr) != strings.TrimSpace(currentStr) {
							needsUpdate = true
							log.Printf("Değişiklik tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentValue, newValue)
						}
					}
				} else {
					// Eğer alan mevcut değilse ve yeni değer boş değilse güncelleme gerekir
					if newValue != nil && newValue != "" {
						needsUpdate = true
						log.Printf("Yeni alan eklendi - Alan: %s, Değer: %v", key, newValue)
					}
				}
			}

			if !needsUpdate {
				log.Printf("Denetçi verileri güncel, güncelleme yapılmayacak: %s", inspector.Name)
				processLog["status"] = "skipped"
				processLog["message"] = "Denetçi verileri güncel"
				processLogs = append(processLogs, processLog)
				continue
			}

			log.Printf("Değişiklik tespit edildi, denetçi güncelleniyor...")

			// Değişiklik varsa güncelle
			updateMutation := `
			mutation UpdateEngineer($ydsid: Int!, $input: CompanyEngineerInput!) {
				updateEngineer(YDSID: $ydsid, input: $input) {
					id
					Name
					YDSID
				}
			}
			`

			var updateResult struct {
				UpdateEngineer struct {
					ID    int    `json:"id"`
					Name  string `json:"Name"`
					YDSID int    `json:"YDSID"`
				} `json:"updateEngineer"`
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
			continue
		}

		// Denetçi yoksa yeni kayıt oluştur
		createMutation := `
		mutation CreateEngineer($input: CompanyEngineerInput!) {
			createEngineer(input: $input) {
				id
				Name
				YDSID
			}
		}
		`

		var createResult struct {
			CreateEngineer struct {
				ID    int    `json:"id"`
				Name  string `json:"Name"`
				YDSID int    `json:"YDSID"`
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
	}

	log.Printf("İşlem tamamlandı. Toplam: %d, Başarılı: %d, Atlanan: %d\n",
		len(simplifiedData),
		successCount,
		len(simplifiedData)-successCount)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d adet denetçiden %d tanesi başarıyla eklendi/güncellendi, %d tanesi zaten güncel",
			len(simplifiedData),
			successCount,
			len(simplifiedData)-successCount),
		"totalCount":   len(simplifiedData),
		"successCount": successCount,
		"skippedCount": len(simplifiedData) - successCount,
		"logs":         processLogs,
	})
}
