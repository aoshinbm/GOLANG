package main

import "fmt"

func main() {
	s := "heloo"
	var b_ptr *string = &s
	fmt.Println(s)
	fmt.Println(b_ptr)

	var a_ptr = &s
	fmt.Println(a_ptr)

	c_ptr := &s
	fmt.Println(c_ptr)
}

/*
OUTPUT:=
we get the same addresses
and this is what pointers are
that store the memory address of another variable

heloo
0xc000102230
0xc000102230
0xc000102230
*/
