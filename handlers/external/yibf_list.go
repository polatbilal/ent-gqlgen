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

	// Test için 4 kayıt al
	requestBody := map[string]interface{}{
		"requireTotalCount": true,
		"searchOperation":   "contains",
		"searchValue":       nil,
		"skip":              20,
		"take":              4, // Test için 4 kayıt
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

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	graphqlClient := client.GraphQLClient{
		URL: fmt.Sprintf("%s://%s/graphql", scheme, c.Request.Host),
	}

	var processedIDs []int
	var failedIDs []int

	for _, item := range response.Items {
		log.Printf("ID %d işleniyor...", item.ID)

		// Önce mevcut kaydı kontrol et
		checkQuery := `
		query CheckJob($yibfNo: Int!) {
			job(yibfNo: $yibfNo) {
				id
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
			"yibfNo": item.ID,
		}

		var checkResult struct {
			Job map[string]interface{} `json:"job"`
		}

		err = graphqlClient.Execute(checkQuery, checkVariables, jwtToken, &checkResult)
		if err != nil && !strings.Contains(err.Error(), "not found") {
			log.Printf("ID %d için iş kontrolü sırasında hata oluştu: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

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

		var detail service.YIBFResponse
		if err := json.Unmarshal(detailBody, &detail); err != nil {
			log.Printf("ID %d için detay parse hatası: %v", item.ID, err)
			failedIDs = append(failedIDs, item.ID)
			continue
		}

		// Input verilerini hazırla
		jobData := map[string]interface{}{
			"YibfNo":      detail.ID,
			"CompanyCode": detail.YDK.FileNumber,
			"Title":       detail.Title,
		}

		var ownerData, contractorData, supervisorData map[string]interface{}

		// Owner verilerini hazırla
		if detail.YibfOwner.Person.ID > 0 {
			ownerData = map[string]interface{}{
				"YibfNo": detail.ID,
				"YDSID":  detail.YibfOwner.Person.ID,
				"Name":   detail.YibfOwner.Person.FullName,
				"TcNo":   detail.YibfOwner.Person.IdentityNumber,
				"Phone":  detail.YibfOwner.Person.LastPhoneNumber,
			}
		}

		// Contractor verilerini hazırla
		if detail.LatestYibfYambis.Yambis.ID > 0 {
			contractorData = map[string]interface{}{
				"YibfNo":     detail.ID,
				"YDSID":      detail.LatestYibfYambis.Yambis.ID,
				"Name":       detail.LatestYibfYambis.Yambis.AdSoyadUnvan,
				"TaxNo":      detail.LatestYibfYambis.Yambis.VergiKimlikNo,
				"Phone":      detail.LatestYibfYambis.Yambis.Telefon,
				"Email":      detail.LatestYibfYambis.Yambis.Eposta,
				"Address":    detail.LatestYibfYambis.Yambis.Adres,
				"RegisterNo": detail.LatestYibfYambis.Yambis.Ybno,
			}
		}

		// Supervisor verilerini hazırla
		if detail.LatestYibfSiteSupervisor.Application.User.ID > 0 {
			tcno := 0
			if tcStr := detail.LatestYibfSiteSupervisor.Application.User.Person.IdentityNumber; tcStr != "" {
				if tcInt, err := strconv.Atoi(tcStr); err == nil {
					tcno = tcInt
				}
			}

			registerNo := 0
			if registerNoStr := detail.LatestYibfSiteSupervisor.Application.OccupationalRegistrationNumber; registerNoStr != "" {
				if registerNoInt, err := strconv.Atoi(registerNoStr); err == nil {
					registerNo = registerNoInt
				}
			}

			securityNo := 0
			if securityNoStr := detail.LatestYibfSiteSupervisor.Application.SocialSecurityNo; securityNoStr != "" {
				if securityNoInt, err := strconv.Atoi(securityNoStr); err == nil {
					securityNo = securityNoInt
				}
			}

			supervisorData = map[string]interface{}{
				"YibfNo":           detail.ID,
				"YDSID":            detail.LatestYibfSiteSupervisor.Application.User.ID,
				"Name":             detail.LatestYibfSiteSupervisor.Application.User.Person.FullName,
				"Address":          detail.LatestYibfSiteSupervisor.Application.User.Person.LastAddress,
				"Phone":            detail.LatestYibfSiteSupervisor.Application.User.Person.LastPhoneNumber,
				"Email":            detail.LatestYibfSiteSupervisor.Application.User.Person.LastEPosta,
				"TcNo":             tcno,
				"Position":         detail.LatestYibfSiteSupervisor.Application.Tasks.Name,
				"Career":           detail.LatestYibfSiteSupervisor.Application.Title.Name,
				"RegisterNo":       registerNo,
				"SocialSecurityNo": securityNo,
				"SchoolGraduation": detail.LatestYibfSiteSupervisor.Application.SchoolGraduation,
			}
		}

		if checkResult.Job != nil {
			// Kayıt var, değişiklik kontrolü yap
			needsUpdate := false

			// Değerleri karşılaştır ve farklılıkları logla
			for key, newValue := range jobData {
				if currentValue, exists := checkResult.Job[key]; exists {
					// Nil değerleri kontrol et
					if newValue == nil || currentValue == nil {
						continue
					}

					// String'e çevirip karşılaştır
					newStr := fmt.Sprintf("%v", newValue)
					currentStr := fmt.Sprintf("%v", currentValue)

					if strings.TrimSpace(newStr) != strings.TrimSpace(currentStr) {
						log.Printf("Değişiklik tespit edildi - Alan: %s, Eski: %v, Yeni: %v", key, currentValue, newValue)
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
				log.Printf("YİBF verileri güncel, güncelleme yapılmayacak: YibfNo %d", item.ID)
				processedIDs = append(processedIDs, item.ID)
				continue
			}

			// Değişiklik varsa güncelle
			updateMutation := `
			mutation UpdateJob($yibfNo: Int!, $input: JobInput!) {
				updateJob(YibfNo: $yibfNo, input: $input) {
					id
					YibfNo
				}
			}
			`

			updateVariables := map[string]interface{}{
				"yibfNo": item.ID,
				"input":  jobData,
			}

			var updateResult struct {
				UpdateJob struct {
					ID     int `json:"id"`
					YibfNo int `json:"YibfNo"`
				} `json:"updateJob"`
			}

			if err := graphqlClient.Execute(updateMutation, updateVariables, jwtToken, &updateResult); err != nil {
				log.Printf("ID %d için güncelleme hatası: %v", item.ID, err)
				failedIDs = append(failedIDs, item.ID)
				continue
			}

			log.Printf("ID %d başarıyla güncellendi", item.ID)
			processedIDs = append(processedIDs, item.ID)
		} else {
			// Kayıt yok, yeni kayıt oluştur
			var mutationParts []string
			createVariables := map[string]interface{}{
				"jobInput": jobData,
			}

			// Job mutation'ı her zaman ekle
			mutationParts = append(mutationParts, `
				createJob(input: $jobInput) {
					YibfNo
					Title
				}`)

			// Owner verisi varsa ekle
			if ownerData != nil {
				mutationParts = append(mutationParts, `
				createOwner(input: $ownerInput) {
					Name
					YDSID
				}`)
				createVariables["ownerInput"] = ownerData
			}

			// Contractor verisi varsa ekle
			if contractorData != nil {
				mutationParts = append(mutationParts, `
				createContractor(input: $contractorInput) {
					Name
					YDSID
				}`)
				createVariables["contractorInput"] = contractorData
			}

			// Supervisor verisi varsa ekle
			if supervisorData != nil {
				mutationParts = append(mutationParts, `
				createSupervisor(input: $supervisorInput) {
					Name
					YDSID
				}`)
				createVariables["supervisorInput"] = supervisorData
			}

			// Mutation parametrelerini oluştur
			var paramParts []string
			paramParts = append(paramParts, "$jobInput: JobInput!")
			if ownerData != nil {
				paramParts = append(paramParts, "$ownerInput: JobOwnerInput!")
			}
			if contractorData != nil {
				paramParts = append(paramParts, "$contractorInput: JobContractorInput!")
			}
			if supervisorData != nil {
				paramParts = append(paramParts, "$supervisorInput: JobSupervisorInput!")
			}

			// Tam mutation string'ini oluştur
			createMutation := fmt.Sprintf(`
			mutation CreateJob(%s) {
				%s
			}
			`, strings.Join(paramParts, ", "), strings.Join(mutationParts, "\n"))

			// Detaylı loglama
			// log.Printf("GraphQL Mutation: %s", createMutation)
			variablesJSON, _ := json.MarshalIndent(createVariables, "", "  ")
			log.Printf("GraphQL Variables: %s", string(variablesJSON))

			var createResult struct {
				CreateJob struct {
					YibfNo int    `json:"YibfNo"`
					Title  string `json:"Title"`
				} `json:"createJob"`
				CreateOwner *struct {
					Name  string `json:"Name"`
					YDSID int    `json:"YDSID"`
				} `json:"createOwner,omitempty"`
				CreateContractor *struct {
					Name  string `json:"Name"`
					YDSID int    `json:"YDSID"`
				} `json:"createContractor,omitempty"`
				CreateSupervisor *struct {
					Name  string `json:"Name"`
					YDSID int    `json:"YDSID"`
				} `json:"createSupervisor,omitempty"`
			}

			if err := graphqlClient.Execute(createMutation, createVariables, jwtToken, &createResult); err != nil {
				log.Printf("ID %d için oluşturma hatası: %v", item.ID, err)
				failedIDs = append(failedIDs, item.ID)
				continue
			}

			log.Printf("ID %d başarıyla oluşturuldu", item.ID)
			processedIDs = append(processedIDs, item.ID)
		}

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
