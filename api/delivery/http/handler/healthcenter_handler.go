package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"encoding/json"
	"github.com/TenaHub/api/healthcenter"
	"fmt"
)

type HealthCenterHandler struct {
	healthCenterService healthcenter.HealthCenterService
}
func NewHealthCenterHandler(adm healthcenter.HealthCenterService) *HealthCenterHandler {
	return &HealthCenterHandler{healthCenterService: adm}
}

func (adm *HealthCenterHandler) GetSingleHealthCenter(w http.ResponseWriter,r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	healthcenter, errs := adm.healthCenterService.HealthCenter(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(healthcenter, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
func (adm *HealthCenterHandler) GetHealthCenter(w http.ResponseWriter,r *http.Request, ps httprouter.Params) {

	agents, errs := adm.healthCenterService.HealthCenters()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(agents, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
func (adm *HealthCenterHandler) DeleteHealthCenter(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	_, errs := adm.healthCenterService.DeleteHealthCenter(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
func (adm *HealthCenterHandler) PutHealthCenter(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println(id)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	healthCenterData, errs := adm.healthCenterService.HealthCenter(uint(id))
	//
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)

	fmt.Println(healthCenterData," is admin data")
	json.Unmarshal(body, &healthCenterData)
	healthCenterData.ID = uint(id)
	fmt.Println(healthCenterData," now admin data")
	healthCenterData, errs = adm.healthCenterService.UpdateHealthCenter(healthCenterData)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(healthCenterData, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(output)
	return
}





