package main

import "fmt"

// SimpleFunction prints a message
//It takes no parameter and returns no values
func SimpleFunction() {
	fmt.Println("Hello World")
}

func printGreetings(strr string) {
	fmt.Println("Hello there ,", strr)
}
func addNumbers(a int, b int) int {
	return a + b
}

func subNumbers(a int, b int) int {
	subt := a - b
	return subt
}

func main() {
	SimpleFunction()
	printGreetings("Monica")
	result := addNumbers(20, 90)
	fmt.Println(result)
	subtraction := subNumbers(90, 30)
	fmt.Println(subtraction)
}
