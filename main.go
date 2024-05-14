package main

import (
	"encoding/json"
	"fmt"
	logger "help_me_unpack/log"
	unpack "help_me_unpack/src"
	"os"
)

func main() {
	var baseURL = "https://hackattic.com"
	var GetPath = "/challenges/help_me_unpack/problem?access_token="
	var PostPath = "/challenges/help_me_unpack/solve?access_token="
	var token = os.Getenv("ACCESS_TOKEN")
	if token == "" {
		logger.Printf("ACCESS_TOKEN couldn't not be found\n", logger.ERROR)
		os.Exit(1)
	}

	data := unpack.Get(fmt.Sprintf("%s%s%s", baseURL, GetPath, token))
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
	unpack.Post(fmt.Sprintf("%s%s%s", baseURL, PostPath, token), result)
}
