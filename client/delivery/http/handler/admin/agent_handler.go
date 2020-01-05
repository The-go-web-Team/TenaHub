package admin

import (
	"net/http"
	"html/template"
	"fmt"
	"encoding/json"
	"github.com/TenaHub/api/entity"
	"bytes"
	"strconv"
	"github.com/TenaHub/client/entity"
	"github.com/TenaHub/client/service"
	"time"
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

func (adh *AgentHandler) AgentPage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("agent")

	if err != nil {
		//adh.temp.ExecuteTemplate(w, "admin.login.layout",nil)
		http.Redirect(w, r, "http://localhost:8282/agent/login", http.StatusSeeOther)
		return
	} else {
		fmt.Println(c.Value)
		fmt.Println(c.MaxAge)
	}
	id, _ := strconv.Atoi(c.Value)
	admin, err := service.FetchAdmin(id)
	admin.Password = ""
	agents, err := service.FetchAgents()
	healthCenters, err := service.FetchHealthCenters()
	users, err := service.FetchUsers()

	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		//http.Redirect(w, r, "http://localhost:8282/admin/login", http.StatusSeeOther)
	}
	adh.temp.ExecuteTemplate(w, "agent_home.layout", data{admin,agents, healthCenters, users})

}



// Login handles Get /login and POST /login
func (ah *AgentHandler) AgentLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Referer())
	if r.Method == http.MethodGet {
		ah.temp.ExecuteTemplate(w, "agent.login.layout", nil)

	} else if r.Method == http.MethodPost {
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		agent := clientEntity.Agent{Email: email, Password: password}
		fmt.Println(agent)

		resp, err := service.AgentAuthenticate(&agent)
		//
		fmt.Println(resp, " error ", err)

		if err != nil {
			if err.Error() == "error" {
				//http.Redirect(w,r,"/admin",http.StatusSeeOther)
				fmt.Println("before executing")
				ah.temp.ExecuteTemplate(w, "agent.login.layout", "incorrect credentials")
				return
			}
		} else {
			fmt.Println(resp ," is the correct one")

			cookie := http.Cookie{
				Name:     "agent",
				Value:    strconv.Itoa(int(resp.ID)),
				MaxAge:   60 * 3,
				Path:     "/",
				HttpOnly: true,
			}

			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "http://localhost:8282/agent", http.StatusSeeOther)
		}
	}
}

// Logout handles GET /logout
func (uh *AgentHandler) AgentLogout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("admin")

	if err != nil {
		http.Redirect(w, r, "http://localhost:8282/agent/login", http.StatusSeeOther)
		return
	}
	if c != nil {
		c = &http.Cookie{
			Name:     "agent",
			Value:    "",
			Path:     "/",
			Expires:  time.Unix(0, 0),
			MaxAge:   -10,
			HttpOnly: true,
		}

		http.SetCookie(w, c)
	}
	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}


