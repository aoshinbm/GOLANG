package main

import "fmt"

func main() {
	//	x := 10
	addX := func(y int) int {
		return x + y
	}
	result := addX(20)
	fmt.Println(result) // Output: 30

	/*	 func add(y int)int  {
		return x+y
	 }
	*/
}
