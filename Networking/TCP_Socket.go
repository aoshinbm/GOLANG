package main

import (
	"fmt"
	"net"
)

func main() {
	// Create a TCP socket
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Connected to", conn.RemoteAddr())
	}
}
