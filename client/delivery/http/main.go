package main

import (
	"html/template"
	"net/http"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/client/delivery/http/handler"
	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/client/delivery/http/handler/admin"
	"github.com/gorilla/mux"
)

var templ = template.Must(template.ParseGlob("../../ui/templates/*"))

func main()  {


	AdminHandler := admin.NewAdminHandler(Templ)
	AgentHandler := admin.NewAgentHandler(Templ)
	HealthCenterHandler := admin.NewHealthCenterHandler(Templ)
	UserHandler := admin.NewUserHandler(Templ)
	ServiceHandler := admin.NewServiceHandler(Templ)
	userHandler := handler.NewUserHandler(templ)
	//FeedbackHandler := admin.NewFeedBackHandlerHandler(Templ)


	router := mux.NewRouter()
	router.PathPrefix("/assets/").
		Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("client/ui/assets"))))
	//router.HandleFunc("/", AdminHandler.AllAgents).Methods("GET")
	router.HandleFunc("/admin/login", AdminHandler.AdminLogin)
	router.HandleFunc("/admin/logout", AdminHandler.AdminLogout)
	router.HandleFunc("/admin", AdminHandler.AdminPage)
	router.HandleFunc("/admin/updateprofile", AdminHandler.EditAdmin).Methods("POST")

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
