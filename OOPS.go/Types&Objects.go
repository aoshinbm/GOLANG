package main

import "fmt"

type myInt int

func main() {
	value := myInt(4)
	fmt.Println(value)
	fmt.Println(double_test(value))

}

func double_test(mi myInt) int {
	return int(mi * 10)
}
