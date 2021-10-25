package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func MergePullRequest(url string) {
	client := &http.Client{}
	request, _ := http.NewRequest("PUT", url, strings.NewReader("Merge this pull request"))
	request.ContentLength = 23
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("   ", response.StatusCode)
		hdr := response.Header
		for key, value := range hdr {
			fmt.Println("   ", key, ":", value)
		}
		fmt.Println(contents)
	}
}
