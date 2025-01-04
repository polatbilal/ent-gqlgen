package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/polatbilal/gqlgen-ent/handlers/client"
	"github.com/polatbilal/gqlgen-ent/handlers/service"

	"github.com/gin-gonic/gin"
)

// String'i int'e dönüştürür, hata durumunda 0 döner
func safeStringToInt(s string) int {
	if s == "" {
		return 0
	}
	if val, err := strconv.Atoi(s); err == nil {
		return val
	}
	return 0
}

// Unix timestamp'i tarihe dönüştürür, geçersiz timestamp için boş string döner
func safeUnixToDate(timestamp int64) string {
	if timestamp > 0 {
		// Milisaniyeyi saniyeye çevir
		seconds := timestamp / 1000
		return time.Unix(seconds, 0).Local().Format("2006-01-02")
	}
	return ""
}

// Koordinat array'ini string'e dönüştürür
func coordinatesToString(coords []float64) string {
	if len(coords) >= 2 {
		return fmt.Sprintf("%.6f, %.6f", coords[0], coords[1])
	}
	return ""
}

// Hata loglarını dosyaya yazar
func logError(message string) {
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Printf("Log dizini oluşturulamadı: %v", err)
		return
	}

	logFile := filepath.Join(logDir, fmt.Sprintf("yibf_errors_%s.log", time.Now().Format("2006-01-02")))
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Log dosyası açılamadı: %v", err)
		return
	}
	defer f.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] %s\n", timestamp, message)

	if _, err := f.WriteString(logMessage); err != nil {
		log.Printf("Log yazılamadı: %v", err)
	}
}

func YibfList(c *gin.Context) {
	// GraphQL için JWT token
	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT Token gerekli"})
		return
	}

	// YDK API için token
	ydkTokenStr := c.GetHeader("YDK-Token")
	if ydkTokenStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "YDK Token gerekli"})
		return
	}

	log.Printf("YDK Token alındı: %s", ydkTokenStr)

	// YDK token'ı JSON'dan parse et
	var ydkToken service.YDKTokenResponse
	if err := json.Unmarshal([]byte(ydkTokenStr), &ydkToken); err != nil {
		log.Printf("YDK Token parse hatası: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "YDK Token parse hatası: " + err.Error()})
		return
	}

	svc := &service.ExternalService{
		BaseURL: os.Getenv("YDK_BASE_URL"),
		Client:  &http.Client{},
	}

	log.Printf("YDK Base URL: %s", svc.BaseURL)

	// Test için 4 kayıt al
	requestBody := map[string]interface{}{
		"requireTotalCount": true,
		"searchOperation":   "contains",
		"searchValue":       nil,
		"skip":              20,
		"take":              4, // Test için 4 kayıt
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

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ydkToken.AccessToken))
	req.Header.Set("Content-Type", "application/json")

	log.Printf("İstek gönderiliyor: %s", svc.BaseURL+service.ENDPOINT_YIBF_ALL)
	log.Printf("Request Body: %s", string(jsonBody))

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

	log.Printf("Toplam kayıt sayısı: %d", response.TotalCount)
	log.Printf("Dönen kayıt sayısı: %d", len(response.Items))

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	graphqlClient := client.GraphQLClient{
		URL: fmt.Sprintf("%s://%s/graphql", scheme, c.Request.Host),
	}

	var processedIDs []int
	var failedIDs []int

	for _, item := range response.Items {
		log.Printf("ID %d işleniyor...", item.ID)

		// Önce mevcut kaydı kontrol et
		checkQuery := `
		query CheckJob($yibfNo: Int!) {
			job(yibfNo: $yibfNo) {
				id
				YibfNo
				Title
				Administration
				State
				Island
				Parcel
				Sheet
				ContractDate
				StartDate
				LicenseDate
				LicenseNo
				CompletionDate
				LandArea
				TotalArea
				ConstructionArea
				LeftArea
				YDSAddress
				Address
				BuildingClass
				BuildingType
				Level
				UnitPrice
				FloorCount
				BKSReferenceNo
				Coordinates
				FolderNo
				UploadedFile
				IndustryArea
				ClusterStructure
				IsLicenseExpired
				IsCompleted
				Note
				Inspector {
					id
					YDSID
				}
				Static {
					id
					YDSID
				}
				Architect {
					id
					YDSID
				}
				Mechanic {
					id
					YDSID
				}
				Electric {
					id
					YDSID
				}
				Controller {
					id
					YDSID
				}
				MechanicController {
					id
					YDSID
				}
				ElectricController {
					id
					YDSID
				}
			}
		}
		`

		checkVariables := map[string]interface{}{
			"yibfNo": item.ID,
		}

		var checkResult struct {
			Job map[string]interface{} `json:"job"`
		}

		err = graphqlClient.Execute(checkQuery, checkVariables, jwtToken, &checkResult)
		if err != nil && !strings.Contains(err.Error(), "not found") {
			log.Printf("ID %d için iş kontrolü sırasında hata oluştu: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		// Detay bilgilerini çek
		detailURL := fmt.Sprintf("%s%s%d", svc.BaseURL, service.ENDPOINT_YIBF_BY_ID, item.ID)
		log.Printf("Detay URL: %s", detailURL)

		detailReq, err := http.NewRequest("GET", detailURL, nil)
		if err != nil {
			log.Printf("ID %d için detay request oluşturma hatası: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		detailReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ydkToken.AccessToken))
		detailResp, err := svc.Client.Do(detailReq)
		if err != nil {
			log.Printf("ID %d için detay isteği başarısız: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		detailBody, err := io.ReadAll(detailResp.Body)
		detailResp.Body.Close()
		if err != nil {
			log.Printf("ID %d için detay yanıtı okuma hatası: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		var detail service.YIBFResponse
		if err := json.Unmarshal(detailBody, &detail); err != nil {
			log.Printf("ID %d için detay parse hatası: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		var ownerData, contractorData, supervisorData, authorData map[string]interface{}

		// Input verilerini hazırla
		jobData := map[string]interface{}{
			"YibfNo":           detail.ID,
			"CompanyCode":      detail.YDK.FileNumber,
			"Title":            detail.Title,
			"Administration":   detail.Administration.Name,
			"State":            detail.State.Name,
			"Island":           detail.Island,
			"Parcel":           detail.Parcel,
			"Sheet":            detail.Sheet,
			"LicenseNo":        detail.LicenseNumber,
			"LandArea":         detail.Megsis.Alan,
			"TotalArea":        detail.YibfStructure.TotalArea,
			"ConstructionArea": detail.YibfStructure.ConstructionArea,
			"LeftArea":         detail.YibfStructure.LeftArea,
			"YDSAddress":       detail.YibfStructure.BuildingAddress,
			"BuildingClass":    detail.YibfStructure.BuildingClass.Name,
			"BuildingType":     detail.YibfStructure.CarrierSystemType.Name,
			"Level":            detail.Level,
			"UnitPrice":        detail.YibfStructure.UnitPrice,
			"FloorCount":       detail.YibfStructure.FloorCount,
			"BKSReferenceNo":   detail.ReferenceNumber,
			"Coordinates":      coordinatesToString(detail.Position.Coordinates),
			"UploadedFile":     detail.UploadedFile,
			"IndustryArea":     detail.YibfStructure.IndustryArea,
			"ClusterStructure": detail.ClusterStructure,
			"IsLicenseExpired": detail.IsLicenseExpired,
			"IsCompleted":      detail.IsCompleted,
		}

		// Tarihleri sadece geçerli değer varsa ekle
		if detail.ContractDate > 0 {
			jobData["ContractDate"] = safeUnixToDate(detail.ContractDate)
		}
		if detail.LicenseDate > 0 {
			jobData["LicenseDate"] = safeUnixToDate(detail.LicenseDate)
		}
		if detail.CompleteDate > 0 {
			jobData["CompletionDate"] = safeUnixToDate(detail.CompleteDate)
		}

		// Owner verilerini hazırla
		if detail.YibfOwner.Person.ID > 0 {
			ownerData = map[string]interface{}{
				"YibfNo":      detail.ID,
				"YDSID":       detail.YibfOwner.Person.ID,
				"Name":        detail.YibfOwner.Person.FullName,
				"TcNo":        safeStringToInt(detail.YibfOwner.Person.IdentityNumber),
				"Address":     detail.YibfOwner.Person.LastAddress,
				"Phone":       detail.YibfOwner.Person.LastPhoneNumber,
				"TaxAdmin":    detail.YibfOwner.Person.TaxAdministration,
				"TaxNo":       safeStringToInt(detail.YibfOwner.Person.TaxAdministrationCode),
				"Shareholder": detail.YibfOwner.ExistsShareholder,
			}
		}

		// Contractor verilerini hazırla
		if detail.LatestYibfYambis.Yambis.ID > 0 {
			contractorData = map[string]interface{}{
				"YibfNo":      detail.ID,
				"YDSID":       detail.LatestYibfYambis.Yambis.ID,
				"Name":        detail.LatestYibfYambis.Yambis.AdSoyadUnvan,
				"TaxNo":       safeStringToInt(detail.LatestYibfYambis.Yambis.VergiKimlikNo),
				"Phone":       detail.LatestYibfYambis.Yambis.Telefon,
				"MobilePhone": detail.LatestYibfYambis.Yambis.CepTelefon,
				"Email":       detail.LatestYibfYambis.Yambis.Eposta,
				"Address":     detail.LatestYibfYambis.Yambis.Adres,
				"RegisterNo":  safeStringToInt(detail.LatestYibfYambis.Yambis.Ybno),
				"PersonType":  detail.LatestYibfYambis.Yambis.KisiTuru,
				"TcNo":        safeStringToInt(detail.LatestYibfYambis.Yambis.TcNo),
			}
		}

		// Supervisor verilerini hazırla
		if detail.LatestYibfSiteSupervisor.Application.User.ID > 0 {
			supervisorData = map[string]interface{}{
				"YibfNo":           detail.ID,
				"YDSID":            detail.LatestYibfSiteSupervisor.Application.User.ID,
				"Name":             detail.LatestYibfSiteSupervisor.Application.User.Person.FullName,
				"Address":          detail.LatestYibfSiteSupervisor.Application.User.Person.LastAddress,
				"Phone":            detail.LatestYibfSiteSupervisor.Application.User.Person.LastPhoneNumber,
				"Email":            detail.LatestYibfSiteSupervisor.Application.User.Person.LastEPosta,
				"TcNo":             safeStringToInt(detail.LatestYibfSiteSupervisor.Application.User.Person.IdentityNumber),
				"Position":         detail.LatestYibfSiteSupervisor.Application.Tasks.Name,
				"Career":           detail.LatestYibfSiteSupervisor.Application.Title.Name,
				"RegisterNo":       safeStringToInt(detail.LatestYibfSiteSupervisor.Application.OccupationalRegistrationNumber),
				"SocialSecurityNo": safeStringToInt(detail.LatestYibfSiteSupervisor.Application.SocialSecurityNo),
				"SchoolGraduation": detail.LatestYibfSiteSupervisor.Application.SchoolGraduation,
			}
		}

		// Author verilerini ayrı endpoint'ten al
		authorURL := fmt.Sprintf("%s%s", svc.BaseURL, service.ENDPOINT_YIBF_PROJECT_OWNER)
		log.Printf("Author URL: %s", authorURL)

		// Author request body hazırla
		authorRequestBody := map[string]interface{}{
			"filter": [][]interface{}{
				{"yibfId", "=", item.ID},
			},
		}

		authorJsonBody, err := json.Marshal(authorRequestBody)
		if err != nil {
			log.Printf("ID %d için author request body oluşturma hatası: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		authorReq, err := http.NewRequest("POST", authorURL, bytes.NewBuffer(authorJsonBody))
		if err != nil {
			log.Printf("ID %d için author request oluşturma hatası: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		authorReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ydkToken.AccessToken))
		authorReq.Header.Set("Content-Type", "application/json")

		authorResp, err := svc.Client.Do(authorReq)
		if err != nil {
			log.Printf("ID %d için author isteği başarısız: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		authorBody, err := io.ReadAll(authorResp.Body)
		authorResp.Body.Close()
		if err != nil {
			log.Printf("ID %d için author yanıtı okuma hatası: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		// Author yanıtını logla
		log.Printf("ID %d için author yanıtı: %s", item.ID, string(authorBody))
		log.Printf("Author yanıt status kodu: %d", authorResp.StatusCode)
		log.Printf("Author yanıt headers: %v", authorResp.Header)

		var authorResponse service.YIBFAuthorResponse
		if err := json.Unmarshal(authorBody, &authorResponse); err != nil {
			log.Printf("ID %d için author parse hatası: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		// Author verilerini hazırla
		if len(authorResponse.Items) > 0 {
			authorData = make(map[string]interface{})
			authorData["YibfNo"] = item.ID

			for _, author := range authorResponse.Items {
				fullName := fmt.Sprintf("%s %s", author.PersonName, author.PersonSurname)

				switch author.TaskId {
				case 4:
					authorData["Electric"] = fullName
				case 5:
					authorData["Mechanic"] = fullName
				case 6:
					authorData["Architect"] = fullName
				case 7:
					authorData["Static"] = fullName
				case 8:
					switch author.TitleId {
					case 7:
						authorData["GeotechnicalGeophysicist"] = fullName
					case 6:
						authorData["GeotechnicalGeologist"] = fullName
					case 4:
						authorData["GeotechnicalEngineer"] = fullName
					}
				}
			}
		}

		if checkResult.Job != nil {
			// Kayıt var, değişiklik kontrolü yap
			needsUpdate := false

			// Değerleri karşılaştır ve farklılıkları logla
			for key, newValue := range jobData {
				if currentValue, exists := checkResult.Job[key]; exists {
					// Nil değerleri kontrol et
					if newValue == nil || currentValue == nil {
						continue
					}

					// String'e çevirip karşılaştır
					newStr := fmt.Sprintf("%v", newValue)
					currentStr := fmt.Sprintf("%v", currentValue)

					if strings.TrimSpace(newStr) != strings.TrimSpace(currentStr) {
						log.Printf("Değişiklik tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentValue, newValue)
						needsUpdate = true
					}
				} else {
					// Eğer alan mevcut değilse ve yeni değer boş değilse güncelleme gerekir
					if newValue != nil && newValue != "" {
						log.Printf("Yeni alan eklendi - Alan: %s, Değer: %v", key, newValue)
						needsUpdate = true
					}
				}
			}

			if !needsUpdate {
				log.Printf("YİBF verileri güncel, güncelleme yapılmayacak: YibfNo %d", item.ID)
				processedIDs = append(processedIDs, item.ID)
				continue
			}

			// Değişiklik varsa güncelle
			updateMutation := `
			mutation UpdateJob($yibfNo: Int!, $input: JobInput!) {
				updateJob(YibfNo: $yibfNo, input: $input) {
					id
					YibfNo
				}
			}
			`

			updateVariables := map[string]interface{}{
				"yibfNo": item.ID,
				"input":  jobData,
			}

			var updateResult struct {
				UpdateJob struct {
					ID     int `json:"id"`
					YibfNo int `json:"YibfNo"`
				} `json:"updateJob"`
			}

			if err := graphqlClient.Execute(updateMutation, updateVariables, jwtToken, &updateResult); err != nil {
				log.Printf("ID %d için güncelleme hatası: %v", item.ID, err)
				failedIDs = append(failedIDs, item.ID)
				continue
			}

			log.Printf("ID %d başarıyla güncellendi", item.ID)
			processedIDs = append(processedIDs, item.ID)
		} else {
			// Kayıt yok, yeni kayıt oluştur
			var mutationParts []string
			createVariables := map[string]interface{}{
				"jobInput": jobData,
			}

			// Job mutation'ı her zaman ekle
			mutationParts = append(mutationParts, `
				createJob(input: $jobInput) {
					YibfNo
					Title
				}`)

			// Owner verisi varsa ekle
			if ownerData != nil {
				mutationParts = append(mutationParts, `
				createOwner(input: $ownerInput) {
					Name
					YDSID
				}`)
				createVariables["ownerInput"] = ownerData
			}

			// Contractor verisi varsa ekle
			if contractorData != nil {
				mutationParts = append(mutationParts, `
				createContractor(input: $contractorInput) {
					Name
					YDSID
				}`)
				createVariables["contractorInput"] = contractorData
			}

			// Supervisor verisi varsa ekle
			if supervisorData != nil {
				mutationParts = append(mutationParts, `
				createSupervisor(input: $supervisorInput) {
					Name
					YDSID
				}`)
				createVariables["supervisorInput"] = supervisorData
			}

			// Author verisi varsa ekle
			if authorData != nil {
				mutationParts = append(mutationParts, `
				createAuthor(input: $authorInput) {
					Static
					Mechanic
					Electric
					Architect
					GeotechnicalEngineer
					GeotechnicalGeologist
					GeotechnicalGeophysicist
				}`)
				createVariables["authorInput"] = authorData
			}

			// Mutation parametrelerini oluştur
			var paramParts []string
			paramParts = append(paramParts, "$jobInput: JobInput!")
			if ownerData != nil {
				paramParts = append(paramParts, "$ownerInput: JobOwnerInput!")
			}
			if contractorData != nil {
				paramParts = append(paramParts, "$contractorInput: JobContractorInput!")
			}
			if supervisorData != nil {
				paramParts = append(paramParts, "$supervisorInput: JobSupervisorInput!")
			}
			if authorData != nil {
				paramParts = append(paramParts, "$authorInput: JobAuthorInput!")
			}

			// Tam mutation string'ini oluştur
			createMutation := fmt.Sprintf(`
			mutation CreateJob(%s) {
				%s
			}
			`, strings.Join(paramParts, ", "), strings.Join(mutationParts, "\n"))

			// Detaylı loglama
			// log.Printf("GraphQL Mutation: %s", createMutation)
			variablesJSON, _ := json.MarshalIndent(createVariables, "", "  ")
			log.Printf("GraphQL Variables: %s", string(variablesJSON))

			var createResult struct {
				CreateJob struct {
					YibfNo int    `json:"YibfNo"`
					Title  string `json:"Title"`
				} `json:"createJob"`
				CreateOwner *struct {
					Name  string `json:"Name"`
					YDSID int    `json:"YDSID"`
				} `json:"createOwner,omitempty"`
				CreateContractor *struct {
					Name  string `json:"Name"`
					YDSID int    `json:"YDSID"`
				} `json:"createContractor,omitempty"`
				CreateSupervisor *struct {
					Name  string `json:"Name"`
					YDSID int    `json:"YDSID"`
				} `json:"createSupervisor,omitempty"`
				CreateAuthor *struct {
					Static                   string `json:"Static"`
					Mechanic                 string `json:"Mechanic"`
					Electric                 string `json:"Electric"`
					Architect                string `json:"Architect"`
					GeotechnicalEngineer     string `json:"GeotechnicalEngineer"`
					GeotechnicalGeologist    string `json:"GeotechnicalGeologist"`
					GeotechnicalGeophysicist string `json:"GeotechnicalGeophysicist"`
				} `json:"createAuthor,omitempty"`
			}

			if err := graphqlClient.Execute(createMutation, createVariables, jwtToken, &createResult); err != nil {
				errMsg := fmt.Sprintf("ID %d için oluşturma hatası: %v", item.ID, err)
				log.Printf("%s", errMsg)
				logError(errMsg)
				failedIDs = append(failedIDs, item.ID)
				continue
			}

			log.Printf("ID %d başarıyla oluşturuldu", item.ID)
			processedIDs = append(processedIDs, item.ID)
		}

		// Yapı Bilgileri loglaması
		log.Printf("Yapı Bilgileri - YibfNo: %d, İdare: %s, Pafta: %s, Ada: %s, Parsel: %s",
			detail.ID,
			detail.Administration.Name,
			detail.Sheet,
			detail.Island,
			detail.Parcel)
	}

	c.JSON(http.StatusOK, gin.H{
		"success":         true,
		"total":           response.TotalCount,
		"processed_count": len(processedIDs),
		"processed_ids":   processedIDs,
		"failed_count":    len(failedIDs),
		"failed_ids":      failedIDs,
	})
}
