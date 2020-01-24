package handler

import (
	"net/http"
	"html/template"
	"fmt"
	"encoding/json"
	"bytes"
	"strconv"
	"github.com/TenaHub/client/service"
	"time"
	"github.com/TenaHub/api/entity"
)


type AgentHandler struct {
	temp *template.Template
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
	password := []byte(r.FormValue("password"))

	hashedPassword,err := HashPassword(password)
	data := entity.Agent{FirstName:firstName, LastName:lastName, UserName:username, Email:email,PhoneNumber:phone,Password:hashedPassword}
	jsonValue, _ := json.Marshal(data)
	_, err = http.Post("http://localhost:8181/v1/agents","application/json",bytes.NewBuffer(jsonValue))
	var status addStatus
	if err != nil {
		status.Success = false
	}else {
		status.Success = true
	}
	http.Redirect(w, r, r.Header.Get("Referer"), 302)
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
	URL := fmt.Sprintf("http://localhost:8181/v1/agents/%d", id)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(jsonValue))
	_, err = client.Do(req)
	var status addStatus
	if err != nil {
		status.Success = false
	}else {
		status.Success = true
	}
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
	}else {
		status.Success = true
	}
	_, err = client.Do(req)

	if err != nil {
		status.Success = false
	}else {
		status.Success = true
	}
	http.Redirect(w, r, r.Header.Get("Referer"), 302)
	}
type addStatus struct {
	Success bool
}
type agentDatas struct {
	Agent entity.Agent
	HealthCenters []entity.HealthCenter
	PendingServices []entity.Service
}
func (adh *AgentHandler) AgentPage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("agent")

	fmt.Println(c.Value, " is value")
	if err != nil {
		http.Redirect(w, r, "http://localhost:8282/agent/login", http.StatusSeeOther)
		return
	}
	id, _ := strconv.Atoi(c.Value)
	agentData, err := service.FetchAgent(id)
	healthcentersByAgent, err := service.FetchHealthCenterByAgentId(uint(id))
	pendingServices, err := service.FetchPendingServices(uint(id))

	fmt.Println("agent are ", agentData)
	fmt.Println("healtcenters are ", healthcentersByAgent)
	fmt.Println("pending services are ", pendingServices)
	data := agentDatas{Agent: *agentData, HealthCenters:healthcentersByAgent, PendingServices:pendingServices}
	fmt.Println("the data is ", data)
	//if err != nil {
	//	w.WriteHeader(http.StatusNoContent)
	//}
	adh.temp.ExecuteTemplate(w, "agent_home.layout",data)
	//adh.temp.ExecuteTemplate(w, "agent_home.layout", data{admin,agents, healthCenters, users})

}



// Login handles Get /login and POST /login
func (ah *AgentHandler) AgentLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ah.temp.ExecuteTemplate(w, "agent.login.layout", nil)

	} else if r.Method == http.MethodPost {
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		agent := entity.Agent{Email: email, Password: password}
		fmt.Println(agent)
		resp, err := service.AgentAuthenticate(&agent)
		if err != nil {
			if err.Error() == "error" {
				fmt.Println("password is not correct")
				ah.temp.ExecuteTemplate(w, "agent.login.layout", "incorrect credentials")
				return
			}
		} else {
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


