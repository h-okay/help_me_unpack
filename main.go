package main

import (
	"encoding/json"
	"fmt"
	logger "help_me_unpack/log"
	unpack "help_me_unpack/src"
	"os"
)

func main() {
	var BASE_URL = "https://hackattic.com"
	var GET_PATH = "/challenges/help_me_unpack/problem?access_token="
	var POST_PATH = "/challenges/help_me_unpack/solve?access_token="
	var TOKEN = os.Getenv("ACCESS_TOKEN")

	data := unpack.Get(fmt.Sprintf("%s%s%s", BASE_URL, GET_PATH, TOKEN))
	bytesString, err := unpack.GetFieldFromJSON(data, "bytes")
	if err != nil {
		logger.Printf("Something went wrong: %v", logger.ERROR, err.Error())
		os.Exit(1)
	}
	unpacked, err := unpack.Unpack((bytesString))
	if err != nil {
		logger.Printf("Something went wrong: %v", logger.ERROR, err.Error())
		os.Exit(1)
	}

	result, err := json.Marshal(&unpacked)
	if err != nil {
		logger.Printf("Something went wrong: %v", logger.ERROR, err.Error())
		os.Exit(1)
	}
	unpack.Post(fmt.Sprintf("%s%s%s", BASE_URL, POST_PATH, TOKEN), result)
}
