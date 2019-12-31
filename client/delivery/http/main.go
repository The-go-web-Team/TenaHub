package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/TenaHub/client/delivery/http/handler"
	"html/template"
)
var Templ *template.Template

func init()  {
	Templ = template.Must(template.ParseGlob("client/ui/templates/*"))
}
func main()  {

	AdminHomeHandler := clientHandler.NewHomeHandler(Templ)
	router := mux.NewRouter()
	router.PathPrefix("/assets/").
		Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("client/ui/assets"))))
	router.HandleFunc("/", AdminHomeHandler.HomePage).Methods("GET")
	http.ListenAndServe(":8282", router)
}
