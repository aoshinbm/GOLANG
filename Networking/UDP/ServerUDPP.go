package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen on UDP address and port
	//passing protocol n port number in ResolveUDPAddr method
	addr, _ := net.ResolveUDPAddr("udp", ":8000")
	//ResolveUDPAddr returns a struct and that struct is passed to ListenUDP method
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	// Continuously read incoming data
	for {
		var buf [1024]byte
		n, _, _ := conn.ReadFromUDP(buf[:])
		fmt.Println(string(buf[:n]))
	}
}
