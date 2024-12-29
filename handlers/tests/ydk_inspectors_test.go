package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gqlgen-ent/handlers"

	"github.com/gin-gonic/gin"
)

func TestYDKInspectors(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/ydk/inspectors", handlers.YDKInspectors)

	// Test senaryoları
	tests := []struct {
		name           string
		token          string
		expectedStatus int
	}{
		{
			name:           "Token olmadan istek yapılması",
			token:          "",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Başarılı istek",
			token:          "Bearer 0fP4RDeflcbvsDEJgcDqtnJz2ZE",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/ydk/inspectors", nil)

			if tt.token != "" {
				req.Header.Set("Authorization", tt.token)
			}
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)
			t.Logf("Test yanıtı: %s", w.Body.String())
		})
	}
}
