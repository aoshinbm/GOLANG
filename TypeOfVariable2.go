package main

import (
	"fmt"
	"reflect"
)

// fmt is used for formatting n printing output
// reflect package is used for the type of function
func main() {
	fmt.Printf("Type : %v \n", reflect.TypeOf(10000))
	fmt.Printf("Type : %v \n", reflect.TypeOf("Vamos Argentina"))
	fmt.Printf("Type : %v \n", reflect.TypeOf(88.36))
	fmt.Printf("Type : %v \n", reflect.TypeOf(true))

	var grades int = 25
	var message string = "hello world"
	//%v is to print value of variable
	//%T is used to print the datatype of the variable
	fmt.Printf("variable grades = %v is of type %v\n", grades, reflect.TypeOf(grades))
	fmt.Printf("variable messgae = %v is of type %v\n", message, reflect.TypeOf(message))

}
