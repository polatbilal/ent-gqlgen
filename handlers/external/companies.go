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

	"github.com/polatbilal/gqlgen-ent/handlers/client"
	"github.com/polatbilal/gqlgen-ent/handlers/service"

	"github.com/gin-gonic/gin"
)

func YDKCompanies(c *gin.Context) {
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

	log.Printf("Şirket bilgileri alınıyor... DepartmentID: %d", companyToken.DepartmentId)

	svc := &service.ExternalService{
		BaseURL: os.Getenv("YDK_BASE_URL"),
		Client:  &http.Client{},
	}

	if svc.BaseURL == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "YDK_BASE_URL environment variable is not set"})
		return
	}

	requestBody := map[string]interface{}{
		"id": companyToken.DepartmentId,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body oluşturma hatası: " + err.Error()})
		return
	}

	url := svc.BaseURL + service.ENDPOINT_COMPANY
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", companyToken.Token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := svc.Client.Do(req)
	if err != nil {
		log.Printf("YDK API isteği başarısız: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "YDK API isteği başarısız: " + err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Yanıt body okuma hatası: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Yanıt okuma hatası: " + err.Error()})
		return
	}

	// HTTP yanıt kodunu kontrol et
	if resp.StatusCode != http.StatusOK {
		var ydkError struct {
			Error            string `json:"error"`
			ErrorDescription string `json:"error_description"`
		}

		if err := json.Unmarshal(body, &ydkError); err != nil {
			c.JSON(resp.StatusCode, gin.H{"error": fmt.Sprintf("YDK API hatası: %s", string(body))})
			return
		}

		// Hata yanıtını doğrudan ilet
		c.JSON(resp.StatusCode, ydkError)
		return
	}

	// Struct'a parse et
	var ydkResponse service.YDKCompanyResponse
	if err := json.Unmarshal(body, &ydkResponse); err != nil {
		log.Printf("JSON parse hatası: %v\nHam veri: %s\n", err, string(body))
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON parse hatası: " + err.Error()})
		return
	}

	if len(ydkResponse.Items) == 0 {
		log.Printf("ydkToken.DepartmentID: %d", companyToken.DepartmentId)
		log.Printf("Şirket bilgisi bulunamadı. Ham yanıt: %s\n", string(body))
		c.JSON(http.StatusNotFound, gin.H{"error": "Şirket bilgisi bulunamadı"})
		return
	}

	// İlk şirketi al
	item := ydkResponse.Items[0]

	// Unix timestamp'i tarihe çevir
	visaDate := service.SafeUnixToDate(item.Department.VisaDate)
	visaEndDate := service.SafeUnixToDate(item.Department.VisaEndDate)

	// GraphQL client oluştur
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	graphqlClient := client.GraphQLClient{
		URL: fmt.Sprintf("%s://%s/graphql", scheme, c.Request.Host),
	}

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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Şirket detayları alınırken hata oluştu: %v", err),
		})
		return
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
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Şirket bilgileri güncel, güncelleme gerekmedi",
		})
		return
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Şirket güncellenirken hata oluştu: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Şirket başarıyla güncellendi",
		"company": gin.H{
			"name": item.Department.Name,
			"code": item.Department.FileNumber,
		},
	})
}
