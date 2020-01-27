package main

import (
	"html/template"
	"net/http"
	"github.com/TenaHub/client/delivery/http/handler"
	"fmt"
	"time"
	"github.com/TenaHub/client/entity"
	"github.com/TenaHub/client/rtoken"
)

var templ = template.Must(template.ParseGlob("client/ui/templates/*.html"))

func main()  {
	csrfSignKey := []byte(rtoken.GenerateRandomID(32))
	sess := configSess()
	fmt.Println(sess)

	UserHandler := handler.NewUserHandler(templ, sess, csrfSignKey)
	AdminHandler := handler.NewAdminHandler(templ, UserHandler, csrfSignKey)
	AgentHandler := handler.NewAgentHandler(templ, csrfSignKey)
	HealthCenterHandler := handler.NewHealthCenterHandler(templ, sess, csrfSignKey )
	ServiceHandler := handler.NewServiceHandler(templ)
	userHandler := handler.NewUserHandler(templ, sess, csrfSignKey)
	//FeedbackHandler := admin.NewFeedBackHandlerHandler(templ)


	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("client/ui/assets"))
	router.Handle("/assets/", http.StripPrefix("/assets/", fs))
	//router.HandleFunc("/", AdminHandler.AllAgents).Methods("GET")
	//router.HandleFunc("/admin/login", AdminHandler.AdminLogin)
	//router.HandleFunc("/admin/logout", AdminHandler.AdminLogout)
	//router.HandleFunc("/admin", AdminHandler.AdminPage)
	//router.HandleFunc("/admin/updateprofile", AdminHandler.EditAdmin)

	//router.HandleFunc("/admin/login", AdminHandler.AdminLogin)
	//router.Handle("/admin/logout", userHandler.Authenticated(http.HandlerFunc(AdminHandler.AdminLogout)) )
	router.Handle("/admin", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(AdminHandler.AdminPage))) )
	router.Handle("/admin/updateprofile", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(AdminHandler.EditAdmin))))
	router.Handle("/admin/agent/addagent",  userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(AgentHandler.AddAgent))))
	router.Handle("/admin/agent/deleteagent",userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(AgentHandler.DeleteAgent))) )
	router.Handle("/admin/user/delete",userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(UserHandler.DeleteUser))))

	//router.HandleFunc("/agent", AgentHandler.AgentPage)
	//router.HandleFunc("/agent/login", AgentHandler.AgentLogin)
	//router.HandleFunc("/agent/logout", AgentHandler.AgentLogout)
	//router.HandleFunc("/agent/editagent", AgentHandler.EditAgent)
	//router.HandleFunc("/healthcenter/delete",HealthCenterHandler.DeleteHealthCenter )
	//
	//router.HandleFunc("/healthcenter/login", HealthCenterHandler.HealthCenterLogin)
	//router.HandleFunc("/healthcenter/logout", HealthCenterHandler.HealthCenterLogout)
	//router.HandleFunc("/healthcenter", HealthCenterHandler.HealthCenterPage)
	//router.HandleFunc("/healthcenter/add", HealthCenterHandler.AddHealthCenter)
	//router.HandleFunc("/healthcenter/updateprofile", HealthCenterHandler.EditHealthCenter)

	router.Handle("/agent", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(AgentHandler.AgentPage))))
	//router.HandleFunc("/agent/login", AgentHandler.AgentLogin)
	//router.HandleFunc("/agent/logout", AgentHandler.AgentLogout)
	router.Handle("/agent/editagent", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(AgentHandler.EditAgent))))
	router.Handle("/agent/healthcenter/delete", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(HealthCenterHandler.DeleteHealthCenter))))
	router.Handle("/agent/healthcenter/add", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(AgentHandler.AddHealthCenter))))


	//router.HandleFunc("/healthcenter/login", HealthCenterHandler.HealthCenterLogin)
	//router.HandleFunc("/healthcenter/logout", HealthCenterHandler.HealthCenterLogout)
	//router.HandleFunc("/healthcenter", HealthCenterHandler.HealthCenterPage)
	//router.HandleFunc("/healthcenter/updateprofile", HealthCenterHandler.EditHealthCenter)
	//router.HandleFunc("/healthcenter/service/addservice",ServiceHandler.AddService)
	//router.HandleFunc("/healthcenter/service/editservice",ServiceHandler.EditService)
	//router.HandleFunc("/healthcenter/service/deleteservice",ServiceHandler.DeleteService)

	router.HandleFunc("/healthcenter/login", HealthCenterHandler.HealthCenterLogin)
	router.Handle("/healthcenter/logout", HealthCenterHandler.Authenticated(http.HandlerFunc(HealthCenterHandler.HealthCenterLogout)))
	router.Handle("/healthcenter", HealthCenterHandler.Authenticated(HealthCenterHandler.Authorized(http.HandlerFunc(HealthCenterHandler.HealthCenterPage))))
	router.Handle("/healthcenter/updateprofile", HealthCenterHandler.Authenticated(HealthCenterHandler.Authorized(http.HandlerFunc(HealthCenterHandler.EditHealthCenter))))
	router.Handle("/healthcenter/service/addservice", HealthCenterHandler.Authenticated(HealthCenterHandler.Authorized(http.HandlerFunc(ServiceHandler.AddService))))
	router.Handle("/healthcenter/service/editservice", HealthCenterHandler.Authenticated(HealthCenterHandler.Authorized(http.HandlerFunc(ServiceHandler.EditService))))
	router.Handle("/healthcenter/service/deleteservice", HealthCenterHandler.Authenticated(HealthCenterHandler.Authorized(http.HandlerFunc(ServiceHandler.DeleteService))))


	// router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("../../ui/assets"))))
	//router.HandleFunc("/",  userHandler.Index)
	//router.HandleFunc("/login", userHandler.Login)
	//router.HandleFunc("/auth", userHandler.Auth)
	//router.HandleFunc("/signup", userHandler.SignUp)
	//router.HandleFunc("/search", userHandler.Search)
	//router.HandleFunc("/home", userHandler.Home)
	//router.HandleFunc("/healthcenters", userHandler.Healthcenters)
	//router.HandleFunc("/logout", userHandler.Logout)
	//router.HandleFunc("/feedback", userHandler.Feedback)

	router.Handle("/", userHandler.Authorized(http.HandlerFunc(userHandler.Index)))
	router.Handle("/login", http.HandlerFunc(userHandler.Login))
	router.Handle("/signup", http.HandlerFunc(userHandler.SignUp))
	router.Handle("/auth", userHandler.Authorized(http.HandlerFunc(userHandler.Auth)))
	router.Handle("/search", userHandler.Authorized(http.HandlerFunc(userHandler.Search)))
	router.Handle("/home", userHandler.Authorized(http.HandlerFunc(userHandler.Home)))
	router.Handle("/healthcenters", userHandler.Authorized(http.HandlerFunc(userHandler.Healthcenters)))
	router.Handle("/logout", userHandler.Authorized(http.HandlerFunc(userHandler.Logout)))
	router.Handle("/feedback", userHandler.Authorized(http.HandlerFunc(userHandler.Feedback)))
	http.ListenAndServe(":8282", router)
}

func configSess() *entity.Session {
	tokenExpires := time.Now().Add(time.Minute * 30).Unix()
	sessionID := rtoken.GenerateRandomID(32)
	signingString, err := rtoken.GenerateRandomString(32)

	if err != nil {
		panic(err)
	}
	signingKey := []byte(signingString)
	return &entity.Session{
		Expires:    tokenExpires,
		SigningKey: signingKey,
		UUID:       sessionID,
	}
}