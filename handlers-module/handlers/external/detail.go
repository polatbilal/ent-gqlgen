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

	"github.com/gofiber/fiber/v2"
	"github.com/polatbilal/gqlgen-ent/api-core/graphql/model"
	"github.com/polatbilal/gqlgen-ent/handlers-module/handlers/client"
	"github.com/polatbilal/gqlgen-ent/handlers-module/handlers/service"
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

// Veri karşılaştırma fonksiyonları
func compareJobData(current *model.JobInput, new map[string]interface{}) map[string]interface{} {
	changes := make(map[string]interface{})

	for key, newValue := range new {
		if key == "YibfNo" || key == "CompanyCode" {
			continue
		}

		currentValue := getFieldValue(current, key)
		if !compareValues(currentValue, newValue, key) {
			changes[key] = newValue
		}
	}

	return changes
}

func compareOwnerData(current *model.JobOwnerInput, new map[string]interface{}) map[string]interface{} {
	changes := make(map[string]interface{})

	for key, newValue := range new {
		if key == "YibfNo" {
			continue
		}

		currentValue := getFieldValue(current, key)
		if !compareValues(currentValue, newValue, key) {
			changes[key] = newValue
		}
	}

	return changes
}

func compareContractorData(current *model.JobContractorInput, new map[string]interface{}) map[string]interface{} {
	changes := make(map[string]interface{})

	for key, newValue := range new {
		if key == "YibfNo" {
			continue
		}

		currentValue := getFieldValue(current, key)
		if !compareValues(currentValue, newValue, key) {
			changes[key] = newValue
		}
	}

	return changes
}

func compareSupervisorData(current *model.JobSupervisorInput, new map[string]interface{}) map[string]interface{} {
	changes := make(map[string]interface{})

	for key, newValue := range new {
		if key == "YibfNo" {
			continue
		}

		currentValue := getFieldValue(current, key)
		if !compareValues(currentValue, newValue, key) {
			changes[key] = newValue
		}
	}

	return changes
}

func compareAuthorData(current *model.JobAuthorInput, new map[string]interface{}) map[string]interface{} {
	changes := make(map[string]interface{})

	for key, newValue := range new {
		if key == "YibfNo" {
			continue
		}

		currentValue := getFieldValue(current, key)
		if !compareValues(currentValue, newValue, key) {
			changes[key] = newValue
		}
	}

	return changes
}

// Yardımcı fonksiyonlar
func getFieldValue(obj interface{}, fieldName string) interface{} {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		return nil
	}

	return field.Interface()
}

func compareValues(current, new interface{}, fieldName string) bool {
	if current == nil && new == nil {
		return true
	}
	if current == nil || new == nil {
		return false
	}

	// Tarih karşılaştırması
	if currentStr, ok := current.(string); ok {
		if strings.Contains(fieldName, "Date") {
			newStr := fmt.Sprintf("%v", new)
			return service.CompareDates(currentStr, newStr)
		}
	}

	// Koordinat karşılaştırması
	if fieldName == "Coordinates" {
		currentStr := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%v", current), " ", ""), ",", ".")
		newStr := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%v", new), " ", ""), ",", ".")
		return currentStr == newStr
	}

	// Genel karşılaştırma
	return fmt.Sprintf("%v", current) == fmt.Sprintf("%v", new)
}

func YibfDetail(c *fiber.Ctx, yibfNumbers []int, companyCodeStr string) error {
	jwtToken := c.Get("Authorization")
	if jwtToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT Token gerekli"})
	}

	if len(yibfNumbers) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "YİBF numaraları gerekli"})
	}

	companyCode := service.SafeStringToInt(companyCodeStr)
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

	graphqlClient := client.GraphQLClient{
		URL: fmt.Sprintf("%s://%s/graphql", scheme, c.Hostname()),
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
			Job        *model.JobInput           `json:"job"`
			Owner      *model.JobOwnerInput      `json:"owner"`
			Contractor *model.JobContractorInput `json:"contractor"`
			Supervisor *model.JobSupervisorInput `json:"supervisor"`
			Author     *model.JobAuthorInput     `json:"author"`
		}

		err = graphqlClient.Execute(checkQuery, checkVariables, jwtToken, &checkResult)

		// Job ve ilişkili verilerin durumunu belirle
		isNewJob := checkResult.Job == nil
		isNewOwner := checkResult.Owner == nil
		isNewContractor := checkResult.Contractor == nil
		isNewSupervisor := checkResult.Supervisor == nil
		isNewAuthor := checkResult.Author == nil

		// Değişiklikleri topla
		var mutationParts []string
		var mutationFields []string
		var variables = make(map[string]interface{})

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

		// Detay verilerini hazırla
		currentResult := map[string]interface{}{
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
		if detail.DistributionDate > 0 {
			currentResult["job"].(map[string]interface{})["DistributionDate"] = service.SafeUnixToDate(detail.DistributionDate)
		}
		if detail.ContractDate > 0 {
			currentResult["job"].(map[string]interface{})["ContractDate"] = service.SafeUnixToDate(detail.ContractDate)
		}
		if detail.LicenseDate > 0 {
			currentResult["job"].(map[string]interface{})["LicenseDate"] = service.SafeUnixToDate(detail.LicenseDate)
		}
		if detail.CompleteDate > 0 {
			currentResult["job"].(map[string]interface{})["CompletionDate"] = service.SafeUnixToDate(detail.CompleteDate)
		}

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

		// Job verilerini hazırla
		if job, ok := currentResult["job"].(map[string]interface{}); ok {
			if isNewJob {
				mutationParts = append(mutationParts, `$jobInput: JobInput!`)
				mutationFields = append(mutationFields, `
					createJob(input: $jobInput) {
						id
						YibfNo
						CompanyCode
						Title
					}`)
				variables["jobInput"] = job
			} else {
				jobChanges := compareJobData(checkResult.Job, job)
				if len(jobChanges) > 0 {
					mutationParts = append(mutationParts, `$jobInput: JobInput!`)
					mutationFields = append(mutationFields, `
						updateJob(yibfNo: $yibfNo, input: $jobInput) {
							id
							YibfNo
							CompanyCode
							Title
						}`)
					variables["jobInput"] = jobChanges
				}
			}
		}

		// Owner verilerini hazırla
		if owner, ok := currentResult["owner"].(map[string]interface{}); ok {
			if isNewOwner {
				mutationParts = append(mutationParts, `$ownerInput: JobOwnerInput!`)
				mutationFields = append(mutationFields, `
					createOwner(input: $ownerInput) {
						id
						Name
						YDSID
					}`)
				variables["ownerInput"] = owner
			} else {
				ownerChanges := compareOwnerData(checkResult.Owner, owner)
				if len(ownerChanges) > 0 {
					mutationParts = append(mutationParts, `$ownerInput: JobOwnerInput!`)
					mutationFields = append(mutationFields, `
						updateOwner(yibfNo: $yibfNo, input: $ownerInput) {
							id
							Name
							YDSID
						}`)
					variables["ownerInput"] = ownerChanges
				}
			}
		}

		// Contractor verilerini hazırla
		if contractor, ok := currentResult["contractor"].(map[string]interface{}); ok {
			if isNewContractor {
				mutationParts = append(mutationParts, `$contractorInput: JobContractorInput!`)
				mutationFields = append(mutationFields, `
					createContractor(input: $contractorInput) {
						id
						Name
						YDSID
					}`)
				variables["contractorInput"] = contractor
			} else {
				contractorChanges := compareContractorData(checkResult.Contractor, contractor)
				if len(contractorChanges) > 0 {
					mutationParts = append(mutationParts, `$contractorInput: JobContractorInput!`)
					mutationFields = append(mutationFields, `
						updateContractor(yibfNo: $yibfNo, input: $contractorInput) {
							id
							Name
							YDSID
						}`)
					variables["contractorInput"] = contractorChanges
				}
			}
		}

		// Supervisor verilerini hazırla
		if supervisor, ok := currentResult["supervisor"].(map[string]interface{}); ok {
			if isNewSupervisor {
				mutationParts = append(mutationParts, `$supervisorInput: JobSupervisorInput!`)
				mutationFields = append(mutationFields, `
					createSupervisor(input: $supervisorInput) {
						id
						Name
						YDSID
					}`)
				variables["supervisorInput"] = supervisor
			} else {
				supervisorChanges := compareSupervisorData(checkResult.Supervisor, supervisor)
				if len(supervisorChanges) > 0 {
					mutationParts = append(mutationParts, `$supervisorInput: JobSupervisorInput!`)
					mutationFields = append(mutationFields, `
						updateSupervisor(yibfNo: $yibfNo, input: $supervisorInput) {
							id
							Name
							YDSID
						}`)
					variables["supervisorInput"] = supervisorChanges
				}
			}
		}

		// Author verilerini hazırla
		if author, ok := currentResult["author"].(map[string]interface{}); ok {
			if isNewAuthor {
				mutationParts = append(mutationParts, `$authorInput: JobAuthorInput!`)
				mutationFields = append(mutationFields, `
					createAuthor(input: $authorInput) {
						id
						Static
						Mechanic
						Electric
						Architect
						GeotechnicalEngineer
						GeotechnicalGeologist
						GeotechnicalGeophysicist
					}`)
				variables["authorInput"] = author
			} else {
				authorChanges := compareAuthorData(checkResult.Author, author)
				if len(authorChanges) > 0 {
					mutationParts = append(mutationParts, `$authorInput: JobAuthorInput!`)
					mutationFields = append(mutationFields, `
						updateAuthor(yibfNo: $yibfNo, input: $authorInput) {
							id
							Static
							Mechanic
							Electric
							Architect
							GeotechnicalEngineer
							GeotechnicalGeologist
							GeotechnicalGeophysicist
						}`)
					variables["authorInput"] = authorChanges
				}
			}
		}

		// Mutation oluştur
		var mutation string
		if len(mutationParts) > 0 {
			if isNewJob {
				mutation = fmt.Sprintf(`
					mutation CreateJobWithRelations(%s) {
						%s
					}`, strings.Join(mutationParts, ", "), strings.Join(mutationFields, "\n"))
			} else {
				variables["yibfNo"] = yibfID
				mutation = fmt.Sprintf(`
					mutation UpdateJobWithRelations($yibfNo: Int!, %s) {
						%s
					}`, strings.Join(mutationParts, ", "), strings.Join(mutationFields, "\n"))
			}

			// Debug logları
			log.Printf("YİBF No: %d için mutation parçaları: %v", yibfID, mutationParts)
			// log.Printf("YİBF No: %d için mutation alanları: %v", yibfID, mutationFields)
			// log.Printf("YİBF No: %d için variables: %+v", yibfID, variables)
			// log.Printf("YİBF No: %d için oluşturulan mutation:\n%s", yibfID, mutation)

			// Mutation'ı çalıştır
			var upsertResult struct {
				Job        *struct{ ID int }
				Owner      *struct{ ID int }
				Contractor *struct{ ID int }
				Supervisor *struct{ ID int }
				Author     *struct{ ID int }
			}

			if err := graphqlClient.Execute(mutation, variables, jwtToken, &upsertResult); err != nil {
				errMsg := fmt.Sprintf("YİBF %d için GraphQL mutation hatası: %v\nMutation: %s\nVariables: %+v",
					yibfID, err, mutation, variables)
				logError(errMsg)
				failedIDs = append(failedIDs, yibfID)
				continue
			}

			// Başarılı işlem logu
			log.Printf("YİBF %d için mutation başarıyla çalıştırıldı. Sonuç: %+v", yibfID, upsertResult)
		} else {
			log.Printf("YİBF %d için değişiklik yok, mutation çalıştırılmayacak", yibfID)
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
