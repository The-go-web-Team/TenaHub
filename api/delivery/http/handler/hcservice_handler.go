package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"
	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/hcservice"
	"github.com/julienschmidt/httprouter"
)

// HcserviceHandler handles hcservice related http requests
type HcserviceHandler struct {
	servService hcservice.ServicesService
}

// NewServiceHandler creates HcserviceHandler object
func NewServiceHandler(serv hcservice.ServicesService) *HcserviceHandler {
	return &HcserviceHandler{servService: serv}
}

// GetServices handles GET /v1/services/:id
func (sh *HcserviceHandler) GetServices(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	id, err := strconv.Atoi(ps.ByName("id"))

	services, errs := sh.servService.Services(uint(id))

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(&services, "", "\n")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return
}

// GetSingleService hanldes GET /v1/service/:id
func (sh *HcserviceHandler) GetSingleService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	service, errs := sh.servService.Service(uint(id))

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(service, "", "\n")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return
}

// PutService handler PUT /v1/services/:id
func (sh *HcserviceHandler) PutService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	service, errs := sh.servService.Service(uint(id))

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	err = json.Unmarshal(body, service)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	service, errs = sh.servService.UpdateService(service)

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(service, "", "\n")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return
}

// DeleteService handles DELETE /v1/services/:id
func (sh *HcserviceHandler) DeleteService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	service, errs := sh.servService.DeleteService(uint(id))

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = json.MarshalIndent(service, "", "\n")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}

// PostService handles POST /v1/services
func (sh *HcserviceHandler) PostService(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "application/json")

	service := &entity.Service{}

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)

	err := json.Unmarshal(body, service)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		return
	}

	service, errs := sh.servService.StoreService(service)

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/services/%d", service.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}
