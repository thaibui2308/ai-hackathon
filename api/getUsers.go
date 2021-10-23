package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/thaibui2308/ai-hackathon/models"
)

func GetUserInfo(url string) (models.User, error) {
	response, err := http.Get(url)
	var userInfo models.User
	if err != nil {
		return models.User{}, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return models.User{}, err
	}
	json.Unmarshal(responseData, &userInfo)
	return userInfo, nil
}
