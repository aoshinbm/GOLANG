package main

import "fmt"

/*
A function is a mapping of zero or more input parameters to zero or more output parameters.
Go functions are first-class citizens.
Functions can be assigned to variables,
passed as arguments to functions or returned from functions.

Functions in Go are created with the func keyword.
We use the return keyword to return values from functions.
The body of the function consists of statements that are executed when the function is called.
The body is delimited with a pair of curly brackets {}.
To call a function, we specify its name followed by round bracktets ().
A function may or may not take parameters.*/

func main() {

	x := 25
	y := 50

	//call the add function; it takes two parameters.
	//The computed value is passed to the z variable.
	z := add(x, y)

	fmt.Printf("Output: %d\n", z)
}

//define a function which adds two values
func add(a int, b int) int {

	//result of the addition operation is returned to the caller with the return keyword
	return a + b
}
