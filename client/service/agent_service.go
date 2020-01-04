package service

import (
	"github.com/TenaHub/client/entity"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

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

