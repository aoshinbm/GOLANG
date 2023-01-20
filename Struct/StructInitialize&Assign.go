package main

import "fmt"

type Student struct {
	name   string
	rollNo int
}

func main() {
	st := Student{
		name:   "Annette",
		rollNo: 22,
	}
	//to print struct field names and values special format specifier is used i.e. +v
	fmt.Printf("%+v", st)
}
