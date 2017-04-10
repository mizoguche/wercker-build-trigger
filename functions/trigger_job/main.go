package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/apex/go-apex"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		client := &http.Client{}
		var buffer bytes.Buffer
		var query = fmt.Sprintf("pipelineId=%v&message=%v", os.Getenv("WERCKER_PIPELINE_ID"), url.QueryEscape(os.Getenv("WERCKER_MESSAGE")))
		buffer.WriteString(query)

		request, err := http.NewRequest("POST", "https://app.wercker.com/api/v3/runs/", &buffer)
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("WERCKER_TOKEN")))
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, err := client.Do(request)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return string(body), nil
	})
}
