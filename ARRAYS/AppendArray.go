package main

import "fmt"

func main() {
	arr := [4]int{14, 24, 34, 44}
	fmt.Println(arr)
	slice := arr[1:3]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 54, 64, 74)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	arr2 := [5]int{21, 31, 41, 51, 61}
	fmt.Println(arr2)
	sliceOfArr2 := arr2[:3]
	fmt.Println(sliceOfArr2)

	new_slice := append(slice, sliceOfArr2...)
	fmt.Println(new_slice)

}
