package main

import (
	"net"
)

func main() {
	// Dial UDP address and port
	conn, _ := net.Dial("udp", "localhost:8000")
	defer conn.Close()

	// Send data
	_, _ = conn.Write([]byte("Hello, UDP!"))
}
