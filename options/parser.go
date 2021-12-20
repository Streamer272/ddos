package options

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func Parse() Options {
	parser := argparse.NewParser("print", "Runs DDOS attack on desired server")

	address := parser.String("a", "address", &argparse.Options{Required: true, Help: "Address to DDOS"})
	delay := parser.Int("d", "delay", &argparse.Options{Required: false, Help: "Packet Delay", Default: 0})
	http := parser.Flag("H", "http", &argparse.Options{Required: false, Help: "Whether to use HTTP", Default: false})

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
