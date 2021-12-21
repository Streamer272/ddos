package options

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

const (
	Version = "1.0"
)

func Parse() Options {
	parser := argparse.NewParser("ddos", "Runs DDOS attack on desired server")
	parser.ExitOnHelp(false)

	delay := parser.Int("d", "delay", &argparse.Options{Required: false, Help: "Packet delay, 0 for none", Default: 10})
	maxRetryCount := parser.Int("r", "max-retry-count", &argparse.Options{Required: false, Help: "Max retry count, 0 for none", Default: 0})
	workerCount := parser.Int("w", "worker-count", &argparse.Options{Required: false, Help: "Worker thread count, 0 for none", Default: 0})
	address := parser.String("a", "address", &argparse.Options{Required: false, Help: "Address to DDOS", Default: ""})
	message := parser.String("m", "message", &argparse.Options{Required: false, Help: "Custom message to send", Default: ""})
	logLevel := parser.Selector("l", "log-level", []string{"NONE", "ERROR", "WARN", "INFO"}, &argparse.Options{Required: false, Help: "Log level", Default: "INFO"})
	http := parser.Flag("H", "http", &argparse.Options{Required: false, Help: "Use HTTP message", Default: false})
	ignoreError := parser.Flag("i", "ignore-error", &argparse.Options{Required: false, Help: "Do not terminate program on error", Default: false})
	noColor := parser.Flag("N", "no-color", &argparse.Options{Required: false, Help: "Do not display colored output", Default: false})
	version := parser.Flag("v", "version", &argparse.Options{Required: false, Help: "Display version info", Default: false})
	outputFile := parser.File("o", "output", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600, &argparse.Options{Required: false, Help: "Additional output file", Default: os.Stdout})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(0)
	}

	if *version {
		fmt.Printf("DDOS version %v\n", Version)
		os.Exit(0)
	}

	if *address == "" {
		fmt.Print(parser.Usage("[-a|--address] is missing"))
		os.Exit(1)
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
		WorkerCount:   *workerCount,
		Address:       *address,
		Message:       *message,
		LogLevel:      *logLevel,
		IgnoreError:   *ignoreError,
		Http:          *http,
		NoColor:       *noColor,
		OutputFile:    outputFile,
	}
}
