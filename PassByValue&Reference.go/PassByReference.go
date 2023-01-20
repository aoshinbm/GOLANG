package main

import "fmt"

func modify(s *string) {
	*s = "folkzz" //here s parameter is copied into same location of memory
}
func main() {
	a := "hello"
	fmt.Println(a)
	modify(&a)
	//wen s value is changed original value  gets affected
	fmt.Println(a)
}
