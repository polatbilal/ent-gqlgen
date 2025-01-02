package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

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

	// YDK API için token
	ydkTokenStr := c.GetHeader("YDK-Token")
	if ydkTokenStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "YDK Token gerekli"})
		return
	}

	// YDK token'ı JSON'dan parse et
	var ydkToken service.YDKTokenResponse
	if err := json.Unmarshal([]byte(ydkTokenStr), &ydkToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "YDK Token parse hatası: " + err.Error()})
		return
	}

	// ID parametresini al
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parametresi gerekli"})
		return
	}

	svc := &service.ExternalService{
		BaseURL: os.Getenv("YDK_BASE_URL"),
		Client:  &http.Client{},
	}

	// YDK API'den detay bilgilerini çek
	url := fmt.Sprintf("%s%s%s", svc.BaseURL, service.ENDPOINT_YIBF_BY_ID, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ydkToken.AccessToken))

	resp, err := svc.Client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var detail service.YIBFResponse
	if err := json.Unmarshal(body, &detail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parse hatası: " + err.Error()})
		return
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
		JobDetail struct {
			ID     string `json:"id"`
			YibfNo int    `json:"YibfNo"`
		} `json:"job"`
	}

	err = graphqlClient.Execute(checkQuery, checkVariables, jwtToken, &checkResult)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "GraphQL hatası: " + err.Error(),
		})
		return
	}

	// CreateJob mutation'ı için input hazırla
	input := map[string]interface{}{
		"input": map[string]interface{}{
			"YibfNo":           detail.ID,
			"CompanyCode":      detail.YDK.FileNumber,
			"Idare":            detail.Administration.Name,
			"Pafta":            detail.Sheet,
			"Ada":              detail.Island,
			"Parsel":           detail.Parcel,
			"State":            detail.State.Name,
			"ContractDate":     time.Unix(detail.ContractDate, 0).Local().Format("2006-01-02"),
			"LicenseDate":      time.Unix(detail.LicenseDate, 0).Local().Format("2006-01-02"),
			"LicenseNo":        detail.LicenseNumber,
			"ConstructionArea": detail.YibfStructure.ConstructionArea,
			"LandArea":         detail.YibfStructure.LeftArea,
			"Address":          detail.YibfStructure.BuildingAddress,
			"BuildingClass":    detail.YibfStructure.BuildingClass.Name,
			"BuildingType":     detail.YibfStructure.CarrierSystemType.Name,
			"Floors":           detail.YibfStructure.FloorCount,
			"Level":            detail.Level,
			"Coordinates":      detail.Position.Coordinates,
			"Supervisor": map[string]interface{}{
				"YdsId":            detail.LatestYibfSiteSupervisor.Application.Person.ID,
				"Name":             detail.LatestYibfSiteSupervisor.Application.Person.FullName,
				"Address":          detail.LatestYibfSiteSupervisor.Application.Person.LastAddress,
				"Phone":            detail.LatestYibfSiteSupervisor.Application.Person.LastPhoneNumber,
				"Email":            detail.LatestYibfSiteSupervisor.Application.Person.LastEPosta,
				"Tcno":             detail.LatestYibfSiteSupervisor.Application.Person.IdentityNumber,
				"Position":         detail.LatestYibfSiteSupervisor.Application.Tasks.Name,
				"Career":           detail.LatestYibfSiteSupervisor.Application.Title.Name,
				"RegNo":            detail.LatestYibfSiteSupervisor.Application.OccupationalRegistrationNumber,
				"SocialSecurityNo": detail.LatestYibfSiteSupervisor.Application.SocialSecurityNo,
				"SchoolGraduation": detail.LatestYibfSiteSupervisor.Application.SchoolGraduation,
			},
			"Owner": []map[string]interface{}{
				{
					"YdsId": detail.YibfOwner.Person.ID,
					"Name":  detail.Title,
					"TcNo":  detail.YibfOwner.Person.IdentityNumber,
					"Phone": detail.YibfOwner.Person.LastPhoneNumber,
				},
			},
			"Contractor": []map[string]interface{}{
				{
					"YdsId":      detail.LatestYibfYambis.Yambis.ID,
					"Name":       detail.LatestYibfYambis.Yambis.AdSoyadUnvan,
					"TaxNo":      detail.LatestYibfYambis.Yambis.VergiKimlikNo,
					"Phone":      detail.LatestYibfYambis.Yambis.Telefon,
					"Email":      detail.LatestYibfYambis.Yambis.Eposta,
					"Address":    detail.LatestYibfYambis.Yambis.Adres,
					"RegisterNo": detail.LatestYibfYambis.Yambis.Ybno,
				},
			},
		},
	}

	// CreateJob mutation'ını çalıştır
	mutation := `
		mutation CreateJob($input: JobInput!) {
			createJob(input: $input) {
				id
				YibfNo
			}
		}
	`

	var result struct {
		CreateJob struct {
			ID     string `json:"id"`
			YibfNo int    `json:"YibfNo"`
		} `json:"createJob"`
	}

	err = graphqlClient.Execute(mutation, input, jwtToken, &result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "GraphQL hatası: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"ydk_detail": detail,
			"db_result":  result,
		},
	})
}
