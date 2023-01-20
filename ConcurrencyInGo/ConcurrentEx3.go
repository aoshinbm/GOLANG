package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printNumbers()
	fmt.Println("Hello from the main function!")
	time.Sleep(1 * time.Second)
}
