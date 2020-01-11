package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/TenaHub/api/healthcenter"
)

// HcHanlder handles healthcenter related http requests
type HcHanlder struct {
	hcServ healthcenter.HealthCenterService
}

// NewHcHandler creates object of HcHandler
func NewHcHandler(hcserv healthcenter.HealthCenterService) *HcHanlder {
	return &HcHanlder{hcServ: hcserv}
}

// GetHealthcenters handles GET /healthcenters
func (hh *HcHanlder) GetHealthcenters(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	searchKey := r.URL.Query().Get("search-key")
	column := r.URL.Query().Get("column")
	fmt.Println(searchKey)

	hcs, errs := hh.hcServ.SearchHealthCenters(searchKey, column)

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(&hcs, "", "\n")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return

}

// GetHealthcenter handles GET/healthcenters/:id
func (hh *HcHanlder) GetHealthcenter(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	healthcenter, errs := hh.hcServ.HealthCenterById(uint(id))

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(healthcenter, "", "\n")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return

}

// GetTop handles Get /healthcenters/top/:amount
func (hh *HcHanlder) GetTop(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	amount, err := strconv.Atoi(ps.ByName("amount"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	results, errs := hh.hcServ.Top(uint(amount))

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(results, "", "\n")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return
}