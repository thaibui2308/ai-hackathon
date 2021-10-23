package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/thaibui2308/ai-hackathon/models"
)

var baseUrl = "https://api.github.com/repos/" // full url should look like this https://api.github.com/repos/cheggaaa/pb/pulls?state=188

func GetPullRequest(username, repo, pullId string) (models.PullRequests, error) {
	url := baseUrl + username + "/" + repo + "/pulls/" + pullId + "/commits"
	var responseData models.PullRequests
	info, err := http.Get(url)
	if err != nil {
		err = errors.New("can not find this pull request")
		return models.PullRequests{}, err
	}

	data, err := ioutil.ReadAll(info.Body)
	if err != nil {
		return models.PullRequests{}, err

	}
	json.Unmarshal(data, &responseData)
	return responseData, nil
}
