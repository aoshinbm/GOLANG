package main

import "fmt"

func main() {
	var y int
	var ptr *int = &y
	fmt.Println(y)
	fmt.Println(*ptr)
}
