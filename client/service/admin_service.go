package service

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/TenaHub/client/entity"
	"fmt"
)
var baseURL = "http://localhost:8181/v1"

func FetchAgent(id int) (*clientEntity.Agent, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s/agent/%d", baseURL, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	userdata := clientEntity.Agent{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &userdata)
	fmt.Println("error is ",err)
	if err != nil {
		return nil, err
	}
	return &userdata, nil
}

func FetchAgents() ([]clientEntity.Agent, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s/agent", baseURL)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var agents []clientEntity.Agent
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &agents)
	if err != nil {
		return nil, err
	}
	return agents, nil
}

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

func FetchHealthCenters() ([]clientEntity.HealthCenter, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s/healthcenter", baseURL)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var healthcenters []clientEntity.HealthCenter
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &healthcenters)
	if err != nil {
		return nil, err
	}
	return healthcenters, nil
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

func FetchUsers() ([]clientEntity.User, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s/user", baseURL)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var users []clientEntity.User
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
