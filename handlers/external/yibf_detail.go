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

	"github.com/polatbilal/gqlgen-ent/handlers/client"
	"github.com/polatbilal/gqlgen-ent/handlers/service"

	"github.com/gin-gonic/gin"
)

func YibfDetail(c *gin.Context) {
	// GraphQL için JWT token
	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT Token gerekli"})
		return
	}

	// ID parametresini al
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parametresi gerekli"})
		return
	}

	// ID'yi integer'a çevir
	yibfID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID formatı"})
		return
	}

	// Detay işlemlerini gerçekleştir
	err = ProcessYibfDetail(c, yibfID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": fmt.Sprintf("ID %d başarıyla işlendi", yibfID),
	})
}

// ProcessYibfDetail YDK'dan gelen iş detaylarını işler ve veritabanına kaydeder
func ProcessYibfDetail(c *gin.Context, yibfID int) error {
	// GraphQL için JWT token
	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		return fmt.Errorf("JWT Token gerekli")
	}

	// CompanyCode parametresini al
	companyCode := c.Query("companyCode")
	if companyCode == "" {
		return fmt.Errorf("CompanyCode parametresi gerekli")
	}

	// CompanyCode'u integer'a çevir
	companyCodeInt, err := strconv.Atoi(companyCode)
	if err != nil {
		return fmt.Errorf("Geçersiz CompanyCode formatı")
	}

	// Token bilgisini veritabanından al
	companyToken, err := service.GetCompanyTokenFromDB(c.Request.Context(), companyCodeInt)
	if err != nil {
		return err
	}

	svc := &service.ExternalService{
		BaseURL: os.Getenv("YDK_BASE_URL"),
		Client:  &http.Client{},
	}

	// YDK API'den detay bilgilerini çek
	url := fmt.Sprintf("%s%s%d", svc.BaseURL, service.ENDPOINT_YIBF_BY_ID, yibfID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", companyToken.Token))

	resp, err := svc.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var detail service.YIBFResponse
	if err := json.Unmarshal(body, &detail); err != nil {
		return fmt.Errorf("Parse hatası: %v", err)
	}

	// GraphQL client oluştur
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	graphqlClient := client.GraphQLClient{
		URL: fmt.Sprintf("%s://%s/graphql", scheme, c.Request.Host),
	}

	// Önce şirketin var olup olmadığını kontrol et
	checkQuery := `
	query JobDetail($YibfNo: Int!) {
		job(YibfNo: $YibfNo) {
			id
			YibfNo
		}
	}
	`

	checkVariables := map[string]interface{}{
		"YibfNo": detail.ID,
	}

	var checkResult struct {
		Job struct {
			ID     string `json:"id"`
			YibfNo int    `json:"YibfNo"`
		} `json:"job"`
	}

	err = graphqlClient.Execute(checkQuery, checkVariables, jwtToken, &checkResult)
	if err != nil {
		return fmt.Errorf("GraphQL hatası: %v", err)
	}

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
		"BKSReferenceNo":   service.SafeStringToInt(detail.ReferenceNumber),
		"Coordinates":      service.CoordinatesToString(detail.Position.Coordinates),
		"UploadedFile":     detail.UploadedFile,
		"IndustryArea":     detail.YibfStructure.IndustryArea,
		"ClusterStructure": detail.ClusterStructure,
		"IsLicenseExpired": detail.IsLicenseExpired,
		"IsCompleted":      detail.IsCompleted,
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

	// Inspector verilerini al
	inspectorURL := fmt.Sprintf("%s%s", svc.BaseURL, service.ENDPOINT_YIBF_DEPARTMENT_EMPLOYEE)
	inspectorRequestBody := map[string]interface{}{
		"filter": [][]interface{}{
			{"yibfId", "=", yibfID},
		},
	}

	inspectorJsonBody, err := json.Marshal(inspectorRequestBody)
	if err != nil {
		return fmt.Errorf("Inspector request body oluşturma hatası: %v", err)
	}

	inspectorReq, err := http.NewRequest("POST", inspectorURL, bytes.NewBuffer(inspectorJsonBody))
	if err != nil {
		return fmt.Errorf("Inspector request oluşturma hatası: %v", err)
	}

	inspectorReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", companyToken.Token))
	inspectorReq.Header.Set("Content-Type", "application/json")

	inspectorResp, err := svc.Client.Do(inspectorReq)
	if err != nil {
		return fmt.Errorf("Inspector isteği başarısız: %v", err)
	}
	defer inspectorResp.Body.Close()

	inspectorBody, err := io.ReadAll(inspectorResp.Body)
	if err != nil {
		return fmt.Errorf("Inspector yanıtı okuma hatası: %v", err)
	}

	var inspectorResponse service.YIBFInspectorResponse
	if err := json.Unmarshal(inspectorBody, &inspectorResponse); err != nil {
		return fmt.Errorf("Inspector parse hatası: %v", err)
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

	// Author verilerini al
	authorURL := fmt.Sprintf("%s%s", svc.BaseURL, service.ENDPOINT_YIBF_PROJECT_OWNER)
	authorRequestBody := map[string]interface{}{
		"filter": [][]interface{}{
			{"yibfId", "=", yibfID},
		},
	}

	authorJsonBody, err := json.Marshal(authorRequestBody)
	if err != nil {
		return fmt.Errorf("Author request body oluşturma hatası: %v", err)
	}

	authorReq, err := http.NewRequest("POST", authorURL, bytes.NewBuffer(authorJsonBody))
	if err != nil {
		return fmt.Errorf("Author request oluşturma hatası: %v", err)
	}

	authorReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", companyToken.Token))
	authorReq.Header.Set("Content-Type", "application/json")

	authorResp, err := svc.Client.Do(authorReq)
	if err != nil {
		return fmt.Errorf("Author isteği başarısız: %v", err)
	}
	defer authorResp.Body.Close()

	authorBody, err := io.ReadAll(authorResp.Body)
	if err != nil {
		return fmt.Errorf("Author yanıtı okuma hatası: %v", err)
	}

	var authorResponse service.YIBFAuthorResponse
	if err := json.Unmarshal(authorBody, &authorResponse); err != nil {
		return fmt.Errorf("Author parse hatası: %v", err)
	}

	// Author verilerini hazırla
	var authorData map[string]interface{}
	if len(authorResponse.Items) > 0 {
		authorData = make(map[string]interface{})
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
	}

	// Owner verilerini hazırla
	var ownerData map[string]interface{}
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
	var contractorData map[string]interface{}
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
	var supervisorData map[string]interface{}
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

	// Her entity için ayrı kontrol ve güncelleme yap
	var mutations []string
	updateVariables := make(map[string]interface{})

	// Job kontrolü ve güncellemesi
	if checkResult.Job.ID != "" {
		mutations = append(mutations, `
			updateJob(yibfNo: $yibfNo, input: $jobInput) {
				id
				YibfNo
			}`)
		updateVariables["jobInput"] = jobData
		updateVariables["yibfNo"] = yibfID
	} else {
		mutations = append(mutations, `
			createJob(input: $jobInput) {
				id
				YibfNo
			}`)
		updateVariables["jobInput"] = jobData
	}

	// Owner verilerini ekle
	if ownerData != nil {
		if checkResult.Job.ID != "" {
			mutations = append(mutations, `
				updateOwner(yibfNo: $yibfNo, input: $ownerInput) {
					Name
					YDSID
				}`)
		} else {
			mutations = append(mutations, `
				createOwner(input: $ownerInput) {
					Name
					YDSID
				}`)
		}
		updateVariables["ownerInput"] = ownerData
	}

	// Contractor verilerini ekle
	if contractorData != nil {
		if checkResult.Job.ID != "" {
			mutations = append(mutations, `
				updateContractor(yibfNo: $yibfNo, input: $contractorInput) {
					Name
					YDSID
				}`)
		} else {
			mutations = append(mutations, `
				createContractor(input: $contractorInput) {
					Name
					YDSID
				}`)
		}
		updateVariables["contractorInput"] = contractorData
	}

	// Supervisor verilerini ekle
	if supervisorData != nil {
		if checkResult.Job.ID != "" {
			mutations = append(mutations, `
				updateSupervisor(yibfNo: $yibfNo, input: $supervisorInput) {
					Name
					YDSID
				}`)
		} else {
			mutations = append(mutations, `
				createSupervisor(input: $supervisorInput) {
					Name
					YDSID
				}`)
		}
		updateVariables["supervisorInput"] = supervisorData
	}

	// Author verilerini ekle
	if authorData != nil {
		if checkResult.Job.ID != "" {
			mutations = append(mutations, `
				updateAuthor(yibfNo: $yibfNo, input: $authorInput) {
					Static
					Mechanic
					Electric
					Architect
					GeotechnicalEngineer
					GeotechnicalGeologist
					GeotechnicalGeophysicist
				}`)
		} else {
			mutations = append(mutations, `
				createAuthor(input: $authorInput) {
					Static
					Mechanic
					Electric
					Architect
					GeotechnicalEngineer
					GeotechnicalGeologist
					GeotechnicalGeophysicist
				}`)
		}
		updateVariables["authorInput"] = authorData
	}

	// Mutation parametrelerini oluştur
	var paramParts []string
	if checkResult.Job.ID != "" {
		paramParts = append(paramParts, "$yibfNo: Int!")
	}
	if _, ok := updateVariables["jobInput"]; ok {
		paramParts = append(paramParts, "$jobInput: JobInput!")
	}
	if _, ok := updateVariables["ownerInput"]; ok {
		paramParts = append(paramParts, "$ownerInput: JobOwnerInput!")
	}
	if _, ok := updateVariables["contractorInput"]; ok {
		paramParts = append(paramParts, "$contractorInput: JobContractorInput!")
	}
	if _, ok := updateVariables["supervisorInput"]; ok {
		paramParts = append(paramParts, "$supervisorInput: JobSupervisorInput!")
	}
	if _, ok := updateVariables["authorInput"]; ok {
		paramParts = append(paramParts, "$authorInput: JobAuthorInput!")
	}

	// Tam mutation string'ini oluştur
	updateMutation := fmt.Sprintf(`
	mutation UpdateJob(%s) {
		%s
	}
	`, strings.Join(paramParts, ", "), strings.Join(mutations, "\n"))

	// Mutation'ı çalıştır
	var updateResult struct {
		UpdateJob *struct {
			ID     string `json:"id"`
			YibfNo int    `json:"YibfNo"`
		} `json:"updateJob,omitempty"`
		CreateJob *struct {
			ID     string `json:"id"`
			YibfNo int    `json:"YibfNo"`
		} `json:"createJob,omitempty"`
	}

	if err := graphqlClient.Execute(updateMutation, updateVariables, jwtToken, &updateResult); err != nil {
		return fmt.Errorf("GraphQL mutation hatası: %v", err)
	}

	return nil
}

// ProcessYibfDetailBatch birden fazla YDK işini toplu olarak işler
func ProcessYibfDetailBatch(c *gin.Context, yibfIDs []int) (processedIDs []int, failedIDs []int) {
	for _, id := range yibfIDs {
		if err := ProcessYibfDetail(c, id); err != nil {
			log.Printf("ID %d işlenirken hata oluştu: %v", id, err)
			failedIDs = append(failedIDs, id)
		} else {
			processedIDs = append(processedIDs, id)
		}
	}
	return processedIDs, failedIDs
}
