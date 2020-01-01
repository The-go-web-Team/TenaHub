package clientHandler

import (
	"net/http"
	"html/template"
	"github.com/TenaHub/client/service"
	"fmt"
)


type AdminHandler struct {
	temp *template.Template
	//userServe menu.UserService
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
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	fmt.Println("something")
	users, err := service.FetchAgents()
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		adh.temp.ExecuteTemplate(w, "admin_home.layout", nil)
	}
	adh.temp.ExecuteTemplate(w, "check.html", users)
}