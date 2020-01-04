package admin

import (
	"net/http"
	"html/template"
	"fmt"
	"strconv"
	"github.com/TenaHub/api/entity"
	"encoding/json"
	"bytes"
	"github.com/TenaHub/client/service"
	"github.com/TenaHub/client/entity"
)


type HealthCenterHandler struct {
	temp *template.Template
}
func NewHealthCenterHandler(T *template.Template) *HealthCenterHandler {
	return &HealthCenterHandler{temp: T}
}
type healthcenterData struct {
	FeedBack []clientEntity.Comment
	Service []clientEntity.Service

}

func (adh *HealthCenterHandler) HealthCenterPage(writer http.ResponseWriter, request *http.Request){
	// cross site scripting is used to secure the endpoint from another server
	//writer.Header().Set("Access-Control-Allow-Origin", request.Header.Get("Origin"))

	feedbacks, err := service.FetchFeedbacks()
	services, err := service.FetchServices()

	fmt.Println(err, " is error")

	var data = healthcenterData{FeedBack:feedbacks, Service:services}

	fmt.Println(data)
	if err != nil {
		writer.WriteHeader(http.StatusNoContent)
		adh.temp.ExecuteTemplate(writer, "healthcenter_home.layout", nil)
	}
	adh.temp.ExecuteTemplate(writer, "healthcenter_home.layout", data)
}

func (adh *HealthCenterHandler) EditHealthCenter(w http.ResponseWriter, r *http.Request) {
	Name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	city := r.FormValue("address")
	password := r.FormValue("password")
	confirm := r.FormValue("confirm")
	//id,_ := strconv.Atoi(r.FormValue("hidden_id"))
	id := 3
	if password != confirm {
		return
	}

	data := entity.HealthCenter{ID:uint(id),Name:Name, Email:email,PhoneNumber:phone,City:city,Password:password}
	jsonValue, _ := json.Marshal(data)
	URL := fmt.Sprintf("http://localhost:8181/v1/healthcenter/%d", id)
	fmt.Println(URL)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(jsonValue))
	_, err = client.Do(req)
	var status addStatus
	if err != nil {
		status.Success = false
		fmt.Println(err)
	}else {
		status.Success = true
	}
	fmt.Println(err)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
	//adh.temp.ExecuteTemplate(w, "admin_home.layout", status)
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

