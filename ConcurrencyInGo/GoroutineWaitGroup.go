package main

import (
	"fmt"
	"sync"
)

func hello(wg *sync.WaitGroup) {
	fmt.Println(" Hello, I am in go routine ")
	wg.Done()
}
func hello1(wg *sync.WaitGroup) {
	fmt.Println(" Learning GOROUTINE ")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Anonymous Function")
		wg.Done()
	}()
	go hello(&wg)
	//	go hello1(&wg)
	wg.Wait()
	fmt.Println("Finishing - In Main ")
}
