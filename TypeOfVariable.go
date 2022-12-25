package main

import "fmt"

func main() {
	var grades int = 25
	var message string = "hello world"
	var isCheck bool = true
	var amount float64 = 5632.89

	//%v is to print value of variable
	//%T is used to print the datatype of the variable
	fmt.Printf("variable grades = %v is of type %T\n", grades, grades)
	fmt.Printf("variable messgae = %v is of type %T\n", message, message)
	fmt.Printf("variable isCheck = %v is of type %T\n", isCheck, isCheck)
	fmt.Printf("variable amount = %v is of type %T\n", amount, amount)
}
