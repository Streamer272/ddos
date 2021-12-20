package main

import (
	"bufio"
	"log"
	"net"
)
import "fmt"

func main() {
	fmt.Println("Start server...")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		log.Printf("GOT CONN %v \n", conn)

		go func() {
			for {
				message, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					log.Printf("Couldn't read message, %v...\n", err)
				}

				log.Printf("Message Received: %v\n", string(message))
			}
		}()
	}
}
