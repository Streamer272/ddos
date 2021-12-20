package options

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func Parse() Options {
	parser := argparse.NewParser("print", "Prints provided string to stdout")
	address := parser.String("a", "address", &argparse.Options{Required: true, Help: "Address to DDOS"})
	delay := parser.Int("d", "delay", &argparse.Options{Required: false, Help: "Packet Delay", Default: 0})
	http := parser.Flag("p", "http", &argparse.Options{Required: false, Help: "Whether to use HTTP", Default: false})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	return Options{
		Address: *address,
		Delay:   *delay,
		Http:    *http,
	}
}
