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

		// log.Printf("ID %d için detay yanıtı: %s", item.ID, string(detailBody))

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
		// Önce işin var olup olmadığını kontrol et
		checkQuery := `
		query CheckJob($yibfNo: Int!) {
			job(yibfNo: $yibfNo) {
				id
				YibfNo
				Idare
				Pafta
				Ada
				Parsel
				State
				ContractDate
				LicenseDate
				LicenseNo
				ConstructionArea
				LandArea
				Address
				BuildingClass
				BuildingType
				Floors
				Level
				Inspector {
					id
					YDSID
				}
				Static {
					id
					YDSID
				}
				Architect {
					id
					YDSID
				}
				Mechanic {
					id
					YDSID
				}
				Electric {
					id
					YDSID
				}
				Controller {
					id
					YDSID
				}
				MechanicController {
					id
					YDSID
				}
				ElectricController {
					id
					YDSID
				}
			}
		}
		`

		checkVariables := map[string]interface{}{
			"yibfNo": detail.ID,
		}

		var checkResult struct {
			Job map[string]interface{} `json:"job"`
		}

		err = graphqlClient.Execute(checkQuery, checkVariables, jwtToken, &checkResult)
		if err != nil {
			log.Printf("İş kontrolü sırasında hata oluştu: %v", err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		// Input verilerini hazırla
		jobData := map[string]interface{}{
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
				YibfId  int `json:"yibfId"`
				UserId  int `json:"userId"`
				TaskId  int `json:"taskId"`
				TitleId int `json:"titleId"`
			} `json:"items"`
		}

		if err := json.Unmarshal(engineerBody, &engineerResponse); err != nil {
			log.Printf("Mühendis yanıtı parse hatası: %v", err)
			continue
		}

		// Her rol için ilgili mühendisi bul ve ata
		for _, eng := range engineerResponse.Items {
			// Inspector: taskId = 1 veya (taskId = 2 ve titleId = 4)
			if eng.TaskId == 1 || (eng.TaskId == 2 && eng.TitleId == 4) {
				jobData["Inspector"] = eng.UserId
			}

			// Static: taskId = 2 ve titleId = 4
			if eng.TaskId == 2 && eng.TitleId == 4 {
				jobData["Static"] = eng.UserId
			}

			// Architect: taskId = 2 ve titleId = 3
			if eng.TaskId == 2 && eng.TitleId == 3 {
				jobData["Architect"] = eng.UserId
			}

			// Mechanic: taskId = 2 ve titleId = 2
			if eng.TaskId == 2 && eng.TitleId == 2 {
				jobData["Mechanic"] = eng.UserId
			}

			// Electric: taskId = 2 ve titleId = 1
			if eng.TaskId == 2 && eng.TitleId == 1 {
				jobData["Electric"] = eng.UserId
			}

			// Controller: taskId = 14 ve (titleId = 3 veya titleId = 4)
			if eng.TaskId == 14 && (eng.TitleId == 3 || eng.TitleId == 4) {
				jobData["Controller"] = eng.UserId
			}

			// MechanicController: taskId = 14 ve titleId = 2
			if eng.TaskId == 14 && eng.TitleId == 2 {
				jobData["MechanicController"] = eng.UserId
			}

			// ElectricController: taskId = 14 ve titleId = 1
			if eng.TaskId == 14 && eng.TitleId == 1 {
				jobData["ElectricController"] = eng.UserId
			}
		}

		// Supervisor bilgilerini kontrol et ve ekle
		if detail.LatestYibfSiteSupervisor.Application.Person.ID > 0 {
			person := detail.LatestYibfSiteSupervisor.Application.Person
			supervisor := detail.LatestYibfSiteSupervisor.Application

			if supervisor.Tasks.Name != "" && supervisor.Title.Name != "" {
				jobData["Supervisor"] = map[string]interface{}{
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
			jobData["Owner"] = map[string]interface{}{
				"YdsId": detail.YibfOwner.Person.ID,
				"Name":  detail.Title,
				"TcNo":  detail.YibfOwner.Person.IdentityNumber,
				"Phone": detail.YibfOwner.Person.LastPhoneNumber,
			}
		}

		// Contractor bilgilerini kontrol et ve ekle
		if detail.LatestYibfYambis.Yambis.ID > 0 {
			yambis := detail.LatestYibfYambis.Yambis
			jobData["Contractor"] = map[string]interface{}{
				"YdsId":      yambis.ID,
				"Name":       yambis.AdSoyadUnvan,
				"TaxNo":      yambis.VergiKimlikNo,
				"Phone":      yambis.Telefon,
				"Email":      yambis.Eposta,
				"Address":    yambis.Adres,
				"RegisterNo": yambis.Ybno,
			}
		}

		var mutation string
		var variables map[string]interface{}

		if checkResult.Job != nil {
			// İş bulundu, değişiklik kontrolü yap
			needsUpdate := false
			existingJob := checkResult.Job

			// Değerleri karşılaştır ve farklılıkları logla
			for key, newValue := range jobData {
				if currentValue, exists := existingJob[key]; exists {
					// Nil değerleri kontrol et
					if newValue == nil && currentValue == nil {
						continue
					}

					// String'e çevirip karşılaştır
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
				log.Printf("İş verileri güncel, güncelleme yapılmayacak: YibfNo %d", detail.ID)
				processedIDs = append(processedIDs, item.ID)
				continue
			}

			// Değişiklik varsa güncelle
			mutation = `
			mutation UpdateJob($yibfNo: Int!, $input: JobInput!) {
				updateJob(YibfNo: $yibfNo, input: $input) {
					id
					YibfNo
				}
			}
			`
			variables = map[string]interface{}{
				"yibfNo": detail.ID,
				"input":  jobData,
			}
		} else {
			// Yeni kayıt oluştur
			mutation = `
			mutation CreateJob($input: JobInput!) {
				createJob(input: $input) {
					id
					YibfNo
				}
			}
			`
			variables = map[string]interface{}{
				"input": jobData,
			}
		}

		var result struct {
			CreateJob *struct {
				ID     int `json:"id"`
				YibfNo int `json:"YibfNo"`
			} `json:"createJob"`
			UpdateJob *struct {
				ID     int `json:"id"`
				YibfNo int `json:"YibfNo"`
			} `json:"updateJob"`
		}

		err = graphqlClient.Execute(mutation, variables, jwtToken, &result)
		if err != nil {
			log.Printf("ID %d için GraphQL mutation hatası: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		processedIDs = append(processedIDs, item.ID)
		log.Printf("ID %d başarıyla işlendi", item.ID)

		// Yapı Bilgileri loglaması
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
