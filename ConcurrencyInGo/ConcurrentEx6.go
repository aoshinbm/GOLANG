package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printNumbers(num int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Goroutine %d: %d\n", num, i)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go printNumbers(1)
	go printNumbers(2)
	fmt.Println("Hello from the main function!")
	wg.Wait()
}
