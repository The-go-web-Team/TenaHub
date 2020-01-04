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

	healthcenters, err := service.GetHealthcenters(searchkey)

	if err != nil {
		uh.templ.ExecuteTemplate(w, "user.error.layout", nil)
		return
	}

	length := len(healthcenters)

	data := struct {
		Length  int
		Content []entity.HealthCenter
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
	frating, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 4.26234), 64)
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
		Comments     []entity.Comment
		Isvalid      string
	}{
		Rating:       frating,
		Healthcenter: *hc,
		Services:     services,
		Comments: comments,
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
