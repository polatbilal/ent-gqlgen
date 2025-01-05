package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/polatbilal/gqlgen-ent/handlers/client"
	"github.com/polatbilal/gqlgen-ent/handlers/service"

	"github.com/gin-gonic/gin"
)

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

	// Varsayılan olarak tüm kayıtları al
	requestBody := map[string]interface{}{
		"requireTotalCount": true,
		"searchOperation":   "contains",
		"searchValue":       nil,
		"skip":              0,
		"take":              3, // Test için 3 kayıt
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

	// log.Printf("FindAll yanıtı (ham veri): %s", string(body))

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
	for i, item := range response.Items {
		log.Printf("ID[%d]: %d", i, item.ID)
	}

	// Her ID için detay bilgilerini çek ve veritabanına yaz
	var processedIDs []int
	var failedIDs []int

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	graphqlClient := client.GraphQLClient{
		URL: fmt.Sprintf("%s://%s/graphql", scheme, c.Request.Host),
	}

	for _, item := range response.Items {
		log.Printf("ID %d için detay bilgileri çekiliyor...", item.ID)

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

		log.Printf("ID %d için detay yanıtı: %s", item.ID, string(detailBody))

		var detail service.YIBFResponse
		if err := json.Unmarshal(detailBody, &detail); err != nil {
			log.Printf("ID %d için detay parse hatası: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		log.Printf("ID %d için detay bilgileri başarıyla parse edildi", item.ID)
		log.Printf("Yapı Bilgileri - YibfNo: %d, İdare: %s, Pafta: %s, Ada: %s, Parsel: %s",
			detail.ID,
			detail.Administration.Name,
			detail.Sheet,
			detail.Island,
			detail.Parcel)

		// GraphQL mutation'ı için input hazırla
		mutation := `
			mutation CreateJob($input: JobInput!) {
				createJob(input: $input) {
					id
					YibfNo
				}
			}
		`

		// Input verilerini hazırla
		input := map[string]interface{}{
			"input": map[string]interface{}{
				"YibfNo":           detail.ID,
				"CompanyCode":      detail.YDK.FileNumber,
				"Idare":            detail.Administration.Name,
				"Pafta":            detail.Sheet,
				"Ada":              detail.Island,
				"Parsel":           detail.Parcel,
				"State":            detail.State.Name,
				"ContractDate":     time.Unix(detail.ContractDate, 0).Format("2006-01-02"),
				"LicenseDate":      time.Unix(detail.LicenseDate, 0).Format("2006-01-02"),
				"LicenseNo":        detail.LicenseNumber,
				"ConstructionArea": fmt.Sprintf("%.2f", detail.YibfStructure.ConstructionArea),
				"LandArea":         fmt.Sprintf("%.2f", detail.YibfStructure.LeftArea),
				"Address":          detail.YibfStructure.BuildingAddress,
				"BuildingClass":    detail.YibfStructure.BuildingClass.Name,
				"BuildingType":     detail.YibfStructure.CarrierSystemType.Name,
				"Floors":           detail.YibfStructure.FloorCount,
				"Level":            detail.Level,
			},
		}

		// Supervisor bilgilerini kontrol et ve ekle
		if detail.LatestYibfSiteSupervisor.Application.Person.ID > 0 {
			person := detail.LatestYibfSiteSupervisor.Application.Person
			supervisor := detail.LatestYibfSiteSupervisor.Application

			input["input"].(map[string]interface{})["Inspector"] = person.ID
			input["input"].(map[string]interface{})["Static"] = person.ID
			input["input"].(map[string]interface{})["Architect"] = person.ID
			input["input"].(map[string]interface{})["Mechanic"] = person.ID
			input["input"].(map[string]interface{})["Electric"] = person.ID
			input["input"].(map[string]interface{})["Controller"] = person.ID
			input["input"].(map[string]interface{})["MechanicController"] = person.ID
			input["input"].(map[string]interface{})["ElectricController"] = person.ID

			if supervisor.Tasks.Name != "" && supervisor.Title.Name != "" {
				input["input"].(map[string]interface{})["Supervisor"] = map[string]interface{}{
					"YdsId":            person.ID,
					"Name":             person.FullName,
					"Address":          person.LastAddress,
					"Phone":            person.LastPhoneNumber,
					"Email":            person.LastEPosta,
					"Tcno":             person.IdentityNumber,
					"Position":         supervisor.Tasks.Name,
					"Career":           supervisor.Title.Name,
					"RegNo":            supervisor.OccupationalRegistrationNumber,
					"SocialSecurityNo": supervisor.SocialSecurityNo,
					"SchoolGraduation": supervisor.SchoolGraduation,
				}
			}
		}

		// Owner bilgilerini kontrol et ve ekle
		if detail.YibfOwner.Person.ID > 0 {
			input["input"].(map[string]interface{})["Owner"] = map[string]interface{}{
				"YdsId": detail.YibfOwner.Person.ID,
				"Name":  detail.Title,
				"TcNo":  detail.YibfOwner.Person.IdentityNumber,
				"Phone": detail.YibfOwner.Person.LastPhoneNumber,
			}
		}

		// Contractor bilgilerini kontrol et ve ekle
		if detail.LatestYibfYambis.Yambis.ID > 0 {
			yambis := detail.LatestYibfYambis.Yambis
			input["input"].(map[string]interface{})["Contractor"] = map[string]interface{}{
				"YdsId":      yambis.ID,
				"Name":       yambis.AdSoyadUnvan,
				"TaxNo":      yambis.VergiKimlikNo,
				"Phone":      yambis.Telefon,
				"Email":      yambis.Eposta,
				"Address":    yambis.Adres,
				"RegisterNo": yambis.Ybno,
			}
		}

		// Mühendis bilgilerini çek ve ekle
		engineerRequestBody := map[string]interface{}{
			"searchOperation": "contains",
			"searchValue":     nil,
			"userData":        struct{}{},
			"filter": [][]interface{}{
				{"yibfId", "=", item.ID},
			},
		}

		engineerJsonBody, err := json.Marshal(engineerRequestBody)
		if err != nil {
			log.Printf("Mühendis request body oluşturma hatası: %v", err)
			continue
		}

		engineerReq, err := http.NewRequest("POST", svc.BaseURL+service.ENDPOINT_YIBF_ENGINEER, bytes.NewBuffer(engineerJsonBody))
		if err != nil {
			log.Printf("Mühendis request oluşturma hatası: %v", err)
			continue
		}

		engineerReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ydkToken.AccessToken))
		engineerReq.Header.Set("Content-Type", "application/json")

		engineerResp, err := svc.Client.Do(engineerReq)
		if err != nil {
			log.Printf("Mühendis bilgileri isteği başarısız: %v", err)
			continue
		}
		defer engineerResp.Body.Close()

		engineerBody, err := io.ReadAll(engineerResp.Body)
		if err != nil {
			log.Printf("Mühendis yanıtı okuma hatası: %v", err)
			continue
		}

		var engineerResponse struct {
			Items []struct {
				ID int `json:"id"`
			} `json:"items"`
		}

		if err := json.Unmarshal(engineerBody, &engineerResponse); err != nil {
			log.Printf("Mühendis yanıtı parse hatası: %v", err)
			continue
		}

		// Eğer mühendis varsa, ilk mühendisin ID'sini al
		if len(engineerResponse.Items) > 0 {
			input["input"].(map[string]interface{})["Engineer"] = engineerResponse.Items[0].ID
		}

		log.Printf("GraphQL mutation: %s", mutation)
		inputJSON, _ := json.MarshalIndent(input, "", "  ")
		log.Printf("GraphQL input: %s", string(inputJSON))

		var result struct {
			CreateJob struct {
				ID     string `json:"id"`
				YibfNo int    `json:"YibfNo"`
			} `json:"createJob"`
		}

		err = graphqlClient.Execute(mutation, input, jwtToken, &result)
		if err != nil {
			log.Printf("ID %d için GraphQL mutation hatası: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		processedIDs = append(processedIDs, item.ID)
		log.Printf("ID %d başarıyla işlendi", item.ID)

		// Yapı Bilgileri loglaması devam ediyor...
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
