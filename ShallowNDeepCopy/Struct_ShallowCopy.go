// Write a program to shallow copy a struct in Go.
package main

import (
	"fmt"
)

/*
Shallow copy in go is an assignment.
Everything in go is passed by value so assigning variable
with a struct to another variable makes a copy of that struct.
The same thing goes if you pass a variable to a function.
Method always gets a copy of your variable
(if you pass in a pointer, it gets copy of the pointer! —
	but not copy of the value underneath the pointer).
*/

type Person struct {
	name string
}

func main() {
	p := Person{name: "Anette"}
	fmt.Printf("Original Struct: %+v\n", p)

	// deep copy
	p2 := &p

	// change the copied name
	p2.name = "Not Dulquer Salman"
	fmt.Printf("Original Struct: %+v\n", p)
	fmt.Printf("Copied Struct :%+v\n", *p2)

}
