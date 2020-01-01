package main

import (
	"net/http"
	"github.com/TenaHub/client/delivery/http/handler"
	"html/template"
	"github.com/gorilla/mux"
)
var Templ *template.Template

func init()  {
	Templ = template.Must(template.ParseGlob("../../ui/templates/*"))
}
func main()  {


	AdminHandler := clientHandler.NewAdminHandler(Templ)
	router := mux.NewRouter()
	//router.PathPrefix("/assets/").
	//	Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("client/ui/assets"))))
	router.HandleFunc("/", AdminHandler.AllAgents).Methods("GET")
	// http.ListenAndServe(":8282", router)

	//router := httprouter.New()
	//router.GET("/", AdminHandler.AllAgents)
	//fmt.Println("something")

	http.ListenAndServe(":8282", router)

}
