package main

import "fmt"

func main() {
	var greeting string = " Hello World"
	fmt.Println(greeting)

	var x int = 10
	fmt.Printf("x is initialized with %d and its type is %T \n", x, x)

	var floatnumber = 10.00
	fmt.Printf("floatnumber is initialized with %.2f and its type is %T \n", floatnumber, floatnumber)

	//Short variable declaration
	//where values can be changed
	// := is short variable assignment operator
	fname := "LUCIAN"
	fmt.Println(fname)
	fname = "ROXANNE"
	fmt.Println(fname)

	lname := "Farewell" //here its assigning n declaring variable with value & datatype
	fmt.Println(lname)
	//once assigned with specific datatype then u cant change its type
	/*
		lname = 12345 //cannot use 12345 (untyped int constant) as string value in assignment
		fmt.Println(lname)
	*/
}
