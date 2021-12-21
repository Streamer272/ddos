package logger

import (
	"ddos/options"
	"fmt"
	"github.com/fatih/color"
	"os"
	"time"
)

type Logger struct {
	opt options.Options
}

func (l Logger) Log(logLevel string, message string, writeToFile bool) {
	if logLevelToInt(logLevel) < logLevelToInt(l.opt.LogLevel) || logLevelToInt(logLevel) == 3 {
		return
	}

	currentTime := time.Now()
	fmt.Printf("[%v] %v: %v\n", getColorFuncByLogLevel(logLevelToInt(logLevel))(logLevel), currentTime.Format("15:04:05"), message)

	if l.opt.OutputFile != "" && writeToFile {
		file, err := os.OpenFile(l.opt.OutputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			l.Log("ERROR", fmt.Sprintf("Couldn't open file, %v...", err), false)
		}
		if _, err = file.WriteString(fmt.Sprintf("[%v] %v: %v\n", logLevel, currentTime.Format("15:04:05"), message)); err != nil {
			l.Log("ERROR", fmt.Sprintf("Couldn't write to file, %v...", err), false)
		}
		if err = file.Close(); err != nil {
			l.Log("ERROR", fmt.Sprintf("Couldn't close file, %v...", err), false)
		}
	}
}

func NewLogger(opt options.Options) Logger {
	color.NoColor = opt.NoColor

	return Logger{
		opt: opt,
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

func getColorFuncByLogLevel(logLevel int) func(a ...interface{}) string {
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
