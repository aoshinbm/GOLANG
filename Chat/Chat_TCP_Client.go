package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//read data from server
	readd := bufio.NewReader(conn)
	for {
		mssge, err := readd.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(mssge)
	}

	//send data to the server
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		mssge := scanner.Text() + "\n"
		_, err := conn.Write([]byte(mssge))
		if err != nil {
			log.Println(err)
			break
		}
	}
}
