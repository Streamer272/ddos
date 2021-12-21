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
	if logLevelToInt(logLevel) < logLevelToInt(l.opt.LogLevel) || logLevelToInt(logLevel) == 3 /* none */ {
		return
	}

	currentTime := time.Now()
	var output = os.Stdout
	if logLevelToInt(logLevel) == 2 /* error */ {
		output = os.Stderr
	}
	if writeToFile && l.opt.OutputFile != os.Stdout {
		output = l.opt.OutputFile
		l.Log(logLevel, message, false)
		color.NoColor = true
	}
	fmt.Fprintf(output, "[%v] %v: %v\n", getColorFuncByLogLevel(logLevelToInt(logLevel))(logLevel), currentTime.Format("15:04:05"), message)

	if writeToFile && l.opt.OutputFile != os.Stdout {
		color.NoColor = l.opt.NoColor
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
