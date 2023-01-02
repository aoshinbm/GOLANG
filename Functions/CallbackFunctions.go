package main

import "fmt"

// Define a function that takes a callback as an argument
func doWork(callback func()) {
	// Perform some work
	fmt.Println("Doing work...")

	// Call the callback
	callback()
}

func main() {
	// Define a function that will be passed as a callback
	done := func() {
		fmt.Println("Work is done!")
	}

	// Call the doWork function and pass the done function as a callback
	doWork(done) // Output: "Doing work...", "Work is done!"
}
