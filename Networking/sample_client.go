// CLIENT
package main

import (
	"fmt"
	"net"
)

func main() {
	lisn, err := net.ResolveUDPAddr("udp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
	}
	connec, err := net.DialUDP("udp", nil, lisn)
	if err != nil {
		fmt.Println(err)
	}
	defer connec.Close()
	connec.Write([]byte("HEloo SERVERRRR"))

	lisn2, err := net.ResolveUDPAddr("udp", "localhost:8082")
	if err != nil {
		fmt.Println(err)
	}
	con, err := net.DialUDP("udp", nil, lisn2)
	if err != nil {
		fmt.Println(err)
	}
	defer con.Close()
	con.Write([]byte("HEloo "))

	buffer := make([]byte, 4096)
	n, _, err := connec.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(n)

	fmt.Println(string(buffer[:n]))
}
