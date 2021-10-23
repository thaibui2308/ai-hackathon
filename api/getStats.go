package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/thaibui2308/ai-hackathon/models"
)

func GetStats(url string) (models.Commit, error) {
	response, err := http.Get(url)
	var userCommit models.Commit
	if err != nil {
		return models.Commit{}, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return models.Commit{}, err
	}
	json.Unmarshal(responseData, &userCommit)
	return userCommit, nil
}
