package main

import (
	"net/http"
	"html/template"
	"github.com/TenaHub/client/delivery/http/handler"
)

func main()  {
	AdminHandler := handler.NewAdminHandler(Templ)
	AgentHandler := handler.NewAgentHandler(Templ)
	HealthCenterHandler := handler.NewHealthCenterHandler(Templ)
	UserHandler := handler.NewUserHandler(Templ)
	ServiceHandler := handler.NewServiceHandler(Templ)
	//FeedbackHandler := handler.NewFeedBackHandlerHandler(Templ)

	fs := http.FileServer(http.Dir("client/ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/admin/login", AdminHandler.AdminLogin)
	http.HandleFunc("/admin/logout", AdminHandler.AdminLogout)
	http.HandleFunc("/admin", AdminHandler.AdminPage)
	http.HandleFunc("/admin/updateprofile", AdminHandler.EditAdmin)

	http.HandleFunc("/agent/addagent", AgentHandler.AddAgent)
	http.HandleFunc("/agent/updateprofile", AgentHandler.UpdateAgent)
	http.HandleFunc("/admin/updateagent", AgentHandler.EditAgent)
	http.HandleFunc("/agent/deleteagent",AgentHandler.DeleteAgent )

	http.HandleFunc("/agent", AgentHandler.AgentPage)
	http.HandleFunc("/agent/login", AgentHandler.AgentLogin)
	http.HandleFunc("/agent/logout", AgentHandler.AgentLogout)

	http.HandleFunc("/user/delete",UserHandler.DeleteUser )

	http.HandleFunc("/healthcenter/login", HealthCenterHandler.HealthCenterLogin)
	http.HandleFunc("/healthcenter/logout", HealthCenterHandler.HealthCenterLogout)
	http.HandleFunc("/healthcenter", HealthCenterHandler.HealthCenterPage)
	http.HandleFunc("/healthcenter/add", HealthCenterHandler.AddHealthCenter)
	http.HandleFunc("/healthcenter/updateprofile", HealthCenterHandler.EditHealthCenter)
	http.HandleFunc("/healthcenter/delete",HealthCenterHandler.DeleteHealthCenter )

	http.HandleFunc("/service/addservice",ServiceHandler.AddService )
	http.HandleFunc("/service/editservice",ServiceHandler.EditService )
	http.HandleFunc("/service/deleteservice",ServiceHandler.DeleteService )

	http.ListenAndServe(":8282", nil)

}
