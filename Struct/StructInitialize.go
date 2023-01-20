package main

import "fmt"

type Student struct {
	name   string
	rollNo int
	marks  []int          //slice
	grades map[string]int //map
}

func main() {
	st := new(Student)
	//to print struct field names and values special format specifier is used i.e. +v
	fmt.Printf("%+v", st)
}

/*
OUTPUT:=
&{name: rollNo:0 marks:[] grades:map[]}

name = by default empty string
rollNo = by default zero
marks = by default empty slice
grades = by default map is empty

*/
