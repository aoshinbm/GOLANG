package main

import "fmt"

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	myByteSlice := []byte(myArray)
	fmt.Println(myByteSlice)
}
