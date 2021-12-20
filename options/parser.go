package options

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func Parse() Options {
	parser := argparse.NewParser("print", "Runs DDOS attack on desired server")
	parser.ExitOnHelp(true)

	delay := parser.Int("d", "delay", &argparse.Options{Required: false, Help: "Packet Delay", Default: 0})
	maxRetryCount := parser.Int("r", "max-retry-count", &argparse.Options{Required: false, Help: "Max retry count, 0 for none", Default: 0})
	address := parser.String("a", "address", &argparse.Options{Required: true, Help: "Address to DDOS"})
	message := parser.String("m", "message", &argparse.Options{Required: false, Help: "Message to send", Default: ""})
	logLevel := parser.Selector("l", "log-level", []string{"NONE", "ERROR", "WARN", "INFO"}, &argparse.Options{Required: false, Help: "Log level", Default: "INFO"})
	http := parser.Flag("H", "http", &argparse.Options{Required: false, Help: "Whether to use HTTP", Default: false})
	ignoreError := parser.Flag("i", "ignore-error", &argparse.Options{Required: false, Help: "Whether to ignore error", Default: false})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Printf("%v", parser.Usage(err))
	}

	if *http && *message != "" {
		fmt.Printf("%v", parser.Usage("Cannot use `[-H|--http]` while using `[-m|--message]`"))
		os.Exit(1)
	}
	if *ignoreError && *maxRetryCount != 0 {
		fmt.Printf("%v", parser.Usage("Cannot use `[-i|--ignore-error]` while using `[-r|--max-retry-count]`"))
		os.Exit(1)
	}

	return Options{
		Delay:         *delay,
		MaxRetryCount: *maxRetryCount,
		Address:       *address,
		Message:       *message,
		LogLevel:      *logLevel,
		IgnoreError:   *ignoreError,
		Http:          *http,
	}
}
