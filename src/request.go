package unpack

import (
	"encoding/json"
	"fmt"
	logger "help_me_unpack/log"
	"net/http"
	"os"
)

type Data struct {
	Bytes string `json:"bytes"`
}

func (d Data) String() string {
	return fmt.Sprintf("Bytes: %s", d.Bytes)
}

func RequestNewData() Data {

	var URL = fmt.Sprintf(
		"https://hackattic.com/challenges/help_me_unpack/problem?access_token=%s",
		os.Getenv("ACCESS_TOKEN"),
	)

	logger.Printf("Making a request\n", logger.INFO)
	resp, err := http.Get(URL)
	if err != nil {
		logger.Printf("Couldn't GET: %v", logger.INFO, err.Error())
		os.Exit(1)
	}

	logger.Printf("Parsing JSON\n", logger.INFO)
	var data Data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		logger.Printf("Couldn't parse JSON: %v", logger.INFO, err.Error())
		os.Exit(1)
	}

	logger.Printf("Got the following data => %s\n", logger.INFO, data.String())
	return data
}
