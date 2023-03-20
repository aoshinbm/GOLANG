package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listenn, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}

	defer listenn.Close()

	fmt.Println("SERVER listening on port 8082")

	for {
		connect, err := listenn.Accept()
		if err != nil {
			fmt.Println("ERROR accepting connection", err)
			continue
		}
		go handleConnect(connect)
	}
}

func handleConnect(conn net.Conn) {
	read := bufio.NewReader()
}
