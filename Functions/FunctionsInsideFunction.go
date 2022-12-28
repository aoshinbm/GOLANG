package main

import "fmt"

func main() {

	test := func() {
		//test is tha variable storing function that has no parameters
		//no arguments and doesnt return anything n simply prints hello
		fmt.Println("Helloo..!!")
	}
	test()

	test1 := func(x int) {
		//test1 is tha variable storing function
		fmt.Println(x)
	}
	test1(259)

	mul :=

		func(x int) int {
			//return value
			return x * 3
		}(12) //call the function directly where it is defined
	fmt.Println(mul) //here instead of storing the function it stored the value returned from function
}
