package service

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/TenaHub/client/entity"
	"fmt"
)

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

