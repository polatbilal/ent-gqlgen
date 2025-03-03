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

	"github.com/gofiber/fiber/v2"
	"github.com/polatbilal/gqlgen-ent/api-core/graphql/model"
	"github.com/polatbilal/gqlgen-ent/handlers-module/client"
	"github.com/polatbilal/gqlgen-ent/handlers-module/service"
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

func YibfDetail(c *fiber.Ctx, yibfNumbers []int, companyCode int) error {
	jwtToken := c.Get("Authorization")
	if jwtToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT Token gerekli"})
	}

	if len(yibfNumbers) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "YİBF numaraları gerekli"})
	}

	// companyCode := service.SafeStringToInt(companyCodeStr)
	if companyCode == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "CompanyCode gerekli"})
	}

	companyToken, err := service.GetCompanyTokenFromDB(c.Context(), companyCode)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if companyToken.Token == "" || companyToken.DepartmentId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçerli token veya department ID bulunamadı"})
	}

	svc := &service.ExternalService{
		BaseURL: os.Getenv("YDK_BASE_URL"),
		Client:  &http.Client{},
	}

	var processedIDs []int
	var failedIDs []int
	var results []map[string]interface{}

	scheme := "http"
	if c.Protocol() == "https" {
		scheme = "https"
	}

	graphqlClient := client.NewGraphQLClient(scheme)

	for _, yibfID := range yibfNumbers {
		// Önce mevcut kaydı kontrol et
		checkQuery := `
			query CheckJobBatch($yibfNo: Int!) {
				jobBatchQuery(yibfNo: $yibfNo) {
					job {
						id
						CompanyCode
						YibfNo
						Title
						Administration
						State
						Island
						Parcel
						Sheet
						DistributionDate
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
					}
					owner {
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
					contractor {
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
					supervisor {
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
					author {
						id
						Static
						Mechanic
						Electric
						Architect
						GeotechnicalEngineer
						GeotechnicalGeologist
						GeotechnicalGeophysicist
					}
					engineer {
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
			}`

		checkVariables := map[string]interface{}{
			"yibfNo": yibfID,
		}

		var checkResponse struct {
			JobBatchQuery struct {
				Job        *model.JobInput           `json:"job"`
				Owner      *model.JobOwnerInput      `json:"owner"`
				Contractor *model.JobContractorInput `json:"contractor"`
				Supervisor *model.JobSupervisorInput `json:"supervisor"`
				Author     *model.JobAuthorInput     `json:"author"`
				Engineer   *model.JobEngineerInput   `json:"engineer"`
			} `json:"jobBatchQuery"`
		}

		log.Printf("YİBF %d için mevcut durum kontrolü yapılıyor...", yibfID)
		err = graphqlClient.Execute(checkQuery, checkVariables, jwtToken, &checkResponse)

		var checkResult = checkResponse.JobBatchQuery

		if err != nil {
			if strings.Contains(err.Error(), "kayıt bulunamadı") || strings.Contains(err.Error(), "parsing time") {
				log.Printf("YİBF %d için kayıt bulunamadı veya tarih parse hatası, yeni kayıt oluşturulacak", yibfID)
			} else {
				errMsg := fmt.Sprintf("YİBF %d için mevcut durum kontrolü başarısız: %v", yibfID, err)
				logError(errMsg)
				failedIDs = append(failedIDs, yibfID)
				continue
			}
		} else if checkResult.Job != nil {
			log.Printf("YİBF %d için mevcut durum kontrolü başarılı, güncelleme yapılacak", yibfID)
		} else {
			log.Printf("YİBF %d için kayıt bulunamadı, yeni kayıt oluşturulacak", yibfID)
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

		// currentResult map'ini oluştur
		currentResult := make(map[string]interface{})

		// Job verilerini hazırla
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
			"BKSReferenceNo":   service.SafeStringToInt(detail.ReferenceNumber),
			"Coordinates":      service.CoordinatesToString(detail.Position.Coordinates),
			"UploadedFile":     detail.UploadedFile,
			"IndustryArea":     detail.YibfStructure.IndustryArea,
			"ClusterStructure": detail.ClusterStructure,
			"IsLicenseExpired": detail.IsLicenseExpired,
			"IsCompleted":      detail.IsCompleted,
		}

		// Unix timestamp'ten YYYY-MM-DD formatına çevir
		if detail.DistributionDate > 0 {
			date := time.Unix(detail.DistributionDate/1000, 0)
			jobData["DistributionDate"] = date.Format("2006-01-02")
		}
		if detail.ContractDate > 0 {
			date := time.Unix(detail.ContractDate/1000, 0)
			jobData["ContractDate"] = date.Format("2006-01-02")
		}
		if detail.LicenseDate > 0 {
			date := time.Unix(detail.LicenseDate/1000, 0)
			jobData["LicenseDate"] = date.Format("2006-01-02")
		}
		if detail.CompleteDate > 0 {
			date := time.Unix(detail.CompleteDate/1000, 0)
			jobData["CompletionDate"] = date.Format("2006-01-02")
		}

		currentResult["job"] = jobData

		// Owner verilerini ekle
		if detail.YibfOwner.Person.ID > 0 {
			currentResult["owner"] = map[string]interface{}{
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
			currentResult["contractor"] = map[string]interface{}{
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
			currentResult["supervisor"] = map[string]interface{}{
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
					case 4:
						authorData["GeotechnicalEngineer"] = fullName
					case 6:
						authorData["GeotechnicalGeologist"] = fullName
					case 7:
						authorData["GeotechnicalGeophysicist"] = fullName
					}
				}
			}
			currentResult["author"] = authorData
		}

		// Engineer verilerini al
		engineerURL := fmt.Sprintf("%s%s", svc.BaseURL, service.ENDPOINT_YIBF_ENGINEER)
		engineerRequestBody := map[string]interface{}{
			"filter": [][]interface{}{
				{"yibfId", "=", yibfID},
			},
		}

		engineerJsonBody, err := json.Marshal(engineerRequestBody)
		if err != nil {
			errMsg := fmt.Sprintf("YİBF %d için engineer request body oluşturma hatası: %v", yibfID, err)
			logError(errMsg)
			continue
		}

		engineerReq, err := http.NewRequest("POST", engineerURL, bytes.NewBuffer(engineerJsonBody))
		if err != nil {
			errMsg := fmt.Sprintf("YİBF %d için engineer request oluşturma hatası: %v", yibfID, err)
			logError(errMsg)
			continue
		}

		engineerReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", companyToken.Token))
		engineerReq.Header.Set("Content-Type", "application/json")

		engineerResp, err := svc.Client.Do(engineerReq)
		if err != nil {
			errMsg := fmt.Sprintf("YİBF %d için engineer isteği başarısız: %v", yibfID, err)
			logError(errMsg)
			continue
		}

		engineerBody, err := io.ReadAll(engineerResp.Body)
		engineerResp.Body.Close()
		if err != nil {
			errMsg := fmt.Sprintf("YİBF %d için engineer yanıtı okuma hatası: %v", yibfID, err)
			logError(errMsg)
			continue
		}

		var engineerResponse service.YIBFInspectorResponse
		if err := json.Unmarshal(engineerBody, &engineerResponse); err != nil {
			errMsg := fmt.Sprintf("YİBF %d için engineer parse hatası: %v", yibfID, err)
			logError(errMsg)
			continue
		}

		// Engineer verilerini ekle
		if len(engineerResponse.Items) > 0 {
			engineerData := make(map[string]interface{})
			engineerData["YibfNo"] = yibfID

			for _, engineer := range engineerResponse.Items {
				// UDIM true ise Inspector olarak ata
				if engineer.Udim {
					engineerData["Inspector"] = engineer.UserId
					engineerData["Static"] = engineer.UserId
					continue
				}

				switch engineer.TaskId {
				case 1:
					engineerData["Inspector"] = engineer.UserId
				case 2:
					switch engineer.TitleId {
					case 1:
						engineerData["Electric"] = engineer.UserId
					case 2:
						engineerData["Mechanic"] = engineer.UserId
					case 3:
						engineerData["Architect"] = engineer.UserId
					case 4:
						engineerData["Static"] = engineer.UserId
					}

				case 14:
					switch engineer.TitleId {
					case 1:
						engineerData["ElectricController"] = engineer.UserId
					case 2:
						engineerData["MechanicController"] = engineer.UserId
					case 3, 4:
						engineerData["Controller"] = engineer.UserId
					}
				}
			}

			// Eğer en az bir değer varsa engineer verilerini ekle
			if len(engineerData) > 1 { // YibfNo dışında en az bir değer varsa
				currentResult["engineer"] = engineerData
			}
		}

		// Batch mutation için verileri hazırla
		jobInput := currentResult["job"].(map[string]interface{})
		batchInput := map[string]interface{}{
			"YibfNo":   yibfID,
			"jobInput": jobInput,
		}

		if owner, ok := currentResult["owner"].(map[string]interface{}); ok {
			batchInput["ownerInput"] = owner
		}
		if contractor, ok := currentResult["contractor"].(map[string]interface{}); ok {
			batchInput["contractorInput"] = contractor
		}
		if supervisor, ok := currentResult["supervisor"].(map[string]interface{}); ok {
			batchInput["supervisorInput"] = supervisor
		}
		if author, ok := currentResult["author"].(map[string]interface{}); ok {
			batchInput["authorInput"] = author
		}
		if engineer, ok := currentResult["engineer"].(map[string]interface{}); ok {
			batchInput["engineerInput"] = engineer
		}

		// Debug logları
		log.Printf("YİBF No: %d için hazırlanan veriler:", yibfID)
		log.Printf("Job verisi: %+v", jobInput)
		if owner, ok := currentResult["owner"].(map[string]interface{}); ok {
			log.Printf("Owner verisi: %+v", owner)
		}
		if contractor, ok := currentResult["contractor"].(map[string]interface{}); ok {
			log.Printf("Contractor verisi: %+v", contractor)
		}
		if supervisor, ok := currentResult["supervisor"].(map[string]interface{}); ok {
			log.Printf("Supervisor verisi: %+v", supervisor)
		}
		if author, ok := currentResult["author"].(map[string]interface{}); ok {
			log.Printf("Author verisi: %+v", author)
		}
		if engineer, ok := currentResult["engineer"].(map[string]interface{}); ok {
			log.Printf("Engineer verisi: %+v", engineer)
		}

		// Batch mutation'ı çalıştır
		mutation := `
			mutation JobBatch($input: JobBatchInput!) {
				jobBatch(input: $input) {
					job {
						id
						YibfNo
						CompanyCode
						Title
					}
					owner {
						id
						Name
						YDSID
					}
					contractor {
						id
						Name
						YDSID
					}
					supervisor {
						id
						Name
						YDSID
					}
					author {
						id
						Static
						Mechanic
						Electric
						Architect
						GeotechnicalEngineer
						GeotechnicalGeologist
						GeotechnicalGeophysicist
					}
					engineer {
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
			}`

		variables := map[string]interface{}{
			"input": batchInput,
		}

		// Debug logları
		log.Printf("YİBF No: %d için mutation çalıştırılıyor", yibfID)
		// log.Printf("Mutation: %s", mutation)
		// log.Printf("Variables: %+v", variables)

		// Mutation'ı çalıştır
		var batchResponse struct {
			JobBatch struct {
				Job struct {
					ID          int    `json:"id"`
					YibfNo      int    `json:"yibfNo"`
					CompanyCode int    `json:"companyCode"`
					Title       string `json:"title"`
				} `json:"job"`
				Owner struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					YDSID int    `json:"ydsid"`
				} `json:"owner"`
				Contractor struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					YDSID int    `json:"ydsid"`
				} `json:"contractor"`
				Supervisor struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					YDSID int    `json:"ydsid"`
				} `json:"supervisor"`
				Author struct {
					ID                       int    `json:"id"`
					Static                   string `json:"static"`
					Mechanic                 string `json:"mechanic"`
					Electric                 string `json:"electric"`
					Architect                string `json:"architect"`
					GeotechnicalEngineer     string `json:"geotechnicalEngineer"`
					GeotechnicalGeologist    string `json:"geotechnicalGeologist"`
					GeotechnicalGeophysicist string `json:"geotechnicalGeophysicist"`
				} `json:"author"`
				Engineer struct {
					Inspector          string `json:"inspector"`
					Static             string `json:"static"`
					Architect          string `json:"architect"`
					Mechanic           string `json:"mechanic"`
					Electric           string `json:"electric"`
					Controller         string `json:"controller"`
					MechanicController string `json:"mechanicController"`
					ElectricController string `json:"electricController"`
				} `json:"engineer"`
			} `json:"jobBatch"`
		}

		if err := graphqlClient.Execute(mutation, variables, jwtToken, &batchResponse); err != nil {
			errMsg := fmt.Sprintf("YİBF %d için GraphQL batch mutation hatası: %v", yibfID, err)
			logError(errMsg)
			log.Printf("Hata detayı - Mutation: %s", mutation)
			log.Printf("Hata detayı - Variables: %+v", variables)
			failedIDs = append(failedIDs, yibfID)
			continue
		}

		// Başarılı işlem logu
		log.Printf("YİBF %d için mutation başarılı", yibfID)

		// Response verilerini logla
		if batchResponse.JobBatch.Job.ID > 0 {
			logBatchResponse(yibfID, batchResponse.JobBatch)
		}

		processedIDs = append(processedIDs, yibfID)
		results = append(results, currentResult)
	}

	result := fiber.Map{
		"success": true,
		"data": fiber.Map{
			"processed_count": len(processedIDs),
			"processed_ids":   processedIDs,
			"failed_count":    len(failedIDs),
			"failed_ids":      failedIDs,
			"results":         results,
		},
	}

	c.Locals("response", result)
	return c.JSON(result)
}

// Response verilerini loglama fonksiyonu
func logBatchResponse(yibfID int, response interface{}) {
	if response == nil {
		return
	}

	// Response verilerini map'e dönüştür
	responseBytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("YİBF %d response marshal hatası: %v", yibfID, err)
		return
	}

	var responseMap map[string]interface{}
	if err := json.Unmarshal(responseBytes, &responseMap); err != nil {
		log.Printf("YİBF %d response unmarshal hatası: %v", yibfID, err)
		return
	}

	// Her bir alanı kontrol et ve logla
	for key, value := range responseMap {
		if value != nil {
			log.Printf("- %s: %+v", key, value)
		}
	}
}

// YibfDetailHandler, fiber.Handler uyumlu bir wrapper fonksiyondur
func YibfDetailHandler(c *fiber.Ctx) error {
	var request struct {
		YibfNumber  int `json:"yibfNo"`
		CompanyCode int `json:"companyCode"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz istek formatı",
		})
	}

	if request.YibfNumber == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "YİBF numarası gerekli",
		})
	}

	if request.CompanyCode == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Şirket kodu gerekli",
		})
	}

	return YibfDetail(c, []int{request.YibfNumber}, request.CompanyCode)
}
