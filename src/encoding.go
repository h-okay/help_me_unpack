package unpack

import (
	"encoding/base64"
	"errors"
	"fmt"
	logger "help_me_unpack/log"
)

var EncodeInvalidErr = "Encode received no bytes\n"
var DecodeInvalidErr = "Decode received no/invalid bytes\n"
var EncodeParseErr = "Couldn't encode\n"
var DecodeParseErr = "Couldn't decode: %v\n"

func Encode(b []byte) ([]byte, error) {
	if b == nil {
		return nil, errors.New(EncodeInvalidErr)
	}

	encoded := base64.StdEncoding.EncodeToString(b)
	if len(encoded) == 0 {
		return nil, errors.New(EncodeParseErr)
	}

	return []byte(encoded), nil
}

func Decode(b []byte) ([]byte, error) {
	if b == nil {
		return nil, errors.New(DecodeInvalidErr)
	}

	decoded, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		return nil, fmt.Errorf(DecodeParseErr, decoded)
	}

	logger.Printf("Decode wrote %d bytes\n", logger.INFO, len(decoded))
	return []byte(decoded), nil
}
