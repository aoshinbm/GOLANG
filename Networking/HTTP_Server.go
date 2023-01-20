package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// Listen on port 8080
	listener, err := net.Listen("tcp", ":8080")
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

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	// Read the incoming request
	request := make([]byte, 1024)
	n, err := conn.Read(request)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	request = request[:n]

	// Log the request
	fmt.Println(string(request))

	// Parse the request
	requestLines := strings.Split(string(request), "\r\n")
	requestLine := strings.Split(requestLines[0], " ")
	method := requestLine[0]
	path := requestLine[1]

	// Handle the GET request
	if method == "GET" {
		fmt.Println("Handling GET request for path:", path)
		switch path {
		case "/":
			handleHome(conn)
		case "/about":
			handleAbout(conn)
		default:
			handle404(conn)
		}
	}
}

func handleHome(conn net.Conn) {
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html\r\n" +
		"Content-Length: 16\r\n" +
		"\r\n" +
		"<h1>Home</h1>"
	conn.Write([]byte(response))
}

func handleAbout(conn net.Conn) {
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html\r\n" +
		"Content-Length: 17\r\n" +
		"\r\n" +
		"<h1>About</h1>"
	conn.Write([]byte(response))
}

func handle404(conn net.Conn) {
	response := "HTTP/1.1 404 Not Found\r\n" +
		"Content-Type: text/html\r\n" +
		"Content-Length: 17\r\n" +
		"\r\n" +
		"<h1>404 Not Found</h1>"
	conn.Write([]byte(response))
}
