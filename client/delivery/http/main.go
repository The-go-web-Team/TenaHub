package main

import (
	"html/template"
	"net/http"

	"github.com/TenaHub/client/delivery/http/handler"
	"github.com/TenaHub/client/delivery/http/handler/admin"
	"github.com/gorilla/mux"
)

var templ = template.Must(template.ParseGlob("client/ui/templates/*.html"))

func main()  {


	AdminHandler := admin.NewAdminHandler(templ)
	AgentHandler := admin.NewAgentHandler(templ)
	HealthCenterHandler := admin.NewHealthCenterHandler(templ)
	UserHandler := admin.NewUserHandler(templ)
	ServiceHandler := admin.NewServiceHandler(templ)
	userHandler := handler.NewUserHandler(templ)
	//FeedbackHandler := admin.NewFeedBackHandlerHandler(templ)


	router := mux.NewRouter()
	router.PathPrefix("/assets/").
		Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("client/ui/assets"))))
	//router.HandleFunc("/", AdminHandler.AllAgents).Methods("GET")
	router.HandleFunc("/admin/login", AdminHandler.AdminLogin)
	router.HandleFunc("/admin/logout", AdminHandler.AdminLogout)
	router.HandleFunc("/admin", AdminHandler.AdminPage)
	router.HandleFunc("/admin/updateprofile", AdminHandler.EditAdmin).Methods("POST")

	router.HandleFunc("/agent", AgentHandler.AgentPage)
	router.HandleFunc("/agent/login", AgentHandler.AgentLogin)
	router.HandleFunc("/agent/logout", AgentHandler.AgentLogout)
	router.HandleFunc("/agent/addagent", AgentHandler.AddAgent).Methods("POST")
	router.HandleFunc("/agent/editagent", AgentHandler.EditAgent).Methods("POST")
	router.HandleFunc("/agent/deleteagent",AgentHandler.DeleteAgent ).Methods("POST")
	router.HandleFunc("/healthcenter/delete",HealthCenterHandler.DeleteHealthCenter ).Methods("POST")
	router.HandleFunc("/user/delete",UserHandler.DeleteUser ).Methods("POST")


	router.HandleFunc("/healthcenter/login", HealthCenterHandler.HealthCenterLogin)
	router.HandleFunc("/healthcenter/logout", HealthCenterHandler.HealthCenterLogout)
	router.HandleFunc("/healthcenter", HealthCenterHandler.HealthCenterPage)
	router.HandleFunc("/healthcenter/updateprofile", HealthCenterHandler.EditHealthCenter).Methods("POST")


	router.HandleFunc("/service/addservice",ServiceHandler.AddService ).Methods("POST")
	router.HandleFunc("/service/editservice",ServiceHandler.EditService ).Methods("POST")
	router.HandleFunc("/service/deleteservice",ServiceHandler.DeleteService ).Methods("POST")


	// router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("../../ui/assets"))))
	router.HandleFunc("/", userHandler.Index)
	router.HandleFunc("/login", userHandler.Login)
	router.HandleFunc("/auth", userHandler.Auth)
	router.HandleFunc("/signup", userHandler.SignUp)
	router.HandleFunc("/search", userHandler.Search)
	router.HandleFunc("/home", userHandler.Home)
	router.HandleFunc("/healthcenters", userHandler.Healthcenters)
	router.HandleFunc("/logout", userHandler.Logout)
	router.HandleFunc("/feedback", userHandler.Feedback)


	http.ListenAndServe(":8282", router)

}
