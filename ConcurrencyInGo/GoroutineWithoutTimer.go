package main

import "fmt"

func hello() {
	fmt.Println(" Hello, I am in go routine ")
}

func main() {
	go hello()
	fmt.Println(" In Main ")
}
