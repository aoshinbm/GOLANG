package main

import "fmt"

func test(myFunc func(int) int) { //takes function as parameter
	//myFunc 	  func(int) 				int
	//parameter   funtc(parameter type)     return type

	fmt.Println(myFunc(77))
}

func main() {

	mul :=

		func(x int) int {
			//return value
			return x * 3
		}

	test(mul)

	test3 := func(y int) int {
		return y * 9
	}
	test(test3)

	//can directly write function and call it immediatey
	func(z int) int {
		fmt.Println("Multiply")
		return z * 10
	}(25)
	//	fmt.Println((25))

}
