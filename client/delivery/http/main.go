package main

import (
	"html/template"
	"net/http"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/client/delivery/http/handler"
	"github.com/gorilla/mux"
)

var templ = template.Must(template.ParseGlob("../../ui/templates/*"))

func main() {

	userHandler := handler.NewUserHandler(templ)

	router := mux.NewRouter()
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("../../ui/assets"))))
	router.HandleFunc("/", userHandler.Index)
	router.HandleFunc("/login", userHandler.Login)
	router.HandleFunc("/signup", userHandler.SignUp)
	router.HandleFunc("/search", userHandler.Search)
	router.HandleFunc("/home", userHandler.Home)
	router.HandleFunc("/healthcenters", userHandler.Healthcenters)
	router.HandleFunc("/logout", userHandler.Logout)

	http.ListenAndServe(":8282", router)
}
