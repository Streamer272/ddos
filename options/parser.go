package options

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func Parse() Options {
	parser := argparse.NewParser("print", "Runs DDOS attack on desired server")
	parser.ExitOnHelp(true)

	address := parser.String("a", "address", &argparse.Options{Required: true, Help: "Address to DDOS"})
	delay := parser.Int("d", "delay", &argparse.Options{Required: false, Help: "Packet Delay", Default: 0})
	http := parser.Flag("H", "http", &argparse.Options{Required: false, Help: "Whether to use HTTP", Default: false})
	maxRetryCount := parser.Int("r", "max-retry-count", &argparse.Options{Required: false, Help: "Max retry count, 0 for none", Default: 0})
	message := parser.String("m", "message", &argparse.Options{Required: false, Help: "Message to send", Default: ""})
	ignoreError := parser.Flag("i", "ignore-error", &argparse.Options{Required: false, Help: "Whether to ignore error", Default: false})
	// TODO: add logging level

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
		Address:       *address,
		Delay:         *delay,
		Http:          *http,
		MaxRetryCount: *maxRetryCount,
		Message:       *message,
		IgnoreError:   *ignoreError,
	}
}
