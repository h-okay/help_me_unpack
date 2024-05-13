package unpack

import (
	"encoding/base64"
	logger "help_me_unpack/log"
)

func Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Decode(s string) string {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		logger.Printf("Coulnd't decode base64", logger.ERROR, err.Error())
	}
	return string(decoded)
}
