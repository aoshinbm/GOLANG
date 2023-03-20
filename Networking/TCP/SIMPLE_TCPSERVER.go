package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Listen on a port
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		// Accept a connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		log.Printf("Client connected: %s\n", conn.RemoteAddr())
		// Handle the connection
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Send data
	conn.Write([]byte("From Server Side : Hello, Client"))

	// Read response
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buffer[:n]))
}
