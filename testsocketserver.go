package main

import "net"
import "fmt"
import "bufio"

func main() {
	fmt.Println("Start server...")

	ln, _ := net.Listen("tcp", ":8080")

	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Message Received: %v\n", string(message))
	}
}
