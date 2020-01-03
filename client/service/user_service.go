package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/client/entity"
)

type cookie struct {
	Key        string
	Expiration time.Time
}

type response struct {
	Status string
	Content interface{}
}

var loggedIn = make([]cookie, 10)

const baseURL string = "http://localhost:8181/v1/"	

func getResponse(request *http.Request) []byte {
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	if 200 != resp.StatusCode {
		panic(errors.New("status not correct"))
	}

	return body
}

// PostUser posts user to api
func PostUser(user *entity.User) error {
	requestBody, err := json.MarshalIndent(user, "", "\n")
	URL := fmt.Sprintf("%s%s", baseURL, "users")

	if err != nil {
		fmt.Println(err)
		return err
	}

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if strings.Compare(string(body), "Not Found") != 0 {
		return errors.New("duplicate")
	}
	fmt.Println(string(body))
	return nil
}

// Authenticate authenticates user
func Authenticate(user *entity.User) (*entity.User, error) {
	// requestBody, err := json.MarshalIndent(user, "", "\n")
	URL := fmt.Sprintf("%s%s", baseURL, "user")

	// fmt.Println(requestBody, URL)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }

	formval := url.Values{}
	formval.Add("email", user.Email)
	formval.Add("password", user.Password)

	resp, err := http.PostForm(URL, formval)


	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
	respjson := struct {
		Status string
		Content entity.User
	}{}

	err = json.Unmarshal(body, &respjson)

	fmt.Println(respjson)


	if respjson.Status == "error" {
		return nil, errors.New("error")
	}



	return &respjson.Content, nil
}
