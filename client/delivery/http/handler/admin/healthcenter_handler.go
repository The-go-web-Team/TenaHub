package admin

import (
	"net/http"
	"html/template"
	"fmt"
	"strconv"
)


type HealthCenterHandler struct {
	temp *template.Template
}
func NewHealthCenterHandler(T *template.Template) *HealthCenterHandler {
	return &HealthCenterHandler{temp: T}
}

func (adh *HealthCenterHandler) DeleteHealthCenter(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	id,_ := strconv.Atoi(r.FormValue("hidden_id"))
	URL := fmt.Sprintf("http://localhost:8181/v1/healthcenter/%d",id)

	req, err := http.NewRequest(http.MethodDelete,URL,nil)

	res, err := client.Do(req)
	var status addStatus
	if err != nil {
		status.Success = false
		fmt.Println(err)
	}else {
		status.Success = true
		fmt.Println(res.StatusCode)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
	}

