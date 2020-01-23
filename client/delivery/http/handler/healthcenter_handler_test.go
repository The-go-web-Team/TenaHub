package handler

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"html/template"
)

func TestHealthCenter(t *testing.T) {
	var templ = template.Must(template.ParseGlob("../../../ui/templates/*.html"))
	hcHandler := NewHealthCenterHandler(templ)
	httprr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/healthcenter", nil)
	if err != nil {
		t.Fatal(err)
	}
	hcHandler.HealthCenterPage(httprr, req)
	resp := httprr.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, resp.StatusCode)
	}
}
