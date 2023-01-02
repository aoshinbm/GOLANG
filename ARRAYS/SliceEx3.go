package main

import "fmt"

func main() {

	// creates a slice of a string slices of various length
	//create a slice of slices
	w := make([][]string, 3)

	//subslices is separately created with make and initialized
	//subslice_1 created and values initialized
	w1 := make([]string, 4)
	w1[0] = "war"
	w1[1] = "water"
	w1[2] = "wrath"
	w1[3] = "wrong"

	//subslice_2 created and values initialized
	w2 := make([]string, 3)
	w2[0] = "car"
	w2[1] = "cup"
	w2[2] = "cloud"

	//subslice_3 created and values initialized
	w3 := make([]string, 2)
	w3[0] = "boy"
	w3[1] = "brown"

	//all the subslices are stored in the main slice
	w[0] = w1
	w[1] = w2
	w[2] = w3

	//n wen the main slice is called all the values of sublices stored in it will be presented
	fmt.Println(w)
}
