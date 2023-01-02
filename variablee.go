package main

import (
	"fmt"
	"reflect"
)

func main() {
	var greeting string = " Hello World"
	fmt.Println(greeting)

	var i int = 1
	var w float64 = 12.5
	fmt.Println(i, w)

	var x int = 10
	fmt.Printf("x is initialized with %d and its type is %T \n", x, x)

	var floatnumber = 10.00
	fmt.Printf("floatnumber is initialized with %.2f and its type is %T \n", floatnumber, floatnumber)

	var z float64
	z = float64(x) * floatnumber
	fmt.Printf("Product is %.2f and type is %T \n", z, z)

	var comp

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

	//in order for the inference to work, the variables must be initialized
	var name = "John Doe"
	var age = 34

	//With the help of the TypeOf function from the reflect package,
	//we print the data types of the two variables.
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(age))
	fmt.Printf("%s is %d years old\n", name, age)

	//declares two variables with the shorhand notation
	Name := "John Doe"
	Age := 34
	fmt.Printf("%s is %d years old", Name, Age)

}
