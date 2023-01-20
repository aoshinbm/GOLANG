package main

import (
	"fmt"
)

func main() {
	go fmt.Println("Hello from a goroutine!")
	fmt.Println("Hello from the main function!")
}
