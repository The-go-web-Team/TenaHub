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

	AdminHandler := handler.NewAdminHandler(templ)
	AgentHandler := handler.NewAgentHandler(templ)
	HealthCenterHandler := handler.NewHealthCenterHandler(templ)
	UserHandler := handler.NewUserHandler(templ, sess, csrfSignKey)
	ServiceHandler := handler.NewServiceHandler(templ)
	userHandler := handler.NewUserHandler(templ, sess, csrfSignKey)
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
	router.Handle("/login", userHandler.Authorized(http.HandlerFunc(userHandler.Login)))
	router.Handle("/auth", userHandler.Authorized(http.HandlerFunc(userHandler.Auth)))
	router.Handle("/signup", userHandler.Authorized(http.HandlerFunc(userHandler.SignUp)))
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