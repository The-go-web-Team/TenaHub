package handler

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
	"github.com/TenaHub/client/service"
	"github.com/TenaHub/client/session"
	"github.com/TenaHub/api/entity"
)

// UserHandler handles user related http requests
type UserHandler struct {
	templ        *template.Template
	userSess     *entity.Session
	loggedInUser *entity.User
	csrfSignKey  []byte
}

// NewUserHandler creates object of UserHandler
func NewUserHandler(tmpl *template.Template) *UserHandler {
	return &UserHandler{templ: tmpl}
}

type contextKey string

var ctxUserSessionKey = contextKey("signed_in_user_session")

// Authenticated checks if a user is authenticated to access a given route
func (uh *UserHandler) Authenticated(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ok := uh.loggedIn(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), ctxUserSessionKey, uh.userSess)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func (uh *UserHandler) loggedIn(r *http.Request) bool {
	if uh.userSess == nil {
		return false
	}
	userSess := uh.userSess
	c, err := r.Cookie(userSess.UUID)

	if err != nil {
		return false
	}

	ok, err := session.Valid(c.Value, userSess.SigningKey)
	if !ok || (err != nil) {
		return false
	}
	return true
}

// Index handles GET /
func (uh *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	hcs, err := service.GetTop(4)

	if err != nil {
		uh.templ.ExecuteTemplate(w, "user.error.layout", nil)
		return
	}

	fmt.Println(hcs)

	uh.templ.ExecuteTemplate(w, "user.index.default.layout", hcs)
}

// Login handles Get /login and POST /login
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Referer())
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

// Auth authenticates user and redirect to referer POST /auth
func (uh *UserHandler) Auth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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
			http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
			// uh.templ.ExecuteTemplate(w, "user.index.auth.layout", resp)
		}

	}
}

// Home handles GET /home
func (uh *UserHandler) Home(w http.ResponseWriter, r *http.Request) {
	hcs, err := service.GetTop(4)

	if err != nil {
		uh.templ.ExecuteTemplate(w, "user.error.layout", nil)
		return
	}

	fmt.Println(hcs)

	c, err := r.Cookie("user")

	if err != nil {
		uh.templ.ExecuteTemplate(w, "user.index.default.layout", hcs)
		return
	} else {
		fmt.Println(c.Value)
		fmt.Println(c.MaxAge)
	}
	uh.templ.ExecuteTemplate(w, "user.index.auth.layout", hcs)
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
	column := r.URL.Query().Get("column")
	if column == "" {
		column = "name"
	}
	fmt.Println(searchkey)

	healthcenters, err := service.GetHealthcenters(searchkey, column)

	if err != nil {
		uh.templ.ExecuteTemplate(w, "user.error.layout", nil)
		return
	}

	length := len(healthcenters)

	data := struct {
		Length  int
		Content []entity.Hcrating
	}{
		Length:  length,
		Content: healthcenters,
	}

	c, err := r.Cookie("user")

	if err != nil {
		uh.templ.ExecuteTemplate(w, "user.result.default.layout", data)
		return
	} else {
		fmt.Println(c.Value)
		fmt.Println(c.MaxAge)
	}

	uh.templ.ExecuteTemplate(w, "user.result.auth.layout", data)
}

// Healthcenters handles GET /healthcenters
func (uh *UserHandler) Healthcenters(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	fmt.Println(id)

	hc, errr := service.GetHealthcenter(uint(id))

	if errr != nil {
		uh.templ.ExecuteTemplate(w, "user.error.layout", nil)
		return
	}

	services, err := service.GetServices(uint(id))

	if err != nil {
		uh.templ.ExecuteTemplate(w, "user.error.layout", nil)
		return
	}

	rating, err := service.GetRating(uint(id))
	frating, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rating), 64)
	fmt.Println("rating: ", rating)

	if err != nil {
		uh.templ.ExecuteTemplate(w, "user.error.layout", nil)
		return
	}

	comments, err := service.GetFeedback(hc.ID)

	if err != nil {
		uh.templ.ExecuteTemplate(w, "user.error.layout", nil)
		return
	}

	fmt.Println(comments)

	data := struct {
		Rating       float64
		Healthcenter entity.HealthCenter
		Services     []entity.Service
		Comments     []entity.UserComment
		Isvalid      string
	}{
		Rating:       frating,
		Healthcenter: *hc,
		Services:     services,
		Comments:     comments,
	}

	c, err := r.Cookie("user")

	if err != nil {
		uh.templ.ExecuteTemplate(w, "user.hc.default.layout", data)
		return
	} else {
		fmt.Println(c.Value)
		fmt.Println(c.MaxAge)
	}
	uid, _ := strconv.Atoi(c.Value)
	validity, err := service.CheckValidity(uint(uid), hc.ID)
	fmt.Println(validity)

	if err != nil {
		uh.templ.ExecuteTemplate(w, "user.hc.default.layout", data)
		return
	}
	data.Isvalid = validity
	uh.templ.ExecuteTemplate(w, "user.hc.auth.layout", data)

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
	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}

// Feedback handles POST /feedback
func (uh *UserHandler) Feedback(w http.ResponseWriter, r *http.Request) {
	rating := r.PostFormValue("userrating")
	comment := r.PostFormValue("usercomment")
	hid := r.PostFormValue("hcid")

	c, err := r.Cookie("user")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	uid := c.Value
	frating, err := strconv.Atoi(rating)

	fuid, err := strconv.Atoi(uid)
	fhid, err := strconv.Atoi(strings.Trim(hid, " "))
	if err != nil {
		fmt.Println(err)
		return
	}
	feedback := entity.Comment{
		Rating:         uint(frating),
		Comment:        comment,
		UserID:         uint(fuid),
		HealthCenterID: uint(fhid),
	}

	fmt.Println(feedback)
	err = service.PostFeedback(&feedback)

	if err != nil {
		w.Write([]byte("error"))
		return
	}

	w.Write([]byte("success"))
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

