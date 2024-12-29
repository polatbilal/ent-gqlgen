package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gqlgen-ent/handlers"

	"github.com/gin-gonic/gin"
)

func TestFindAllAndFindById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/sync/findAll", handlers.FindAll)
	r.GET("/sync/findById/:id", handlers.FindById)

	// Önce FindAll'ı çağır
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync/findAll", nil)
	req.Header.Set("Authorization", "Bearer 0fP4RDeflcbvsDEJgcDqtnJz2ZE")
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	// FindAll yanıtını parse et
	var findAllResponse struct {
		Success bool  `json:"success"`
		Ids     []int `json:"ids"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &findAllResponse)
	if err != nil {
		t.Fatalf("FindAll yanıtı parse edilemedi: %v", err)
	}

	t.Logf("FindAll'dan gelen ID'ler: %v", findAllResponse.Ids)

	// Her ID için FindById'yi çağır
	for _, id := range findAllResponse.Ids {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/sync/findById/%d", id), nil)
		req.Header.Set("Authorization", "Bearer 0fP4RDeflcbvsDEJgcDqtnJz2ZE")

		r.ServeHTTP(w, req)

		// Önce ham yanıtı logla
		t.Logf("Ham FindById yanıtı (ID: %d): %s", id, w.Body.String())

		var findByIdResponse struct {
			Success bool `json:"success"`
			ID      int  `json:"id"`
			Village struct {
				Parent struct {
					Name   string `json:"name"`
					Parent struct {
						Name string `json:"name"`
					} `json:"parent"`
				} `json:"parent"`
			} `json:"village"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &findByIdResponse)
		if err != nil {
			t.Errorf("ID %d için FindById yanıtı parse edilemedi: %v", id, err)
			continue
		}

		t.Logf("ID: %d, İl: %s, İlçe: %s",
			findByIdResponse.ID,
			findByIdResponse.Village.Parent.Parent.Name,
			findByIdResponse.Village.Parent.Name,
		)
	}
}
