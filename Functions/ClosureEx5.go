package main

import "fmt"

// Define a function that takes a closure as an argument
func doLater(f func()) {
	// Store the closure in a variable for later use
	deferredFunc := f

	// Perform some work
	fmt.Println("Doing work...")

	// Execute the closure at a later time
	deferredFunc()
}

func main() {
	// Define a closure that will be passed to the doLater function
	done := func() {
		fmt.Println("Work is done!")
	}

	// Call the doLater function and pass the closure as an argument
	doLater(done) // Output: "Doing work...", "Work is done!"
}
