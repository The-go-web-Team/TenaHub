package main

import (
	"net/http"
	"github.com/TenaHub/client/delivery/http/handler/admin"
	"html/template"
	"github.com/gorilla/mux"
)
var Templ *template.Template

func init()  {
	Templ = template.Must(template.ParseGlob("client/ui/templates/*"))
}
func main()  {


	AdminHandler := admin.NewAdminHandler(Templ)
	AgentHandler := admin.NewAgentHandler(Templ)
	HealthCenterHandler := admin.NewHealthCenterHandler(Templ)
	UserHandler := admin.NewUserHandler(Templ)
	ServiceHandler := admin.NewServiceHandler(Templ)
	//FeedbackHandler := admin.NewFeedBackHandlerHandler(Templ)


	router := mux.NewRouter()
	router.PathPrefix("/assets/").
		Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("client/ui/assets"))))
	//router.HandleFunc("/", AdminHandler.AllAgents).Methods("GET")
	router.HandleFunc("/admin", AdminHandler.AdminPage).Methods("GET")
	router.HandleFunc("/admin/updateprofile", AdminHandler.EditAdmin).Methods("POST")
	router.HandleFunc("/agent/addagent", AgentHandler.AddAgent).Methods("POST")
	router.HandleFunc("/agent/editagent", AgentHandler.EditAgent).Methods("POST")
	router.HandleFunc("/agent/deleteagent",AgentHandler.DeleteAgent ).Methods("POST")
	router.HandleFunc("/healthcenter/delete",HealthCenterHandler.DeleteHealthCenter ).Methods("POST")
	router.HandleFunc("/user/delete",UserHandler.DeleteUser ).Methods("POST")


	router.HandleFunc("/healthcenter", HealthCenterHandler.HealthCenterPage).Methods("GET")
	router.HandleFunc("/healthcenter/updateprofile", HealthCenterHandler.EditHealthCenter).Methods("POST")


	router.HandleFunc("/service/addservice",ServiceHandler.AddService ).Methods("POST")
	router.HandleFunc("/service/editservice",ServiceHandler.EditService ).Methods("POST")
	router.HandleFunc("/service/deleteservice",ServiceHandler.DeleteService ).Methods("POST")


	http.ListenAndServe(":8282", router)

}
