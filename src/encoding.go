package unpack

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	logger "help_me_unpack/log"
)

// EncodeInvalidErr thrown when invalid values passed to Encode
var EncodeInvalidErr = "Encode received no bytes\n"

// DecodeInvalidErr thrown when invalid values passed to Decode
var DecodeInvalidErr = "Decode received no/invalid bytes\n"

// EncodeParseErr thrown when given input couldn't be encoded correctly
var EncodeParseErr = "Couldn't encode %v\n"

// DecodeParseErr thrown when given input couldn't be decoded correctly
var DecodeParseErr = "Couldn't decode %v: %v\n"

// Encode encodes the input to base64
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

// Decode decodes the base64 input
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

// GetFieldFromJSON gets a certain field and returns it as a string
func GetFieldFromJSON(b []byte, field string) (string, error) {
	var jsn map[string]any
	err := json.Unmarshal(b, &jsn)
	if err != nil {
		return "", err
	}

	value, ok := jsn[field]
	if !ok {
		return "", err
	}

	return value.(string), nil
}
