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

	currentTime := time.Now()
	fmt.Printf("[%v] %v: %v\n", logLevel, currentTime.Format("15:04:05"), fmt.Sprintf(message, formats...))
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
