package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		//Accept incoming client connections
		connec, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		//goroutine to handle the communication with the client
		go handleConnection(connec)
	}
}
func handleConnection(connec net.Conn) {
	defer connec.Close()

	clients := &Client{connec, make(chan string)}
	clients[client] = true

	readData := bufio.NewReader(connec)
	for {
		message, err := readData.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		//broadcast message to all clients
		for c := range clients {
			c.ch <- message
		}
	}
	//remove the client from the connected clients
	delete(clients, client)
}
