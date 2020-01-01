package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"encoding/json"
	"github.com/TenaHub/api/admin"
	"fmt"
)

type AdminHandler struct {
	adminService admin.AdminService
}
func NewAdminHandler(adm admin.AdminService) *AdminHandler {
	return &AdminHandler{adminService: adm}
}

func (adm *AdminHandler) GetAdmin(w http.ResponseWriter,r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println("Before Recieving")
	admin, errs := adm.adminService.Admin(uint(id))

	fmt.Println("After Recieving", admin)
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(admin, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
