package youtrack

import (
	"encoding/json"
	"fmt"
	"github.com/darmshot/vd/config"
	"io"
	"log"
	"net/http"
)

type Issue struct {
	Summary string `json:"summary"`
}

func GetIssueSummary(issueIdOrKey string) string {
	//SPORTDEFEND-121 добавить loading=lazy в виджет товаров
	URL := fmt.Sprintf("%s/issues/%s", config.YoutrackBaseUrl, issueIdOrKey)

	client := http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("fields", "summary")
	req.URL.RawQuery = q.Encode()

	req.Header = http.Header{
		"Accept":        {"application/json"},
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + config.YoutrackKey},
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

	return response.Summary
}
