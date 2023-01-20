package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello, server!"))
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := buf[:n]
	newUser := &User{}
	json.Unmarshal(data, newUser) // deserialize json to struct
	fmt.Println(newUser)

}
