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
	// Listen on port 8080
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Listening on :8080...")

	for {
		// Wait for a connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	u1 := User{1, "Eric", "er@gmail.co"}
	u2 := User{2, "Rose", "rose@gmail.co"}
	u3 := User{3, "Rio", "rio@gmail.co"}
	users := []User{
		u1,
		u2,
		u3,
	}

	//user := &User{ID: 1, Name: "John Smith", Email: "john.smith@example.com"}
	fmt.Println(users)
	fmt.Printf("%T \n", users)
	// serialize the struct to json
	data, _ := json.MarshalIndent(users, " ", "\t")
	fmt.Println(string(data))
	fmt.Printf("\n %T", data)

	// Echo incoming data
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
	}
	_, err = conn.Write([]byte("Hello, client!"))
	if err != nil {
		fmt.Println(err)
		return
	}

}
