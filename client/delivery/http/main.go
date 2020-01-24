package main

import (
	"html/template"
	"net/http"
	"github.com/TenaHub/client/delivery/http/handler"
)

var templ = template.Must(template.ParseGlob("client/ui/templates/*.html"))

func main()  {


	AdminHandler := handler.NewAdminHandler(templ)
	AgentHandler := handler.NewAgentHandler(templ)
	HealthCenterHandler := handler.NewHealthCenterHandler(templ)
	UserHandler := handler.NewUserHandler(templ)
	ServiceHandler := handler.NewServiceHandler(templ)
	userHandler := handler.NewUserHandler(templ)
	//FeedbackHandler := admin.NewFeedBackHandlerHandler(templ)


	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("client/ui/assets"))
	router.Handle("/assets/", http.StripPrefix("/assets/", fs))
	//router.HandleFunc("/", AdminHandler.AllAgents).Methods("GET")
	router.HandleFunc("/admin/login", AdminHandler.AdminLogin)
	router.HandleFunc("/admin/logout", AdminHandler.AdminLogout)
	router.HandleFunc("/admin", AdminHandler.AdminPage)
	router.HandleFunc("/admin/updateprofile", AdminHandler.EditAdmin)

	router.HandleFunc("/agent", AgentHandler.AgentPage)
	router.HandleFunc("/agent/login", AgentHandler.AgentLogin)
	router.HandleFunc("/agent/logout", AgentHandler.AgentLogout)
	router.HandleFunc("/agent/addagent", AgentHandler.AddAgent)
	router.HandleFunc("/agent/editagent", AgentHandler.EditAgent)
	router.HandleFunc("/agent/deleteagent",AgentHandler.DeleteAgent )
	router.HandleFunc("/healthcenter/delete",HealthCenterHandler.DeleteHealthCenter )
	router.HandleFunc("/user/delete",UserHandler.DeleteUser )


	router.HandleFunc("/healthcenter/login", HealthCenterHandler.HealthCenterLogin)
	router.HandleFunc("/healthcenter/logout", HealthCenterHandler.HealthCenterLogout)
	router.HandleFunc("/healthcenter", HealthCenterHandler.HealthCenterPage)
	router.HandleFunc("/healthcenter/add", HealthCenterHandler.AddHealthCenter)
	router.HandleFunc("/healthcenter/updateprofile", HealthCenterHandler.EditHealthCenter)


	router.HandleFunc("/service/addservice",ServiceHandler.AddService )
	router.HandleFunc("/service/editservice",ServiceHandler.EditService )
	router.HandleFunc("/service/deleteservice",ServiceHandler.DeleteService )


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
