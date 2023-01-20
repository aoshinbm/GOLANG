package main

import "fmt"

//eventhough havent mentioned any pointer the value changes
func modify(sl []int) {
	sl[0] = 120
}

//this happened bcuz slice itself is an reference to an underlying array
func main() {
	slice := []int{10, 40, 70}
	fmt.Println(slice)
	modify(slice)
	//as slice is passed to function the pointer of slice will still refer to
	//the same underlaying array
	//any change through slice done in function is reflected even after the function scope
	fmt.Println(slice)
}
