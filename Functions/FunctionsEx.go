package main

import "fmt"

func test() {
	fmt.Println("HElOOO .......")
}
func test1(x int) {
	fmt.Println("HElOOO .......", x)
}
func main() {
	//calling a normal function
	test()
	/*
		test() ==> () this parentheses signifies call i.e. call the function
		test ==> here it says refernce the function where actually it doesnt call the function so
	*/
	x := test //here function is assigned it to x variable
	x()       //here it is calling the function same like test() function

	//both the above will work
	//this way v can store functions
	y := test1
	y(55) //passing parameters

	//we can assign any function to variable
	//we can create functions anywhere we want
	//we can return function
	//we can pass function in as parameters
}
