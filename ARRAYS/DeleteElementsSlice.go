package main

import "fmt"

func main() {
	arr := [5]int{14, 24, 34, 44, 54}
	fmt.Println(arr)

	i := 2

	slice_1 := arr[:i]
	fmt.Println(slice_1)
	slice_2 := arr[i+1:]
	fmt.Println(slice_2)

	SSlice1 := append(slice_1, slice_2...)
	fmt.Println(SSlice1)
}
