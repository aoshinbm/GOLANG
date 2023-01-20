package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, i := range interfaces {
		fmt.Printf("Name: %v\n", i.Name)
		fmt.Printf("MTU: %v\n", i.MTU)
		fmt.Printf("HardwareAddr: %v\n", i.HardwareAddr)
		fmt.Printf("Flags: %v\n", i.Flags)
		fmt.Printf("Addrs: %v\n", i.Addrs)
		fmt.Println("================Divider===================")
	}
}
