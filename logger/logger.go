package logger

import (
	"fmt"
	"github.com/Streamer272/ddos/options"
	"github.com/fatih/color"
	"os"
	"time"
)

type Logger struct {
	opt      options.Options
	disabled bool
}

func (log *Logger) Log(logLevel string, message string, writeToFile bool) {
	if logLevelToInt(logLevel) < logLevelToInt(log.opt.LogLevel) || logLevelToInt(logLevel) == 3 /* none */ || log.disabled {
		return
	}

	currentTime := time.Now()
	output := os.Stdout
	if logLevelToInt(logLevel) == 2 /* error */ {
		output = os.Stderr
	}
	if writeToFile && log.opt.OutputFile != "" {
		file, err := os.OpenFile(log.opt.OutputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			log.Log("ERROR", fmt.Sprintf("Couldn't write to %v...", log.opt.OutputFile), false)
		}
		output = file
		log.Log(logLevel, message, false)
		color.NoColor = true
	}
	fmt.Fprintf(output, "[%v] %v: %v\n", getColorFuncByLogLevel(logLevelToInt(logLevel))(logLevel), currentTime.Format("15:04:05"), message)

	if writeToFile && log.opt.OutputFile != "" {
		color.NoColor = log.opt.NoColor
	}
}

func (log *Logger) Disable() {
	log.disabled = true
}

func (log *Logger) Enable() {
	log.disabled = false
}

func NewLogger(opt options.Options) Logger {
	color.NoColor = opt.NoColor

	return Logger{
		opt:      opt,
		disabled: false,
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
