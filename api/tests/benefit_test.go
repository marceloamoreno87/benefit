package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marceloamoreno87/api/routes"
	"github.com/stretchr/testify/assert"
)

func TestGetBenefitWithoutDoc(t *testing.T) {
	r := routes.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/benefit", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestGetBenefitWithWrongDoc(t *testing.T) {
	r := routes.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/benefit?doc=12345678978", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestGetBenefitWithRightDoc(t *testing.T) {
	r := routes.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/benefit?doc=927.100.938-04", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetBenefitWithRightDocWithoutCaracters(t *testing.T) {
	r := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/benefit?doc=92710093804", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
