package main

import "fmt"

func main() {
	//declaring an anonymous function
	//NOTE:= function doesnt have a name
	x := func(l int, b int) int {
		return l * b
	}
	fmt.Printf("%T \n", x)
	fmt.Println(x(30, 50)) //v r calling the function

	y := func(l int, b int) int {
		return l * b
	}(50, 60)
	fmt.Printf("%T \n", y) //y stores the value that is returned by the function i.e y=3000 &  y is int
	fmt.Println(y)         //v r calling the function

}
