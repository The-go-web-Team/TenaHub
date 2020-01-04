package service

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/TenaHub/client/entity"
	"fmt"
)

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
