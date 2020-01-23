package handler

import (
	"net/http/httptest"
	"net/http"
	"testing"
	"html/template"
)

func TestService(t *testing.T) {
	var templ = template.Must(template.ParseGlob("../../../ui/templates/*.html"))
	serviceHandler := NewServiceHandler(templ)
	httprr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/service/editservice", nil)
	if err != nil {
		t.Fatal(err)
	}
	serviceHandler.EditService(httprr, req)
	resp := httprr.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, resp.StatusCode)
	}
}