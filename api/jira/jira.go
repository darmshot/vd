package jira

import (
	"encoding/json"
	"fmt"
	"github.com/darmshot/vd/config"
	"io"
	"log"
	"net/http"
)

type Issue struct {
	Fields IssueFields `json:"fields"`
}

type IssueFields struct {
	Summary string `json:"summary"`
}

func GetIssueSummary(issueIdOrKey string) string {
	//https://ontid.atlassian.net/browse/CREOS-859
	URL := fmt.Sprintf("https://ontid.atlassian.net/rest/api/3/issue/%s", issueIdOrKey)

	client := http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Accept":        {"application/json"},
		"Authorization": {"Basic " + config.JiraKey},
	}

	resp, err := client.Do(req)
	if err != nil {
		//Handle Error
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	bytes, err := io.ReadAll(resp.Body)

	//log.Println(string(bytes))

	var response Issue
	errUnmarshal := json.Unmarshal(bytes, &response)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}

	return response.Fields.Summary
}
