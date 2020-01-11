package handler

import (
	"net/http"
	"html/template"
	"fmt"
	"strconv"
)


type UserHandler struct {
	temp *template.Template
	//userServe menu.UserService
}
func NewUserHandler(T *template.Template) *UserHandler {
	return &UserHandler{temp: T}
}


func (adh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	client := &http.Client{}
	id,_ := strconv.Atoi(r.FormValue("hidden_id"))
	URL := fmt.Sprintf("http://localhost:8181/v1/user/%d",id)

	req, err := http.NewRequest(http.MethodDelete,URL,nil)

	res, err := client.Do(req)
	var status addStatus
	if err != nil {
		status.Success = false
		fmt.Println(err)
	}else {
		status.Success = true
		fmt.Println(res.StatusCode)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

