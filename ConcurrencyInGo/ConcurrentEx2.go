package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printGreeting() {
	fmt.Println("Hello from a goroutine!")
	wg.Done()
}

func main() {
	wg.Add(1)
	go printGreeting()
	fmt.Println("Hello from the main function!")
	wg.Wait()
}

/*
It utilizes the built-in "sync" package to create a WaitGroup,
which is used to wait for a collection of goroutines to finish executing.

In the main function, the WaitGroup's counter is incremented by 1 using the Add() method.
Then a goroutine is created using the "go" keyword, which runs the printGreeting() function concurrently
with the main function. The main function then prints "Hello from the main function!" and
the WaitGroup's Wait() method is called, which blocks execution until the WaitGroup counter is 0.

In the printGreeting() function, "Hello from a goroutine!" is printed to the console,
and then the Done() method is called on the WaitGroup, which decrements the counter by 1.
Once the Wait() method in the main function sees that the counter is 0, it unblocks and the program exits.

*/
