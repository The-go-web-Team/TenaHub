package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/client/service"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/client/entity"
)

// UserHandler handles user related http requests
type UserHandler struct {
	templ *template.Template
}

// NewUserHandler creates object of UserHandler
func NewUserHandler(tmpl *template.Template) *UserHandler {
	return &UserHandler{templ: tmpl}
}

// Index handles GET /
func (uh *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	uh.templ.ExecuteTemplate(w, "user.index.default.layout", nil)
}

// Login handles Get /login and POST /login
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		uh.templ.ExecuteTemplate(w, "user.login.layout", nil)
	} else if r.Method == http.MethodPost {
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		user := entity.User{Email: email, Password: password}
		fmt.Println(user)

		err := service.Authenticate(&user)

		fmt.Println(err)
	}
}

// SignUp handles Get /signup and POST /signup
func (uh *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		uh.templ.ExecuteTemplate(w, "user.signup.layout", nil)
	} else if r.Method == http.MethodPost {
		firstname := r.PostFormValue("firstname")
		lastname := r.PostFormValue("lastname")
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		phonenum := r.PostFormValue("phonenum")

		user := entity.User{FirstName: firstname, LastName: lastname, Email: email, Password: password, PhoneNumber: phonenum, Role: "user"}
		fmt.Println(user)

		err := service.PostUser(&user)
		if err != nil {
			if strings.Compare(err.Error(), "duplicate") == 0 {
				fmt.Println("duplicate")
				uh.templ.ExecuteTemplate(w, "user.signup.layout", "email or phone number is already taken")
				return
			}
			w.Write([]byte("failed"))
			return
		} else {
			w.Write([]byte("success"))
			w.Header().Set("Location:", "http://locahost:8282/login")
		}
	}
}
