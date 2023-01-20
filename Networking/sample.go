/*
Q1) Write a Go program that creates a UDP server that listens on port 8000.
The server should respond to incoming datagrams by sending the number of bytes
received back to the client and then closing the connection.

Advance Question:
Instead of sending back only the number of bytes received,
we will send back the number of bytes received, plus the
content of the data sent by the client, and also, we will
listen for the incoming data on multiple ports and the
client must specify the port from which it will send the data.
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	lisn, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	connec, err := net.ListenUDP("udp", lisn)
	if err != nil {
		fmt.Println(err)
	}
	defer connec.Close()

	for {
		buffer := make([]byte, 4096)
		n, lisn, err := connec.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleData(connec, lisn, buffer[:n])
	}

	lisn2, err := net.ResolveUDPAddr("udp", ":8082")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := net.ListenUDP("udp", lisn2)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	for {
		buffer := make([]byte, 4096)
		n, lisn2, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleData(conn, lisn2, buffer[:n])
	}

}
func handleData(conn *net.UDPConn, lisnn *net.UDPAddr, data []byte) {
	conn.WriteToUDP([]byte("HELOO CLIENT! hwz u ??"), lisnn)

	fmt.Println("Listening :", string(data))
	fmt.Println(data)
}
