package main

import "fmt"

func modify(s string) {
	s = "folkzz" //here s parameter is copied into another location of memory
}
func main() {
	a := "hello"
	fmt.Println(a)
	modify(a)
	//wen s value is changed original value didnt get affected
	fmt.Println(a)
}
