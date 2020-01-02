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

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/client/entity"
)

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
func Authenticate(user *entity.User) error {
	requestBody, err := json.MarshalIndent(user, "", "\n")
	URL := fmt.Sprintf("%s%s", baseURL, "user")

	fmt.Println(requestBody, URL)
	if err != nil {
		fmt.Println(err)
		return err
	}

	formval := url.Values{}
	formval.Add("email", user.Email)
	formval.Add("password", user.Password)

	resp, err := http.PostForm(URL, formval)

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

	fmt.Println(string(body))
	return nil
}
