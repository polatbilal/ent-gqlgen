package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/polatbilal/gqlgen-ent/handlers-module/client"
	"github.com/polatbilal/gqlgen-ent/handlers-module/service"
)

func ProgressPayments(c *fiber.Ctx) error {
	// GraphQL için JWT token
	jwtToken := c.Get("Authorization")
	if jwtToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT Token gerekli"})
	}

	// Request body'den parametreleri al
	var requestParams service.FrontendRequest
	if err := c.BodyParser(&requestParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz request body: " + err.Error(),
		})
	}

	// Parametreleri kontrol et
	if requestParams.YibfNo == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "yibfNo parametresi gerekli"})
	}

	// İş kaydından şirket kodunu al
	companyCode, err := service.GetCompanyCodeFromYibf(c.Context(), requestParams.YibfNo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	companyToken, err := service.GetCompanyTokenFromDB(c.Context(), companyCode)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if companyToken.Token == "" || companyToken.DepartmentId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçerli token veya department ID bulunamadı"})
	}

	// API'ye istek at
	svc := &service.ExternalService{
		BaseURL: os.Getenv("YDK_BASE_URL"),
		Client:  &http.Client{},
	}

	if svc.BaseURL == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "YDK_BASE_URL environment variable is not set"})
	}

	// Sabit request body'i oluştur
	requestBody := map[string]interface{}{
		"requireTotalCount": true,
		"searchOperation":   "contains",
		"searchValue":       nil,
		"skip":              0,
		"take":              100,
		"userData":          map[string]interface{}{},
		"sort": []map[string]interface{}{
			{
				"selector": "progressPaymentNumber",
				"desc":     true,
			},
		},
		"filter": []interface{}{
			"yibf.id",
			"=",
			float64(requestParams.YibfNo),
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request body oluşturma hatası: " + err.Error()})
	}

	url := fmt.Sprintf("%s%s", svc.BaseURL, service.ENDPOINT_YIBF_PAYMENT)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", companyToken.Token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := svc.Client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer resp.Body.Close()

	var progressPaymentResponse service.ProgressPaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&progressPaymentResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// GraphQL client oluştur
	scheme := "http"
	if c.Protocol() == "https" {
		scheme = "https"
	}
	graphqlClient := client.NewGraphQLClient(scheme)

	successCount := 0
	skippedCount := 0
	var processLogs []map[string]interface{}

	for _, payment := range progressPaymentResponse.Items {
		processLog := map[string]interface{}{
			"progressPaymentNumber": payment.ProgressPaymentNumber,
			"time":                  time.Now().Format("2006-01-02 15:04:05"),
		}

		// Unix timestamp'i tarihe çevir
		paymentDate := service.SafeUnixToDate(payment.ProgressPaymentDate)

		// Hakediş verilerini hazırla
		paymentData := map[string]interface{}{
			"PaymentNo":    payment.ProgressPaymentNumber,
			"PaymentDate":  paymentDate,
			"PaymentType":  payment.ProgressPaymentType.Name,
			"TotalPayment": payment.TotalPayment,
			"LevelRequest": payment.LevelRequest,
			"LevelApprove": payment.LevelApprove,
			"Amount":       payment.Amount,
			"State":        payment.State.Name,
			"yibfNo":       requestParams.YibfNo,
		}

		// Önce hakediş kaydını sorgula
		detailQuery := `
		query GetJobPayments($yibfNo: Int!) {
			jobPayments(yibfNo: $yibfNo) {
				id
				PaymentNo
				PaymentDate
				PaymentType
				TotalPayment
				LevelRequest
				LevelApprove
				Amount
				State
			}
		}
		`

		var detailResult struct {
			JobPayments []struct {
				ID           int     `json:"id"`
				PaymentNo    int     `json:"PaymentNo"`
				PaymentDate  string  `json:"PaymentDate"`
				PaymentType  string  `json:"PaymentType"`
				TotalPayment float64 `json:"TotalPayment"`
				LevelRequest float64 `json:"LevelRequest"`
				LevelApprove float64 `json:"LevelApprove"`
				Amount       float64 `json:"Amount"`
				State        string  `json:"State"`
			} `json:"jobPayments"`
		}

		err = graphqlClient.Execute(detailQuery, map[string]interface{}{"yibfNo": requestParams.YibfNo}, jwtToken, &detailResult)

		// Mevcut kayıt var mı kontrol et
		existingPayment := false
		if err == nil && len(detailResult.JobPayments) > 0 {
			for _, existing := range detailResult.JobPayments {
				if existing.PaymentNo == payment.ProgressPaymentNumber {
					existingPayment = true
					// Değişiklikleri kontrol et ve logla
					needsUpdate := false
					var changes []string

					if existing.LevelApprove != payment.LevelApprove {
						changes = append(changes, fmt.Sprintf("LevelApprove: %.2f -> %.2f", existing.LevelApprove, payment.LevelApprove))
						needsUpdate = true
					}
					if existing.TotalPayment != payment.TotalPayment {
						changes = append(changes, fmt.Sprintf("TotalPayment: %.2f -> %.2f", existing.TotalPayment, payment.TotalPayment))
						needsUpdate = true
					}
					if existing.Amount != payment.Amount {
						changes = append(changes, fmt.Sprintf("Amount: %.2f -> %.2f", existing.Amount, payment.Amount))
						needsUpdate = true
					}
					if existing.State != payment.State.Name {
						changes = append(changes, fmt.Sprintf("State: %s -> %s", existing.State, payment.State.Name))
						needsUpdate = true
					}

					if needsUpdate {
						log.Printf("  Tespit edilen değişiklikler:")
						for _, change := range changes {
							log.Printf("    - %s", change)
						}

						updateMutation := `
						mutation UpdateJobPayments($id: Int!, $input: JobPaymentsInput!) {
							updateJobPayments(id: $id, input: $input) {
								id
								PaymentNo
								PaymentType
								PaymentDate
								TotalPayment
								Amount
								State
							}
						}
						`

						var updateResult struct {
							UpdateJobPayments struct {
								ID           int     `json:"id"`
								PaymentNo    int     `json:"PaymentNo"`
								PaymentType  string  `json:"PaymentType"`
								PaymentDate  string  `json:"PaymentDate"`
								TotalPayment float64 `json:"TotalPayment"`
								Amount       float64 `json:"Amount"`
								State        string  `json:"State"`
							} `json:"updateJobPayments"`
						}

						updateVariables := map[string]interface{}{
							"id":    existing.ID,
							"input": paymentData,
						}

						if err := graphqlClient.Execute(updateMutation, updateVariables, jwtToken, &updateResult); err != nil {
							errMsg := fmt.Sprintf("Hakediş güncellenirken hata oluştu: %v", err)
							log.Printf("Hata: %s - Hakediş #%d\n", errMsg, payment.ProgressPaymentNumber)
							processLog["status"] = "error"
							processLog["error"] = errMsg
							processLogs = append(processLogs, processLog)
							continue
						}

						log.Printf("Hakediş başarıyla güncellendi - ID: %d, PaymentNo: %d", updateResult.UpdateJobPayments.ID, updateResult.UpdateJobPayments.PaymentNo)
						successCount++
						processLog["status"] = "success"
						processLog["message"] = "Hakediş başarıyla güncellendi"
						processLog["data"] = updateResult.UpdateJobPayments
					} else {
						skippedCount++
						processLog["status"] = "skipped"
						processLog["message"] = "Hakediş verileri güncel"
					}
					break
				}
			}
		}

		// Kayıt yoksa yeni ekle
		if !existingPayment {
			createMutation := `
			mutation CreateJobPayments($input: JobPaymentsInput!) {
				createJobPayments(input: $input) {
					id
					PaymentNo
					PaymentType
					PaymentDate
					TotalPayment
					Amount
					State
				}
			}
			`

			var createResult struct {
				CreateJobPayments struct {
					ID           int     `json:"id"`
					PaymentNo    int     `json:"PaymentNo"`
					PaymentType  string  `json:"PaymentType"`
					PaymentDate  string  `json:"PaymentDate"`
					TotalPayment float64 `json:"TotalPayment"`
					Amount       float64 `json:"Amount"`
					State        string  `json:"State"`
				} `json:"createJobPayments"`
			}

			if err := graphqlClient.Execute(createMutation, map[string]interface{}{"input": paymentData}, jwtToken, &createResult); err != nil {
				errMsg := fmt.Sprintf("Hakediş eklenirken hata oluştu: %v", err)
				log.Printf("Hata: %s - Hakediş #%d\n", errMsg, payment.ProgressPaymentNumber)
				processLog["status"] = "error"
				processLog["error"] = errMsg
				processLogs = append(processLogs, processLog)
				continue
			}

			log.Printf("Hakediş başarıyla eklendi - ID: %d, PaymentNo: %d", createResult.CreateJobPayments.ID, createResult.CreateJobPayments.PaymentNo)
			successCount++
			processLog["status"] = "success"
			processLog["message"] = "Hakediş başarıyla eklendi"
			processLog["data"] = createResult.CreateJobPayments
		}

		processLogs = append(processLogs, processLog)
	}

	result := fiber.Map{
		"yibfNo":      requestParams.YibfNo,
		"companyCode": companyCode,
		"message": fmt.Sprintf("%d adet hakediş kaydından %d tanesi başarıyla eklendi/güncellendi, %d tanesi zaten güncel",
			len(progressPaymentResponse.Items),
			successCount,
			skippedCount),
		"totalCount":   len(progressPaymentResponse.Items),
		"successCount": successCount,
		"skippedCount": skippedCount,
		"logs":         processLogs,
	}

	return c.JSON(result)
}
