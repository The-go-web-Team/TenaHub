package service

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/TenaHub/client/entity"
	"fmt"
)

func FetchFeedbacks() ([]clientEntity.Comment, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s/feedback/23", baseURL)
	fmt.Println(URL)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var feedbacks []clientEntity.Comment
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &feedbacks)
	if err != nil {
		return nil, err
	}
	return feedbacks, nil
}


