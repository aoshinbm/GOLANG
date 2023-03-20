package main

import "fmt"

func main() {
	s := 100
	var ptr *string = &s
	fmt.Println(s)
	*ptr += 100
	fmt.Println(s)
}
