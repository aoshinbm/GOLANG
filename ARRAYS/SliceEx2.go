package main

import (
	"fmt"
)

func main() {
	/*The make is a built-in function which allocates and initializes an object of type slice, map, or channel.
	func make(t Type, size ...IntegerType) Type
	In case of a slice, the size specifies its length.
	A second integer may be specified to set the capacity of the slice;
	it must not be smaller than the length.
	The make([]int, 0, 5) allocates an underlying array of size 5 and
	returns a slice of length 0 and
	capacity 5 that is backed by this underlying array.*/

	//create slice .
	//We create a new slice with five elements initialized to 0.
	vals := make([]int, 5)
	fmt.Println(vals)

	//fill the slice with some data.
	vals[0] = 12
	vals[1] = 18
	vals[2] = 13
	vals[3] = 19
	vals[4] = 38
	fmt.Println(vals)

	//initialize the slice when we define it with literal notation.
	vals2 := []int{12, 18, 13, 19, 38}
	fmt.Println(vals2)

	//Here we create an empty slice.
	vals3 := []int{}
	fmt.Println(vals3)

	//It is not possible to assign values to an empty slice.
	// vals3[0] = 12
	// vals3[1] = 18

	//To add new values, the append function can be uzed.
	vals3 = append(vals3, 1)
	vals3 = append(vals3, 2)
	vals3 = append(vals3, 5)
	vals3 = append(vals3, 6)

	fmt.Println(vals3)
}
