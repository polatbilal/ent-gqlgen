package handlers

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

	"github.com/gin-gonic/gin"
)

func YDKCompanies(c *gin.Context) {
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

	// YDK token'ı JSON'dan parse et
	var ydkToken YDKTokenResponse
	if err := json.Unmarshal([]byte(ydkTokenStr), &ydkToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "YDK Token parse hatası: " + err.Error()})
		return
	}

	service := &ExternalService{
		baseURL: os.Getenv("YDK_BASE_URL"),
		client:  &http.Client{},
	}

	if service.baseURL == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "YDK_BASE_URL environment variable is not set"})
		return
	}

	requestBody := map[string]interface{}{
		"id": ydkToken.DepartmentID,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body oluşturma hatası: " + err.Error()})
		return
	}

	url := service.baseURL + ENDPOINT_BY_DEPARTMENT
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ydkToken.AccessToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := service.client.Do(req)
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
		log.Printf("YDK API hata döndürdü. Status: %d, Body: %s\n", resp.StatusCode, string(body))
		c.JSON(resp.StatusCode, gin.H{"error": fmt.Sprintf("YDK API hatası: %s", string(body))})
		return
	}

	// Ham veriyi logla
	// log.Printf("YDK API Ham Yanıt: %s\n", string(body))

	// Struct'a parse et
	var ydkResponse YDKCompanyResponse
	if err := json.Unmarshal(body, &ydkResponse); err != nil {
		log.Printf("JSON parse hatası: %v\nHam veri: %s\n", err, string(body))
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON parse hatası: " + err.Error()})
		return
	}

	if len(ydkResponse.Items) == 0 {
		log.Printf("ydkToken.AccessToken: %s", ydkToken.AccessToken)
		log.Printf("ydkToken.DepartmentID: %d", ydkToken.DepartmentID)
		log.Printf("Şirket bilgisi bulunamadı. Ham yanıt: %s\n", string(body))
		c.JSON(http.StatusNotFound, gin.H{"error": "Şirket bilgisi bulunamadı"})
		return
	}

	// İlk şirketi al
	item := ydkResponse.Items[0]

	// Unix timestamp'i tarihe çevir
	visaDate := time.Unix(item.Department.VisaDate/1000, 0).Local().Format("2006-01-02")
	visaEndDate := time.Unix(item.Department.VisaEndDate/1000, 0).Local().Format("2006-01-02")

	// GraphQL client oluştur
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	graphqlClient := GraphQLClient{
		URL: fmt.Sprintf("%s://%s/graphql", scheme, c.Request.Host),
	}

	// Önce şirketin var olup olmadığını kontrol et
	checkQuery := `
	query CheckCompany($companyCode: Int!) {
		companyByCode(companyCode: $companyCode) {
			id
			CompanyCode
			Name
		}
	}
	`

	checkVariables := map[string]interface{}{
		"companyCode": item.Department.FileNumber,
	}

	var checkResult struct {
		CompanyByCode struct {
			ID          int    `json:"id"`
			CompanyCode int    `json:"CompanyCode"`
			Name        string `json:"Name"`
		} `json:"companyByCode"`
	}

	err = graphqlClient.Execute(checkQuery, checkVariables, jwtToken, &checkResult)

	// Şirket verilerini hazırla
	companyData := map[string]interface{}{
		"Name":                      item.Department.Name,
		"CompanyCode":               item.Department.FileNumber,
		"Address":                   item.Department.Person.AddressStr,
		"Phone":                     item.Department.Person.LastPhoneNumber,
		"Email":                     item.Department.Person.LastEPosta,
		"Website":                   item.Department.Person.LastWebAddress,
		"TaxAdmin":                  item.Department.Person.TaxAdministration,
		"TaxNo":                     item.Department.Person.IdentityNumber,
		"ChamberInfo":               item.Department.ChamberInfo,
		"ChamberRegNo":              item.Department.RegistrationNumber,
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
		"OwnerRegNo":                item.OccupationalRegistrationNumber,
		"OwnerBirthDate":            item.Person.BirthDateString,
		"OwnerCareer":               item.Title.Name,
	}

	// Şirket var mı kontrolü
	if err == nil && checkResult.CompanyByCode.CompanyCode > 0 {
		// Mevcut şirket detaylarını al
		detailQuery := `
		query GetCompanyDetail($companyCode: Int!) {
			companyByCode(companyCode: $companyCode) {
				id
				CompanyCode
				Name
				Address
				Phone
				Email
				Website
				TaxAdmin
				TaxNo
				ChamberInfo
				ChamberRegNo
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
				OwnerRegNo
				OwnerBirthDate
				OwnerCareer
			}
		}
		`

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

		// Sayısal alanların listesi
		numericFields := map[string]bool{
			"CompanyCode": true,
			"TaxNo":       true,
		}

		// Değerleri karşılaştır ve farklılıkları logla
		for key, newValue := range companyData {
			if currentValue, exists := detailResult.CompanyByCode[key]; exists {
				// Nil değerleri kontrol et
				if newValue == nil && currentValue == nil {
					continue
				}

				// Sayısal alan kontrolü
				if numericFields[key] {
					// Sayısal değerlere çevir
					var newFloat, currentFloat float64

					switch v := newValue.(type) {
					case int:
						newFloat = float64(v)
					case float64:
						newFloat = v
					case string:
						newFloat, _ = strconv.ParseFloat(v, 64)
					}

					switch v := currentValue.(type) {
					case int:
						currentFloat = float64(v)
					case float64:
						currentFloat = v
					case string:
						currentFloat, _ = strconv.ParseFloat(v, 64)
					}

					// Sayısal karşılaştırma
					if math.Abs(newFloat-currentFloat) > 0.0001 { // küçük farkları tolere et
						log.Printf("Değişiklik tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentValue, newValue)
						needsUpdate = true
					}
					continue
				}

				// String'e çevirip karşılaştır (sayısal olmayan alanlar için)
				newStr := fmt.Sprintf("%v", newValue)
				currentStr := fmt.Sprintf("%v", currentValue)

				if newStr != currentStr {
					log.Printf("Değişiklik tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentStr, newStr)
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
			log.Printf("Şirket verileri güncel, güncelleme yapılmayacak")
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
		return
	}

	// Eğer hata varsa ve bu "not found" hatası değilse
	if err != nil && !strings.Contains(err.Error(), "not found") {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Şirket kontrolü sırasında hata oluştu: %v", err),
		})
		return
	}

	// Şirket yoksa, yeni kayıt oluştur
	mutation := `
	mutation CreateCompany($input: CompanyDetailInput!) {
		createCompany(input: $input) {
			id
			CompanyCode
			Name
		}
	}
	`

	variables := map[string]interface{}{
		"input": map[string]interface{}{
			"Name":                      item.Department.Name,
			"CompanyCode":               item.Department.FileNumber,
			"Address":                   item.Department.Person.AddressStr,
			"Phone":                     item.Department.Person.LastPhoneNumber,
			"Email":                     item.Department.Person.LastEPosta,
			"Website":                   item.Department.Person.LastWebAddress,
			"TaxAdmin":                  item.Department.Person.TaxAdministration,
			"TaxNo":                     item.Department.Person.IdentityNumber,
			"ChamberInfo":               item.Department.ChamberInfo,
			"ChamberRegNo":              item.Department.RegistrationNumber,
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
			"OwnerRegNo":                item.OccupationalRegistrationNumber,
			"OwnerBirthDate":            item.Person.BirthDateString,
			"OwnerCareer":               item.Title.Name,
		},
	}

	var result struct {
		CreateCompany struct {
			ID          int    `json:"id"`
			CompanyCode int    `json:"CompanyCode"`
			Name        string `json:"Name"`
		} `json:"createCompany"`
	}

	if err := graphqlClient.Execute(mutation, variables, jwtToken, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Şirket kaydedilirken hata oluştu: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Şirket başarıyla kaydedildi",
		"company": gin.H{
			"name": item.Department.Name,
			"code": item.Department.FileNumber,
		},
	})
}
