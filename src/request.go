package unpack

import (
	"bytes"
	logger "help_me_unpack/log"
	"io"
	"net/http"
	"os"
	"time"
)

// Get makes a GET request to url and returns the data
func Get(url string) []byte {
	netClient := &http.Client{
		Timeout: time.Second * 10,
	}

	logger.Printf("Making a GET request\n", logger.INFO)
	resp, err := netClient.Get(url)
	if err != nil {
		logger.Printf("Couldn't GET: %v", logger.WARNING, err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	logger.Printf("Reading byte stream from response body\n", logger.INFO)
	bodyBytes, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		logger.Printf("Failed to read byte stream: %v", logger.ERROR, err.Error())
		os.Exit(1)
	}

	return bodyBytes
}

// Post makes a POST request to url
func Post(url string, data []byte) {
	netClient := &http.Client{
		Timeout: time.Second * 10,
	}

	logger.Printf("Making a POST request\n", logger.INFO)
	resp, err := netClient.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		logger.Printf("Couldn't POST results: %v", logger.WARNING, err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	logger.Printf("Response Status: %s\n", logger.INFO, resp.Status)
	logger.Printf("Response Headers: %s\n", logger.INFO, resp.Header)

	body, _ := io.ReadAll(resp.Body)
	logger.Printf("Response Body: %s\n", logger.INFO, string(body))
}
