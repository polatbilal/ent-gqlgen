package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gqlgen-ent/handlers"

	"github.com/gin-gonic/gin"
)

func TestFindById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/sync/findById/:id", handlers.FindById)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sync/findById/2307446", nil)
	req.Header.Set("Authorization", "Bearer FjCLd1ezJzAkH0T2XDFajHgNefU")

	r.ServeHTTP(w, req)
	t.Logf("Test yaniti: %s", w.Body.String())
}
