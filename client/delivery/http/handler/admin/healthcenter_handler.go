package admin

import (
	"net/http"
	"html/template"
	"fmt"
	"strconv"
	// "github.com/TenaHub/api/entity"
	//"encoding/json"
	//"bytes"
	"github.com/TenaHub/client/service"
	"github.com/TenaHub/client/entity"
	"time"
	"os"
	"io"
	"path/filepath"
	"encoding/json"
	"bytes"
)


type HealthCenterHandler struct {
	temp *template.Template
}
func NewHealthCenterHandler(T *template.Template) *HealthCenterHandler {
	return &HealthCenterHandler{temp: T}
}
type healthcenterData struct {
	HealthCenter *entity.HealthCenter
	FeedBack []entity.Comment
	Service []entity.Service

}

func (adh *HealthCenterHandler) EditHealthCenter(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("healthcenter")
	id, _ := strconv.Atoi(c.Value)

	Name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	city := r.FormValue("address")
	password := r.FormValue("password")
	confirm := r.FormValue("confirm")

	if password != confirm {
		return
	}
	fileName, err := FileUpload(r,"healthcenter_uploads")
	if err != nil{
		fmt.Println(err)
	}
	data := entity.HealthCenter{ID:uint(id),Name:Name, Email:email,PhoneNumber:phone,City:city,Password:password,ProfilePic:fileName}
	jsonValue, _ := json.Marshal(data)
	URL := fmt.Sprintf("http://localhost:8181/v1/healthcenter/%d", id)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(jsonValue))
	_, err = client.Do(req)
	var status addStatus
	if err != nil {
		status.Success = false
		fmt.Println(err)
	}else {
		status.Success = true
	}
	fmt.Println(err)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
	adh.temp.ExecuteTemplate(w, "admin_home.layout", status)
}


func (adh *HealthCenterHandler) DeleteHealthCenter(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	id,_ := strconv.Atoi(r.FormValue("hidden_id"))
	URL := fmt.Sprintf("http://localhost:8181/v1/healthcenter/%d",id)

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


func (adh *HealthCenterHandler) HealthCenterPage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("healthcenter")

	if err != nil {
		http.Redirect(w, r, "http://localhost:8282/healthcenter/login", http.StatusSeeOther)
		return
	} else {
		fmt.Println(c.Value)
		fmt.Println(c.MaxAge)
	}
	id, _ := strconv.Atoi(c.Value)
	healthcenter, err := service.FetchHealthCenter(uint(id))
	feedbacks, err := service.FetchFeedbacks(uint(id))
	services, err := service.FetchService(uint(id))

	var data = healthcenterData{HealthCenter:healthcenter, FeedBack:feedbacks, Service:services}
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		adh.temp.ExecuteTemplate(w, "healthcenter_home.layout", nil)
	}
	adh.temp.ExecuteTemplate(w, "healthcenter_home.layout", data)
}

// Login handles Get /login and POST /login
func (ah *HealthCenterHandler) HealthCenterLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Referer())
	if r.Method == http.MethodGet {
		ah.temp.ExecuteTemplate(w, "healthcenter.login.layout", nil)

	} else if r.Method == http.MethodPost {
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		healthcenter := entity.HealthCenter{Email: email,Password : password}
		resp, err := service.HealthCenterAuthenticate(&healthcenter)
		if err != nil {
			if err.Error() == "error" {
				ah.temp.ExecuteTemplate(w, "healthcenter.login.layout", "incorrect credentials")
				return
			}
		} else {
			cookie := http.Cookie{
				Name:     "healthcenter",
				Value:    strconv.Itoa(int(resp.ID)),
				MaxAge:   60 * 30,
				Path:     "/",
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "http://localhost:8282/healthcenter", http.StatusSeeOther)
		}
	}
}
// Logout handles GET /logout
func (uh *HealthCenterHandler) HealthCenterLogout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("healthcenter")
	if err != nil {
		http.Redirect(w, r, "http://localhost:8282/healthcenter/login", http.StatusSeeOther)
		return
	}
	if c != nil {
		c = &http.Cookie{
			Name:     "healthcenter",
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

func FileUpload(r *http.Request, folderName string) (string, error) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("upload_image")
	if err != nil {
		return "",err
	}
	defer file.Close()
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "client","ui", "assets", "img", "uploads",folderName, handler.Filename)

	f, err := os.OpenFile(path,os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "",err
	}
	defer f.Close()
	io.Copy(f, file)
	return handler.Filename, nil
}