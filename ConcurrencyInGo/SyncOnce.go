package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go dostuff()
	go dostuff()
	wg.Wait()
}

func dostuff() {

	//it will run only once
	once.Do(setup)

	fmt.Println("hello, I am doing an important stuff")
	fmt.Println("Making you learn once initialization from goroutine")

	wg.Done()
}

func setup() {
	fmt.Println("Doing Initialiation now...")
}
