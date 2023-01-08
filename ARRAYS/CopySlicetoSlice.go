package main

import "fmt"

func main() {
	src_Slice := [6]int{14, 24, 34, 44, 54, 64}
	fmt.Println(src_Slice)

	dest_Slice := make([]int, 3)

	num := copy(dest_Slice, src_Slice[:])
	fmt.Println(dest_Slice)
	fmt.Println("No of elements copied : ", num)

	/*OUTPUT:-
	[14 24 34 44 54 64]
	[14 24 34]
	No of elements copied :  3

	Since the length of the destination slice was three,
	and hence it also got the capacity as three,
	we see that only three numbers were copied into the destination slice.*/

	//looping through slice
	for index, element := range src_Slice {
		fmt.Println(index, "===>", element)
	}
	//Now for use cases where you don't need the index or value, simply replace it with the underscore.
	for _, element := range src_Slice {
		fmt.Println("===>", element)
	}
}
