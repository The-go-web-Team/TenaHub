package handler

import (
	"net/http"
	"html/template"
)
type AdminHandler struct {
	temp *template.Template
	//userServe menu.UserService
}
func NewHomeHandler(T *template.Template) *AdminHandler {
	return &AdminHandler{temp: T}
}

func (adh *AdminHandler) HomePage(writer http.ResponseWriter, request *http.Request){
	adh.temp.ExecuteTemplate(writer, "admin_home.layout", nil)
}