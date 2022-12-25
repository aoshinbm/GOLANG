package main

import "fmt"

func main() {
	//remember to use PRINTF() for FORMAT SPECIFIER (like for %d)

	var s string = "Ferugson"
	//to print the value of the string within a statement using %v
	fmt.Printf("Hey Mr. %v ", s)
	fmt.Print("\n")

	var grades int = 86
	//to print the value of the string within a statement using %v
	fmt.Printf("Marks: %d ", grades)
	fmt.Print("\n")

	var name string = "Jonathan"
	var u int = 87
	fmt.Printf("Hey MASTER %v your MARKS are %d ", name, u)
}
