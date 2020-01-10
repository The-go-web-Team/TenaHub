package service

import (
	"github.com/TenaHub/client/entity"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func FetchPendingServices(id uint) ([]clientEntity.Service, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s/services/pending/%d", baseURL, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var services []clientEntity.Service
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &services)
	if err != nil {
		return nil, err
	}
	return services, nil
}
func FetchService(id uint) ([]clientEntity.Service, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s/service/%d", baseURL, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var service []clientEntity.Service
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &service)
	if err != nil {
		return nil, err
	}
	return service, nil
}

