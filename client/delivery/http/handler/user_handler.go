package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

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

		resp, err := service.Authenticate(&user)

		if err != nil {
			if err.Error() == "error" {
				uh.templ.ExecuteTemplate(w, "user.login.layout", "incorrect credentials")
				return
			}
		} else {
			fmt.Println(resp)

			cookie := http.Cookie{
				Name:     "user",
				Value:    strconv.Itoa(int(resp.ID)),
				MaxAge:   60 * 3,
				Path:     "/",
				HttpOnly: true,
			}

			http.SetCookie(w, &cookie)
			// w.Header().Set("Location:", "https://locahost:8282/home")
			http.Redirect(w, r, "http://localhost:8282/home", http.StatusSeeOther)
			// uh.templ.ExecuteTemplate(w, "user.index.auth.layout", resp)
		}

	}
}

// Home handles GET /home
func (uh *UserHandler) Home(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user")

	if err != nil {
		uh.templ.ExecuteTemplate(w, "user.index.default.layout", "yhe")
		return
	} else {
		fmt.Println(c.Value)
		fmt.Println(c.MaxAge)
	}
	uh.templ.ExecuteTemplate(w, "user.index.auth.layout", "yhe")
}

// SignUp handles GET /signup and POST /signup
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

// Search handles GET /search
func (uh *UserHandler) Search(w http.ResponseWriter, r *http.Request) {
	searchkey := r.URL.Query().Get("search-key")

	fmt.Println(searchkey)

	uh.templ.ExecuteTemplate(w, "user.result.auth.layout", "data")
}

// Healthcenters handles GET /healthcenters
func (uh *UserHandler) Healthcenters(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println(id)

	c, err := r.Cookie("user")

	if err != nil {
		uh.templ.ExecuteTemplate(w, "user.hc.default.layout", "yhe")
		return
	} else {
		fmt.Println(c.Value)
		fmt.Println(c.MaxAge)
	}
	uh.templ.ExecuteTemplate(w, "user.hc.auth.layout", "yhe")

}

// Logout handles GET /logout
func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user")

	if err != nil {
		http.Redirect(w, r, "http://localhost:8282/login", http.StatusSeeOther)
		return
	}
	if c != nil {
		c = &http.Cookie{
			Name:     "user",
			Value:    "",
			Path:     "/",
			Expires:  time.Unix(0, 0),
			MaxAge:   -10,
			HttpOnly: true,
		}

		http.SetCookie(w, c)
	}
	http.Redirect(w, r, "http://localhost:8282/login", http.StatusSeeOther)
}
