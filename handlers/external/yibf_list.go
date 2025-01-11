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

// Inspector alanlarını kontrol eden yardımcı fonksiyon
func isInspectorField(field string) bool {
	inspectorFields := []string{
		"Inspector",
		"Static",
		"Architect",
		"Mechanic",
		"Electric",
		"Controller",
		"MechanicController",
		"ElectricController",
	}
	for _, f := range inspectorFields {
		if field == f {
			return true
		}
	}
	return false
}

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

	log.Printf("YDK Base URL: %s", svc.BaseURL)

	// Test için 4 kayıt al
	requestBody := map[string]interface{}{
		"requireTotalCount": true,
		"searchOperation":   "contains",
		"searchValue":       nil,
		"skip":              20,
		"take":              5, // Test için 4 kayıt
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

	log.Printf("İstek gönderiliyor: %s", svc.BaseURL+service.ENDPOINT_YIBF_ALL)

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
				CompanyCode
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
					Name
				}
				Static {
					id
					YDSID
					Name
				}
				Architect {
					id
					YDSID
					Name
				}
				Mechanic {
					id
					YDSID
					Name
				}
				Electric {
					id
					YDSID
					Name
				}
				Controller {
					id
					YDSID
					Name
				}
				MechanicController {
					id
					YDSID
					Name
				}
				ElectricController {
					id
					YDSID
					Name
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

		detailReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", companyToken.Token))
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
			"YibfNo":         detail.ID,
			"CompanyCode":    detail.YDK.FileNumber,
			"Title":          detail.Title,
			"Administration": detail.Administration.Name,
			"State":          detail.State.Name,
			"Island": func() string {
				if detail.Island == "" || detail.Island == "Mevcut Değil" {
					return "-"
				}
				return detail.Island
			}(),
			"Parcel": func() string {
				if detail.Parcel == "" || detail.Parcel == "Mevcut Değil" {
					return "-"
				}
				return detail.Parcel
			}(),
			"Sheet": func() string {
				if detail.Sheet == "" || detail.Sheet == "Mevcut Değil" {
					return "-"
				}
				return detail.Sheet
			}(),
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
			"Coordinates":      service.CoordinatesToString(detail.Position.Coordinates),
			"UploadedFile":     detail.UploadedFile,
			"IndustryArea":     detail.YibfStructure.IndustryArea,
			"ClusterStructure": detail.ClusterStructure,
			"IsLicenseExpired": detail.IsLicenseExpired,
			"IsCompleted":      detail.IsCompleted,
		}

		// Inspector verilerini al
		inspectorURL := fmt.Sprintf("%s%s", svc.BaseURL, service.ENDPOINT_YIBF_DEPARTMENT_EMPLOYEE)
		log.Printf("Inspector URL: %s", inspectorURL)

		// Inspector request body hazırla
		inspectorRequestBody := map[string]interface{}{
			"filter": [][]interface{}{
				{"yibfId", "=", item.ID},
			},
		}

		inspectorJsonBody, err := json.Marshal(inspectorRequestBody)
		if err != nil {
			log.Printf("ID %d için inspector request body oluşturma hatası: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		inspectorReq, err := http.NewRequest("POST", inspectorURL, bytes.NewBuffer(inspectorJsonBody))
		if err != nil {
			log.Printf("ID %d için inspector request oluşturma hatası: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		inspectorReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", companyToken.Token))
		inspectorReq.Header.Set("Content-Type", "application/json")

		inspectorResp, err := svc.Client.Do(inspectorReq)
		if err != nil {
			log.Printf("ID %d için inspector isteği başarısız: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		inspectorBody, err := io.ReadAll(inspectorResp.Body)
		inspectorResp.Body.Close()
		if err != nil {
			log.Printf("ID %d için inspector yanıtı okuma hatası: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		var inspectorResponse service.YIBFInspectorResponse
		if err := json.Unmarshal(inspectorBody, &inspectorResponse); err != nil {
			log.Printf("ID %d için inspector parse hatası: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		// Inspector verilerini job input'a ekle
		for _, inspector := range inspectorResponse.Items {
			switch inspector.TaskId {
			case 1:
				jobData["Inspector"] = inspector.UserId
			case 2:
				switch inspector.TitleId {
				case 4:
					jobData["Static"] = inspector.UserId
				case 3:
					jobData["Architect"] = inspector.UserId
				case 2:
					jobData["Mechanic"] = inspector.UserId
				case 1:
					jobData["Electric"] = inspector.UserId
				}
			case 14:
				switch inspector.TitleId {
				case 3, 4:
					jobData["Controller"] = inspector.UserId
				case 2:
					jobData["MechanicController"] = inspector.UserId
				case 1:
					jobData["ElectricController"] = inspector.UserId
				}
			}
		}

		// Tarihleri sadece geçerli değer varsa ekle
		if detail.ContractDate > 0 {
			jobData["ContractDate"] = service.SafeUnixToDate(detail.ContractDate)
		}
		if detail.LicenseDate > 0 {
			jobData["LicenseDate"] = service.SafeUnixToDate(detail.LicenseDate)
		}
		if detail.CompleteDate > 0 {
			jobData["CompletionDate"] = service.SafeUnixToDate(detail.CompleteDate)
		}

		// Owner verilerini hazırla
		if detail.YibfOwner.Person.ID > 0 {
			ownerData = map[string]interface{}{
				"YibfNo":      detail.ID,
				"YDSID":       detail.YibfOwner.Person.ID,
				"Name":        detail.YibfOwner.Person.FullName,
				"TcNo":        service.SafeStringToInt(detail.YibfOwner.Person.IdentityNumber),
				"Address":     detail.YibfOwner.Person.LastAddress,
				"Phone":       detail.YibfOwner.Person.LastPhoneNumber,
				"TaxAdmin":    detail.YibfOwner.Person.TaxAdministration,
				"TaxNo":       service.SafeStringToInt(detail.YibfOwner.Person.TaxAdministrationCode),
				"Shareholder": detail.YibfOwner.ExistsShareholder,
			}
		}

		// Contractor verilerini hazırla
		if detail.LatestYibfYambis.Yambis.ID > 0 {
			contractorData = map[string]interface{}{
				"YibfNo":      detail.ID,
				"YDSID":       detail.LatestYibfYambis.Yambis.ID,
				"Name":        detail.LatestYibfYambis.Yambis.AdSoyadUnvan,
				"TaxNo":       service.SafeStringToInt(detail.LatestYibfYambis.Yambis.VergiKimlikNo),
				"Phone":       detail.LatestYibfYambis.Yambis.Telefon,
				"MobilePhone": detail.LatestYibfYambis.Yambis.CepTelefon,
				"Email":       detail.LatestYibfYambis.Yambis.Eposta,
				"Address":     detail.LatestYibfYambis.Yambis.Adres,
				"RegisterNo":  service.SafeStringToInt(detail.LatestYibfYambis.Yambis.Ybno),
				"PersonType":  detail.LatestYibfYambis.Yambis.KisiTuru,
				"TcNo":        service.SafeStringToInt(detail.LatestYibfYambis.Yambis.TcNo),
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
				"TcNo":             service.SafeStringToInt(detail.LatestYibfSiteSupervisor.Application.User.Person.IdentityNumber),
				"Position":         detail.LatestYibfSiteSupervisor.Application.Tasks.Name,
				"Career":           detail.LatestYibfSiteSupervisor.Application.Title.Name,
				"RegisterNo":       service.SafeStringToInt(detail.LatestYibfSiteSupervisor.Application.OccupationalRegistrationNumber),
				"SocialSecurityNo": service.SafeStringToInt(detail.LatestYibfSiteSupervisor.Application.SocialSecurityNo),
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

		authorReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", companyToken.Token))
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

					// YibfNo ve CompanyCode için özel karşılaştırma
					if key == "YibfNo" || key == "CompanyCode" {
						// Float to int conversion
						var currentIDInt int
						switch v := currentValue.(type) {
						case float64:
							currentIDInt = int(v)
						case int:
							currentIDInt = v
						}

						// Yeni değeri int'e çevir
						var newIDInt int
						switch v := newValue.(type) {
						case float64:
							newIDInt = int(v)
						case int:
							newIDInt = v
						}

						if currentIDInt == newIDInt {
							continue // Değerler eşleşiyorsa değişiklik yok
						}
						log.Printf("%s değişikliği tespit edildi - Eski: %v, Yeni: %v", key, currentValue, newValue)
						needsUpdate = true
						continue
					}

					// Inspector alanlarını özel olarak karşılaştır
					if isInspectorField(key) {
						// Mevcut değer bir map ise (GraphQL'den gelen nested yapı)
						if currentMap, ok := currentValue.(map[string]interface{}); ok {
							// YDSID'yi kontrol et
							if currentYDSID, exists := currentMap["YDSID"]; exists {
								// Float to int conversion
								var currentIDInt int
								switch v := currentYDSID.(type) {
								case float64:
									currentIDInt = int(v)
								case int:
									currentIDInt = v
								}

								// Yeni değeri int'e çevir
								var newIDInt int
								switch v := newValue.(type) {
								case float64:
									newIDInt = int(v)
								case int:
									newIDInt = v
								}

								if currentIDInt == newIDInt {
									continue // YDSID'ler eşleşiyorsa değişiklik yok
								}
							}
						}
						needsUpdate = true
						log.Printf("Inspector değişikliği tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentValue, newValue)
						continue
					}

					// Diğer alanlar için normal karşılaştırma
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
			var mutationParts []string
			updateVariables := map[string]interface{}{
				"yibfNo":   item.ID,
				"jobInput": jobData,
			}

			// Job mutation'ı her zaman ekle
			mutationParts = append(mutationParts, `
				updateJob(yibfNo: $yibfNo, input: $jobInput) {
					id
					YibfNo
				}`)

			// Owner verisi varsa ekle
			if ownerData != nil {
				mutationParts = append(mutationParts, `
				updateOwner(yibfNo: $yibfNo, input: $ownerInput) {
					Name
					YDSID
				}`)
				updateVariables["ownerInput"] = ownerData
			}

			// Contractor verisi varsa ekle
			if contractorData != nil {
				mutationParts = append(mutationParts, `
				updateContractor(yibfNo: $yibfNo, input: $contractorInput) {
					Name
					YDSID
				}`)
				updateVariables["contractorInput"] = contractorData
			}

			// Supervisor verisi varsa ekle
			if supervisorData != nil {
				mutationParts = append(mutationParts, `
				updateSupervisor(yibfNo: $yibfNo, input: $supervisorInput) {
					Name
					YDSID
				}`)
				updateVariables["supervisorInput"] = supervisorData
			}

			// Author verisi varsa ekle
			if authorData != nil {
				mutationParts = append(mutationParts, `
				updateAuthor(yibfNo: $yibfNo, input: $authorInput) {
					Static
					Mechanic
					Electric
					Architect
					GeotechnicalEngineer
					GeotechnicalGeologist
					GeotechnicalGeophysicist
				}`)
				updateVariables["authorInput"] = authorData
			}

			// Mutation parametrelerini oluştur
			var paramParts []string
			paramParts = append(paramParts, "$yibfNo: Int!", "$jobInput: JobInput!")
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
			updateMutation := fmt.Sprintf(`
			mutation UpdateJob(%s) {
				%s
			}
			`, strings.Join(paramParts, ", "), strings.Join(mutationParts, "\n"))

			var updateResult struct {
				UpdateJob struct {
					ID     int `json:"id"`
					YibfNo int `json:"YibfNo"`
				} `json:"updateJob"`
				UpdateOwner *struct {
					Name  string `json:"Name"`
					YDSID int    `json:"YDSID"`
				} `json:"updateOwner,omitempty"`
				UpdateContractor *struct {
					Name  string `json:"Name"`
					YDSID int    `json:"YDSID"`
				} `json:"updateContractor,omitempty"`
				UpdateSupervisor *struct {
					Name  string `json:"Name"`
					YDSID int    `json:"YDSID"`
				} `json:"updateSupervisor,omitempty"`
				UpdateAuthor *struct {
					Static                   string `json:"Static"`
					Mechanic                 string `json:"Mechanic"`
					Electric                 string `json:"Electric"`
					Architect                string `json:"Architect"`
					GeotechnicalEngineer     string `json:"GeotechnicalEngineer"`
					GeotechnicalGeologist    string `json:"GeotechnicalGeologist"`
					GeotechnicalGeophysicist string `json:"GeotechnicalGeophysicist"`
				} `json:"updateAuthor,omitempty"`
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
			// variablesJSON, _ := json.MarshalIndent(createVariables, "", "  ")
			// log.Printf("GraphQL Variables: %s", string(variablesJSON))

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
