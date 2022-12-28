package main

import "fmt"

func test() func(x, y int) int {
	f := func(x, y int) int {
		return x + y
	}
	return f
}

func main() {
	t := test()
	fmt.Println(t(27, 7))
}
