package main

import "fmt"

func main() {
	s := "aoshin"
	fmt.Printf("%v %T\n", s, s)

	ps := &s
	fmt.Println("Address of variable s:=", ps)
	//modify value of variable using dereference operator with the pointer
	*ps = "changetheWORLD"
	fmt.Printf("%v %T\n", s, s)

}
