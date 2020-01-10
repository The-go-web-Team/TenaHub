package service

import (
	"github.com/TenaHub/client/entity"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"net/url"
	"errors"
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

// Authenticate authenticates user
func AgentAuthenticate(agent *clientEntity.Agent) (*clientEntity.Agent, error) {
	URL := fmt.Sprintf("%s/%s", baseURL, "agent")

	formval := url.Values{}
	formval.Add("email", agent.Email)
	formval.Add("password", agent.Password)

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
		Content clientEntity.Agent
	}{}

	err = json.Unmarshal(body, &respjson)

	fmt.Println(respjson)

	if respjson.Status == "error" {
		return nil, errors.New("error")
	}
	return &respjson.Content, nil
}


