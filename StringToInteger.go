package main

import (
	"fmt"
	"strconv"
)

// first import string conversion package

func main() {
	//convert integer into string
	fmt.Println("Converting int into string :")
	var strr string = "200"
	i, err := strconv.Atoi(strr)
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", err, err)
	fmt.Println("----------------------------------------")

	var str string = "200Bucks" //Atoi() function will return as an error for this conversion
	a, err := strconv.Atoi(str)
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", err, err)
	//The error data type is a custom data type in the string conversion package.
	//error message ==>> invalid syntax because we cannot convert 200Bucks to a pure integer.

}
