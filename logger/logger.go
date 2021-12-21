package logger

import (
	"ddos/options"
	"fmt"
	"github.com/fatih/color"
	"time"
)

type Logger struct {
	DesiredLogLevel string
}

func (l Logger) Log(logLevel string, message string) {
	if logLevelToInt(logLevel) < logLevelToInt(l.DesiredLogLevel) || logLevelToInt(logLevel) == 3 {
		return
	}

	currentTime := time.Now()
	fmt.Printf("[%v] %v: %v\n", getColorByLogLevel(logLevelToInt(logLevel))(logLevel), currentTime.Format("15:04:05"), message)
}

func NewLogger(opt options.Options) Logger {
	color.NoColor = opt.NoColor

	return Logger{
		DesiredLogLevel: opt.LogLevel,
	}
}

func logLevelToInt(logLevel string) int {
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

func getColorByLogLevel(logLevel int) func(a ...interface{}) string {
	switch logLevel {
	case 0:
		return color.New(color.FgGreen).SprintFunc()
	case 1:
		return color.New(color.FgYellow).SprintFunc()
	case 2:
		return color.New(color.FgRed).SprintFunc()
	default:
		return color.New(color.FgWhite).SprintFunc()
	}
}
