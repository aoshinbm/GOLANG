package main

import (
	"fmt"
	"net"
)

func main() {
	// Resolve server address
	addr, err := net.ResolveUDPAddr("udp", "localhost:8082")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Dial to server
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Send data
	conn.Write([]byte("Hello, server!"))

	// Read response
	buffer := make([]byte, 4096)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buffer[:n]))
}
