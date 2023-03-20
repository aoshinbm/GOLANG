// Write a program to shallow copy a struct in Go.
package main

import (
	"fmt"
)

type Person struct {
	name string
}

func main() {

	p := Person{name: "Anette"}

	// shallow copy
	p2 := p
	fmt.Printf("Original Struct: %+v\n", p)

	// change the copied name
	p2.name = "Not Dulquer Salman"
	fmt.Printf("Original Struct: %+v\n", p)
	fmt.Printf("Copied Struct :%+v\n", p2)
}
