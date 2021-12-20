package main

import (
	"io/ioutil"
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

		fmt.Printf("GOT CONN %v \n", conn)

		go func() {
			for {
				message, err := ioutil.ReadAll(conn)
				if err != nil {
					continue
				}

				fmt.Printf("Message Received: %v\n", string(message))
			}
		}()
	}
}
