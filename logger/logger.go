package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	DesiredLogLevel string
}

func (l Logger) Log(logLevel string, message string, formats ...interface{}) {
	if LogLevelToInt(logLevel) < LogLevelToInt(l.DesiredLogLevel) {
		return
	}

	fmt.Printf("[%v] %v: %v\n", logLevel, time.Now(), fmt.Sprintf(message, formats))
}

func NewLogger(desiredLogLevel string) Logger {
	return Logger{
		DesiredLogLevel: desiredLogLevel,
	}
}

func LogLevelToInt(logLevel string) int {
	switch logLevel {
	case "INFO":
		return 0
	case "WARN":
		return 1
	case "ERROR":
		return 2
	default:
		return 3
	}
}
