package service

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/TenaHub/client/entity"
	"fmt"
)
var baseURL = "http://localhost:8181/v1"

func FetchHealthCenter(id int) (*clientEntity.HealthCenter, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s/healthcenter/%d", baseURL, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	healthcenter := &clientEntity.HealthCenter{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, healthcenter)
	if err != nil {
		return nil, err
	}
	return healthcenter, nil
}

func FetchUser(id int) (*clientEntity.User, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s/user/%d", baseURL, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	userdata := &clientEntity.User{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, userdata)
	if err != nil {
		return nil, err
	}
	return userdata, nil
}

