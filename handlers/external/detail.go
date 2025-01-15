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
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/polatbilal/gqlgen-ent/handlers/client"
	"github.com/polatbilal/gqlgen-ent/handlers/service"
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

type YibfDetailRequest struct {
	YibfNumbers []int `json:"yibfNumbers"`
	CompanyCode int   `json:"companyCode"`
}

func YibfDetail(c *gin.Context, yibfNumbers []int, companyCodeStr string) {
	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT Token gerekli"})
		return
	}

	if len(yibfNumbers) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "YİBF numaraları gerekli"})
		return
	}

	companyCode := service.SafeStringToInt(companyCodeStr)
	if companyCode == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CompanyCode gerekli"})
		return
	}

	companyToken, err := service.GetCompanyTokenFromDB(c.Request.Context(), companyCode)
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

	var processedIDs []int
	var failedIDs []int
	var results []map[string]interface{}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	graphqlClient := client.GraphQLClient{
		URL: fmt.Sprintf("%s://%s/graphql", scheme, c.Request.Host),
	}

	for _, yibfID := range yibfNumbers {
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
		}`

		checkVariables := map[string]interface{}{
			"yibfNo": yibfID,
		}

		var checkResult struct {
			Job map[string]interface{} `json:"job"`
		}

		err = graphqlClient.Execute(checkQuery, checkVariables, jwtToken, &checkResult)
		if err != nil && !strings.Contains(err.Error(), "iş bulunamadı veya bu işe erişim yetkiniz yok") {
			log.Printf("ID %d için iş kontrolü sırasında hata oluştu: %v", yibfID, err)
			failedIDs = append(failedIDs, yibfID)
			continue
		}

		// YDK'dan detay bilgilerini çek
		detailURL := fmt.Sprintf("%s%s%d", svc.BaseURL, service.ENDPOINT_YIBF_BY_ID, yibfID)
		log.Printf("Detay URL: %s", detailURL)

		detailReq, err := http.NewRequest("GET", detailURL, nil)
		if err != nil {
			errMsg := fmt.Sprintf("YİBF %d için detay request oluşturma hatası: %v", yibfID, err)
			logError(errMsg)
			failedIDs = append(failedIDs, yibfID)
			continue
		}

		detailReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", companyToken.Token))
		detailResp, err := svc.Client.Do(detailReq)
		if err != nil {
			errMsg := fmt.Sprintf("YİBF %d için detay isteği başarısız: %v", yibfID, err)
			logError(errMsg)
			failedIDs = append(failedIDs, yibfID)
			continue
		}
		defer detailResp.Body.Close()

		detailBody, err := io.ReadAll(detailResp.Body)
		if err != nil {
			errMsg := fmt.Sprintf("YİBF %d için detay yanıtı okuma hatası: %v", yibfID, err)
			logError(errMsg)
			failedIDs = append(failedIDs, yibfID)
			continue
		}

		var rawDetail map[string]interface{}
		if err := json.Unmarshal(detailBody, &rawDetail); err != nil {
			errMsg := fmt.Sprintf("YİBF %d için raw JSON parse hatası: %v", yibfID, err)
			logError(errMsg)
			failedIDs = append(failedIDs, yibfID)
			continue
		}

		// Position ve coordinates kontrolü
		if position, ok := rawDetail["position"].(map[string]interface{}); ok {
			if coords, exists := position["coordinates"]; exists {
				var firstCoords []float64
				switch v := coords.(type) {
				case []interface{}:
					if len(v) >= 2 {
						if f1, ok := v[0].(float64); ok {
							if f2, ok := v[1].(float64); ok {
								firstCoords = []float64{f1, f2}
							}
						}
					} else if len(v) > 0 {
						if innerArray, ok := v[0].([]interface{}); ok {
							if len(innerArray) > 0 {
								if deepArray, ok := innerArray[0].([]interface{}); ok {
									if len(deepArray) >= 2 {
										if f1, ok := deepArray[0].(float64); ok {
											if f2, ok := deepArray[1].(float64); ok {
												firstCoords = []float64{f1, f2}
											}
										}
									}
								}
							}
						}
					}
				}
				if len(firstCoords) == 2 {
					position["coordinates"] = firstCoords
				} else {
					position["coordinates"] = []float64{}
				}
			}
		}

		var detail service.YIBFResponse
		updatedJSON, err := json.Marshal(rawDetail)
		if err != nil {
			errMsg := fmt.Sprintf("YİBF %d için JSON marshal hatası: %v", yibfID, err)
			logError(errMsg)
			failedIDs = append(failedIDs, yibfID)
			continue
		}

		if err := json.Unmarshal(updatedJSON, &detail); err != nil {
			errMsg := fmt.Sprintf("YİBF %d için detay parse hatası: %v", yibfID, err)
			logError(errMsg)
			failedIDs = append(failedIDs, yibfID)
			continue
		}

		result := map[string]interface{}{
			"yibfNo": yibfID,
			"job": map[string]interface{}{
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
				"BKSReferenceNo":   service.SafeStringToInt(detail.ReferenceNumber),
				"Coordinates":      service.CoordinatesToString(detail.Position.Coordinates),
				"UploadedFile":     detail.UploadedFile,
				"IndustryArea":     detail.YibfStructure.IndustryArea,
				"ClusterStructure": detail.ClusterStructure,
				"IsLicenseExpired": detail.IsLicenseExpired,
				"IsCompleted":      detail.IsCompleted,
			},
		}

		// Tarihleri ekle
		if detail.ContractDate > 0 {
			result["job"].(map[string]interface{})["ContractDate"] = service.SafeUnixToDate(detail.ContractDate)
		}
		if detail.LicenseDate > 0 {
			result["job"].(map[string]interface{})["LicenseDate"] = service.SafeUnixToDate(detail.LicenseDate)
		}
		if detail.CompleteDate > 0 {
			result["job"].(map[string]interface{})["CompletionDate"] = service.SafeUnixToDate(detail.CompleteDate)
		}

		// Owner verilerini ekle
		if detail.YibfOwner.Person.ID > 0 {
			result["owner"] = map[string]interface{}{
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

		// Contractor verilerini ekle
		if detail.LatestYibfYambis.Yambis.ID > 0 {
			result["contractor"] = map[string]interface{}{
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

		// Supervisor verilerini ekle
		if detail.LatestYibfSiteSupervisor.Application.User.ID > 0 {
			result["supervisor"] = map[string]interface{}{
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

		// Author verilerini al
		authorURL := fmt.Sprintf("%s%s", svc.BaseURL, service.ENDPOINT_YIBF_PROJECT_OWNER)
		authorRequestBody := map[string]interface{}{
			"filter": [][]interface{}{
				{"yibfId", "=", yibfID},
			},
		}

		authorJsonBody, err := json.Marshal(authorRequestBody)
		if err != nil {
			errMsg := fmt.Sprintf("YİBF %d için author request body oluşturma hatası: %v", yibfID, err)
			logError(errMsg)
			continue
		}

		authorReq, err := http.NewRequest("POST", authorURL, bytes.NewBuffer(authorJsonBody))
		if err != nil {
			errMsg := fmt.Sprintf("YİBF %d için author request oluşturma hatası: %v", yibfID, err)
			logError(errMsg)
			continue
		}

		authorReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", companyToken.Token))
		authorReq.Header.Set("Content-Type", "application/json")

		authorResp, err := svc.Client.Do(authorReq)
		if err != nil {
			errMsg := fmt.Sprintf("YİBF %d için author isteği başarısız: %v", yibfID, err)
			logError(errMsg)
			continue
		}

		authorBody, err := io.ReadAll(authorResp.Body)
		authorResp.Body.Close()
		if err != nil {
			errMsg := fmt.Sprintf("YİBF %d için author yanıtı okuma hatası: %v", yibfID, err)
			logError(errMsg)
			continue
		}

		var authorResponse service.YIBFAuthorResponse
		if err := json.Unmarshal(authorBody, &authorResponse); err != nil {
			errMsg := fmt.Sprintf("YİBF %d için author parse hatası: %v", yibfID, err)
			logError(errMsg)
			continue
		}

		// Author verilerini ekle
		if len(authorResponse.Items) > 0 {
			authorData := make(map[string]interface{})
			authorData["YibfNo"] = yibfID

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
			result["author"] = authorData
		}

		results = append(results, result)

		// Verileri hazırla
		jobData := result["job"].(map[string]interface{})
		ownerData := result["owner"].(map[string]interface{})
		contractorData := result["contractor"].(map[string]interface{})
		supervisorData := result["supervisor"].(map[string]interface{})
		authorData := result["author"].(map[string]interface{})

		// Job kontrolü ve güncellemesi
		if checkResult.Job != nil {
			// Kayıt var, değişiklik kontrolü yap
			needsUpdate := false

			// Job verilerini karşılaştır
			for key, newValue := range jobData {
				if currentValue, exists := checkResult.Job[key]; exists {
					if newValue == nil || currentValue == nil {
						continue
					}

					if key == "YibfNo" || key == "CompanyCode" {
						var currentIDInt int
						switch v := currentValue.(type) {
						case float64:
							currentIDInt = int(v)
						case int:
							currentIDInt = v
						}

						var newIDInt int
						switch v := newValue.(type) {
						case float64:
							newIDInt = int(v)
						case int:
							newIDInt = v
						}

						if currentIDInt == newIDInt {
							continue
						}
						log.Printf("%s değişikliği tespit edildi - Eski: %v, Yeni: %v", key, currentValue, newValue)
						needsUpdate = true
						continue
					}

					if isInspectorField(key) {
						if currentMap, ok := currentValue.(map[string]interface{}); ok {
							if currentYDSID, exists := currentMap["YDSID"]; exists {
								var currentIDInt int
								switch v := currentYDSID.(type) {
								case float64:
									currentIDInt = int(v)
								case int:
									currentIDInt = v
								}

								var newIDInt int
								switch v := newValue.(type) {
								case float64:
									newIDInt = int(v)
								case int:
									newIDInt = v
								}

								if currentIDInt == newIDInt {
									continue
								}
							}
						}
						needsUpdate = true
						log.Printf("Inspector değişikliği tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentValue, newValue)
						continue
					}

					if strings.Contains(key, "Date") {
						currentStr := fmt.Sprintf("%v", currentValue)
						newStr := fmt.Sprintf("%v", newValue)
						if service.CompareDates(currentStr, newStr) {
							continue
						}
					}

					if key == "Coordinates" {
						currentStr := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%v", currentValue), " ", ""), ",", ".")
						newStr := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%v", newValue), " ", ""), ",", ".")
						if currentStr == newStr {
							continue
						}
					}

					newStr := fmt.Sprintf("%v", newValue)
					currentStr := fmt.Sprintf("%v", currentValue)

					if strings.TrimSpace(newStr) != strings.TrimSpace(currentStr) {
						log.Printf("Değişiklik tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentValue, newValue)
						needsUpdate = true
					}
				} else {
					if newValue != nil && newValue != "" {
						log.Printf("Yeni alan eklendi - Alan: %s, Değer: %v", key, newValue)
						needsUpdate = true
					}
				}
			}

			if !needsUpdate {
				log.Printf("YİBF verileri güncel, güncelleme yapılmayacak: YibfNo %d", yibfID)
				processedIDs = append(processedIDs, yibfID)
				continue
			}

			// Job güncelleme
			jobMutation := `
			mutation UpdateJob($yibfNo: Int!, $jobInput: JobInput!) {
				updateJob(yibfNo: $yibfNo, input: $jobInput) {
					id
					YibfNo
					CompanyCode
					Title
					Administration
					State
				}
			}`

			jobVariables := map[string]interface{}{
				"yibfNo":   yibfID,
				"jobInput": jobData,
			}

			var jobResult struct {
				UpdateJob struct {
					ID          int    `json:"id"`
					YibfNo      int    `json:"YibfNo"`
					CompanyCode int    `json:"CompanyCode"`
					Title       string `json:"Title"`
				} `json:"updateJob"`
			}

			if err := graphqlClient.Execute(jobMutation, jobVariables, jwtToken, &jobResult); err != nil {
				log.Printf("ID %d için job güncelleme hatası: %v", yibfID, err)
				failedIDs = append(failedIDs, yibfID)
				continue
			}

			// Owner güncelleme
			if ownerData != nil {
				ownerMutation := `
				mutation UpdateOwner($yibfNo: Int!, $ownerInput: JobOwnerInput!) {
					updateOwner(yibfNo: $yibfNo, input: $ownerInput) {
						id
						Name
						YDSID
					}
				}`

				ownerVariables := map[string]interface{}{
					"yibfNo":     yibfID,
					"ownerInput": ownerData,
				}

				var ownerResult struct {
					UpdateOwner struct {
						ID    int    `json:"id"`
						Name  string `json:"Name"`
						YDSID int    `json:"YDSID"`
					} `json:"updateOwner"`
				}

				if err := graphqlClient.Execute(ownerMutation, ownerVariables, jwtToken, &ownerResult); err != nil {
					log.Printf("ID %d için owner güncelleme hatası: %v", yibfID, err)
				}
			}

			// Contractor güncelleme
			if contractorData != nil {
				contractorMutation := `
				mutation UpdateContractor($yibfNo: Int!, $contractorInput: JobContractorInput!) {
					updateContractor(yibfNo: $yibfNo, input: $contractorInput) {
						id
						Name
						YDSID
					}
				}`

				contractorVariables := map[string]interface{}{
					"yibfNo":          yibfID,
					"contractorInput": contractorData,
				}

				var contractorResult struct {
					UpdateContractor struct {
						ID    int    `json:"id"`
						Name  string `json:"Name"`
						YDSID int    `json:"YDSID"`
					} `json:"updateContractor"`
				}

				if err := graphqlClient.Execute(contractorMutation, contractorVariables, jwtToken, &contractorResult); err != nil {
					log.Printf("ID %d için contractor güncelleme hatası: %v", yibfID, err)
				}
			}

			// Supervisor güncelleme
			if supervisorData != nil {
				supervisorMutation := `
				mutation UpdateSupervisor($yibfNo: Int!, $supervisorInput: JobSupervisorInput!) {
					updateSupervisor(yibfNo: $yibfNo, input: $supervisorInput) {
						id
						Name
						YDSID
					}
				}`

				supervisorVariables := map[string]interface{}{
					"yibfNo":          yibfID,
					"supervisorInput": supervisorData,
				}

				var supervisorResult struct {
					UpdateSupervisor struct {
						ID    int    `json:"id"`
						Name  string `json:"Name"`
						YDSID int    `json:"YDSID"`
					} `json:"updateSupervisor"`
				}

				if err := graphqlClient.Execute(supervisorMutation, supervisorVariables, jwtToken, &supervisorResult); err != nil {
					log.Printf("ID %d için supervisor güncelleme hatası: %v", yibfID, err)
				}
			}

			// Author güncelleme
			if authorData != nil {
				authorMutation := `
				mutation UpdateAuthor($yibfNo: Int!, $authorInput: JobAuthorInput!) {
					updateAuthor(yibfNo: $yibfNo, input: $authorInput) {
						id
						Static
						Mechanic
						Electric
						Architect
						GeotechnicalEngineer
						GeotechnicalGeologist
						GeotechnicalGeophysicist
					}
				}`

				authorVariables := map[string]interface{}{
					"yibfNo":      yibfID,
					"authorInput": authorData,
				}

				var authorResult struct {
					UpdateAuthor struct {
						ID                       int    `json:"id"`
						Static                   string `json:"Static"`
						Mechanic                 string `json:"Mechanic"`
						Electric                 string `json:"Electric"`
						Architect                string `json:"Architect"`
						GeotechnicalEngineer     string `json:"GeotechnicalEngineer"`
						GeotechnicalGeologist    string `json:"GeotechnicalGeologist"`
						GeotechnicalGeophysicist string `json:"GeotechnicalGeophysicist"`
					} `json:"updateAuthor"`
				}

				if err := graphqlClient.Execute(authorMutation, authorVariables, jwtToken, &authorResult); err != nil {
					log.Printf("ID %d için author güncelleme hatası: %v", yibfID, err)
				}
			}

			log.Printf("ID %d başarıyla güncellendi", yibfID)
			processedIDs = append(processedIDs, yibfID)
		} else {
			// Kayıt yok, yeni kayıt oluştur
			// Job oluşturma
			jobMutation := `
			mutation CreateJob($jobInput: JobInput!) {
				createJob(input: $jobInput) {
					id
					YibfNo
					CompanyCode
					Title
				}
			}`

			jobVariables := map[string]interface{}{
				"jobInput": jobData,
			}

			var jobResult struct {
				CreateJob struct {
					ID          int    `json:"id"`
					YibfNo      int    `json:"YibfNo"`
					CompanyCode int    `json:"CompanyCode"`
					Title       string `json:"Title"`
				} `json:"createJob"`
			}

			if err := graphqlClient.Execute(jobMutation, jobVariables, jwtToken, &jobResult); err != nil {
				errMsg := fmt.Sprintf("ID %d için job oluşturma hatası: %v", yibfID, err)
				log.Printf("%s", errMsg)
				logError(errMsg)
				failedIDs = append(failedIDs, yibfID)
				continue
			}

			// Owner oluşturma
			if ownerData != nil {
				ownerMutation := `
				mutation CreateOwner($ownerInput: JobOwnerInput!) {
					createOwner(input: $ownerInput) {
						id
						Name
						YDSID
					}
				}`

				ownerVariables := map[string]interface{}{
					"ownerInput": ownerData,
				}

				var ownerResult struct {
					CreateOwner struct {
						ID    int    `json:"id"`
						Name  string `json:"Name"`
						YDSID int    `json:"YDSID"`
					} `json:"createOwner"`
				}

				if err := graphqlClient.Execute(ownerMutation, ownerVariables, jwtToken, &ownerResult); err != nil {
					log.Printf("ID %d için owner oluşturma hatası: %v", yibfID, err)
				}
			}

			// Contractor oluşturma
			if contractorData != nil {
				contractorMutation := `
				mutation CreateContractor($contractorInput: JobContractorInput!) {
					createContractor(input: $contractorInput) {
						id
						Name
						YDSID
					}
				}`

				contractorVariables := map[string]interface{}{
					"contractorInput": contractorData,
				}

				var contractorResult struct {
					CreateContractor struct {
						ID    int    `json:"id"`
						Name  string `json:"Name"`
						YDSID int    `json:"YDSID"`
					} `json:"createContractor"`
				}

				if err := graphqlClient.Execute(contractorMutation, contractorVariables, jwtToken, &contractorResult); err != nil {
					log.Printf("ID %d için contractor oluşturma hatası: %v", yibfID, err)
				}
			}

			// Supervisor oluşturma
			if supervisorData != nil {
				supervisorMutation := `
				mutation CreateSupervisor($supervisorInput: JobSupervisorInput!) {
					createSupervisor(input: $supervisorInput) {
						id
						Name
						YDSID
					}
				}`

				supervisorVariables := map[string]interface{}{
					"supervisorInput": supervisorData,
				}

				var supervisorResult struct {
					CreateSupervisor struct {
						ID    int    `json:"id"`
						Name  string `json:"Name"`
						YDSID int    `json:"YDSID"`
					} `json:"createSupervisor"`
				}

				if err := graphqlClient.Execute(supervisorMutation, supervisorVariables, jwtToken, &supervisorResult); err != nil {
					log.Printf("ID %d için supervisor oluşturma hatası: %v", yibfID, err)
				}
			}

			// Author oluşturma
			if authorData != nil {
				authorMutation := `
				mutation CreateAuthor($authorInput: JobAuthorInput!) {
					createAuthor(input: $authorInput) {
						id
						Static
						Mechanic
						Electric
						Architect
						GeotechnicalEngineer
						GeotechnicalGeologist
						GeotechnicalGeophysicist
					}
				}`

				authorVariables := map[string]interface{}{
					"authorInput": authorData,
				}

				var authorResult struct {
					CreateAuthor struct {
						ID                       int    `json:"id"`
						Static                   string `json:"Static"`
						Mechanic                 string `json:"Mechanic"`
						Electric                 string `json:"Electric"`
						Architect                string `json:"Architect"`
						GeotechnicalEngineer     string `json:"GeotechnicalEngineer"`
						GeotechnicalGeologist    string `json:"GeotechnicalGeologist"`
						GeotechnicalGeophysicist string `json:"GeotechnicalGeophysicist"`
					} `json:"createAuthor"`
				}

				if err := graphqlClient.Execute(authorMutation, authorVariables, jwtToken, &authorResult); err != nil {
					log.Printf("ID %d için author oluşturma hatası: %v", yibfID, err)
				}
			}

			log.Printf("ID %d başarıyla oluşturuldu", yibfID)
			processedIDs = append(processedIDs, yibfID)
		}
	}

	result := gin.H{
		"success": true,
		"data": map[string]interface{}{
			"processed_count": len(processedIDs),
			"processed_ids":   processedIDs,
			"failed_count":    len(failedIDs),
			"failed_ids":      failedIDs,
			"results":         results,
		},
	}

	c.Set("response", result)
	c.JSON(http.StatusOK, result)
}
