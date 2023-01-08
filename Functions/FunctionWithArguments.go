package main

import "fmt"

// Function accepting arguments i.e two arguments of int type
func add(x int, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

func main() {
	// Passing arguments
	add(20, 30)
}
