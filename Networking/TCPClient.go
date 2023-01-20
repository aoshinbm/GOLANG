package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8082")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Send data
	strOfArray := []string{"AOSHIN", "WAKANDA", "FOREVER", "NAMOR"}
	//b := []byte(strOfArray)
	for _, strr := range strOfArray {
		conn.Write([]byte(strr))
	}

	// Read response
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buffer[:n]))
}
