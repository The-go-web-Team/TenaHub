package admin

import (
	"net/http"
	"html/template"
	"fmt"
	"encoding/json"
	"github.com/TenaHub/api/entity"
	"bytes"
	"strconv"
)


type AgentHandler struct {
	temp *template.Template
	//userServe menu.UserService
}
func NewAgentHandler(T *template.Template) *AgentHandler {
	return &AgentHandler{temp: T}
}

func (adh *AgentHandler) AddAgent(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("firstname")
	lastName := r.FormValue("lastname")
	username := r.FormValue("username")
	email := r.FormValue("email")
	phone := r.FormValue("phonenum")
	password := r.FormValue("password")

	data := entity.Agent{FirstName:firstName, LastName:lastName, UserName:username, Email:email,PhoneNumber:phone,Password:password}
	jsonValue, _ := json.Marshal(data)
	response, err := http.Post("http://localhost:8181/v1/agent","application/json",bytes.NewBuffer(jsonValue))
	var status addStatus
	if err != nil {
		status.Success = false
		fmt.Println(err)
	}else {
		status.Success = true
		fmt.Println(response.StatusCode)
	}
	http.Redirect(w, r, r.Header.Get("Referer"), 302)

	//adh.temp.ExecuteTemplate(w, "admin_home.layout", status)
}


func (adh *AgentHandler) EditAgent(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("firstname")
	lastName := r.FormValue("lastname")
	username := r.FormValue("username")
	email := r.FormValue("email")
	phone := r.FormValue("phonenumber")
	password := r.FormValue("password")
	id,_ := strconv.Atoi(r.FormValue("hidden_id"))


	data := entity.Agent{ID:uint(id), FirstName:firstName, LastName:lastName, UserName:username, Email:email,PhoneNumber:phone,Password:password}
	jsonValue, _ := json.Marshal(data)
	URL := fmt.Sprintf("http://localhost:8181/v1/agent/%d", id)
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

	//adh.temp.ExecuteTemplate(w, "admin_home.layout", status)
	http.Redirect(w, r, r.Header.Get("Referer"), 302)

	}



func (adh *AgentHandler) DeleteAgent(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	id,_ := strconv.Atoi(r.FormValue("hidden_id"))
	URL := fmt.Sprintf("http://localhost:8181/v1/agent/%d",id)

	req, err := http.NewRequest(http.MethodDelete,URL,nil)
	var status addStatus

	if err != nil {
		status.Success = false
		fmt.Println(err)
	}else {
		status.Success = true
		fmt.Println(req.URL.Path)
	}
	res, err := client.Do(req)

	if err != nil {
		status.Success = false
		fmt.Println(err)
	}else {
		status.Success = true
		fmt.Println(res.StatusCode)
	}
	//adh.temp.ExecuteTemplate(w, "admin_home.layout", status)
	http.Redirect(w, r, r.Header.Get("Referer"), 302)

	}
type addStatus struct {
	Success bool
}
