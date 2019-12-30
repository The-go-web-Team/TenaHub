package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/TenaHub/api/delivery/http/handler"
	"html/template"
)
var Templ *template.Template

func init()  {
	Templ = template.Must(template.ParseGlob("ui/templates/*"))
}
func main()  {

	AdminHomeHandler := handler.NewHomeHandler(Templ)
	router := mux.NewRouter()
	router.PathPrefix("/assets/").
		Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("ui/assets"))))
	router.HandleFunc("/", AdminHomeHandler.HomePage).Methods("GET")
	http.ListenAndServe(":8181", router)
}
