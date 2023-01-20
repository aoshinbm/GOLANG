package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println(" Hello, I am in go routine ")
}

func main() {
	go hello()
	time.Sleep(time.Second) //jz a hack not gud practice
	fmt.Println(" In Main ")
}
