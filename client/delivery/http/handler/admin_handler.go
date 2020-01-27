package handler

import (
	"net/http"
	"html/template"
	"github.com/TenaHub/client/service"
	"fmt"
	"encoding/json"
	"bytes"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
	"github.com/TenaHub/client/entity"
	"net/url"
	"github.com/TenaHub/client/rtoken"
	"github.com/TenaHub/client/form"
)


type AdminHandler struct {
	temp *template.Template
	userHandl *UserHandler
	CsrfSignKey  []byte
}
func NewAdminHandler(T *template.Template, uh *UserHandler, csk []byte) *AdminHandler {
	return &AdminHandler{temp: T, userHandl:uh, CsrfSignKey:csk}
}


func (adh *AdminHandler) AllAgents(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	users, err := service.FetchAgent(6)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		adh.temp.ExecuteTemplate(w, "admin_home.layout", nil)
	}
	adh.temp.ExecuteTemplate(w, "check.html", users)
}
type data struct {
	Admin *entity.Admin
	Agent []entity.User
	HealthCenter []entity.HealthCenter
	User []entity.User
	Form form.Input
}
func (adh *AdminHandler) AdminPage(w http.ResponseWriter, r *http.Request) {
	//c, err := r.Cookie("admin")

	//if err != nil {
	//	//adh.temp.ExecuteTemplate(w, "admin.login.layout",nil)
	//	http.Redirect(w, r, "http://localhost:8282/admin/login", http.StatusSeeOther)
	//	return
	//} else {
	//	fmt.Println(c.Value)
	//	fmt.Println(c.MaxAge)
	//}
	//	id, _ := strconv.Atoi(c.Value)
	//	usr := adh.userHandl.LoggedInUser
	//	fmt.Println(usr)
		token, err := rtoken.CSRFToken(adh.CsrfSignKey)
		agentForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		admin, err := service.FetchAdmin(id)
		fmt.Println(admin)
		agents, err := service.FetchAgents()
		healthCenters, err := service.FetchHealthCenters()
		users, err := service.FetchUsers()

		if err != nil {
			fmt.Println("here")
			fmt.Println(err)
			w.WriteHeader(http.StatusNoContent)
			return
			//http.Redirect(w, r, "http://localhost:8282/admin/login", http.StatusSeeOther)
		}
		adh.temp.ExecuteTemplate(w, "admin_home.layout", data{admin,agents, healthCenters, users, agentForm})
		return
}

func (adh *AdminHandler) EditAdmin(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("admin")
	id, _ := strconv.Atoi(c.Value)

	firstName := r.FormValue("firstname")
	lastName := r.FormValue("lastname")
	username := r.FormValue("username")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	password := r.FormValue("password")
	confirm := r.FormValue("confirm")

	if password != confirm {
		return
	}
	fileName, err := FileUpload(r,"admin_uploads")
	if err != nil{
		fmt.Println(err)
	}

	data := entity.Admin{FirstName:firstName, LastName:lastName, UserName:username, Email:email,PhoneNumber:phone,Password:password,ProfilePic:fileName}
	jsonValue, _ := json.Marshal(data)
	client := &http.Client{}

	URL := fmt.Sprintf("http://localhost:8181/v1/admin/%d", id)

	req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(jsonValue))
	_, err = client.Do(req)
	var status addStatus
	if err != nil {
		status.Success = false
	}else {
		status.Success = true
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}


//func (u *entity.Admin) Prepare() {
//	//u.ID = 0
//	//u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
//	//u.Email = html.EscapeString(strings.TrimSpace(u.Email))
//	//u.CreatedAt = time.Now()
//	//u.UpdatedAt = time.Now()
//}

//func (u *User) Validate(action string) error {
//	switch strings.ToLower(action) {
//	case "update":
//		if u.Nickname == "" {
//			return errors.New("Required Nickname")
//		}
//		if u.Password == "" {
//			return errors.New("Required Password")
//		}
//		if u.Email == "" {
//			return errors.New("Required Email")
//		}
//		if err := checkmail.ValidateFormat(u.Email); err != nil {
//			return errors.New("Invalid Email")
//		}
//
//		return nil
//	case "login":
//		if u.Password == "" {
//			return errors.New("Required Password")
//		}
//		if u.Email == "" {
//			return errors.New("Required Email")
//		}
//		if err := checkmail.ValidateFormat(u.Email); err != nil {
//			return errors.New("Invalid Email")
//		}
//		return nil
//
//	default:
//		if u.Nickname == "" {
//			return errors.New("Required Nickname")
//		}
//		if u.Password == "" {
//			return errors.New("Required Password")
//		}
//		if u.Email == "" {
//			return errors.New("Required Email")
//		}
//		if err := checkmail.ValidateFormat(u.Email); err != nil {
//			return errors.New("Invalid Email")
//		}
//		return nil
//	}
//}


// Login handles Get /login and POST /login
func (ah *AdminHandler) AdminLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Referer())
	if r.Method == http.MethodGet {
		ah.temp.ExecuteTemplate(w, "admin.login.layout", nil)

	} else if r.Method == http.MethodPost {
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		admin := entity.Admin{Email: email, Password: password}
		fmt.Println(admin)

		resp, err := service.AdminAuthenticate(&admin)
		//
		fmt.Println(resp, " error ", err)

		if err != nil {
			if err.Error() == "error" {
				//http.Redirect(w,r,"/admin",http.StatusSeeOther)
				fmt.Println("before executing")
				ah.temp.ExecuteTemplate(w, "admin.login.layout", "incorrect credentials")
				return
			}
		} else {
			fmt.Println(resp ," is the correct one")

			cookie := http.Cookie{
				Name:     "admin",
				Value:    strconv.Itoa(int(resp.ID)),
				MaxAge:   60 * 30,
				Path:     "/",
				HttpOnly: true,
			}

			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "http://localhost:8282/admin", http.StatusSeeOther)
		}
	}
}

// Logout handles GET /logout
func (uh *AdminHandler) AdminLogout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("admin")

	if err != nil {
		http.Redirect(w, r, "http://localhost:8282/admin/login", http.StatusSeeOther)
		return
	}
	if c != nil {
		c = &http.Cookie{
			Name:     "admin",
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


func HashPassword(password []byte)(string, error){
	hashedPassword,err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return string(hashedPassword), err
}
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
