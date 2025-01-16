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
	"reflect"
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
			query CheckJobWithRelations($yibfNo: Int!) {
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
				owner(yibfNo: $yibfNo) {
					id
					YDSID
					Name
					TcNo
					Address
					Phone
					TaxAdmin
					TaxNo
					Shareholder
				}
				contractor(yibfNo: $yibfNo) {
					id
					YDSID
					Name
					TaxNo
					Phone
					MobilePhone
					Email
					Address
					RegisterNo
					PersonType
					TcNo
				}
				supervisor(yibfNo: $yibfNo) {
					id
					YDSID
					Name
					Address
					Phone
					Email
					TcNo
					Position
					Career
					RegisterNo
					SocialSecurityNo
					SchoolGraduation
				}
				author(yibfNo: $yibfNo) {
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

		checkVariables := map[string]interface{}{
			"yibfNo": yibfID,
		}

		var checkResult struct {
			Job *struct {
				ID               int     `json:"id"`
				YibfNo           int     `json:"YibfNo"`
				CompanyCode      int     `json:"CompanyCode"`
				Title            string  `json:"Title"`
				Administration   string  `json:"Administration"`
				State            string  `json:"State"`
				Island           string  `json:"Island"`
				Parcel           string  `json:"Parcel"`
				Sheet            string  `json:"Sheet"`
				ContractDate     string  `json:"ContractDate"`
				StartDate        string  `json:"StartDate"`
				LicenseDate      string  `json:"LicenseDate"`
				LicenseNo        string  `json:"LicenseNo"`
				CompletionDate   string  `json:"CompletionDate"`
				LandArea         float64 `json:"LandArea"`
				TotalArea        float64 `json:"TotalArea"`
				ConstructionArea float64 `json:"ConstructionArea"`
				LeftArea         float64 `json:"LeftArea"`
				YDSAddress       string  `json:"YDSAddress"`
				Address          string  `json:"Address"`
				BuildingClass    string  `json:"BuildingClass"`
				BuildingType     string  `json:"BuildingType"`
				Level            int     `json:"Level"`
				UnitPrice        float64 `json:"UnitPrice"`
				FloorCount       int     `json:"FloorCount"`
				BKSReferenceNo   int     `json:"BKSReferenceNo"`
				Coordinates      string  `json:"Coordinates"`
				FolderNo         string  `json:"FolderNo"`
				UploadedFile     bool    `json:"UploadedFile"`
				IndustryArea     bool    `json:"IndustryArea"`
				ClusterStructure bool    `json:"ClusterStructure"`
				IsLicenseExpired bool    `json:"IsLicenseExpired"`
				IsCompleted      bool    `json:"IsCompleted"`
				Note             string  `json:"Note"`
				Inspector        struct {
					ID    int    `json:"id"`
					YDSID int    `json:"YDSID"`
					Name  string `json:"Name"`
				} `json:"Inspector"`
				Static struct {
					ID    int    `json:"id"`
					YDSID int    `json:"YDSID"`
					Name  string `json:"Name"`
				} `json:"Static"`
				Architect struct {
					ID    int    `json:"id"`
					YDSID int    `json:"YDSID"`
					Name  string `json:"Name"`
				} `json:"Architect"`
				Mechanic struct {
					ID    int    `json:"id"`
					YDSID int    `json:"YDSID"`
					Name  string `json:"Name"`
				} `json:"Mechanic"`
				Electric struct {
					ID    int    `json:"id"`
					YDSID int    `json:"YDSID"`
					Name  string `json:"Name"`
				} `json:"Electric"`
				Controller struct {
					ID    int    `json:"id"`
					YDSID int    `json:"YDSID"`
					Name  string `json:"Name"`
				} `json:"Controller"`
				MechanicController struct {
					ID    int    `json:"id"`
					YDSID int    `json:"YDSID"`
					Name  string `json:"Name"`
				} `json:"MechanicController"`
				ElectricController struct {
					ID    int    `json:"id"`
					YDSID int    `json:"YDSID"`
					Name  string `json:"Name"`
				} `json:"ElectricController"`
			} `json:"job"`
			Owner *struct {
				ID          int    `json:"id"`
				YDSID       int    `json:"YDSID"`
				Name        string `json:"Name"`
				TcNo        int    `json:"TcNo"`
				Address     string `json:"Address"`
				Phone       string `json:"Phone"`
				TaxAdmin    string `json:"TaxAdmin"`
				TaxNo       int    `json:"TaxNo"`
				Shareholder bool   `json:"Shareholder"`
			} `json:"owner"`
			Contractor *struct {
				ID          int    `json:"id"`
				YDSID       int    `json:"YDSID"`
				Name        string `json:"Name"`
				TaxNo       int    `json:"TaxNo"`
				Phone       string `json:"Phone"`
				MobilePhone string `json:"MobilePhone"`
				Email       string `json:"Email"`
				Address     string `json:"Address"`
				RegisterNo  int    `json:"RegisterNo"`
				PersonType  string `json:"PersonType"`
				TcNo        int    `json:"TcNo"`
			} `json:"contractor"`
			Supervisor *struct {
				ID               int    `json:"id"`
				YDSID            int    `json:"YDSID"`
				Name             string `json:"Name"`
				Address          string `json:"Address"`
				Phone            string `json:"Phone"`
				Email            string `json:"Email"`
				TcNo             int    `json:"TcNo"`
				Position         string `json:"Position"`
				Career           string `json:"Career"`
				RegisterNo       int    `json:"RegisterNo"`
				SocialSecurityNo int    `json:"SocialSecurityNo"`
				SchoolGraduation string `json:"SchoolGraduation"`
			} `json:"supervisor"`
			Author *struct {
				ID                       int    `json:"id"`
				Static                   string `json:"Static"`
				Mechanic                 string `json:"Mechanic"`
				Electric                 string `json:"Electric"`
				Architect                string `json:"Architect"`
				GeotechnicalEngineer     string `json:"GeotechnicalEngineer"`
				GeotechnicalGeologist    string `json:"GeotechnicalGeologist"`
				GeotechnicalGeophysicist string `json:"GeotechnicalGeophysicist"`
			} `json:"author"`
		}

		err = graphqlClient.Execute(checkQuery, checkVariables, jwtToken, &checkResult)
		var notFoundTables []string

		if err != nil {
			log.Printf("ID %d için GraphQL sorgusu:\n%s", yibfID, checkQuery)
			checkVariablesJSON, _ := json.MarshalIndent(checkVariables, "", "  ")
			log.Printf("Variables:\n%s", string(checkVariablesJSON))
			log.Printf("Ham hata mesajı: %v", err)

			errStr := err.Error()
			// Önce ham hata mesajını kontrol et
			if strings.Contains(errStr, "job_contractor not found") {
				notFoundTables = append(notFoundTables, "contractor")
			}
			if strings.Contains(errStr, "job_owner not found") {
				notFoundTables = append(notFoundTables, "owner")
			}
			if strings.Contains(errStr, "job_supervisor not found") {
				notFoundTables = append(notFoundTables, "supervisor")
			}
			if strings.Contains(errStr, "job_author not found") {
				notFoundTables = append(notFoundTables, "author")
			}
			if strings.Contains(errStr, "iş bulunamadı veya bu işe erişim yetkiniz yok") {
				notFoundTables = []string{"job", "owner", "contractor", "supervisor", "author"}
				log.Printf("İş bulunamadı, tüm tablolar için create işlemi yapılacak")
			}

			// GraphQL hata mesajını parse etmeyi dene
			var errResp struct {
				Errors []struct {
					Message string   `json:"message"`
					Path    []string `json:"path"`
				} `json:"errors"`
			}

			if strings.HasPrefix(errStr, "{") {
				if err := json.Unmarshal([]byte(errStr), &errResp); err == nil {
					for _, e := range errResp.Errors {
						log.Printf("GraphQL Hata Detayı - Path: %v, Message: %s", e.Path, e.Message)

						// Job tablosu için özel kontrol
						if len(e.Path) > 0 && e.Path[0] == "job" && strings.Contains(e.Message, "iş bulunamadı veya bu işe erişim yetkiniz yok") {
							notFoundTables = []string{"job", "owner", "contractor", "supervisor", "author"}
							log.Printf("İş bulunamadı, tüm tablolar için create işlemi yapılacak")
							break
						}

						// Diğer tablolar için not found kontrolü
						if strings.Contains(e.Message, "not found") && len(e.Path) > 0 {
							notFoundTables = append(notFoundTables, e.Path[0])
						}
					}
				}
			}

			log.Printf("Not found tabloları: %v", notFoundTables)
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
					switch author.TitleId {
					case 4:
						authorData["GeotechnicalEngineer"] = fullName
					case 6:
						authorData["GeotechnicalGeologist"] = fullName
					case 7:
						authorData["GeotechnicalGeophysicist"] = fullName
					}
				}
			}
			result["author"] = authorData
		}

		results = append(results, result)

		// Verileri hazırla ve değişiklikleri kontrol et
		var jobChanges map[string]interface{}
		var ownerChanges map[string]interface{}
		var contractorChanges map[string]interface{}
		var supervisorChanges map[string]interface{}
		var authorChanges map[string]interface{}

		// Job değişikliklerini kontrol et
		if job, ok := result["job"].(map[string]interface{}); ok && checkResult.Job != nil {
			jobChanges = make(map[string]interface{})
			jobObj := reflect.ValueOf(checkResult.Job).Elem()

			for key, newValue := range job {
				if key == "YibfNo" || key == "CompanyCode" {
					continue // Bu alanları güncellemeye gerek yok
				}

				if strings.Contains(key, "Date") {
					currentValue := jobObj.FieldByName(key).String()
					newStr := fmt.Sprintf("%v", newValue)
					if !service.CompareDates(currentValue, newStr) {
						jobChanges[key] = newValue
						log.Printf("Job tarih değişikliği tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentValue, newValue)
					}
					continue
				}

				if key == "Coordinates" {
					currentValue := jobObj.FieldByName(key).String()
					currentStr := strings.ReplaceAll(strings.ReplaceAll(currentValue, " ", ""), ",", ".")
					newStr := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%v", newValue), " ", ""), ",", ".")
					if currentStr != newStr {
						jobChanges[key] = newValue
						log.Printf("Job koordinat değişikliği tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentValue, newValue)
					}
					continue
				}

				field := jobObj.FieldByName(key)
				if !field.IsValid() {
					jobChanges[key] = newValue
					log.Printf("Job yeni alan eklendi - Alan: %s, Değer: %v", key, newValue)
					continue
				}

				newStr := fmt.Sprintf("%v", newValue)
				currentStr := fmt.Sprintf("%v", field.Interface())
				if strings.TrimSpace(newStr) != strings.TrimSpace(currentStr) {
					jobChanges[key] = newValue
					log.Printf("Job değişiklik tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentStr, newStr)
				}
			}
		}

		// Owner değişikliklerini kontrol et
		if owner, ok := result["owner"].(map[string]interface{}); ok && checkResult.Owner != nil {
			ownerChanges = make(map[string]interface{})
			ownerObj := reflect.ValueOf(checkResult.Owner).Elem()

			for key, newValue := range owner {
				if key == "YibfNo" {
					continue
				}

				field := ownerObj.FieldByName(key)
				if !field.IsValid() {
					ownerChanges[key] = newValue
					log.Printf("Owner yeni alan eklendi - Alan: %s, Değer: %v", key, newValue)
					continue
				}

				newStr := fmt.Sprintf("%v", newValue)
				currentStr := fmt.Sprintf("%v", field.Interface())
				if strings.TrimSpace(newStr) != strings.TrimSpace(currentStr) {
					ownerChanges[key] = newValue
					log.Printf("Owner değişiklik tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentStr, newStr)
				}
			}
		}

		// Contractor değişikliklerini kontrol et
		if contractor, ok := result["contractor"].(map[string]interface{}); ok && checkResult.Contractor != nil {
			contractorChanges = make(map[string]interface{})
			contractorObj := reflect.ValueOf(checkResult.Contractor).Elem()

			for key, newValue := range contractor {
				if key == "YibfNo" {
					continue
				}

				field := contractorObj.FieldByName(key)
				if !field.IsValid() {
					contractorChanges[key] = newValue
					log.Printf("Contractor yeni alan eklendi - Alan: %s, Değer: %v", key, newValue)
					continue
				}

				newStr := fmt.Sprintf("%v", newValue)
				currentStr := fmt.Sprintf("%v", field.Interface())
				if strings.TrimSpace(newStr) != strings.TrimSpace(currentStr) {
					contractorChanges[key] = newValue
					log.Printf("Contractor değişiklik tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentStr, newStr)
				}
			}
		}

		// Supervisor değişikliklerini kontrol et
		if supervisor, ok := result["supervisor"].(map[string]interface{}); ok && checkResult.Supervisor != nil {
			supervisorChanges = make(map[string]interface{})
			supervisorObj := reflect.ValueOf(checkResult.Supervisor).Elem()

			for key, newValue := range supervisor {
				if key == "YibfNo" {
					continue
				}

				field := supervisorObj.FieldByName(key)
				if !field.IsValid() {
					supervisorChanges[key] = newValue
					log.Printf("Supervisor yeni alan eklendi - Alan: %s, Değer: %v", key, newValue)
					continue
				}

				newStr := fmt.Sprintf("%v", newValue)
				currentStr := fmt.Sprintf("%v", field.Interface())
				if strings.TrimSpace(newStr) != strings.TrimSpace(currentStr) {
					supervisorChanges[key] = newValue
					log.Printf("Supervisor değişiklik tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentStr, newStr)
				}
			}
		}

		// Author değişikliklerini kontrol et
		if author, ok := result["author"].(map[string]interface{}); ok && checkResult.Author != nil {
			authorChanges = make(map[string]interface{})
			authorObj := reflect.ValueOf(checkResult.Author).Elem()

			for key, newValue := range author {
				if key == "YibfNo" {
					continue
				}

				field := authorObj.FieldByName(key)
				if !field.IsValid() {
					authorChanges[key] = newValue
					log.Printf("Author yeni alan eklendi - Alan: %s, Değer: %v", key, newValue)
					continue
				}

				newStr := fmt.Sprintf("%v", newValue)
				currentStr := fmt.Sprintf("%v", field.Interface())
				if strings.TrimSpace(newStr) != strings.TrimSpace(currentStr) {
					authorChanges[key] = newValue
					log.Printf("Author değişiklik tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentStr, newStr)
				}
			}
		}

		// Mutation parçalarını oluştur
		var mutationParts []string
		var mutationFields []string
		var hasUpdate bool

		// Job mutation'ı
		if len(jobChanges) > 0 {
			hasUpdate = true
			mutationParts = append(mutationParts, `$jobInput: JobInput!`)
			mutationFields = append(mutationFields, `
				job: updateJob(yibfNo: $yibfNo, input: $jobInput) {
					id
					YibfNo
					CompanyCode
					Title
				}`)
		}

		// Owner mutation'ı
		if len(ownerChanges) > 0 {
			hasUpdate = true
			mutationParts = append(mutationParts, `$ownerInput: JobOwnerInput!`)
			mutationFields = append(mutationFields, `
				owner: updateOwner(yibfNo: $yibfNo, input: $ownerInput) {
					id
					Name
					YDSID
				}`)
		} else if result["owner"] != nil {
			for _, table := range notFoundTables {
				if table == "owner" {
					if owner, ok := result["owner"].(map[string]interface{}); ok {
						mutationParts = append(mutationParts, `$ownerInput: JobOwnerInput!`)
						mutationFields = append(mutationFields, `
							owner: createOwner(input: $ownerInput) {
								id
								Name
								YDSID
							}`)
						ownerChanges = owner
					}
					break
				}
			}
		}

		// Contractor mutation'ı
		if len(contractorChanges) > 0 {
			hasUpdate = true
			mutationParts = append(mutationParts, `$contractorInput: JobContractorInput!`)
			mutationFields = append(mutationFields, `
				contractor: updateContractor(yibfNo: $yibfNo, input: $contractorInput) {
					id
					Name
					YDSID
				}`)
		} else if result["contractor"] != nil {
			for _, table := range notFoundTables {
				if table == "contractor" {
					if contractor, ok := result["contractor"].(map[string]interface{}); ok {
						mutationParts = append(mutationParts, `$contractorInput: JobContractorInput!`)
						mutationFields = append(mutationFields, `
							contractor: createContractor(input: $contractorInput) {
								id
								Name
								YDSID
							}`)
						contractorChanges = contractor
						break
					}
				}
			}
		}

		// Supervisor mutation'ı
		if len(supervisorChanges) > 0 {
			hasUpdate = true
			mutationParts = append(mutationParts, `$supervisorInput: JobSupervisorInput!`)
			mutationFields = append(mutationFields, `
				supervisor: updateSupervisor(yibfNo: $yibfNo, input: $supervisorInput) {
					id
					Name
					YDSID
				}`)
		} else if result["supervisor"] != nil {
			for _, table := range notFoundTables {
				if table == "supervisor" {
					if supervisor, ok := result["supervisor"].(map[string]interface{}); ok {
						mutationParts = append(mutationParts, `$supervisorInput: JobSupervisorInput!`)
						mutationFields = append(mutationFields, `
							supervisor: createSupervisor(input: $supervisorInput) {
								id
								Name
								YDSID
							}`)
						supervisorChanges = supervisor
						break
					}
				}
			}
		}

		// Author mutation'ı
		if len(authorChanges) > 0 {
			hasUpdate = true
			mutationParts = append(mutationParts, `$authorInput: JobAuthorInput!`)
			mutationFields = append(mutationFields, `
				author: updateAuthor(yibfNo: $yibfNo, input: $authorInput) {
					id
					Static
					Mechanic
					Electric
					Architect
					GeotechnicalEngineer
					GeotechnicalGeologist
					GeotechnicalGeophysicist
				}`)
		} else if result["author"] != nil {
			for _, table := range notFoundTables {
				if table == "author" {
					if author, ok := result["author"].(map[string]interface{}); ok {
						mutationParts = append(mutationParts, `$authorInput: JobAuthorInput!`)
						mutationFields = append(mutationFields, `
							author: createAuthor(input: $authorInput) {
								id
								Static
								Mechanic
								Electric
								Architect
								GeotechnicalEngineer
								GeotechnicalGeologist
								GeotechnicalGeophysicist
							}`)
						authorChanges = author
						break
					}
				}
			}
		}

		// Eğer hiç değişiklik yoksa devam et
		if len(mutationParts) == 0 {
			log.Printf("ID %d için değişiklik yok, güncelleme yapılmayacak", yibfID)
			processedIDs = append(processedIDs, yibfID)
			continue
		}

		// Mutation string'ini oluştur
		var updateMutation string
		if hasUpdate {
			updateMutation = fmt.Sprintf(`
			mutation UpsertJobWithRelations(
				$yibfNo: Int!
				%s
			) {
				%s
			}`, strings.Join(mutationParts, "\n\t\t\t\t"), strings.Join(mutationFields, "\n"))
		} else {
			// Tüm tablolarda not found hatası var mı kontrol et
			allTablesNotFound := len(notFoundTables) >= 5 && // En az 5 tablo (job, owner, contractor, supervisor, author)
				contains(notFoundTables, "job") &&
				contains(notFoundTables, "owner") &&
				contains(notFoundTables, "contractor") &&
				contains(notFoundTables, "supervisor") &&
				contains(notFoundTables, "author")

			if allTablesNotFound {
				log.Printf("Tüm tablolarda not found hatası var, create mutation'ı çalıştırılacak")
				// Create mutation'ı için tüm verileri ekle
				if job, ok := result["job"].(map[string]interface{}); ok {
					jobChanges = job
					mutationParts = append(mutationParts, `$jobInput: JobInput!`)
					mutationFields = append(mutationFields, `
						job: createJob(input: $jobInput) {
							id
							YibfNo
							CompanyCode
							Title
						}`)
				}
				if owner, ok := result["owner"].(map[string]interface{}); ok {
					ownerChanges = owner
					mutationParts = append(mutationParts, `$ownerInput: JobOwnerInput!`)
					mutationFields = append(mutationFields, `
						owner: createOwner(input: $ownerInput) {
							id
							Name
							YDSID
						}`)
				}
				if contractor, ok := result["contractor"].(map[string]interface{}); ok {
					contractorChanges = contractor
					mutationParts = append(mutationParts, `$contractorInput: JobContractorInput!`)
					mutationFields = append(mutationFields, `
						contractor: createContractor(input: $contractorInput) {
							id
							Name
							YDSID
						}`)
				}
				if supervisor, ok := result["supervisor"].(map[string]interface{}); ok {
					supervisorChanges = supervisor
					mutationParts = append(mutationParts, `$supervisorInput: JobSupervisorInput!`)
					mutationFields = append(mutationFields, `
						supervisor: createSupervisor(input: $supervisorInput) {
							id
							Name
							YDSID
						}`)
				}
				if author, ok := result["author"].(map[string]interface{}); ok {
					authorChanges = author
					mutationParts = append(mutationParts, `$authorInput: JobAuthorInput!`)
					mutationFields = append(mutationFields, `
						author: createAuthor(input: $authorInput) {
							id
							Static
							Mechanic
							Electric
							Architect
							GeotechnicalEngineer
							GeotechnicalGeologist
							GeotechnicalGeophysicist
						}`)
				}
			}

			updateMutation = fmt.Sprintf(`
			mutation CreateJobWithRelations(
				%s
			) {
				%s
			}`, strings.Join(mutationParts, "\n\t\t\t\t"), strings.Join(mutationFields, "\n"))
		}

		// Variables'ı hazırla
		variables := map[string]interface{}{
			"yibfNo": yibfID,
		}

		if len(jobChanges) > 0 {
			variables["jobInput"] = jobChanges
		}
		if len(ownerChanges) > 0 {
			variables["ownerInput"] = ownerChanges
		}
		if len(contractorChanges) > 0 {
			variables["contractorInput"] = contractorChanges
		}
		if len(supervisorChanges) > 0 {
			variables["supervisorInput"] = supervisorChanges
		}
		if len(authorChanges) > 0 {
			variables["authorInput"] = authorChanges
		}

		// Log mutation ve variables
		log.Printf("Mutation sorgusu:\n%s", updateMutation)
		variablesJSON, _ := json.MarshalIndent(variables, "", "  ")
		log.Printf("Variables:\n%s", string(variablesJSON))

		var upsertResult struct {
			Job *struct {
				ID          int    `json:"id"`
				YibfNo      int    `json:"YibfNo"`
				CompanyCode int    `json:"CompanyCode"`
				Title       string `json:"Title"`
			} `json:"job"`
			Owner *struct {
				ID    int    `json:"id"`
				Name  string `json:"Name"`
				YDSID int    `json:"YDSID"`
			} `json:"owner"`
			Contractor *struct {
				ID    int    `json:"id"`
				Name  string `json:"Name"`
				YDSID int    `json:"YDSID"`
			} `json:"contractor"`
			Supervisor *struct {
				ID    int    `json:"id"`
				Name  string `json:"Name"`
				YDSID int    `json:"YDSID"`
			} `json:"supervisor"`
			Author *struct {
				ID                       int    `json:"id"`
				Static                   string `json:"Static"`
				Mechanic                 string `json:"Mechanic"`
				Electric                 string `json:"Electric"`
				Architect                string `json:"Architect"`
				GeotechnicalEngineer     string `json:"GeotechnicalEngineer"`
				GeotechnicalGeologist    string `json:"GeotechnicalGeologist"`
				GeotechnicalGeophysicist string `json:"GeotechnicalGeophysicist"`
			} `json:"author"`
		}

		if err := graphqlClient.Execute(updateMutation, variables, jwtToken, &upsertResult); err != nil {
			log.Printf("GraphQL mutation hatası - YİBF No: %d", yibfID)
			log.Printf("Mutation sorgusu:\n%s", updateMutation)
			log.Printf("Variables:\n%s", string(variablesJSON))

			// GraphQL hata detaylarını parse et
			var graphqlError struct {
				Errors []struct {
					Message string   `json:"message"`
					Path    []string `json:"path"`
				} `json:"errors"`
			}

			if jsonErr := json.Unmarshal([]byte(err.Error()), &graphqlError); jsonErr == nil {
				for _, e := range graphqlError.Errors {
					log.Printf("GraphQL Hata Detayı - Path: %v, Message: %s", e.Path, e.Message)
					errMsg := fmt.Sprintf("YİBF %d için GraphQL hatası - Path: %v, Message: %s", yibfID, e.Path, e.Message)
					logError(errMsg)
				}
			} else {
				log.Printf("Ham hata mesajı: %v", err)
				errMsg := fmt.Sprintf("YİBF %d için işlem hatası: %v", yibfID, err)
				logError(errMsg)
			}

			// Hata durumunda bile başarılı olan işlemleri kontrol et
			if upsertResult.Job != nil || upsertResult.Owner != nil || upsertResult.Contractor != nil ||
				upsertResult.Supervisor != nil || upsertResult.Author != nil {
				log.Printf("ID %d için bazı veriler başarıyla işlendi, ancak hata oluştu", yibfID)
				processedIDs = append(processedIDs, yibfID)
			} else {
				log.Printf("ID %d için hiçbir veri işlenemedi", yibfID)
				failedIDs = append(failedIDs, yibfID)
			}
			continue
		}

		log.Printf("ID %d için tüm veriler başarıyla işlendi", yibfID)
		processedIDs = append(processedIDs, yibfID)
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

// contains yardımcı fonksiyonu
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
