package unpack

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	logger "help_me_unpack/log"
)

var EncodeInvalidErr = "Encode received no bytes\n"
var DecodeInvalidErr = "Decode received no/invalid bytes\n"
var EncodeParseErr = "Couldn't encode %v\n"
var DecodeParseErr = "Couldn't decode %v: %v\n"

func Encode(b []byte) ([]byte, error) {
	if b == nil {
		return nil, errors.New(EncodeInvalidErr)
	}

	encoded := base64.StdEncoding.EncodeToString(b)
	if len(encoded) == 0 {
		return nil, fmt.Errorf(EncodeParseErr, b)
	}

	return []byte(encoded), nil
}

func Decode(s string) ([]byte, error) {
	if s == "" {
		return nil, errors.New(DecodeInvalidErr)
	}

	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf(DecodeParseErr, s, err)
	}

	logger.Printf("Decode wrote %d bytes\n", logger.INFO, len(decoded))
	return decoded, nil
}

func GetFieldFromJSON(b []byte, field string) (string, error) {
	var jsn map[string]any
	err := json.Unmarshal(b, &jsn)
	if err != nil {
		return "", err
	}
	return jsn[field].(string), nil
}
