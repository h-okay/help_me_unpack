package log

import (
	"fmt"
	"time"
)

// Level is the log level
type Level int

// Allowed values for Level
const (
	INFO Level = iota
	WARNING
	ERROR
)

// Desired time format for log strings
var timeFormat = "2006-02-01T15:04:05Z"

// Printf uses given timeFormat and arguments to construct a log string
func Printf(format string, logLevel Level, args ...any) {
	convertedLogType := convertLogType(logLevel)
	format = fmt.Sprintf("%s [%s] %s", time.Now().Format(timeFormat), convertedLogType, format)
	fmt.Printf(format, args...)
}

// convertLogType converts log enums to string representations
func convertLogType(logType Level) string {
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
