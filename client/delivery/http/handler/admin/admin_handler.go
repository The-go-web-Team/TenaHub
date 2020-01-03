package admin

import (
	"net/http"
	"html/template"
	"github.com/TenaHub/client/service"
	"fmt"
	"github.com/TenaHub/client/entity"
	"github.com/TenaHub/api/entity"
	"encoding/json"
	"bytes"
)


type AdminHandler struct {
	temp *template.Template
}
func NewAdminHandler(T *template.Template) *AdminHandler {
	return &AdminHandler{temp: T}
}

func (adh *AdminHandler) HomePage(writer http.ResponseWriter, request *http.Request){
	// cross site scripting is used to secure the endpoint from another server
	//writer.Header().Set("Access-Control-Allow-Origin", request.Header.Get("Origin"))
	adh.temp.ExecuteTemplate(writer, "admin_home.layout", nil)
}

func (adh *AdminHandler) AllAgents(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	fmt.Println("something")
	users, err := service.FetchAgent(6)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		adh.temp.ExecuteTemplate(w, "admin_home.layout", nil)
	}
	adh.temp.ExecuteTemplate(w, "check.html", users)
}
type data struct {
	Agent []clientEntity.Agent
	HealthCenter []clientEntity.HealthCenter
	User []clientEntity.User

}
func (adh *AdminHandler) AdminPage(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	fmt.Println("something")
	agents, err := service.FetchAgents()
	healthCenters, err := service.FetchHealthCenters()
	users, err := service.FetchUsers()

	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		adh.temp.ExecuteTemplate(w, "admin_home.layout", nil)
	}
	adh.temp.ExecuteTemplate(w, "admin_home.layout", data{agents, healthCenters, users})
}

func (adh *AdminHandler) EditAdmin(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("firstname")
	lastName := r.FormValue("lastname")
	username := r.FormValue("username")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	password := r.FormValue("password")
	confirm := r.FormValue("confirm")


	if password != confirm {
		return
	}


	data := entity.Admin{FirstName:firstName, LastName:lastName, UserName:username, Email:email,PhoneNumber:phone,Password:password}
	fmt.Println(data," is data")
	jsonValue, _ := json.Marshal(data)
	client := &http.Client{}

	URL := fmt.Sprintf("http://localhost:8181/v1/admin/%d", 1)
	fmt.Println(URL)
	fmt.Println(data)

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

