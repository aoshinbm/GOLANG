package main

import "fmt"

func main() {
	var y int
	var ptr *int = &y

	*ptr = 0
	fmt.Println(y)

	*ptr += 5
	fmt.Println(y)

}
