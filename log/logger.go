package log

import (
	"fmt"
	"time"
)

type LogType int

const (
	INFO LogType = iota
	WARNING
	ERROR
)

var TimeFormat = "2006-02-01T15:04:05Z"

func Printf(format string, logType LogType, args ...any) {
	convertedLogType := convertLogType(logType)
	format = fmt.Sprintf("%s [%s] %s", time.Now().Format(TimeFormat), convertedLogType, format)
	fmt.Printf(format, args...)
}

func convertLogType(logType LogType) string {
	switch logType {
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}
