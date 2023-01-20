package main

import (
	"log"
	"net"
)

func main() {
	// create listener
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		//catching the error
		log.Fatalf("Failed to create listener: %v", err)
	}
	defer listener.Close()

	// set up logger
	log.Println("Server starting...")

	for {
		// accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		// handle connection goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// read data from connection
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		log.Printf("Failed to read data: %v", err)
		return
	}

	// log received data
	log.Printf("Received data: %s", string(data[:n]))

	// write response
	_, err = conn.Write([]byte("Hello, client"))
	if err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
