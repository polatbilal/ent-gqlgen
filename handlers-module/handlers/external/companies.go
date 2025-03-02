package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/polatbilal/gqlgen-ent/handlers-module/handlers/client"
	"github.com/polatbilal/gqlgen-ent/handlers-module/handlers/service"
)

func YDKCompanies(c *fiber.Ctx) error {
	// GraphQL için JWT token
	jwtToken := c.Get("Authorization")
	if jwtToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT Token gerekli"})
	}

	// // CompanyCode parametresini al
	// companyCode := c.Query("companyCode")
	// if companyCode == "" {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "CompanyCode parametresi gerekli"})
	// }

	// // CompanyCode'u integer'a çevir
	// companyCodeInt, err := strconv.Atoi(companyCode)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz CompanyCode formatı"})
	// }

	// Request body'den parametreleri al
	var requestParams service.ProgressPaymentRequest
	if err := c.BodyParser(&requestParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz request body: " + err.Error(),
		})
	}

	// Parametreleri kontrol et
	if requestParams.CompanyCode == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "companyCode parametresi gerekli"})
	}

	// Token bilgisini veritabanından al
	companyToken, err := service.GetCompanyTokenFromDB(c.Context(), requestParams.CompanyCode)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if companyToken.Token == "" || companyToken.DepartmentId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçerli token veya department ID bulunamadı"})
	}

	log.Printf("Şirket bilgileri alınıyor... DepartmentID: %d", companyToken.DepartmentId)

	svc := &service.ExternalService{
		BaseURL: os.Getenv("YDK_BASE_URL"),
		Client:  &http.Client{},
	}

	if svc.BaseURL == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "YDK_BASE_URL environment variable is not set"})
	}

	requestBody := map[string]interface{}{
		"id": companyToken.DepartmentId,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request body oluşturma hatası: " + err.Error()})
	}

	url := svc.BaseURL + service.ENDPOINT_COMPANY
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", companyToken.Token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := svc.Client.Do(req)
	if err != nil {
		log.Printf("YDK API isteği başarısız: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "YDK API isteği başarısız: " + err.Error()})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Yanıt body okuma hatası: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Yanıt okuma hatası: " + err.Error()})
	}

	// HTTP yanıt kodunu kontrol et
	if resp.StatusCode != http.StatusOK {
		var ydkError struct {
			Error            string `json:"error"`
			ErrorDescription string `json:"error_description"`
		}

		if err := json.Unmarshal(body, &ydkError); err != nil {
			return c.Status(resp.StatusCode).JSON(fiber.Map{"error": fmt.Sprintf("YDK API hatası: %s", string(body))})
		}

		// Hata yanıtını doğrudan ilet
		return c.Status(resp.StatusCode).JSON(ydkError)
	}

	// Struct'a parse et
	var ydkResponse service.YDKCompanyResponse
	if err := json.Unmarshal(body, &ydkResponse); err != nil {
		log.Printf("JSON parse hatası: %v\nHam veri: %s\n", err, string(body))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "JSON parse hatası: " + err.Error()})
	}

	if len(ydkResponse.Items) == 0 {
		log.Printf("ydkToken.DepartmentID: %d", companyToken.DepartmentId)
		log.Printf("Şirket bilgisi bulunamadı. Ham yanıt: %s\n", string(body))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Şirket bilgisi bulunamadı"})
	}

	// İlk şirketi al
	item := ydkResponse.Items[0]

	// Unix timestamp'i tarihe çevir
	visaDate := service.SafeUnixToDate(item.Department.VisaDate)
	visaEndDate := service.SafeUnixToDate(item.Department.VisaEndDate)

	// GraphQL client oluştur
	scheme := "http"
	if c.Protocol() == "https" {
		scheme = "https"
	}
	graphqlClient := client.NewGraphQLClient(scheme)

	// Şirket verilerini hazırla
	companyData := map[string]interface{}{
		"CompanyCode":               item.Department.FileNumber,
		"Name":                      item.Department.Name,
		"Address":                   item.Department.Person.AddressStr,
		"MobilePhone":               item.Department.Person.LastPhoneNumber,
		"Email":                     item.Department.Person.LastEPosta,
		"Website":                   item.Department.Person.LastWebAddress,
		"TaxAdmin":                  item.Department.Person.TaxAdministration,
		"TaxNo":                     item.Department.Person.IdentityNumber,
		"ChamberInfo":               item.Department.ChamberInfo,
		"ChamberRegisterNo":         item.Department.RegistrationNumber,
		"VisaDate":                  visaDate,
		"VisaEndDate":               visaEndDate,
		"visa_finished_for_90days":  item.Department.VisaFinishedFor90Days,
		"core_person_absent_90days": item.Department.CorePersonAbsent90Days,
		"isClosed":                  item.Department.IsClosed,
		"OwnerName":                 item.Person.FullName,
		"OwnerTcNo":                 item.Person.IdentityNumber,
		"OwnerAddress":              item.Person.AddressStr,
		"OwnerPhone":                item.Person.LastPhoneNumber,
		"OwnerEmail":                item.Person.LastEPosta,
		"OwnerRegisterNo":           item.OccupationalRegistrationNumber,
		"OwnerCareer":               item.Title.Name,
	}

	// Şirket detaylarını al
	detailQuery := `
	query GetCompanyDetail($companyCode: Int!) {
		companyByCode(companyCode: $companyCode) {
			id
			CompanyCode
			Name
			Address
			Phone
			Fax
			MobilePhone
			Email
			Website
			TaxAdmin
			TaxNo
			ChamberInfo
			ChamberRegisterNo
			VisaDate
			VisaEndDate
			visa_finished_for_90days
			core_person_absent_90days
			isClosed
			OwnerName
			OwnerTcNo
			OwnerAddress
			OwnerPhone
			OwnerEmail
			OwnerRegisterNo
			OwnerCareer
		}
	}
	`

	checkVariables := map[string]interface{}{
		"companyCode": item.Department.FileNumber,
	}

	var detailResult struct {
		CompanyByCode map[string]interface{} `json:"companyByCode"`
	}

	err = graphqlClient.Execute(detailQuery, checkVariables, jwtToken, &detailResult)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Şirket detayları alınırken hata oluştu: %v", err),
		})
	}

	// Değişiklik var mı kontrol et
	needsUpdate := false

	// Değerleri karşılaştır ve farklılıkları logla
	for key, newValue := range companyData {
		if currentValue, exists := detailResult.CompanyByCode[key]; exists {
			if changed, logMessage := service.CompareFieldValues(key, currentValue, newValue); changed {
				needsUpdate = true
				log.Printf("Değişiklik: %s", logMessage)
			}
		}
	}

	if !needsUpdate {
		result := fiber.Map{
			"status":  "success",
			"message": "Şirket bilgileri güncel, güncelleme gerekmedi",
		}
		c.Locals("response", result)
		return c.JSON(result)
	}

	log.Printf("Değişiklik tespit edildi, şirket güncelleniyor...")

	// Değişiklik varsa güncelle
	updateMutation := `
	mutation UpdateCompany($input: CompanyDetailInput!) {
		updateCompany(input: $input) {
			id
			CompanyCode
			Name
		}
	}
	`

	var updateResult struct {
		UpdateCompany struct {
			ID          int    `json:"id"`
			CompanyCode int    `json:"CompanyCode"`
			Name        string `json:"Name"`
		} `json:"updateCompany"`
	}

	if err := graphqlClient.Execute(updateMutation, map[string]interface{}{"input": companyData}, jwtToken, &updateResult); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Şirket güncellenirken hata oluştu: %v", err),
		})
	}

	result := fiber.Map{
		"message": "Şirket başarıyla güncellendi",
		"company": fiber.Map{
			"name": item.Department.Name,
			"code": item.Department.FileNumber,
		},
	}
	c.Locals("response", result)
	return c.JSON(result)
}
