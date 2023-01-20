// Client Socket
package main

import (
	"fmt"
	"net"
)

func main() {
	// Create a socket
	socket, err := net.Dial("tcp", "www.google.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer socket.Close()

	// Send data
	//\r\n are escape sequence
	socket.Write([]byte("GET / HTTP/1.1\r\n\r\n"))

	// Read response
	buffer := make([]byte, 4096)
	n, err := socket.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buffer[:n]))
}
