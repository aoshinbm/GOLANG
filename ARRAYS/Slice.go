package main

import "fmt"

func main() {
	//a slice is declared and initialized with values
	slice := []int{10, 30, 50}
	fmt.Println(slice)

	//create an array of 10 elements
	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// then slice it from index one to index eight.
	//This means we want to include elements from index one to index seven
	//and storing the slice in a variable called slice1
	slice1 := arr[1:8]
	fmt.Println(slice1)

	//declare and initialize a slice using another slice.
	sub_slice := slice1[2:4]
	fmt.Println(sub_slice)

	//create a slice using the make function.
	sliceArray := make([]int, 5, 8) //OUTPUT:= [0 0 0 0 0] as its empty slice
	//print the slice
	fmt.Println(sliceArray)
	//print the length of the slice
	fmt.Println(len(sliceArray))
	//print capacity of the slice
	fmt.Println(cap(sliceArray))

	//capacity function is mostly used with slice.
	//but it works for arrays as well
	//The only difference is that for arrays, capacity is equal to the length of the array.
	fmt.Println(arr)
	fmt.Println(cap(arr))
	fmt.Println(cap(slice1))

	arrayIntegers := [5]int{11, 22, 33, 44, 55}
	fmt.Println(arrayIntegers)
	sliceOfArrayIntegers := arrayIntegers[:3]
	fmt.Println(sliceOfArrayIntegers)

	sliceOfArrayIntegers[1] = 99
	fmt.Println("After modification")
	fmt.Println(sliceOfArrayIntegers)
	fmt.Println(arrayIntegers)

	arraa := [5]int{10, 20, 90, 70, 60}
	sliceee := arraa[:3]
	fmt.Println(cap(sliceee))

	slice_2 := make([]int, 5, 20)
	new_slice := append(slice, slice_2...)
	fmt.Println(cap(new_slice))
}
