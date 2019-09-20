package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/fatih/color"
)

func makePetition(method, url string, body []byte, token *string, params map[string]string) (map[string]interface{}, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Origin", siteFrontendURL)

	if token != nil {
		req.Header.Add("Authorization", *token)
	}

	q := req.URL.Query()

	for key, value := range params {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := make(map[string]interface{})

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		data, _ := json.Marshal(response)

		red := color.New(color.FgRed).SprintFunc()
		log.Printf("\nThe server has responded with: \"%s\" to the petition: %s on: %s\n", red(string(data[:])), green(req.Method), green(req.URL))
	}

	return response, nil
}
