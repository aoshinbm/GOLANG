package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(250 * time.Millisecond)
	}
}

func printLetters() {
	for i := 'a'; i <= 'j'; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(400 * time.Millisecond)
	}
}

func main() {
	// Start a new goroutine
	go printNumbers()
	// Do other work concurrently
	go printLetters()
	fmt.Println("main function done!")
	// sleep for a while to let goroutines finish
	time.Sleep(5 * time.Second)
}

/*
It's also important to note that when the main function completes,
it terminates the program and kills any running goroutines.
Here, we added a sleep time
after main function done to let the goroutines print their output.
*/
