package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gqlgen-ent/handlers"

	"github.com/gin-gonic/gin"
)

func TestFindAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/sync/findAll", handlers.FindAll)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync/findAll", nil)
	req.Header.Set("Authorization", "Bearer FjCLd1ezJzAkH0T2XDFajHgNefU")
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	t.Logf("Test yaniti: %s", w.Body.String())
}
