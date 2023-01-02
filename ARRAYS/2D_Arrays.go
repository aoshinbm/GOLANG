package main

import "fmt"

func main() {

	//It's like an array of arrays or basically an array that has multiple levels
	//The simplest multi-dimensional array is the 2-d array or two-dimensional array
	arr := [3][2]int{{2, 4}, {5, 6}, {8, 9}}
	fmt.Println(arr[2][1])
}
