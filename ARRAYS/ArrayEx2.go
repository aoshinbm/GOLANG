package main

import "fmt"

func main() {

	//only declared the array of integers
	var grades [5]int
	fmt.Println(grades) //OUTPUT:- [0 0 0 0 0] bcuz havent initialized yet,
	//zero bcuz datatype is integer

	var fruits [3]string
	fmt.Println(fruits) //OUTPUT:- [  ] empty string

	//declare an integer array intArr of size 3,
	//and then assign values to the elements of this integer array
	var intArr [3]int
	intArr[0] = 100
	intArr[1] = 0020
	intArr[2] = 310
	fmt.Println(intArr)

	//Integer Array which is declared and initialized in a single line
	arr := [3]int{10, 20, 30}
	fmt.Println(arr)

	//declare and initialize the array of strings
	var Fruits [2]string = [2]string{"apples", "pineapples"}
	fmt.Println(Fruits) //OUTPUT:- [  ] empty string
	fmt.Println(Fruits[0])

	//initializing array using short hand operator
	marks := [6]int{25, 65, 75, 22, 45, 59}
	fmt.Println(marks)
	//change value of an element in array at specified index
	marks[2] = 100
	fmt.Println(marks)
	//loopin through an array
	for i := 0; i < len(marks); i++ {
		fmt.Println(marks[i])
	}

	name := [...]string{"Kathleen", "Dianna", "Meghan", "Lillibet"}
	fmt.Println(name)
	fmt.Println("Length of name array :", len(name))
	fmt.Println(name[3])
	//looping through an array using range keyword
	//The range keyword is mainly used in for loops, in order to iterate over all the elements
	//of an array,slice or map
	for index, element := range name {
		fmt.Println(index, "=>", element)
	}

	arr11 := [5]bool{true, true, true, true}
	for i := 0; i < len(arr11); i++ {
		if arr11[i] {
			fmt.Println(i)
		}
	}

	arrayy := [5]bool{true, true}
	fmt.Println(arrayy)
}
