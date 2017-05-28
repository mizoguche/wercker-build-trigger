package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/apex/go-apex"
)

type payload struct {
	PipelineId string `json:"pipelineId"`
	Message    string `json:"message"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		client := &http.Client{}

		p := payload{
			PipelineId: os.Getenv("WERCKER_PIPELINE_ID"),
			Message:    os.Getenv("WERCKER_MESSAGE"),
		}
		b, _ := json.Marshal(p)

		var buffer bytes.Buffer
		buffer.WriteString(string(b))

		request, err := http.NewRequest("POST", "https://app.wercker.com/api/v3/runs/", &buffer)
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("WERCKER_TOKEN")))
		request.Header.Add("Content-Type", "application/json")

		resp, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		return string(body), nil
	})
}
